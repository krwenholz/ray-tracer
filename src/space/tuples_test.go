package space

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAPoint(t *testing.T) {
	var a Tuple
	a = &Point{4.3, -4.2, 3.1}
	assert.Equal(t, a.X(), 4.3)
	assert.Equal(t, a.Y(), -4.2)
	assert.Equal(t, a.Z(), 3.1)
	assert.Equal(t, a.W(), 1.0)
	assert.IsType(t, &Point{}, a)
}

func TestIsAVector(t *testing.T) {
	var a Tuple
	a = &Vector{4.3, -4.2, 3.1}
	assert.Equal(t, a.X(), 4.3)
	assert.Equal(t, a.Y(), -4.2)
	assert.Equal(t, a.Z(), 3.1)
	assert.Equal(t, a.W(), 0.0)
	assert.IsType(t, &Vector{}, a)
}

func TestIsTupelEq(t *testing.T) {
	assert.True(t, TupleEq(&Vector{4.3, -4.2, 3.1}, &Vector{4.3, -4.2, 3.1}), "definitely equal")
	assert.True(t, TupleEq(&Vector{4.3, -4.2, 3.1}, &Vector{4.3, -4.2, 3.100009}), "very close equal")
	assert.False(t, TupleEq(&Vector{4.3, -4.2, 3.1}, &Vector{4.3, -4.2, 3.10002}), "not equal")
	assert.True(t, TupleEq(&Point{4.3, -4.2, 3.1}, &Point{4.3, -4.2, 3.1}), "point equal")
}

func TestTupleAdd(t *testing.T) {
	a1 := Point{3, -2, 5}
	a2 := Vector{-2, 3, 1}
	assert.True(t, TupleEq(TupleAdd(&a1, &a2), &Point{1, 1, 6}))
}

func TestTupleSubtractPoints(t *testing.T) {
	a1 := Point{3, 2, 1}
	a2 := Point{5, 6, 7}
	assert.True(t, TupleEq(TupleSubtract(&a1, &a2), &Vector{-2, -4, -6}))
}

func TestTupleSubtractPointAndVector(t *testing.T) {
	a1 := Point{3, 2, 1}
	a2 := Vector{5, 6, 7}
	assert.True(t, TupleEq(TupleSubtract(&a1, &a2), &Point{-2, -4, -6}))
}

func TestTupleSubtractVectors(t *testing.T) {
	a1 := Vector{3, 2, 1}
	a2 := Vector{5, 6, 7}
	assert.True(t, TupleEq(TupleSubtract(&a1, &a2), &Vector{-2, -4, -6}))
}

func TestSubtractFromZeroVector(t *testing.T) {
	a1 := Vector{0, 0, 0}
	a2 := Vector{1, -2, 3}
	assert.True(t, TupleEq(TupleSubtract(&a1, &a2), &Vector{-1, 2, -3}))
}

func TestTupleNegate(t *testing.T) {
	a1 := Vector{1, -2, 3}
	a2 := Point{1, -2, 3}
	assert.True(t, TupleEq(TupleNegate(&a1), &Vector{-1, 2, -3}))
	assert.True(t, TupleEq(TupleNegate(&a2), &Point{-1, 2, -3}))
}

func TestTupleMultiply(t *testing.T) {
	a := Vector{1, 2, 3}
	assert.True(t, TupleEq(TupleMultiply(&a, 2), &BaseTuple{2, 4, 6, 0}), "vector")
	b := BaseTuple{1, -2, 3, -4}
	assert.True(t, TupleEq(TupleMultiply(&b, 3.5), &BaseTuple{3.5, -7, 10.5, -14}), "scalar")
	assert.True(t, TupleEq(TupleMultiply(&b, 0.5), &BaseTuple{0.5, -1, 1.5, -2}), "fraction")
}

func TestTupleDivide(t *testing.T) {
	b := BaseTuple{1, -2, 3, -4}
	assert.True(t, TupleEq(TupleDivide(&b, 2), &BaseTuple{0.5, -1, 1.5, -2}))
}

func TestVectorMagnitude(t *testing.T) {
	type opt struct {
		v Vector
		m float64
		s string
	}
	vs := []opt{
		{Vector{1, 0, 0}, 1, "x"},
		{Vector{0, 1, 0}, 1, "y"},
		{Vector{0, 0, 1}, 1, "z"},
		{Vector{0, 0, 0}, 0, "zero"},
		{Vector{1, 2, 3}, math.Sqrt(14), "all some"},
		{Vector{-1, -2, -3}, math.Sqrt(14), "all neg"},
	}
	for _, o := range vs {
		assert.Equal(t, o.v.Magnitude(), o.m, o.s)
	}
}

func TestVectorNormalize(t *testing.T) {
	type opt struct {
		v Vector
		r Vector
		s string
	}
	vs := []opt{
		{Vector{4, 0, 0}, Vector{1, 0, 0}, "x"},
		{Vector{1, 2, 3}, Vector{1 / math.Sqrt(14), 02 / math.Sqrt(14), 03 / math.Sqrt(14)}, "all"},
	}
	for _, o := range vs {
		assert.True(t, TupleEq(o.v.Normalize(), o.r), o.s)
		assert.Equal(t, o.v.Normalize().Magnitude(), 1.0, o.s)
	}
}

func TestVectorDotProduct(t *testing.T) {
	type opt struct {
		v1 Vector
		v2 Vector
		p  float64
		s  string
	}
	vs := []opt{
		{Vector{1, 2, 3}, Vector{2, 3, 4}, 20, "simple"},
	}
	for _, o := range vs {
		assert.Equal(t, o.v1.DotProduct(o.v2), o.p, o.s)
	}
}

func TestVectorCrossProduct(t *testing.T) {
	type opt struct {
		v1 Vector
		v2 Vector
		r  Vector
		s  string
	}
	vs := []opt{
		{Vector{1, 2, 3}, Vector{2, 3, 4}, Vector{-1, 2, -1}, "simple"},
	}
	for _, o := range vs {
		assert.True(t, TupleEq(o.v1.CrossProduct(o.v2), o.r), o.s)
		assert.True(t, TupleEq(o.v2.CrossProduct(o.v1), TupleNegate(o.r)), o.s)
	}
}
