package context

import (
	"math/rand"

	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func newRngCtx(parent Context, r Ranger, callback func(ctx Context)) {
	fixedRand := rand.Float64()
	iter := r.Iterate()
	for iter.Next() {
		ctx := rngCtx{
			Context:   parent,
			rng:       iter,
			fixedRand: fixedRand,
		}
		callback(ctx)
	}
}

type rngCtx struct {
	Context
	rng       timer.Range
	fixedRand float64
}

// private
func (c rngCtx) baseCtx() bool { return false }

func (c rngCtx) FixedRand() float64 { return c.fixedRand }
func (c rngCtx) Apply(e evt.Event) {
	e.SetBeat(c.B() + c.BOffset())
	c.AddEvents(e)
}

// public
func (c rngCtx) WSeq(s Sequencer, callback func(ctx Context)) { newSeqCtx(c, s, callback) }
func (c rngCtx) WRng(r Ranger, callback func(ctx Context))    { newRngCtx(c, r, callback) }
func (c rngCtx) WLight(l Light, callback func(ctx LightContext, e evt.Events)) {
	newLightCtx(c, l, callback)
}
func (c rngCtx) WBOffset(o float64) Context { return newBOffsetCtx(c, o) }

// passthrough to range timer
func (c rngCtx) B() float64                 { return c.rng.B() }
func (c rngCtx) BOffset() float64           { return c.Context.BOffset() + c.Context.B() }
func (c rngCtx) T() float64                 { return c.rng.T() }
func (c rngCtx) Ordinal() int               { return c.rng.Ordinal() }
func (c rngCtx) StartB() float64            { return c.rng.StartB() }
func (c rngCtx) EndB() float64              { return c.rng.EndB() }
func (c rngCtx) Duration() float64          { return c.rng.Duration() }
func (c rngCtx) First() bool                { return c.rng.First() }
func (c rngCtx) Last() bool                 { return c.rng.Last() }
func (c rngCtx) Next() bool                 { return c.rng.Next() }
func (c rngCtx) ToRange() timer.Range       { return c.rng.ToRange() }
func (c rngCtx) ToSequence() timer.Sequence { return c.rng.ToSequence() }

// pass up to closest sequence timer, fallback to conversion if closest sequence timer is base
func (c rngCtx) SeqT() float64 {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqT()
	}
	return c.Context.SeqT()
}

func (c rngCtx) SeqOrdinal() int {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqOrdinal()
	}
	return c.Context.SeqOrdinal()
}

func (c rngCtx) SeqLen() int {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqLen()
	}
	return c.Context.SeqLen()
}

func (c rngCtx) SeqNextB() float64 {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqNextB()
	}
	return c.Context.SeqNextB()
}

func (c rngCtx) SeqNextBOffset() float64 {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqNextBOffset()
	}
	return c.Context.SeqNextBOffset()
}

func (c rngCtx) SeqPrevB() float64 {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqPrevB()
	}
	return c.Context.SeqPrevB()
}

func (c rngCtx) SeqPrevBOffset() float64 {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqPrevBOffset()
	}
	return c.Context.SeqPrevBOffset()
}

func (c rngCtx) SeqFirst() bool {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqFirst()
	}
	return c.Context.SeqFirst()
}

func (c rngCtx) SeqLast() bool {
	if c.Context.baseCtx() {
		return c.rng.ToSequence().SeqLast()
	}
	return c.Context.SeqLast()
}
