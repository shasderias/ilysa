package evt

import (
	"image/color"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/lightid"
)

// NewLighting returns a new base game lighting event.
func NewLighting(opts ...Option) *Lighting {
	e := &Lighting{NewBase()}
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

type Lighting struct {
	Base
}

func (e *Lighting) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}

type ChromaLighting struct {
	Base
	chroma.Lighting
}

// NewChromaLighting returns a new Chroma lighting event.
func NewChromaLighting(opts ...Option) *ChromaLighting {
	e := &ChromaLighting{Base: NewBase()}
	e.SetValue(ValueLightRedOn)
	for _, opt := range opts {
		opt.Apply(e)
	}
	return e
}

func (e *ChromaLighting) Apply(opts ...Option) {
	for _, opt := range opts {
		opt.Apply(e)
	}
}
func (e ChromaLighting) EventV220() beatsaber.EventV220 {
	cd, err := e.Lighting.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV220WithCD(cd)
}

func (e ChromaLighting) EventV250() beatsaber.EventV250 {
	cd, err := e.Lighting.CustomData()
	if err != nil {
		panic(err)
	}
	return e.Base.EventV250WithCD(cd)
}

func (e *ChromaLighting) Color() color.Color {
	return e.Lighting.Color
}
func (e *ChromaLighting) SetColor(c color.Color) {
	e.Lighting.Color = c
}
func OptColor(c color.Color) Option {
	return NewFuncOpt(func(e Event) {
		cl, ok := e.(*ChromaLighting)
		if !ok {
			return
		}
		cl.SetColor(c)
	})
}

type Alphaer interface {
	Alpha() float64
	SetAlpha(float64)
}

func (e *ChromaLighting) Alpha() float64 {
	c := colorful.FromColor(e.Color())
	return c.A
}
func (e *ChromaLighting) SetAlpha(a float64) {
	c := colorful.FromColor(e.Lighting.Color)
	c.A = a
	e.SetColor(c)
}
func OptAlpha(a float64) Option {
	return NewFuncOpt(func(e Event) {
		ae, ok := e.(Alphaer)
		if !ok {
			return
		}
		ae.SetAlpha(a)
	})
}

type LightIDer interface {
	LightID() lightid.ID
	SetLightID(id lightid.ID)
}

func (e *ChromaLighting) SetSingleLightID(id int)  { e.SetLightID(lightid.New(id)) }
func (e *ChromaLighting) LightID() lightid.ID      { return lightid.ID(e.Lighting.LightID) }
func (e *ChromaLighting) SetLightID(id lightid.ID) { e.Lighting.LightID = chroma.LightID(id) }
func OptLightID(id lightid.ID) Option {
	return NewFuncOpt(func(e Event) {
		lid, ok := e.(LightIDer)
		if !ok {
			return
		}
		lid.SetLightID(id)
	})
}
