package context

import (
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

func (c eventer) NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting {
	e := evt.NewRGBLighting()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewLaser(opts ...evt.LaserOpt) *evt.Laser {
	e := evt.NewLaser()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser {
	e := evt.NewPreciseLaser()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewRotation(opts ...evt.RotationOpt) *evt.Rotation {
	e := evt.NewRotation()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation {
	e := evt.NewPreciseRotation()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewZoom(opts ...evt.ZoomOpt) *evt.Zoom {
	e := evt.NewZoom()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom {
	e := evt.NewPreciseZoom()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	evt.Apply(&e, c.defaultOptsPost...)
	c.ctx.addEvents(&e)
	return &e
}
