package shapes

import (
	"log"

	"github.com/segmentio/ksuid"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
)

type Ray struct {
	Id        ksuid.KSUID
	Origin    *tuples.Tuple
	Direction *tuples.Tuple
}

func InitRay(o, d *tuples.Tuple) *Ray {
	r := Ray{ksuid.New(), o, d}
	if o.W != 1 {
		log.Fatal("Attempted to create a ray origin with a non-point")
	}
	if d.W != 0 {
		log.Fatal("Attempted to create a ray direction with a non-vector")
	}
	return &r
}

func (r *Ray) Position(t float64) *tuples.Tuple {
	return r.Origin.Add(r.Direction.MultiplyScalar(t))
}

func (r *Ray) Transform(m *matrix.Matrix) *Ray {
	return InitRay(m.MultiplyTuple(r.Origin), m.MultiplyTuple(r.Direction))
}
