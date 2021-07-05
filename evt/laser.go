package evt

import (
	"encoding/json"

	"github.com/shasderias/ilysa/chroma"
)

func NewLaser(opts ...LaserOpt) Laser {
	e := Laser{NewBase(WithInvalidDefaults())}
	for _, opt := range opts {
		opt.applyLaser(&e)
	}
	return e
}

type Laser struct {
	Base
}

type LaserOpt interface {
	applyLaser(*Laser)
}

func (e *Laser) Apply(opts ...LaserOpt) {
	for _, opt := range opts {
		opt.applyLaser(e)
	}
}

func (e Laser) CustomData() (json.RawMessage, error) { return nil, nil }

func NewPreciseLaser(opts ...PreciseLaserOpt) PreciseLaser {
	e := PreciseLaser{Base: NewBase(WithInvalidDefaults())}
	for _, opt := range opts {
		opt.applyPreciseLaser(&e)
	}
	return e
}

type PreciseLaser struct {
	Base
	chroma.PreciseLaser
}

type PreciseLaserOpt interface {
	applyPreciseLaser(*PreciseLaser)
}

func (e *PreciseLaser) Apply(opts ...PreciseLaserOpt) {
	for _, opt := range opts {
		opt.applyPreciseLaser(e)
	}
}
