package evt

import "github.com/shasderias/ilysa/beatsaber"

type BaseOpt interface {
	applyBase(e *Base)
}

func WithBeat(b float64) withBeatOpt {
	return withBeatOpt{b}
}

type withBeatOpt struct{ b float64 }

func (o withBeatOpt) apply(e Event) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyBase(e *Base) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyLighting(e *Lighting) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyRGBLighting(e *RGBLighting) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyRotation(e *Rotation) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyPreciseRotation(e *PreciseRotation) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyLaser(e *Laser) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyPreciseLaser(e *PreciseLaser) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyZoom(e *Zoom) {
	e.SetBeat(o.b)
}

func (o withBeatOpt) applyPreciseZoom(e *PreciseZoom) {
	e.SetBeat(o.b)
}

func WithBeatOffset(o float64) withBeatOffsetOpt {
	return withBeatOffsetOpt{o}
}

type withBeatOffsetOpt struct {
	o float64
}

func (o withBeatOffsetOpt) apply(e Event) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyBase(e *Base) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyLighting(e *Lighting) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyRGBLighting(e *RGBLighting) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyRotation(e *Rotation) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyPreciseRotation(e *PreciseRotation) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyLaser(e *Laser) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyPreciseLaser(e *PreciseLaser) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyZoom(e *Zoom) {
	e.SetBeat(e.Beat() + o.o)
}

func (o withBeatOffsetOpt) applyPreciseZoom(e *PreciseZoom) {
	e.SetBeat(e.Beat() + o.o)
}

type withTypeOpt struct{ t beatsaber.EventType }

func WithType(t beatsaber.EventType) BaseOpt {
	return withTypeOpt{t}
}

func (o withTypeOpt) applyBase(e *Base) {
	e.SetType(o.t)
}

type withValueOpt struct{ t beatsaber.EventValue }

func WithValue(t beatsaber.EventValue) BaseOpt {
	return withValueOpt{t}
}

func (o withValueOpt) applyBase(e *Base) {
	e.SetValue(o.t)
}

type withInvalidBaseOpt struct{}

func WithInvalidDefaults() BaseOpt {
	return withInvalidBaseOpt{}
}

func (o withInvalidBaseOpt) applyBase(e *Base) {
	e.SetBeat(-1)
	e.SetType(beatsaber.EventTypeInvalid)
	e.SetValue(beatsaber.EventValueInvalid)
}

type withRGBLightingDefault struct{}

func WithRGBLightingDefaults() BaseOpt {
	return withRGBLightingDefault{}
}

func (o withRGBLightingDefault) applyBase(e *Base) {
	e.SetBeat(-1)
	e.SetType(beatsaber.EventTypeInvalid)
	e.SetValue(beatsaber.EventValueLightRedOn)
}
