package shapes

import (
	"github.com/segmentio/ksuid"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
)

type Shape interface {
	Equals(i any) bool
	Transform() *matrix.Matrix
	Material() *Material
	Intersect(r *Ray) *Intersections
	NormalAt(p *tuples.Tuple) *tuples.Tuple
}

type ShapeEmbed struct {
	Id               ksuid.KSUID
	transform        *matrix.Matrix
	transformInverse *matrix.Matrix
	material         *Material
}

func InitShapeEmbed(t *matrix.Matrix, m *Material) *ShapeEmbed {
	if t == nil {
		t = matrix.InitMatrixIdentity(4)
	}
	if m == nil {
		m = DefaultMaterial()
	}
	return &ShapeEmbed{
		ksuid.New(),
		t,
		t.Inverse(),
		m,
	}
}

func (s *ShapeEmbed) Transform() *matrix.Matrix {
	return s.transform
}

func (s *ShapeEmbed) TransformInverse() *matrix.Matrix {
	return s.transformInverse
}

func (s *ShapeEmbed) SetTransform(t *matrix.Matrix) {
	s.transform = t
	s.transformInverse = t.Inverse()
}

func (s *ShapeEmbed) Material() *Material {
	return s.material
}

func (s *ShapeEmbed) SetMaterial(m *Material) {
	s.material = m
}

func (s *ShapeEmbed) prepIntersect(r *Ray) *Ray {
	return r.Transform(s.transformInverse)
}

func (s *ShapeEmbed) normalAtPre(p *tuples.Tuple) *tuples.Tuple {
	return s.TransformInverse().MultiplyTuple(p)
}

func (s *ShapeEmbed) normalAtPost(localNormal *tuples.Tuple) *tuples.Tuple {
	worldNormal := s.TransformInverse().Transpose().MultiplyTuple(localNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}
