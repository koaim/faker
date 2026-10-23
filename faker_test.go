package faker

import (
	"context"
	"net"
	"net/url"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"
	"unicode/utf8"
	"uuid"
)

func TestMake_UUID(t *testing.T) {
	t.Parallel()

	v := Make[uuid.UUID](t)

	defaultUUID := uuid.UUID{}

	if v == defaultUUID {
		t.Fatal("expected non-default uuid.UUID")
	}
	if _, err := uuid.Parse(v.String()); err != nil {
		t.Fatal(err)
	}
}

func TestMake_TypeWrap(t *testing.T) {
	t.Parallel()

	type MyInt int

	v := Make[MyInt](t,
		WithBuilder(func() int { return 10 }),
		WithBuilder(func() MyInt { return 20 }),
	)

	if v != 20 {
		t.Fatalf("expected 20, got %v", v)
	}
}

func TestMake_TypeAlias(t *testing.T) {
	t.Parallel()

	type MyInt = int

	v := Make[MyInt](t, WithBuilder(func() int { return 10 }))

	if v != 10 {
		t.Fatalf("expected 10, got %v", v)
	}
}

func TestMake_EmptyStruct(t *testing.T) {
	t.Parallel()

	type User struct{}

	defaultUser := User{}

	u := Make[User](t)

	if u != defaultUser {
		t.Fatal("expected zero User value")
	}
}

func TestMake_AnonymousStruct(t *testing.T) {
	t.Parallel()

	defaultValue := struct {
		ID   int
		Name string
	}{}

	u := Make[struct {
		ID   int
		Name string
	}](t)

	if u == defaultValue {
		t.Fatal("expected non-default value")
	}
}

func TestMake_Int(t *testing.T) {
	t.Parallel()

	v := Make[int](t)

	if v == 0 {
		t.Fatal("expected non-zero int")
	}
}

func TestMake_IPMask(t *testing.T) {
	t.Parallel()

	v := Make[net.IPMask](t)

	if len(v) == 0 {
		t.Fatal("expected non-zero net.IPMask")
	}
	if _, bits := v.Size(); bits != 32 {
		t.Fatalf("expected contiguous IPv4 mask, got %v", v)
	}
}

func TestMake_HardwareAddr(t *testing.T) {
	t.Parallel()

	v := Make[net.HardwareAddr](t)

	if _, err := net.ParseMAC(v.String()); err != nil {
		t.Fatal(err)
	}
	if len(v) == 0 {
		t.Fatal("expected non-zero net.HardwareAddr")
	}
}

func TestMake_IP(t *testing.T) {
	t.Parallel()

	v := Make[net.IP](t)

	ip := net.ParseIP(v.String())
	if ip == nil {
		t.Fatalf("expected valid net.IP, got %v", v.String())
	}
}

func TestMake_Byte(t *testing.T) {
	t.Parallel()

	v := Make[byte](t)

	if v == 0 {
		t.Fatal("expected non-zero byte")
	}
}

func TestMake_URL(t *testing.T) {
	t.Parallel()

	v := Make[url.URL](t)

	_, err := url.Parse(v.String())
	if err != nil {
		t.Fatal(err)
	}
}

func TestMake_URL_WithOptions(t *testing.T) {
	t.Parallel()

	var (
		domainZones = []string{"info"}
		schemes     = []string{"http"}
	)

	v := Make[url.URL](t, WithAllowedDomainZones(domainZones), WithAllowedUrlSchemes(schemes), WithUrlUserInfo(true))

	if !slices.Contains(schemes, v.Scheme) {
		t.Fatalf("invalid scheme %v", v.Scheme)
	}
	if v.User == nil || v.User.String() == "" {
		t.Fatal("expected non-nil and non-zero User")
	}
	if !strings.HasSuffix(v.Host, "info") {
		t.Fatalf("invalid domain zone %v", v.Host)
	}
}

