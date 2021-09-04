package context

import (
	"math/rand"

	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func WithRange(parent Context, r timer.Ranger, callback func(ctx Context)) {
	fixedRand := rand.Float64()
	iter := r.Iterate()
	for iter.Next() {
		ctx := rangeTimerCtx{
			parent:    parent,
			rng:       iter,
			fixedRand: fixedRand,
		}
		ctx.eventer = newEventer(ctx)
		callback(ctx)
	}
}

type rangeTimerCtx struct {
	parent    Context
	rng       timer.Range
	fixedRand float64
	eventer
}

// public
func (c rangeTimerCtx) BOffset(o float64) Context { return WithBOffset(c, o) }
func (c rangeTimerCtx) Sequence(s timer.Sequencer, callback func(ctx Context)) {
	WithSequence(c, s, callback)
}
func (c rangeTimerCtx) Range(r timer.Ranger, callback func(ctx Context)) {
	WithRange(c, r, callback)
}
func (c rangeTimerCtx) Beat(beat float64, callback func(ctx Context)) {
	WithSequence(c, timer.Beat(beat), callback)
}
func (c rangeTimerCtx) BeatSequence(seq []float64, ghostBeat float64, callback func(ctx Context)) {
	WithSequence(c, timer.Seq(seq, ghostBeat), callback)
}
func (c rangeTimerCtx) BeatRange(startB, endB float64, steps int, easeFn ease.Func, callback func(ctx Context)) {
	WithRange(c, timer.Rng(startB, endB, steps, easeFn), callback)
}
func (c rangeTimerCtx) BeatRangeInterval(startB, endB, interval float64, easeFn ease.Func, callback func(ctx Context)) {
	WithRange(c, timer.RngInterval(startB, endB, interval, easeFn), callback)
}
func (c rangeTimerCtx) Light(l Light, callback func(ctx LightContext)) {
	WithLight(c, l, callback)
}

// private
func (c rangeTimerCtx) base() base      { return c.parent.base() }
func (c rangeTimerCtx) baseTimer() bool { return false }
func (c rangeTimerCtx) offset() float64 { return c.parent.offset() + c.rng.B() }

// pass up
func (c rangeTimerCtx) FixedRand() float64 {
	return c.fixedRand
}

// passthrough to base
func (c rangeTimerCtx) FilterEvents(f func(e evt.Event) bool) *[]evt.Event {
	return c.parent.FilterEvents(f)
}
func (c rangeTimerCtx) MapEvents(f func(e evt.Event) evt.Event) {
	c.parent.MapEvents(f)
}
func (c rangeTimerCtx) MaxLightID(t evt.LightType) int { return c.parent.MaxLightID(t) }
func (c rangeTimerCtx) addEvents(events ...evt.Event)  { c.parent.addEvents(events...) }

// passthrough to range timer
func (c rangeTimerCtx) B() float64                 { return c.rng.B() }
func (c rangeTimerCtx) T() float64                 { return c.rng.T() }
func (c rangeTimerCtx) Ordinal() int               { return c.rng.Ordinal() }
func (c rangeTimerCtx) StartB() float64            { return c.rng.StartB() }
func (c rangeTimerCtx) EndB() float64              { return c.rng.EndB() }
func (c rangeTimerCtx) Duration() float64          { return c.rng.Duration() }
func (c rangeTimerCtx) First() bool                { return c.rng.First() }
func (c rangeTimerCtx) Last() bool                 { return c.rng.Last() }
func (c rangeTimerCtx) Next() bool                 { return c.rng.Next() }
func (c rangeTimerCtx) ToRange() timer.Range       { return c.rng.ToRange() }
func (c rangeTimerCtx) ToSequence() timer.Sequence { return c.rng.ToSequence() }

// pass up to closest sequence timer, fallback to conversion if closest sequence timer is base
func (c rangeTimerCtx) SeqT() float64 {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqT()
	}
	return c.parent.SeqT()
}

func (c rangeTimerCtx) SeqOrdinal() int {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqOrdinal()
	}
	return c.parent.SeqOrdinal()
}

func (c rangeTimerCtx) SeqLen() int {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqLen()
	}
	return c.parent.SeqLen()
}

func (c rangeTimerCtx) SeqNextB() float64 {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqNextB()
	}
	return c.parent.SeqNextB()
}

func (c rangeTimerCtx) SeqNextBOffset() float64 {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqNextBOffset()
	}
	return c.parent.SeqNextBOffset()
}

func (c rangeTimerCtx) SeqPrevB() float64 {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqPrevB()
	}
	return c.parent.SeqPrevB()
}

func (c rangeTimerCtx) SeqPrevBOffset() float64 {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqPrevBOffset()
	}
	return c.parent.SeqPrevBOffset()
}

func (c rangeTimerCtx) SeqFirst() bool {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqFirst()
	}
	return c.parent.SeqFirst()
}

func (c rangeTimerCtx) SeqLast() bool {
	if c.parent.baseTimer() {
		return c.rng.ToSequence().SeqLast()
	}
	return c.parent.SeqLast()
}
