package event

import "ilysa/pkg/beatsaber"

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
