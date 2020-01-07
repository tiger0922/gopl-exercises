// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
    "strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.RGBA{240,52,52,1}, color.RGBA{41, 241, 195, 1}, color.RGBA{30, 139, 195, 1}, color.RGBA{255, 255, 126, 1}} 

const (
	whiteIndex = 0 // first color in palette
	redIndex = 1 // next color in palette
    greenIndex = 2
    blueIndex = 3
    yellowIndex = 4
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
            cycles  := 5
            res     := 0.001
            size    := 100
            nframe  := 64
            delay   := 8
            if err := r.ParseForm(); err != nil {
                log.Print(err)
            }
            for k, v := range r.Form {
                value := v[0]
                if k == "cycles" {
                    if i, err := strconv.Atoi(value); err == nil {
                        cycles = i
                    }
                }
                if k == "res" {
                    if i, err := strconv.ParseFloat(value, 64); err == nil {
                        res = i
                    }
                }
                if k == "size" {
                    if i, err := strconv.Atoi(value); err == nil {
                        size = i
                    }
                }
                if k == "nframe" {
                    if i, err := strconv.Atoi(value); err == nil {
                        nframe = i
                    }
                }
                if k == "delay" {
                    if i, err := strconv.Atoi(value); err == nil {
                        delay = i
                    }
                }
            }
			lissajous(w, cycles, res, size, nframe, delay)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, 5, 0.001, 100, 64, 8)
}

func lissajous(out io.Writer, Cycles int, Res float64, Size int, Nframes int, Delay int) {

    cycles  := Cycles    // number of complete x oscillator revolutions
    res     := Res       // angular resolution
    size    := Size      // image canvas covers [-size..+size]
    nframes := Nframes   // number of animation frames
    delay   := Delay     // delay between frames in 10ms units

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5),
				uint8(t)/5%4+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
