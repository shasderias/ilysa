package context

import (
	"math/rand"

	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type Project interface {
	MaxLightID(t evt.LightType) int
	AddEvents(events ...evt.Event)
	Events() *[]evt.Event
	FilterEvents(func(e evt.Event) bool) *[]evt.Event
	MapEvents(func(e evt.Event) evt.Event)
}

func Base(project Project) Context {
	ctx := base{
		project:   project,
		fixedRand: rand.Float64(),
	}
	ctx.eventer = newEventer(ctx)
	return ctx
}

type base struct {
	project   Project
	fixedRand float64
	eventer
}

func (b base) base() base {
	return b
}

func (b base) FixedRand() float64 {
	return b.fixedRand
}

func (b base) MaxLightID(t evt.LightType) int {
	return b.project.MaxLightID(t)
}

func (b base) addEvents(events ...evt.Event) {
	b.project.AddEvents(events...)
}

func (b base) BOffset(o float64) Context {
	return WithBOffset(b, o)
}
func (b base) Sequence(s timer.Sequencer, callback func(ctx Context)) {
	WithSequence(b, s, callback)
}
func (b base) Range(r timer.Ranger, callback func(ctx Context)) {
	WithRange(b, r, callback)
}
func (b base) Beat(beat float64, callback func(ctx Context)) {
	WithSequence(b, timer.Beat(beat), callback)
}
func (b base) BeatSequence(seq []float64, ghostBeat float64, callback func(ctx Context)) {
	WithSequence(b, timer.Seq(seq, ghostBeat), callback)
}
func (b base) BeatInterval(startBeat, duration float64, count int, callback func(ctx Context)) {
	WithSequence(b, timer.Interval(startBeat, duration, count), callback)
}
func (b base) BeatRange(startB, endB float64, steps int, easeFn ease.Func, callback func(ctx Context)) {
	WithRange(b, timer.Rng(startB, endB, steps, easeFn), callback)
}
func (b base) BeatRangeInterval(startB, endB, interval float64, easeFn ease.Func, callback func(ctx Context)) {
	WithRange(b, timer.RngInterval(startB, endB, interval, easeFn), callback)
}
func (b base) Light(l Light, callback func(ctx LightContext)) {
	WithLight(b, l, callback)
}

func (b base) FilterEvents(f func(e evt.Event) bool) *[]evt.Event {
	return b.project.FilterEvents(f)
}
func (b base) MapEvents(f func(e evt.Event) evt.Event) {
	b.project.MapEvents(f)
}

func (b base) baseTimer() bool   { return true }
func (b base) offset() float64   { return 0 }
func (b base) B() float64        { return 0 }
func (b base) T() float64        { panic("implement me") }
func (b base) Ordinal() int      { panic("implement me") }
func (b base) StartB() float64   { panic("implement me") }
func (b base) EndB() float64     { panic("implement me") }
func (b base) Duration() float64 { panic("implement me") }
func (b base) First() bool       { panic("implement me") }
func (b base) Last() bool        { panic("implement me") }

func (b base) SeqT() float64              { panic("implement me") }
func (b base) SeqOrdinal() int            { panic("implement me") }
func (b base) SeqLen() int                { panic("implement me") }
func (b base) SeqNextB() float64          { panic("implement me") }
func (b base) SeqNextBOffset() float64    { panic("implement me") }
func (b base) SeqPrevB() float64          { panic("implement me") }
func (b base) SeqPrevBOffset() float64    { panic("implement me") }
func (b base) SeqFirst() bool             { panic("implement me") }
func (b base) SeqLast() bool              { panic("implement me") }
func (b base) ToRange() timer.Range       { return b }
func (b base) ToSequence() timer.Sequence { return b }
func (b base) Next() bool                 { return false }
