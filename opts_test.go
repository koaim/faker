package faker

import (
	"math/rand/v2"
	"slices"
	"testing"
)

func Test_WithRecursiveDepth(t *testing.T) {
	t.Parallel()

	depth := rand.Uint()

	o := DefaultOption()
	WithRecursiveDepth(depth)(&o)

	if o.MaxRecursiveDepth != depth {
		t.Fatalf("\nwant %v\nhave %v", depth, o.MaxRecursiveDepth)
	}
}

func Test_WithMinUint64Value(t *testing.T) {
	t.Parallel()

	minVal := rand.Uint64()

	o := DefaultOption()
	WithMinUint64Value(minVal)(&o)

	if o.MinUint64Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinUint64Value)
	}
}

func Test_WithMinUint32Value(t *testing.T) {
	t.Parallel()

	minVal := rand.Uint32()

	o := DefaultOption()
	WithMinUint32Value(minVal)(&o)

	if o.MinUint32Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinUint32Value)
	}
}

func Test_WithMinUintValue(t *testing.T) {
	t.Parallel()

	minVal := rand.Uint()

	o := DefaultOption()
	WithMinUintValue(minVal)(&o)

	if o.MinUintValue != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinUintValue)
	}
}

func Test_WithMinIntValue(t *testing.T) {
	t.Parallel()

	minVal := rand.Int64()

	o := DefaultOption()
	WithMinIntValue(minVal)(&o)

	if o.MinIntValue != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinIntValue)
	}
}

func Test_WithMinInt32Value(t *testing.T) {
	t.Parallel()

	minVal := rand.Int32()

	o := DefaultOption()
	WithMinInt32Value(minVal)(&o)

	if o.MinInt32Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinIntValue)
	}
}

func Test_WithContainersLen(t *testing.T) {
	t.Parallel()

	min, max := rand.Uint(), rand.Uint()

	o := DefaultOption()
	WithContainersLen(min, max)(&o)

	if o.MinContainersLen != int(min) {
		t.Fatalf("\nwant %v\nhave %v", min, o.MinContainersLen)
	}
	if o.MaxContainersLen != int(max) {
		t.Fatalf("\nwant %v\nhave %v", min, o.MaxContainersLen)
	}
}

func Test_WithIgnoreFields(t *testing.T) {
	t.Parallel()

	fields := []string{"ID", "CreatedAt", "UpdatedAt"}

	o := DefaultOption()
	WithIgnoreFields(fields...)(&o)

	if !slices.Equal(o.IgnoreFields, fields) {
		t.Fatalf("\nwant %v\nhave %v", fields, o.IgnoreFields)
	}
}

func Test_WithAllowedRunes(t *testing.T) {
	t.Parallel()

	allowedRunes := []rune("abc123")

	o := DefaultOption()
	WithAllowedRunes(allowedRunes)(&o)

	if !slices.Equal(o.AllowedRunes, allowedRunes) {
		t.Fatalf("\nwant %v\nhave %v", allowedRunes, o.AllowedRunes)
	}
}

func Test_WithAllowedDomainZones(t *testing.T) {
	t.Parallel()

	zones := []string{"dev", "test"}

	o := DefaultOption()
	WithAllowedDomainZones(zones)(&o)

	if !slices.Equal(o.Url.AllowedDomainZones, zones) {
		t.Fatalf("\nwant %v\nhave %v", zones, o.Url.AllowedDomainZones)
	}
}

func Test_WithAllowedUrlSchemes(t *testing.T) {
	t.Parallel()

	schemes := []string{"ftp", "ssh"}

	o := DefaultOption()
	WithAllowedUrlSchemes(schemes)(&o)

	if !slices.Equal(o.Url.AllowedSchemes, schemes) {
		t.Fatalf("\nwant %v\nhave %v", schemes, o.Url.AllowedSchemes)
	}
}

func Test_WithUrlUserInfo(t *testing.T) {
	t.Parallel()

	withUserInfo := false

	o := DefaultOption()
	WithUrlUserInfo(withUserInfo)(&o)

	if o.Url.WithUserInfo != withUserInfo {
		t.Fatalf("\nwant %v\nhave %v", withUserInfo, o.Url.WithUserInfo)
	}
}

