package evt

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

func NewRotation(opts ...RotationOpt) Rotation {
	e := Rotation{NewBase(
		WithType(beatsaber.EventTypeRingSpin),
	)}
	for _, opt := range opts {
		opt.applyRotation(&e)
	}
	return e
}

type Rotation struct {
	Base
}

type RotationOpt interface {
	applyRotation(*Rotation)
}

func (e Rotation) CustomData() (json.RawMessage, error) { return nil, nil }

func (e Rotation) Apply(opts ...RotationOpt) {
	for _, opt := range opts {
		opt.applyRotation(&e)
	}
}

func NewPreciseRotation(opts ...PreciseRotationOpt) PreciseRotation {
	e := PreciseRotation{
		Base: NewBase(
			WithType(beatsaber.EventTypeRingSpin),
		)}
	for _, opt := range opts {
		opt.applyPreciseRotation(&e)
	}
	return e
}

type PreciseRotation struct {
	Base
	chroma.PreciseRotation
}

type PreciseRotationOpt interface {
	applyPreciseRotation(*PreciseRotation)
}

func (e *PreciseRotation) Apply(opts ...PreciseRotationOpt) {
	for _, opt := range opts {
		opt.applyPreciseRotation(e)
	}
}
