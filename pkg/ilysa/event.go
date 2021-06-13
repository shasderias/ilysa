package ilysa

import (
	"encoding/json"
	"fmt"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
)

type Event interface {
	Base() BaseEvent
	CustomData() (json.RawMessage, error)
}

type BaseEvent struct {
	Beat  float64
	Type  beatsaber.EventType
	Value beatsaber.EventValue
}

func (e BaseEvent) Base() BaseEvent {
	return e
}

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

func (c *Context) NewRGBLightingEvent(typ beatsaber.EventType, val beatsaber.EventValue) *RGBLightingEvent {
	if !beatsaber.IsLightingEvent(typ) {
		panic(fmt.Sprintf("context.NewRGBLightingEvent: %v is not a lighting event", typ))
	}

	e := &RGBLightingEvent{
		BaseEvent: BaseEvent{
			Beat:  c.B,
			Type:  typ,
			Value: val,
		},
	}
	c.events = append(c.events, e)
	return e
}

type RotationEvent struct {
	BaseEvent
}

func (c *Context) NewRotationEvent() *RotationEvent {
	event := &RotationEvent{BaseEvent: BaseEvent{
		Beat:  c.B,
		Type:  beatsaber.EventTypeRingSpin,
		Value: 0,
	}}

	c.events = append(c.events, event)

	return event
}

func (e *RotationEvent) CustomData() (json.RawMessage, error) { return nil, nil }

type PreciseRotationEvent struct {
	BaseEvent
	chroma.PreciseRotation
}

func (c *Context) NewPreciseRotationEvent() *PreciseRotationEvent {
	event := &PreciseRotationEvent{
		BaseEvent: BaseEvent{
			Beat:  c.B,
			Type:  beatsaber.EventTypeRingSpin,
			Value: 0,
		},
	}

	c.events = append(c.events, event)

	return event
}

type ZoomEvent struct {
	BaseEvent
}

func (c *Context) NewZoomEvent() *ZoomEvent {
	e := &ZoomEvent{
		BaseEvent{
			Beat:  c.B,
			Type:  beatsaber.EventTypeRingZoom,
			Value: 0,
		},
	}
	c.events = append(c.events, e)
	return e
}

func (e *ZoomEvent) CustomData() (json.RawMessage, error) { return nil, nil }

type PreciseZoomEvent struct {
	BaseEvent
	chroma.PreciseZoom
}

func (c *Context) NewPreciseZoomEvent() *PreciseZoomEvent {
	e := &PreciseZoomEvent{
		BaseEvent: BaseEvent{
			Beat:  c.B,
			Type:  beatsaber.EventTypeRingZoom,
			Value: 0,
		},
	}
	c.events = append(c.events, e)
	return e
}

type RotationSpeedEvent struct {
	BaseEvent
}

type DirectionalLaser int

const (
	LeftLaser  DirectionalLaser = 0
	RightLaser DirectionalLaser = 1
)

func (c *Context) NewRotationSpeedEvent(d DirectionalLaser, value int) *RotationSpeedEvent {
	var typ beatsaber.EventType
	switch d {
	case LeftLaser:
		typ = beatsaber.EventTypeLeftRotatingLasersRotationSpeed
	case RightLaser:
		typ = beatsaber.EventTypeRightRotatingLasersRotationSpeed
	default:
		panic(fmt.Sprintf("NewRotationSpeedEvent: unsupported direction %v", typ))
	}

	e := &RotationSpeedEvent{
		BaseEvent{
			Beat:  c.B,
			Type:  typ,
			Value: beatsaber.EventValue(value),
		},
	}
	c.events = append(c.events, e)
	return e
}

func (e *RotationSpeedEvent) CustomData() (json.RawMessage, error) { return nil, nil }

type PreciseRotationSpeedEvent struct {
	BaseEvent
	chroma.PreciseLaser
}

func (c Context) NewPreciseRotationSpeedEvent(l DirectionalLaser, value int) *PreciseRotationSpeedEvent {
	var typ beatsaber.EventType
	switch l {
	case LeftLaser:
		typ = beatsaber.EventTypeLeftRotatingLasersRotationSpeed
	case RightLaser:
		typ = beatsaber.EventTypeRightRotatingLasersRotationSpeed
	default:
		panic(fmt.Sprintf("NewRotationSpeedEvent: unsupported direction %v", typ))
	}

	e := &PreciseRotationSpeedEvent{
		BaseEvent: BaseEvent{
			Beat:  c.B,
			Type:  typ,
			Value: beatsaber.EventValue(value),
		},
	}
	c.events = append(c.events, e)
	return e
}
