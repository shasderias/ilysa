package fx

import (
	ilysa2 "github.com/shasderias/ilysa"
	ease2 "github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

type RippleOpt interface {
	applyRipple(ilysa2.TimeLightContext, *ilysa2.CompoundRGBLightingEvent)
}

func Ripple(ctx ilysa2.TimeLightContext, events *ilysa2.CompoundRGBLightingEvent, step float64, opts ...RippleOpt) {
	for _, e := range *events {
		e.ShiftBeat(ctx.LightIDT() * step)
	}

	for _, opt := range opts {
		opt.applyRipple(ctx, events)
	}
}

type withFadeInOpt struct {
	startT     float64
	endT       float64
	startAlpha float64
	endAlpha float64
	easeFn   ease2.Func
}

func WithAlphaBlend(startT, endT, startAlpha, endAlpha float64, easeFn ease2.Func) withFadeInOpt {
	return withFadeInOpt{
		startT:     startT,
		endT:       endT,
		startAlpha: startAlpha,
		endAlpha:   endAlpha,
		easeFn:     easeFn,
	}
}

func (o withFadeInOpt) apply(ctx ilysa2.TimeLightContext, events *ilysa2.CompoundRGBLightingEvent) {
	if ctx.T() < o.startT || ctx.T() > o.endT {
		return
	}
	alphaScale := scale.Clamped(o.startT, o.endT, o.startAlpha, o.endAlpha)
	events.Mod(ilysa2.WithAlpha(events.GetAlpha() * o.easeFn(alphaScale(ctx.T()))))
}

func (o withFadeInOpt) applyRipple(ctx ilysa2.TimeLightContext, events *ilysa2.CompoundRGBLightingEvent) {
	o.apply(ctx, events)
}

func (o withFadeInOpt) applyColorSweep(ctx ilysa2.TimeLightContext, events *ilysa2.CompoundRGBLightingEvent) {
	o.apply(ctx, events)
}

func AlphaBlend(ctx ilysa2.TimeLightContext, events *ilysa2.CompoundRGBLightingEvent, startT, endT, startAlpha, endAlpha float64, easeFn ease2.Func) {
	if ctx.T() < startT || ctx.T() > endT {
		return
	}
	alphaScale := scale.Clamped(startT, endT, startAlpha, endAlpha)
	events.Mod(ilysa2.WithAlpha(events.GetAlpha() * easeFn(alphaScale(ctx.T()))))
}
