package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

func NewLeadIn(p *ilysa.Project, startBeat float64) LeadIn {
	return LeadIn{
		p:       p,
		Context: p.Offset(startBeat),
		g: gradient.NewSet(
			shirayukiSingleGradient,
			sukoyaSingleGradient,
		),
	}
}

type LeadIn struct {
	context.Context
	p *ilysa.Project
	g gradient.Set
}

func (l LeadIn) Play() {
	l.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewLaser(evt.WithDirectionalLaser(evt.LeftLaser), evt.WithIntValue(3))
		ctx.NewLaser(evt.WithDirectionalLaser(evt.RightLaser), evt.WithIntValue(3))
	})

	l.BrokenChord(0)
	l.BrokenChord(1.5)
	l.BrokenChord(4)
	l.BrokenChord(5.5)
	l.BrokenChord(8)
	l.BrokenChord(9.5)
}

func (l LeadIn) BrokenChord(startBeat float64) {
	ctx := l.Offset(startBeat)

	seqLight := light.NewSequence(
		light.NewBasic(ctx, evt.LeftRotatingLasers),
		light.NewBasic(ctx, evt.RightRotatingLasers),
	)

	ctx.Sequence(timer.NewSequencer([]float64{0, 0.25, 0.50, 0.75}, 0), func(ctx context.Context) {
		ctx.Light(seqLight, func(ctx context.LightContext) {
			ctx.NewRGBLighting(
				evt.WithLightValue(evt.LightRedFade), evt.WithColor(magnetGradient.Lerp(ctx.SeqT())),
			)
		})
	})

	ctx.Sequence(timer.Beat(0.75), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(8),
			evt.WithProp(20),
			evt.WithRotationSpeed(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
			evt.WithCounterSpin(false),
		)
	})

	var (
		backLasers = transform.Light(
			light.NewBasic(ctx, evt.BackLasers),
			transform.DivideSingle(),
		)
		duration          = 1.2
		steps             = 12
		colorSweepSpeed   = 2.7
		shimmerSweepSpeed = 1.2
		intensity         = 0.8
		grad              = magnetRainbowPale
	)

	rng := timer.NewRanger(0.75, 0.75+duration, steps, ease.Linear)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(backLasers, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, colorSweepSpeed, grad)
			fx.AlphaFade(ctx, e, intensity, 0, ease.InSin)
			fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})
}
