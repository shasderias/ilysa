package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/scale"
)

type Verse struct {
	ilysa.BaseContext
	project *ilysa.Project
	offset  float64
}

func NewVerse1a(p *ilysa.Project, offset float64) Verse {
	ctx := p.WithBeatOffset(offset)
	return Verse{
		BaseContext: ctx,
		project:     p,
		offset:      offset,
	}
}

func (p Verse) Play() {
	p.EventForBeat(-0.01, func(ctx ilysa.TimeContext) {
		fx.OffAll(ctx)
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithSpeed(1.5))
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithSpeed(1.5))
	})

	p.Rhythm(0, false)
	p.Rhythm(4, false)
	p.Rhythm(8, false)
	p.Rhythm(12, false)
	p.Rhythm(16, false)
	p.Rhythm(20, false)
	p.Rhythm(24, false)
	p.Rhythm(28, true)

	p.PianoBackstep(7)
	p.PianoBackstep(15)
	p.PianoBackstep(23)
	p.RinPun(27)
}

func (p Verse) Rhythm(startBeat float64, kickOnly bool) {
	var (
		kickDrumLight = ilysa.NewSequenceLight(
			p.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers),
			p.NewBasicLight(beatsaber.EventTypeRightRotatingLasers),
		)
		kickDrumSequence = []float64{0, 2.5}
		kickDrumColors   = magnetColors
	)
	p.EventsForSequence(startBeat, kickDrumSequence, func(ctx ilysa.SequenceContext) {
		ctx.WithLight(kickDrumLight, func(ctx ilysa.SequenceLightContext) {
			e := ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(kickDrumColors.Next()),
			)

			if startBeat == 0 && ctx.Ordinal() == 0 {
				e.SetAlpha(10)
			} else {
				e.SetAlpha(0.7)
			}
		})
	})

	if kickOnly {
		return
	}

	const (
		rippleDuration = 2
	)

	var (
		rippleStart  = startBeat
		rippleEnd    = rippleStart + rippleDuration
		rippleLights = p.NewBasicLight(beatsaber.EventTypeRingLights).Transform(ilysa.DivideSingle)
		rippleStep   = 0.6
		grad         = gradient.Table{
			{shirayukiPurple, 0.0},
			{shirayukiGold, 0.3},
			{shirayukiGold, 0.7},
			{shirayukiPurple, 1.0},
		}
	)

	p.EventForBeat(rippleStart, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(22.5),
			ilysa.WithSpeed(2),
			ilysa.WithProp(0.3),
		)
	})

	p.EventsForRange(rippleStart, rippleEnd, 30, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(rippleLights, func(ctx ilysa.TimeLightContext) {
			events := fx.ColorSweep(ctx, 1.4, grad)
			events.Mod(ilysa.WithAlpha(1.5))
			for _, ee := range *events {
				ee.ShiftBeat(ctx.LightIDT() * rippleStep)
			}
			switch {
			case ctx.T() <= 0.5:
				alphaScale := scale.ToUnitIntervalClamped(0, 0.5)
				events.Mod(ilysa.WithAlpha(events.GetAlpha() * ease.InOutQuart(alphaScale(ctx.T()))))
			case ctx.T() > 0.8:
				alphaScale := scale.ToUnitIntervalClamped(0.8, 1)
				events.Mod(ilysa.WithAlpha(events.GetAlpha() * ease.InExpo(1-alphaScale(ctx.T()))))
			}
		})
	})
}

