package basic_ray_cast

import (
	"image"
	"image/color"
	"image/color/palette"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/schollz/progressbar/v3"
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
}

func Init(h, w int, c viz.Color, s *shapes.Sphere) *BasicCast {
	b := BasicCast{
		H:            h,
		W:            w,
		DefaultColor: c,
		Sphere:       s,
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
				if c.RayHits(source, x, y) {
					img.Set(
						x,
						y,
						color.RGBA{
							uint8(viz.ScaledColorValue256(c.DefaultColor.R())),
							uint8(viz.ScaledColorValue256(c.DefaultColor.G())),
							uint8(viz.ScaledColorValue256(c.DefaultColor.B())),
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

func (c *BasicCast) RayHits(source *tuples.Tuple, x, y int) bool {
	v := tuples.InitVectorFromPoints(tuples.InitPoint(float64(x), float64(y), 100), source)
	r := shapes.InitRay(source, v)
	return c.Sphere.Intersect(r).Hit() != nil
}

func (c *BasicCast) DrawRGBA() *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, int(c.W), int(c.H)), palette.Plan9)
	for y := 0; y < c.Height(); y++ {
		for x := 0; x < c.Width(); x++ {
			img.Set(
				x,
				c.Height()-y,
				color.RGBA{
					uint8(viz.ScaledColorValue256(c.DefaultColor.R())),
					uint8(viz.ScaledColorValue256(c.DefaultColor.G())),
					uint8(viz.ScaledColorValue256(c.DefaultColor.B())),
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

func BasicRayCast(c *gin.Context) {
	s := shapes.InitSphere()
	s.SetTransform(
		matrix.Chain(
			matrix.Scaling(2, 2, 2),
			matrix.Translation(100, 100, 0),
		),
	)
	scene := Init(200, 200, viz.InitColor(255, 255, 0), s)
	viz.EncodeGIF(
		c.Writer,
		[]*image.Paletted{
			scene.Shine(tuples.InitPoint(80, 80, -30)),
			scene.Shine(tuples.InitPoint(90, 90, -30)),
			scene.Shine(tuples.InitPoint(100, 100, -30)),
			scene.Shine(tuples.InitPoint(110, 110, -30)),
			scene.Shine(tuples.InitPoint(120, 120, -30)),
		},
		50,
	)
}
