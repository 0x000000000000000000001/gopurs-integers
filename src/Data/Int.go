import (
	"math"
	"strconv"
)

func FromNumberImpl(just func(int) any, nothing any, n float64) any {
	if math.IsNaN(n) || math.IsInf(n, 0) || math.Trunc(n) != n {
		return nothing
	}
	return just(int(n))
}

func ToNumber(n int) float64 {
	return float64(n)
}

func FromStringAsImpl(just func(int) any, nothing any, radix int, s string) any {
	val, err := strconv.ParseInt(s, radix, 64)
	if err != nil {
		return nothing
	}
	return just(int(val))
}

func ToStringAs(radix int, i int) string {
	return strconv.FormatInt(int64(i), radix)
}

func Quot(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}

func Rem(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x % y
}

func Pow(x int, y int) int {
	if y < 0 {
		return 0
	}
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}
