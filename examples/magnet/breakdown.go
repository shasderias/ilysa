package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light2"
)

type Breakdown struct {
	ilysa.BaseContext
	p *ilysa.Project
}

func NewBreakdown(p *ilysa.Project, startBeat float64) Breakdown {
	return Breakdown{
		BaseContext: p.WithBeatOffset(startBeat),
		p:           p,
	}
}

func (b Breakdown) Play() {
	b.BrokenChord(0)
	b.Chord()
	b.GuitarPlucks()
}

func (b Breakdown) BrokenChord(startBeat float64) {
	ll := light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, b)
	rl := light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, b)

	light := light2.TransformLight(
		light2.NewSequenceLight(ll, rl),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(light2.SequenceLight)

	b.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
		)
	})

	b.EventsForBeats(4, 1, 8, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(22.5),
			evt.WithRotationStep(float64(ctx.Ordinal())*2.5),
			evt.WithPreciseLaserSpeed(20),
			evt.WithProp(1.2),
			evt.WithLaserDirection(chroma.Clockwise),
		)
	})

	b.EventsForBeats(0, 1, 12, func(ctx ilysa.RangeContext) {
		l := light.Index(ctx.Ordinal())
		ctx.Range(ctx.B(), ctx.B()+0.80, 12, ease.Linear, func(ctx ilysa.RangeContext) {
			grad := gradient.New(crossickColors.Next(), crossickColors.Next())
			ctx.WithLight(l, func(ctx ilysa.TimeLightContext) {
				e := fx.Gradient(ctx, grad)
				fx.Ripple(ctx, e, 1.2,
					fx.WithAlphaBlend(0, 0.3, 0, 0.6, ease.InSine),
					fx.WithAlphaBlend(0.3, 1, 0.6, 0, ease.OutSine),
				)
			})
		})

	})
}

func (b Breakdown) Chord() {
	seq := []float64{
		0, 1, 2, 3, 4, 5, 6,
		7, 7.5, 7.75,
		8, 8.5,
		9.25,
		10.0, 10.5,
		11.0, 11.25, 11.5, 11.75,
		12.0, 12.25, 12.5, 12.75,
	}

	light := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeRingLights, b),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	b.Sequence(0, seq, func(ctx ilysa.SequenceContext) {
		grad := gradient.New(
			crossickColors.Index(ctx.Ordinal()),
			crossickColors.Index(ctx.Ordinal()),
		)

		nb := ctx.SequenceNextB() - 0.25
		if ctx.Last() {
			nb = ctx.B() + 0.5
		}

		alpha := ease.InCubic(ctx.T()) * 6

		ctx.EventsForRange(ctx.B(), nb, 24, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
				//e := fx.Gradient(ctx, grad)
				e := fx.ColorSweep(ctx, 1.2, grad)
				fx.Ripple(ctx, e, 1.4)
				fx.AlphaBlend(ctx, e, 0, 1, alpha, 0, ease.OutQuart)
			})
		})
	})
}

func (b Breakdown) GuitarPlucks() {
	var (
		backLasers        = b.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.DivideSingle)
		colorSweepSpeed   = 2.2
		shimmerSweepSpeed = 0.8
		grad              = magnetRainbowPale
		startBeat         = 12.75
		endBeat           = startBeat + 1.5
	)

	b.EventForBeat(startBeat, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(8),
			evt.WithPreciseLaserSpeed(20),
			evt.WithProp(0.4),
			evt.WithLaserDirection(chroma.CounterClockwise),
		)
	})

	b.Range(startBeat, endBeat, 30, ease.Linear, func(ctx ilysa.RangeContext) {
		ctx.WithLight(backLasers, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, colorSweepSpeed, grad)
			fx.AlphaBlend(ctx, e, 0, 0.5, 0, 0.8, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.5, 1.0, 0.8, 0, ease.InCubic)
			fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})
}
