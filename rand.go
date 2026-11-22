package faker

import (
	"math/rand/v2"
)

func randIn[T any](v ...T) T {
	randInd := rand.IntN(len(v))
	return v[randInd]
}

func randInRange[T int | int64 | uint | uint64](min, max T) T {
	n := uint64(max) - uint64(min) + 1
	if n == 0 {
		return T(rand.Uint64())
	}

	return min + T(rand.Uint64N(n))
}

func randStr(r []rune, length int) string {
	str := make([]rune, length)

	for i := range len(str) {
		index := rand.IntN(len(r))
		str[i] = r[index]
	}

	return string(str)
}

func randLen(opt Option) int {
	if opt.MaxContainersLen < opt.MinContainersLen {
		return opt.MinContainersLen
	}
	return opt.MinContainersLen + rand.IntN(opt.MaxContainersLen-opt.MinContainersLen+1)
}
