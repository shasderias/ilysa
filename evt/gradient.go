package evt

import "github.com/shasderias/ilysa/chroma"

type ChromaGradient struct {
	Base
	chroma.Gradient
}

type ChromaGradientOpt interface {
	applyChromaGradient(*ChromaGradient)
}

func NewChromaGradient(opts ...ChromaGradientOpt) ChromaGradient {
	e := ChromaGradient{Base: NewBase(WithRGBLightingDefaults())}
	for _, opt := range opts {
		opt.applyChromaGradient(&e)
	}
	return e
}

func (e *ChromaGradient) Apply(opts ...ChromaGradientOpt) {
	for _, opt := range opts {
		opt.applyChromaGradient(e)
	}
}
