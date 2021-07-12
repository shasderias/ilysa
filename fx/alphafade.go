package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/scale"
)

func AlphaFade(ctx context.LightContext, events evt.RGBLightingEvents, startAlpha, endAlpha float64, easeFn ease.Func) {
	alpha := scale.FromUnitClamp(startAlpha, endAlpha)(easeFn(ctx.T()))

	events.Apply(evt.WithAlpha(alpha))
}

func AlphaFadeEx(ctx context.LightContext, events evt.RGBLightingEvents, startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) {
	if ctx.T() < startT || ctx.T() > endT {
		return
	}

	tScale := scale.ToUnitClamp(startT, endT)
	alphaScale := scale.FromUnitClamp(startAlpha, endAlpha)
	newAlpha := alphaScale(tScale(easeFn(ctx.T())))
	events.Apply(evt.WithAlpha(newAlpha))
}

func NewAlphaFader(startT, endT, startAlpha, endAlpha float64, fn ease.Func) func(context.LightContext, evt.RGBLightingEvents) {
	return func(ctx context.LightContext, e evt.RGBLightingEvents) {
		AlphaFadeEx(ctx, e, startT, endT, startAlpha, endAlpha, fn)
	}
}
