package projectile

import (
	"image"
	"image/color"
	"image/color/palette"
	"log"
	"math"
	"sync"

	progressbar "github.com/schollz/progressbar/v3"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

type Projectile struct {
	Pos      *tuples.Tuple
	Velocity *tuples.Tuple
}

type Environment struct {
	Gravity *tuples.Tuple
	Wind    *tuples.Tuple
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

func (s *Scene) DrawRGBA(t int) *image.Paletted {
	p := s.ProjectileSnapshots[t]
	x := int(p.Pos.X)
	y := int(p.Pos.Y)
	img := image.NewPaletted(image.Rect(0, 0, int(s.MaxWidth), int(s.MaxHeight)), palette.Plan9)
	for oy := 0; oy < 6; oy++ {
		for ox := 0; ox < 6; ox++ {
			img.Set(
				x+ox,
				int(s.MaxHeight)-(y+oy),
				color.RGBA{
					uint8(viz.ScaledColorValue(s.DefaultColor.R())),
					uint8(viz.ScaledColorValue(s.DefaultColor.G())),
					uint8(viz.ScaledColorValue(s.DefaultColor.B())),
					0,
				},
			)
		}
	}
	return img
}

func (s *Scene) DrawLast() viz.Canvas {
	return s.Draw(len(s.ProjectileSnapshots) - 1)
}

func (s *Scene) DrawAll() []viz.Canvas {
	res := []viz.Canvas{}
	cs := sync.Map{}

	wg := sync.WaitGroup{}
	wg.Add(len(s.ProjectileSnapshots))
	bar := progressbar.Default(int64(len(s.ProjectileSnapshots)))
	bar.Describe("Drawing scene")

	for i := range s.ProjectileSnapshots {
		t := i
		go func() {
			cs.Store(t, s.Draw(t))
			bar.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := range s.ProjectileSnapshots {
		v, ok := cs.Load(i)
		if ok {
			if c, ok := v.(viz.Canvas); ok {
				res = append(res, c)
			}
		}
		if !ok {
			log.Fatal("Map read failed for snapshot", i)
		}
	}
	return res
}

func (s *Scene) DrawAllRGBA() []*image.Paletted {
	res := []*image.Paletted{}
	imgs := sync.Map{}

	wg := sync.WaitGroup{}
	wg.Add(len(s.ProjectileSnapshots))
	bar := progressbar.Default(int64(len(s.ProjectileSnapshots)))
	bar.Describe("Drawing scene")

	for i := range s.ProjectileSnapshots {
		t := i
		go func() {
			imgs.Store(t, s.DrawRGBA(t))
			bar.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := range s.ProjectileSnapshots {
		v, ok := imgs.Load(i)
		if ok {
			if img, ok := v.(*image.Paletted); ok {
				res = append(res, img)
			}
		}
		if !ok {
			log.Fatal("Map read failed for snapshot", i)
		}
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
