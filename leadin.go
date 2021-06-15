package main

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/chroma/lightid"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ease"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/ilysa/fx"
)

func LeadIn(p *ilysa.Project) {
	p.EventForBeat(4, func(ctx ilysa.Context) {
		ctx.NewRotationSpeedEvent(ilysa.LeftLaser, 3)
		ctx.NewRotationSpeedEvent(ilysa.RightLaser, 3)
	})

	LeadInBrokenChord(p, 4)
	LeadInBrokenChord(p, 5.5)
	LeadInBrokenChord(p, 8)
	LeadInBrokenChord(p, 9.5)
	LeadInBrokenChord(p, 12)
	LeadInBrokenChord(p, 13.5)
}

func LeadInBrokenChord(p *ilysa.Project, startBeat float64) {
	p.EventsForRange(0, 10, 30, ease.Linear, func(ctx ilysa.Context) {
		ctx.UseLight(backLasers, func(ctx ilysa.Context) {
			ctx.NewRGBLightingEvent().SetValue(off).SetColor()

		})
	})

	p.EventsForBeats(startBeat, 0.25, 4, func(ctx ilysa.Context) {
		lights := beatsaber.NewEventTypeSet(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventTypeRightRotatingLasers)
		values := beatsaber.NewEventValueSet(
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightBlueFade,
			beatsaber.EventValueLightBlueFade,
		)

		e := ctx.NewRGBLightingEvent().
			SetLight(lights.Pick(ctx.Ordinal)).
			SetValue(values.Pick(ctx.Ordinal))

		e.SetColor(allColors.Next())
	})

	p.EventForBeat(startBeat+0.75, func(ctx ilysa.Context) {
		ctx.NewPreciseRotationEvent().PreciseRotation = chroma.PreciseRotation{
			Rotation:    45,
			Step:        9,
			Prop:        20,
			Speed:       1.2,
			Direction:   chroma.Clockwise,
			CounterSpin: false,
		}
	})

	var (
		backLasers        = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		colorSweepSpeed   = 2.2
		shimmerSweepSpeed = 0.8
		intensity         = 0.8
		grad              = gradient.Rainbow
		shimmerStart      = startBeat + 0.75
		shimmerEnd        = startBeat + 0.75 + 1.1
		fadeIn            = startBeat + 0.75 + 0.3
		fadeOut           = startBeat + 0.75 + 0.7
	)

	p.EventsForRange(shimmerStart, shimmerEnd, 30, ease.Linear, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(backLasers, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
			e := fx.ColorSweep(ctx, intensity, colorSweepSpeed, grad)
			fx.AlphaShimmer(ctx, e, shimmerSweepSpeed)
		})
	})

	p.ModEventsInRange(shimmerStart, fadeIn, ilysa.FilterRGBLight(backLasers), func(ctx ilysa.Context, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
	})

	p.ModEventsInRange(fadeOut, shimmerEnd, ilysa.FilterRGBLight(backLasers), func(ctx ilysa.Context, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
	})
}