func TestMake_Uint(t *testing.T) {
	t.Parallel()

	v := Make[uint](t)

	if v == 0 {
		t.Fatal("expected non-zero uint")
	}
}

func TestMake_Uintptr(t *testing.T) {
	t.Parallel()

	v := Make[uintptr](t)

	if v == 0 {
		t.Fatal("expected non-zero uintptr")
	}
}

func TestMake_Time(t *testing.T) {
	t.Parallel()

	v := Make[time.Time](t)

	if v.Equal(time.Time{}) {
		t.Fatal("expected non-zero time.Time")
	}
}

func TestMake_TimePointer(t *testing.T) {
	t.Parallel()

	v := Make[*time.Time](t)

	if v == nil {
		t.Fatal("expected non-nil value")
	}
	if v.Equal(time.Time{}) {
		t.Fatal("expected non-zero time.Time")
	}
}

func TestMake_Bool(t *testing.T) {
	t.Parallel()

	v := Make[bool](t)

	if v == false {
		t.Fatal("expected non-zero bool")
	}
}

func TestMake_Float64(t *testing.T) {
	t.Parallel()

	v := Make[float64](t)

	if v == 0 {
		t.Fatal("expected non-zero float64")
	}
}

func TestMake_Complex64(t *testing.T) {
	t.Parallel()

	v := Make[complex64](t)

	if v == 0 {
		t.Fatal("expected non-zero complex64")
	}
}

func TestMake_Slice(t *testing.T) {
	t.Parallel()

	v := Make[[]int](t)

	if len(v) < DefaultOption().MinContainersLen || len(v) > DefaultOption().MaxContainersLen {
		t.Fatalf("expected slice length in range [%d, %d], got %d", DefaultOption().MinContainersLen, DefaultOption().MaxContainersLen, len(v))
	}
	for _, v := range v {
		if v == 0 {
			t.Fatal("expected non-zero int")
		}
	}
}

func TestMake_Map(t *testing.T) {
	t.Parallel()

	v := Make[map[string]string](t)

	if len(v) < DefaultOption().MinContainersLen || len(v) > DefaultOption().MaxContainersLen {
		t.Fatalf("expected map length in range [%d, %d], got %d", DefaultOption().MinContainersLen, DefaultOption().MaxContainersLen, len(v))
	}
	for k, v := range v {
		if k == "" {
			t.Fatal("expected non-zero string key")
		}
		if v == "" {
			t.Fatal("expected non-zero string value")
		}
	}
}

func TestMake_Array(t *testing.T) {
	t.Parallel()

	const length = 4

	v := Make[[length]int](t)

	if len(v) != length {
		t.Fatalf("expected array length %d, got %d", length, len(v))
	}
	if cap(v) != length {
		t.Fatalf("expected array capacity %d, got %d", length, cap(v))
	}
	for _, v := range v {
		if v == 0 {
			t.Fatal("expected non-zero int")
		}
	}
}

func TestMake_String(t *testing.T) {
	t.Parallel()

	v := Make[string](t)

	if utf8.RuneCountInString(v) != DefaultOption().StrLen {
		t.Fatalf("expected string length %d, got %d", DefaultOption().StrLen, len(v))
	}
	for _, v := range v {
		if !slices.Contains(DefaultOption().AllowedRunes, v) {
			t.Fatalf("rune %q is not in the alphabet", v)
		}
	}
}

func TestMake_Chan(t *testing.T) {
	t.Parallel()

	v := Make[chan int](t)

	if len(v) == 0 {
		t.Fatal("expected non-empty chan")
	}
}

func TestMake_Error(t *testing.T) {
	t.Parallel()

	v := Make[error](t)

	if v == nil {
		t.Fatal("expected non-nil error")
	}
}

