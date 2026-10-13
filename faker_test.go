package faker

import (
	"context"
	"net"
	"net/url"
	"slices"
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
	_, err := uuid.Parse(v.String())
	if err != nil {
		t.Fatal(err)
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
}

func TestMake_HardwareAddr(t *testing.T) {
	t.Parallel()

	v := Make[net.HardwareAddr](t)

	_, err := net.ParseMAC(v.String())

	if len(v) == 0 {
		t.Fatal("expected non-zero net.HardwareAddr")
	}
	if err != nil {
		t.Fatal(err)
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
		t.Fatal("expected non-zero float46")
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
		t.Fatal("expected non-zero User value")
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

	u := Make[User](t)

	if u.Ctx != nil {
		t.Fatalf("expected nil u.Ctx, got %v", u.Ctx)
	}
	if u.Any != nil {
		t.Fatalf("expected nil u.Any, got %v", u.Any)
	}
	if u.F != nil {
		t.Fatalf("expected nil u.F")
	}
	if u.private != 0 {
		t.Fatal("expected zero u.private")
	}
}

func TestMake_Struct_CustomFakers(t *testing.T) {
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

	u := Make[User](t,
		WithCustomFaker(func() string { return strVal }),
		WithCustomFaker(func() int { return intVal }),
	)

	if u.ID != intVal {
		t.Fatalf("expected ID = %v, actual = %v", intVal, u.ID)
	}
	if *u.Age != intVal {
		t.Fatalf("expected Age = %v, actual = %v", intVal, u.Age)
	}
	if u.Name != strVal {
		t.Fatalf("expected Age = %v, actual = %v", intVal, u.Name)
	}
	if *u.City != strVal {
		t.Fatalf("expected Age = %v, actual = %v", intVal, u.City)
	}
}

func TestMakeAndOverride(t *testing.T) {
	t.Parallel()

	const (
		customName = "custom name"
		customID   = 10
	)

	type User struct {
		ID   int
		Name string
	}

	withCustomName := func(u *User) { u.Name = customName }
	withCustomID := func(u *User) { u.ID = customID }

	defaultUser := User{}

	u := MakeAndOverride(t, withCustomName, withCustomID)

	if u == defaultUser {
		t.Fatal("expected not default value")
	}
	if u.Name != customName {
		t.Fatalf("expected Name = %v, actual = %v", customName, u.Name)
	}
	if u.ID != customID {
		t.Fatalf("expected ID = %v, actual = %v", customID, u.ID)
	}
}

func TestMakeWithOption(t *testing.T) {
	t.Parallel()

	type User struct {
		Numbers []int
		Name    string
		Ch      chan int
		Skip1   string
		Skip2   int
	}

	t.Run("close channels", func(t *testing.T) {
		t.Parallel()

		opt := DefaultOption()
		opt.CloseChannels = true

		u := MakeWithOption[User](t, opt)

		for range u.Ch {
		}

		if _, ok := <-u.Ch; ok {
			t.Fatal("expected channel to be closed")
		}
	})

	t.Run("allowed runes", func(t *testing.T) {
		t.Parallel()

		opt := DefaultOption()
		opt.AllowedRunes = []rune("123abc")

		u := MakeWithOption[User](t, opt)

		for _, v := range u.Name {
			if !slices.Contains(opt.AllowedRunes, v) {
				t.Fatalf("symbol %d not exists in allowed runes '%v'", v, string(opt.AllowedRunes))
			}
		}
	})

	t.Run("str len", func(t *testing.T) {
		t.Parallel()

		opt := DefaultOption()
		opt.StrLen = 3

		u := MakeWithOption[User](t, opt)

		if utf8.RuneCountInString(u.Name) != opt.StrLen {
			t.Fatalf("expected string length %v, got %v", opt.StrLen, utf8.RuneCountInString(u.Name))
		}
	})

	t.Run("containers length", func(t *testing.T) {
		t.Parallel()

		opt := DefaultOption()
		opt.MinContainersLen = 10
		opt.MaxContainersLen = 15

		u := MakeWithOption[User](t, opt)
		if len(u.Numbers) < opt.MinContainersLen || len(u.Numbers) > opt.MaxContainersLen {
			t.Fatalf("expected slice length in range [%d, %d], got %d", DefaultOption().MinContainersLen, DefaultOption().MaxContainersLen, len(u.Numbers))
		}
	})

	t.Run("skip fields", func(t *testing.T) {
		t.Parallel()

		opt := DefaultOption()
		opt.IgnoreFields = []string{"Skip1", "Skip2"}

		u := MakeWithOption[User](t, opt)
		if u.Skip1 != "" {
			t.Fatalf("expected default string value, got %v", u.Skip1)
		}
		if u.Skip2 != 0 {
			t.Fatalf("expected default int value, got %v", u.Skip2)
		}
	})
}
