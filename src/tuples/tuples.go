package tuples

import (
	"math"

	"happymonday.dev/ray-tracer/src/maths"
)

type Tuple interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
}

type BaseTuple struct {
	x float64
	y float64
	z float64
	w float64
}

func (t BaseTuple) X() float64 {
	return t.x
}

func (t BaseTuple) Y() float64 {
	return t.y
}

func (t BaseTuple) Z() float64 {
	return t.z
}

func (t BaseTuple) W() float64 {
	return t.w
}

type Point struct {
	x float64
	y float64
	z float64
}

func (p Point) X() float64 {
	return p.x
}

func (p Point) Y() float64 {
	return p.y
}

func (p Point) Z() float64 {
	return p.z
}

func (p Point) W() float64 {
	return 1
}

type Vector struct {
	x float64
	y float64
	z float64
}

func (v Vector) X() float64 {
	return v.x
}

func (v Vector) Y() float64 {
	return v.y
}

func (v Vector) Z() float64 {
	return v.z
}

func (v Vector) W() float64 {
	return 0
}

func (v Vector) Magnitude() float64 {
	return math.Sqrt(math.Pow(v.x, 2) + math.Pow(v.y, 2) + math.Pow(v.z, 2))
}

func (v Vector) Normalize() Vector {
	m := v.Magnitude()
	return Vector{v.x / m, v.y / m, v.z / m}
}

func (v Vector) DotProduct(v2 Vector) float64 {
	return v.x*v2.x + v.y*v2.y + v.z*v2.z + v.W()*v2.W()
}

func (v Vector) CrossProduct(v2 Vector) Vector {
	return Vector{v.y*v2.z - v.z*v2.y, v.z*v2.x - v.x*v2.z, v.x*v2.y - v.y*v2.x}
}

func TupleEq(t1, t2 Tuple) bool {
	return maths.FuzzyEq(t1.X(), t2.X()) && maths.FuzzyEq(t1.Y(), t2.Y()) && maths.FuzzyEq(t1.Z(), t2.Z()) && maths.FuzzyEq(t1.W(), t2.W())
}

func TupleAdd(t1, t2 Tuple) Point {
	return Point{t1.X() + t2.X(), t1.Y() + t2.Y(), t1.Z() + t2.Z()}
}

func TupleSubtract(t1, t2 Tuple) Tuple {
	x := t1.X() - t2.X()
	y := t1.Y() - t2.Y()
	z := t1.Z() - t2.Z()
	w := t1.W() - t2.W()
	if w == 1 {
		return Point{x, y, z}
	}
	return Vector{x, y, z}
}

func TupleNegate(t Tuple) Tuple {
	if t.W() == 1 {
		return Point{-t.X(), -t.Y(), -t.Z()}
	}
	return Vector{-t.X(), -t.Y(), -t.Z()}
}

func TupleMultiply(t Tuple, s float64) Tuple {
	return BaseTuple{t.X() * s, t.Y() * s, t.Z() * s, t.W() * s}
}

func TupleDivide(t Tuple, s float64) Tuple {
	return TupleMultiply(t, 1/s)
}
