package faker

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"
)

const (
	minContainersLen       = 1
	maxContainersLen       = 3
	strLen                 = 6
	timeYearsBack          = 200
	timeYearsForward       = 200
	maxTimeYears     int64 = math.MaxInt64 / (366 * 24 * 60 * 60)

	urlQueryCount   = 2
	urlWithUserInfo = true
	pathLen         = 3
)

var (
	allowedDomainZones = []string{"com", "net", "org", "en", "site"}
	allowedUrlSchemes  = []string{"http", "https"}
	allowedRunes       = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789123456789")
)

// Custom describes a user-defined generator for values of a specific type.
type Custom struct {
	// Type is the type whose values F generates.
	Type reflect.Type
	// F generates a value of Type.
	F func() any
}

// OptionFunc is a functional option that configures an Option.
type OptionFunc func(o *Option)

// Option holds the settings used when generating random values.
type Option struct {
	// MinContainersLen is the minimum requested length of containers (at least 1).
	// Map key collisions can reduce the resulting map length.
	MinContainersLen int
	// MaxContainersLen is the maximum length of slices, maps, and channels.
	MaxContainersLen int
	// StrLen is the length of generated strings in runes (at least 1).
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

	// Url is the setting for generating random URL values.
	Url UrlOption
	// Time is the setting for generating random time.Time values.
	Time TimeOption
}

func (o Option) validate(t *testing.T) {
	t.Helper()

	if err := o.validationError(); err != nil {
		t.Fatal(err)
	}
}

func (o Option) validationError() error {
	var errs []error

	if o.MinContainersLen < 1 {
		errs = append(errs, fmt.Errorf("MinContainersLen must be >= 1"))
	}

	if o.MaxContainersLen < o.MinContainersLen {
		errs = append(errs, fmt.Errorf("MaxContainersLen must be >= MinContainersLen"))
	}

	if o.MaxContainersLen-o.MinContainersLen == math.MaxInt {
		errs = append(errs, fmt.Errorf("container length range is too large"))
	}

	if o.StrLen < 1 {
		errs = append(errs, fmt.Errorf("StrLen must be >= 1"))
	}

	if o.Time.YearsBack < 0 {
		errs = append(errs, fmt.Errorf("Time.YearsBack must be non-negative"))
	}
	if o.Time.YearsForward < 0 {
		errs = append(errs, fmt.Errorf("Time.YearsForward must be non-negative"))
	}
	if o.Time.YearsBack >= 0 && o.Time.YearsForward >= 0 && int64(o.Time.YearsBack) > maxTimeYears-int64(o.Time.YearsForward) {
		errs = append(errs, fmt.Errorf("time range is too large"))
	}

	if len(o.AllowedRunes) == 0 {
		errs = append(errs, fmt.Errorf("AllowedRunes must not be empty"))
	}

	for i, r := range o.AllowedRunes {
		if !utf8.ValidRune(r) {
			errs = append(errs, fmt.Errorf("AllowedRunes[%d] must be a valid Unicode rune", i))
		}
	}

	seenTypes := make(map[reflect.Type]int, len(o.CustomFakers))
	for i, f := range o.CustomFakers {
		if f.Type == nil {
			errs = append(errs, fmt.Errorf("CustomFakers[%d].Type must not be nil", i))
		} else if first, ok := seenTypes[f.Type]; ok {
			errs = append(errs, fmt.Errorf("CustomFakers[%d].Type duplicates CustomFakers[%d].Type (%v)", i, first, f.Type))
		} else {
			seenTypes[f.Type] = i
		}
		if f.F == nil {
			errs = append(errs, fmt.Errorf("CustomFakers[%d].F must not be nil", i))
		}
	}

	if o.Url.QueryCount < 0 {
		errs = append(errs, fmt.Errorf("Url.QueryCount must be non-negative"))
	}

	if o.Url.PathLen < 0 {
		errs = append(errs, fmt.Errorf("Url.PathLen must be non-negative"))
	}

	if len(o.Url.AllowedDomainZones) == 0 {
		errs = append(errs, fmt.Errorf("Url.AvailableDomainZones must not be empty"))
	}

	for i, zone := range o.Url.AllowedDomainZones {
		if strings.TrimSpace(zone) == "" {
			errs = append(errs, fmt.Errorf("Url.AvailableDomainZones[%d] must not be empty", i))
		}
	}

	if len(o.Url.AllowedSchemes) == 0 {
		errs = append(errs, fmt.Errorf("Url.AvailableSchemes must not be empty"))
	}

	for i, scheme := range o.Url.AllowedSchemes {
		if strings.TrimSpace(scheme) == "" {
			errs = append(errs, fmt.Errorf("Url.AllowedSchemes[%d] must not be empty", i))
		}
	}

	return errors.Join(errs...)
}

