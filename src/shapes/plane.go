package shapes

import (
	"math"
	"sync"

	"happymonday.dev/ray-tracer/src/maths"
	"happymonday.dev/ray-tracer/src/tuples"
)

type Plane struct {
	*ShapeEmbed
	xs sync.Map
}

func InitPlane() *Plane {
	return &Plane{
		InitShapeEmbed(nil, nil),
		sync.Map{},
	}
}

func (s Plane) Intersect(r *Ray) *Intersections {
	if v, ok := s.xs.Load(r.Id); ok {
		if xs, ok := v.(*Intersections); ok {
			return xs
		}
	}
	xs := InitIntersections()
	s.xs.Store(r.Id, &xs)

	r = s.prepIntersect(r)

	if math.Abs(r.Direction.Y) < maths.EPSILON {
		return xs
	}

	t := -r.Origin.Y / r.Direction.Y

	return InitIntersections(InitIntersection(t, s))
}

func (s Plane) Equals(s2 any) bool {
	if v, ok := s2.(Plane); ok {
		return s.Id == v.Id
	}
	if v, ok := s2.(*Plane); ok {
		return s.Id == v.Id
	}
	return false
}

func (s Plane) NormalAt(p *tuples.Tuple) *tuples.Tuple {
	return s.normalAtPost(tuples.InitPoint(0, 1, 0))
}
