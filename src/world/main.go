package world

import (
	"happymonday.dev/ray-tracer/src/lights"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

type World struct {
	Objects []shapes.Object
	Light   *lights.PointLight
}

func InitWorld() *World {
	return &World{}
}

func InitDefaultWorld() *World {
	l := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.SetMaterial(shapes.InitMaterial(viz.InitColor(0.8, 1.0, 0.6), 0, 0.7, 0.2, 0))
	s2 := shapes.InitSphere()
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))
	return &World{[]shapes.Object{s1, s2}, l}
}
