package context

import (
	"math/rand"

	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func newSeqCtx(parent Context, s Sequencer, callback func(ctx Context)) {
	iter := s.Iterate()
	for iter.Next() {
		ctx := seqCtx{
			Context:   parent,
			seq:       iter,
			fixedRand: rand.Float64(),
		}
		callback(ctx)
	}
}

type seqCtx struct {
	Context
	seq       timer.Sequence
	fixedRand float64
}

func (c seqCtx) Apply(e evt.Event) {
	e.SetBeat(c.B() + c.BOffset())
	c.AddEvents(e)
}

// private
func (c seqCtx) baseCtx() bool { return false }

// public
func (c seqCtx) WSeq(s Sequencer, callback func(ctx Context)) { newSeqCtx(c, s, callback) }
func (c seqCtx) WRng(r Ranger, callback func(ctx Context))    { newRngCtx(c, r, callback) }
func (c seqCtx) WLight(l Light, callback func(ctx LightContext, e evt.Events)) {
	newLightCtx(c, l, callback)
}
func (c seqCtx) WBOffset(o float64) Context { return newBOffsetCtx(c, o) }

// passthrough to sequence timer
func (c seqCtx) SeqT() float64              { return c.seq.SeqT() }
func (c seqCtx) SeqOrdinal() int            { return c.seq.SeqOrdinal() }
func (c seqCtx) SeqLen() int                { return c.seq.SeqLen() }
func (c seqCtx) SeqNextB() float64          { return c.seq.SeqNextB() }
func (c seqCtx) SeqNextBOffset() float64    { return c.seq.SeqNextBOffset() }
func (c seqCtx) SeqPrevB() float64          { return c.seq.SeqPrevB() }
func (c seqCtx) SeqPrevBOffset() float64    { return c.seq.SeqPrevBOffset() }
func (c seqCtx) SeqFirst() bool             { return c.seq.SeqFirst() }
func (c seqCtx) SeqLast() bool              { return c.seq.SeqLast() }
func (c seqCtx) Next() bool                 { return c.seq.Next() }
func (c seqCtx) ToRange() timer.Range       { return c.seq.ToRange() }
func (c seqCtx) ToSequence() timer.Sequence { return c.seq.ToSequence() }

// pass up to range context, fallback to conversion if closest range timer is base
func (c seqCtx) B() float64 {
	//if c.Context.baseCtx() {
	//	b := c.seq.ToRange().B()
	//	return b
	//}
	return c.seq.ToRange().B()
}

func (c seqCtx) BOffset() float64 { return c.Context.BOffset() + c.Context.B() }

func (c seqCtx) T() float64 {
	if c.Context.baseCtx() {
		return c.seq.ToRange().T()
	}
	return c.Context.T()
}

func (c seqCtx) Ordinal() int {
	if c.Context.baseCtx() {
		return c.seq.ToRange().Ordinal()
	}
	return c.Context.Ordinal()
}

func (c seqCtx) StartB() float64 {
	if c.Context.baseCtx() {
		return c.seq.ToRange().StartB()
	}
	return c.Context.StartB()
}

func (c seqCtx) EndB() float64 {
	if c.Context.baseCtx() {
		return c.seq.ToRange().EndB()
	}
	return c.Context.EndB()
}

func (c seqCtx) Duration() float64 {
	if c.Context.baseCtx() {
		return c.seq.ToRange().Duration()
	}
	return c.Context.Duration()
}

func (c seqCtx) First() bool {
	if c.Context.baseCtx() {
		return c.seq.ToRange().First()
	}
	return c.Context.First()
}

func (c seqCtx) Last() bool {
	if c.Context.baseCtx() {
		return c.seq.ToRange().Last()
	}
	return c.Context.Last()
}
