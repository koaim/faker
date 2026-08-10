package faker

import (
	"math/rand/v2"
	"reflect"
	"slices"
)

// Make generates and returns a value of type T filled with fake data.
//
// It starts from the default options and applies the provided opts in
// order. Supported types: int/uint, float/complex, string, bool, struct,
// slice, array, map, chan, and pointer. Pointers are never nil. Struct
// fields listed in IgnoreFields and unaddressable fields are skipped.
func Make[T any](opts ...OptionF) T {
	o := DefaultOption()
	for _, v := range opts {
		v(&o)
	}

	var v T
	fill(&v, o)

	return v
}

// MakeAndOverride generates a value of type T using the default rules,
// then applies the override f to it. The override receives a pointer to
// the generated value and can mutate any part of it before it is
// returned.
func MakeAndOverride[T any](f func(v *T)) T {
	v := Make[T]()
	f(&v)

	return v
}

// MakeWithOption generates a value of type T using a single preconfigured
// Option, exactly as given. Unlike Make, it does not start from the
// defaults and applies no overrides.
func MakeWithOption[T any](opt Option) T {
	var v T
	fill(&v, opt)

	return v
}

func fill(v any, opt Option) {
	rv := reflect.ValueOf(v).Elem()

	switch rv.Kind() {
	case reflect.String:
		fillString(rv, opt)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		rv.SetInt(rand.Int64())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		rv.SetUint(rand.Uint64())
	case reflect.Struct:
		fillStruct(rv, opt)
	case reflect.Bool:
		rv.SetBool(true)
	case reflect.Float64, reflect.Float32:
		rv.SetFloat(rand.NormFloat64())
	case reflect.Complex64, reflect.Complex128:
		rv.SetComplex(complex(rand.NormFloat64(), rand.NormFloat64()))
	case reflect.Slice:
		fillSlice(rv, opt)
	case reflect.Array:
		fillArray(rv, opt)
	case reflect.Map:
		fillMap(rv, opt)
	case reflect.Chan:
		fillChan(rv, opt)
	case reflect.Pointer:
		fillPointer(rv, opt)
	}
}

func fillArray(rv reflect.Value, opt Option) {
	for i := 0; i < rv.Cap(); i++ {
		fill(rv.Index(i).Addr().Interface(), opt)
	}
}

func fillSlice(rv reflect.Value, opt Option) {
	length := randLen(opt)
	rv.Set(reflect.MakeSlice(rv.Type(), length, length))

	for i := 0; i < rv.Len(); i++ {
		fill(rv.Index(i).Addr().Interface(), opt)
	}
}

func fillStruct(rv reflect.Value, opt Option) {
	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)
		sf := rv.Type().Field(i)

		if !field.CanSet() || slices.Contains(opt.IgnoreFields, sf.Name) {
			continue
		}

		fill(field.Addr().Interface(), opt)
	}
}

func fillMap(rv reflect.Value, opt Option) {
	length := randLen(opt)
	rv.Set(reflect.MakeMapWithSize(rv.Type(), length))

	kType := rv.Type().Key()
	vType := rv.Type().Elem()

	for range length {
		k := reflect.New(kType).Elem()
		v := reflect.New(vType).Elem()

		fill(k.Addr().Interface(), opt)
		fill(v.Addr().Interface(), opt)

		rv.SetMapIndex(k, v)
	}
}

func fillChan(rv reflect.Value, opt Option) {
	if opt.CloseChannels {
		defer rv.Close()
	}

	length := randLen(opt)
	rv.Set(reflect.MakeChan(rv.Type(), length))

	elemType := rv.Type().Elem()

	for range length {
		elem := reflect.New(elemType).Elem()

		fill(elem.Addr().Interface(), opt)

		rv.Send(elem)
	}
}

func fillPointer(rv reflect.Value, opt Option) {
	ptr := reflect.New(rv.Type().Elem())
	fill(ptr.Interface(), opt)

	rv.Set(ptr)
}

func fillString(rv reflect.Value, opt Option) {
	r := make([]rune, opt.StrLen)

	for i := range len(r) {
		index := rand.IntN(len(opt.AllowedRunes))
		r[i] = opt.AllowedRunes[index]
	}

	rv.SetString(string(r))
}

func randLen(opt Option) int {
	if opt.MaxLen < opt.MinLen {
		return opt.MinLen
	}
	return opt.MinLen + rand.IntN(opt.MaxLen-opt.MinLen+1)
}
