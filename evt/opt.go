package evt

import (
	"image/color"

	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/lightid"
)

type optType struct {
	t Type
}

func (o optType) Apply(evt Event) {
	evt.SetType(o.t)
}

func OptType(t Type) Option {
	return optType{t}
}

type optValue struct {
	v Value
}

func (o optValue) Apply(evt Event) {
	evt.SetValue(o.v)
}

func OptValue(v Value) Option {
	return optValue{v}
}

type optIntValue struct {
	v int
}

func (o optIntValue) Apply(evt Event) { evt.SetValue(Value(o.v)) }
func OptIntValue(v int) Option        { return optIntValue{v} }

type colorer interface {
	Color() color.Color
	SetColor(color.Color)
}

func OptColor(c colorful.Color) Option { return optColor{c} }

type optColor struct {
	colorful.Color
}

func (o optColor) Apply(evt Event) {
	e, ok := evt.(colorer)
	if !ok {
		return
	}
	e.SetColor(o.Color)
}

type optLightID struct {
	lightID lightid.ID
}

func (o optLightID) Apply(evt Event) {
	e, ok := evt.(lightIDer)
	if !ok {
		return
	}
	e.SetLightID(o.lightID)
}

func OptLightID(id lightid.ID) Option {
	return optLightID{id}
}

type lightIDer interface {
	LightID() lightid.ID
	SetLightID(id lightid.ID)
}

type optShiftB struct {
	b float64
}

func OptShiftB(b float64) Option {
	return optShiftB{b}
}

func (o optShiftB) Apply(e Event) {
	e.SetBeat(e.Beat() + o.b)
}

type alphaer interface {
	Alpha() float64
	SetAlpha(float64)
}

type optAlpha struct {
	a float64
}

func OptAlpha(a float64) Option {
	return optAlpha{a}
}

func (o optAlpha) Apply(e Event) {
	ae, ok := e.(alphaer)
	if !ok {
		return
	}
	ae.SetAlpha(o.a)
}

type optFloatValue struct {
	f float64
}

func OptFloatValue(f float64) Option {
	return optFloatValue{f}
}

func (o optFloatValue) Apply(e Event) {
	e.SetFloatValue(o.f)
}
