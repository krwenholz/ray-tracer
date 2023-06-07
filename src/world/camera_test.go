package world

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

func TestConstructingACamera(t *testing.T) {
	hsize := 160
	vsize := 120
	fov := math.Pi / 2
	c := InitCamera(hsize, vsize, fov)
	assert.Equal(t, hsize, c.HSize)
	assert.Equal(t, vsize, c.VSize)
	assert.Equal(t, fov, c.FOV)
	assert.True(t, matrix.InitMatrixIdentity(4).Equals(c.transform))
}

func TestPixelSizeForAHorizontalCanvas(t *testing.T) {
	c := InitCamera(200, 125, math.Pi/2)
	assert.Equal(t, 0.01, c.PixelSize)
}

func TestPixelSizeForAVerticalCanvas(t *testing.T) {
	c := InitCamera(125, 200, math.Pi/2)
	assert.Equal(t, 0.01, c.PixelSize)
}

func TestConstructingARayThroughTheCenterOfTheCanvas(t *testing.T) {
	c := InitCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(100, 50)
	assert.True(t, tuples.InitPoint(0, 0, 0).Equals(r.Origin))
	assert.True(t, tuples.InitVector(0, 0, -1).Equals(r.Direction))
}

func TestConstructingARayThroughACornerOfTheCanvas(t *testing.T) {
	c := InitCamera(201, 101, math.Pi/2)
	r := c.RayForPixel(0, 0)
	assert.True(t, tuples.InitPoint(0, 0, 0).Equals(r.Origin))
	assert.True(t, tuples.InitVector(0.66519, 0.33259, -0.66851).Equals(r.Direction))
}

func TestConstructingARayWhenCameraIsTransformed(t *testing.T) {
	c := InitCamera(201, 101, math.Pi/2.0)
	c.SetTransform(matrix.RotationY(1.0 / 4.0).Multiply(matrix.Translation(0, -2, 5)))
	r := c.RayForPixel(100, 50)
	assert.True(t, tuples.InitPoint(0, 2, -5).Equals(r.Origin))
	fmt.Println(r.Direction)
	assert.True(t, tuples.InitVector(math.Sqrt(2)/2, 0, -math.Sqrt(2)/2).Equals(r.Direction))
}
func TestRenderingAWorldWithACamera(t *testing.T) {
	w := InitDefaultWorld()
	c := InitCamera(11, 11, math.Pi/2.0)
	from := tuples.InitPoint(0, 0, -5)
	to := tuples.InitPoint(0, 0, 0)
	up := tuples.InitVector(0, 1, 0)
	c.SetTransform(ViewTransformation(from, to, up))
	image := c.Render(w)
	assert.True(t, viz.InitColor(0.38066, 0.47583, 0.2855).Equals(image.Pixel(5, 5)))
}
