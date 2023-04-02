package shapes

import (
	"log"

	"happymonday.dev/ray-tracer/src/viz"
)

type Material struct {
	Color     viz.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func DefaultMaterial() *Material {
	return &Material{viz.InitColor(1, 1, 1), 0.1, 0.9, 0.9, 200.0}
}

func InitMaterial(c viz.Color, a, d, sp, sh float64) *Material {
	if a < 0 || d < 0 || sp < 0 || sh < 0 {
		log.Fatal("Material creation attempted with negative values", a, d, sp, sh)
	}
	return &Material{c, a, d, sp, sh}
}

func (m *Material) Equals(m2 *Material) bool {
	return (m.Color.Equals(m2.Color.Tuple) &&
		m.Ambient == m2.Ambient &&
		m.Diffuse == m2.Diffuse &&
		m.Specular == m2.Specular &&
		m.Shininess == m2.Shininess)
}
