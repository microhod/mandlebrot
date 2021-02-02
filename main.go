package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"strconv"
)

var (
	planeTopLeft            = -2 + 1.5i
	planeWidth              = 3.0
	imgWidth                = 1000
	mandlebrotMaxIterations = 250
	colorPalette            = DarkPurple
)

func main() {
	// parse first parameter as width (otherwise just use the default)
	var err error
	if len(os.Args) > 1 {
		imgWidth, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("width must be an integer")
			os.Exit(1)
		}
	}

	plane := NewComplexPlane(planeTopLeft, planeWidth)

	img := CreateImage(imgWidth, func(p image.Point, img image.Image) color.Color {
		c := plane.pointAtPixel(p, img)
		if i := MandlebrotIterations(c, mandlebrotMaxIterations); i >= 0 {
			return GetPixelColor(i, colorPalette)
		}
		return color.Black
	})

	f, _ := os.Create("image.png")
	png.Encode(f, img)
}
