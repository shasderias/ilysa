package evt

import (
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

func (e *ChromaRingRotation) ChromaRingRotationNameFilter() string {
	return e.RingRotation.NameFilter
}
func (e *ChromaRingRotation) SetChromaRingRotationNameFilter(f string) {
	e.RingRotation.NameFilter = f
}
func (e *ChromaRingRotation) ChromaRingRotationReset() bool {
	return e.RingRotation.Reset
}
func (e *ChromaRingRotation) SetChromaRingRotationReset(b bool) {
	e.RingRotation.Reset = b
}
func (e *ChromaRingRotation) ChromaRingRotation() float64 {
	return e.RingRotation.Rotation
}
func (e *ChromaRingRotation) SetChromaRingRotation(r float64) {
	e.RingRotation.Rotation = r
}
func (e *ChromaRingRotation) ChromaRingRotationStep() float64 {
	return e.RingRotation.Step
}
func (e *ChromaRingRotation) SetChromaRingRotationStep(r float64) {
	e.RingRotation.Step = r
}
func (e *ChromaRingRotation) ChromaRingRotationProp() float64 {
	return e.RingRotation.Prop
}
func (e *ChromaRingRotation) SetChromaRingRotationProp(p float64) {
	e.RingRotation.Prop = p
}
func (e *ChromaRingRotation) ChromaRingRotationSpeed() float64 {
	return e.RingRotation.Speed
}
func (e *ChromaRingRotation) SetChromaRingRotationSpeed(s float64) {
	e.RingRotation.Speed = s
}
func (e *ChromaRingRotation) ChromaRingRotationDirection() chroma.SpinDirection {
	return e.RingRotation.Direction
}
func (e *ChromaRingRotation) SetChromaRingRotationDirection(d chroma.SpinDirection) {
	e.RingRotation.Direction = d
}
func (e *ChromaRingRotation) ChromaRingRotationCounterSpin() bool {
	return e.RingRotation.CounterSpin
}
func (e *ChromaRingRotation) SetChromaRingRotationCounterSpin(b bool) {
	e.RingRotation.CounterSpin = b
}
func (e *ChromaRingRotation) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}
