package context

import (
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type offsetCtx struct {
	Context
	o float64
	eventer
}

func WithBOffset(parent Context, offset float64) Context {
	ctx := offsetCtx{
		Context: parent,
		o:       offset,
	}
	ctx.eventer = newEventer(ctx)
	return ctx
}

func (c offsetCtx) Sequence(s timer.Sequencer, callback func(ctx Context)) {
	WithSequence(c, s, callback)
}
func (c offsetCtx) Range(r timer.Ranger, callback func(ctx Context)) {
	WithRange(c, r, callback)
}
func (c offsetCtx) Light(l Light, callback func(ctx LightContext)) {
	WithLight(c, l, callback)
}
func (c offsetCtx) BOffset(o float64) Context {
	return WithBOffset(c, o)
}

func (c offsetCtx) offset() float64 { return c.Context.offset() + c.o }

func (c offsetCtx) NewLighting(opts ...evt.LightingOpt) *evt.Lighting {
	return c.eventer.NewLighting(opts...)
}
func (c offsetCtx) NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting {
	return c.eventer.NewRGBLighting(opts...)
}
func (c offsetCtx) NewLaser(opts ...evt.LaserOpt) *evt.Laser {
	return c.eventer.NewLaser(opts...)
}
func (c offsetCtx) NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser {
	return c.eventer.NewPreciseLaser(opts...)
}
func (c offsetCtx) NewRotation(opts ...evt.RotationOpt) *evt.Rotation {
	return c.eventer.NewRotation(opts...)
}
func (c offsetCtx) NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation {
	return c.eventer.NewPreciseRotation(opts...)
}
func (c offsetCtx) NewZoom(opts ...evt.ZoomOpt) *evt.Zoom {
	return c.eventer.NewZoom(opts...)
}
func (c offsetCtx) NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom {
	return c.eventer.NewPreciseZoom(opts...)
}
