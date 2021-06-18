package main

import (
	"math/rand"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/colorful"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ease"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/ilysa/fx"
)

type Intro struct {
	*ilysa.Project
	startBeat float64
}

func (p Intro) Play() {
	p.PianoDoubles(16)
	p.LeadinDrums(18.25)
	p.BassTwang(18.5)
	p.StartSplash(20)
	p.Rhythm(20, 23)
	p.Rhythm(24, 27)
	p.Rhythm(28, 31)
	p.Melody1(20)
	p.Melody2(23.25, false)
	p.Melody1(24)
	p.Melody2(27.25, true)
	p.Melody3(28)
	p.Chorus(32)
	p.PianoRoll(36.5, 6)
	p.Trill(38.5)
	p.Climb(39.5)
	p.Trill(42.5)
	p.Fall(43.25)
	p.Trill(44.5)
	p.Bridge(45.0)
	p.Rhythm(46.0, 50)
	p.Outro(46.5)
	p.OutroSplash(50.0)
}

func (p Intro) PianoDoubles(startBeat float64) {
	set := colorful.NewSet(magnetPurple, magnetPink, magnetWhite, colorful.Black)

	lights := []beatsaber.EventType{
		beatsaber.EventTypeCenterLights,
		beatsaber.EventTypeRingLights,
		beatsaber.EventTypeBackLasers,
	}

	cl := ilysa.NewCombinedLight()

	for _, light := range lights {
		cl.Add(p.NewBasicLight(light).Transform(ilysa.DivideSingle))
	}

	values := beatsaber.NewEventValueSet(
		beatsaber.EventValueLightRedOn,
		beatsaber.EventValueLightBlueOn,
		beatsaber.EventValueLightRedOn,
		beatsaber.EventValueLightBlueOn,
		beatsaber.EventValueLightOff,
	)

	p.EventsForSequence(startBeat, []float64{0, 0.75, 1.25, 1.75, 2.25}, func(ctx ilysa.SequenceContext) {
		grad := gradient.Table{
			{set.Pick(ctx.Ordinal()), 0.0},
			{set.Pick(ctx.Ordinal() + 1), 1.0},
		}

		ctx.UseLight(cl, func(ctx ilysa.SequenceContextWithLight) {
			if !ctx.Last() {
				e := fx.Gradient(ctx, grad)
				e.SetValue(values[ctx.Ordinal()])
			} else {
				ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			}
		})
	})
}

func (p Intro) LeadinDrums(startBeat float64) {
	light := ilysa.NewSequenceLight(
		p.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers).
			Transform(ilysa.Fan(2)),
		p.NewBasicLight(beatsaber.EventTypeRightRotatingLasers).
			Transform(ilysa.Fan(2)),
	)

	p.EventsForSequence(startBeat, []float64{0, 0.25, 0.75, 1, 1.5}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithValue(1),
			ilysa.WithLockPosition(false), ilysa.WithSpeed(0), ilysa.WithDirection(chroma.Clockwise),
		)

		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithValue(1),
			ilysa.WithLockPosition(false), ilysa.WithSpeed(0), ilysa.WithDirection(chroma.Clockwise),
		)

		ctx.UseLight(light, func(ctx ilysa.SequenceContextWithLight) {
			if ctx.Ordinal() < 4 {
				ctx.NewRGBLightingEvent(ilysa.WithColor(magnetColors.Pick(ctx.Ordinal())))
			} else {
				ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			}

			if ctx.Ordinal() > 0 {
				e := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
				e.ShiftBeat(ctx.NextBOffset())
			}
		})
	})
}

func (p Intro) BassTwang(startBeat float64) {
	const (
		steps      = 60
		intensity  = 0.65
		sweepSpeed = 1.1
	)

	var (
		endBeat    = startBeat + 1.495
		midPoint   = (endBeat - startBeat) / 2.0
		backLasers = p.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.DivideSingle)
		grad       = magnetGradient
	)

	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(backLasers, func(ctx ilysa.TimingContextWithLight) {
			e := fx.BiasedColorSweep(ctx, sweepSpeed, grad)
			e.Mod(ilysa.WithAlpha(intensity))
		})
	})

	p.ModEventsInRange(startBeat, startBeat+midPoint-0.001, ilysa.FilterRGBLight(backLasers),
		func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
		})

	p.ModEventsInRange(startBeat+midPoint, endBeat, ilysa.FilterRGBLight(backLasers),
		func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutBounce)
		})
}