// UrlOption holds the settings used when generating random URL values.
type UrlOption struct {
	// QueryCount is the number of query parameters to add to generated URLs.
	QueryCount int
	// WithUserInfo reports whether generated URLs include user info (login and password).
	WithUserInfo bool
	// AllowedDomainZones lists the domain zones used to build URL hosts.
	AllowedDomainZones []string
	// AllowedSchemes lists the URL schemes used when generating URLs.
	AllowedSchemes []string
	// PathLen is the length of the generated URL path in runes.
	PathLen int
}

// TimeOption sets the calendar-year range around now for time.Time values.
// Both zero means the current Unix second, with random nanoseconds.
// Negative years and ranges that risk Unix-second overflow are rejected.
type TimeOption struct {
	// YearsBack is the number of calendar years before now (inclusive).
	YearsBack int
	// YearsForward is the number of calendar years after now (inclusive).
	YearsForward int
}

// DefaultOption returns an Option with the default generation settings.
func DefaultOption() Option {
	return Option{
		MinContainersLen: minContainersLen,
		MaxContainersLen: maxContainersLen,
		AllowedRunes:     slices.Clone(allowedRunes),
		StrLen:           strLen,
		Time: TimeOption{
			YearsBack:    timeYearsBack,
			YearsForward: timeYearsForward,
		},
		Url: UrlOption{
			QueryCount:         urlQueryCount,
			WithUserInfo:       urlWithUserInfo,
			AllowedDomainZones: slices.Clone(allowedDomainZones),
			AllowedSchemes:     slices.Clone(allowedUrlSchemes),
			PathLen:            pathLen,
		},
	}
}

// WithIgnoreFields adds the given fields to the
// list of struct fields to skip when generating values.
func WithIgnoreFields(fields ...string) OptionFunc {
	return func(o *Option) {
		o.IgnoreFields = append(o.IgnoreFields, fields...)
	}
}

// WithAllowedRunes sets the runes used when generating random strings.
func WithAllowedRunes(r []rune) OptionFunc {
	return func(o *Option) {
		o.AllowedRunes = r
	}
}

// WithAllowedDomainZones sets the domain zones used to build URL hosts.
func WithAllowedDomainZones(z []string) OptionFunc {
	return func(o *Option) {
		o.Url.AllowedDomainZones = z
	}
}

// WithAllowedUrlSchemes sets the schemes used when generating URLs.
func WithAllowedUrlSchemes(s []string) OptionFunc {
	return func(o *Option) {
		o.Url.AllowedSchemes = s
	}
}

// WithUrlUserInfo sets whether generated URLs include
// user info (login and password).
func WithUrlUserInfo(val bool) OptionFunc {
	return func(o *Option) {
		o.Url.WithUserInfo = val
	}
}

// WithUrlQueryCount sets the number of query parameters to add to generated URLs.
func WithUrlQueryCount(c int) OptionFunc {
	return func(o *Option) {
		o.Url.QueryCount = c
	}
}

// WithUrlPathLen sets the length of generated URL paths in runes.
func WithUrlPathLen(l int) OptionFunc {
	return func(o *Option) {
		o.Url.PathLen = l
	}
}

// WithContainersLen sets the min and max length of
// slices, maps, and channels. The bounds must satisfy 1 <= min <= max.
func WithContainersLen(min, max uint) OptionFunc {
	return func(o *Option) {
		o.MinContainersLen = int(min)
		o.MaxContainersLen = int(max)
	}
}

// WithStrLen sets the length of generated strings in runes. l must be at least 1.
func WithStrLen(l uint) OptionFunc {
	return func(o *Option) {
		o.StrLen = int(l)
	}
}

// WithCloseChannels enables closing generated channels.
func WithCloseChannels() OptionFunc {
	return func(o *Option) {
		o.CloseChannels = true
	}
}

// WithTimeYearsBack sets the non-negative number of calendar years before now.
func WithTimeYearsBack(years int) OptionFunc {
	return func(o *Option) {
		o.Time.YearsBack = years
	}
}

// WithTimeYearsForward sets the non-negative number of calendar years after now.
func WithTimeYearsForward(years int) OptionFunc {
	return func(o *Option) {
		o.Time.YearsForward = years
	}
}

// WithCustomFaker appends a Custom generator for
// values of type T. When a value of that exact type is generated, f is
// called and its result is used instead of the default generation rules.
func WithCustomFaker[T any](f func() T) OptionFunc {
	return func(o *Option) {
		if f == nil {
			return
		}

		o.CustomFakers = append(o.CustomFakers, Custom{
			Type: reflect.TypeFor[T](),
			F:    func() any { return f() },
		})
	}
}
