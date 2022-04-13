package evt

import (
	"encoding/json"

	"github.com/shasderias/ilysa/beatsaber"
)

// Base represents the attributes common to all Beat Saber events - beat, type
// and value. Beat here refers to the scaled beat (i.e. after BPM changes).
// Ilysa unscales the beat on save.
//
// Base is meant to be embedded and unless you are writing a low level extension,
// you do not want this.
type Base struct {
	beat     float64
	offset   float64
	typ      Type
	val      Value
	floatVal float64

	tags map[string]struct{}
}

func NewBase(opts ...Option) Base {
	e := Base{
		beat:     0,
		offset:   0,
		typ:      -1,
		val:      0,
		floatVal: 1.0,
		tags:     make(map[string]struct{}),
	}
	be := baseWrapper{&e}
	for _, opt := range opts {
		opt.Apply(be)
	}
	return e
}

// baseWrapper exists to allow options to be applied to a Base event as Base
// being designed to be implemented, cannot implement the Apply method.
type baseWrapper struct {
	*Base
}

func (w baseWrapper) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(w)
	}
}

func (e *Base) Beat() float64           { return e.beat }
func (e *Base) SetBeat(b float64)       { e.beat = b }
func (e *Base) Offset() float64         { return e.offset }
func (e *Base) SetOffset(o float64)     { e.offset = o }
func (e *Base) Type() Type              { return e.typ }
func (e *Base) SetType(t Type)          { e.typ = t }
func (e *Base) Value() Value            { return e.val }
func (e *Base) SetValue(v Value)        { e.val = v }
func (e *Base) IntValue() int           { return int(e.val) }
func (e *Base) SetIntValue(v int)       { e.val = Value(v) }
func (e *Base) FloatValue() float64     { return e.floatVal }
func (e *Base) SetFloatValue(v float64) { e.floatVal = v }

// SetTag labels the event with a string. Tags can be used to distinguish
// one event from another. This can be useful when constructing filters
// for FilterEvents. Tags are ignored when saving events.
func (e *Base) SetTag(tags ...string) {
	for _, t := range tags {
		e.tags[t] = struct{}{}
	}
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

func (e *Base) Copy() Base {
	return Base{
		beat:     e.beat,
		offset:   e.offset,
		typ:      e.typ,
		val:      e.val,
		floatVal: e.floatVal,
		tags:     make(map[string]struct{}),
	}
}

func (e Base) EventV220() beatsaber.EventV220 {
	return beatsaber.NewEventV220(e.beat, int(e.typ), int(e.val), nil)
}

func (e Base) EventV250() beatsaber.EventV250 {
	return beatsaber.NewEventV250(e.beat, int(e.typ), int(e.val), e.floatVal, nil)
}

func (e Base) EventV220WithCD(cd json.RawMessage) beatsaber.EventV220 {
	return beatsaber.NewEventV220(e.beat, int(e.typ), int(e.val), cd)
}

func (e Base) EventV250WithCD(cd json.RawMessage) beatsaber.EventV250 {
	return beatsaber.NewEventV250(e.beat, int(e.typ), int(e.val), e.floatVal, cd)
}
