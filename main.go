package main

import (
	"context"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"net/http"
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
	log.Println("Authorized github email", os.Getenv("EMAIL"))
	tun, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(
			config.WithOAuth("github", config.WithAllowOAuthEmail(os.Getenv("EMAIL"))),
			config.WithDomain(os.Getenv("DOMAIN")),
		),
		ngrok.WithAuthtokenFromEnv(),
	)
	r := gin.Default()
	r.Run()
	if err != nil {
		return err
	}

	log.Println("tunnel created:", tun.URL())

	http.HandleFunc("/main.go", getMainGo)
	http.HandleFunc("/projectile", simulateProjectile)
	http.HandleFunc("/clock", simulateClock)
	http.HandleFunc("/basic_cast", basicRayCast)
	http.HandleFunc("/basic_3d", threeDRayCast)
	http.HandleFunc("/basic_3d_light", threeDRayCastLightMoves)
	http.HandleFunc("/basic_3d_jpeg", threeDRayCastLightJpeg)
	return http.Serve(tun, nil)
}

func getMainGo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./main.go")
}

func simulateProjectile(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting simulation")
	scene := &projectile.Scene{
		ProjectileSnapshots: []projectile.Projectile{{Pos: tuples.InitPoint(0, 1, 0), Velocity: tuples.InitVector(1, 1.8, 0).Normalize().MultiplyScalar(11.25)}},
		E:                   projectile.Environment{Gravity: tuples.InitVector(0, -0.1, 0), Wind: tuples.InitVector(-0.01, 0, 0)},
		DefaultColor:        viz.InitColor(255, 255, 0),
	}
	for i := 0; i < 200; i++ {
		scene.Tick()
	}
	viz.EncodeGIF(w, viz.DrawAllRGBA(scene), 10)
	log.Println("Done simulating")
}

func simulateClock(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting simulation")
	c := clock.Init(200, 200, viz.InitColor(255, 255, 0))
	viz.EncodeGIF(w, viz.DrawAllRGBA(c), 50)
	log.Println("Done simulating")
}

func basicRayCast(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting basic cast")
	s := shapes.InitSphere()
	s.SetTransform(
		matrix.Chain(
			matrix.Scaling(2, 2, 2),
			matrix.Translation(100, 100, 0),
		),
	)
	c := basic_ray_cast.Init(200, 200, viz.InitColor(255, 255, 0), s)
	viz.EncodeGIF(
		w,
		[]*image.Paletted{
			c.Shine(tuples.InitPoint(80, 80, -30)),
			c.Shine(tuples.InitPoint(90, 90, -30)),
			c.Shine(tuples.InitPoint(100, 100, -30)),
			c.Shine(tuples.InitPoint(110, 110, -30)),
			c.Shine(tuples.InitPoint(120, 120, -30)),
		},
		50,
	)
	log.Println("Done simulating")
}

func threeDRayCast(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting 3D ray cast")
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
	c := three_d_ray_cast.Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	imgs := []*image.Paletted{}
	steps := 3.0
	for i := 0.0; i < steps; i++ {
		location := tuples.InitPoint(size/steps*i, size/steps*i, -size/4)
		imgs = append(imgs, c.Shine(location))
	}
	viz.EncodeGIF(
		w,
		imgs,
		50,
	)
	log.Println("Done simulating")
}

func threeDRayCastLightMoves(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting 3D ray cast")
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
	c := three_d_ray_cast.Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	step := 5.0
	imgs := []*image.Paletted{}
	for i := 0.0; i < 5; i++ {
		imgs = append(imgs, c.Shine(tuples.InitPoint(size/step*i, size/step*i, -size/4)))
	}
	viz.EncodeGIF(
		w,
		imgs,
		50,
	)
	log.Println("Done simulating")
}

func threeDRayCastLightJpeg(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting 3D jpeg ray cast")
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
	c := three_d_ray_cast.Init(int(size), int(size), viz.InitColor(255, 255, 0), s, l)
	img := c.Shine(tuples.InitPoint(size/3, size/2, -size/4))
	jpeg.Encode(
		w,
		img,
		nil,
	)
	log.Println("Done")
}

func writeTemp(s string, t string) string {
	f, err := os.CreateTemp("", fmt.Sprintf("ray-tracer-*.%s", t))
	if err != nil {
		log.Fatal(err)
	}

	// close and remove the temporary file at the end of the program
	defer f.Close()

	// write data to the temporary file
	data := []byte(s)
	if _, err := f.Write(data); err != nil {
		log.Fatal(err)
	}

	log.Print("tmp file written ", f.Name())
	return f.Name()
}
