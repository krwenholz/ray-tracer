package viz

import (
	"image"
	"image/gif"
	"io"
	"log"
	"sync"

	"github.com/schollz/progressbar/v3"
)

type RGBADrawable interface {
	DrawRGBA(t int) *image.Paletted
	Len() int
	Height() int
	Width() int
}

func EncodeGIF(w io.Writer, imgs []*image.Paletted, delay int) {
	delays := []int{}
	disposals := []byte{}

	for range imgs {
		delays = append(delays, delay)
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

func DrawAllRGBA(d RGBADrawable) []*image.Paletted {
	res := []*image.Paletted{}
	imgs := sync.Map{}

	wg := sync.WaitGroup{}
	wg.Add(d.Len())
	bar := progressbar.Default(int64(d.Len()))
	bar.Describe("Drawing")

	for i := 0; i < d.Len(); i++ {
		t := i
		go func() {
			imgs.Store(t, d.DrawRGBA(t))
			bar.Add(1)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := 0; i < d.Len(); i++ {
		v, ok := imgs.Load(i)
		if ok {
			if img, ok := v.(*image.Paletted); ok {
				res = append(res, img)
			}
		}
		if !ok {
			log.Fatal("Map read failed for draw", i)
		}
	}
	return res
}
