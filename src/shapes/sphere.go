package shapes

import (
	"math"
	"sync"

	"happymonday.dev/ray-tracer/src/tuples"
)

type Sphere struct {
	*ShapeEmbed
	xs sync.Map
}

func InitSphere() *Sphere {
	return &Sphere{
		InitShapeEmbed(nil, nil),
		sync.Map{},
	}
}

func (s Sphere) Intersect(r *Ray) *Intersections {
	if v, ok := s.xs.Load(r.Id); ok {
		if xs, ok := v.(*Intersections); ok {
			return xs
		}
	}
	xs := InitIntersections()
	s.xs.Store(r.Id, &xs)

	r = s.prepIntersect(r)

	sphereToRay := r.Origin.Subtract(tuples.InitPoint(0, 0, 0))
	a := r.Direction.DotProduct(r.Direction)
	b := 2 * r.Direction.DotProduct(sphereToRay)
	c := sphereToRay.DotProduct(sphereToRay) - 1
	d := math.Pow(b, 2) - (4 * a * c)
	if d < 0 {
		return xs
	}

	xs.Add(InitIntersection((-b-math.Sqrt(d))/(2*a), s))
	xs.Add(InitIntersection((-b+math.Sqrt(d))/(2*a), s))
	return xs
}

func (s Sphere) Equals(s2 any) bool {
	if v, ok := s2.(Sphere); ok {
		return s.Id == v.Id
	}
	if v, ok := s2.(*Sphere); ok {
		return s.Id == v.Id
	}
	return false
}

func (s Sphere) NormalAt(p *tuples.Tuple) *tuples.Tuple {
	return s.normalAtPost(s.normalAtPre(p).Subtract(tuples.InitPoint(0, 0, 0)))
}
