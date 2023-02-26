package viz

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io"
	"sync"

	"github.com/schollz/progressbar/v3"
)

func EncodeGIF(w io.Writer, cs []Canvas) {
	mappedImgs := make(map[int]*image.Paletted)
	delays := []int{}
	disposals := []byte{}

	wg := sync.WaitGroup{}
	wg.Add(len(cs))
	bar := progressbar.Default(int64(len(cs)))
	bar.Describe("Encoding GIF")

	for i := range cs {
		t := i
		go func() {
			mappedImgs[t] = CanvasToRGBA(cs[t])
			wg.Done()
			bar.Add(1)
		}()
	}
	wg.Wait()

	imgs := []*image.Paletted{}
	for i := range cs {
		imgs = append(imgs, mappedImgs[i])
		delays = append(delays, 50)
		disposals = append(disposals, gif.DisposalPrevious)
	}
	g := gif.GIF{
		Image:     imgs,
		Delay:     delays,
		Disposal:  disposals,
		LoopCount: 0,
	}
	gif.EncodeAll(w, &g)
}

func CanvasToRGBA(c Canvas) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, c.Width, c.Height), palette.Plan9)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			p := c.Pixel(x, y)
			for oy := 0; oy < 2; oy++ {
				for ox := 0; ox < 2; ox++ {
					img.Set(
						x+ox,
						y+oy,
						color.RGBA{
							uint8(scaledColorValue(p.R())),
							uint8(scaledColorValue(p.G())),
							uint8(scaledColorValue(p.B())),
							0,
						},
					)
				}
			}
		}
	}
	return img
}
