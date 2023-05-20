package shapes

import (
	"math"
	"sync"

	"github.com/segmentio/ksuid"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
)

type Sphere struct {
	Id               ksuid.KSUID
	transform        *matrix.Matrix
	transformInverse *matrix.Matrix
	material         *Material
	xs               sync.Map
}

func InitSphere() *Sphere {
	return &Sphere{
		ksuid.New(),
		matrix.InitMatrixIdentity(4),
		matrix.InitMatrixIdentity(4),
		DefaultMaterial(),
		sync.Map{},
	}
}

func (s Sphere) Transform() *matrix.Matrix {
	return s.transform
}

func (s Sphere) Material() *Material {
	return s.material
}

func (s Sphere) Intersect(r *Ray) *Intersections {
	if v, ok := s.xs.Load(r.Id); ok {
		if xs, ok := v.(*Intersections); ok {
			return xs
		}
	}
	xs := InitIntersections()
	s.xs.Store(r.Id, &xs)

	r = r.Transform(s.transformInverse)

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

func (s *Sphere) SetTransform(t *matrix.Matrix) {
	s.transform = t
	s.transformInverse = t.Inverse()
}

func (s *Sphere) SetMaterial(m *Material) {
	s.material = m
}

func (s Sphere) NormalAt(p *tuples.Tuple) *tuples.Tuple {
	objectPoint := s.transformInverse.MultiplyTuple(p)
	objectNormal := objectPoint.Subtract(tuples.InitPoint(0, 0, 0))
	worldNormal := s.transformInverse.Transpose().MultiplyTuple(objectNormal)
	worldNormal.W = 0
	return worldNormal.Normalize()
}
