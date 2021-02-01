package main

import "image"

// ComplexPlane describes a square in the complex plane
type ComplexPlane struct {
	topLeft complex128
	width   float64
}

// NewComplexPlane creates a new square in the complex
func NewComplexPlane(topLeft complex128, width float64) ComplexPlane {
	return ComplexPlane{
		topLeft: topLeft,
		width:   width,
	}
}

func (plane ComplexPlane) pointAtPixel(p image.Point, image image.Image) complex128 {

	imgWidth := image.Bounds().Dx()
	imgHeight := image.Bounds().Dy()

	// percentages for 'how far to the right' and 'how far down' the image the pixel is
	var imageXPerc float64 = float64(p.X) / float64(imgWidth)
	var imageYPerc float64 = float64(p.Y) / float64(imgHeight)

	// use the percentages relative to topLeft to compute the coords of the complex value
	real := real(plane.topLeft) + plane.width*imageXPerc
	imag := imag(plane.topLeft) - plane.width*imageYPerc

	return complex(real, imag)
}
