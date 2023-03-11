package matrix

import (
	"math"

	"happymonday.dev/ray-tracer/src/tuples"
)

func Translation(x, y, z float64) *Matrix {
	t := InitMatrixIdentity(4)
	t.Set(0, 3, x)
	t.Set(1, 3, y)
	t.Set(2, 3, z)
	return t
}

func Scaling(x, y, z float64) *Matrix {
	t := InitMatrixIdentity(4)
	t.Set(0, 0, x)
	t.Set(1, 1, y)
	t.Set(2, 2, z)
	return t
}

func ReflectingX(t *tuples.Tuple) *tuples.Tuple {
	return Scaling(-1, 1, 1).MultiplyTuple(t)
}

func ReflectingY(t *tuples.Tuple) *tuples.Tuple {
	return Scaling(1, -1, 1).MultiplyTuple(t)
}

func ReflectingZ(t *tuples.Tuple) *tuples.Tuple {
	return Scaling(1, 1, -1).MultiplyTuple(t)
}

func RotationX(pis float64) *Matrix {
	t := InitMatrixIdentity(4)
	t.Set(1, 1, math.Cos(pis*math.Pi))
	t.Set(1, 2, -math.Sin(pis*math.Pi))
	t.Set(2, 1, math.Sin(pis*math.Pi))
	t.Set(2, 2, math.Cos(pis*math.Pi))
	return t
}

func RotationY(pis float64) *Matrix {
	t := InitMatrixIdentity(4)
	t.Set(0, 0, math.Cos(pis*math.Pi))
	t.Set(0, 2, math.Sin(pis*math.Pi))
	t.Set(2, 0, -math.Sin(pis*math.Pi))
	t.Set(2, 2, math.Cos(pis*math.Pi))
	return t
}

func RotationZ(pis float64) *Matrix {
	t := InitMatrixIdentity(4)
	t.Set(0, 0, math.Cos(pis*math.Pi))
	t.Set(0, 1, -math.Sin(pis*math.Pi))
	t.Set(1, 0, math.Sin(pis*math.Pi))
	t.Set(1, 1, math.Cos(pis*math.Pi))
	return t
}

func Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	t := InitMatrixIdentity(4)
	t.Set(0, 1, xy)
	t.Set(0, 2, xz)
	t.Set(1, 0, yx)
	t.Set(1, 2, yz)
	t.Set(2, 0, zx)
	t.Set(2, 1, zy)
	return t
}

func Chain(ts ...*Matrix) *Matrix {
	res := InitMatrixIdentity(4)
	for i := range ts {
		res = res.Multiply(ts[len(ts)-(i+1)])
	}
	return res
}

func (m *Matrix) Translation(x, y, z float64) *Matrix {
	return Translation(x, y, z).Multiply(m)
}

func (m *Matrix) Scaling(x, y, z float64) *Matrix {
	return Scaling(x, y, z).Multiply(m)
}

func (m *Matrix) RotationX(pis float64) *Matrix {
	return RotationX(pis).Multiply(m)
}

func (m *Matrix) RotationY(pis float64) *Matrix {
	return RotationY(pis).Multiply(m)
}

func (m *Matrix) RotationZ(pis float64) *Matrix {
	return RotationZ(pis).Multiply(m)
}

func (m *Matrix) Shearing(xy, xz, yx, yz, zx, zy float64) *Matrix {
	return Shearing(xy, xz, yx, yz, zx, zy).Multiply(m)
}
