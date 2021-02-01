package main

import (
	"image/color"
)

var (
	// DarkPurple has a dark purple hue around the outside and other colours in the middle
	DarkPurple = color.RGBA64{470, 230, 1270, 65535}
	// DarkPink has a dark pink hue around the outside and other colours in the middle
	DarkPink = color.RGBA64{1063, 230, 1270, 65535}
	// Greyscale is just greyscale...that's it
	Greyscale = color.RGBA64{1000, 1000, 1000, 65535}
)

// GetPixelColor gets the pixel color based on number of iterations and the passed in multiplier
// the multiplier sets the hue
func GetPixelColor(iterations int, multiplier color.RGBA64) color.RGBA64 {
	return color.RGBA64{
		uint16((uint16(iterations) * multiplier.R) % 65535),
		uint16((uint16(iterations) * multiplier.G) % 65535),
		uint16((uint16(iterations) * multiplier.B) % 65535),
		65535,
	}
}
