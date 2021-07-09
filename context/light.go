package context

import (
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func WithLight(parent Context, l Light, callback func(ctx LightContext)) {
	iter := timer.NewLighter(l).Iterate()
	for iter.Next() {
		ctx := lightTimerCtx{
			Context: parent,
			Light:   iter,
			l:       l,
		}
		callback(ctx)
	}
}

type lightTimerCtx struct {
	Context
	timer.Light
	l Light
}

func (c lightTimerCtx) Next() bool {
	return c.Light.Next()
}

//func (c lightTimerCtx) offset() float64 {
//	return c.Context.offset()
//}

func (c lightTimerCtx) NewRGBLighting(opts ...evt.RGBLightingOpt) evt.RGBLightingEvents {
	e := c.l.NewRGBLighting(newLightCtx(c.Context, c.Light, c.l, opts))
	return e
}

func newLightCtx(ctx Context, lightTimer timer.Light, l Light, opts []evt.RGBLightingOpt) lightCtx {
	lightCtx := lightCtx{
		Context: ctx,
		Light:   lightTimer,
		l:       l,
	}

	lightCtx.defaultOptsPre = []evt.Opt{evt.WithBeat(ctx.B())}
	lightCtx.userOpts = opts
	lightCtx.defaultOptsPost = []evt.Opt{evt.WithBeatOffset(ctx.offset() - ctx.B())}
	return lightCtx
}

type lightCtx struct {
	Context
	timer.Light
	l Light

	defaultOptsPre  []evt.Opt
	userOpts        []evt.RGBLightingOpt
	defaultOptsPost []evt.Opt
}

func (c lightCtx) NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting {
	e := evt.NewRGBLighting()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	e.Apply(c.userOpts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.addEvents(&e)
	return &e
}

func (c lightCtx) Next() bool {
	return c.Light.Next()
}
