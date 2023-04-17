package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.White,
	color.RGBA64{0xFFFF, 0x0000, 0x0000, 0xFF},
	color.RGBA64{0x0000, 0xFFFF, 0x0000, 0xFF},
	color.RGBA64{0x0000, 0x0000, 0xFFFF, 0xFF},
}

type LissajousConfig struct {
	NFrames           int
	AngularResolution float64
	ImageSize         int
	FrameTimeMs       int
	OscilatoroCycles  float64
	Frequency         float64
}

func LissajousConfigCreate() LissajousConfig {
	config := LissajousConfig{}
	config.NFrames = 120
	config.OscilatoroCycles = 10
	config.FrameTimeMs = 10
	config.AngularResolution = 0.001
	config.ImageSize = 120
	return config
}

func Lissajous(fouot io.Writer, config *LissajousConfig) {
	nframes := config.NFrames
	angular_resolution := config.AngularResolution
	image_size := config.ImageSize
	frame_time_ms := config.FrameTimeMs
	oscilator_cycles := config.OscilatoroCycles
	freq := config.Frequency

	anim := gif.GIF{LoopCount: nframes}
	phase_diff := 0.0

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*image_size+1, 2*image_size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < oscilator_cycles*2.0*math.Pi; t += angular_resolution {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase_diff)

			colorIndex := uint8(rand.Int()%len(palette)-1) + 1

			x_ := image_size + int(x*(float64(image_size)+0.5))
			y_ := image_size + int(y*(float64(image_size)+0.5))
			img.SetColorIndex(x_, y_, colorIndex)
		}
		phase_diff += 0.1

		anim.Delay = append(anim.Delay, frame_time_ms)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(fouot, &anim)
}

// This siple program will generate gifs of Lissajous figures
// http://www.fotoacustica.fis.ufba.br/daniele/FIS3/roteiro%208%20oscilosc%C3%B3pio%20Digital%20FigurasLissajous.pdf
func handler(w http.ResponseWriter, r *http.Request) {
	config := LissajousConfigCreate()

	if err := r.ParseForm(); err == nil {
		form := r.Form
		if v, ok := form["NFrames"]; ok == true {
			v, err := strconv.Atoi(v[0])
			if err != nil {
				log.Print(err)
			}
			config.NFrames = v
		}

		if v, ok := form["Frequency"]; ok == true {
			v, err := strconv.ParseFloat(v[0], 64)
			if err != nil {
				log.Print(err)
			}
			config.Frequency = v
		}

		/*
		   NFrames           int
		   AngularResolution float64
		   ImageSize         int
		   FrameTimeMs       int
		   OscilatoroCycles  float64
		   Frequency         float64
		*/
	}
	fmt.Fprintf(os.Stdout, "%v\n", config)
	Lissajous(w, &config)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	// Lissajous(os.Stdout)

	listen := "localhost:3000"
	http.HandleFunc("/", handler)

	fmt.Fprintf(os.Stdout, "Listening on '%s'\n", listen)
	err := http.ListenAndServe(listen, nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

}
