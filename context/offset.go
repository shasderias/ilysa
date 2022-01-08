package context

import "github.com/shasderias/ilysa/evt"

type bOffsetCtx struct {
	Context
	o float64
}

func newBOffsetCtx(parent Context, offset float64) Context {
	ctx := bOffsetCtx{
		Context: parent,
		o:       offset,
	}
	return ctx
}

func (c bOffsetCtx) Apply(e evt.Event) {
	e.SetBeat(c.B() + c.BOffset())
	c.AddEvents(e)
}
func (c bOffsetCtx) BOffset() float64 { return c.Context.BOffset() + c.o }

func (c bOffsetCtx) WSeq(s Sequencer, callback func(ctx Context)) { newSeqCtx(c, s, callback) }
func (c bOffsetCtx) WRng(r Ranger, callback func(ctx Context))    { newRngCtx(c, r, callback) }
func (c bOffsetCtx) WLight(l Light, callback func(ctx LightContext, e evt.Events)) {
	newLightCtx(c, l, callback)
}
func (c bOffsetCtx) WBOffset(o float64) Context { return newBOffsetCtx(c, o) }
