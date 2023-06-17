package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/tuples"
)

func TestNormalOfAPlaneIsAConstantEverywhere(t *testing.T) {
	p := InitPlane()
	n1 := p.NormalAt(tuples.InitPoint(0, 0, 0))
	n2 := p.NormalAt(tuples.InitPoint(10, 0, -10))
	n3 := p.NormalAt(tuples.InitPoint(-5, 0, 150))

	assert.True(t, tuples.InitVector(0, 1, 0).Equals(n1))
	assert.True(t, tuples.InitVector(0, 1, 0).Equals(n2))
	assert.True(t, tuples.InitVector(0, 1, 0).Equals(n3))
}

func TestPlaneIntersections(t *testing.T) {
	type opt struct {
		s       string
		r       *Ray
		xsCount int
		xsT     float64
	}
	opts := []opt{
		{
			s: "Intersect with a ray parallel to the plane",
			r: InitRay(tuples.InitPoint(0, 10, 0), tuples.InitVector(0, 0, 1)),
		},
		{
			s: "Intersect with a coplanar ray",
			r: InitRay(tuples.InitPoint(0, 0, 0), tuples.InitVector(0, 0, 1)),
		},
		{
			s:       "Intersect with a ray from above",
			r:       InitRay(tuples.InitPoint(0, 1, 0), tuples.InitVector(0, -1, 0)),
			xsCount: 1,
			xsT:     1,
		},
		{
			s:       "Intersect with a ray from below",
			r:       InitRay(tuples.InitPoint(0, -1, 0), tuples.InitVector(0, 1, 0)),
			xsCount: 1,
			xsT:     1,
		},
	}
	for _, o := range opts {
		p := InitPlane()
		xs := p.Intersect(o.r)
		if o.xsCount > 0 {
			assert.Equal(t, o.xsCount, len(xs.Intersections), o.s)
			assert.Equal(t, o.xsT, xs.Hit().T, o.s)
			assert.True(t, p.Equals(xs.Hit().Object), o.s)
		} else {
			assert.Nil(t, xs.Hit(), o.s)
		}
	}
}
