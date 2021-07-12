package evt

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

func NewZoom(opts ...ZoomOpt) Zoom {
	e := Zoom{NewBase(
		WithType(beatsaber.EventTypeRingZoom),
	)}
	for _, opt := range opts {
		opt.applyZoom(&e)
	}
	return e
}

type Zoom struct {
	Base
}

type ZoomOpt interface {
	applyZoom(*Zoom)
}

func (e Zoom) CustomData() (json.RawMessage, error) { return nil, nil }

func (e *Zoom) Apply(opts ...ZoomOpt) {
	for _, opt := range opts {
		opt.applyZoom(e)
	}
}

func NewPreciseZoom(opts ...PreciseZoomOpt) PreciseZoom {
	e := PreciseZoom{Base: NewBase(
		WithType(beatsaber.EventTypeRingZoom),
	)}
	for _, opt := range opts {
		opt.applyPreciseZoom(&e)
	}
	return e
}

type PreciseZoom struct {
	Base
	chroma.PreciseZoom
}

type PreciseZoomOpt interface {
	applyPreciseZoom(*PreciseZoom)
}

func (e *PreciseZoom) Apply(opts ...PreciseZoomOpt) {
	for _, opt := range opts {
		opt.applyPreciseZoom(e)
	}
}