func Test_WithUrlQueryCount(t *testing.T) {
	t.Parallel()

	queryCount := rand.Uint()

	o := DefaultOption()
	WithUrlQueryCount(queryCount)(&o)

	if o.Url.QueryCount != int(queryCount) {
		t.Fatalf("\nwant %v\nhave %v", queryCount, o.Url.QueryCount)
	}
}

func Test_WithUrlPathLen(t *testing.T) {
	t.Parallel()

	pathLen := rand.Uint()

	o := DefaultOption()
	WithUrlPathLen(pathLen)(&o)

	if o.Url.PathLen != int(pathLen) {
		t.Fatalf("\nwant %v\nhave %v", pathLen, o.Url.PathLen)
	}
}

func Test_WithStrLen(t *testing.T) {
	t.Parallel()

	strLen := rand.Uint()

	o := DefaultOption()
	WithStrLen(strLen)(&o)

	if o.StrLen != int(strLen) {
		t.Fatalf("\nwant %v\nhave %v", strLen, o.StrLen)
	}
}

func Test_WithCloseChannels(t *testing.T) {
	t.Parallel()

	o := DefaultOption()
	WithCloseChannels()(&o)

	if !o.CloseChannels {
		t.Fatal("expected CloseChannels to be true")
	}
}

func Test_WithTimeYearsBack(t *testing.T) {
	t.Parallel()

	yearsBack := rand.Uint()

	o := DefaultOption()
	WithTimeYearsBack(yearsBack)(&o)

	if o.Time.YearsBack != int(yearsBack) {
		t.Fatalf("\nwant %v\nhave %v", yearsBack, o.Time.YearsBack)
	}
}

func Test_WithTimeYearsForward(t *testing.T) {
	t.Parallel()

	yearsForward := rand.Uint()

	o := DefaultOption()
	WithTimeYearsForward(yearsForward)(&o)

	if o.Time.YearsForward != int(yearsForward) {
		t.Fatalf("\nwant %v\nhave %v", yearsForward, o.Time.YearsForward)
	}
}

func Test_WithMinInt8Value(t *testing.T) {
	t.Parallel()

	minVal := int8(rand.Int())

	o := DefaultOption()
	WithMinInt8Value(minVal)(&o)

	if o.MinInt8Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinInt8Value)
	}
}

func Test_WithMinInt16Value(t *testing.T) {
	t.Parallel()

	minVal := int16(rand.Int())

	o := DefaultOption()
	WithMinInt16Value(minVal)(&o)

	if o.MinInt16Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinInt16Value)
	}
}

func Test_WithMinInt64Value(t *testing.T) {
	t.Parallel()

	minVal := rand.Int64()

	o := DefaultOption()
	WithMinInt64Value(minVal)(&o)

	if o.MinInt64Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinInt64Value)
	}
}

func Test_WithMinUint8Value(t *testing.T) {
	t.Parallel()

	minVal := uint8(rand.Uint())

	o := DefaultOption()
	WithMinUint8Value(minVal)(&o)

	if o.MinUint8Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinUint8Value)
	}
}

func Test_WithMinUint16Value(t *testing.T) {
	t.Parallel()

	minVal := uint16(rand.Uint())

	o := DefaultOption()
	WithMinUint16Value(minVal)(&o)

	if o.MinUint16Value != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinUint16Value)
	}
}

func Test_WithMinUintptrValue(t *testing.T) {
	t.Parallel()

	minVal := uintptr(rand.Uint())

	o := DefaultOption()
	WithMinUintptrValue(minVal)(&o)

	if o.MinUintptrValue != minVal {
		t.Fatalf("\nwant %v\nhave %v", minVal, o.MinUintptrValue)
	}
}

func Test_WithBuilder(t *testing.T) {
	t.Parallel()

	const want = "custom value"

	o := DefaultOption()
	WithBuilder(func() string { return want })(&o)

	if o.CustomFakers[0].Type.Name() != "string" {
		t.Fatalf("want custom faker for string, have %v", o.CustomFakers[0].Type)
	}
	if have := o.CustomFakers[0].F(); have != want {
		t.Fatalf("\nwant %v\nhave %v", want, have)
	}
}
