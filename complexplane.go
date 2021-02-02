package main

import "image"

// ComplexPlane describes a square in the complex plane
type ComplexPlane struct {
	center complex128
	radius float64
}

// NewComplexPlane creates a new square in the complex
func NewComplexPlane(center complex128, radius float64) ComplexPlane {
	return ComplexPlane{
		center: center,
		radius: radius,
	}
}

// pointAtPixel translates a pixel in an image to a point in the complex plane
func (plane ComplexPlane) pointAtPixel(p image.Point, image image.Image) complex128 {

	topLeft := plane.center + complex(-1 * plane.radius, plane.radius)
	width := 2 * plane.radius

	imgWidth := image.Bounds().Dx()
	imgHeight := image.Bounds().Dy()

	// percentages for 'how far to the right' and 'how far down' the image the pixel is
	var imageXPerc float64 = float64(p.X) / float64(imgWidth)
	var imageYPerc float64 = float64(p.Y) / float64(imgHeight)

	// use the percentages relative to topLeft to compute the coords of the complex value
	real := real(topLeft) + width*imageXPerc
	imag := imag(topLeft) - width*imageYPerc

	return complex(real, imag)
}
