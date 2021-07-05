package context

import (
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func WithSequence(parent Context, s timer.Sequencer, callback func(ctx Context)) {
	iter := s.Iterate(parent.Offset())
	for iter.Next() {
		ctx := seqTimerCtx{
			parent:   parent,
			Sequence: iter,
		}
		ctx.eventer = newEventer(ctx)
		callback(ctx)
	}
}

type seqTimerCtx struct {
	parent Context
	timer.Sequence
	eventer
}

func (c seqTimerCtx) Offset() float64 {
	return c.B()
}

func (c seqTimerCtx) baseTimer() bool {
	return false
}

// passthrough methods
func (c seqTimerCtx) addEvents(events ...evt.Event) {
	c.parent.addEvents(events...)
}

// timer methods
func (c seqTimerCtx) B() float64 {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().B()
	}
	return c.parent.B()
}

func (c seqTimerCtx) NoOffsetB() float64 {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().NoOffsetB()
	}
	return c.parent.NoOffsetB()
}

func (c seqTimerCtx) T() float64 {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().T()
	}
	return c.parent.T()
}

func (c seqTimerCtx) Ordinal() int {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().Ordinal()
	}
	return c.parent.Ordinal()
}

func (c seqTimerCtx) StartB() float64 {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().StartB()
	}
	return c.parent.StartB()
}

func (c seqTimerCtx) EndB() float64 {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().EndB()
	}
	return c.parent.EndB()
}

func (c seqTimerCtx) Duration() float64 {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().Duration()
	}
	return c.parent.Duration()
}

func (c seqTimerCtx) First() bool {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().First()
	}
	return c.parent.First()
}

func (c seqTimerCtx) Last() bool {
	if c.parent.baseTimer() {
		return c.Sequence.ToRange().Last()
	}
	return c.parent.Last()
}