func (p Intro) StartSplash(startBeat float64) {
	p.EventForBeat(startBeat, func(ctx ilysa.TimingContext) {
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithValue(8))
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithValue(8))

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
			ilysa.WithColor(sukoyaPink))
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightRedFlash),
			ilysa.WithColor(shirayukiPurple))
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeCenterLights), ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
			ilysa.WithColor(magnetPurple))

		ctx.NewZoomEvent()
	})
}

func (p Intro) Rhythm(startBeat, endBeat float64) {
	var (
		steps = int(endBeat-startBeat) + 1
	)

	p.EventsForBeats(startBeat, 1.0, steps, func(ctx ilysa.TimingContext) {
		set := magnetColors

		switch {
		case ctx.Ordinal() == 0:
			ctx.NewPreciseRotationEvent(
				ilysa.WithRotation(180),
				ilysa.WithStep(0),
				ilysa.WithProp(1),
				ilysa.WithSpeed(24),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRingLights),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(set.Pick(ctx.Ordinal())),
			)

		case ctx.Ordinal()%2 == 1:
			ctx.NewPreciseRotationEvent(
				ilysa.WithRotation(12.5),
				ilysa.WithStep(10+3*ctx.T()),
				ilysa.WithProp(20),
				ilysa.WithSpeed(20),
				ilysa.WithDirection(chroma.CounterClockwise),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(magnetPurple),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
				ilysa.WithColor(magnetPink),
			)

		case ctx.Ordinal()%2 == 0:
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRingLights),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
				ilysa.WithColor(set.Pick(ctx.Ordinal())),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
				ilysa.WithColor(magnetPink),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(magnetPurple),
			)
		}
	})
}

//
////
////func IntroRhythmSplash(p *ilysa.Project, startBeat, endBeat float64) {
////	var (
////		steps = int(endBeat - startBeat)
////	)
////	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Timing) {
////		set := magnetColors
////
////		br := ctx.NewRGBLightingEvent(beatsaber.EventTypeRingLights, beatsaber.EventValueLightRedFlash)
////		br.SetColor(set.Index(ctx.Ordinal()))
////
////	})
////}
////
//
func (p Intro) Melody1(startBeat float64) {
	var (
		sequence = []float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75}
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Divide(3))
	)

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.UseLight(light, func(ctx ilysa.SequenceContextWithLight) {
			//if ctx.Ordinal()%ctx.LightIDLen() != ctx.LightIDOrdinal() {
			//	return
			//}
			//fx.Gradient(ctx, beatsaber.EventValueLightBlueOn, magnetGradient)
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightBlueOn),
				ilysa.WithColor(magnetPurple),
			)

			e := ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightOff),
			)
			if !ctx.Last() {
				e.ShiftBeat(ctx.NextBOffset())
			} else {
				e.ShiftBeat(0.249)
			}
			//ilysa.WithLightID(ctx.LightIDCur()),

			//if ctx.Ordinal() > 0 {
			//	ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff).
			//		SetTypeID(ctx.PreLightID)
			//}
		})
	})

	//p.EventForBeat(startBeat+2.999, func(ctx ilysa.TimingContext) {
	//	ctx.NewRGBLightingEvent().SetType(light).SetValue(beatsaber.EventValueLightOff)
	//})
}

func (p Intro) Melody2(startBeat float64, reverseZoom bool) {
	var (
		sequence = []float64{0, 0.25, 0.50}
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Fan(3))
	)

	p.EventForBeat(startBeat-0.01, func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
		})
	})

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.UseLight(light, func(ctx ilysa.SequenceContextWithLight) {
			ctx.NewRGBLightingEvent(
				ilysa.WithColor(magnetPink),
			)

			oe := ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightOff),
			)
			if !ctx.Last() {
				oe.ShiftBeat(ctx.NextBOffset())
			} else {
				oe.ShiftBeat(0.245)
			}
		})

		ze := ctx.NewPreciseZoomEvent()
		if reverseZoom {
			ze.Mod(ilysa.WithStep(0.3))
		} else {
			ze.Mod(ilysa.WithStep(-0.3))
		}
	})
}

