package simpler_world

import (
	"fmt"
	"image/jpeg"
	"math"

	"github.com/gin-gonic/gin"
	"happymonday.dev/ray-tracer/src/lights"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
	"happymonday.dev/ray-tracer/src/world"
)

func worldOne() *world.World {
	l := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.Material().Color = viz.InitColor(0.8, 1.0, 0.6)
	s1.Material().Diffuse = 0.7
	s1.Material().Specular = 0.2
	s2 := shapes.InitSphere()
	s2.Material().Color = viz.InitColor(0.8, 0.5, 0.6)
	s2.Material().Diffuse = 0.7
	s2.Material().Specular = 0.2
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5).Multiply(matrix.Translation(5, 0, 0)))
	return &world.World{Objects: []shapes.Shape{s1, s2}, Lights: []*lights.PointLight{l}}
}

func worldOneBacked() *world.World {
	l := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.Material().Color = viz.InitColor(0.8, 1.0, 0.6)
	s1.Material().Diffuse = 0.7
	s1.Material().Specular = 0.2

	s2 := shapes.InitSphere()
	s2.Material().Color = viz.InitColor(0.8, 0.5, 0.6)
	s2.Material().Diffuse = 0.7
	s2.Material().Specular = 0.2
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5).Multiply(matrix.Translation(5, 0, 0)))

	sbx := shapes.InitSphere()
	sbx.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sbx.Material().Diffuse = 0.7
	sbx.Material().Specular = 0.2
	sbx.SetTransform(matrix.Chain(matrix.Scaling(0.01, 100, 100), matrix.Translation(-10, 0, 0)))

	sby := shapes.InitSphere()
	sby.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sby.Material().Diffuse = 0.7
	sby.Material().Specular = 0.2
	sby.SetTransform(matrix.Chain(matrix.Scaling(100, 0.01, 100), matrix.Translation(0, -10, 0)))

	sbz := shapes.InitSphere()
	sbz.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sbz.Material().Diffuse = 0.7
	sbz.Material().Specular = 0.2
	sbz.SetTransform(matrix.Chain(matrix.Scaling(100, 100, 0.01), matrix.Translation(0, 0, -10)))

	return &world.World{Objects: []shapes.Shape{s1, s2, sbx, sby, sbz}, Lights: []*lights.PointLight{l}}
}

func worldFull() *world.World {
	l1 := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))
	l2 := lights.InitPointLight(tuples.InitPoint(-9.5, 10, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.Material().Color = viz.InitColor(0.8, 1.0, 0.6)
	s1.Material().Diffuse = 0.7
	s1.Material().Specular = 0.2

	s2 := shapes.InitSphere()
	s2.Material().Color = viz.InitColor(0.8, 0.5, 0.6)
	s2.Material().Diffuse = 0.7
	s2.Material().Specular = 0.2
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5).Multiply(matrix.Translation(5, 0, 0)))

	s3 := shapes.InitSphere()
	s3.Material().Color = viz.InitColor(0.1, 0.1, 0.8)
	s3.Material().Diffuse = 0.7
	s3.Material().Specular = 0.2
	s3.SetTransform(matrix.Scaling(0.5, 2.0, 0.5).Multiply(matrix.Translation(-2, 0, 2)))

	sbx := shapes.InitSphere()
	sbx.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sbx.Material().Diffuse = 0.7
	sbx.Material().Specular = 0.2
	sbx.SetTransform(matrix.Chain(matrix.Scaling(0.01, 100, 100), matrix.Translation(-20, 0, 0)))

	sby := shapes.InitSphere()
	sby.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sby.Material().Diffuse = 0.7
	sby.Material().Specular = 0.2
	sby.SetTransform(matrix.Chain(matrix.Scaling(100, 0.01, 100), matrix.Translation(0, -10, 0)))

	sbz := shapes.InitSphere()
	sbz.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sbz.Material().Diffuse = 0.7
	sbz.Material().Specular = 0.2
	sbz.SetTransform(matrix.Chain(matrix.Scaling(100, 100, 0.01), matrix.Translation(0, 0, 10)))

	return &world.World{Objects: []shapes.Shape{s1, s2, s3, sbx, sby, sbz}, Lights: []*lights.PointLight{l1, l2}}
}

