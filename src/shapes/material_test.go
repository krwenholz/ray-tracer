package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/viz"
)

func TestDefaultMaterial(t *testing.T) {
	m := DefaultMaterial()
	assert.True(t, m.Color.Equals(viz.InitColor(1, 1, 1).Tuple))
	assert.Equal(t, 0.1, m.Ambient)
	assert.Equal(t, 0.9, m.Diffuse)
	assert.Equal(t, 0.9, m.Specular)
	assert.Equal(t, 200.0, m.Shininess)
}
