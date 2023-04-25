package world

import (
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
	assert.Empty(t, w.Light)
}

func TestDefaultWorld(t *testing.T) {
	w := InitDefaultWorld()
	l := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.SetMaterial(shapes.InitMaterial(viz.InitColor(0.8, 1.0, 0.6), 0, 0.7, 0.2, 0))
	s2 := shapes.InitSphere()
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))

	assert.Equal(t, 2, len(w.Objects))
	assert.True(t, s1.Transform().Equals(w.Objects[0].Transform()))
	assert.True(t, s1.Material().Equals(w.Objects[0].Material()))
	assert.True(t, s2.Transform().Equals(w.Objects[1].Transform()))
	assert.True(t, s2.Material().Equals(w.Objects[1].Material()))
	assert.True(t, l.Equals(w.Light))
}