func worldHeart() *world.World {
	l := lights.InitPointLight(tuples.InitPoint(-0, 2, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.Material().Color = viz.InitColor(1.0, 0.0, 0.0)
	s1.Material().Diffuse = 0.7
	s1.Material().Specular = 0.2
	s1.SetTransform(
		matrix.Chain(
			matrix.Shearing(1, -1, 0, 0, 0, 0),
			matrix.RotationZ(1.0/4.0),
			matrix.Translation(0.5, 0, 0),
		))

	s2 := shapes.InitSphere()
	s2.Material().Color = viz.InitColor(1.0, 0.0, 0.0)
	s2.Material().Diffuse = 0.7
	s2.Material().Specular = 0.2
	s2.SetTransform(
		matrix.Chain(
			matrix.Shearing(-1, 1, 0, 0, 0, 0),
			matrix.RotationZ(-1.0/4.0),
			matrix.Translation(-0.5, 0, 0),
		))

	sbx := shapes.InitSphere()
	sbx.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sbx.Material().Diffuse = 0.7
	sbx.Material().Specular = 0.2
	sbx.SetTransform(matrix.Chain(matrix.Scaling(0.01, 100, 100), matrix.Translation(-10, 0, 0)))

	sby := shapes.InitSphere()
	sby.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sby.Material().Diffuse = 0.7
	sby.Material().Specular = 0.2
	sby.SetTransform(matrix.Chain(matrix.Scaling(100, 0.01, 100), matrix.Translation(0, -10, 0)))

	sbz := shapes.InitSphere()
	sbz.Material().Color = viz.InitColor(0.9, 0.9, 1.0)
	sbz.Material().Diffuse = 0.7
	sbz.Material().Specular = 0.2
	sbz.SetTransform(matrix.Chain(matrix.Scaling(100, 100, 0.01), matrix.Translation(0, 0, 10)))

	return &world.World{Objects: []shapes.Shape{s1, s2, sbx, sby, sbz}, Lights: []*lights.PointLight{l}}
}

func worldBook() *world.World {
	l := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))

	floor := shapes.InitSphere()
	floor.Material().Color = viz.InitColor(1.0, 0.9, 0.9)
	floor.Material().Specular = 0.0
	floor.SetTransform(
		matrix.Chain(
			matrix.Scaling(10, 0.01, 10),
		))

	left_wall := shapes.InitSphere()
	left_wall.SetMaterial(floor.Material())
	left_wall.SetTransform(
		matrix.Chain(
			matrix.Scaling(10, 0.01, 10),
			matrix.RotationX(-1.0/2.0),
			matrix.RotationY(-1.0/4.0),
			matrix.Translation(-0.0, 0, 5),
		))

	right_wall := shapes.InitSphere()
	right_wall.SetMaterial(floor.Material())
	right_wall.SetTransform(
		matrix.Chain(
			matrix.Scaling(10, 0.01, 10),
			matrix.RotationX(-1.0/2.0),
			matrix.RotationY(1.0/4.0),
			matrix.Translation(0.0, 0, 5),
		))

	middle := shapes.InitSphere()
	middle.Material().Color = viz.InitColor(0.1, 1.0, 0.5)
	middle.Material().Diffuse = 0.7
	middle.Material().Specular = 0.3
	middle.SetTransform(
		matrix.Chain(
			matrix.Translation(-0.5, 1, 0.5),
		))

	right := shapes.InitSphere()
	right.Material().Color = viz.InitColor(0.5, 1.0, 0.1)
	right.Material().Diffuse = 0.7
	right.Material().Specular = 0.3
	right.SetTransform(
		matrix.Chain(
			matrix.Scaling(0.5, 0.5, 0.5),
			matrix.Translation(1.5, 0.5, -0.5),
		))

	left := shapes.InitSphere()
	left.Material().Color = viz.InitColor(1.0, 0.8, 0.1)
	left.Material().Diffuse = 0.7
	left.Material().Specular = 0.3
	left.SetTransform(
		matrix.Chain(
			matrix.Scaling(0.33, 0.33, 0.33),
			matrix.Translation(-1.5, 0.33, -0.75),
		))

	return &world.World{Objects: []shapes.Shape{floor, left_wall, right_wall, middle, left, right}, Lights: []*lights.PointLight{l}}
}

func SimplerWorld(ctx *gin.Context) {
	s := 500
	c := world.InitCamera(s, s, math.Pi/2.0)
	from := tuples.InitPoint(0, 0, -5)
	to := tuples.InitPoint(0, 0, 0)
	up := tuples.InitVector(0, 1, 0)
	var w *world.World
	switch wName := ctx.Query("world"); wName {
	case "one":
		fmt.Println("Displaying world one")
		w = worldOne()
	case "one_backed":
		fmt.Println("Displaying world one backed")
		w = worldOneBacked()
	case "full":
		fmt.Println("Displaying world full")
		w = worldFull()
	case "heart":
		fmt.Println("Displaying world heart")
		w = worldHeart()
	case "book":
		fmt.Println("Displaying world book")
		w = worldBook()
		c = world.InitCamera(s, s, math.Pi/3.0)
		from = tuples.InitPoint(0, 1.5, -5)
		to = tuples.InitPoint(0, 1, 0)
	default:
		fmt.Println("Displaying world default")
		w = world.InitDefaultWorld()
	}
	c.SetTransform(world.ViewTransformation(from, to, up))
	img := c.Render(w)
	jpeg.Encode(
		ctx.Writer,
		img.DrawRGBA(),
		nil,
	)
}