func (p Intro) Melody3(startBeat float64) {
	var (
		sequence = []float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75, 3.00, 3.25, 3.50}
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Divide(3))
	)

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.UseLight(light, func(ctx ilysa.SequenceContextWithLight) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightBlueOn),
				ilysa.WithColor(magnetPurple),
			)

			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			if !ctx.Last() {
				oe.ShiftBeat(ctx.NextBOffset())
			} else {
				oe.ShiftBeat(0.45)
			}
		})
	})
}

func (p Intro) Chorus(startBeat float64) {
	var (
		sequence  = []float64{0, 1, 2, 2.75, 3.5, 4}
		light     = p.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.DivideSingle)
		colorGrad = allColorsGradient
	)

	p.EventForBeat(startBeat, func(ctx ilysa.TimingContext) {
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightOff))

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightOff))
	})

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseZoomEvent(ilysa.WithStep(0.2))

		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithValue(1),
			ilysa.WithLockPosition(false),
			ilysa.WithSpeed(0),
			ilysa.WithDirection(chroma.Clockwise),
		)

		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithValue(1),
			ilysa.WithLockPosition(false),
			ilysa.WithSpeed(0),
			ilysa.WithDirection(chroma.CounterClockwise),
		)

		re := ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45),
			ilysa.WithStep(5+(1.5*float64(ctx.Ordinal()))),
			ilysa.WithProp(20),
			ilysa.WithSpeed(4),
		)

		if ctx.Ordinal()%2 == 0 {
			re.Mod(ilysa.WithDirection(chroma.Clockwise))
		} else {
			re.Mod(ilysa.WithDirection(chroma.CounterClockwise))
		}

		if ctx.Ordinal() == 5 {
			re.Mod(ilysa.WithRotation(360))
		}

		grad := append(gradient.Table{}, colorGrad...)
		rand.Shuffle(len(colorGrad), func(i, j int) {
			grad[i].Col, grad[j].Col = grad[j].Col, grad[i].Col
		})

		ctx.UseLight(light, func(ctx ilysa.SequenceContextWithLight) {
			e := fx.Gradient(ctx, grad)
			e.ShiftBeat(1.0 * float64(ctx.Ordinal()) / 32)
		})

		//for i := 1; i <= LightIDMax; i++ {
		//	gradientPos := util.Scale(1, float64(LightIDMax), 0, 1)
		//	color := grad.GetInterpolatedColorFor(gradientPos(float64(i)))
		//
		//	e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
		//	e.SetSingleLightID(i)
		//	e.SetColor(color)
		//	e.Beat += 1.0 / 64.0
		//}
	})
}

func (p Intro) PianoRoll(startBeat float64, count int) {
	var (
		light = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Divide(count))
	)

	p.EventsForBeats(startBeat, 0.25, count, func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff)
		})
	})
}
func (p Intro) Trill(startBeat float64) {
	var (
		backLasers = p.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.DivideSingle)
		ringLasers = p.NewBasicLight(beatsaber.EventTypeRingLights).Transform(ilysa.DivideSingle)
		step       = 0.125
		count      = 5
		ratio      = 0.666
		//lightCount = int(ratio * float64(LightIDMax))
	)

	p.EventsForBeats(startBeat, step, count, func(ctx ilysa.TimingContext) {
		ctx.UseLight(backLasers, func(ctx ilysa.TimingContextWithLight) {
			if rand.Float64() > ratio {
				return
			}

			ctx.NewRGBLightingEvent(ilysa.WithColor(
				allColorsGradient.GetInterpolatedColorFor(rand.Float64()),
			))

			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			oe.ShiftBeat(step / 2)
		})
		//for i := 0; i < lightCount; i++ {
		//	e := ctx.NewRGBLightingEvent(backLasers, beatsaber.EventValueLightRedOn)
		//	e.SetSingleLightID(rand.Intn(LightIDMax) + 1)
		//	e.SetColor(allColorsGradient.GetInterpolatedColorFor(rand.Float64()))
		//
		//	oe := ctx.NewRGBLightingEvent(backLasers, beatsaber.EventValueLightOff)
		//	oe.Beat += step / 2
		//
		//	if !ctx.Last {
		//		continue
		//	}
		//
		//}
	})
	p.EventsForRange(startBeat+0.5, startBeat+0.5+1.2, 30, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(ringLasers, func(ctx ilysa.TimingContextWithLight) {
			fx.ColorSweep(ctx, 0.4, gradient.Rainbow)
		})
	})

	p.ModEventsInRange(startBeat+0.5, startBeat+0.5+0.6-0.001, ilysa.FilterRGBLight(ringLasers),
		func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
		})

	p.ModEventsInRange(startBeat+0.5+0.6, startBeat+0.5+1.2, ilysa.FilterRGBLight(ringLasers),
		func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
		})
}

