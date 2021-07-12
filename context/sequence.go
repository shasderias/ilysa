package context

import (
	"math/rand"

	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func WithSequence(parent Context, s timer.Sequencer, callback func(ctx Context)) {
	iter := s.Iterate()
	for iter.Next() {
		ctx := seqTimerCtx{
			parent:    parent,
			seq:       iter,
			fixedRand: rand.Float64(),
		}
		ctx.eventer = newEventer(ctx)
		callback(ctx)
	}
}

type seqTimerCtx struct {
	parent    Context
	seq       timer.Sequence
	fixedRand float64
	eventer
}

// public
func (c seqTimerCtx) BOffset(o float64) Context { return WithOffset(c, o) }
func (c seqTimerCtx) Sequence(s timer.Sequencer, callback func(ctx Context)) {
	WithSequence(c, s, callback)
}
func (c seqTimerCtx) Range(r timer.Ranger, callback func(ctx Context)) {
	WithRange(c, r, callback)
}
func (c seqTimerCtx) Light(l Light, callback func(ctx LightContext)) {
	WithLight(c, l, callback)
}

func (c seqTimerCtx) FixedRand() float64 {
	return c.fixedRand
}

// private
func (c seqTimerCtx) baseTimer() bool { return false }
func (c seqTimerCtx) offset() float64 { return c.parent.offset() + c.seq.ToRange().B() }

// passthrough to base
func (c seqTimerCtx) addEvents(events ...evt.Event)  { c.parent.addEvents(events...) }
func (c seqTimerCtx) MaxLightID(t evt.LightType) int { return c.parent.MaxLightID(t) }

// passthrough to sequence timer
func (c seqTimerCtx) SeqT() float64              { return c.seq.SeqT() }
func (c seqTimerCtx) SeqOrdinal() int            { return c.seq.SeqOrdinal() }
func (c seqTimerCtx) SeqLen() int                { return c.seq.SeqLen() }
func (c seqTimerCtx) SeqNextB() float64          { return c.seq.SeqNextB() }
func (c seqTimerCtx) SeqNextBOffset() float64    { return c.seq.SeqNextBOffset() }
func (c seqTimerCtx) SeqPrevB() float64          { return c.seq.SeqPrevB() }
func (c seqTimerCtx) SeqPrevBOffset() float64    { return c.seq.SeqPrevBOffset() }
func (c seqTimerCtx) SeqFirst() bool             { return c.seq.SeqFirst() }
func (c seqTimerCtx) SeqLast() bool              { return c.seq.SeqLast() }
func (c seqTimerCtx) Next() bool                 { return c.seq.Next() }
func (c seqTimerCtx) ToRange() timer.Range       { return c.seq.ToRange() }
func (c seqTimerCtx) ToSequence() timer.Sequence { return c.seq.ToSequence() }

// pass up to range timer, fallback to conversion if closest range timer is base
func (c seqTimerCtx) B() float64 {
	if c.parent.baseTimer() {
		return c.seq.ToRange().B()
	}
	return c.parent.B()
}

func (c seqTimerCtx) T() float64 {
	if c.parent.baseTimer() {
		return c.seq.ToRange().T()
	}
	return c.parent.T()
}

func (c seqTimerCtx) Ordinal() int {
	if c.parent.baseTimer() {
		return c.seq.ToRange().Ordinal()
	}
	return c.parent.Ordinal()
}

func (c seqTimerCtx) StartB() float64 {
	if c.parent.baseTimer() {
		return c.seq.ToRange().StartB()
	}
	return c.parent.StartB()
}

func (c seqTimerCtx) EndB() float64 {
	if c.parent.baseTimer() {
		return c.seq.ToRange().EndB()
	}
	return c.parent.EndB()
}

func (c seqTimerCtx) Duration() float64 {
	if c.parent.baseTimer() {
		return c.seq.ToRange().Duration()
	}
	return c.parent.Duration()
}

func (c seqTimerCtx) First() bool {
	if c.parent.baseTimer() {
		return c.seq.ToRange().First()
	}
	return c.parent.First()
}

func (c seqTimerCtx) Last() bool {
	if c.parent.baseTimer() {
		return c.seq.ToRange().Last()
	}
	return c.parent.Last()
}
