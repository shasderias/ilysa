package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/fx"
)

type Verse1b struct {
	ilysa.BaseContext
	offset float64
}

func NewVerse1b(project *ilysa.Project, offset float64) Verse1b {
	return Verse1b{
		BaseContext: project.WithBeatOffset(offset),
	}
}

func (v Verse1b) Play() {
	v.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(5),
		)
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(5),
		)
	})

	v.Rhythm(0)
	v.Rhythm(4)
	v.Rhythm(8)
	v.Rhythm(12)
	v.Rhythm(16)
	v.Rhythm(20)
	v.Rhythm(24)
}

func (v Verse1b) Rhythm(startBeat float64) {
	var (
		leftLaser  = ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, v)
		rightLaser = ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, v)
		beatLight  = ilysa.NewSequenceLight(leftLaser, rightLaser)
		color      = crossickColors
	)

	v.EventsForSequence(startBeat, []float64{0, 1, 2, 3}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(22.5),
			ilysa.WithSpeed(2),
			ilysa.WithProp(0.3),
			ilysa.WithDirection(chroma.CounterClockwise),
		)

		ctx.WithLight(beatLight, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(color.Next()),
			)
		})
	})

	var (
		rippleDuration = 1.0
		rippleStart    = startBeat + 2
		rippleEnd      = rippleStart + rippleDuration
		rippleLights   = v.NewBasicLight(beatsaber.EventTypeRingLights).Transform(ilysa.DivideSingle)
		rippleStep     = 0.8
		grad           = gradient.Table{
			{magnetColors.Index(0), 0.0},
			{magnetColors.Index(1), 0.05},
			{magnetColors.Index(2), 0.5},
			{magnetColors.Index(3), 1.0},
		}
	)

	v.EventsForRange(rippleStart, rippleEnd, 30, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(rippleLights, func(ctx ilysa.TimeLightContext) {
			events := fx.BiasedColorSweep(ctx, 3, grad)

			fx.Ripple(ctx, events, rippleStep,
				fx.WithAlphaBlend(0, 0.2, 0, 2, ease.InCubic),
				fx.WithAlphaBlend(0.8, 1, 2, 0, ease.OutCubic),
			)
			//events.Mod(ilysa.WithAlpha(1.5))
			//for _, ee := range *events {
			//	ee.ShiftBeat(ctx.LightIDT() * rippleStep)
			//}
			//switch {
			//case ctx.T() <= 0.5:
			//	alphaScale := scale.ClampedToUnitInterval(0, 0.5)
			//	events.Mod(ilysa.WithAlpha(events.GetAlpha() * ease.InOutQuart(alphaScale(ctx.T()))))
			//case ctx.T() > 0.8:
			//	alphaScale := scale.ClampedToUnitInterval(0.8, 1)
			//	events.Mod(ilysa.WithAlpha(events.GetAlpha() * ease.InExpo(1-alphaScale(ctx.T()))))
			//}
		})
	})
}
