package evt

import (
	"github.com/shasderias/ilysa/beatsaber"
)

type Base struct {
	beat float64
	typ  beatsaber.EventType
	val  beatsaber.EventValue
}

func NewBase(opts ...BaseOpt) Base {
	e := Base{
		beat: -1,
		typ:  beatsaber.EventTypeInvalid,
		val:  beatsaber.EventValueInvalid,
	}
	for _, opt := range opts {
		opt.applyBase(&e)
	}
	return e
}

func (e Base) ShiftBeat(offset float64) {
	e.beat += offset
}

func (e Base) Beat() float64 {
	return e.beat
}

func (e *Base) SetBeat(b float64) {
	e.beat = b
}

func (e Base) Type() beatsaber.EventType {
	return e.typ
}

func (e *Base) SetType(t beatsaber.EventType) {
	e.typ = t
}

func (e Base) Value() beatsaber.EventValue {
	return e.val
}

func (e *Base) SetValue(v beatsaber.EventValue) {
	e.val = v
}
