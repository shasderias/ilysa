package fx

import (
	"ilysa/pkg/ease"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/util"
)

type RippleOpt interface {
	applyRipple(ilysa.TimingContextWithLight, *ilysa.CompoundRGBLightingEvent)
}

func Ripple(ctx ilysa.TimingContextWithLight, events *ilysa.CompoundRGBLightingEvent, step float64, opts ...RippleOpt) {
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

func (o withFadeInOpt) applyRipple(ctx ilysa.TimingContextWithLight, events *ilysa.CompoundRGBLightingEvent) {
	if ctx.T() < o.startT || ctx.T() > o.endT {
		return
	}
	alphaScale := util.Scale(o.startT, o.endT, o.startAlpha, o.endAlpha)
	events.Mod(ilysa.WithAlpha(events.GetAlpha() * o.easeFn(alphaScale(ctx.T()))))
}
