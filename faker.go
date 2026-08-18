// Package faker generates fake data for arbitrary Go types.
//
// It fills structs, slices, maps, and primitive values with random
// values using reflection. See the Make function for details.
package faker

import (
	"math"
	"math/rand/v2"
	"net"
	"net/url"
	"reflect"
	"slices"
	"time"
)

// Make generates and returns a value of type T filled with fake data.
//
// It starts from the default options and applies the provided opts in
// order. Supported types: int/uint, float/complex, string, bool, struct,
// slice, array, map, chan, time.Time, net.IP, url.URL and pointer. Pointers are
// never nil. Struct fields listed in IgnoreFields and unaddressable fields are
// skipped.
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
func MakeAndOverride[T any](funcs ...func(v *T)) T {
	v := Make[T]()
	for _, f := range funcs {
		f(&v)
	}

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
	rvType := rv.Type()
	rvKind := rv.Kind()

	for _, f := range opt.CustomFakers {
		if f.Type == rvType {
			rv.Set(reflect.ValueOf(f.F()))
			return
		}
	}

	switch rvType {
	case reflect.TypeFor[string]():
		fillString(rv, opt)
	case reflect.TypeFor[int](), reflect.TypeFor[int8](), reflect.TypeFor[int16](), reflect.TypeFor[int32](), reflect.TypeFor[int64]():
		fillInt(rv)
	case reflect.TypeFor[uint](), reflect.TypeFor[uint8](), reflect.TypeFor[uint16](), reflect.TypeFor[uint32](), reflect.TypeFor[uint64](), reflect.TypeFor[uintptr]():
		fillUint(rv)
	case reflect.TypeFor[time.Time]():
		fillTime(rv)
	case reflect.TypeFor[url.URL]():
		fillURL(rv, opt)
	case reflect.TypeFor[net.IPNet]():
		fillIPNet(rv)
	case reflect.TypeFor[net.IP]():
		fillIP(rv)
	case reflect.TypeFor[net.IPMask]():
		fillIPMask(rv)
	case reflect.TypeFor[net.HardwareAddr]():
		fillHardwareAddr(rv)
	case reflect.TypeFor[bool]():
		rv.SetBool(true)
	case reflect.TypeFor[float64](), reflect.TypeFor[float32]():
		rv.SetFloat(rand.NormFloat64())
	case reflect.TypeFor[complex64](), reflect.TypeFor[complex128]():
		rv.SetComplex(complex(rand.NormFloat64(), rand.NormFloat64()))
	default:
		switch rvKind {
		case reflect.Struct:
			fillStruct(rv, opt)
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
}

func fillUint(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Uint8:
		rv.SetUint(rand.Uint64N(math.MaxUint8))
	case reflect.Uint16:
		rv.SetUint(rand.Uint64N(math.MaxUint16))
	case reflect.Uint32:
		rv.SetUint(rand.Uint64N(math.MaxUint32))
	case reflect.Uint64:
		rv.SetUint(rand.Uint64())
	case reflect.Uint, reflect.Uintptr:
		rv.SetUint(uint64(rand.Uint()))
	}
}

func fillInt(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Int8:
		rv.SetInt(rand.Int64N(math.MaxInt8))
	case reflect.Int16:
		rv.SetInt(rand.Int64N(math.MaxInt16))
	case reflect.Int32:
		rv.SetInt(rand.Int64N(math.MaxInt32))
	case reflect.Int64:
		rv.SetInt(rand.Int64N(math.MaxInt64))
	case reflect.Int:
		rv.SetInt(rand.Int64N(math.MaxInt))
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
	rv.SetString(randStr(opt.AllowedRunes, opt.StrLen))
}
