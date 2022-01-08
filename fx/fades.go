package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/scale"
)

func AlphaFade(ctx context.LightContext, e evt.Events,
	startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) {
	if ctx.T() < startT || ctx.T() > endT {
		return
	}

	var (
		tScale   = scale.ToUnitClamp(startT, endT)
		aScale   = scale.FromUnitClamp(startAlpha, endAlpha)
		newAlpha = aScale(easeFn(tScale(ctx.T())))
	)
	e.Apply(evt.OptAlpha(newAlpha))
}

func FloatValueFade(ctx context.LightContext, e evt.Events,
	startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) {
	if ctx.T() < startT || ctx.T() > endT {
		return
	}

	var (
		tScale        = scale.ToUnitClamp(startT, endT)
		fScale        = scale.FromUnitClamp(startAlpha, endAlpha)
		newFloatValue = fScale(easeFn(tScale(ctx.T())))
	)
	e.Apply(evt.OptFloatValue(newFloatValue))
}

func AlphaMultiply(ctx context.LightContext, e evt.Events,
	startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) {
	if ctx.T() < startT || ctx.T() > endT {
		return
	}

	var (
		tScale          = scale.ToUnitClamp(startT, endT)
		aScale          = scale.FromUnitClamp(startAlpha, endAlpha)
		alphaMultiplier = aScale(easeFn(tScale(ctx.T())))
	)

	for _, evt := range e {
		ae, ok := evt.(alphaer)
		if !ok {
			continue
		}
		ae.SetAlpha(ae.Alpha() * alphaMultiplier)
	}
}

type alphaer interface {
	Alpha() float64
	SetAlpha(a float64)
}
