package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"happymonday.dev/ray-tracer/src/basic_ray_cast"
	"happymonday.dev/ray-tracer/src/clock"
	"happymonday.dev/ray-tracer/src/projectile"
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
	r.StaticFile("favicon.ico", "static/favicon.ico")
	r.GET("/main.go", handleMainGo)
	r.GET("/projectile", handleProjectile)
	r.GET("/clock", handleClock)
	r.GET("/basic_cast", basic_ray_cast.BasicRayCast)
	r.GET("/basic_3d", three_d_ray_cast.ThreeDRayCast)
	r.GET("/basic_3d_light", three_d_ray_cast.ThreeDRayCastLightMoves)
	r.GET("/basic_3d_jpeg", three_d_ray_cast.ThreeDRayCastLightJpeg)
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
