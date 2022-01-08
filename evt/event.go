// Package evt is a package for creating Beat Saber events.
package evt

import "encoding/json"

type Option interface {
	Apply(evt Event)
}

type Event interface {
	Beat() float64
	SetBeat(b float64)
	Type() Type
	SetType(t Type)
	Value() Value
	SetValue(v Value)
	FloatValue() float64
	SetFloatValue(v float64)
	SetTag(tags ...string)
	HasTag(tag ...string) bool

	Apply(opts ...Option)
}

type CustomDataer interface {
	CustomData() (json.RawMessage, error)
}

type (
	Type  int
	Value int
)

const (
	TypeInvalid         Type = -1
	TypeBackLasers           = 0
	TypeRingLights           = 1
	TypeLeftLasers           = 2
	TypeRightLasers          = 3
	TypeCenterLights         = 4
	TypeBoostLights          = 5
	TypeRingRotation         = 8
	TypeRingZoom             = 9
	TypeBPMChange            = 10
	TypeLeftLaserSpeed       = 12
	TypeRightLaserSpeed      = 13
	TypeEarlyRotation        = 14
	TypeLateRotation         = 15
)

const (
	ValueInvalid        Value = -1
	ValueLightOff             = 0
	ValueLightBlueOn          = 1
	ValueLightBlueFlash       = 2
	ValueLightBlueFade        = 3
	ValueLightUnused4         = 4
	ValueLightRedOn           = 5
	ValueLightRedFlash        = 6
	ValueLightRedFade         = 7
)

const (
	ValueBoostOff Value = 0
	ValueBoostOn        = 1
)

const (
	Value60CCW Value = iota
	Value45CCW
	Value30CCW
	Value15CCW
	Value15CW
	Value30CW
	Value45CW
	Value60CW
)

type Custom struct {
	Base
}

func (e *Custom) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

func New(opts ...Option) Event {
	ce := &Custom{
		Base: Base{
			beat:     0,
			offset:   0,
			typ:      -1,
			val:      0,
			floatVal: 1.0,
			tags:     make(map[string]struct{}),
		},
	}

	for _, opt := range opts {
		opt.Apply(ce)
	}

	return ce
}

type Events []Event

func NewEvents(events ...Event) Events {
	return events
}

func (e *Events) Add(newEvents ...Event) {
	*e = append(*e, newEvents...)
}

func (e *Events) Apply(opts ...Option) {
	for _, evt := range *e {
		evt.Apply(opts...)
	}
}

func (e *Events) Filter(keepFunc func(e Event) bool) {
	keepEvents := NewEvents()
	for _, evt := range *e {
		if keepFunc(evt) {
			keepEvents = append(keepEvents, evt)
		}
	}
	*e = keepEvents
}
