package ilysa

import (
	"github.com/shasderias/ilysa/beatsaber"
)

type Light interface {
	CreateRGBEvent(ctx LightContext) *CompoundRGBLightingEvent
	EventTypeSet() beatsaber.EventTypeSet
	LightIDLen() int
}

type LightContext interface {
	Timer
	Lighter
	LightTimer
}

type LightIDMaxer interface {
	LightIDMax(beatsaber.EventType) int
}
