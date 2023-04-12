package three_d_ray_cast

import (
	"image"
	"image/color"
	"image/color/palette"
	"sync"

	"github.com/schollz/progressbar/v3"
	"happymonday.dev/ray-tracer/src/lights"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

var tock *matrix.Matrix = matrix.RotationZ(-1 / 12.0)

type BasicCast struct {
	H            int
	W            int
	DefaultColor viz.Color
	Sphere       *shapes.Sphere
	Light        *lights.PointLight
}

func Init(h, w int, c viz.Color, s *shapes.Sphere, l *lights.PointLight) *BasicCast {
	b := BasicCast{
		H:            h,
		W:            w,
		DefaultColor: c,
		Sphere:       s,
		Light:        l,
	}
	return &b
}

func (c *BasicCast) Shine(source *tuples.Tuple) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, int(c.W), int(c.H)), palette.Plan9)
	wg := sync.WaitGroup{}
	wg.Add(c.Height() * c.Width())
	bar := progressbar.Default(int64(c.Height() * c.Width()))
	bar.Describe("Casting")

	for iy := 0; iy < c.Height(); iy++ {
		for ix := 0; ix < c.Width(); ix++ {
			y := iy
			x := ix
			go func() {
				hitColor := c.RayHits(source, x, y)
				if hitColor != nil {
					img.Set(
						x,
						y,
						color.RGBA{
							uint8(hitColor.R()),
							uint8(hitColor.G()),
							uint8(hitColor.B()),
							0,
						},
					)
				}
				bar.Add(1)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return img
}

func (c *BasicCast) RayHits(source *tuples.Tuple, x, y int) *viz.Color {
	v := tuples.InitVectorFromPoints(tuples.InitPoint(float64(x), float64(y), 100), source).Normalize()
	r := shapes.InitRay(source, v)
	h := c.Sphere.Intersect(r).Hit()
	if h == nil {
		return nil
	}
	point := r.Position(h.T)
	normal := c.Sphere.NormalAt(point)
	eye := r.Direction.Negate()
	return c.Light.Lighting(c.Sphere.Material, point, eye, normal)
}

func (c *BasicCast) DrawRGBA() *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, int(c.W), int(c.H)), palette.Plan9)
	for y := 0; y < c.Height(); y++ {
		for x := 0; x < c.Width(); x++ {
			img.Set(
				x,
				c.Height()-y,
				color.RGBA{
					uint8(viz.ScaledColorValue(c.DefaultColor.R())),
					uint8(viz.ScaledColorValue(c.DefaultColor.G())),
					uint8(viz.ScaledColorValue(c.DefaultColor.B())),
					0,
				},
			)
		}
	}
	return img
}

func (c *BasicCast) DrawPpm() *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, int(c.W), int(c.H)), palette.Plan9)
	for y := 0; y < c.Height(); y++ {
		for x := 0; x < c.Width(); x++ {
			img.Set(
				x,
				c.Height()-y,
				color.RGBA{
					uint8(viz.ScaledColorValue(c.DefaultColor.R())),
					uint8(viz.ScaledColorValue(c.DefaultColor.G())),
					uint8(viz.ScaledColorValue(c.DefaultColor.B())),
					0,
				},
			)
		}
	}
	return img
}

func (c *BasicCast) Height() int {
	return c.H
}

func (c *BasicCast) Width() int {
	return c.W
}

func (c *BasicCast) Len() int {
	return 1
}
