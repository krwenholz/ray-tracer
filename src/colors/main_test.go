package colors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAColor(t *testing.T) {
	var a Color
	a = InitColor(-0.5, 0.4, 1.7)
	assert.Equal(t, a.R(), -0.5)
	assert.Equal(t, a.G(), 0.4)
	assert.Equal(t, a.B(), 1.7)
	assert.Equal(t, a.W, 0.0)
}

func TestColorsAdd(t *testing.T) {
	a := InitColor(0.9, 0.6, 0.75)
	b := InitColor(0.7, 0.1, 0.25)
	assert.True(t, a.Add(b).Eq(InitColor(1.6, 0.7, 1.0)))
}

func TestColorsSubtract(t *testing.T) {
	a := InitColor(0.9, 0.6, 0.75)
	b := InitColor(0.7, 0.1, 0.25)
	assert.True(t, a.Subtract(b).Eq(InitColor(0.2, 0.5, 0.5)))
}

func TestColorsMultiplyScalar(t *testing.T) {
	a := InitColor(0.2, 0.3, 0.4)
	assert.True(t, a.MultiplyScalar(2).Eq(InitColor(0.4, 0.6, 0.8)))
}

func TestColorsMultiply(t *testing.T) {
	a := InitColor(1.0, 0.2, 0.4)
	b := InitColor(0.9, 1.0, 0.1)
	assert.True(t, a.Multiply(b).Eq(InitColor(0.9, 0.2, 0.04)))
}
