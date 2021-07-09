package evt

import (
	"fmt"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

type DirectionalLaser int

const (
	LeftLaser  DirectionalLaser = 0
	RightLaser DirectionalLaser = 1
)

// WithLockPosition will not reset laser positions when true
func WithLockPosition(b bool) withLockPositionOpt {
	return withLockPositionOpt{b}
}

type withLockPositionOpt struct {
	b bool
}

func (o withLockPositionOpt) apply(e Event) {
	plse, ok := e.(*PreciseLaser)
	if !ok {
		return
	}
	plse.LockPosition = o.b
}

func (o withLockPositionOpt) applyPreciseLaser(e *PreciseLaser) {
	e.LockPosition = o.b
}

func WithLaserSpeed(v int) withLaserSpeedOpt {
	return withLaserSpeedOpt{v}
}

type withLaserSpeedOpt struct {
	v int
}

func (o withLaserSpeedOpt) apply(e Event) {
	switch lse := e.(type) {
	case *Laser:
		lse.SetValue(beatsaber.EventValue(o.v))
	case *PreciseLaser:
		lse.SetValue(beatsaber.EventValue(o.v))
	}
}

func (o withLaserSpeedOpt) applyLaser(e *Laser) {
	e.SetValue(beatsaber.EventValue(o.v))
}

func (o withLaserSpeedOpt) applyPreciseLaser(e *PreciseLaser) {
	e.SetValue(beatsaber.EventValue(o.v))
}

// WithPreciseLaserSpeed is identical to just setting value, but allows for decimals.
// Will overwrite value (Because the game will randomize laser position on
// anything other than value 0, a small trick you can do is set value to 1 and
// _preciseSpeed to 0, creating 0 s lasers with a randomized position).
func WithPreciseLaserSpeed(s float64) withPreciseLaserSpeedOpt {
	return withPreciseLaserSpeedOpt{s}
}

type withPreciseLaserSpeedOpt struct {
	s float64
}

func (o withPreciseLaserSpeedOpt) apply(e Event) {
	lse, ok := e.(*PreciseLaser)
	if !ok {
		return
	}
	lse.Speed = o.s
}

func (o withPreciseLaserSpeedOpt) applyPreciseLaser(e *PreciseLaser) {
	e.Speed = o.s
}

// WithLaserDirection set the spin direction
func WithLaserDirection(d chroma.SpinDirection) withDirectionOpt {
	return withDirectionOpt{d}
}

type withDirectionOpt struct {
	direction chroma.SpinDirection
}

func (o withDirectionOpt) applyPreciseLaser(e *PreciseLaser) {
	e.Direction = o.direction
}

func WithDirectionalLaser(dl DirectionalLaser) withDirectionalLaserOpt {
	switch dl {
	case LeftLaser:
		return withDirectionalLaserOpt{beatsaber.EventTypeLeftRotatingLasersRotationSpeed}
	case RightLaser:
		return withDirectionalLaserOpt{beatsaber.EventTypeRightRotatingLasersRotationSpeed}
	default:
		panic(fmt.Sprintf("WithDirectionalLaser: unsupported direction %v", dl))
	}
}

type withDirectionalLaserOpt struct {
	t beatsaber.EventType
}

func (o withDirectionalLaserOpt) applyLaser(e *Laser) {
	e.SetType(o.t)
}

func (o withDirectionalLaserOpt) applyPreciseLaser(e *PreciseLaser) {
	e.SetType(o.t)
}
