package maths

import "math"

const EPSILON float64 = 0.00001

func FuzzyEquals(x, y float64) bool {
	return math.Abs(x-y) < EPSILON
}
