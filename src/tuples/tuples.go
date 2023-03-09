package tuples

import (
	"math"

	"happymonday.dev/ray-tracer/src/maths"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
}

func InitPoint(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 1}
}

func InitVector(x, y, z float64) *Tuple {
	return &Tuple{x, y, z, 0}
}
func (t *Tuple) Add(t2 *Tuple) *Tuple {
	return &Tuple{t.X + t2.X, t.Y + t2.Y, t.Z + t2.Z, t.W + t2.W}
}

func (t *Tuple) Subtract(t2 *Tuple) *Tuple {
	return &Tuple{t.X - t2.X, t.Y - t2.Y, t.Z - t2.Z, t.W - t2.W}
}

func (t *Tuple) Magnitude() float64 {
	return math.Sqrt(math.Pow(t.X, 2) + math.Pow(t.Y, 2) + math.Pow(t.Z, 2))
}

func (t *Tuple) Normalize() *Tuple {
	m := t.Magnitude()
	return InitVector(t.X/m, t.Y/m, t.Z/m)
}

func (t *Tuple) DotProduct(t2 *Tuple) float64 {
	return t.X*t2.X + t.Y*t2.Y + t.Z*t2.Z + t.W*t2.W
}

func (t *Tuple) CrossProduct(t2 *Tuple) *Tuple {
	return InitVector(t.Y*t2.Z-t.Z*t2.Y, t.Z*t2.X-t.X*t2.Z, t.X*t2.Y-t.Y*t2.X)
}

func (t *Tuple) Equals(t2 *Tuple) bool {
	return maths.FuzzyEquals(t.X, t2.X) && maths.FuzzyEquals(t.Y, t2.Y) && maths.FuzzyEquals(t.Z, t2.Z) && maths.FuzzyEquals(t.W, t2.W)
}

func (t *Tuple) Negate() *Tuple {
	return &Tuple{-t.X, -t.Y, -t.Z, t.W}
}

func (t *Tuple) MultiplyScalar(s float64) *Tuple {
	return &Tuple{t.X * s, t.Y * s, t.Z * s, t.W * s}
}

func (t *Tuple) Divide(s float64) *Tuple {
	return t.MultiplyScalar(1 / s)
}