//func (p Verse) Lyrics(startBeat float64) {
//	var (
//		// 52-58.5
//		sequence = []float64{0, 0.5, 0.75, 1.25, 1.75, 2.25, 2.75, 3.0, 3.5, 4.0, 4.5, 4.75, 5.25, 5.75, 6.25, 6.5}
//		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers)
//	)
//
//	p.EventsForSequence(startBeat+, sequence, func(ctx ilysa.Context) {
//		ctx.WithLight(light, lightid.GroupDivide(3), func(ctx ilysa.ContextWithLight) {
//			if ctx.ordinal%ctx.LightIDSetLen != ctx.LightIDOrdinal {
//				return
//			}
//
//			ctx.NewRGBLightingEvent().
//				SetValue(beatsaber.EventValueLightOff).
//				SetLightID(ctx.PreLightID)
//
//			ctx.NewRGBLightingEvent().
//				SetValue(beatsaber.EventValueLightRedOn).
//				SetColor(allColors.Next())
//		})
//	})
//}
//
func (p Verse) PianoBackstep(startBeat float64) {
	var (
		sequence = []float64{0, 0.5}
		//backLight = p.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.Divide(2))
	)

	backLasers := p.NewBasicLight(beatsaber.EventTypeBackLasers)
	backLasersSplit := backLasers.LightIDSequenceTransform(ilysa.Divide(2))
	backLasersSplitSingle := backLasersSplit.(ilysa.SequenceLight).LightIDTransform(ilysa.DivideSingle)

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		e := ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(5),
			ilysa.WithStep(15),
			ilysa.WithProp(12),
			ilysa.WithSpeed(10),
		)

		if ctx.Ordinal()%2 == 0 {
			e.Mod(ilysa.WithDirection(chroma.Clockwise))
		} else {
			e.Mod(ilysa.WithDirection(chroma.CounterClockwise))
		}

		ctx.EventsForRange(ctx.B(), ctx.B()+0.25, 6, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(backLasersSplitSingle, func(ctx ilysa.TimeLightContext) {
				fx.Gradient(ctx, magnetGradient)
			})
		})

		//ctx.ModEventsInRange(ctx.b, ctx.b+0.10, ilysa.FilterRGBLight(backLight), func(ctx ilysa.Context, event ilysa.Event) {
		//	fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCirc)
		//})
		ctx.ModEventsInRange(ctx.B(), ctx.B()+0.25, ilysa.FilterRGBLight(backLasersSplitSingle), func(ctx ilysa.TimeContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
		})

		//ctx.WithLight(backLight, lightid.GroupDivide(2), func(ctx ilysa.ContextWithLight) {
		//	if ctx.ordinal != ctx.LightIDOrdinal {
		//		return
		//	}
		//
		//	fx.BiasedColorSweep(ctx, 2, 0.8, gradient.Rainbow)
		//})
	})
}

