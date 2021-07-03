package ilysa

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
)

type RotationEvent struct {
	BaseEvent
}

type RotationEventOpt interface {
	applyRotationEvent(*RotationEvent)
}

func (c baseContext) NewRotationEvent(opts ...RotationEventOpt) *RotationEvent {
	e := &RotationEvent{BaseEvent: BaseEvent{
		beat: c.B(),
		typ:  beatsaber.EventTypeRingSpin,
		val:  0,
	}}

	e.Mod(opts...)
	c.addEvent(e)
	return e
}

func (e RotationEvent) CustomData() (json.RawMessage, error) { return nil, nil }

func (e RotationEvent) Mod(opts ...RotationEventOpt) {
	for _, opt := range opts {
		opt.applyRotationEvent(&e)
	}
}

type PreciseRotationEvent struct {
	BaseEvent
	chroma.PreciseRotation
}

type PreciseRotationEventOpt interface {
	applyPreciseRotationEvent(*PreciseRotationEvent)
}

func (c baseContext) NewPreciseRotationEvent(opts ...PreciseRotationEventOpt) *PreciseRotationEvent {
	e := &PreciseRotationEvent{
		BaseEvent: BaseEvent{
			beat: c.B(),
			typ:  beatsaber.EventTypeRingSpin,
			val:  0,
		},
	}

	e.Mod(opts...)
	c.addEvent(e)
	return e
}

func (e *PreciseRotationEvent) Mod(opts ...PreciseRotationEventOpt) {
	for _, opt := range opts {
		opt.applyPreciseRotationEvent(e)
	}
}

type withNameFilterOpt struct {
	nameFilter string
}

func WithNameFilter(nf string) withNameFilterOpt {
	return withNameFilterOpt{nf}
}

func (o withNameFilterOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.NameFilter = o.nameFilter
}

type withResetOpt struct {
	reset bool
}

func WithReset(r bool) withResetOpt {
	return withResetOpt{r}
}

func (o withResetOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.Reset = o.reset
}

type withRotationOpt struct {
	rotation float64
}

func WithRotation(r float64) withRotationOpt {
	return withRotationOpt{r}
}

func (o withRotationOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.Rotation = o.rotation
}

type withPropOpt struct {
	prop float64
}

func WithProp(p float64) withPropOpt {
	return withPropOpt{p}
}

func (o withPropOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.Prop = o.prop
}

type withCounterSpinOpt struct {
	counterSpin bool
}

func WithCounterSpin(c bool) withCounterSpinOpt {
	return withCounterSpinOpt{c}
}

func (o withCounterSpinOpt) applyPreciseRotationEvent(e *PreciseRotationEvent) {
	e.CounterSpin = o.counterSpin
}
