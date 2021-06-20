package ilysa

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

type ZoomEventOpt interface {
	applyZoomEvent(*ZoomEvent)
}

type ZoomEvent struct {
	BaseEvent
}

func (c baseContext) NewZoomEvent(opts ...ZoomEventOpt) *ZoomEvent {
	e := &ZoomEvent{
		BaseEvent{
			Beat:  c.B(),
			Type:  beatsaber.EventTypeRingZoom,
			Value: 0,
		},
	}
	e.Mod(opts...)
	c.addEvent(e)
	return e
}

func (e ZoomEvent) CustomData() (json.RawMessage, error) { return nil, nil }

func (e *ZoomEvent) Mod(opts ...ZoomEventOpt) {
	for _, opt := range opts {
		opt.applyZoomEvent(e)
	}
}

type PreciseZoomEventOpt interface {
	applyPreciseZoomEvent(*PreciseZoomEvent)
}

type PreciseZoomEvent struct {
	BaseEvent
	chroma.PreciseZoom
}

func (c baseContext) NewPreciseZoomEvent(opts ...PreciseZoomEventOpt) *PreciseZoomEvent {
	e := &PreciseZoomEvent{
		BaseEvent: BaseEvent{
			Beat:  c.B(),
			Type:  beatsaber.EventTypeRingZoom,
			Value: 0,
		},
	}
	e.Mod(opts...)
	c.addEvent(e)
	return e
}

func (e *PreciseZoomEvent) Mod(opts ...PreciseZoomEventOpt) {
	for _, opt := range opts {
		opt.applyPreciseZoomEvent(e)
	}
}
