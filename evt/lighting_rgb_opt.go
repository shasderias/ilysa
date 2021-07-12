package evt

import (
	"image/color"

	"github.com/shasderias/ilysa/lightid"
)

type withLightIDOpt struct {
	l lightid.ID
}

func WithLightID(id lightid.ID) withLightIDOpt {
	return withLightIDOpt{id}
}

func (o withLightIDOpt) applyRGBLighting(e *RGBLighting) {
	e.SetLightID(o.l)
}

type withAlphaOpt struct {
	a float64
}

func WithAlpha(a float64) withAlphaOpt {
	return withAlphaOpt{a}
}

func (o withAlphaOpt) apply(e Event) {
	switch te := e.(type) {
	case *RGBLighting:
		te.SetAlpha(o.a)
	}
}

func (o withAlphaOpt) applyRGBLighting(e *RGBLighting) {
	e.SetAlpha(o.a)
}

type withColorOpt struct {
	c color.Color
}

func WithColor(c color.Color) withColorOpt {
	return withColorOpt{c}
}

func (o withColorOpt) applyRGBLighting(e *RGBLighting) {
	e.SetColor(o.c)
}
