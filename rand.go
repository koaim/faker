package faker

import (
	"math/rand/v2"
)

func randIn[T any](v ...T) T {
	randInd := rand.IntN(len(v))
	return v[randInd]
}

func randInRange[T int | int64](min, max T) T {
	return min + T(rand.Int64N(int64(max-min+1)))
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