func (p Intro) Climb(startBeat float64) {
	var (
		light           = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Divide(7))
		step            = 0.25
		lightIDSequence = []int{6, 7, 5, 8, 4, 9, 3, 10, 2, 11, 1, 12}
		count           = len(lightIDSequence)
		//LightIDMax      = p.ActiveDifficultyProfile().LightIDMax(light)
		//lightIDs        = light2.FromInterval(1, LightIDMax)
		//lightIDSets     = light2.DivideIntoGroupsOf(lightIDs, LightIDMax/2)

		backGrad = gradient.Table{
			{magnetPink, 0.0},
			{magnetWhite, 1.0},
		}
		sideGrad = gradient.Table{
			{magnetWhite, 0.0},
			{magnetPurple, 1.0},
		}
	)

	p.EventForBeat(startBeat, func(ctx ilysa.TimingContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(360),
			ilysa.WithStep(15),
			ilysa.WithSpeed(1.3),
			ilysa.WithProp(13),
		)

		ctx.NewZoomEvent()
	})

	p.EventsForBeats(startBeat, step, count, func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			ctx.NewRGBLightingEvent(
				ilysa.WithColor(backGrad.GetInterpolatedColorFor(ctx.T())),
			)
		})

		switch {
		case ctx.Last():
			const exitValue = 3
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
				ilysa.WithColor(magnetPurple),
			)

			ctx.NewPreciseRotationSpeedEvent(
				ilysa.WithDirectionalLaser(ilysa.LeftLaser),
				ilysa.WithValue(exitValue),
				ilysa.WithLockPosition(true),
				ilysa.WithSpeed(exitValue),
				ilysa.WithDirection(chroma.CounterClockwise),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(magnetPurple),
			)

			ctx.NewPreciseRotationSpeedEvent(
				ilysa.WithDirectionalLaser(ilysa.RightLaser),
				ilysa.WithValue(exitValue),
				ilysa.WithLockPosition(true),
				ilysa.WithSpeed(exitValue),
				ilysa.WithDirection(chroma.Clockwise),
			)

		case ctx.Ordinal()%2 == 0:
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
				ilysa.WithColor(sideGrad.GetInterpolatedColorFor(ctx.T())),
			)

			ctx.NewPreciseRotationSpeedEvent(
				ilysa.WithDirectionalLaser(ilysa.LeftLaser),
				ilysa.WithValue(beatsaber.EventValue(ctx.Ordinal())),
				ilysa.WithLockPosition(true),
				ilysa.WithSpeed(float64(ctx.Ordinal())),
				ilysa.WithDirection(chroma.Clockwise),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightOff),
			)

		case ctx.Ordinal()%2 == 1:
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightOff),
			)

			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFlash),
				ilysa.WithColor(sideGrad.GetInterpolatedColorFor(ctx.T())),
			)

			ctx.NewPreciseRotationSpeedEvent(
				ilysa.WithDirectionalLaser(ilysa.RightLaser),
				ilysa.WithValue(beatsaber.EventValue(ctx.Ordinal())),
				ilysa.WithLockPosition(false),
				ilysa.WithSpeed(float64(ctx.Ordinal())),
				ilysa.WithDirection(chroma.CounterClockwise),
			)
		}
	})
}

