package world

import (
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
)

func ViewTransformation(from, to, up *tuples.Tuple) *matrix.Matrix {
	forward := to.Subtract(from).Normalize()
	upn := up.Normalize()
	left := forward.CrossProduct(upn)
	trueUp := left.CrossProduct(forward)

	orientation := matrix.InitMatrixIdentity(4)
	orientation.Set(0, 0, left.X)
	orientation.Set(0, 1, left.Y)
	orientation.Set(0, 2, left.Z)
	orientation.Set(1, 0, trueUp.X)
	orientation.Set(1, 1, trueUp.Y)
	orientation.Set(1, 2, trueUp.Z)
	orientation.Set(2, 0, -forward.X)
	orientation.Set(2, 1, -forward.Y)
	orientation.Set(2, 2, -forward.Z)

	return orientation.Multiply(
		matrix.Translation(-from.X, -from.Y, -from.Z),
	)
}
