package fx

import (
	"math/rand"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/scale"
)

func AlphaFade(ctx context.LightContext,
	startT, endT, startA, endA float64, easeFn ease.Func) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if ctx.T() < startT || ctx.T() > endT {
			return
		}

		var (
			tScale   = scale.ToUnitClamp(startT, endT)
			aScale   = scale.FromUnitClamp(startA, endA)
			newAlpha = aScale(easeFn(tScale(ctx.T())))
		)
		e.Apply(evt.OptAlpha(newAlpha))
	})
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

func FloatValueFade(ctx context.LightContext,
	startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if ctx.T() < startT || ctx.T() > endT {
			return
		}

		var (
			tScale        = scale.ToUnitClamp(startT, endT)
			fScale        = scale.FromUnitClamp(startAlpha, endAlpha)
			newFloatValue = fScale(easeFn(tScale(ctx.T())))
		)
		e.Apply(evt.OptFloatValue(newFloatValue))
	})
}

func AlphaMultiply(ctx context.LightContext,
	startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if ctx.T() < startT || ctx.T() > endT {
			return
		}

		var (
			tScale          = scale.ToUnitClamp(startT, endT)
			aScale          = scale.FromUnitClamp(startAlpha, endAlpha)
			alphaMultiplier = aScale(easeFn(tScale(ctx.T())))
		)

		ae, ok := e.(alphaer)
		if !ok {
			return
		}
		ae.SetAlpha(ae.Alpha() * alphaMultiplier)
	})
}

type alphaer interface {
	Alpha() float64
	SetAlpha(a float64)
}
