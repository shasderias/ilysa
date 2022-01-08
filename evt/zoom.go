package evt

import (
	"github.com/shasderias/ilysa/chroma"
)

// NewZoom returns a new base game zoom event.
func NewZoom(opts ...Option) Zoom {
	e := Zoom{NewBase()}
	e.SetType(TypeRingZoom)
	for _, opt := range opts {
		opt.Apply(&e)
	}
	return e
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
func NewChromaZoom(opts ...Option) ChromaZoom {
	e := ChromaZoom{Base: NewBase()}
	e.SetType(TypeRingZoom)

	for _, opt := range opts {
		opt.Apply(&e)
	}
	return e
}

type ChromaZoom struct {
	Base
	chroma.Zoom
}

func (e *ChromaZoom) ChromaZoomStep() float64              { return e.Step }
func (e *ChromaZoom) SetChromaZoomStep(zoomStep float64)   { e.Step = zoomStep }
func (e *ChromaZoom) ChromaZoomSpeed() float64             { return e.Speed }
func (e *ChromaZoom) SetChromaZoomSpeed(zoomSpeed float64) { e.Speed = zoomSpeed }
func (e *ChromaZoom) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}
