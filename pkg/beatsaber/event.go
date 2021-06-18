package beatsaber

import (
	"encoding/json"

	"github.com/shasderias/ilysa/pkg/swallowjson"
)

type Event struct {
	Time  Time       `json:"_time"`
	Type  EventType  `json:"_type"`
	Value EventValue `json:"_value"`

	CustomData json.RawMessage `json:"_customData,omitempty"`

	Extra map[string]*json.RawMessage `json:"-"`
}

func (e *Event) UnmarshalJSON(raw []byte) error {
	return swallowjson.UnmarshalWith(e, "Extra", raw)
}

func (e Event) MarshalJSON() ([]byte, error) {
	mj, err := swallowjson.MarshalWith(e, "Extra")
	return mj, err
}

type CustomData interface {
	json.Marshaler
}

type EventValue float64

const (
	EventValueLightOff       EventValue = 0
	EventValueLightBlueOn    EventValue = 1
	EventValueLightBlueFlash EventValue = 2
	EventValueLightBlueFade  EventValue = 3
	EventValueLightUnused4   EventValue = 4
	EventValueLightRedOn     EventValue = 5
	EventValueLightRedFlash  EventValue = 6
	EventValueLightRedFade   EventValue = 7
)

type EventValueSet []EventValue

func NewEventValueSet(eventValues ...EventValue) EventValueSet {
	return append(EventValueSet{}, eventValues...)
}

func (s EventValueSet) Pick(n int) EventValue {
	return s[n%len(s)]
}

func IsLightingEvent(e EventType) bool {
	switch e {
	case EventTypeBackLasers:
		fallthrough
	case EventTypeRingLights:
		fallthrough
	case EventTypeLeftRotatingLasers:
		fallthrough
	case EventTypeRightRotatingLasers:
		fallthrough
	case EventTypeCenterLights:
		return true
	}
	return false
}
