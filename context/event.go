package context

import (
	"fmt"

	"github.com/shasderias/ilysa/evt"
)

type eventer struct {
	ctx             Context
	defaultOptsPre  []evt.Opt
	defaultOptsPost []evt.Opt
}

func newEventer(ctx Context) eventer {
	ec := eventer{ctx: ctx}

	ec.defaultOptsPre = []evt.Opt{evt.WithBeat(ctx.NoOffsetB())}
	ec.defaultOptsPost = []evt.Opt{evt.WithBeatOffset(ctx.Offset() - ctx.NoOffsetB())}

	return ec
}

func (c eventer) NewLighting(opts ...evt.LightingOpt) *evt.Lighting {
	e := evt.NewLighting()
	evt.Apply(&e, c.defaultOptsPre...)
	e.Apply(opts...)
	fmt.Println(e.Beat())
	evt.Apply(&e, c.defaultOptsPost...)
	fmt.Println(e.Beat())
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting {
	e := evt.NewRGBLighting()
	e.Apply(opts...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewLaser(opts ...evt.LaserOpt) *evt.Laser {
	e := evt.NewLaser()
	e.Apply(opts...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser {
	e := evt.NewPreciseLaser()
	e.Apply(opts...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewRotation(opts ...evt.RotationOpt) *evt.Rotation {
	e := evt.NewRotation()
	e.Apply(opts...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation {
	e := evt.NewPreciseRotation()
	e.Apply(opts...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewZoom(opts ...evt.ZoomOpt) *evt.Zoom {
	e := evt.NewZoom()
	e.Apply(opts...)
	c.ctx.addEvents(&e)
	return &e
}

func (c eventer) NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom {
	e := evt.NewPreciseZoom()
	e.Apply(opts...)
	c.ctx.addEvents(&e)
	return &e
}
