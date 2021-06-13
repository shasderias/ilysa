package main

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/ilysa"
)

func LeadIn(p *ilysa.Project) {
	p.EventForBeat(4, func(ctx *ilysa.Context) {
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
	p.EventsForBeats(startBeat, 0.25, 4, func(ctx *ilysa.Context) {
		lights := beatsaber.NewEventTypeSet(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventTypeRightRotatingLasers)
		values := beatsaber.NewEventValueSet(
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightBlueFade,
			beatsaber.EventValueLightBlueFade,
		)

		e := ctx.NewRGBLightingEvent(lights.Pick(ctx.Ordinal), values.Pick(ctx.Ordinal))
		e.SetColor(allColors.Next())
	})

	p.EventForBeat(startBeat+0.75, func(ctx *ilysa.Context) {
		ctx.NewPreciseRotationEvent().PreciseRotation = chroma.PreciseRotation{
			Rotation:    45,
			Step:        9,
			Prop:        20,
			Speed:       1.2,
			Direction:   chroma.Clockwise,
			CounterSpin: false,
		}
	})

	Shimmer(p, startBeat+0.75, startBeat+1.85, 40, beatsaber.EventTypeBackLasers, 2.5, 1.3)
	//SimpleShimmer(p, startBeat + 0.75, startBeat + 0.75 + 1.1)
}
