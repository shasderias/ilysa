package fx

import (
	"math/rand"

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

type alphaFadeOpt struct {
	context.LightContext
	startT, endT float64
	startA, endA float64
	easeFn       ease.Func
}

func (a alphaFadeOpt) Apply(e evt.Event) {
	if a.T() < a.startT || a.T() > a.endT {
		return
	}

	var (
		tScale   = scale.ToUnitClamp(a.startT, a.endT)
		aScale   = scale.FromUnitClamp(a.startA, a.endA)
		newAlpha = aScale(a.easeFn(tScale(a.T())))
	)
	e.Apply(evt.OptAlpha(newAlpha))
}

func AlphaFade2(ctx context.LightContext,
	startT, endT, startA, endA float64, easeFn ease.Func) evt.Option {
	return alphaFadeOpt{
		LightContext: ctx,
		startT:       startT, endT: endT,
		startA: startA, endA: endA,
		easeFn: easeFn,
	}

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

func AlphaJitter(ctx context.LightContext, maxJitter float64) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		ae, ok := e.(alphaer)
		if !ok {
			return
		}
		ae.SetAlpha(ae.Alpha() + (rand.Float64()*2-1)*maxJitter)
	})
}

type alphaer interface {
	Alpha() float64
	SetAlpha(a float64)
}
