package tuples

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAPoint(t *testing.T) {
	a := InitPoint(4.3, -4.2, 3.1)
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 1.0)
}

func TestIsAVector(t *testing.T) {
	a := InitVector(4.3, -4.2, 3.1)
	assert.Equal(t, a.X, 4.3)
	assert.Equal(t, a.Y, -4.2)
	assert.Equal(t, a.Z, 3.1)
	assert.Equal(t, a.W, 0.0)
}

func TestIsTupelEq(t *testing.T) {
	assert.True(t, InitVector(4.3, -4.2, 3.1).Equals(InitVector(4.3, -4.2, 3.1)), "definitely equal")
	assert.True(t, InitVector(4.3, -4.2, 3.1).Equals(InitVector(4.3, -4.2, 3.100009)), "very close equal")
	assert.False(t, InitVector(4.3, -4.2, 3.1).Equals(InitVector(4.3, -4.2, 3.10002)), "not equal")
	assert.True(t, InitPoint(4.3, -4.2, 3.1).Equals(InitPoint(4.3, -4.2, 3.1)), "point equal")
}

func TestTupleAdd(t *testing.T) {
	a1 := InitPoint(3, -2, 5)
	a2 := InitVector(-2, 3, 1)
	assert.True(t, a1.Add(a2).Equals(InitPoint(1, 1, 6)))
}

func TestTupleSubtractPoints(t *testing.T) {
	a1 := InitPoint(3, 2, 1)
	a2 := InitPoint(5, 6, 7)
	assert.True(t, a1.Subtract(a2).Equals(InitVector(-2, -4, -6)))
}

func TestTupleSubtractPointAndVector(t *testing.T) {
	a1 := InitPoint(3, 2, 1)
	a2 := InitVector(5, 6, 7)
	assert.True(t, a1.Subtract(a2).Equals(InitPoint(-2, -4, -6)))
}

func TestTupleSubtractVectors(t *testing.T) {
	a1 := InitVector(3, 2, 1)
	a2 := InitVector(5, 6, 7)
	assert.True(t, a1.Subtract(a2).Equals(InitVector(-2, -4, -6)))
}

func TestSubtractFromZeroVector(t *testing.T) {
	a1 := InitVector(0, 0, 0)
	a2 := InitVector(1, -2, 3)
	assert.True(t, a1.Subtract(a2).Equals(InitVector(-1, 2, -3)))
}

func TestNegate(t *testing.T) {
	a1 := InitVector(1, -2, 3)
	a2 := InitPoint(1, -2, 3)
	assert.True(t, a1.Negate().Equals(InitVector(-1, 2, -3)))
	assert.True(t, a2.Negate().Equals(InitPoint(-1, 2, -3)))
}

func TestTupleMultiply(t *testing.T) {
	a := InitVector(1, 2, 3)
	assert.True(t, a.MultiplyScalar(2).Equals(InitVector(2, 4, 6)), "vector")
	b := Tuple{1, -2, 3, -4}
	assert.True(t, b.MultiplyScalar(3.5).Equals(&Tuple{3.5, -7, 10.5, -14}), "scalar")
	assert.True(t, b.MultiplyScalar(0.5).Equals(&Tuple{0.5, -1, 1.5, -2}), "fraction")
}

func TestTupleDivide(t *testing.T) {
	b := Tuple{1, -2, 3, -4}
	assert.True(t, b.Divide(2).Equals(&Tuple{0.5, -1, 1.5, -2}))
}

func TestVectorMagnitude(t *testing.T) {
	type opt struct {
		v *Tuple
		m float64
		s string
	}
	vs := []opt{
		{InitVector(1, 0, 0), 1, "x"},
		{InitVector(0, 1, 0), 1, "y"},
		{InitVector(0, 0, 1), 1, "z"},
		{InitVector(0, 0, 0), 0, "zero"},
		{InitVector(1, 2, 3), math.Sqrt(14), "all some"},
		{InitVector(-1, -2, -3), math.Sqrt(14), "all neg"},
	}
	for _, o := range vs {
		assert.Equal(t, o.v.Magnitude(), o.m, o.s)
	}
}

func TestVectorNormalize(t *testing.T) {
	type opt struct {
		v *Tuple
		r *Tuple
		s string
	}
	vs := []opt{
		{InitVector(4, 0, 0), InitVector(1, 0, 0), "x"},
		{InitVector(1, 2, 3), InitVector(1/math.Sqrt(14), 02/math.Sqrt(14), 03/math.Sqrt(14)), "all"},
	}
	for _, o := range vs {
		assert.True(t, o.v.Normalize().Equals(o.r), o.s)
		assert.Equal(t, o.v.Normalize().Magnitude(), 1.0, o.s)
	}
}

func TestVectorDotProduct(t *testing.T) {
	type opt struct {
		v1 *Tuple
		v2 *Tuple
		p  float64
		s  string
	}
	vs := []opt{
		{InitVector(1, 2, 3), InitVector(2, 3, 4), 20, "simple"},
	}
	for _, o := range vs {
		assert.Equal(t, o.v1.DotProduct(o.v2), o.p, o.s)
	}
}

func TestVectorCrossProduct(t *testing.T) {
	type opt struct {
		v1 *Tuple
		v2 *Tuple
		r  *Tuple
		s  string
	}
	vs := []opt{
		{InitVector(1, 2, 3), InitVector(2, 3, 4), InitVector(-1, 2, -1), "simple"},
	}
	for _, o := range vs {
		assert.True(t, o.v1.CrossProduct(o.v2).Equals(o.r), o.s)
		assert.True(t, o.v2.CrossProduct(o.v1).Equals(o.r.Negate()), o.s)
	}
}

func TestReflectingVectors(t *testing.T) {
	type opt struct {
		v   *Tuple
		n   *Tuple
		res *Tuple
		s   string
	}
	vs := []opt{
		{
			InitVector(1, -1, 0),
			InitVector(0, 1, 0),
			InitVector(1, 1, 0),
			"Reflecting a vector approaching at 45 degrees",
		},
		{
			InitVector(0, -1, 0),
			InitVector(math.Sqrt(2)/2, math.Sqrt(2)/2, 0),
			InitVector(1, 0, 0),
			"Reflecting a vector off a scaled surface",
		},
	}
	for _, o := range vs {
		r := o.v.Reflect(o.n)
		assert.True(t, o.res.Equals(r), o.s)
	}
}
