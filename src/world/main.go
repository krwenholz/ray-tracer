package world

import (
	"sync"

	"happymonday.dev/ray-tracer/src/lights"
	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/shapes"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

type World struct {
	Objects []shapes.Object
	Lights  []*lights.PointLight
}

func InitWorld() *World {
	return &World{}
}

func InitDefaultWorld() *World {
	l := lights.InitPointLight(tuples.InitPoint(-10, 10, -10), viz.InitColor(1, 1, 1))
	s1 := shapes.InitSphere()
	s1.Material().Color = viz.InitColor(0.8, 1.0, 0.6)
	s1.Material().Diffuse = 0.7
	s1.Material().Specular = 0.2
	s2 := shapes.InitSphere()
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))
	return &World{[]shapes.Object{s1, s2}, []*lights.PointLight{l}}
}

func (w *World) Intersections(r *shapes.Ray) *shapes.Intersections {
	ress := make(chan []*shapes.Intersection)
	wg := sync.WaitGroup{}
	wg.Add(len(w.Objects))
	for _, o := range w.Objects {
		obj := o
		go func() {
			ress <- obj.Intersect(r).Intersections
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ress)
	}()
	xs := []*shapes.Intersection{}
	for res := range ress {
		xs = append(xs, res...)
	}
	return shapes.InitIntersections(xs...)
}

func (w *World) ShadeHit(c *shapes.IntersectionComputations) *viz.Color {
	res := viz.Black()
	for _, l := range w.Lights {
		res = res.Add(l.Lighting(c.Object.Material(), c.Point, c.EyeV, c.NormalV))
	}
	return res
}

func (w *World) ColorAt(r *shapes.Ray) *viz.Color {
	is := w.Intersections(r)
	h := is.Hit()
	if h == nil {
		return viz.Black()
	}
	return w.ShadeHit(h.PrepareComputations(r))
}
