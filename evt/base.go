package evt

import (
	"github.com/shasderias/ilysa/beatsaber"
)

type Base struct {
	beat float64
	typ  beatsaber.EventType
	val  beatsaber.EventValue

	tags map[string]struct{}
}

func NewBase(opts ...BaseOpt) Base {
	e := Base{
		beat: -1,
		typ:  beatsaber.EventTypeInvalid,
		val:  beatsaber.EventValueInvalid,
		tags: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt.applyBase(&e)
	}
	return e
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

func (e *Base) SetTag(tag string) {
	e.tags[tag] = struct{}{}
}

// HasTag returns true if the event is tagged with any one of tag.
func (e *Base) HasTag(tag ...string) bool {
	for _, t := range tag {
		if _, ok := e.tags[t]; ok {
			return true
		}
	}
	return false
}

func (e *Base) ClearTag(tag string) {
	delete(e.tags, tag)
}
