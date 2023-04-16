package main

import (
	"image"
	"image/color"
	"image/gif"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.White,
	color.RGBA64{0xFFFF, 0x0000, 0x0000, 0xFF},
	color.RGBA64{0x0000, 0xFFFF, 0x0000, 0xFF},
	color.RGBA64{0x0000, 0x0000, 0xFFFF, 0xFF},
}

func lissajous(fouot *os.File) {
	const nframes = 120
	const angular_resolution = 0.0001
	const image_size = 120
	const frame_time_ms = 4
	const oscilator_cycles = 10

	freq := rand.Float64() * 3

	anim := gif.GIF{LoopCount: nframes}
	phase_diff := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*image_size+1, 2*image_size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < oscilator_cycles*2.0*math.Pi; t += angular_resolution {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase_diff)

			colorIndex := uint8(rand.Int()%len(palette)-1) + 1

			img.SetColorIndex(
				image_size+int(x*image_size+0.5),
				image_size+int(y*image_size+0.5),
				colorIndex,
			)
		}
		phase_diff += 0.1

		anim.Delay = append(anim.Delay, frame_time_ms)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(fouot, &anim)
}

// This siple program will generate gifs of Lissajous figures
// http://www.fotoacustica.fis.ufba.br/daniele/FIS3/roteiro%208%20oscilosc%C3%B3pio%20Digital%20FigurasLissajous.pdf
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}
