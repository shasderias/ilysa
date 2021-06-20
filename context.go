package ilysa

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/ease"
)

type EventModifier func(e Event)
type EventModder func(ctx TimeContext, event Event)
type EventFilter func(event Event) bool

type BaseContext interface {
	EventForBeat(beat float64, callback func(ctx TimeContext))
	EventsForBeats(startBeat, duration float64, count int, callback func(ctx TimeContext))
	EventsForRange(startBeat, endBeat float64, steps int, fn ease.Func, callback func(ctx TimeContext))
	EventsForSequence(startBeat float64, sequence []float64, callback func(ctx SequenceContext))
	ModEventsInRange(startBeat, endBeat float64, filter EventFilter, eventModder func(ctx TimeContext, event Event))
	NewBasicLight(eventType beatsaber.EventType) BasicLight
	LightIDMax(beatsaber.EventType) int
}

type Timer interface {
	B() float64         // current beat
	T() float64         // current time in the current sequence, on a 0-1 scale
	Ordinal() int       // ordinal number of the current iteration, starting from 0
	StartBeat() float64 // first beat of the current sequence
	EndBeat() float64   // last beat of the current sequence
	Duration() float64  // duration of the current sequence, in beats
	First() bool        // true if this is the first iteration
	Last() bool         // true if this is the last iteration
	FixedRand() float64 // a number from 0-1, fixed for the current sequence, but different for every sequence
}

type LightTimer interface {
	LightIDLen() int
	LightIDCur() int
	LightIDOrdinal() int
	LightIDT() float64
}

type SequenceTimer interface {
	SequenceIndex(idx int) float64
	NextB() float64
	NextBOffset() float64
	PrevB() float64
	PrevBOffset() float64
}

type TimeContext interface {
	BaseContext
	Timer
	Eventer
	Lighter
	WithLight(Light, func(TimeLightContext))
}

type TimeLightContext interface {
	Timer
	LightTimer
	CompoundLighter
}

type SequenceContext interface {
	BaseContext
	Timer
	SequenceTimer
	Lighter
	Eventer
	WithLight(Light, func(SequenceLightContext))
}

type SequenceLightContext interface {
	Timer
	LightTimer
	SequenceTimer
	CompoundLighter
}

type Lighter interface {
	NewLightingEvent(opts ...BasicLightingEventOpt) *BasicLightingEvent
	NewRGBLightingEvent(opts ...RGBLightingEventOpt) *RGBLightingEvent
}

type CompoundLighter interface {
	NewLightingEvent(opts ...BasicLightingEventOpt) *CompoundBasicLightingEvent
	NewRGBLightingEvent(opts ...CompoundRGBLightingEventOpt) *CompoundRGBLightingEvent
}

type Eventer interface {
	NewRotationEvent(opts ...RotationEventOpt) *RotationEvent
	NewPreciseRotationEvent(opts ...PreciseRotationEventOpt) *PreciseRotationEvent
	NewRotationSpeedEvent(opts ...RotationSpeedEventOpt) *RotationSpeedEvent
	NewPreciseRotationSpeedEvent(opts ...PreciseRotationSpeedEventOpt) *PreciseRotationSpeedEvent
	NewZoomEvent(opts ...ZoomEventOpt) *ZoomEvent
	NewPreciseZoomEvent(opts ...PreciseZoomEventOpt) *PreciseZoomEvent
}
