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

func NewLeadIn(p *ilysa.Project, startBeat float64) LeadIn {
	return LeadIn{
		p:           p,
		BaseContext: p.WithBeatOffset(startBeat),
		g: gradient.NewSet(
			shirayukiSingleGradient,
			sukoyaSingleGradient,
		),
	}
}

type LeadIn struct {
	p *ilysa.Project
	ilysa.BaseContext
	g gradient.Set
}

func (l LeadIn) Play() {
	l.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(3),
		)
		ctx.NewLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(3),
		)

	})

	l.BrokenChord(0)
	l.BrokenChord(1.5)
	l.BrokenChord(4)
	l.BrokenChord(5.5)
	l.BrokenChord(8)
	l.BrokenChord(9.5)
}

func (l LeadIn) BrokenChord(startBeat float64) {
	ctx := l.WithBeatOffset(startBeat)

	seqLight := light2.NewSequenceLight(
		light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, l),
		light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, l),
	)

	ctx.EventsForSequence(0, []float64{0, 0.25, 0.50}, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.WithLight(seqLight, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				evt.WithColor(magnetGradient.Ierp(seqCtx.T())),
			)
		})
	})

	comLight := light2.NewCombinedLight(
		light2.TransformLight(
			light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, l),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
		light2.TransformLight(
			light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, l),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
	)

	ctx.EventForBeat(0.75, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(9),
			evt.WithProp(20),
			evt.WithPreciseLaserSpeed(1.2),
			evt.WithLaserDirection(chroma.Clockwise),
			evt.WithCounterSpin(false),
		)

		grad := l.g.Next()

		ctx.Range(ctx.B(), ctx.B()+0.75, 12, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(comLight, func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 6, grad)
				fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.InCirc)
			})
		})
	})

	var (
		backLasers = light2.TransformLight(
			l.NewBasicLight(beatsaber.EventTypeBackLasers),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		)
		colorSweepSpeed   = 2.2
		//shimmerSweepSpeed = 0.8
		intensity         = 0.8
		grad              = magnetRainbowPale
	)

	ctx.EventsForRange(0.75, 0.75+0.7, 30, ease.Linear, func(ctx ilysa.RangeContext) {
		ctx.WithLight(backLasers, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, colorSweepSpeed, grad)
			fx.AlphaBlend(ctx, e, 0, 0.3, 0, intensity, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.4, 1, intensity, 0, ease.InCirc)
			//fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})
}
