package main

import (
	"image"
	"image/color"
)

// CreateImage creates a square image with the specified width, using the getColor to set the color of each pixel
func CreateImage(width int, getColor func(image.Point, image.Image) color.Color) image.Image {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{width, width}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	done := make(chan bool, width * width)
	for x := 0; x < width; x++ {
		for y := 0; y < width; y++ {
			go func(x, y int) {
				img.Set(x, y, getColor(image.Point{X: x, Y: y}, img))
				done <- true
			}(x, y)
		}
	}

	// wait for all colors to be set
	for i := 0; i < width * width; i++ {
		_ = <- done
	}
	close(done)

	return img
}