func TestMake_Struct(t *testing.T) {
	t.Parallel()

	type Order[T any] struct {
		ID        int64
		Name      string
		DeletedAt *time.Time
		Type      T
	}

	type User struct {
		ID        int
		Name      string
		Order     *Order[string]
		CreatedAt time.Time
	}

	defaultUser := User{}
	defaultOrder := Order[string]{}

	u := Make[User](t)

	if u == defaultUser {
		t.Fatal("expected non-default User value")
	}
	if u.Order == nil || *u.Order == defaultOrder {
		t.Fatal("expected non-nil and non-empty Order value")
	}
}

func TestMake_Struct_NotSupportedTypes(t *testing.T) {
	t.Parallel()

	type User struct {
		Ctx     context.Context
		Any     any
		F       func()
		private int
	}

	defaultUser := User{}

	u := Make[User](t)

	if !reflect.DeepEqual(u, defaultUser) {
		t.Fatalf("\nwant %+v\ngot  %+v", defaultUser, u)
	}
}

func TestMake_Struct_WithBuilders(t *testing.T) {
	t.Parallel()

	const (
		strVal = "string"
		intVal = 11
	)

	type User struct {
		ID   int
		Age  *int
		Name string
		City *string
	}

	want := User{intVal, new(intVal), strVal, new(strVal)}

	u := Make[User](t,
		WithBuilder(func() string { return strVal }),
		WithBuilder(func() int { return intVal }),
	)

	if !reflect.DeepEqual(u, want) {
		t.Fatalf("\nwant %+v\ngot  %+v", want, u)
	}
}

func TestMakeWithOverrides(t *testing.T) {
	t.Parallel()

	const (
		customName = "custom name"
		customID   = 10
	)

	type User struct {
		ID   int
		Name string
	}

	want := User{customID, customName}

	withCustomName := func(u *User) { u.Name = customName }
	withCustomID := func(u *User) { u.ID = customID }

	u := MakeWithOverrides(withCustomName, withCustomID)

	if u != want {
		t.Fatalf("\nwant %+v\ngot  %+v", want, u)
	}
}

func TestMakeWithOption_CloseChannels(t *testing.T) {
	t.Parallel()

	opt := DefaultOption()
	opt.CloseChannels = true

	v := MakeWithOption[chan int](t, opt)

	for range v {
	}

	if _, ok := <-v; ok {
		t.Fatal("expected channel to be closed")
	}
}

func TestMakeWithOption_AllowedRunes(t *testing.T) {
	t.Parallel()

	opt := DefaultOption()
	opt.AllowedRunes = []rune("123abc")

	v := MakeWithOption[string](t, opt)

	for _, v := range v {
		if !slices.Contains(opt.AllowedRunes, v) {
			t.Fatalf("symbol %d not exists in allowed runes '%v'", v, string(opt.AllowedRunes))
		}
	}
}

func TestMakeWithOption_StrLen(t *testing.T) {
	t.Parallel()

	opt := DefaultOption()
	opt.StrLen = 3

	v := MakeWithOption[string](t, opt)

	if utf8.RuneCountInString(v) != opt.StrLen {
		t.Fatalf("expected string length %v, got %v", opt.StrLen, utf8.RuneCountInString(v))
	}
}

func TestMakeWithOption_ContainersLen(t *testing.T) {
	t.Parallel()

	opt := DefaultOption()
	opt.MinContainersLen = 10
	opt.MaxContainersLen = 15

	v := MakeWithOption[[]int](t, opt)
	if len(v) < opt.MinContainersLen || len(v) > opt.MaxContainersLen {
		t.Fatalf("expected slice length in range [%d, %d], got %d", opt.MinContainersLen, opt.MaxContainersLen, len(v))
	}
}

func TestMakeWithOption_IgnoreFields(t *testing.T) {
	t.Parallel()

	type Struct struct {
		Field1 string
		Field2 int
	}

	opt := DefaultOption()
	opt.IgnoreFields = []string{"Field1", "Field2"}

	defaultStruct := Struct{}

	v := MakeWithOption[Struct](t, opt)

	if v != defaultStruct {
		t.Fatalf("expected default Struct, got %+v", v)
	}
}
