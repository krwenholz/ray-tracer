package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
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
