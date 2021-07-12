package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

func NewLeadInOut(p *ilysa.Project, startBeat float64) LeadIn {
	return LeadIn{p.Offset(startBeat)}
}

type LeadIn struct {
	context.Context
}

func (l LeadIn) PlayIn() {
	l.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewLaser(evt.WithDirectionalLaser(evt.LeftLaser), evt.WithLaserSpeed(3))
		ctx.NewLaser(evt.WithDirectionalLaser(evt.RightLaser), evt.WithLaserSpeed(3))
	})

	l.BrokenChord(0)
	l.BrokenChord(1.5)
	l.BrokenChord(4)
	l.BrokenChord(5.5)
	l.BrokenChord(8)
	l.BrokenChord(9.5)
}

func (l LeadIn) PlayOut() {
	l.BrokenChordOnly(0)
	l.BrokenChordOnly(1.5)
}

func (l LeadIn) BrokenChord(startBeat float64) {
	ctx := l.BOffset(startBeat)

	seqLight := light.NewSequence(
		light.NewBasic(ctx, evt.LeftRotatingLasers),
		light.NewBasic(ctx, evt.RightRotatingLasers),
	)

	ctx.Sequence(timer.Seq([]float64{0, 0.25, 0.50, 0.75}, 0), func(ctx context.Context) {
		ctx.Light(seqLight, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithLightValue(evt.LightRedFade),
				evt.WithColor(magnetGradient.Lerp(ctx.SeqT())),
				evt.WithAlpha(0.4),
			)
		})
	})

	ctx.Sequence(timer.Beat(0.75), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(8),
			evt.WithRotationSpeed(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
			evt.WithProp(20),
			evt.WithCounterSpin(false),
		)
	})

	var (
		backLasers = transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.DivideSingle(),
		)
		duration          = 1.2
		steps             = 12
		colorSweepSpeed   = 2.7
		shimmerSweepSpeed = 1.2
		intensity         = 0.8
		grad              = magnetRainbowPale
	)

	ctx.Range(timer.Rng(0.75, 0.75+duration, steps, ease.Linear), func(ctx context.Context) {
		ctx.Light(backLasers, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, colorSweepSpeed, grad)
			fx.AlphaFade(ctx, e, intensity, 0, ease.InCirc)
			fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})
}

func (l LeadIn) BrokenChordOnly(startBeat float64) {
	ctx := l.BOffset(startBeat)

	seqLight := light.NewSequence(
		light.NewBasic(ctx, evt.LeftRotatingLasers),
		light.NewBasic(ctx, evt.RightRotatingLasers),
	)

	ctx.Sequence(timer.Seq([]float64{0, 0.25, 0.50, 0.75}, 0), func(ctx context.Context) {
		ctx.Light(seqLight, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithLightValue(evt.LightRedFade),
				evt.WithColor(magnetGradient.Lerp(ctx.SeqT())),
				evt.WithAlpha(0.4),
			)
		})
	})
}
