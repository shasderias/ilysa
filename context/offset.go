package context

import "github.com/shasderias/ilysa/evt"

type offsetCtx struct {
	Context
	offset float64
	eventer
}

func WithOffset(parent Context, offset float64) Context {
	ctx := offsetCtx{
		Context: parent,
		offset:  offset,
	}
	ctx.eventer = newEventer(ctx)
	return ctx
}

func (ctx offsetCtx) Offset() float64 {
	return ctx.Context.Offset() + ctx.offset
}

func (ctx offsetCtx) NewLighting(opts ...evt.LightingOpt) *evt.Lighting {
	return ctx.eventer.NewLighting(opts...)
}

func (ctx offsetCtx) NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting {
	return ctx.eventer.NewRGBLighting(opts...)
}

//func (ctx offsetCtx) NewLaser(opts ...evt.LaserOpt) *evt.Laser {
//	return ctx.eventer.NewLaser(opts...)
//}
//
//func (ctx offsetCtx) NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser {
//	return ctx.eventer.NewPreciseLaser(opts...)
//}
//
//func (ctx offsetCtx) NewRotation(opts ...evt.RotationOpt) *evt.Rotation {
//	return ctx.eventer.NewRotation(opts...)
//}
//
//func (ctx offsetCtx) NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation {
//	return ctx.eventer.NewPreciseRotation(opts...)
//}
//
//func (ctx offsetCtx) NewZoom(opts ...evt.ZoomOpt) *evt.Zoom {
//	return ctx.eventer.NewZoom(opts...)
//}
//
//func (ctx offsetCtx) NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom {
//	return ctx.eventer.NewPreciseZoom(opts...)
//}
