package evt

import (
	"github.com/shasderias/ilysa/chroma"
)

// NewLaserSpeed returns a laser speed event.
func NewLaserSpeed(opts ...Option) *LaserSpeed {
	e := &LaserSpeed{NewBase()}
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

type LaserSpeed struct {
	Base
}

func (e *LaserSpeed) SetLaserSpeed(speed int) { e.SetIntValue(speed) }
func (e *LaserSpeed) LaserSpeed() int         { return e.IntValue() }

func (e *LaserSpeed) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

// NewChromaLaserSpeed returns a Chroma precise laser event.
func NewChromaLaserSpeed(opts ...Option) *ChromaLaserSpeed {
	e := &ChromaLaserSpeed{Base: NewBase()}
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

type ChromaLaserSpeed struct {
	Base
	chroma.LaserSpeed
}

func (e *ChromaLaserSpeed) ChromaLaserSpeedLockPosition() bool     { return e.LaserSpeed.LockPosition }
func (e *ChromaLaserSpeed) SetChromaLaserSpeedLockPosition(l bool) { e.LaserSpeed.LockPosition = l }
func (e *ChromaLaserSpeed) ChromaLaserSpeed() float64              { return e.LaserSpeed.Speed }
func (e *ChromaLaserSpeed) SetChromaLaserSpeed(speed float64)      { e.LaserSpeed.Speed = speed }
func (e *ChromaLaserSpeed) ChromaLaserSpeedSpinDirection() chroma.SpinDirection {
	return e.LaserSpeed.Direction
}
func (e *ChromaLaserSpeed) SetChromaLaserSpeedSpinDirection(dir chroma.SpinDirection) {
	e.LaserSpeed.Direction = dir
}

func (e *ChromaLaserSpeed) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

type optChromaLaserSpeed struct {
	s float64
}

func (o optChromaLaserSpeed) Apply(evt Event) {
	clse, ok := evt.(*ChromaLaserSpeed)
	if !ok {
		return
	}
	clse.SetChromaLaserSpeed(o.s)
}

func OptChromaLaserSpeed(s float64) Option {
	return &optChromaLaserSpeed{s}
}

type optChromaLaserSpeedLockPosition struct {
	l bool
}

func (o optChromaLaserSpeedLockPosition) Apply(evt Event) {
	clse, ok := evt.(*ChromaLaserSpeed)
	if !ok {
		return
	}
	clse.SetChromaLaserSpeedLockPosition(o.l)
}

func OptChromaLaserSpeedLockPosition(l bool) Option {
	return &optChromaLaserSpeedLockPosition{l}
}
