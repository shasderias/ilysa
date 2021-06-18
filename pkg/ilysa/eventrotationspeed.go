package ilysa

import (
	"encoding/json"
	"fmt"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/ilysa/event"
)

type RotationSpeedEvent struct {
	event.BaseEvent
}

type RotationSpeedEventOpt interface {
	applyRotationSpeedEvent(*RotationSpeedEvent)
}

func (c baseContext) NewRotationSpeedEvent(opts ...RotationSpeedEventOpt) *RotationSpeedEvent {
	e := &RotationSpeedEvent{
		event.BaseEvent{
			Beat:  c.B(),
			Type:  beatsaber.EventTypeLeftRotatingLasersRotationSpeed,
			Value: 0,
		},
	}
	e.Mod(opts...)
	c.addEvent(e)
	return e
}

func (e *RotationSpeedEvent) Mod(opts ...RotationSpeedEventOpt) {
	for _, opt := range opts {
		opt.applyRotationSpeedEvent(e)
	}
}

func (e RotationSpeedEvent) CustomData() (json.RawMessage, error) { return nil, nil }

type PreciseRotationSpeedEvent struct {
	event.BaseEvent
	chroma.PreciseLaser
}

type PreciseRotationSpeedEventOpt interface {
	applyPreciseRotationSpeedEvent(event *PreciseRotationSpeedEvent)
}

func (c baseContext) NewPreciseRotationSpeedEvent(opts ...PreciseRotationSpeedEventOpt) *PreciseRotationSpeedEvent {
	e := &PreciseRotationSpeedEvent{
		BaseEvent: event.BaseEvent{
			Beat: c.B(),
		},
	}
	for _, opt := range opts {
		opt.applyPreciseRotationSpeedEvent(e)
	}
	c.addEvent(e)
	return e
}

func (e *PreciseRotationSpeedEvent) Mod(opts ...PreciseRotationSpeedEventOpt) {
	for _, opt := range opts {
		opt.applyPreciseRotationSpeedEvent(e)
	}
}

type withLockPositionOpt struct {
	lp bool
}

func WithLockPosition(lp bool) withLockPositionOpt {
	return withLockPositionOpt{lp}
}

func (o withLockPositionOpt) applyPreciseRotationSpeedEvent(e *PreciseRotationSpeedEvent) {
	e.LockPosition = o.lp
}

type withDirectionalLaserOpt struct {
	dl  DirectionalLaser
	typ beatsaber.EventType
}

func WithDirectionalLaser(dl DirectionalLaser) withDirectionalLaserOpt {
	o := withDirectionalLaserOpt{dl: dl}
	switch dl {
	case LeftLaser:
		o.typ = beatsaber.EventTypeLeftRotatingLasersRotationSpeed
	case RightLaser:
		o.typ = beatsaber.EventTypeRightRotatingLasersRotationSpeed
	default:
		panic(fmt.Sprintf("WithDirectionalLaser: unsupported direction %v", dl))
	}
	return o
}

func (o withDirectionalLaserOpt) applyRotationSpeedEvent(e *RotationSpeedEvent) {
	e.SetType(o.typ)
}

func (o withDirectionalLaserOpt) applyPreciseRotationSpeedEvent(e *PreciseRotationSpeedEvent) {
	e.SetType(o.typ)
}
