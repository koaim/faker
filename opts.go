package faker

const (
	minContainersLen = 1
	maxContainersLen = 3
	strLen           = 6
)

var (
	allowedRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

type OptionF func(o *Option)

type Option struct {
	MinLen        int
	MaxLen        int
	StrLen        int
	AllowedRunes  []rune
	IgnoreFields  []string
	CloseChannels bool
}

func DefaultOption() Option {
	return Option{
		MinLen:       minContainersLen,
		MaxLen:       maxContainersLen,
		AllowedRunes: allowedRunes,
		StrLen:       strLen,
	}
}

func WithIgnoreFields(fields ...string) OptionF {
	return func(o *Option) {
		o.IgnoreFields = append(o.IgnoreFields, fields...)
	}
}

func WithAllowedRunes(r []rune) OptionF {
	return func(o *Option) {
		o.AllowedRunes = r
	}
}
