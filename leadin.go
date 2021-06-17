package main

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ease"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/ilysa/fx"
)

func LeadIn(p *ilysa.Project) {
	p.EventForBeat(4, func(ctx ilysa.TimingContext) {
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithValue(3),
		)
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithValue(3),
		)

	})

	LeadInBrokenChord(p, 4)
	LeadInBrokenChord(p, 5.5)
	LeadInBrokenChord(p, 8)
	LeadInBrokenChord(p, 9.5)
	LeadInBrokenChord(p, 12)
	LeadInBrokenChord(p, 13.5)
}

func LeadInBrokenChord(p *ilysa.Project, startBeat float64) {
	//p.EventsForRange(0, 10, 30, ease.Linear, func(ctx ilysa.Timing) {
	//	ctx.UseLight(backLasers, func(ctx ilysa.Timing) {
	//		ctx.NewRGBLightingEvent().SetValue(off).SetColor()
	//	})
	//})

	p.EventsForBeats(startBeat, 0.25, 4, func(ctx ilysa.TimingContext) {
		lights := beatsaber.NewEventTypeSet(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventTypeRightRotatingLasers)
		values := beatsaber.NewEventValueSet(
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightBlueFade,
			beatsaber.EventValueLightBlueFade,
		)

		ctx.NewRGBLightingEvent(
			ilysa.WithType(lights.Pick(ctx.Ordinal())),
			ilysa.WithValue(values.Pick(ctx.Ordinal())),
			ilysa.WithColor(allColors.Next()),
		)
	})

	p.EventForBeat(startBeat+0.75, func(ctx ilysa.TimingContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45),
			ilysa.WithStep(9),
			ilysa.WithProp(20),
			ilysa.WithSpeed(1.2),
			ilysa.WithDirection(chroma.Clockwise),
			ilysa.WithCounterSpin(false),
		)
	})

	var (
		backLasers        = p.NewBasicLight(beatsaber.EventTypeBackLasers).Split(ilysa.DivideSingle)
		colorSweepSpeed   = 2.2
		shimmerSweepSpeed = 0.8
		intensity         = 0.8
		grad              = gradient.Rainbow
		shimmerStart      = startBeat + 0.75
		shimmerEnd        = startBeat + 0.75 + 1.1
		fadeIn            = startBeat + 0.75 + 0.3
		fadeOut           = startBeat + 0.75 + 0.7
	)

	p.EventsForRange(shimmerStart, shimmerEnd, 30, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(backLasers, func(ctx ilysa.TimingContextWithLight) {
			e := fx.ColorSweep(ctx, colorSweepSpeed, grad)
			e.Mod(ilysa.WithAlpha(intensity))
			fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})

	p.ModEventsInRange(shimmerStart, fadeIn, ilysa.FilterRGBLight(backLasers), func(ctx ilysa.TimingContext, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
	})

	p.ModEventsInRange(fadeOut, shimmerEnd, ilysa.FilterRGBLight(backLasers), func(ctx ilysa.TimingContext, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
	})
}
