package main

import (
	"math/cmplx"
)

// MandlebrotIterations finds n such that |z_n| > 2 where z_n = (z_n-1 ^ 2) + c and z_0 = 0
// returns -1 if not found after maxIterations iterations
func MandlebrotIterations(c complex128, maxIterations int) int {
	var prev complex128 = 0

	for i := 0; i < maxIterations; i++ {
		if cmplx.Abs(prev) > 2 {
			return i
		}
		prev = cmplx.Pow(prev, 2) + c
	}
	return -1
}
