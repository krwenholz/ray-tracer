package projectile

import (
	"image"
	"image/color"
	"image/color/palette"
	"math"

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

func (s *Scene) DrawRGBA(t int) *image.Paletted {
	p := s.ProjectileSnapshots[t]
	x := int(p.Pos.X)
	y := int(p.Pos.Y)
	img := image.NewPaletted(image.Rect(0, 0, s.Width(), s.Height()), palette.Plan9)
	for oy := 0; oy < 6; oy++ {
		for ox := 0; ox < 6; ox++ {
			img.Set(
				x+ox,
				s.Height()-(y+oy),
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

func (s *Scene) Height() int {
	return int(s.MaxHeight)
}

func (s *Scene) Width() int {
	return int(s.MaxWidth)
}

func (s *Scene) Len() int {
	return len(s.ProjectileSnapshots)
}
