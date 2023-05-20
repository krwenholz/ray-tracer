package shapes

import (
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
)

type Object interface {
	Equals(i any) bool
	Transform() *matrix.Matrix
	Material() *Material
	Intersect(r *Ray) *Intersections
	NormalAt(p *tuples.Tuple) *tuples.Tuple
}
