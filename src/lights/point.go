package lights

import (
	"math"

	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

type PointLight struct {
	Position  *tuples.Tuple
	Intensity *viz.Color
}

func InitPointLight(p *tuples.Tuple, i *viz.Color) *PointLight {
	return &PointLight{p, i}
}

func (p *PointLight) Lighting(m *shapes.Material, point *tuples.Tuple, eyev *tuples.Tuple, normalv *tuples.Tuple) *viz.Color {
	var ambient, diffuse, specular *viz.Color
	// combine the surface color with the light's color/intensity
	effectiveColor := m.Color.Multiply(p.Intensity)
	// find the direction of the light source
	lightv := p.Position.Subtract(point).Normalize()
	ambient = effectiveColor.MultiplyScalar(m.Ambient)
	// lightDotNormal represents the cosine of the angle between the
	// light vector and the normal vector. A negative number means the light is on the other
	// side of the surface.
	lightDotNormal := lightv.DotProduct(normalv)
	if lightDotNormal < 0 {
		diffuse = viz.Black()
		specular = viz.Black()
	} else {
		// compute the diffuse contribution
		diffuse = effectiveColor.MultiplyScalar(m.Diffuse * lightDotNormal)

		// reflectDotEye represents the cosine of the angle between the reflection vector
		// and the eye vector. A negative number means the light reflects away from the eye.
		reflectv := lightv.Negate().Reflect(normalv)
		reflectDotEye := reflectv.DotProduct(eyev)
		if reflectDotEye <= 0 {
			specular = viz.Black()
		} else {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = p.Intensity.MultiplyScalar(m.Specular * factor)
		}
	}
	return ambient.Add(diffuse).Add(specular)
}

func (p *PointLight) Equals(p2 *PointLight) bool {
	return p.Intensity.Equals(p2.Intensity) && p.Position.Equals(p2.Position)
}
