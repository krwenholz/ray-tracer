package viz

import "math"

type Canvas struct {
	Height int
	Width  int
	pixels [][]Color
}

func InitCanvas(w, h int) Canvas {
	ps := make([][]Color, h)
	for i := 0; i < h; i++ {
		ps[i] = make([]Color, w)
		for j := 0; j < w; j++ {
			ps[i][j] = InitColor(0, 0, 0)
		}
	}
	return Canvas{Height: h, Width: w, pixels: ps}
}

func (c *Canvas) Pixel(x, y int) Color {
	return c.pixels[y][x]
}

func (c *Canvas) SetPixel(col Color, x, y int) {
	c.pixels[y][x] = col
}

func ScaledColorValue256(v float64) int {
	return int(math.Round(math.Max(math.Min(v*255, 255), 0)))
}

func ScaledColorValueRGBA(v float64) int {
	return int(math.Round(math.Max(math.Min(v*0xff, 0xff), 0)))
}
