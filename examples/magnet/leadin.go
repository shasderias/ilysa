package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/fx"
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
	l.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(3),
		)
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(3),
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

	seqLight := ilysa.NewSequenceLight(
		ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, l),
		ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, l),
	)

	ctx.EventsForSequence(0, []float64{0, 0.25, 0.50}, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.WithLight(seqLight, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(magnetGradient.Ierp(seqCtx.T())),
			)
		})
	})

	comLight := ilysa.NewCombinedLight(
		ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, l),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
		ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, l),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
	)

	ctx.EventForBeat(0.75, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45),
			ilysa.WithStep(9),
			ilysa.WithProp(20),
			ilysa.WithSpeed(1.2),
			ilysa.WithDirection(chroma.Clockwise),
			ilysa.WithCounterSpin(false),
		)

		grad := l.g.Next()

		ctx.EventsForRange(ctx.B(), ctx.B()+0.75, 12, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(comLight, func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 6, grad)
				fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.InCirc)
			})
		})
	})

	var (
		backLasers = ilysa.TransformLight(
			l.NewBasicLight(beatsaber.EventTypeBackLasers),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		)
		colorSweepSpeed   = 2.2
		//shimmerSweepSpeed = 0.8
		intensity         = 0.8
		grad              = magnetRainbowPale
	)

	ctx.EventsForRange(0.75, 0.75+0.7, 30, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(backLasers, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, colorSweepSpeed, grad)
			fx.AlphaBlend(ctx, e, 0, 0.3, 0, intensity, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.4, 1, intensity, 0, ease.InCirc)
			//fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})
}
