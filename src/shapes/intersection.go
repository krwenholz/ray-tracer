package shapes

import (
	"sort"

	"happymonday.dev/ray-tracer/src/maths"
	"happymonday.dev/ray-tracer/src/tuples"
)

type Intersection struct {
	T      float64
	Object Shape
}

type IntersectionComputations struct {
	T         float64
	Object    Shape
	Point     *tuples.Tuple
	EyeV      *tuples.Tuple
	NormalV   *tuples.Tuple
	OverPoint *tuples.Tuple
	Inside    bool
}

func InitIntersection(t float64, o Shape) *Intersection {
	return &Intersection{t, o}
}

func (i *Intersection) Equals(i2 *Intersection) bool {
	return i.T == i2.T && i.Object.Equals(i2.Object)
}

func (i *Intersection) PrepareComputations(r *Ray) *IntersectionComputations {
	c := IntersectionComputations{T: i.T, Object: i.Object}
	c.Point = r.Position(c.T)
	c.NormalV = c.Object.NormalAt(c.Point)
	c.EyeV = r.Direction.Negate()
	c.Inside = c.NormalV.DotProduct(c.EyeV) < 0
	if c.Inside {
		c.NormalV = c.NormalV.Negate()
	}

	c.OverPoint = c.Point.Add(c.NormalV.MultiplyScalar(maths.EPSILON))
	return &c
}

type Intersections struct {
	Intersections []*Intersection
	hit           int
}

func InitIntersections(is ...*Intersection) *Intersections {
	res := &Intersections{is, -1}
	res.Sort()
	return res
}

func (is *Intersections) Add(i *Intersection) {
	is.Intersections = append(is.Intersections, i)
	is.Sort()
}

func (is *Intersections) Sort() {
	sort.Slice(is.Intersections, func(i, j int) bool {
		return is.Intersections[i].T < is.Intersections[j].T
	})
	for idx, i := range is.Intersections {
		if i.T > 0 {
			is.hit = idx
			return
		}
	}
}

func (is *Intersections) Hit() *Intersection {
	if is.hit < 0 {
		return nil
	}
	return is.Intersections[is.hit]
}
