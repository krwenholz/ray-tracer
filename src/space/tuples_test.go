package space

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAPoint(t *testing.T) {
	var a Tuple
	a = &Point{4.3, -4.2, 3.1}
	assert.Equal(t, a.X(), 4.3)
	assert.Equal(t, a.Y(), -4.2)
	assert.Equal(t, a.Z(), 3.1)
	assert.Equal(t, a.W(), 1.0)
	assert.IsType(t, &Point{}, a)
}

func TestIsAVector(t *testing.T) {
	var a Tuple
	a = &Vector{4.3, -4.2, 3.1}
	assert.Equal(t, a.X(), 4.3)
	assert.Equal(t, a.Y(), -4.2)
	assert.Equal(t, a.Z(), 3.1)
	assert.Equal(t, a.W(), 0.0)
	assert.IsType(t, &Vector{}, a)
}
