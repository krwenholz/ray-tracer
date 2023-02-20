package maths

import "math"

func FuzzyEq(x, y float64) bool {
	epsilon := 0.00001
	return math.Abs(x-y) < epsilon
}
