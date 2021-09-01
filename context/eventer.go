package context

import (
	"math"

	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/evt"
)

type eventer struct {
	ctx             Context
	defaultOptsPre  []evt.Opt
	defaultOptsPost []evt.Opt
}

func newEventer(ctx Context) eventer {
	ec := eventer{ctx: ctx}

	ec.defaultOptsPre = []evt.Opt{evt.WithBeat(ctx.B())}
	ec.defaultOptsPost = []evt.Opt{evt.WithBOffset(ctx.offset() - ctx.B())}

	return ec
}

func (c eventer) NewLighting(opts ...evt.LightingOpt) *evt.Lighting {
	e := evt.NewLighting()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZLighting(typ evt.LightType, val evt.LightValue) *evt.Lighting {
	return c.NewLighting(
		evt.WithLight(typ), evt.WithLightValue(val),
	)
}

func (c eventer) NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting {
	e := evt.NewRGBLighting()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZRGBLighting(color colorful.Color) *evt.RGBLighting {
	return c.NewRGBLighting(
		evt.WithColor(color),
	)
}

func (c eventer) NewLaser(opts ...evt.LaserOpt) *evt.Laser {
	e := evt.NewLaser()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZLaser(laser evt.DirectionalLaser, speed int) *evt.Laser {
	return c.NewLaser(
		evt.WithDirectionalLaser(laser),
		evt.WithLaserSpeed(speed),
	)
}

func (c eventer) NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser {
	e := evt.NewPreciseLaser()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZPreciseLaser(laser evt.DirectionalLaser, speed float64) *evt.PreciseLaser {
	return c.NewPreciseLaser(
		evt.WithDirectionalLaser(laser),
		evt.WithLaserSpeed(int(math.RoundToEven(speed))),
		evt.WithPreciseLaserSpeed(speed),
	)
}

func (c eventer) NewRotation(opts ...evt.RotationOpt) *evt.Rotation {
	e := evt.NewRotation()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZRotation() *evt.Rotation {
	return c.NewRotation()
}

func (c eventer) NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation {
	e := evt.NewPreciseRotation()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZPreciseRotation(rotation, step, prop, speed float64, direction chroma.SpinDirection) *evt.PreciseRotation {
	return c.NewPreciseRotation(
		evt.WithRotation(rotation),
		evt.WithRotationStep(step),
		evt.WithProp(prop),
		evt.WithRotationSpeed(speed),
		evt.WithRotationDirection(direction),
	)
}

func (c eventer) NewZoom(opts ...evt.ZoomOpt) *evt.Zoom {
	e := evt.NewZoom()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZZoom() *evt.Zoom {
	return c.NewZoom()
}

func (c eventer) NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom {
	e := evt.NewPreciseZoom()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) EZPreciseZoom(step float64) *evt.PreciseZoom {
	return c.NewPreciseZoom(
		evt.WithZoomStep(step),
	)
}

func (c eventer) NewChromaGradient(opts ...evt.ChromaGradientOpt) *evt.ChromaGradient {
	e := evt.NewChromaGradient()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}
