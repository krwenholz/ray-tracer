package clock

import (
	"image"
	"image/color"
	"image/color/palette"
	"log"

	"happymonday.dev/ray-tracer/src/matrix"
	"happymonday.dev/ray-tracer/src/tuples"
	"happymonday.dev/ray-tracer/src/viz"
)

var tock *matrix.Matrix = matrix.RotationZ(-1 / 12.0)

type Clock struct {
	Positions    []*tuples.Tuple
	H            int
	W            int
	DefaultColor viz.Color
}

func Init(h, w int, c viz.Color) *Clock {
	return &Clock{
		Positions: []*tuples.Tuple{
			tuples.InitPoint(0, 1, 0), // 1200
		},
		H:            h,
		W:            w,
		DefaultColor: c,
	}
}

func (c *Clock) Tick() {
	prev := c.Positions[len(c.Positions)-1]
	c.Positions = append(c.Positions, tock.MultiplyTuple(prev))
	log.Println(c.Positions[len(c.Positions)-1])
}

func (c *Clock) DrawRGBA(t int) *image.Paletted {
	p := c.Positions[t]
	p = matrix.Chain(matrix.Scaling(float64(c.W)/3, float64(c.H)/3, 0), matrix.Translation(float64(c.W)/2, float64(c.H)/2, 0)).MultiplyTuple(p)
	x := int(p.X)
	y := int(p.Y)
	img := image.NewPaletted(image.Rect(0, 0, int(c.W), int(c.H)), palette.Plan9)
	for oy := 0; oy < 6; oy++ {
		for ox := 0; ox < 6; ox++ {
			img.Set(
				x+ox,
				int(c.H)-(y+oy),
				color.RGBA{
					uint8(viz.ScaledColorValue256(c.DefaultColor.R())),
					uint8(viz.ScaledColorValue256(c.DefaultColor.G())),
					uint8(viz.ScaledColorValue256(c.DefaultColor.B())),
					0,
				},
			)
		}
	}
	return img
}

func (c *Clock) Height() int {
	return c.H
}

func (c *Clock) Width() int {
	return c.W
}

func (c *Clock) Len() int {
	return len(c.Positions)
}
