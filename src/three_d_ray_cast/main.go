package three_d_ray_cast

import (
	"image"
	"image/color"
	"image/jpeg"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
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
	DefaultColor *viz.Color
	Sphere       *shapes.Sphere
	Light        *lights.PointLight
}

func Init(h, w int, c *viz.Color, s *shapes.Sphere, l *lights.PointLight) *BasicCast {
	b := BasicCast{
		H:            h,
		W:            w,
		DefaultColor: c,
		Sphere:       s,
		Light:        l,
	}
	return &b
}

func (c *BasicCast) DrawRGBA(source *tuples.Tuple) image.RGBA64Image {
	img := image.NewRGBA64(image.Rect(0, 0, int(c.W), int(c.H)))
	wg := sync.WaitGroup{}
	wg.Add(c.Height() * c.Width())
	bar := progressbar.Default(int64(c.Height() * c.Width()))
	bar.Describe("Casting")

	for iy := 0; iy < c.Height(); iy++ {
		for ix := 0; ix < c.Width(); ix++ {
			y := iy
			x := ix
			go func() {
				hitColor := c.HitColor(source, x, y)
				if hitColor != nil {
					img.Set(
						x,
						y,
						color.RGBA{
							uint8(hitColor.RRGBA()),
							uint8(hitColor.GRGBA()),
							uint8(hitColor.BRGBA()),
							0xff,
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

func (c *BasicCast) HitColor(source *tuples.Tuple, x, y int) *viz.Color {
	dir := tuples.InitVectorFromPoints(tuples.InitPoint(float64(x), float64(y), 100), source).Normalize()
	r := shapes.InitRay(source, dir)
	h := c.Sphere.Intersect(r).Hit()
	if h == nil {
		return nil
	}
	point := r.Position(h.T)
	normal := c.Sphere.NormalAt(point)
	eye := r.Direction.Negate()
	return c.Light.Lighting(c.Sphere.Material, point, eye, normal)
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

func ThreeDRayCast(c *gin.Context) {
	size := 500.0
	s := shapes.InitSphere()
	s.Material.Color = viz.InitColor(1, 0.2, 1)
	s.SetTransform(
		matrix.Chain(
			matrix.Scaling(size/7, size/7, size/7),
			matrix.Translation(size/2, size/2, size/5),
		),
	)
	color := viz.InitColor(1, 1, 1)
	l := lights.InitPointLight(tuples.InitPoint(size/2, size/2, -size/5), color)
	rc := Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	imgs := []image.RGBA64Image{}
	steps := 10.0
	for i := 0.0; i < steps; i++ {
		location := tuples.InitPoint(size/steps*i, size/steps*i, -size/4)
		imgs = append(imgs, rc.DrawRGBA(location))
	}
	data := viz.EncodeX264FromRBA64(int(size), int(size), 5, imgs)
	c.Header("Content-Disposition", `attachment; filename=3d_ray_cast.264`)
	c.Data(
		http.StatusOK,
		"video/H264",
		data.Bytes(),
	)
}

func ThreeDRayCastLightMoves(c *gin.Context) {
	size := 500.0
	s := shapes.InitSphere()
	s.Material.Color = viz.InitColor(1, 0.2, 1)
	s.SetTransform(
		matrix.Chain(
			matrix.Scaling(size/7, size/7, size/7),
			matrix.Translation(size/2, size/2, 5),
		),
	)
	color := viz.InitColor(1, 1, 1)
	l := lights.InitPointLight(tuples.InitPoint(size/2, size/2, -size/5), color)
	rc := Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	steps := 10.0
	imgs := []image.RGBA64Image{}
	location := tuples.InitPoint(size/2, size/2, -size/4)
	for i := 0.0; i < steps; i++ {
		ll := tuples.InitPoint(size/steps*i, size/steps*i, -size/4)
		rc.Light = lights.InitPointLight(ll, color)
		imgs = append(imgs, rc.DrawRGBA(location))
	}
	data := viz.EncodeX264FromRBA64(int(size), int(size), 5, imgs)
	c.Header("Content-Disposition", `attachment; filename=3d_ray_cast_light_moves.264`)
	c.Data(
		http.StatusOK,
		"video/H264",
		data.Bytes(),
	)
}

func ThreeDRayCastLightJpeg(c *gin.Context) {
	size := 500.0
	s := shapes.InitSphere()
	s.Material.Color = viz.InitColor(1, 0.2, 1)
	s.SetTransform(
		matrix.Chain(
			matrix.Scaling(size/7, size/7, size/7),
			matrix.Translation(size/2, size/2, 5),
		),
	)
	lightColor := viz.InitColor(1, 1, 1)
	light := lights.InitPointLight(tuples.InitPoint(-size/5, size/2, -size/5), lightColor)
	rc := Init(int(size), int(size), viz.InitColor(255, 255, 0), s, light)
	img := rc.DrawRGBA(tuples.InitPoint(size/3, size/2, -size/4))
	jpeg.Encode(
		c.Writer,
		img,
		nil,
	)
}
