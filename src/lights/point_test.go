package lights

import (
	"log"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

func TestAPointLightHasAPositionAndIntensity(t *testing.T) {
	intensity := viz.InitColor(1, 1, 1)
	position := tuples.InitPoint(0, 0, 0)
	light := InitPointLight(position, intensity)
	assert.True(t, light.Position.Equals(position))
	assert.True(t, light.Intensity.Equals(intensity))
}

func TestLighting(t *testing.T) {
	type opt struct {
		eyev    *tuples.Tuple
		normalv *tuples.Tuple
		point   *tuples.Tuple
		color   *viz.Color
		exp     *viz.Color
		msg     string
	}
	m := shapes.DefaultMaterial()
	position := tuples.InitPoint(0, 0, 0)
	opts := []opt{
		{
			eyev:    tuples.InitVector(0, 0, -1),
			normalv: tuples.InitVector(0, 0, -1),
			point:   tuples.InitPoint(0, 0, -10),
			color:   viz.InitColor(1, 1, 1),
			exp:     viz.InitColor(1.9, 1.9, 1.9),
			msg:     "Lighting with the eye betwen the light and the surface",
		},
		{
			eyev:    tuples.InitVector(0, math.Sqrt(2)/2, -math.Sqrt(2)/2),
			normalv: tuples.InitVector(0, 0, -1),
			point:   tuples.InitPoint(0, 0, -10),
			color:   viz.InitColor(1, 1, 1),
			exp:     viz.InitColor(1.0, 1.0, 1.0),
			msg:     "Lighting with the eye between light and surface, eye offset 45 degrees",
		},
		{
			eyev:    tuples.InitVector(0, 0, -1),
			normalv: tuples.InitVector(0, 0, -1),
			point:   tuples.InitPoint(0, 10, -10),
			color:   viz.InitColor(1, 1, 1),
			exp:     viz.InitColor(0.7364, 0.7364, 0.7364),
			msg:     "Lighting with eye opposite surface, light offset 45 degrees",
		},
		{
			eyev:    tuples.InitVector(0, -math.Sqrt(2)/2, -math.Sqrt(2)/2),
			normalv: tuples.InitVector(0, 0, -1),
			point:   tuples.InitPoint(0, 10, -10),
			color:   viz.InitColor(1, 1, 1),
			exp:     viz.InitColor(1.6364, 1.6364, 1.6364),
			msg:     "Lighting with eye in the path of the reflection vector",
		},
		{
			eyev:    tuples.InitVector(0, 0, -1),
			normalv: tuples.InitVector(0, 0, -1),
			point:   tuples.InitPoint(0, 0, 10),
			color:   viz.InitColor(1, 1, 1),
			exp:     viz.InitColor(0.1, 0.1, 0.1),
			msg:     "Lighting with the light behind the surface",
		},
	}
	for _, o := range opts {
		light := InitPointLight(o.point, o.color)
		result := light.Lighting(m, position, o.eyev, o.normalv)
		log.Println(result.Tuple)
		assert.True(t, o.exp.Equals(result), o.msg)
	}
}
