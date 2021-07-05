package evt

import (
	"image/color"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/light"
)

type RGBLighting struct {
	Base
	chroma.RGB
}

type RGBLightingOpt interface {
	applyRGBLighting(*RGBLighting)
}

func NewRGBLighting(opts ...RGBLightingOpt) RGBLighting {
	e := RGBLighting{Base: NewBase(WithRGBLightingDefaults())}
	for _, opt := range opts {
		opt.applyRGBLighting(&e)
	}
	return e
}

func (e *RGBLighting) SetLight(lightType LightType) *RGBLighting {
	e.SetType(beatsaber.EventType(lightType))
	return e
}

func (e *RGBLighting) Color() color.Color {
	return e.RGB.Color
}

func (e *RGBLighting) SetColor(c color.Color) {
	e.RGB.Color = c
}

func (e *RGBLighting) Alpha() float64 {
	c := colorful.FromColor(e.Color())
	return c.A
}

func (e *RGBLighting) SetAlpha(a float64) {
	c := colorful.FromColor(e.RGB.Color)
	c.A = a
	e.SetColor(c)
}

func (e *RGBLighting) SetSingleLightID(id int) {
	e.SetLightID(light.NewID(id))
}

func (e *RGBLighting) SetLightID(id light.ID) {
	e.RGB.LightID = chroma.LightID(id)
}

func (e *RGBLighting) Apply(opts ...RGBLightingOpt) {
	for _, opt := range opts {
		opt.applyRGBLighting(e)
	}
}
