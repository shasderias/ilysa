// Package scale implements functions for scaling numbers and ranges.
//
// Range scaling formula adapted from https://stats.stackexchange.com/questions/281162/scale-a-number-between-a-range
package scale

import "math"

type Fn func(m float64) float64

func clamp(t, min, max float64) float64 {
	min, max = math.Min(min, max), math.Max(min, max)
	return math.Max(math.Min(t, max), min)
}

// Unclamp returns a function that scales a number from the interval [rMin,rMax]
// to the interval [tMin,tMax].
func Unclamp(rMin, rMax, tMin, tMax float64) func(m float64) float64 {
	return func(m float64) float64 {
		if rMin == rMax {
			return 0
		}
		return (m-rMin)/(rMax-rMin)*(tMax-tMin) + tMin
	}
}

// FromUnitUnclamp returns a function that scales a number from the unit interval
// ([0,1]) to the interval [tMin,tMax].
func FromUnitUnclamp(tMin, tMax float64) func(m float64) float64 {
	return Unclamp(0, 1, tMin, tMax)
}

// ToUnitUnclamp returns a function that scales a number from the interval [rMin,rMax]
// to the unit interval ([0,1]).
func ToUnitUnclamp(rMin, rMax float64) func(m float64) float64 {
	return Unclamp(rMin, rMax, 0, 1)
}

// Clamp returns a function that scales a number from the interval [rMin,rMax]
// to the interval [tMin,tMax], if the result falls outside [tMin,tMax], it is
// clamped to tMin or tMax.
func Clamp(rMin, rMax, tMin, tMax float64) func(m float64) float64 {
	return func(m float64) float64 {
		if rMin == rMax {
			return 0
		}
		t := (m-rMin)/(rMax-rMin)*(tMax-tMin) + tMin

		return clamp(t, tMin, tMax)
	}
}

// FromUnitClamp returns a function that scales a number from the unit interval
// ([0,1]) to the interval [tMin,tMax], if the result falls outside [tMin,tMax],
// it is clamped to tMin or tMax.
func FromUnitClamp(tMin, tMax float64) func(m float64) float64 {
	return Clamp(0, 1, tMin, tMax)
}

// ToUnitClamp returns a function that scales a number from the interval [rMin,rMax]
// to the unit interval ([0,1]), if the result falls outside [0,1], it is clamped
// to 0 or 1.
func ToUnitClamp(rMin, rMax float64) func(m float64) float64 {
	return Clamp(rMin, rMax, 0, 1)
}

// StepsForFPS returns the number of steps required to attain targetFPS, or minFPS,
// whichever is greater.
func StepsForFPS(startBeat, endBeat, targetFPS, minFPS float64) int {
	steps := math.RoundToEven((endBeat - startBeat) * targetFPS)

	return int(math.Max(steps, minFPS))
}
