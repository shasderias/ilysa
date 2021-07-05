package ilysa

import (
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
)

type sequenceContext struct {
	baseContext
	sequence timer.Sequencer
	timer.SequenceIterator
	timer.Range
}

func newSequenceContext(c baseContext, s timer.Sequencer) sequenceContext {
	return sequenceContext{
		baseContext: c,
		sequence:    s,
	}
}

type sequenceContextOpt func(ctx *sequenceContext)

func (c sequenceContext) with(opts ...sequenceContextOpt) sequenceContext {
	newCtx := sequenceContext{
		baseContext:      c.baseContext,
		sequence:         c.sequence,
		SequenceIterator: c.SequenceIterator,
		Range:            c.Range,
	}

	for _, opt := range opts {
		opt(&newCtx)
	}

	return newCtx
}

func (c sequenceContext) iterate(callback func(ctx SequenceContext)) {
	iter := c.sequence.Iterate()
	for iter.Next() {
		callback(c.with(func(ctx *sequenceContext) {
			timer := iter.ToRangeTimer()
			ctx.baseContext = ctx.baseContext.withDefaultOpts(evt.WithBeat(timer.B()))
			ctx.SequenceIterator = iter
			ctx.Range = timer
		}))
	}
}

func (c sequenceContext) EventsForRange(startBeat, endBeat float64, steps int, fn ease.Func, callback func(ctx SequenceTimeContext)) {

}

func (c sequenceContext) WithLight(l light.Light, callback func(ctx SequenceLightContext)) {
	//for i := 0; i < l.LightIDLen(); i++ {
	//	callback(newSequenceLightContext(c.baseContext, c, l, i))
	//}
}
