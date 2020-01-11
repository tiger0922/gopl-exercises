// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
    "fmt"
    "math"
    "os"
    "image/color"
    "net/http"
    "log"
)

const (
    width, height = 600, 320            // canvas size in pixels
    cells         = 100                 // number of grid cells
    xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
    xyscale       = width / 2 / xyrange // pixels per x or y unit
    zscale        = height * 0.4        // pixels per z unit
    angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
    handler := func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "image/svg+xml")
        fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
        "style='stroke: grey; fill: white; stroke-width: 0.7' "+
        "width='%d' height='%d'>", width, height)
        for i := 0; i < cells; i++ {
            for j := 0; j < cells; j++ {
                ax, ay := corner(i+1, j)
                bx, by := corner(i, j)
                cx, cy := corner(i, j+1)
                dx, dy := corner(i+1, j+1)

                if math.IsNaN(ax) || math.IsNaN(ay) || math.IsNaN(bx) || math.IsNaN(by) || math.IsNaN(cx) || math.IsNaN(cy) || math.IsNaN(dx) || math.IsNaN(dy) {
                    continue
                }

                r,g,b,_ := dye(i,j).RGBA()

                fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill: #%.2x%.2x%.2x'/>\n",
                ax, ay, bx, by, cx, cy, dx, dy, uint8(r), uint8(g), uint8(b))
            }
        }
        fmt.Fprintln(w, "</svg>")
    }
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func dye(i, j int) color.Color {
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    z := f(x, y)

    r := 255 * (z+1)/2
    b := 255 * (1-z)/2
    return color.RGBA{uint8(r), 0x00, uint8(b), 0x00}
}

func corner(i, j int) (float64, float64) {
    // Find point (x,y) at corner of cell (i,j).
    x := xyrange * (float64(i)/cells - 0.5)
    y := xyrange * (float64(j)/cells - 0.5)

    // Compute surface height z.
    z := f(x, y)

    // Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
    sx := width/2 + (x-y)*cos30*xyscale
    sy := height/2 + (x+y)*sin30*xyscale - z*zscale
    return sx, sy
}

func eggbox(x, y float64) float64 {
    return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
    return math.Pow(x, 2)/(25*25) - math.Pow(y, 2)/(17*17)
}

func moguls(x, y float64) float64 {
    return 0.03 * (math.Cos(x) + math.Cos(y))
}

func origin(x, y float64) float64 {
    r := math.Hypot(x, y)
    return math.Sin(r) / r
}

func f(x, y float64) float64 {
    var r float64
    if len(os.Args) < 2 {
        return origin(x, y)
    }
    switch os.Args[1] {
        case "saddle":
            r = saddle(x, y)
        case "eggbox":
            r = eggbox(x, y)
        case "moguls":
            r = moguls(x, y)
        default:
            fmt.Fprintln(os.Stderr)
            os.Exit(1)
    }
    return r
}

//!-
