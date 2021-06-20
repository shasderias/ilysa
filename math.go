package ilysa

import (
	"math"

	"github.com/shasderias/ilysa/ease"
)

func Ierp(min, max float64, pos float64, easeFunc ease.Func) float64 {
	return min + (max-min)*easeFunc(pos)
}

func IntIerp(min, max int, pos float64, easeFunc ease.Func) int {
	return int(math.RoundToEven(Ierp(float64(min), float64(max), pos, easeFunc)))
}
