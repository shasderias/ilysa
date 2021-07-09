package light2

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/rework"
	"github.com/shasderias/ilysa/timer"
)

type Light interface {
	CreateRGBLightingEvent(ctx LightContext) *CompoundRGBLightingEvent
	EventTypeSet() beatsaber.EventTypeSet
	LightIDLen() int
}

type LightContext interface {
	timer.Range
	rework.Lighter
	rework.LightTimer
}

type LightIDMaxer interface {
	MaxLightID(beatsaber.EventType) int
}
