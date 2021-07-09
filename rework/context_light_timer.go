package rework

import (
	"github.com/shasderias/ilysa/light"
)

type lightTimer struct {
	light.Light
	lightIDOrdinal int
}

func newLightTimer(l light.Light, lightIDOrdinal int) lightTimer {
	return lightTimer{
		Light:          l,
		lightIDOrdinal: lightIDOrdinal,
	}
}

func (c lightTimer) LightIDOrdinal() int {
	return c.lightIDOrdinal
}

func (c lightTimer) LightIDCur() int {
	return c.lightIDOrdinal + 1
}

func (c lightTimer) LightIDT() float64 {
	return float64(c.LightIDOrdinal()) / 1 // float64(c.LightIDLen())
}
