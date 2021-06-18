package ilysa

import (
	"github.com/shasderias/ilysa/pkg/beatsaber"
)

type Light interface {
	CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent
	EventTypeSet() beatsaber.EventTypeSet
	LightIDLen() int
}

type LightIDMaxer interface {
	LightIDMax(beatsaber.EventType) int
}

