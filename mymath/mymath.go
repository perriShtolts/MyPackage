package mymath

import "math"

func Sqrt(x float64) int {
	return int(math.Sqrt(x))
}

func Abs(x float64) int {
	return int(math.Abs(x))
}

func Max(a float64, b float64) int {
	return int(math.Max(a, b))
}

func Yn(n int, x float64) int {
	return int(math.Yn(n, x))
}
