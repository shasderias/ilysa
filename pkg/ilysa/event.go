package ilysa

import (
	"encoding/json"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
)

type DirectionalLaser int

const (
	LeftLaser  DirectionalLaser = 0
	RightLaser DirectionalLaser = 1
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

func (e *BaseEvent) ScaleBeat(scaler func(float64) float64) {
	e.Beat = scaler(e.Beat)
}

func (e *BaseEvent) ShiftBeat(offset float64) {
	e.Beat += offset
}

func (e *BaseEvent) Base() *BaseEvent {
	return e
}

func (e *BaseEvent) SetValue(v beatsaber.EventValue) {
	e.Value = v
}

func (e *BaseEvent) SetType(typ beatsaber.EventType) {
	e.Type = typ
}

type withStepOpt struct {
	step float64
}

func WithStep(s float64) withStepOpt {
	return withStepOpt{s}
}

func (o withStepOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.Step = o.step
}

func (o withStepOpt) applyPreciseZoomEvent(e *PreciseZoomEvent) {
	e.Step = o.step
}

type withSpeedOpt struct {
	speed float64
}

func WithSpeed(s float64) withSpeedOpt {
	return withSpeedOpt{s}
}

func (o withSpeedOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.Speed = o.speed
}

func (o withSpeedOpt) applyPreciseRotationSpeedEvent(e *PreciseRotationSpeedEvent) {
	e.Speed = o.speed
}

type withDirectionOpt struct {
	direction chroma.SpinDirection
}

func WithDirection(d chroma.SpinDirection) withDirectionOpt {
	return withDirectionOpt{d}
}

func (o withDirectionOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.Direction = o.direction
}

func (o withDirectionOpt) applyPreciseRotationSpeedEvent(e *PreciseRotationSpeedEvent) {
	e.Direction = o.direction
}
