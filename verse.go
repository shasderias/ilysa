package main

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ease"
	"ilysa/pkg/gen"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/ilysa/fx"
	"ilysa/pkg/util"
)

type Verse struct {
	ilysa.BareContext
	project *ilysa.Project
	offset  float64
}

func NewVerse(p *ilysa.Project, offset float64) Verse {
	ctx := p.WithBeatOffset(offset)
	return Verse{
		BareContext: ctx,
		project:     p,
		offset:      offset,
	}
}

func (p Verse) Play1a() {
	p.EventForBeat(-0.01, func(ctx ilysa.TimingContext) {
		gen.OffAll(ctx)
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithSpeed(1.5))
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithSpeed(1.5))
	})

	p.Rhythm(0)
	p.Rhythm(4)
	p.Rhythm(8)
	p.Rhythm(12)
	p.Rhythm(16)
	p.Rhythm(20)
	p.Rhythm(24)

	p.PianoBackstep(7)
	p.PianoBackstep(15)
	p.PianoBackstep(23)
	p.RinPun(27)
}

func (p Verse) Rhythm(startBeat float64) {
	var (
		kickDrumLight = ilysa.NewSequenceLight(
			p.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers),
			p.NewBasicLight(beatsaber.EventTypeRightRotatingLasers),
		)
		kickDrumSequence = []float64{0, 2.5}
		kickDrumColors   = magnetColors
	)
	p.EventsForSequence(startBeat, kickDrumSequence, func(ctx ilysa.SequenceContext) {
		ctx.UseLight(kickDrumLight, func(ctx ilysa.SequenceContextWithLight) {
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
	p.EventForBeat(rippleStart, func(ctx ilysa.TimingContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(22.5),
			ilysa.WithSpeed(2),
			ilysa.WithProp(0.3),
		)
	})

	p.EventsForRange(rippleStart, rippleEnd, 30, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(rippleLights, func(ctx ilysa.TimingContextWithLight) {
			events := fx.ColorSweep(ctx, 1.4, grad)
			events.Mod(ilysa.WithAlpha(1.5))
			for _, ee := range *events {
				ee.ShiftBeat(ctx.LightIDT() * rippleStep)
			}
			switch {
			case ctx.T() <= 0.5:
				alphaScale := util.ScaleToUnitInterval(0, 0.5)
				events.Mod(ilysa.WithAlpha(events.GetAlpha() * ease.InOutQuart(alphaScale(ctx.T()))))
			case ctx.T() > 0.8:
				alphaScale := util.ScaleToUnitInterval(0.8, 1)
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
//		ctx.UseLight(light, lightid.GroupDivide(3), func(ctx ilysa.ContextWithLight) {
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
	backLasersSplit := backLasers.LightIDTransformSequence(ilysa.Divide(2))
	backLasersSplitSingle := backLasersSplit.ApplyLightIDTransform(ilysa.DivideSingle)

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

		ctx.EventsForRange(ctx.B(), ctx.B()+0.25, 6, ease.Linear, func(ctx ilysa.TimingContext) {
			ctx.UseLight(backLasersSplitSingle, func(ctx ilysa.TimingContextWithLight) {
				fx.Gradient(ctx, magnetGradient)
			})
		})

		//ctx.ModEventsInRange(ctx.b, ctx.b+0.10, ilysa.FilterRGBLight(backLight), func(ctx ilysa.Context, event ilysa.Event) {
		//	fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCirc)
		//})
		ctx.ModEventsInRange(ctx.B(), ctx.B()+0.25, ilysa.FilterRGBLight(backLasersSplitSingle), func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
		})

		//ctx.UseLight(backLight, lightid.GroupDivide(2), func(ctx ilysa.ContextWithLight) {
		//	if ctx.ordinal != ctx.LightIDOrdinal {
		//		return
		//	}
		//
		//	fx.BiasedColorSweep(ctx, 2, 0.8, gradient.Rainbow)
		//})
	})
}

func (p Verse) RinPun(startBeat float64) {
	light := p.NewBasicLight(beatsaber.EventTypeBackLasers).LightIDTransform(ilysa.DivideSingle)

	p.EventsForRange(startBeat, startBeat+0.95, 30, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			fx.BiasedColorSweep(ctx, 1.2, gradient.Rainbow)
		})
	})

	p.ModEventsInRange(startBeat, startBeat+0.25, ilysa.FilterRGBLight(light), func(ctx ilysa.TimingContext, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCirc)
	})
	p.ModEventsInRange(startBeat+0.25, startBeat+0.40, ilysa.FilterRGBLight(light), func(ctx ilysa.TimingContext, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 1, 0.6, ease.OutCirc)
	})
	p.ModEventsInRange(startBeat+0.40, startBeat+0.50, ilysa.FilterRGBLight(light), func(ctx ilysa.TimingContext, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 0.6, 3, ease.InExpo)
	})
	p.ModEventsInRange(startBeat+0.5, startBeat+0.95, ilysa.FilterRGBLight(light), func(ctx ilysa.TimingContext, event ilysa.Event) {
		fx.RGBAlphaBlend(ctx, event, 3, 0, ease.OutCirc)
	})

	p.EventForBeat(startBeat, func(ctx ilysa.TimingContext) {
		ctx.NewZoomEvent()
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(360),
			ilysa.WithStep(12),
			ilysa.WithSpeed(3),
			ilysa.WithProp(2),
		)
	})

	seqLight := p.NewBasicLight(beatsaber.EventTypeBackLasers).LightIDTransformSequence(ilysa.Fan(2))
	seqLight = seqLight.ApplyLightIDTransform(ilysa.DivideSingle)

	backLasers := ilysa.FanBasicLight(2)(ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p.project.ActiveDifficultyProfile()))

	p.EventsForSequence(startBeat+1, []float64{0, 1}, func(ctx ilysa.SequenceContext) {
		seqOrdinal := ctx.Ordinal()
		ctx.EventsForRange(ctx.B(), ctx.B()+1, 30, ease.Linear, func(ctx ilysa.TimingContext) {
			ctx.UseLight(backLasers[seqOrdinal].LightIDTransform(ilysa.DivideSingle), func(ctx ilysa.TimingContextWithLight) {
				events := fx.ColorSweep(ctx, 1.2, gradient.Rainbow)
				fx.Ripple(ctx, events, 0.3,
					fx.WithAlphaBlend(0, 0.3, 0, 2, ease.InCubic),
					fx.WithAlphaBlend(0.5, 1, 2, 0, ease.OutCirc),
				)
			})
		})
	})
}
