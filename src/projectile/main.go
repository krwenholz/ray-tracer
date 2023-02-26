package projectile

import (
	"math"
	"sync"

	progressbar "github.com/schollz/progressbar/v3"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

type Projectile struct {
	Pos      tuples.Tuple
	Velocity tuples.Tuple
}

type Environment struct {
	Gravity tuples.Tuple
	Wind    tuples.Tuple
}

type Scene struct {
	ProjectileSnapshots []Projectile
	E                   Environment
	MaxHeight           float64
	MaxWidth            float64
	DefaultColor        viz.Color
}

func (s *Scene) Tick() {
	prev := s.ProjectileSnapshots[len(s.ProjectileSnapshots)-1]
	cur := Projectile{}
	cur.Pos = prev.Pos.Add(prev.Velocity)
	if cur.Pos.X < 0 {
		cur.Pos.X = 0
	}
	if cur.Pos.Y < 0 {
		cur.Pos.Y = 0
	}

	cur.Velocity = prev.Velocity.Add(s.E.Gravity).Add(s.E.Wind)
	s.ProjectileSnapshots = append(s.ProjectileSnapshots, cur)
	s.MaxHeight = math.Max(s.MaxHeight, cur.Pos.Y)
	s.MaxWidth = math.Max(s.MaxWidth, cur.Pos.X)
}

func (s *Scene) Draw(t int) viz.Canvas {
	c := viz.InitCanvas(int(s.MaxWidth)+1, int(s.MaxHeight)+1)
	p := s.ProjectileSnapshots[t]
	c.SetPixel(s.DefaultColor, int(p.Pos.X), int(s.MaxHeight-p.Pos.Y))
	return c
}

func (s *Scene) DrawLast() viz.Canvas {
	return s.Draw(len(s.ProjectileSnapshots) - 1)
}

func (s *Scene) DrawAll() []viz.Canvas {
	res := []viz.Canvas{}
	cs := make(map[int]viz.Canvas)

	wg := sync.WaitGroup{}
	wg.Add(len(s.ProjectileSnapshots))
	bar := progressbar.Default(int64(len(s.ProjectileSnapshots)))
	bar.Describe("Drawing canvases")

	for i := range s.ProjectileSnapshots {
		t := i
		go func() {
			cs[t] = s.Draw(t)
			bar.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()
	for i := range s.ProjectileSnapshots {
		res = append(res, cs[i])
	}
	return res
}

func (s *Scene) DrawAllCollapsed() viz.Canvas {
	c := viz.InitCanvas(int(s.MaxWidth)+1, int(s.MaxHeight)+1)
	for _, p := range s.ProjectileSnapshots {
		c.SetPixel(s.DefaultColor, int(p.Pos.X), int(s.MaxHeight-p.Pos.Y))
	}
	return c
}
