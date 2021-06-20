package ilysa

import (
	"image/color"

	"github.com/shasderias/ilysa/beatsaber"
)

type CompoundRGBLightingEvent []*RGBLightingEvent
type CompoundRGBLightingEventOpt interface {
	applyCompoundRGBLightingEvent(*CompoundRGBLightingEvent)
}

func NewCompoundRGBLightingEvent(events ...*RGBLightingEvent) *CompoundRGBLightingEvent {
	compoundEvent := CompoundRGBLightingEvent{}
	compoundEvent = append(compoundEvent, events...)
	return &compoundEvent
}

func (e *CompoundRGBLightingEvent) Add(events ...*RGBLightingEvent) {
	*e = append(*e, events...)
}

func (e *CompoundRGBLightingEvent) ShiftBeat(offset float64) {
	for i := range *e {
		(*e)[i].Beat += offset
	}
}

func (e *CompoundRGBLightingEvent) SetValue(val beatsaber.EventValue) {
	for i := range *e {
		(*e)[i].SetValue(val)
	}
}

func (e *CompoundRGBLightingEvent) GetColor() color.Color {
	return (*e)[0].GetColor()
}

func (e *CompoundRGBLightingEvent) SetColor(c color.Color) {
	for i := range *e {
		(*e)[i].SetColor(c)
	}
}

func (e *CompoundRGBLightingEvent) GetAlpha() float64 {
	return (*e)[0].GetAlpha()
}

func (e *CompoundRGBLightingEvent) SetAlpha(a float64) {
	for i := range *e {
		(*e)[i].SetAlpha(a)
	}
}

func (e *CompoundRGBLightingEvent) Mod(opts ...CompoundRGBLightingEventOpt) {
	for _, opt := range opts {
		opt.applyCompoundRGBLightingEvent(e)
	}
}
