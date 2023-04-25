package viz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCanvas(t *testing.T) {
	c := InitCanvas(20, 10)
	assert.Equal(t, c.Width, 20)
	assert.Equal(t, c.Height, 10)

	zeroColor := InitColor(0, 0, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			assert.True(t, c.Pixel(j, i).Equals(zeroColor))
		}
	}
}

func TestSetColor(t *testing.T) {
	c := InitCanvas(20, 10)

	assert.True(t, c.Pixel(2, 3).Equals(InitColor(0, 0, 0)))

	red := InitColor(1, 0, 0)
	c.SetPixel(red, 2, 3)

	assert.True(t, c.Pixel(2, 3).Equals(red))
}
