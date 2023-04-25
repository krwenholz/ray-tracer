package shapes

import "sort"

type Intersection struct {
	T      float64
	Object Object
}

func InitIntersection(t float64, o Object) *Intersection {
	return &Intersection{t, o}
}

func (i *Intersection) Equals(i2 *Intersection) bool {
	return i.T == i2.T && i.Object.Equals(i2.Object)
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
