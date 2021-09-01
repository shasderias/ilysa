package context

import (
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/ease"
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
	Beat(beat float64, callback func(ctx Context))
	BeatSequence(seq []float64, ghostBeat float64, callback func(ctx Context))
	BeatRange(startB, endB float64, steps int, easeFn ease.Func, callback func(ctx Context))
	Light(l Light, callback func(ctx LightContext))

	addEvents(events ...evt.Event)
	base() base
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
	NewChromaGradient(opts ...evt.ChromaGradientOpt) *evt.ChromaGradient

	EZLighting(typ evt.LightType, val evt.LightValue) *evt.Lighting
	EZRGBLighting(color colorful.Color) *evt.RGBLighting
	EZLaser(laser evt.DirectionalLaser, speed int) *evt.Laser
	EZPreciseLaser(laser evt.DirectionalLaser, speed float64) *evt.PreciseLaser
	EZRotation() *evt.Rotation
	EZPreciseRotation(rotation, step, prop, speed float64, direction chroma.SpinDirection) *evt.PreciseRotation
	EZZoom() *evt.Zoom
	EZPreciseZoom(step float64) *evt.PreciseZoom
}

type LightEventer interface {
	NewRGBLighting(opts ...evt.RGBLightingOpt) evt.RGBLightingEvents
	EZRGBLighting(color colorful.Color) evt.RGBLightingEvents
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
