package main

import (
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

func RGBLightingEventsIdentity(ctx context.LightContext, e evt.RGBLightingEvents) {
	return
}

func RingRipple(ctx context.Context, rng timer.Ranger, grad gradient.Table, opts ...ringRippleOpts) evt.RGBLightingEvents {
	o := ringRippleOpt{
		sweepSpeed: 1,
		rippleTime: 1,
		reverse:    false,
		fadeIn:     RGBLightingEventsIdentity,
		fadeOut:    RGBLightingEventsIdentity,
	}

	for _, opt := range opts {
		opt.apply(&o)
	}

	transformOpts := []transform.LightTransformer{
		transform.DivideSingle(),
	}

	if o.reverse {
		transformOpts = append(transformOpts, transform.ReverseSet())
	}

	l := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transformOpts...,
	)

	var e evt.RGBLightingEvents

	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e = fx.ColorSweep(ctx, o.sweepSpeed, grad)
			fx.RippleT(ctx, e, o.rippleTime)
			o.fadeIn(ctx, e)
			o.fadeOut(ctx, e)
		})
	})

	return e
}

type ringRippleOpt struct {
	sweepSpeed float64
	rippleTime float64
	reverse    bool
	fadeIn     func(ctx context.LightContext, e evt.RGBLightingEvents)
	fadeOut    func(ctx context.LightContext, e evt.RGBLightingEvents)
}

type ringRippleOpts interface {
	apply(o *ringRippleOpt)
}

type withFadeInOpt struct {
	fadeIn func(ctx context.LightContext, e evt.RGBLightingEvents)
}

func WithFadeIn(fn func(ctx context.LightContext, e evt.RGBLightingEvents)) withFadeInOpt {
	return withFadeInOpt{fn}
}

func (f withFadeInOpt) apply(o *ringRippleOpt) {
	o.fadeIn = f.fadeIn
}

type withFadeOutOpt struct {
	fadeOut func(ctx context.LightContext, e evt.RGBLightingEvents)
}

func WithFadeOut(fn func(ctx context.LightContext, e evt.RGBLightingEvents)) withFadeOutOpt {
	return withFadeOutOpt{fn}
}

func (f withFadeOutOpt) apply(o *ringRippleOpt) {
	o.fadeOut = f.fadeOut
}

type withRippleTimeOpt struct {
	t float64
}

func WithRippleTime(t float64) withRippleTimeOpt {
	return withRippleTimeOpt{t}
}

func (w withRippleTimeOpt) apply(o *ringRippleOpt) {
	o.rippleTime = w.t
}

type withSweepSpeedOpt struct {
	t float64
}

func WithSweepSpeed(t float64) withSweepSpeedOpt {
	return withSweepSpeedOpt{t}
}

func (w withSweepSpeedOpt) apply(o *ringRippleOpt) {
	o.sweepSpeed = w.t
}

type withReverseOpt struct {
	b bool
}

func WithReverse(b bool) withReverseOpt {
	return withReverseOpt{b}
}

func (w withReverseOpt) apply(o *ringRippleOpt) {
	o.reverse = w.b
}
