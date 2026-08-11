package faker

const (
	minContainersLen = 1
	maxContainersLen = 3
	strLen           = 6
	yearsForward     = 200
)

var (
	allowedRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789123456789")
)

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
func WithContainersLen(min, max int) OptionF {
	return func(o *Option) {
		o.MinContainersLen = min
		o.MaxContainersLen = max
	}
}

// WithStrLen returns an OptionF that sets the length of generated strings.
func WithStrLen(l int) OptionF {
	return func(o *Option) {
		o.StrLen = l
	}
}

// WithCloseChannels returns an OptionF that enables closing generated
// channels.
func WithCloseChannels() OptionF {
	return func(o *Option) {
		o.CloseChannels = true
	}
}
