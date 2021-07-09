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

type Breakdown struct {
	context.Context
	p *ilysa.Project
}

func NewBreakdown(p *ilysa.Project, startBeat float64) Breakdown {
	return Breakdown{
		Context: p.Offset(startBeat),
		p:       p,
	}
}

func (b Breakdown) Play() {
	b.BrokenChord(0)
	b.Chord()
	b.GuitarPlucks()
}

func (b Breakdown) BrokenChord(startBeat float64) {
	ctx := b.Offset(startBeat)
	l := transform.Light(
		light.NewSequence(
			light.NewBasic(ctx, evt.LeftRotatingLasers),
			light.NewBasic(ctx, evt.RightRotatingLasers),
		),
		transform.DivideSingle(),
	)

	//ll := light.NewBasic(beatsaber.EventTypeLeftRotatingLasers, b)
	//rl := light.NewBasic(beatsaber.EventTypeRightRotatingLasers, b)
	//
	//light := transform.Light(
	//	light2.NewSequenceLight(ll, rl),
	//	rework.ToLightTransformer(rework.DivideSingle),
	//).(light2.SequenceLight)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser),
			evt.WithLaserSpeed(3), evt.WithPreciseLaserSpeed(4.5),
		)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser),
			evt.WithLaserSpeed(3), evt.WithPreciseLaserSpeed(4.5),
		)
	})

	seq := timer.Interval(4, 1, 8)
	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(22.5),
			evt.WithRotationStep(float64(ctx.Ordinal())*2.5),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)
	})

	seq = timer.Interval(0, 1, 12)
	ctx.Sequence(seq, func(ctx context.Context) {
		rng := timer.NewRanger(0, 0.8, 12, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			grad := gradient.New(crossickColors.Next(), crossickColors.Next())
			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				fx.RippleT(ctx, e, 1.2)
				fx.AlphaFadeEx(ctx, e, 0, 0.3, 0, 0.6, ease.InSin)
				fx.AlphaFadeEx(ctx, e, 0.3, 1, 0.6, 0, ease.OutSin)
			})
		})
	})
}

func (b Breakdown) Chord() {
	seq := timer.NewSequencer([]float64{
		0, 1, 2, 3, 4, 5, 6,
		7, 7.5, 7.75,
		8, 8.5,
		9.25,
		10.0, 10.5,
		11.0, 11.25, 11.5, 11.75,
		12.0, 12.25, 12.5, 12.75,
	}, 13.50)

	ctx := b.Offset(0)

	l := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.DivideSingle(),
	)

	b.Sequence(seq, func(ctx context.Context) {
		grad := gradient.New(
			crossickColors.Idx(ctx.Ordinal()),
			crossickColors.Idx(ctx.Ordinal()),
		)

		//nb := ctx.SequenceNextB() - 0.25
		//if ctx.Last() {
		//	nb = ctx.B() + 0.5
		//}

		alpha := ease.InCubic(ctx.T()) * 6

		rng := timer.NewRanger(0, ctx.SeqNextBOffset(), 24, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(l, func(ctx context.LightContext) {
				//e := fx.Gradient(ctx, grad)
				e := fx.ColorSweep(ctx, 1.2, grad)
				fx.RippleT(ctx, e, 1.4)
				fx.AlphaFadeEx(ctx, e, 0, 1, alpha, 0, ease.OutQuart)
			})
		})
	})
}

func (b Breakdown) GuitarPlucks() {
	ctx := b.Offset(12.75)

	var (
		backLasers = transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.DivideSingle())
		colorSweepSpeed   = 2.2
		shimmerSweepSpeed = 0.8
		grad              = magnetRainbowPale
	)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(8),
			evt.WithRotationSpeed(20),
			evt.WithProp(0.4),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)
	})

	rng := timer.NewRanger(0, 1.5, 30, ease.Linear)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(backLasers, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, colorSweepSpeed, grad)
			fx.AlphaFadeEx(ctx, e, 0, 0.5, 0, 0.8, ease.InCubic)
			fx.AlphaFadeEx(ctx, e, 0.5, 1.0, 0.8, 0, ease.InCubic)
			fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})
}
