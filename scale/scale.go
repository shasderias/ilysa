package scale

type Fn func(m float64) float64

// Clamped returns a function that scales a number from the interval [rMin,rMax]
// to the interval [tMin,tMax]
// https://stats.stackexchange.com/questions/281162/scale-a-number-between-a-range
func Clamped(rMin, rMax, tMin, tMax float64) func(m float64) float64 {
	return func(m float64) float64 {
		if rMin == rMax {
			return 0
		}
		// clamp to range
		switch {
		case m > rMax:
			return rMax
		case m < rMin:
			return rMin
		}
		return (m-rMin)/(rMax-rMin)*(tMax-tMin) + tMin
	}
}

// FromUnitIntervalClamped returns a function that scales a number from the unit
// interval ([0,1]) to the interval [tMin,tMax]
func FromUnitIntervalClamped(tMin, tMax float64) func(m float64) float64 {
	return Clamped(0, 1, tMin, tMax)
}

// ToUnitIntervalClamped returns a function that scales a number from the interval
// [rMin,rMax] to the unit interval ([0,1])
func ToUnitIntervalClamped(rMin, rMax float64) func(m float64) float64 {
	return Clamped(rMin, rMax, 0, 1)
}
