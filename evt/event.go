package evt

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
)

type Event interface {
	Beat() float64
	SetBeat(float64)
	Type() beatsaber.EventType
	SetType(beatsaber.EventType)
	Value() beatsaber.EventValue
	SetValue(value beatsaber.EventValue)

	CustomData() (json.RawMessage, error)
}
