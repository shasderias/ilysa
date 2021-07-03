package ilysa

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

type DirectionalLaser int

const (
	LeftLaser  DirectionalLaser = 0
	RightLaser DirectionalLaser = 1
)

type Event interface {
	Beat() float64
	SetBeat(b float64)
	Type() beatsaber.EventType
	SetType(t beatsaber.EventType)
	Value() beatsaber.EventValue
	SetValue(v beatsaber.EventValue)

	CustomData() (json.RawMessage, error)
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

type BaseEvent struct {
	beat float64
	typ  beatsaber.EventType
	val  beatsaber.EventValue
}

func (e *BaseEvent) ScaleBeat(scaler func(float64) float64) {
	e.beat = scaler(e.beat)
}

func (e *BaseEvent) ShiftBeat(offset float64) {
	e.beat += offset
}

func (e BaseEvent) Beat() float64 {
	return e.beat
}

func (e *BaseEvent) SetBeat(b float64) {
	e.beat = b
}

func (e BaseEvent) Type() beatsaber.EventType {
	return e.typ
}

func (e *BaseEvent) SetType(t beatsaber.EventType) {
	e.typ = t
}

func (e BaseEvent) Value() beatsaber.EventValue {
	return e.val
}

func (e *BaseEvent) SetValue(v beatsaber.EventValue) {
	e.val = v
}
