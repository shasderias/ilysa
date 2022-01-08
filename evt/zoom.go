package evt

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

// NewZoom returns a new base game zoom event.
func NewZoom(opts ...Option) *Zoom {
	e := Zoom{NewBase()}
	e.SetType(TypeRingZoom)
	for _, opt := range opts {
		opt.Apply(&e)
	}
	return &e
}

type Zoom struct {
	Base
}

func (e *Zoom) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

// NewChromaZoom returns a Chroma precise zoom event.
func NewChromaZoom(opts ...Option) *ChromaZoom {
	e := ChromaZoom{Base: NewBase()}
	e.SetType(TypeRingZoom)

	for _, opt := range opts {
		opt.Apply(&e)
	}
	return &e
}

type ChromaZoom struct {
	Base
	chroma.Zoom
}

func (e *ChromaZoom) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

func (e ChromaZoom) EventV220() beatsaber.EventV220 {
	cd, err := e.Zoom.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV220WithCD(cd)
}

func (e ChromaZoom) EventV250() beatsaber.EventV250 {
	cd, err := e.Zoom.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV250WithCD(cd)
}

func (e *ChromaZoom) ChromaZoomStep() float64            { return e.Step.Value }
func (e *ChromaZoom) SetChromaZoomStep(zoomStep float64) { e.Step.Set(zoomStep) }
func OptChromaZoomStep(s float64) Option {
	return NewFuncOpt(func(e Event) {
		cz, ok := e.(*ChromaZoom)
		if !ok {
			return
		}
		cz.SetChromaZoomStep(s)
	})
}

func (e *ChromaZoom) ChromaZoomSpeed() float64             { return e.Speed.Value }
func (e *ChromaZoom) SetChromaZoomSpeed(zoomSpeed float64) { e.Speed.Set(zoomSpeed) }
func OptChromaZoomSpeed(s float64) Option {
	return NewFuncOpt(func(e Event) {
		cz, ok := e.(*ChromaZoom)
		if !ok {
			return
		}
		cz.SetChromaZoomSpeed(s)
	})
}