func (p Intro) Fall(startBeat float64) {
	var (
		//lightIDs    = light2.FromInterval(1, LightIDMax)
		//lightIDSets = light2.DivideIntoGroupsOf(lightIDs, count)
		step     = 0.25
		count    = 4
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Divide(count))
		colorSet = colorful.NewSet(magnetPurple, magnetPink)
		values   = beatsaber.NewEventValueSet(
			beatsaber.EventValueLightRedOn,
			beatsaber.EventValueLightOff,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedFlash,
		)
	)

	p.EventsForBeats(startBeat, step, count, func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(values.Pick(ctx.Ordinal())),
				ilysa.WithColor(colorSet.Next()),
			)
		})
	})
}

func (p Intro) Bridge(startBeat float64) {
	p.EventForBeat(startBeat, func(ctx ilysa.TimingContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(180),
			ilysa.WithStep(12.5),
			ilysa.WithDirection(chroma.CounterClockwise),
			ilysa.WithSpeed(3),
			ilysa.WithProp(5),
			ilysa.WithCounterSpin(true),
		)
	})

	p.EventsForRange(startBeat, startBeat+1, 30, ease.OutCubic, func(ctx ilysa.TimingContext) {
		if !ctx.Last() {
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeBackLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueOn),
				ilysa.WithColor(magnetPurple),
				ilysa.WithAlpha(1-ctx.T()),
			)
		} else {
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeBackLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedOn),
				ilysa.WithColor(magnetWhite),
				ilysa.WithAlpha(8),
			)
		}
	})
}

func (p Intro) Outro(startBeat float64) {
	var (
		sequence = []float64{0, 0.25, 0.50, 1.0, 1.25, 1.50, 2.0, 2.25, 2.50, 2.75, 3.25}
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Fan(len(sequence)))
	)

	p.EventForBeat(startBeat-0.01, func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			ctx.NewRGBLightingEvent(
				ilysa.WithColor(allColorsGradient.GetInterpolatedColorFor(ctx.LightIDT())),
			)
		})
	})

	p.EventsForSequence(startBeat, sequence[:len(sequence)-1], func(ctx ilysa.SequenceContext) {
		ctx.UseLight(light, func(ctx ilysa.SequenceContextWithLight) {
			ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
		})
	})

	p.EventForBeat(startBeat+sequence[len(sequence)-1], func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
		})
	})
}

func (p Intro) OutroSplash(startBeat float64) {
	var (
		intensity  = 2.0
		sweepSpeed = 0.6
		grad       = gradient.Rainbow
		sequence   = []float64{0, 0.75, 1.5}
		leftLaser  = p.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers).Transform(ilysa.DivideSingle)
		rightLaser = p.NewBasicLight(beatsaber.EventTypeRightRotatingLasers).Transform(ilysa.DivideSingle)
		backLaser  = p.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.DivideSingle)
	)

	p.EventsForRange(startBeat, startBeat+2, 60, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(backLaser, func(ctx ilysa.TimingContextWithLight) {
			e := fx.BiasedColorSweep(ctx, sweepSpeed, grad)
			e.Mod(ilysa.WithAlpha(intensity))
		})
	})

	p.ModEventsInRange(startBeat, startBeat+1, ilysa.FilterRGBLight(backLaser),
		func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
		})

	p.ModEventsInRange(startBeat+1, startBeat+2, ilysa.FilterRGBLight(backLaser),
		func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
		})

	const (
		splashDuration = 0.495
	)

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		g := append(gradient.Table{}, grad...)
		rand.Shuffle(len(g), func(i, j int) {
			g[i].Col, g[j].Col = g[j].Col, g[i].Col
		})
		p.EventsForRange(ctx.B(), ctx.B()+splashDuration, 15, ease.Linear, func(ctx ilysa.TimingContext) {
			ctx.UseLight(leftLaser, func(ctx ilysa.TimingContextWithLight) {
				e := fx.Gradient(ctx, g)
				e.Mod(ilysa.WithAlpha(float64(ctx.Ordinal()) * 0.5))
			})
			ctx.UseLight(rightLaser, func(ctx ilysa.TimingContextWithLight) {
				e := fx.Gradient(ctx, g)
				e.Mod(ilysa.WithAlpha(float64(ctx.Ordinal()) * 0.5))
			})
		})

		p.ModEventsInRange(ctx.B(), ctx.B()+splashDuration, ilysa.FilterRGBLights(), func(ctx ilysa.TimingContext, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.InCirc)
		})
	})
}
