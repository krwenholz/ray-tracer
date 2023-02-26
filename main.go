package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
	"happymonday.dev/ray-tracer/src/projectile"
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
	if err != nil {
		return err
	}

	log.Println("tunnel created:", tun.URL())

	http.HandleFunc("/main.go", getMainGo)
	http.HandleFunc("/simulateProjectile", simulateProjectile)
	return http.Serve(tun, nil)
}

func getMainGo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./main.go")
}

func simulateProjectile(w http.ResponseWriter, r *http.Request) {
	log.Println("Starting simulation")
	scene := projectile.Scene{
		ProjectileSnapshots: []projectile.Projectile{{Pos: tuples.InitPoint(0, 1, 0), Velocity: tuples.InitVector(1, 1.8, 0).Normalize().MultiplyScalar(11.25)}},
		E:                   projectile.Environment{Gravity: tuples.InitVector(0, -0.1, 0), Wind: tuples.InitVector(-0.01, 0, 0)},
		MaxHeight:           5,
		MaxWidth:            5,
		DefaultColor:        viz.InitColor(255, 255, 0),
	}
	for i := 0; i < 200; i++ {
		scene.Tick()
	}
	viz.EncodeGIF(w, scene.DrawAllRGBA())
	log.Println("Done simulating")
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
