package ilysa

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
)

type BasicLightingEvent struct {
	BaseEvent
}
type BasicLightingEventOpt interface {
	applyBasicLightingEvent(event *BasicLightingEvent)
}

func (c baseContext) NewLightingEvent(opts ...BasicLightingEventOpt) *BasicLightingEvent {
	e := &BasicLightingEvent{
		BaseEvent: BaseEvent{
			beat: c.B(),
			typ:  beatsaber.EventTypeBackLasers,
			val:  0,
		},
	}
	for _, opt := range opts {
		opt.applyBasicLightingEvent(e)
	}
	c.addEvent(e)
	return e
}

func (e *BasicLightingEvent) CustomData() (json.RawMessage, error) { return nil, nil }
