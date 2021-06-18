package ilysa

import "ilysa/pkg/beatsaber"

type withTypeOpt struct {
	typ beatsaber.EventType
}

func WithType(t beatsaber.EventType) withTypeOpt {
	return withTypeOpt{t}
}

func (w withTypeOpt) applyBasicLightingEvent(e *BasicLightingEvent) {
	e.SetType(w.typ)
}

func (w withTypeOpt) applyRGBLightingEvent(e *RGBLightingEvent) {
	e.SetType(w.typ)
}

//func (w withTypeOpt) applyRotationEvent(e *RotationEvent) {
//	e.SetType(w.typ)
//}
//
//func (w withTypeOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
//	e.SetType(w.typ)
//}
//
//func (w withTypeOpt) applyZoomEvent(e *ZoomEvent) {
//	e.SetType(w.typ)
//}
//
//func (w withTypeOpt) applyPreciseZoomEvent(e *PreciseZoomEvent) {
//	e.SetType(w.typ)
//}
//
//func (w withTypeOpt) applyRotationSpeedEvent(e *RotationSpeedEvent) {
//	e.SetType(w.typ)
//}
//
//func (w withTypeOpt) applyPreciseRotationSpeedEvent(e *PreciseRotationSpeedEvent) {
//	e.SetType(w.typ)
//}

type withValueOpt struct {
	typ beatsaber.EventValue
}

func WithValue(t beatsaber.EventValue) withValueOpt {
	return withValueOpt{t}
}

func (w withValueOpt) applyBasicLightingEvent(e *BasicLightingEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyRGBLightingEvent(e *RGBLightingEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyCompoundRGBLightingEvent(e *CompoundRGBLightingEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyRotationEvent(e *RotationEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyZoomEvent(e *ZoomEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyPreciseZoomEvent(e *PreciseZoomEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyRotationSpeedEvent(e *RotationSpeedEvent) {
	e.SetValue(w.typ)
}

func (w withValueOpt) applyPreciseRotationSpeedEvent(e *PreciseRotationSpeedEvent) {
	e.SetValue(w.typ)
}