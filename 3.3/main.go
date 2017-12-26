package main

import (
	"fmt"
	"image/color"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, _ := corner(i+1, j)
			bx, by, color := corner(i, j)
			cx, cy, _ := corner(i, j+1)
			dx, dy, _ := corner(i+1, j+1)
			r, g, b, _ := color.RGBA()
			fmt.Printf("<polygon style='fill: #%.2x%.2x%.2x;' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				uint8(r), uint8(g), uint8(b), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, color.Color) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	r := 255 * (z + 1) / 2
	b := 255 * (1 - z) / 2

	return sx, sy, color.RGBA{uint8(r), 0x00, uint8(b), 0x00}
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	if r != 0 {
		return math.Sin(r) / r
	}

	return 0
}
