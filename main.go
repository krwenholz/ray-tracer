package main

import (
	"context"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"happymonday.dev/ray-tracer/src/basic_ray_cast"
	"happymonday.dev/ray-tracer/src/clock"
	"happymonday.dev/ray-tracer/src/lights"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/projectile"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/three_d_ray_cast"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithOAuth("github", config.WithAllowOAuthEmail(os.Getenv("EMAIL"))),
			config.WithDomain(os.Getenv("DOMAIN")),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	if err != nil {
		return err
	}

	log.Println("tunnel created:", tun.URL())

	r := gin.Default()
	r.Use(gin.ErrorLogger())
	r.GET("/main.go", handleMainGo)
	r.GET("/projectile", handleProjectile)
	r.GET("/clock", handleClock)
	r.GET("/basic_cast", basicRayCast)
	r.GET("/basic_3d", threeDRayCast)
	r.GET("/basic_3d_light", threeDRayCastLightMoves)
	r.GET("/basic_3d_jpeg", threeDRayCastLightJpeg)
	return r.RunListener(tun)
}

func handleMainGo(c *gin.Context) {
	c.File("./main.go")
}

func handleProjectile(c *gin.Context) {
	scene := &projectile.Scene{
		ProjectileSnapshots: []projectile.Projectile{{Pos: tuples.InitPoint(0, 1, 0), Velocity: tuples.InitVector(1, 1.8, 0).Normalize().MultiplyScalar(11.25)}},
		E:                   projectile.Environment{Gravity: tuples.InitVector(0, -0.1, 0), Wind: tuples.InitVector(-0.01, 0, 0)},
		DefaultColor:        viz.InitColor(255, 255, 0),
	}
	for i := 0; i < 200; i++ {
		scene.Tick()
	}
	viz.EncodeGIF(c.Writer, viz.DrawAllRGBA(scene), 10)
}

func handleClock(c *gin.Context) {
	sim := clock.Init(200, 200, viz.InitColor(255, 255, 0))
	viz.EncodeGIF(c.Writer, viz.DrawAllRGBA(sim), 50)
}

func basicRayCast(c *gin.Context) {
	s := shapes.InitSphere()
	s.SetTransform(
		matrix.Chain(
			matrix.Scaling(2, 2, 2),
			matrix.Translation(100, 100, 0),
		),
	)
	scene := basic_ray_cast.Init(200, 200, viz.InitColor(255, 255, 0), s)
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

func threeDRayCast(c *gin.Context) {
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
	l := lights.InitPointLight(tuples.InitPoint(size/2, size/2, -size/5), &color)
	rc := three_d_ray_cast.Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	imgs := []*image.Paletted{}
	steps := 3.0
	for i := 0.0; i < steps; i++ {
		location := tuples.InitPoint(size/steps*i, size/steps*i, -size/4)
		imgs = append(imgs, rc.Shine(location))
	}
	viz.EncodeGIF(
		c.Writer,
		imgs,
		50,
	)
}

func threeDRayCastLightMoves(c *gin.Context) {
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
	l := lights.InitPointLight(tuples.InitPoint(size/2, size/2, -size/5), &color)
	rc := three_d_ray_cast.Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	step := 5.0
	imgs := []*image.Paletted{}
	for i := 0.0; i < 5; i++ {
		imgs = append(imgs, rc.Shine(tuples.InitPoint(size/step*i, size/step*i, -size/4)))
	}
	viz.EncodeGIF(
		c.Writer,
		imgs,
		50,
	)
}

func threeDRayCastLightJpeg(c *gin.Context) {
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
	l := lights.InitPointLight(tuples.InitPoint(size/2, size/2, -size/5), &color)
	rc := three_d_ray_cast.Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	img := rc.Shine(tuples.InitPoint(size/3, size/2, -size/4))
	jpeg.Encode(
		c.Writer,
		img,
		nil,
	)
}
