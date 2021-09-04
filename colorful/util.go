package colorful

import "math"

func sq(v float64) float64 {
	return v * v
}

func cub(v float64) float64 {
	return v * v * v
}

// clamp01 clamps from 0 to 1.
func clamp01(v float64) float64 {
	return math.Max(0.0, math.Min(v, 1.0))
}

// lerpAngle interpolates between two angles in the range [0,360].
// Adapted from http://stackoverflow.com/a/14498790/2366315.
// With potential proof that it works here: http://math.stackexchange.com/a/2144499
func lerpAngle(a0, a1, t float64) float64 {
	delta := math.Mod(math.Mod(a1-a0, 360.0)+540, 360.0) - 180.0
	return math.Mod(a0+t*delta+360.0, 360.0)
}
