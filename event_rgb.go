package ilysa

import (
	"fmt"
	"image/color"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
)

type RGBLightingEvent struct {
	BaseEvent
	chroma.RGB
}

type RGBLightingEventOpt interface {
	applyRGBLightingEvent(*RGBLightingEvent)
}

func (c baseContext) NewRGBLightingEvent(opts ...RGBLightingEventOpt) *RGBLightingEvent {
	e := &RGBLightingEvent{
		BaseEvent: BaseEvent{
			beat: c.B(),
			val:  beatsaber.EventValueLightRedOn,
		},
	}
	for _, opt := range opts {
		opt.applyRGBLightingEvent(e)
	}
	c.addEvent(e)
	return e
}

func (e *RGBLightingEvent) SetLight(typ beatsaber.EventTyper) *RGBLightingEvent {
	if !beatsaber.IsLightingEvent(typ.EventType()) {
		panic(fmt.Sprintf("context.NewRGBLightingEvent: %v is not c lighting event", typ))
	}
	e.typ = typ.EventType()
	return e
}

func (e *RGBLightingEvent) GetColor() color.Color {
	return e.Color
}

func (e *RGBLightingEvent) SetColor(c color.Color) {
	e.Color = c
}

func (e *RGBLightingEvent) GetAlpha() float64 {
	c := colorful.FromColor(e.Color)
	return c.A
}

func (e *RGBLightingEvent) SetAlpha(a float64) {
	c := colorful.FromColor(e.Color)
	c.A = a
	e.Color = c
}

func (e *RGBLightingEvent) SetSingleLightID(id int) *RGBLightingEvent {
	e.LightID = []int{id}
	return e
}

func (e *RGBLightingEvent) SetLightID(id LightID) *RGBLightingEvent {
	e.LightID = chroma.LightID(id)
	return e
}

func (e *RGBLightingEvent) Mod(opts ...RGBLightingEventOpt) {
	for _, opt := range opts {
		opt.applyRGBLightingEvent(e)
	}
}