func (p Verse) RinPun(startBeat float64) {
	backLasers := ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p.project.ActiveDifficultyProfile())
	sl := ilysa.TransformLight(backLasers,
		ilysa.ToSequenceLightTransformer(ilysa.Fan(2)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(ilysa.SequenceLight)

	sl = ilysa.NewSequenceLight(
		sl.Index(0),
		ilysa.TransformLight(sl.Index(1),
			ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
		))

	p.EventForBeat(startBeat, func(ctx ilysa.TimeContext) {
		ctx.NewZoomEvent()
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45),
			ilysa.WithStep(12),
			ilysa.WithSpeed(3),
			ilysa.WithProp(2),
			ilysa.WithDirection(chroma.CounterClockwise),
		)
	})
	p.EventsForSequence(startBeat, []float64{0, 0.5}, func(ctx ilysa.SequenceContext) {

		seqOrdinal := ctx.Ordinal()
		ctx.EventsForRange(ctx.B(), ctx.B()+0.7, 30, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(sl.Index(seqOrdinal), func(ctx ilysa.TimeLightContext) {
				events := fx.ColorSweep(ctx, 0.8, gradient.Rainbow)
				fx.Ripple(ctx, events, 0.30,
					fx.WithAlphaBlend(0, 0.5, 0, 2, ease.InCubic),
					fx.WithAlphaBlend(0.5, 1.0, 2, 0, ease.OutCirc),
					//fx.WithAlphaBlend(0.4, 1.0, 0, 0, ease.OutCirc),
				)
			})
		})
	})

	//p.EventsForRange(startBeat, startBeat+0.95, 30, ease.Linear, func(ctx ilysa.TimeContext) {
	//	ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
	//		fx.BiasedColorSweep(ctx, 1.2, gradient.Rainbow)
	//	})
	//})

	//p.ModEventsInRange(startBeat, startBeat+0.25, ilysa.FilterRGBLight(light), func(ctx ilysa.TimeContext, event ilysa.Event) {
	//	fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCirc)
	//})
	//p.ModEventsInRange(startBeat+0.25, startBeat+0.40, ilysa.FilterRGBLight(light), func(ctx ilysa.TimeContext, event ilysa.Event) {
	//	fx.RGBAlphaBlend(ctx, event, 1, 0.6, ease.OutCirc)
	//})
	//p.ModEventsInRange(startBeat+0.40, startBeat+0.50, ilysa.FilterRGBLight(light), func(ctx ilysa.TimeContext, event ilysa.Event) {
	//	fx.RGBAlphaBlend(ctx, event, 0.6, 3, ease.InExpo)
	//})
	//p.ModEventsInRange(startBeat+0.5, startBeat+0.95, ilysa.FilterRGBLight(light), func(ctx ilysa.TimeContext, event ilysa.Event) {
	//	fx.RGBAlphaBlend(ctx, event, 3, 0, ease.OutCirc)
	//})

	//p.EventForBeat(startBeat, func(ctx ilysa.TimeContext) {
	//	ctx.NewZoomEvent()
	//	ctx.NewPreciseRotationEvent(
	//		ilysa.WithRotation(45),
	//		ilysa.WithStep(12),
	//		ilysa.WithSpeed(3),
	//		ilysa.WithProp(2),
	//	)
	//})

	p.EventsForSequence(startBeat+1, []float64{0, 1}, func(ctx ilysa.SequenceContext) {
		ctx.NewZoomEvent()
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45),
			ilysa.WithStep(12),
			ilysa.WithSpeed(3),
			ilysa.WithProp(2),
			ilysa.WithDirection(chroma.Clockwise),
		)

		seqOrdinal := ctx.Ordinal()
		ctx.EventsForRange(ctx.B(), ctx.B()+0.3, 30, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(sl.Index(seqOrdinal), func(ctx ilysa.TimeLightContext) {
				events := fx.ColorSweep(ctx, 0.8, gradient.Rainbow)
				fx.Ripple(ctx, events, 0.65,
					fx.WithAlphaBlend(0, 0.5, 0, 2, ease.InCubic),
					fx.WithAlphaBlend(0.5, 1.0, 2, 0, ease.OutCirc),
					//fx.WithAlphaBlend(0.4, 1.0, 0, 0, ease.OutCirc),
				)
			})
		})
	})

	// tsuketa
	{
		sl := ilysa.TransformLight(backLasers,
			ilysa.ToLightTransformer(ilysa.Fan(2)),
			ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		).(ilysa.SequenceLight)

		sl = ilysa.NewSequenceLight(
			sl.Index(1),
			sl.Index(0),
			sl.Index(2),
		)

		p.EventsForSequence(startBeat+3, []float64{0, 0.5, 0.75}, func(ctx ilysa.SequenceContext) {
			ctx.NewZoomEvent()
			ctx.NewPreciseRotationEvent(
				ilysa.WithRotation(22.5),
				ilysa.WithStep(12),
				ilysa.WithSpeed(0.5),
				ilysa.WithProp(2),
				ilysa.WithDirection(chroma.CounterClockwise),
			)
			seqOrdinal := ctx.Ordinal()
			ctx.EventsForRange(ctx.B(), ctx.B()+0.6, 30, ease.Linear, func(ctx ilysa.TimeContext) {
				ctx.WithLight(sl.Index(seqOrdinal), func(ctx ilysa.TimeLightContext) {
					events := fx.ColorSweep(ctx, 0.8, gradient.Rainbow)
					fx.Ripple(ctx, events, 0.65,
						fx.WithAlphaBlend(0, 0.5, 0, 2, ease.InCubic),
						fx.WithAlphaBlend(0.5, 1.0, 2, 0, ease.OutCirc),
					)
				})
			})
		})
	}
}