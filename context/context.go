package context

import (
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type Context interface {
	Timer
	Eventer
	Offset() float64

	addEvents(events ...evt.Event)
	baseTimer() bool
}

type Timer interface {
	timer.Sequence
	timer.Range
}

type LightContext interface {
	Timer
	LightTimer
}

type LightTimer interface {
	LightIDT() float64 // current time in light ID sequence, 0-1
	LightIDOrdinal() int
	LightIDLen() int
	LightIDCur() int
}

type Eventer interface {
	NewLighting(opts ...evt.LightingOpt) *evt.Lighting
	NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting
}
