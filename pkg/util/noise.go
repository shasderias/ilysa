package util

import "math"

var (
	DefaultNoise = Noise(0.8, 0.25)
)

func Noise(wiggle, offset float64) func(x float64) float64 {
	var (
		pi  = math.Pi
		sin = math.Sin
	)

	return func(x float64) float64 {
		r := Scale(-2, 2, -0.1, 1)(((sin(2*x) + sin(pi*x)) * wiggle) + offset)
		if r < 0 {
			r = 0
		}
		return r
	}
}
