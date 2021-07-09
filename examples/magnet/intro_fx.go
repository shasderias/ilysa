package main

import (
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

func (p Intro) PianoGlow(ctx context.Context, seq timer.Sequencer, divisor int, duration float64, divideSingle, shuffle bool) {
	transforms := []transform.LightTransformer{
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(divisor).Sequence(),
	}

	if divideSingle {
		transforms = append(transforms, transform.DivideSingle())
	}

	lightSweepDiv := transform.Light(
		light.NewBasic(ctx, evt.BackLasers),
		transforms...,
	).(light.Sequence)

	if shuffle {
		lightSweepDiv = lightSweepDiv.Shuffle()
	}

	var (
		steps      = 8
		sweepSpeed = 3.0
	)

	ctx.Sequence(seq, func(ctx context.Context) {
		grad := magnetRainbowPale.RotateRand()
		ctx.Range(timer.NewRanger(0, duration, steps, ease.Linear), func(ctx context.Context) {
			ctx.Light(lightSweepDiv, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, sweepSpeed, grad)
				fx.RippleT(ctx, e, 0.05)
				fx.AlphaFadeEx(ctx, e, 0, 1, 1, 0, ease.OutCubic)
			})
		})
	})
}

func (p Intro) PianoTransmute(ctx context.Context, sequence timer.Sequencer, divisor int, shuffle bool, grad gradient.Table) {
	backLasers := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(divisor).Sequence(),
		transform.DivideSingle(),
	).(light.Sequence)

	if shuffle {
		backLasers = backLasers.Shuffle()
	}

	ctx.Sequence(sequence, func(ctx context.Context) {
		rng := timer.NewRanger(0, 0.435, 12, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(backLasers, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaFadeEx(ctx, e, 0, 0.3, 1, 3, ease.OutCubic)
				fx.AlphaFadeEx(ctx, e, 0.3, 1, 3, 1, ease.InCubic)
			})
		})
	})
}

func (p Intro) Rush(ctx context.Context, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	ringLasers := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.Reverse(),
		transform.DivideSingle(),
	)

	rng := timer.NewRanger(startBeat, endBeat, 30, ease.InExpo)

	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(ringLasers, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, 2, grad)
			fx.RippleT(ctx, e, step)
			fx.AlphaFadeEx(ctx, e, 0, 0.6, 1, peakAlpha, ease.OutCubic)
			fx.AlphaFadeEx(ctx, e, 0.6, 1.0, peakAlpha, 0, ease.InCubic)
		})
	})
}
