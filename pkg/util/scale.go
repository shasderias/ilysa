package util

// Scale returns a function that scales a number from the interval [rMin,rMax]
// to the interval [tMin,tMax]
// https://stats.stackexchange.com/questions/281162/scale-a-number-between-a-range
func Scale(rMin, rMax, tMin, tMax float64) func(m float64) float64 {
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

// ScaleFromUnitInterval returns a function that scales a number from the unit
// interval ([0,1]) to the interval [tMin,tMax]
func ScaleFromUnitInterval(tMin, tMax float64) func(m float64) float64 {
	return Scale(0, 1, tMin, tMax)
}

// ScaleToUnitInterval returns a function that scales a number from the interval
// [rMin,rMax] to the unit interval ([0,1])
func ScaleToUnitInterval(rMin, rMax float64) func(m float64) float64 {
	return Scale(rMin, rMax, 0, 1)
}
