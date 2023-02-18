package space

type Tuple interface {
	X() float64
	Y() float64
	Z() float64
	W() float64
}

type Point struct {
	x float64
	y float64
	z float64
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func (p *Point) Z() float64 {
	return p.z
}

func (p *Point) W() float64 {
	return 1
}

type Vector struct {
	x float64
	y float64
	z float64
}

func (v *Vector) X() float64 {
	return v.x
}

func (v *Vector) Y() float64 {
	return v.y
}

func (v *Vector) Z() float64 {
	return v.z
}

func (v *Vector) W() float64 {
	return 0
}
