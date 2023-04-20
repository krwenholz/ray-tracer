package world

import (
	"happymonday.dev/ray-tracer/src/lights"
	"happymonday.dev/ray-tracer/src/shapes"
)

type World struct {
	Objects []*shapes.Intersectable
	Light   lights.PointLight
}

func InitWorld() *World {
	return &World{}
}
