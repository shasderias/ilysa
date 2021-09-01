package context

import (
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/ease"
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
func (c offsetCtx) Beat(beat float64, callback func(ctx Context)) {
	WithSequence(c, timer.Beat(beat), callback)
}
func (c offsetCtx) BeatSequence(seq []float64, ghostBeat float64, callback func(ctx Context)) {
	WithSequence(c, timer.Seq(seq, ghostBeat), callback)
}
func (c offsetCtx) BeatInterval(startBeat, duration float64, count int, callback func(ctx Context)) {
	WithSequence(c, timer.Interval(startBeat, duration, count), callback)
}
func (c offsetCtx) BeatRange(startB, endB float64, steps int, easeFn ease.Func, callback func(ctx Context)) {
	WithRange(c, timer.Rng(startB, endB, steps, easeFn), callback)
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
func (c offsetCtx) NewChromaGradient(opts ...evt.ChromaGradientOpt) *evt.ChromaGradient {
	return c.eventer.NewChromaGradient(opts...)
}

func (c offsetCtx) EZLighting(typ evt.LightType, val evt.LightValue) *evt.Lighting {
	return c.eventer.EZLighting(typ, val)
}
func (c offsetCtx) EZRGBLighting(color colorful.Color) *evt.RGBLighting {
	return c.eventer.EZRGBLighting(color)
}
func (c offsetCtx) EZLaser(laser evt.DirectionalLaser, speed int) *evt.Laser {
	return c.eventer.EZLaser(laser, speed)
}
func (c offsetCtx) EZPreciseLaser(laser evt.DirectionalLaser, speed float64) *evt.PreciseLaser {
	return c.eventer.EZPreciseLaser(laser, speed)
}
func (c offsetCtx) EZRotation() *evt.Rotation {
	return c.eventer.EZRotation()
}
func (c offsetCtx) EZPreciseRotation(rotation, step, prop, speed float64, direction chroma.SpinDirection) *evt.PreciseRotation {
	return c.eventer.EZPreciseRotation(rotation, step, prop, speed, direction)
}
func (c offsetCtx) EZZoom() *evt.Zoom {
	return c.eventer.EZZoom()
}
func (c offsetCtx) EZPreciseZoom(step float64) *evt.PreciseZoom {
	return c.eventer.EZPreciseZoom(step)
}
