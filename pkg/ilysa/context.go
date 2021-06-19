package ilysa

import (
	"math/rand"

	"github.com/shasderias/ilysa/pkg/beatsaber"
	"github.com/shasderias/ilysa/pkg/ease"
	"github.com/shasderias/ilysa/pkg/util"
)

type BaseContext interface {
	EventForBeat(beat float64, callback func(ctx TimingContext))
	EventsForBeats(startBeat, duration float64, count int, callback func(ctx TimingContext))
	EventsForRange(startBeat, endBeat float64, steps int, fn ease.Func, callback func(ctx TimingContext))
	EventsForSequence(startBeat float64, sequence []float64, callback func(ctx SequenceContext))
	ModEventsInRange(startBeat, endBeat float64, filter EventFilter, modder EventModder)
	NewBasicLight(eventType beatsaber.EventType) BasicLight
	LightIDMax(beatsaber.EventType) int
}

type Timing interface {
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

type LightContext interface {
	LightIDLen() int
	LightIDCur() int
	LightIDOrdinal() int
	LightIDT() float64
}

type TimingContext interface {
	BaseContext
	Timing
	Eventer
	Lighter
	UseLight(Light, func(TimingContextWithLight))
}

type TimingContextWithLight interface {
	Timing
	CompoundLighter
	LightContext
}

type TimingContextForLight interface {
	Timing
	Lighter
	LightContext
}

type Sequencer interface {
	SequenceIndex(idx int) float64
	NextB() float64
	NextBOffset() float64
	PrevB() float64
	PrevBOffset() float64
}

type SequenceContext interface {
	BaseContext
	Timing
	Lighter
	Eventer
	UseLight(Light, func(SequenceContextWithLight))
	Sequencer
}

type SequenceContextWithLight interface {
	Timing
	LightContext
	CompoundLighter
	Sequencer
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

type baseContext struct {
	*Project
	timing

	beatOffset float64
	fixedRand  float64
	modifiers  []EventModifier
}

func (c baseContext) FixedRand() float64 {
	return c.fixedRand
}

func (c baseContext) LightIDMax(typ beatsaber.EventType) int {
	return c.Project.ActiveDifficultyProfile().LightIDMax(typ)
}

func (c baseContext) WithBeatOffset(offset float64) baseContext {
	return baseContext{
		Project: c.Project,
		timing:  c.timing,

		beatOffset: offset,
		fixedRand:  c.fixedRand,
		modifiers:  c.modifiers,
	}
}

type timing struct {
	b         float64
	startBeat float64
	endBeat   float64
	ordinal   int
}

func (t timing) B() float64 {
	return t.b
}

func (t timing) T() float64 {
	return util.ScaleToUnitInterval(t.startBeat, t.endBeat)(t.b)
}

func (t timing) StartBeat() float64 {
	return t.startBeat
}

func (t timing) EndBeat() float64 {
	return t.endBeat
}

func (t timing) Duration() float64 {
	return t.endBeat - t.startBeat
}

func (t timing) Ordinal() int {
	return t.ordinal
}

func (t timing) First() bool {
	return t.b == t.startBeat
}

func (t timing) Last() bool {
	return t.b == t.endBeat
}

func newBaseContext(p *Project) baseContext {
	return baseContext{
		Project:   p,
		fixedRand: rand.Float64(),
	}
}

func (c baseContext) withTiming(beat, startBeat, endBeat float64, ordinal int) baseContext {
	return baseContext{
		Project: c.Project,
		timing:  newTiming(beat, startBeat, endBeat, ordinal),

		beatOffset: c.beatOffset,
		fixedRand:  c.fixedRand,
		modifiers:  c.modifiers,
	}
}

func newTiming(beat, startBeat, endBeat float64, ordinal int) timing {
	return timing{
		b:         beat,
		startBeat: startBeat,
		endBeat:   endBeat,
		ordinal:   ordinal,
	}
}

func (c baseContext) WithModifier(modifiers ...EventModifier) baseContext {
	return baseContext{
		Project: c.Project,
		timing:  c.timing,

		beatOffset: c.beatOffset,
		fixedRand:  c.fixedRand,
		modifiers:  append([]EventModifier{}, modifiers...),
	}
}

func (c baseContext) addEvent(e Event) {
	c.applyModifiers(e)
	c.Project.events = append(c.Project.events, e)
}

func (c baseContext) applyModifiers(e Event) {
	if len(c.modifiers) == 0 {
		return
	}

	for _, m := range c.modifiers {
		m(e)
	}
}

func (c baseContext) withLight(light Light) contextWithLight {
	return contextWithLight{
		baseContext: c,
		lightContext: lightContext{
			Light:          light,
			lightIDOrdinal: 0,
		},
	}
}

func (c baseContext) withSequence(sequence []float64) sequenceContext {
	return sequenceContext{
		baseContext: c,
		sequence:    sequence,
	}
}

func (c baseContext) UseLight(light Light, callback func(ctx TimingContextWithLight)) {
	ctx := c.withLight(light)
	for i := 0; i < light.LightIDLen(); i++ {
		callback(ctx.withLightIDOrdinal(i))
	}
}

type lightContext struct {
	Light
	lightIDOrdinal int
}

type contextWithLight struct {
	baseContext
	lightContext
}

func (c contextWithLight) withLightIDOrdinal(ordinal int) contextWithLight {
	return contextWithLight{
		baseContext: c.baseContext,
		lightContext: lightContext{
			Light:          c.Light,
			lightIDOrdinal: ordinal,
		},
	}
}

type timingContextForLight struct {
	baseContext
	lightContext
}

func (c contextWithLight) forLight() timingContextForLight {
	return timingContextForLight{
		baseContext:  c.baseContext,
		lightContext: c.lightContext,
	}
}

func (c lightContext) LightIDOrdinal() int {
	return c.lightIDOrdinal
}

func (c lightContext) LightIDCur() int {
	return c.lightIDOrdinal + 1
}

func (c lightContext) LightIDT() float64 {
	return float64(c.LightIDOrdinal()) / float64(c.LightIDLen())
}

func (c contextWithLight) NewLightingEvent(opts ...BasicLightingEventOpt) *CompoundBasicLightingEvent {
	events := CompoundBasicLightingEvent{}

	if c.LightIDOrdinal() > 0 {
		return &events
	}

	for _, et := range c.Light.EventTypeSet() {
		opts := append([]BasicLightingEventOpt{WithType(et)}, opts...)
		events.Add(c.baseContext.NewLightingEvent(opts...))
	}

	return &events
}

func (c contextWithLight) NewRGBLightingEvent(opts ...CompoundRGBLightingEventOpt) *CompoundRGBLightingEvent {
	e := c.Light.CreateRGBEvent(c.forLight())
	e.Mod(opts...)
	return e
}

type sequenceContext struct {
	baseContext
	sequence []float64
}

func (c sequenceContext) SequenceIndex(idx int) float64 {
	return c.sequence[util.WraparoundIdx(len(c.sequence), idx)]
}

func (c sequenceContext) OrdinalOffset(offset int) float64 {
	return c.sequence[util.WraparoundIdx(len(c.sequence), c.Ordinal()+offset)]
}

func (c sequenceContext) NextB() float64 {
	return c.OrdinalOffset(1)
}

func (c sequenceContext) NextBOffset() float64 {
	return c.OrdinalOffset(1) - c.OrdinalOffset(0)
}

func (c sequenceContext) PrevB() float64 {
	return c.OrdinalOffset(-1)
}

func (c sequenceContext) PrevBOffset() float64 {
	return c.OrdinalOffset(0) - c.OrdinalOffset(-1)

}

func (c sequenceContext) UseLight(light Light, callback func(ctx SequenceContextWithLight)) {
	ctx := c.withLight(light)
	for i := 0; i < light.LightIDLen(); i++ {
		callback(ctx.withLightIDOrdinal(i))
	}
}

func (c sequenceContext) withLight(light Light) sequenceContextWithLight {
	return sequenceContextWithLight{
		baseContext: c.baseContext,
		lightContext: lightContext{
			Light:          light,
			lightIDOrdinal: 0,
		},
		sequenceContext: c,
	}
}

type sequenceContextWithLight struct {
	baseContext
	lightContext
	sequenceContext
}

//func (c sequenceContextWithLight) EventsForRange(startBeat, endBeat float64, steps int, easeFn ease.Func, callback func(light TimingContextWithLight)) {
//tScaler := util.ScaleToUnitInterval(0, float64(steps-1))
//
//ctx := timingContextWithLight{
//
//}
//
//for i := 0; i < steps; i++ {
//	beat := Ierp(startBeat, endBeat, tScaler(float64(i)), easeFn)
//	callback(c.withTiming(beat, startBeat, endBeat, i).WithBeatOffset(0))
//}
//}

func (c sequenceContextWithLight) withLightIDOrdinal(ordinal int) sequenceContextWithLight {
	return sequenceContextWithLight{
		baseContext: c.baseContext,
		lightContext: lightContext{
			Light:          c.Light,
			lightIDOrdinal: ordinal,
		},
		sequenceContext: c.sequenceContext,
	}
}

func (c sequenceContextWithLight) NewLightingEvent(opts ...BasicLightingEventOpt) *CompoundBasicLightingEvent {
	events := CompoundBasicLightingEvent{}

	if c.LightIDOrdinal() > 0 {
		return &events
	}

	for _, et := range c.Light.EventTypeSet() {
		opts := append([]BasicLightingEventOpt{WithType(et)}, opts...)
		events.Add(c.baseContext.NewLightingEvent(opts...))
	}

	return &events
}

func (c sequenceContextWithLight) NewRGBLightingEvent(opts ...CompoundRGBLightingEventOpt) *CompoundRGBLightingEvent {
	e := c.Light.CreateRGBEvent(c.forLight())
	e.Mod(opts...)
	return e
}

func (c sequenceContextWithLight) forLight() timingContextForLight {
	return timingContextForLight{
		baseContext:  c.baseContext,
		lightContext: c.lightContext,
	}
}

type EventModifier func(e Event)
type EventGenerator func(ctx Timing)
type EventModder func(ctx TimingContext, event Event)
type EventFilter func(event Event) bool
