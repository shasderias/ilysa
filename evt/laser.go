package evt

import (
	"github.com/shasderias/ilysa/beatsaber"
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

type LaserSpeeder interface {
	LaserSpeed() int
	SetLaserSpeed(int)
}

func OptLaserSpeed(speed int) Option { return OptIntValue(speed) }

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

func (e *ChromaLaserSpeed) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}
func (e ChromaLaserSpeed) EventV220() beatsaber.EventV220 {
	cd, err := e.LaserSpeed.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV220WithCD(cd)
}
func (e ChromaLaserSpeed) EventV250() beatsaber.EventV250 {
	cd, err := e.LaserSpeed.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV250WithCD(cd)
}

type ChromaLaserSpeeder interface {
	ChromaLaserSpeed() float64
	SetChromaLaserSpeed(float64)
}

func (e *ChromaLaserSpeed) ChromaLaserSpeed() float64 {
	return e.LaserSpeed.Speed.Value
}
func (e *ChromaLaserSpeed) SetChromaLaserSpeed(speed float64) {
	e.LaserSpeed.Speed.Set(speed)
}
func OptChromaLaserSpeed(s float64) Option {
	return NewFuncOpt(func(e Event) {
		clse, ok := e.(ChromaLaserSpeeder)
		if !ok {
			return
		}
		clse.SetChromaLaserSpeed(s)
	})
}

type ChromaLaserSpeedLockPositioner interface {
	ChromaLaserSpeedLockPosition() bool
	SetChromaLaserSpeedLockPosition(bool)
}

func (e *ChromaLaserSpeed) ChromaLaserSpeedLockPosition() bool {
	return e.LaserSpeed.LockPosition.Value
}
func (e *ChromaLaserSpeed) SetChromaLaserSpeedLockPosition(l bool) {
	e.LaserSpeed.LockPosition.Set(l)
}
func OptChromaLaserSpeedLockPosition(l bool) Option {
	return NewFuncOpt(func(e Event) {
		clse, ok := e.(ChromaLaserSpeedLockPositioner)
		if !ok {
			return
		}
		clse.SetChromaLaserSpeedLockPosition(l)
	})
}

type ChromaLaserSpeedSpinDirectioner interface {
	ChromaLaserSpeedSpinDirection() chroma.SpinDirection
	SetChromaLaserSpeedSpinDirection(chroma.SpinDirection)
}

func (e *ChromaLaserSpeed) ChromaLaserSpeedSpinDirection() chroma.SpinDirection {
	return e.LaserSpeed.Direction.Value
}
func (e *ChromaLaserSpeed) SetChromaLaserSpeedSpinDirection(dir chroma.SpinDirection) {
	e.LaserSpeed.Direction.Set(dir)
}
func OptChromaLaserSpeedSpinDirection(dir chroma.SpinDirection) Option {
	return NewFuncOpt(func(e Event) {
		clse, ok := e.(ChromaLaserSpeedSpinDirectioner)
		if !ok {
			return
		}
		clse.SetChromaLaserSpeedSpinDirection(dir)
	})
}
