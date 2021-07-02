package fx

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

type RippleOpt interface {
	applyRipple(ilysa.TimeLightContext, *ilysa.CompoundRGBLightingEvent)
}

func Ripple(ctx ilysa.TimeLightContext, events *ilysa.CompoundRGBLightingEvent, step float64, opts ...RippleOpt) {
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
	endAlpha   float64
	easeFn     ease.Func
}

func WithAlphaBlend(startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) withFadeInOpt {
	return withFadeInOpt{
		startT:     startT,
		endT:       endT,
		startAlpha: startAlpha,
		endAlpha:   endAlpha,
		easeFn:     easeFn,
	}
}

func (o withFadeInOpt) apply(ctx ilysa.TimeLightContext, events *ilysa.CompoundRGBLightingEvent) {
	if ctx.T() < o.startT || ctx.T() > o.endT {
		return
	}
	alphaScale := scale.Clamped(o.startT, o.endT, o.startAlpha, o.endAlpha)
	events.Mod(ilysa.WithAlpha(events.GetAlpha() * o.easeFn(alphaScale(ctx.T()))))
}

func (o withFadeInOpt) applyRipple(ctx ilysa.TimeLightContext, events *ilysa.CompoundRGBLightingEvent) {
	o.apply(ctx, events)
}

func (o withFadeInOpt) applyColorSweep(ctx ilysa.TimeLightContext, events *ilysa.CompoundRGBLightingEvent) {
	o.apply(ctx, events)
}

func AlphaBlend(ctx ilysa.TimeLightContext, events *ilysa.CompoundRGBLightingEvent, startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) {
	if ctx.T() < startT || ctx.T() > endT {
		return
	}
	tScale := scale.ToUnitIntervalClamped(startT, endT)
	alphaScale := scale.FromUnitIntervalClamped(startAlpha, endAlpha)
	events.Mod(ilysa.WithAlpha(alphaScale(tScale(easeFn(ctx.T())))))
}
