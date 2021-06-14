package ilysa

import (
	"encoding/json"
	"fmt"
	"image/color"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/colorful"
)

type LightingEvent struct {
	BaseEvent
}

func (c *Context) NewLightingEvent(typ beatsaber.EventType, val beatsaber.EventValue) *LightingEvent {
	e := &LightingEvent{
		BaseEvent: BaseEvent{
			Beat:  c.B,
			Type:  typ,
			Value: val,
		},
	}

	c.events = append(c.events, e)

	return e
}

func (e *LightingEvent) CustomData() (json.RawMessage, error) { return nil, nil }

type RGBLightingEvent struct {
	BaseEvent
	chroma.RGB
}

func (c *Context) NewRGBLightingEvent() *RGBLightingEvent {
	e := &RGBLightingEvent{
		BaseEvent: BaseEvent{
			Beat: c.B,
		},
	}
	c.applyModifiers(e)
	c.events = append(c.events, e)
	return e
}

func (e *RGBLightingEvent) SetLight(typ beatsaber.EventTyper) *RGBLightingEvent {
	if !beatsaber.IsLightingEvent(typ.EventType()) {
		panic(fmt.Sprintf("context.NewRGBLightingEvent: %v is not a lighting event", typ))
	}
	e.Type = typ.EventType()
	return e
}

func (e *RGBLightingEvent) SetValue(val beatsaber.EventValue) *RGBLightingEvent {
	e.Value = val
	return e
}

func (e *RGBLightingEvent) SetColor(c color.Color) *RGBLightingEvent {
	e.Color = c
	return e
}

func (e *RGBLightingEvent) GetAlpha() float64 {
	c := colorful.FromColor(e.Color)
	return c.A
}

func (e *RGBLightingEvent) SetAlpha(a float64) *RGBLightingEvent {
	c := colorful.FromColor(e.Color)
	c.A = a
	e.Color = c
	return e
}

func (e *RGBLightingEvent) FirstLightID() int {
	if len(e.LightID) == 0 {
		panic("RGB.FirstLightID: lightID is nil or contains no lightIDs")
	}
	return e.LightID[0]
}

func (e *RGBLightingEvent) SetSingleLightID(id int) *RGBLightingEvent {
	e.LightID = []int{id}
	return e
}

func (e *RGBLightingEvent) SetLightID(ids chroma.LightID) *RGBLightingEvent {
	e.LightID = ids
	return e
}
