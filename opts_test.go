package faker

import (
	"math/rand/v2"
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

func Test_WithMinInt64Value(t *testing.T) {
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
