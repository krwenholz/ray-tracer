package viz

import (
	"image"
	"image/color"
	"math"
	"sync"

	"github.com/schollz/progressbar/v3"
)

type Canvas struct {
	Height int
	Width  int
	pixels [][]*Color
}

func InitCanvas(w, h int) Canvas {
	ps := make([][]*Color, h)
	for i := 0; i < h; i++ {
		ps[i] = make([]*Color, w)
		for j := 0; j < w; j++ {
			ps[i][j] = InitColor(0, 0, 0)
		}
	}
	return Canvas{Height: h, Width: w, pixels: ps}
}

func (c *Canvas) Pixel(x, y int) *Color {
	return c.pixels[y][x]
}

func (c *Canvas) SetPixel(col *Color, x, y int) {
	c.pixels[y][x] = col
}

func ScaledColorValue256(v float64) int {
	return int(math.Round(math.Max(math.Min(v*255, 255), 0)))
}

func ScaledColorValueRGBA(v float64) int {
	return int(math.Round(math.Max(math.Min(v*0xff, 0xff), 0)))
}

func (c *Canvas) DrawRGBA() image.RGBA64Image {
	img := image.NewRGBA64(image.Rect(0, 0, int(c.Width), int(c.Height)))
	wg := sync.WaitGroup{}
	size := int(c.Height * c.Width)
	wg.Add(size)
	bar := progressbar.Default(int64(size))
	bar.Describe("Drawing canvas")

	for iy := 0; iy < c.Height; iy++ {
		for ix := 0; ix < c.Width; ix++ {
			y := iy
			x := ix
			go func() {
				pc := c.Pixel(x, y)
				img.Set(
					x,
					y,
					color.RGBA{
						uint8(pc.RRGBA()),
						uint8(pc.GRGBA()),
						uint8(pc.BRGBA()),
						0xff,
					},
				)
				bar.Add(1)
				wg.Done()
			}()
		}
	}
	wg.Wait()
	return img
}
