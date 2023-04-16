package viz

import (
	"bytes"
	"image"
	"log"

	x264 "github.com/gen2brain/x264-go"
)

func EncodeX264FromRBA64(width, height, repeat int, imgs []image.RGBA64Image) *bytes.Buffer {
	buf := bytes.NewBuffer(make([]byte, 0))

	opts := x264.Options{
		Width:     width,
		Height:    height,
		FrameRate: 1,
		Tune:      "zerolatency",
		Preset:    "veryfast",
		Profile:   "baseline",
		LogLevel:  x264.LogDebug,
	}
	enc, err := x264.NewEncoder(buf, &opts)
	if err != nil {
		log.Fatal(err)
	}

	for r := 0; r < repeat; r++ {
		for _, i := range imgs {
			err = enc.Encode(i)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	err = enc.Close()
	if err != nil {
		log.Fatal(err)
	}

	return buf
}
