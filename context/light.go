package context

import (
	"github.com/shasderias/ilysa/evt"
)

type lightCtx struct {
	Context
	ordinal int
	l       Light
}

func (c *lightCtx) LightT() float64 {
	if c.LightLen() == 1 {
		return 1
	}
	return float64(c.ordinal) / float64(c.LightLen()-1)
}
func (c *lightCtx) LightOrdinal() int { return c.ordinal }
func (c *lightCtx) LightLen() int     { return c.l.LightLen() }
func (c *lightCtx) LightCur() int     { return c.ordinal + 1 }
func (c *lightCtx) Next() bool {
	c.ordinal++
	if c.ordinal == c.l.LightLen() {
		return false
	}
	return true
}
func (c *lightCtx) Apply(e evt.Event) {
	e.SetBeat(c.B() + c.BOffset())
	e.SetTag(c.l.Name()...)
	c.AddEvents(e)
}

func newLightCtx(parent Context, l Light, callback func(ctx LightContext, e evt.Events)) {
	lctx := &lightCtx{
		Context: parent,
		ordinal: -1,
		l:       l,
	}
	for lctx.Next() {
		events := l.GenerateEvents(lctx)
		callback(lctx, events)
	}
}

func LightContextAtOrdinal(ctx LightContext, l Light, ordinal int) LightContext {
	lctx, ok := ctx.(*lightCtx)
	if !ok {
		panic("LightContextAtOrdinal: invalid context")
	}
	return &lightCtx{
		Context: lctx.Context,
		ordinal: ordinal,
		l:       l,
	}
}
