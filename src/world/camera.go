package world

import (
	"math"
	"sync"

	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

type Camera struct {
	HSize            int
	VSize            int
	FOV              float64
	PixelSize        float64
	HalfWidth        float64
	HalfHeight       float64
	transform        *matrix.Matrix
	transformInverse *matrix.Matrix
}

func InitCamera(hsize, vsize int, fov float64) *Camera {
	halfView := math.Tan(fov / 2.0)
	aspect := float64(hsize) / float64(vsize)
	halfWidth := 0.0
	halfHeight := 0.0
	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfHeight = halfView
		halfWidth = halfView * aspect
	}
	pixelSize := halfWidth * 2.0 / float64(hsize)

	return &Camera{
		HSize:            hsize,
		VSize:            vsize,
		FOV:              fov,
		PixelSize:        pixelSize,
		HalfWidth:        halfWidth,
		HalfHeight:       halfHeight,
		transform:        matrix.InitMatrixIdentity(4),
		transformInverse: matrix.InitMatrixIdentity(4),
	}
}

func (c *Camera) SetTransform(t *matrix.Matrix) {
	c.transform = t
	c.transformInverse = t.Inverse()
}

func (c *Camera) RayForPixel(px, py int) *shapes.Ray {
	// the offset from the edge of the canvas to the pixel's center
	xOffset := (float64(px) + 0.5) * c.PixelSize
	yOffset := (float64(py) + 0.5) * c.PixelSize
	// the untransformed coordinates of teh pixel in world space
	// (camera looks toward -z, so +x is te the "left")
	worldX := c.HalfWidth - xOffset
	worldY := c.HalfHeight - yOffset
	// using the camera matrix, transform the canvas point and the origin,
	// then compute the ray's direction vector
	// (the canvas is at z=-1)
	pixel := c.transformInverse.MultiplyTuple(tuples.InitPoint(worldX, worldY, -1))
	origin := c.transformInverse.MultiplyTuple(tuples.InitPoint(0, 0, 0))
	direction := pixel.Subtract(origin).Normalize()
	return shapes.InitRay(origin, direction)
}

func (c *Camera) Render(w *World) *viz.Canvas {
	image := viz.InitCanvas(c.HSize, c.VSize)
	wg := sync.WaitGroup{}
	wg.Add(c.HSize * c.VSize)

	for iy := 0; iy < c.VSize; iy++ {
		for ix := 0; ix < c.HSize; ix++ {
			y := iy
			x := ix
			go func() {
				r := c.RayForPixel(x, y)
				color := w.ColorAt(r)
				image.SetPixel(color, x, y)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return &image
}
