package context

import (
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func WithRange(parent Context, r timer.Ranger, callback func(ctx Context)) {
	iter := r.Iterate(parent.Offset())
	for iter.Next() {
		ctx := rangeTimerCtx{
			parent: parent,
			Range:  iter,
		}
		ctx.eventer = newEventer(ctx)
		callback(ctx)
	}
}

type rangeTimerCtx struct {
	parent Context
	timer.Range
	eventer
}

func (c rangeTimerCtx) Offset() float64 {
	return c.B()
}

func (c rangeTimerCtx) baseTimer() bool {
	return false
}

// passthrough methods
func (c rangeTimerCtx) addEvents(events ...evt.Event) {
	c.parent.addEvents(events...)
}

// timer methods
func (c rangeTimerCtx) SeqT() float64 {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqT()
	}
	return c.parent.SeqT()
}

func (c rangeTimerCtx) SeqOrdinal() int {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqOrdinal()
	}
	return c.parent.SeqOrdinal()
}

func (c rangeTimerCtx) SeqLen() int {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqLen()
	}
	return c.parent.SeqLen()
}

func (c rangeTimerCtx) SeqNextB() float64 {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqNextB()
	}
	return c.parent.SeqNextB()
}

func (c rangeTimerCtx) SeqNextBOffset() float64 {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqNextBOffset()
	}
	return c.parent.SeqNextBOffset()
}

func (c rangeTimerCtx) SeqPrevB() float64 {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqPrevB()
	}
	return c.parent.SeqPrevB()
}

func (c rangeTimerCtx) SeqPrevBOffset() float64 {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqPrevBOffset()
	}
	return c.parent.SeqPrevBOffset()
}

func (c rangeTimerCtx) SeqFirst() bool {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqFirst()
	}
	return c.parent.SeqFirst()
}

func (c rangeTimerCtx) SeqLast() bool {
	if c.parent.baseTimer() {
		return c.Range.ToSequence().SeqLast()
	}
	return c.parent.SeqLast()
}
