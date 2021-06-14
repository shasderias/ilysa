package ilysa

import (
	"encoding/json"
	"fmt"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
)

type Event interface {
	Base() *BaseEvent
	CustomData() (json.RawMessage, error)
}

type BaseEvent struct {
	Beat  float64
	Type  beatsaber.EventType
	Value beatsaber.EventValue
}

func (e *BaseEvent) Base() *BaseEvent {
	return e
}

type RotationEvent struct {
	BaseEvent
}

func (c *Context) NewRotationEvent() *RotationEvent {
	e := &RotationEvent{BaseEvent: BaseEvent{
		Beat:  c.B,
		Type:  beatsaber.EventTypeRingSpin,
		Value: 0,
	}}

	c.applyModifiers(e)
	c.events = append(c.events, e)
	return e
}

func (e *RotationEvent) CustomData() (json.RawMessage, error) { return nil, nil }

type PreciseRotationEvent struct {
	BaseEvent
	chroma.PreciseRotation
}

func (c *Context) NewPreciseRotationEvent() *PreciseRotationEvent {
	e := &PreciseRotationEvent{
		BaseEvent: BaseEvent{
			Beat:  c.B,
			Type:  beatsaber.EventTypeRingSpin,
			Value: 0,
		},
	}

	c.applyModifiers(e)
	c.events = append(c.events, e)
	return e
}

func (e *PreciseRotationEvent) SetNameFilter(nf string) *PreciseRotationEvent {
	e.NameFilter = nf
	return e
}
func (e *PreciseRotationEvent) SetReset(r bool) *PreciseRotationEvent {
	e.Reset = r
	return e
}
func (e *PreciseRotationEvent) SetRotation(r float64) *PreciseRotationEvent {
	e.Rotation = r
	return e
}
func (e *PreciseRotationEvent) SetStep(s float64) *PreciseRotationEvent {
	e.Step = s
	return e
}
func (e *PreciseRotationEvent) SetProp(p float64) *PreciseRotationEvent {
	e.Prop = p
	return e
}
func (e *PreciseRotationEvent) SetSpeed(s float64) *PreciseRotationEvent {
	e.Speed = s
	return e
}
func (e *PreciseRotationEvent) SetDirection(d chroma.SpinDirection) *PreciseRotationEvent {
	e.Direction = d
	return e
}
func (e *PreciseRotationEvent) SetCounterSpin(c bool) *PreciseRotationEvent {
	e.CounterSpin = c
	return e
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
	c.applyModifiers(e)
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
	c.applyModifiers(e)
	c.events = append(c.events, e)
	return e
}

func (e *PreciseZoomEvent) SetStep(s float64) *PreciseZoomEvent {
	e.Step = s
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

func (c Context) NewPreciseRotationSpeedEvent() *PreciseRotationSpeedEvent {
	e := &PreciseRotationSpeedEvent{
		BaseEvent: BaseEvent{
			Beat: c.B,
		},
	}
	c.applyModifiers(e)
	c.events = append(c.events, e)
	return e
}

func (e *PreciseRotationSpeedEvent) SetLaser(l DirectionalLaser) *PreciseRotationSpeedEvent {
	var typ beatsaber.EventType
	switch l {
	case LeftLaser:
		typ = beatsaber.EventTypeLeftRotatingLasersRotationSpeed
	case RightLaser:
		typ = beatsaber.EventTypeRightRotatingLasersRotationSpeed
	default:
		panic(fmt.Sprintf("NewRotationSpeedEvent: unsupported direction %v", typ))
	}

	e.Type = typ
	return e
}

func (e *PreciseRotationSpeedEvent) SetValue(value int) *PreciseRotationSpeedEvent {
	e.Value = beatsaber.EventValue(value)
	return e
}

func (e *PreciseRotationSpeedEvent) SetLockPosition(lp bool) *PreciseRotationSpeedEvent {
	e.LockPosition = lp
	return e
}

func (e *PreciseRotationSpeedEvent) SetSpeed(s float64) *PreciseRotationSpeedEvent {
	e.Speed = s
	return e
}

func (e *PreciseRotationSpeedEvent) SetDirection(d chroma.SpinDirection) *PreciseRotationSpeedEvent {
	e.Direction = d
	return e
}
