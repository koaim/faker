package faker

import (
	"slices"
	"testing"
	"unicode/utf8"
)

func TestMake(t *testing.T) {
	t.Parallel()

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		v := Make[int]()
		if v == 0 {
			t.Fatal("expected non-zero int")
		}
	})

	t.Run("uint", func(t *testing.T) {
		t.Parallel()

		v := Make[uint]()
		if v == 0 {
			t.Fatal("expected non-zero uint")
		}
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()

		v := Make[bool]()
		if v == false {
			t.Fatal("expected non-zero bool")
		}
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()

		v := Make[float64]()
		if v == 0 {
			t.Fatal("expected non-zero float46")
		}
	})

	t.Run("complex64", func(t *testing.T) {
		t.Parallel()

		v := Make[complex64]()
		if v == 0 {
			t.Fatal("expected non-zero complex64")
		}
	})

	t.Run("slice", func(t *testing.T) {
		t.Parallel()

		v := Make[[]int]()
		if len(v) < DefaultOption().MinLen || len(v) > DefaultOption().MaxLen {
			t.Fatalf("expected slice length in range [%d, %d], got %d", DefaultOption().MinLen, DefaultOption().MaxLen, len(v))
		}
		for _, v := range v {
			if v == 0 {
				t.Fatal("expected non-zero int")
			}
		}
	})

	t.Run("map", func(t *testing.T) {
		t.Parallel()

		v := Make[map[string]string]()
		if len(v) < DefaultOption().MinLen || len(v) > DefaultOption().MaxLen {
			t.Fatalf("expected map length in range [%d, %d], got %d", DefaultOption().MinLen, DefaultOption().MaxLen, len(v))
		}
		for k, v := range v {
			if k == "" {
				t.Fatal("expected non-zero string key")
			}
			if v == "" {
				t.Fatal("expected non-zero string value")
			}
		}
	})

	t.Run("array", func(t *testing.T) {
		t.Parallel()

		const length = 4

		v := Make[[length]int]()
		if len(v) != length {
			t.Fatalf("expected array length %d, got %d", length, len(v))
		}
		if cap(v) != length {
			t.Fatalf("expected array length %d, got %d", length, cap(v))
		}
	})

	t.Run("string", func(t *testing.T) {
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
	})

	t.Run("pointer", func(t *testing.T) {
		t.Parallel()

		v := Make[*string]()
		if v == nil {
			t.Fatal("expected non-nil pointer")
		}
		if *v == "" {
			t.Fatal("expected non-empty string")
		}
	})

	t.Run("chan", func(t *testing.T) {
		t.Parallel()

		v := Make[chan int]()

		if len(v) == 0 {
			t.Fatal("expected non-empty chan")
		}
	})

	t.Run("struct", func(t *testing.T) {
		t.Parallel()

		type Order struct {
			ID   int64
			Name string
		}

		type User struct {
			ID    int
			Name  string
			Order *Order
		}

		defaultUser := User{}
		defaultOrder := Order{}

		u := Make[User]()
		if u == defaultUser {
			t.Fatal("expected non-zero User value")
		}
		if u.Order == nil || *u.Order == defaultOrder {
			t.Fatal("expected non-nil and non-empty Order value")
		}
	})
}

func TestMakeAndOverride(t *testing.T) {
	t.Parallel()

	type User struct {
		ID   int
		Name string
	}

	defaultUser := User{}

	u := MakeAndOverride(func(v *User) {
		v.Name = "custom name"
	})

	if u == defaultUser {
		t.Fatal("expected not default value")
	}
	if u.Name != "custom name" {
		t.Fatalf("expected Name = %v, actual = %v", "custom name", u.Name)
	}
}
