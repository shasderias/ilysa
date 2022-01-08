package evt

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

// NewRingRotation returns a new base game ring rotation event.
func NewRingRotation(opts ...Option) *Rotation {
	e := &Rotation{NewBase()}
	e.SetType(TypeRingRotation)
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

type Rotation struct {
	Base
}

func (e Rotation) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(&e)
	}
}

// NewChromaRingRotation returns a new Chroma ring rotation event.
func NewChromaRingRotation(opts ...Option) *ChromaRingRotation {
	e := &ChromaRingRotation{Base: NewBase()}
	e.SetType(TypeRingRotation)
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

type ChromaRingRotation struct {
	Base
	chroma.RingRotation
}

func (e *ChromaRingRotation) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

func (e ChromaRingRotation) EventV220() beatsaber.EventV220 {
	cd, err := e.RingRotation.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV220WithCD(cd)
}

func (e ChromaRingRotation) EventV250() beatsaber.EventV250 {
	cd, err := e.RingRotation.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV250WithCD(cd)
}

func (e *ChromaRingRotation) ChromaRingRotationNameFilter() string {
	return e.RingRotation.NameFilter.Value
}
func (e *ChromaRingRotation) SetChromaRingRotationNameFilter(f string) {
	e.RingRotation.NameFilter.Set(f)
}
func OptChromaRingRotationNameFilter(f string) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotationNameFilter(f)
	})

}

func (e *ChromaRingRotation) ChromaRingRotationReset() bool {
	return e.RingRotation.Reset.Value
}
func (e *ChromaRingRotation) SetChromaRingRotationReset(b bool) {
	e.RingRotation.Reset.Set(b)
}
func OptChromaRingRotationReset(b bool) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotationReset(b)
	})
}

func (e *ChromaRingRotation) ChromaRingRotation() float64 {
	return e.RingRotation.Rotation.Value
}
func (e *ChromaRingRotation) SetChromaRingRotation(r float64) {
	e.RingRotation.Rotation.Set(r)
}
func OptChromaRingRotation(r float64) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotation(r)
	})
}

func (e *ChromaRingRotation) ChromaRingRotationStep() float64 {
	return e.RingRotation.Step.Value
}
func (e *ChromaRingRotation) SetChromaRingRotationStep(r float64) {
	e.RingRotation.Step.Set(r)
}
func OptChromaRingRotationStep(r float64) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotationStep(r)
	})
}

func (e *ChromaRingRotation) ChromaRingRotationProp() float64 {
	return e.RingRotation.Prop.Value
}
func (e *ChromaRingRotation) SetChromaRingRotationProp(p float64) {
	e.RingRotation.Prop.Set(p)
}
func OptChromaRingRotationProp(p float64) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotationProp(p)
	})
}

func (e *ChromaRingRotation) ChromaRingRotationSpeed() float64 {
	return e.RingRotation.Speed.Value
}
func (e *ChromaRingRotation) SetChromaRingRotationSpeed(s float64) {
	e.RingRotation.Speed.Set(s)
}
func OptChromaRingRotationSpeed(s float64) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotationSpeed(s)
	})
}

func (e *ChromaRingRotation) ChromaRingRotationDirection() chroma.SpinDirection {
	return e.RingRotation.Direction.Value
}
func (e *ChromaRingRotation) SetChromaRingRotationDirection(d chroma.SpinDirection) {
	e.RingRotation.Direction.Set(d)
}
func OptChromaRingRotationDirection(d chroma.SpinDirection) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotationDirection(d)
	})
}

func (e *ChromaRingRotation) ChromaRingRotationCounterSpin() bool {
	return e.RingRotation.CounterSpin.Value
}
func (e *ChromaRingRotation) SetChromaRingRotationCounterSpin(b bool) {
	e.RingRotation.CounterSpin.Set(b)
}
func OptChromaRingRotationCounterSpin(b bool) Option {
	return NewFuncOpt(func(e Event) {
		crre, ok := e.(*ChromaRingRotation)
		if !ok {
			return
		}
		crre.SetChromaRingRotationCounterSpin(b)
	})
}
