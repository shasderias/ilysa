package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
)

// Ripple offsets each successive lightID by step beats
func Ripple(ctx context.LightContext, events evt.RGBLightingEvents, step float64) {
	events.Apply(evt.WithBeatOffset(float64(ctx.LightIDOrdinal()) * step))
}

// RippleT offsets each successive lightID by a number of beats such that the last
// lightID triggers delay beats after the first
func RippleT(ctx context.LightContext, events evt.RGBLightingEvents, delay float64, opts ...rippleTOpt) {
	t := ctx.LightIDT()
	for _, o := range opts {
		t = o.easeT(t)
	}

	events.Apply(evt.WithBeatOffset(t * delay))
}

type rippleTOpt interface {
	easeT(float64) float64
}

type easeTOpt struct {
	fn ease.Func
}

func EaseT(f ease.Func) easeTOpt {
	return easeTOpt{f}
}

func (o easeTOpt) easeT(t float64) float64 {
	if o.fn == nil {
		return ease.Linear(t)
	}
	return o.fn(t)
}

//
//type withFadeInOpt struct {
//	startT     float64
//	endT       float64
//	startAlpha float64
//	endAlpha   float64
//	easeFn     ease.Func
//}
//
//func WithAlphaBlend(startT, endT, startAlpha, endAlpha float64, easeFn ease.Func) withFadeInOpt {
//	return withFadeInOpt{
//		startT:     startT,
//		endT:       endT,
//		startAlpha: startAlpha,
//		endAlpha:   endAlpha,
//		easeFn:     easeFn,
//	}
//}
//
//func (o withFadeInOpt) apply(ctx context.LightContext, events evt.RGBLightingEvents) {
//	if ctx.T() < o.startT || ctx.T() > o.endT {
//		return
//	}
//	alphaScale := scale.Clamped(o.startT, o.endT, o.startAlpha, o.endAlpha)
//	events.Mod(evt.WithAlpha(events.Alpha() * o.easeFn(alphaScale(ctx.T()))))
//}
//
//func (o withFadeInOpt) applyRipple(ctx context.LightContext, events evt.RGBLightingEvents) {
//	o.apply(ctx, events)
//}
//
//func (o withFadeInOpt) applyColorSweep(ctx context.LightContext, events evt.RGBLightingEvents) {
//	o.apply(ctx, events)
//}
//
