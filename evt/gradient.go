package evt

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

type ChromaGradient struct {
	Base
	chroma.Gradient
}

// NewChromaGradient creates a new Chroma 2.0 gradient.
func NewChromaGradient(opts ...Option) *ChromaGradient {
	e := &ChromaGradient{Base: NewBase()}
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

func (e *ChromaGradient) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

func (e ChromaGradient) EventV220() beatsaber.EventV220 {
	cd, err := e.Gradient.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV220WithCD(cd)
}

func (e ChromaGradient) EventV250() beatsaber.EventV250 {
	cd, err := e.Gradient.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV250WithCD(cd)
}
