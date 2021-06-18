package ilysa

import (
	"encoding/json"

	"github.com/shasderias/ilysa/pkg/beatsaber"
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
			Beat:  c.B(),
			Type:  beatsaber.EventTypeBackLasers,
			Value: 0,
		},
	}
	for _, opt := range opts {
		opt.applyBasicLightingEvent(e)
	}
	c.addEvent(e)
	return e
}

func (e *BasicLightingEvent) CustomData() (json.RawMessage, error) { return nil, nil }
