package faker

import (
	"context"
	"net/url"
	"slices"
	"testing"
	"time"
	"unicode/utf8"
)

func TestMake_Int(t *testing.T) {
	t.Parallel()

	v := Make[int]()
	if v == 0 {
		t.Fatal("expected non-zero int")
	}
}

func TestMake_URL(t *testing.T) {
	t.Parallel()

	v := Make[url.URL]()

	defaultUrl := url.URL{}
	if v == defaultUrl {
		t.Fatal("expected non-zero url.URL")
	}

	_, err := url.Parse(v.String())
	if err != nil {
		t.Fatal(err)
	}
}

func TestMake_Uint(t *testing.T) {
	t.Parallel()

	v := Make[uint]()
	if v == 0 {
		t.Fatal("expected non-zero uint")
	}
}

func TestMake_Uintptr(t *testing.T) {
	t.Parallel()

	v := Make[uintptr]()
	if v == 0 {
		t.Fatal("expected non-zero uintptr")
	}
}

func TestMake_Time(t *testing.T) {
	t.Parallel()

	v := Make[time.Time]()
	if v.Equal(time.Time{}) {
		t.Fatal("expected non-zero time.Time")
	}
}

func TestMake_TimePointer(t *testing.T) {
	t.Parallel()

	v := Make[*time.Time]()
	if v == nil {
		t.Fatal("expected non-nil value")
	}
	if v.Equal(time.Time{}) {
		t.Fatal("expected non-zero time.Time")
	}
}

func TestMake_Bool(t *testing.T) {
	t.Parallel()

	v := Make[bool]()
	if v == false {
		t.Fatal("expected non-zero bool")
	}
}

func TestMake_Float64(t *testing.T) {
	t.Parallel()

	v := Make[float64]()
	if v == 0 {
		t.Fatal("expected non-zero float46")
	}
}

func TestMake_Complex64(t *testing.T) {
	t.Parallel()

	v := Make[complex64]()
	if v == 0 {
		t.Fatal("expected non-zero complex64")
	}
}

func TestMake_Slice(t *testing.T) {
	t.Parallel()

	v := Make[[]int]()
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

	v := Make[map[string]string]()
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

	v := Make[[length]int]()
	if len(v) != length {
		t.Fatalf("expected array length %d, got %d", length, len(v))
	}
	if cap(v) != length {
		t.Fatalf("expected array length %d, got %d", length, cap(v))
	}
}

func TestMake_String(t *testing.T) {
	t.Parallel()

	v := Make[string]()
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

	v := Make[chan int]()
	if len(v) == 0 {
		t.Fatal("expected non-empty chan")
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

	u := Make[User]()
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
		Ctx context.Context
		Any any
		Err error
		F   func()
	}

	u := Make[User]()

	if u.Ctx != nil {
		t.Fatalf("expected nil Ctx, got %v", u.Ctx)
	}
	if u.Any != nil {
		t.Fatalf("expected nil Any, got %v", u.Any)
	}
	if u.Err != nil {
		t.Fatalf("expected nil Err, got %v", u.Err)
	}
	if u.F != nil {
		t.Fatalf("expected nil F")
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
		Age  int
		Name string
		City string
	}

	u := Make[User](
		WithCustomFaker(func() string {
			return strVal
		}),
		WithCustomFaker(func() int {
			return intVal
		}),
	)

	if u.ID != intVal {
		t.Fatalf("expected ID = %v, actual = %v", intVal, u.ID)
	}
	if u.Age != intVal {
		t.Fatalf("expected Age = %v, actual = %v", intVal, u.Age)
	}
	if u.Name != strVal {
		t.Fatalf("expected Age = %v, actual = %v", intVal, u.Name)
	}
	if u.City != strVal {
		t.Fatalf("expected Age = %v, actual = %v", intVal, u.City)
	}
}

func TestMakeAndOverride(t *testing.T) {
	t.Parallel()

	const customName = "custom name"

	type User struct {
		ID   int
		Name string
	}

	defaultUser := User{}

	u := MakeAndOverride(func(v *User) {
		v.Name = customName
	})

	if u == defaultUser {
		t.Fatal("expected not default value")
	}
	if u.Name != customName {
		t.Fatalf("expected Name = %v, actual = %v", customName, u.Name)
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

		u := MakeWithOption[User](opt)

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

		u := MakeWithOption[User](opt)

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

		u := MakeWithOption[User](opt)

		if utf8.RuneCountInString(u.Name) != opt.StrLen {
			t.Fatalf("expected string length %v, got %v", opt.StrLen, utf8.RuneCountInString(u.Name))
		}
	})

	t.Run("containers length", func(t *testing.T) {
		t.Parallel()

		opt := DefaultOption()
		opt.MinContainersLen = 10
		opt.MaxContainersLen = 15

		u := MakeWithOption[User](opt)
		if len(u.Numbers) < opt.MinContainersLen || len(u.Numbers) > opt.MaxContainersLen {
			t.Fatalf("expected slice length in range [%d, %d], got %d", DefaultOption().MinContainersLen, DefaultOption().MaxContainersLen, len(u.Numbers))
		}
	})

	t.Run("skip fields", func(t *testing.T) {
		t.Parallel()

		opt := DefaultOption()
		opt.IgnoreFields = []string{"Skip1", "Skip2"}

		u := MakeWithOption[User](opt)
		if u.Skip1 != "" {
			t.Fatalf("expected default string value, got %v", u.Skip1)
		}
		if u.Skip2 != 0 {
			t.Fatalf("expected default int value, got %v", u.Skip2)
		}
	})
}
