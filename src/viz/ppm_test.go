package viz

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanvasToPPMHeader(t *testing.T) {
	c := InitCanvas(5, 3)
	ppm := CanvasToPPM(c)
	assert.Equal(t, "P3\n5 3\n255\n", ppm[:11])
}

func TestCanvasToPPMEndsWithNewline(t *testing.T) {
	c := InitCanvas(5, 3)
	ppm := CanvasToPPM(c)
	assert.True(t, strings.HasSuffix(ppm, "\n"))
}

func TestCanvasToPPMPixelData(t *testing.T) {
	c := InitCanvas(5, 3)
	c.SetPixel(InitColor(1.5, 0, 0), 0, 0)
	c.SetPixel(InitColor(0, 0.5, 0), 2, 1)
	c.SetPixel(InitColor(-0.5, 0, 1), 4, 2)
	ppm := CanvasToPPM(c)
	pixelData := "255 0 0 0 0 0 0 0 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 128 0 0 0 0 0 0 0\n0 0 0 0 0 0 0 0 0 0 0 0 0 0 255\n"
	assert.Equal(t, pixelData, ppm[11:])
}

func TestCanvasToPPMSplitLongLines(t *testing.T) {
	c := InitCanvas(10, 2)
	for i := 0; i < c.Height; i++ {
		for j := 0; j < c.Width; j++ {
			c.SetPixel(InitColor(1, 0.8, 0.6), j, i)
		}
	}
	ppm := CanvasToPPM(c)
	pixelData := "255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n255 204 153 255 204 153 255 204 153 255 204 153 255 204 153 255 204\n153 255 204 153 255 204 153 255 204 153 255 204 153\n"
	assert.Equal(t, pixelData, ppm[12:])
}
