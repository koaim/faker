package faker

import "reflect"

const (
	minContainersLen = 1
	maxContainersLen = 3
	strLen           = 6
	yearsForward     = 200
)

var (
	allowedRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789123456789")
)

// Custom describes a user-defined generator for values of a specific type.
type Custom struct {
	// Type is the type whose values F generates.
	Type reflect.Type
	// F generates a value of Type.
	F func() any
}

// OptionF is a functional option that configures an Option.
type OptionF func(o *Option)

// Option holds the settings used when generating random values.
type Option struct {
	// MinContainersLen is the minimum length of generated containers (slices, maps).
	MinContainersLen int
	// MaxContainersLen is the maximum length of generated containers (slices, maps).
	MaxContainersLen int
	// StrLen is the length of generated strings.
	StrLen int
	// AllowedRunes are the characters used to build random strings.
	AllowedRunes []rune
	// IgnoreFields lists struct fields to skip when generating values.
	IgnoreFields []string
	// CloseChannels reports whether generated channels should be closed.
	CloseChannels bool
	// CustomFakers lists custom generators applied to values of the types
	// they specify, overriding the default generation rules.
	CustomFakers []Custom
}

// DefaultOption returns an Option with the default generation settings.
func DefaultOption() Option {
	return Option{
		MinContainersLen: minContainersLen,
		MaxContainersLen: maxContainersLen,
		AllowedRunes:     allowedRunes,
		StrLen:           strLen,
	}
}

// WithIgnoreFields returns an OptionF that adds the given fields to the
// list of struct fields to skip when generating values.
func WithIgnoreFields(fields ...string) OptionF {
	return func(o *Option) {
		o.IgnoreFields = append(o.IgnoreFields, fields...)
	}
}

// WithAllowedRunes returns an OptionF that overrides the runes used when
// generating random strings.
func WithAllowedRunes(r []rune) OptionF {
	return func(o *Option) {
		o.AllowedRunes = r
	}
}

// WithContainersLen returns an OptionF that sets the min and max length of
// generated containers (slices, maps).
func WithContainersLen(min, max uint) OptionF {
	return func(o *Option) {
		o.MinContainersLen = int(min)
		o.MaxContainersLen = int(max)
	}
}

// WithStrLen returns an OptionF that sets the length of generated strings.
func WithStrLen(l uint) OptionF {
	return func(o *Option) {
		o.StrLen = int(l)
	}
}

// WithCloseChannels returns an OptionF that enables closing generated
// channels.
func WithCloseChannels() OptionF {
	return func(o *Option) {
		o.CloseChannels = true
	}
}

// WithCustomFaker returns an OptionF that appends a Custom generator for
// values of type T. When a value of that exact type is generated, f is
// called and its result is used instead of the default generation rules.
func WithCustomFaker[T any](f func() T) OptionF {
	t := reflect.TypeFor[T]()
	return func(o *Option) {
		o.CustomFakers = append(o.CustomFakers, Custom{
			Type: t,
			F:    func() any { return f() },
		})
	}
}
