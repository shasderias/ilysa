package ilysa

import (
	"image/color"
)

type withLightIDOpt struct {
	lightID LightID
}

func WithLightID(lightID LightID) withLightIDOpt {
	return withLightIDOpt{lightID}
}

func (w withLightIDOpt) applyRGBLightingEvent(e *RGBLightingEvent) {
	e.SetLightID(w.lightID)
}

type withAlphaOpt struct {
	a float64
}

func WithAlpha(a float64) withAlphaOpt {
	return withAlphaOpt{a}
}

func (w withAlphaOpt) applyRGBLightingEvent(e *RGBLightingEvent) {
	e.SetAlpha(w.a)
}

func (w withAlphaOpt) applyCompoundRGBLightingEvent(e *CompoundRGBLightingEvent) {
	e.SetAlpha(w.a)
}

type withColorOpt struct {
	c color.Color
}

func WithColor(c color.Color) withColorOpt {
	return withColorOpt{c}
}

func (w withColorOpt) applyRGBLightingEvent(e *RGBLightingEvent) {
	e.SetColor(w.c)
}

func (w withColorOpt) applyCompoundRGBLightingEvent(e *CompoundRGBLightingEvent) {
	e.SetColor(w.c)
}
