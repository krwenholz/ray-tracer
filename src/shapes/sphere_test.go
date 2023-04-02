package shapes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

func TestARayIntersectsASphereAtTwoPoints(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs.Intersections))
	assert.Equal(t, 4.0, xs.Intersections[0].T)
	assert.Equal(t, 6.0, xs.Intersections[1].T)
}

func TestARayIntersectsASphereAtATangent(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 1, -5), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs.Intersections))
	assert.Equal(t, 5.0, xs.Intersections[0].T)
	assert.Equal(t, 5.0, xs.Intersections[1].T)
}

func TestARayMissesASphere(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 2, -5), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 0, len(xs.Intersections))
}

func TestARayOriginatesInsideASphere(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 0, 0), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs.Intersections))
	assert.Equal(t, -1.0, xs.Intersections[0].T)
	assert.Equal(t, 1.0, xs.Intersections[1].T)
}

func TestARayStartsPastASphere(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 0, 5), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs.Intersections))
	assert.Equal(t, -6.0, xs.Intersections[0].T)
	assert.Equal(t, -4.0, xs.Intersections[1].T)
}

func TestIntersectSetsTheObjectOnTheIntersection(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs.Intersections))
	assert.True(t, s.Equals(xs.Intersections[0].Object))
	assert.True(t, s.Equals(xs.Intersections[1].Object))
}

func TestSphereDefaultTranformation(t *testing.T) {
	s := InitSphere()
	assert.True(t, s.Transform.Equals(matrix.InitMatrixIdentity(4)))
}

func TestSphereSetTransform(t *testing.T) {
	s := InitSphere()
	s.SetTransform(matrix.Translation(2, 3, 4))
	assert.True(t, s.Transform.Equals(matrix.Translation(2, 3, 4)))
}

func TestIntersectingAScaledSphere(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	s.SetTransform(matrix.Scaling(2, 2, 2))
	xs := s.Intersect(r)
	assert.Equal(t, 2, len(xs.Intersections))
	assert.Equal(t, 3.0, xs.Intersections[0].T)
	assert.Equal(t, 7.0, xs.Intersections[1].T)
}

func TestIntersectingATranslatedSphere(t *testing.T) {
	r := InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	s := InitSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))
	xs := s.Intersect(r)
	assert.Equal(t, 0, len(xs.Intersections))
}

func TestNormalOnASphere(t *testing.T) {
	type opt struct {
		s string
		p *tuples.Tuple
		v *tuples.Tuple
	}
	opts := []opt{
		{
			s: "The normal on a sphere at a point on the x axis",
			p: tuples.InitPoint(1, 0, 0),
			v: tuples.InitVector(1, 0, 0),
		},
		{
			s: "The normal on a sphere at a point on the y axis",
			p: tuples.InitPoint(0, 1, 0),
			v: tuples.InitVector(0, 1, 0),
		},
		{
			s: "The normal on a sphere at a point on the z axis",
			p: tuples.InitPoint(0, 0, 1),
			v: tuples.InitVector(0, 0, 1),
		},
		{
			s: "The normal on a sphere at a nonaxial point",
			p: tuples.InitPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
			v: tuples.InitVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3),
		},
	}
	for _, o := range opts {
		s := InitSphere()
		n := s.NormalAt(o.p)
		assert.True(t, o.v.Equals(n), o.s)
	}
}

func TestNormalIsANormalizedVector(t *testing.T) {
	s := InitSphere()
	n := s.NormalAt(tuples.InitPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)))
	assert.True(t, n.Normalize().Equals(n))
}

func TestNormalOnATransformedSphere(t *testing.T) {
	type opt struct {
		s string
		t *matrix.Matrix
		p *tuples.Tuple
		v *tuples.Tuple
	}
	opts := []opt{
		{
			s: "Computing the normal on a translated sphere",
			t: matrix.Translation(0, 1, 0),
			p: tuples.InitPoint(0, 1.70711, -0.70711),
			v: tuples.InitVector(0, 0.70711, -0.70711),
		},
		{
			s: "Computing the normal on a transformed sphere",
			t: matrix.Scaling(1, 0.5, 1).Multiply(matrix.RotationZ(1 / 5)),
			p: tuples.InitPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			v: tuples.InitVector(0, 0.97014, -0.24254),
		},
	}
	for _, o := range opts {
		s := InitSphere()
		s.SetTransform(o.t)
		n := s.NormalAt(o.p)
		assert.True(t, o.v.Equals(n), o.s)
	}
}

func TestSphereHasADefaultMaterial(t *testing.T) {
	s := InitSphere()
	assert.True(t, s.Material.Equals(DefaultMaterial()))
}

func TestSphereMayBeAssignedAMaterial(t *testing.T) {
	s := InitSphere()
	s.Material = InitMaterial(viz.Black(), 1, 1, 1, 1)
	m := InitMaterial(viz.Black(), 1, 1, 1, 1)
	assert.True(t, s.Material.Equals(m))
}
