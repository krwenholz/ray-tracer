package world

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/lights"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

func TestCreatingAWorld(t *testing.T) {
	w := InitWorld()
	assert.Empty(t, w.Objects)
	assert.Empty(t, w.Lights)
}

func TestDefaultWorld(t *testing.T) {
	w := InitDefaultWorld()
	l := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.Material().Color = viz.InitColor(0.8, 1.0, 0.6)
	s1.Material().Diffuse = 0.7
	s1.Material().Specular = 0.2
	s2 := shapes.InitSphere()
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	assert.Equal(t, 2, len(w.Objects))
	assert.True(t, s1.Transform().Equals(w.Objects[0].Transform()))
	assert.True(t, s1.Material().Equals(w.Objects[0].Material()))
	assert.True(t, s2.Transform().Equals(w.Objects[1].Transform()))
	assert.True(t, s2.Material().Equals(w.Objects[1].Material()))
	assert.True(t, l.Equals(w.Lights[0]))
}

func TestIntersectDefaultWorld(t *testing.T) {
	w := InitDefaultWorld()
	r := shapes.InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	xs := w.Intersections(r)
	assert.Equal(t, 4.0, xs.Intersections[0].T)
	assert.Equal(t, 4.5, xs.Intersections[1].T)
	assert.Equal(t, 5.5, xs.Intersections[2].T)
	assert.Equal(t, 6.0, xs.Intersections[3].T)
}

func TestShadingAnIntersection(t *testing.T) {
	w := InitDefaultWorld()
	r := shapes.InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	i := shapes.InitIntersection(4, w.Objects[0])
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	log.Println("c", c)
	log.Println("wanted", 0.38066, 0.47583, 0.2855)
	assert.True(t, viz.InitColor(0.38066, 0.47583, 0.2855).Equals(c))
}

func TestShadingAnIntersectionFromTheInside(t *testing.T) {
	w := InitDefaultWorld()
	w.Lights[0] = lights.InitPointLight(tuples.InitPoint(0, 0.25, 0), viz.InitColor(1, 1, 1))
	r := shapes.InitRay(tuples.InitPoint(0, 0, 0), tuples.InitVector(0, 0, 1))
	i := shapes.InitIntersection(0.5, w.Objects[1])
	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)
	assert.True(t, viz.InitColor(0.90498, 0.90498, 0.90498).Equals(c))
}

func TestColorWhenARayMisses(t *testing.T) {
	w := InitDefaultWorld()
	r := shapes.InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 1, 0))
	c := w.ColorAt(r)
	assert.True(t, viz.Black().Equals(c))
}

func TestColorWhenARayHits(t *testing.T) {
	w := InitDefaultWorld()
	r := shapes.InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	c := w.ColorAt(r)
	assert.True(t, viz.InitColor(0.38066, 0.47583, 0.2855).Equals(c))
}

func TestColorWithAnIntersectionBehindTheRay(t *testing.T) {
	w := InitDefaultWorld()
	outer := w.Objects[0]
	outer.Material().Ambient = 1
	inner := w.Objects[1]
	inner.Material().Ambient = 1
	r := shapes.InitRay(tuples.InitPoint(0, 0, 0.75), tuples.InitVector(0, 0, -1))
	c := w.ColorAt(r)
	assert.True(t, inner.Material().Color.Equals(c))
}

func TestShadows(t *testing.T) {
	type opt struct {
		w   *World
		p   *tuples.Tuple
		exp bool
		msg string
	}
	opts := []opt{
		{
			InitDefaultWorld(),
			tuples.InitPoint(0, 10, 0),
			false,
			"There is no shadow when nothing is colinear with point and light",
		},
		{
			InitDefaultWorld(),
			tuples.InitPoint(10, -10, 10),
			true,
			"The shadow when the object is between the point and light",
		},
		{
			InitDefaultWorld(),
			tuples.InitPoint(-20, 20, -20),
			false,
			"There is no shadow when object is behind the light",
		},
		{
			InitDefaultWorld(),
			tuples.InitPoint(-2, 2, -2),
			false,
			"There is no shadow when an object is behind the point",
		},
	}
	for _, o := range opts {
		assert.Equal(t, o.exp, o.w.IsShadowed(o.p), o.msg)
	}
}

func TestShadeHitIsGivenAnIntersectionInShadow(t *testing.T) {
	w := InitWorld()
	l := lights.InitPointLight(tuples.InitPoint(0, 0, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s2 := shapes.InitSphere()
	s2.SetTransform(matrix.Translation(0, 0, 10))
	w.Objects = []shapes.Object{s1, s2}
	w.Lights = []*lights.PointLight{l}

	r := shapes.InitRay(tuples.InitPoint(0, 0, 5), tuples.InitVector(0, 0, 1))
	i := shapes.InitIntersection(4, s2)

	comps := i.PrepareComputations(r)
	c := w.ShadeHit(comps)

	assert.True(t, c.Equals(viz.InitColor(0.1, 0.1, 0.1)))
}
