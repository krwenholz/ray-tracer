package shapes

import "happymonday.dev/ray-tracer/src/matrix"

type Object interface {
	Equals(i any) bool
	Transform() *matrix.Matrix
	Material() *Material
}
