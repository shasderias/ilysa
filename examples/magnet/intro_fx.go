package main

import (
	"math"

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
		steps      = 6
		sweepSpeed = 3.0
	)

	ctx.Sequence(seq, func(ctx context.Context) {
		grad := magnetRainbowPale.RotateRand()
		ctx.Range(timer.Rng(0, duration, steps, ease.Linear), func(ctx context.Context) {
			ctx.Light(lightSweepDiv, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, sweepSpeed, grad)
				fx.RippleT(ctx, e, 0.1)
				fx.AlphaFadeEx(ctx, e, 0, 1, 0.6, 0, ease.InCirc)
			})
		})
	})
}

func (p Intro) SlowMotionLasers(ctx context.Context, rng timer.Ranger, startSpeed, endSpeed float64) {
	ctx.Range(rng, func(ctx context.Context) {
		speed := (1-ctx.T())*startSpeed + endSpeed
		intSpeed := int(math.Round(speed))

		opts := evt.NewOpts(evt.WithLaserSpeed(intSpeed), evt.WithPreciseLaserSpeed(speed))
		if !ctx.First() {
			opts.Add(evt.WithLockPosition(true))
		}

		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser), opts)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser), opts)
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
		rng := timer.Rng(0, 0.435, 7, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(backLasers, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaFadeEx(ctx, e, 0, 0.3, 1, 3, ease.OutCubic)
				fx.AlphaFadeEx(ctx, e, 0.3, 1, 3, 1, ease.InCubic)
			})
		})
	})
}

func (p Intro) Rush(ctx context.Context, startBeat, endBeat, rippleDuration, peakAlpha float64, grad gradient.Table) {
	ringLasers := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.Reverse(),
		transform.DivideSingle(),
	)

	steps := int(math.Max(math.RoundToEven(endBeat-startBeat)*12, 12))

	rng := timer.Rng(startBeat, endBeat, steps, ease.InExpo)

	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(ringLasers, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, 2, grad)
			fx.RippleT(ctx, e, rippleDuration)
			fx.AlphaFadeEx(ctx, e, 0, 0.6, 1, peakAlpha, ease.OutCirc)
			fx.AlphaFadeEx(ctx, e, 0.6, 1.0, peakAlpha, 0, ease.InCirc)
		})
	})
}
