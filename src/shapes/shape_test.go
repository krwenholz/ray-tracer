package shapes

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

type TestShape struct {
	*ShapeEmbed
}

func InitTestShape() *TestShape {
	return &TestShape{
		InitShapeEmbed(nil, nil),
	}
}

func (s *TestShape) TestPrepIntersect(r *Ray) *Ray {
	return s.prepIntersect(r)
}

func (s *TestShape) NormalAt(p *tuples.Tuple) *tuples.Tuple {
	localNormal := s.normalAtPre(p)
	localNormal = tuples.InitVector(localNormal.X, localNormal.Y, localNormal.Z)
	return s.normalAtPost(localNormal)
}

func TestDefaultTransformation(t *testing.T) {
	s := InitTestShape()
	assert.True(t, s.transform.Equals(matrix.InitMatrixIdentity(4)))
}

func TestAssigningATransformation(t *testing.T) {
	s := InitTestShape()
	m := matrix.Translation(2, 3, 4)
	s.SetTransform(m)
	assert.True(t, s.transform.Equals(m))
	assert.True(t, s.transformInverse.Equals(m.Inverse()))
}

func TestShapeHasADefaultMaterial(t *testing.T) {
	s := InitTestShape()
	assert.True(t, s.Material().Equals(DefaultMaterial()))
}

func TestShapeMayBeAssignedAMaterial(t *testing.T) {
	s := InitTestShape()
	s.material = InitMaterial(viz.Black(), 1, 1, 1, 1)
	m := InitMaterial(viz.Black(), 1, 1, 1, 1)
	assert.True(t, s.Material().Equals(m))
}

func TestIntersectingAScaledShapeWithARay(t *testing.T) {
	s := InitTestShape()
	r := InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	m := matrix.Scaling(2, 2, 2)
	s.SetTransform(m)
	res := s.TestPrepIntersect(r)
	assert.True(t, res.Origin.Equals(tuples.InitPoint(0, 0, -2.5)))
	assert.True(t, res.Direction.Equals(tuples.InitVector(0, 0, 0.5)))
}

func TestIntersectingATranslatedShapeWithARay(t *testing.T) {
	s := InitTestShape()
	r := InitRay(tuples.InitPoint(0, 0, -5), tuples.InitVector(0, 0, 1))
	m := matrix.Translation(5, 0, 0)
	s.SetTransform(m)
	res := s.TestPrepIntersect(r)
	assert.True(t, res.Origin.Equals(tuples.InitPoint(-5, 0, -5)))
	assert.True(t, res.Direction.Equals(tuples.InitVector(0, 0, 1)))
}

func TestNormalOnATransformedShape(t *testing.T) {
	type opt struct {
		s string
		t *matrix.Matrix
		p *tuples.Tuple
		v *tuples.Tuple
	}
	opts := []opt{
		{
			s: "Computing the normal on a translated shape",
			t: matrix.Translation(0, 1, 0),
			p: tuples.InitPoint(0, 1.70711, -0.70711),
			v: tuples.InitVector(0, 0.70711, -0.70711),
		},
		{
			s: "Computing the normal on a transformed shape",
			t: matrix.Scaling(1, 0.5, 1).Multiply(matrix.RotationZ(1 / 5)),
			p: tuples.InitPoint(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			v: tuples.InitVector(0, 0.97014, -0.24254),
		},
	}
	for _, o := range opts {
		s := InitTestShape()
		s.SetTransform(o.t)
		n := s.NormalAt(o.p)
		assert.True(t, o.v.Equals(n), o.s)
	}
}
