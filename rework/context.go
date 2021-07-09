package rework

import (
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
)

type EventModifier func(e evt.Event)
type EventModder func(ctx RangeContext, event evt.Event)
type EventFilter func(event evt.Event) bool

type Context interface {
	timer.Range
	timer.Sequence
}

type LightTimer interface {
	LightIDT() float64 // current time in light ID sequence, 0-1
	LightIDOrdinal() int
	LightIDLen() int
	LightIDCur() int
}
type BaseContext interface {
	Sequence(sequence timer.Sequencer, callback func(ctx SequenceContext))
	Range(startBeat, endBeat float64, steps int, fn ease.Func, callback func(ctx RangeContext))
	//ModEventsInRange(startBeat, endBeat float64, filter EventFilter, eventModder func(ctx RangeContext, event Event))
	//NewBasicLight(eventType beatsaber.EventType) Basic
	//MaxLightID(beatsaber.EventType) int
	//Offset(o float64) Context
}

type RangeContext interface {
	BaseContext
	timer.Range
	Eventer
	Lighter
	WithLight(light.Light, func(ctx TimeLightContext))
}

type TimeLightContext interface {
	timer.Range
	LightTimer
	CompoundLighter
}

type SequenceContext interface {
	timer.Range
	timer.Sequence
	Lighter
	Eventer
	EventsForRange(startBeat, endBeat float64, steps int, fn ease.Func, callback func(ctx SequenceTimeContext))
	WithLight(light.Light, func(ctx SequenceLightContext))
}

type SequenceTimeContext interface {
	timer.Sequence
	timer.Range
	WithLight(light.Light, func(ctx SequenceTimeLightContext))
}

type SequenceLightContext interface {
	timer.Sequence
	timer.Range
	LightTimer
	CompoundLighter
}

type SequenceTimeLightContext interface {
	timer.Sequence
	timer.Range
	LightTimer
	CompoundLighter
}

type Lighter interface {
	NewLighting(opts ...evt.LightingOpt) *evt.Lighting
	NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting
}

type CompoundLighter interface {
	NewLighting(opts ...evt.LightingOpt) evt.Events
	NewRGBLighting(opts ...evt.RGBLightingOpt) evt.Events
}

type Eventer interface {
	NewLaser(opts ...evt.LaserOpt) *evt.Laser
	NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser
	NewRotation(opts ...evt.RotationOpt) *evt.Rotation
	NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation
	NewZoom(opts ...evt.ZoomOpt) *evt.Zoom
	NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom
}
