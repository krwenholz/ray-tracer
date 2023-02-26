package viz

import (
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io"
)

func EncodeGIF(w io.Writer, cs []Canvas) *gif.GIF {
	imgs := []*image.Paletted{}
	delays := []int{}
	disposals := []byte{}
	for _, c := range cs {
		imgs = append(imgs, CanvasToRGBA(c))
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
	return nil
}

func CanvasToRGBA(c Canvas) *image.Paletted {
	img := image.NewPaletted(image.Rect(0, 0, c.Width, c.Height), palette.Plan9)
	for y := 0; y < c.Height; y++ {
		for x := 0; x < c.Width; x++ {
			p := c.Pixel(x, y)
			img.Set(
				x,
				y,
				color.RGBA{
					uint8(scaledColorValue(p.R())),
					uint8(scaledColorValue(p.G())),
					uint8(scaledColorValue(p.B())),
					0,
				},
			)
		}
	}
	return img
}
