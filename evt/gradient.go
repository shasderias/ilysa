package evt

import (
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
