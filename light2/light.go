package light2

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/timer"
)

type Light interface {
	CreateRGBLightingEvent(ctx LightContext) *CompoundRGBLightingEvent
	EventTypeSet() beatsaber.EventTypeSet
	LightIDLen() int
}

type LightContext interface {
	timer.Range
	ilysa.Lighter
	ilysa.LightTimer
}

type LightIDMaxer interface {
	LightIDMax(beatsaber.EventType) int
}
