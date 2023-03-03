package viz

import "happymonday.dev/ray-tracer/src/tuples"

type Color struct {
	*tuples.Tuple
}

func InitColor(r, g, b float64) Color {
	return Color{&tuples.Tuple{X: r, Y: g, Z: b, W: 0}}
}

func (c Color) R() float64 {
	return c.X
}

func (c Color) G() float64 {
	return c.Y
}

func (c Color) B() float64 {
	return c.Z
}

func (c Color) Add(c2 Color) Color {
	return Color{c.Tuple.Add(c2.Tuple)}
}

func (c Color) Subtract(c2 Color) Color {
	return Color{c.Tuple.Subtract(c2.Tuple)}
}

func (c Color) MultiplyScalar(s float64) Color {
	return Color{c.Tuple.MultiplyScalar(s)}
}

func (c Color) Multiply(c2 Color) Color {
	return Color{&tuples.Tuple{X: c.R() * c2.R(), Y: c.G() * c2.G(), Z: c.B() * c2.B(), W: 0}}
}

func (c Color) Eq(c2 Color) bool {
	return c.Tuple.Eq(c2.Tuple)
}
