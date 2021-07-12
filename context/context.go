package context

import (
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type Context interface {
	timer.Sequence
	timer.Range
	FixedRand() float64

	Eventer
	MaxLightID(t evt.LightType) int

	BOffset(float64) Context
	Sequence(s timer.Sequencer, callback func(ctx Context))
	Range(r timer.Ranger, callback func(ctx Context))
	Light(l Light, callback func(ctx LightContext))

	addEvents(events ...evt.Event)
	baseTimer() bool
	offset() float64
}

type LightContext interface {
	timer.Sequence
	timer.Range
	timer.Light
	FixedRand() float64

	LightEventer
}

type Eventer interface {
	NewLighting(opts ...evt.LightingOpt) *evt.Lighting
	NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting
	NewLaser(opts ...evt.LaserOpt) *evt.Laser
	NewPreciseLaser(opts ...evt.PreciseLaserOpt) *evt.PreciseLaser
	NewRotation(opts ...evt.RotationOpt) *evt.Rotation
	NewPreciseRotation(opts ...evt.PreciseRotationOpt) *evt.PreciseRotation
	NewZoom(opts ...evt.ZoomOpt) *evt.Zoom
	NewPreciseZoom(opts ...evt.PreciseZoomOpt) *evt.PreciseZoom
}

type LightEventer interface {
	NewRGBLighting(opts ...evt.RGBLightingOpt) evt.RGBLightingEvents
}

type LightRGBLightingContext interface {
	timer.Sequence
	timer.Range
	timer.Light
	FixedRand() float64
	NewRGBLighting(opts ...evt.RGBLightingOpt) *evt.RGBLighting
}

type Light interface {
	NewRGBLighting(ctx LightRGBLightingContext) evt.RGBLightingEvents
	LightIDLen() int
}
