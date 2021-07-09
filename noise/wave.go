package noise

import (
	"math"

	"github.com/shasderias/ilysa/scale"
)

var (
	DefaultWave = Wave(0.25)
)

func Wave(offset float64) func(x float64) float64 {
	var (
		pi  = math.Pi
		sin = math.Sin
	)

	return func(x float64) float64 {
		r := scale.Clamp(-2, 2, -0.1, 1)((sin(2*x) + sin(pi*x)) + offset)
		if r < 0 {
			r = 0
		}
		return r
	}
}
