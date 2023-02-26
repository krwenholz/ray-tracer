package viz

import (
	"image"
	"image/gif"
	"io"
)

func EncodeGIF(w io.Writer, imgs []*image.Paletted) {
	delays := []int{}
	disposals := []byte{}

	for range imgs {
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
