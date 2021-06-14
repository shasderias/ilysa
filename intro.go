package main

import (
	"math/rand"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/chroma/lightid"
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
	p.EventsForSequence(startBeat, []float64{0, 0.75, 1.25, 1.75, 2.25}, func(ctx ilysa.Context) {
		set := colorful.NewSet(magnetPurple, magnetPink, magnetWhite, colorful.Black)

		lights := []beatsaber.EventType{
			beatsaber.EventTypeCenterLights,
			beatsaber.EventTypeRingLights,
			beatsaber.EventTypeBackLasers,
		}

		values := beatsaber.NewEventValueSet(
			beatsaber.EventValueLightRedOn,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightOff,
		)

		for _, light := range lights {
			grad := gradient.Table{
				{set.Pick(ctx.Ordinal), 0.0},
				{set.Pick(ctx.Ordinal + 1), 1.0},
			}

			ctx.RangeLightIDs(p.NewBasicLight(light), lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
				fx.Gradient(ctx, values.Pick(ctx.Ordinal), grad)
			})
		}
	})
}

func (p Intro) LeadinDrums(startBeat float64) {
	p.EventsForSequence(startBeat, []float64{0, 0.25, 0.75, 1, 1.5}, func(ctx ilysa.Context) {
		var (
			leftLaser            = beatsaber.EventTypeLeftRotatingLasers
			rightLaser           = beatsaber.EventTypeRightRotatingLasers
			leftLaserMaxLightID  = ctx.ActiveDifficultyProfile().MaxLightID(leftLaser)
			rightLaserMaxLightID = ctx.ActiveDifficultyProfile().MaxLightID(rightLaser)
			ordinal              = ctx.Ordinal
		)

		lightIDGroups := [][]int{
			lightid.EveryNthLightID(1, leftLaserMaxLightID, 2, 0),
			lightid.EveryNthLightID(1, leftLaserMaxLightID, 2, 1),
			lightid.EveryNthLightID(1, rightLaserMaxLightID, 2, 0),
			lightid.EveryNthLightID(1, rightLaserMaxLightID, 2, 1),
		}

		types := beatsaber.NewEventTypeSet(leftLaser, rightLaser)

		values := beatsaber.NewEventValueSet(
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedOn,
		)

		ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.LeftLaser).SetValue(1).
			SetLockPosition(false).SetSpeed(0).SetDirection(chroma.Clockwise)

		ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.RightLaser).SetValue(1).
			SetLockPosition(false).SetSpeed(0).SetDirection(chroma.CounterClockwise)

		if ordinal < 4 {
			ctx.NewRGBLightingEvent().
				SetLight(types.Pick(ctx.Ordinal)).
				SetValue(values.Pick(ctx.Ordinal)).
				SetLightID(lightIDGroups[ordinal]).
				SetColor(magnetColors.Pick(ordinal))
		} else {
			ctx.NewRGBLightingEvent().SetLight(leftLaser).SetValue(beatsaber.EventValueLightOff)
			ctx.NewRGBLightingEvent().SetLight(rightLaser).SetValue(beatsaber.EventValueLightOff)
		}

		if ordinal > 0 {
			ctx.NewRGBLightingEvent().SetLight(types.Pick(ordinal - 1)).SetValue(beatsaber.EventValueLightOff)
		}
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
		backLasers = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		grad       = magnetGradient
	)

	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(backLasers, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
			e := fx.BiasedColorSweep(ctx, intensity, sweepSpeed, grad)
			e.SetAlpha(intensity)
		})
	})

	p.ModEventsInRange(startBeat, startBeat+midPoint-0.001, ilysa.FilterRGBLight(backLasers),
		func(ctx ilysa.Context, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
		})

	p.ModEventsInRange(startBeat+midPoint, endBeat, ilysa.FilterRGBLight(backLasers),
		func(ctx ilysa.Context, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutBounce)
		})

	//p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Context) {
	//	for i := 1; i <= maxLightID; i++ {
	//		e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightBlueOn)
	//		e.SetSingleLightID(i)
	//		e.SetColor(magnetGradient.GetInterpolatedColorFor(
	//			sin(ctx.Pos*3 + (float64(i)/float64(maxLightID))*pi + 4),
	//		))
	//		e.SetAlpha(3)
	//	}
	//})
	//
	//fadeScale := util.Scale(startBeat, endBeat, 0, 1)
	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
	//modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutBounce)
}

func (p Intro) StartSplash(startBeat float64) {
	p.EventForBeat(startBeat, func(ctx ilysa.Context) {
		ctx.NewRotationSpeedEvent(ilysa.LeftLaser, 8)
		ctx.NewRotationSpeedEvent(ilysa.RightLaser, 8)

		ctx.NewRGBLightingEvent().
			SetLight(beatsaber.EventTypeLeftRotatingLasers).
			SetValue(beatsaber.EventValueLightRedFlash).
			SetColor(sukoyaPink)

		ctx.NewRGBLightingEvent().
			SetLight(beatsaber.EventTypeRightRotatingLasers).
			SetValue(beatsaber.EventValueLightBlueFlash).
			SetColor(shirayukiPurple)

		ctx.NewRGBLightingEvent().
			SetLight(beatsaber.EventTypeCenterLights).
			SetValue(beatsaber.EventValueLightRedFlash).
			SetColor(magnetPurple)

		ctx.NewZoomEvent()
	})
}

func (p Intro) Rhythm(startBeat, endBeat float64) {
	var (
		steps = int(endBeat-startBeat) + 1
	)

	p.EventsForBeats(startBeat, 1.0, steps, func(ctx ilysa.Context) {
		set := magnetColors

		switch {
		case ctx.Ordinal == 0:
			ctx.NewPreciseRotationEvent().
				SetRotation(180).
				SetStep(0).
				SetProp(1).
				SetSpeed(24)

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeRingLights).
				SetValue(beatsaber.EventValueLightRedFade).
				SetColor(set.Pick(ctx.Ordinal))

		case ctx.Ordinal%2 == 1:
			ctx.NewPreciseRotationEvent().
				SetRotation(12.5).
				SetStep(10 + 3*ctx.Pos).
				SetProp(20).
				SetSpeed(20).
				SetDirection(chroma.CounterClockwise)

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeLeftRotatingLasers).
				SetValue(beatsaber.EventValueLightRedFade).
				SetColor(magnetPurple)

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeRightRotatingLasers).
				SetValue(beatsaber.EventValueLightBlueFade).
				SetColor(magnetPink)

		case ctx.Ordinal%2 == 0:
			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeRingLights).
				SetValue(beatsaber.EventValueLightBlueFade).
				SetColor(set.Pick(ctx.Ordinal))

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeLeftRotatingLasers).
				SetValue(beatsaber.EventValueLightBlueFade).
				SetColor(magnetPink)

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeRightRotatingLasers).
				SetValue(beatsaber.EventValueLightRedFade).
				SetColor(magnetPurple)
		}
	})
}

//
//func IntroRhythmSplash(p *ilysa.Project, startBeat, endBeat float64) {
//	var (
//		steps = int(endBeat - startBeat)
//	)
//	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Context) {
//		set := magnetColors
//
//		br := ctx.NewRGBLightingEvent(beatsaber.EventTypeRingLights, beatsaber.EventValueLightRedFlash)
//		br.SetColor(set.Pick(ctx.Ordinal))
//
//	})
//}
//

func (p Intro) Melody1(startBeat float64) {
	var (
		sequence    = []float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75}
		light       = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		lightGroups = 3
	)

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.GroupDivide(lightGroups), func(ctx ilysa.RangeLightIDContext) {
			if ctx.Ordinal%ctx.LightIDSetLen != ctx.LightIDOrdinal {
				return
			}
			//fx.Gradient(ctx, beatsaber.EventValueLightBlueOn, magnetGradient)
			ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightBlueOn).
				SetColor(magnetPurple).
				SetLightID(ctx.CurLightID)

			if ctx.Ordinal > 0 {
				ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff).
					SetLightID(ctx.PreLightID)
			}
		})
	})

	p.EventForBeat(startBeat+2.999, func(ctx ilysa.Context) {
		ctx.NewRGBLightingEvent().SetLight(light).SetValue(beatsaber.EventValueLightOff)
	})
}

func (p Intro) Melody2(startBeat float64, reverseZoom bool) {
	var (
		sequence    = []float64{0, 0.25, 0.50}
		light       = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		lightGroups = 3
	)

	p.EventForBeat(startBeat-0.001, func(ctx ilysa.Context) {
		ctx.NewRGBLightingEvent().SetLight(light).SetValue(beatsaber.EventValueLightOff)
	})

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.Fan(lightGroups), func(ctx ilysa.RangeLightIDContext) {
			if ctx.Ordinal%ctx.LightIDSetLen != ctx.LightIDOrdinal {
				return
			}

			ctx.NewRGBLightingEvent().
				SetValue(beatsaber.EventValueLightBlueOn).
				SetColor(magnetPink)

			if ctx.Ordinal > 0 {
				ctx.NewRGBLightingEvent().
					SetValue(beatsaber.EventValueLightOff).
					SetLightID(ctx.PreLightID)
			}
		})
		//e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightBlueOn)
		//e.SetColor(magnetPink)
		//e.SetLightID(lightIDSet.Pick(ctx.Ordinal))

		ze := ctx.NewPreciseZoomEvent()
		if reverseZoom {
			ze.Step = 0.3
		} else {
			ze.Step = -0.3
		}

	})

	p.EventForBeat(startBeat+0.749, func(ctx ilysa.Context) {
		ctx.NewRGBLightingEvent().SetLight(light).SetValue(beatsaber.EventValueLightOff)
	})
}

func (p Intro) Melody3(startBeat float64) {
	var (
		sequence = []float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75, 3.00, 3.25, 3.50}
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers)
	)

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.GroupDivide(3), func(ctx ilysa.RangeLightIDContext) {
			if ctx.Ordinal%ctx.LightIDSetLen != ctx.LightIDOrdinal {
				return
			}

			ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightBlueOn).
				SetColor(magnetPurple)

			if ctx.Ordinal > 0 {
				ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff).
					SetLightID(ctx.PreLightID)
			}
		})
	})

	p.EventForBeat(startBeat+3.999, func(ctx ilysa.Context) {
		ctx.NewRGBLightingEvent().SetLight(light).SetValue(beatsaber.EventValueLightOff)
	})
}

func (p Intro) Chorus(startBeat float64) {
	var (
		sequence  = []float64{0, 1, 2, 2.75, 3.5, 4}
		light     = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		colorGrad = allColorsGradient
	)

	p.EventForBeat(startBeat, func(ctx ilysa.Context) {
		ctx.NewRGBLightingEvent().SetLight(beatsaber.EventTypeLeftRotatingLasers).
			SetValue(beatsaber.EventValueLightOff)
		ctx.NewRGBLightingEvent().SetLight(beatsaber.EventTypeRightRotatingLasers).
			SetValue(beatsaber.EventValueLightOff)
	})

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.Context) {
		ctx.NewPreciseZoomEvent().SetStep(0.2)

		ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.LeftLaser).SetValue(1).
			SetLockPosition(false).SetSpeed(0).SetDirection(chroma.Clockwise)

		ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.RightLaser).SetValue(1).
			SetLockPosition(false).SetSpeed(0).SetDirection(chroma.CounterClockwise)

		re := ctx.NewPreciseRotationEvent().
			SetRotation(45).
			SetStep(5 + (1.5 * float64(ctx.Ordinal))).
			SetProp(20).
			SetSpeed(4)

		if ctx.Ordinal%2 == 0 {
			re.Direction = chroma.Clockwise
		} else {
			re.Direction = chroma.CounterClockwise
		}

		if ctx.Ordinal == 5 {
			re.Rotation = 360
		}

		grad := append(gradient.Table{}, colorGrad...)
		rand.Shuffle(len(colorGrad), func(i, j int) {
			grad[i].Col, grad[j].Col = grad[j].Col, grad[i].Col
		})

		ctx.RangeLightIDs(light, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
			e := fx.Gradient(ctx, beatsaber.EventValueLightRedOn, grad)
			e.Beat += 1.0 * float64(ctx.Ordinal) / 32
		})

		//for i := 1; i <= maxLightID; i++ {
		//	gradientPos := util.Scale(1, float64(maxLightID), 0, 1)
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
		light = p.NewBasicLight(beatsaber.EventTypeBackLasers)
	)

	p.EventsForBeats(startBeat, 0.25, count, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.GroupDivide(count), func(ctx ilysa.RangeLightIDContext) {
			if ctx.Ordinal%ctx.LightIDSetLen != ctx.LightIDOrdinal {
				return
			}
			ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff)
		})
	})
}

func (p Intro) Trill(startBeat float64) {
	var (
		backLasers = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		ringLasers = p.NewBasicLight(beatsaber.EventTypeRingLights)
		step       = 0.125
		count      = 5
		ratio      = 0.666
		//lightCount = int(ratio * float64(maxLightID))
	)

	p.EventsForBeats(startBeat, step, count, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(backLasers, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
			if rand.Float64() > ratio {
				return
			}

			ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightRedOn).
				SetColor(allColorsGradient.GetInterpolatedColorFor(rand.Float64()))

			oe := ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff)
			oe.Beat += step / 2
		})
		//for i := 0; i < lightCount; i++ {
		//	e := ctx.NewRGBLightingEvent(backLasers, beatsaber.EventValueLightRedOn)
		//	e.SetSingleLightID(rand.Intn(maxLightID) + 1)
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
	p.EventsForRange(startBeat+0.5, startBeat+0.5+1.2, 30, ease.Linear, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(ringLasers, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
			fx.ColorSweep(ctx, 1, 0.4, gradient.Rainbow)
		})
	})

	p.ModEventsInRange(startBeat+0.5, startBeat+0.5+0.6-0.001, ilysa.FilterRGBLight(ringLasers),
		func(ctx ilysa.Context, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
		})

	p.ModEventsInRange(startBeat+0.5+0.6, startBeat+0.5+1.2, ilysa.FilterRGBLight(ringLasers),
		func(ctx ilysa.Context, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
		})

}

func (p Intro) Climb(startBeat float64) {
	var (
		light           = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		step            = 0.25
		lightIDSequence = []int{6, 7, 5, 8, 4, 9, 3, 10, 2, 11, 1, 12}
		count           = len(lightIDSequence)
		//maxLightID      = p.ActiveDifficultyProfile().MaxLightID(light)
		//lightIDs        = light2.FromInterval(1, maxLightID)
		//lightIDSets     = light2.Divide(lightIDs, maxLightID/2)

		backGrad = gradient.Table{
			{magnetPink, 0.0},
			{magnetWhite, 1.0},
		}
		sideGrad = gradient.Table{
			{magnetWhite, 0.0},
			{magnetPurple, 1.0},
		}
	)

	p.EventForBeat(startBeat, func(ctx ilysa.Context) {
		ctx.NewPreciseRotationEvent().
			SetRotation(360).
			SetStep(15).
			SetSpeed(1.3).
			SetProp(13)

		ctx.NewZoomEvent()
	})

	p.EventsForBeats(startBeat, step, count, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.GroupDivide(7), func(ctx ilysa.RangeLightIDContext) {
			ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightRedOn).
				SetColor(backGrad.GetInterpolatedColorFor(ctx.Pos))
		})

		switch {
		case ctx.Last:
			const exitValue = 3
			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeLeftRotatingLasers).SetValue(beatsaber.EventValueLightBlueFade).
				SetColor(magnetPurple)

			ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.LeftLaser).SetValue(exitValue).
				SetLockPosition(true).
				SetSpeed(exitValue).
				SetDirection(chroma.CounterClockwise)

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeRightRotatingLasers).SetValue(beatsaber.EventValueLightRedFade).
				SetColor(magnetPurple)

			ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.RightLaser).SetValue(exitValue).
				SetLockPosition(true).
				SetSpeed(exitValue).
				SetDirection(chroma.Clockwise)

		case ctx.Ordinal%2 == 0:
			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeLeftRotatingLasers).SetValue(beatsaber.EventValueLightBlueFlash).
				SetColor(sideGrad.GetInterpolatedColorFor(ctx.Pos))

			ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.LeftLaser).SetValue(ctx.Ordinal).
				SetLockPosition(true).SetSpeed(float64(ctx.Ordinal)).SetDirection(chroma.Clockwise)

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeRightRotatingLasers).SetValue(beatsaber.EventValueLightOff)

		case ctx.Ordinal%2 == 1:
			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeLeftRotatingLasers).SetValue(beatsaber.EventValueLightOff)

			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeRightRotatingLasers).SetValue(beatsaber.EventValueLightRedFlash).
				SetColor(sideGrad.GetInterpolatedColorFor(ctx.Pos))

			ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.RightLaser).SetValue(ctx.Ordinal).
				SetLockPosition(false).SetSpeed(float64(ctx.Ordinal)).SetDirection(chroma.CounterClockwise)
		}
	})
}

func (p Intro) Fall(startBeat float64) {
	var (
		//lightIDs    = light2.FromInterval(1, maxLightID)
		//lightIDSets = light2.Divide(lightIDs, count)
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		step     = 0.25
		count    = 4
		colorSet = colorful.NewSet(magnetPurple, magnetPink)
		values   = beatsaber.NewEventValueSet(
			beatsaber.EventValueLightRedOn,
			beatsaber.EventValueLightOff,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedFlash,
		)
	)

	p.EventsForBeats(startBeat, step, count, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.GroupDivide(count), func(ctx ilysa.RangeLightIDContext) {
			if ctx.Ordinal%ctx.LightIDSetLen != ctx.LightIDOrdinal {
				return
			}
			ctx.NewRGBLightingEvent().
				SetValue(values.Pick(ctx.Ordinal)).
				SetColor(colorSet.Next())
		})
		//if ctx.Ordinal <= 3 {
		//	e.SetLightID(lightIDSets[ctx.Ordinal])
		//}
	})
}

func (p Intro) Bridge(startBeat float64) {
	p.EventForBeat(startBeat, func(ctx ilysa.Context) {
		ctx.NewPreciseRotationEvent().
			SetRotation(180).
			SetStep(12.5).
			SetDirection(chroma.CounterClockwise).
			SetSpeed(3).
			SetProp(5).
			SetCounterSpin(true)
	})

	p.EventsForRange(startBeat, startBeat+1, 30, ease.OutCubic, func(ctx ilysa.Context) {
		if !ctx.Last {
			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeBackLasers).SetValue(beatsaber.EventValueLightBlueOn).
				SetColor(magnetPurple).
				SetAlpha(1 - ctx.Pos)
		} else {
			ctx.NewRGBLightingEvent().
				SetLight(beatsaber.EventTypeBackLasers).SetValue(beatsaber.EventValueLightRedFade).
				SetColor(magnetWhite)
		}
	})
}

func (p Intro) Outro(startBeat float64) {
	var (
		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers)
		sequence = []float64{0, 0.25, 0.50, 1.0, 1.25, 1.50, 2.0, 2.25, 2.50, 2.75, 3.25}
		//lightIDSet = light2.Fan(light2.FromInterval(1, maxLaserID), len(sequence))
	)

	p.EventForBeat(startBeat-0.001, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.Fan(len(sequence)), func(ctx ilysa.RangeLightIDContext) {
			ctx.NewRGBLightingEvent().
				SetValue(beatsaber.EventValueLightRedOn).
				SetColor(allColorsGradient.GetInterpolatedColorFor(ctx.LightIDPos))

		})
		//for i := 1; i <= maxLaserID; i++ {
		//	e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedFlash)
		//	e.SetSingleLightID(i)
		//	e.SetColor(allColorsGradient.GetInterpolatedColorFor(float64(i) / float64(maxLaserID)))
		//}
	})

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(light, lightid.Fan(len(sequence)), func(ctx ilysa.RangeLightIDContext) {
			if ctx.Ordinal%ctx.LightIDSetLen != ctx.LightIDOrdinal {
				return
			}
			ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff)
		})

		if ctx.Last {
			ctx.NewRGBLightingEvent().SetLight(light).SetValue(beatsaber.EventValueLightOff)
		}
	})
}

func (p Intro) OutroSplash(startBeat float64) {
	var (
		intensity  = 2.0
		sweepSpeed = 0.6
		grad       = gradient.Rainbow
		sequence   = []float64{0, 0.75, 1.5}
		leftLaser  = p.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers)
		rightLaser = p.NewBasicLight(beatsaber.EventTypeRightRotatingLasers)
		backLaser  = p.NewBasicLight(beatsaber.EventTypeBackLasers)
	)

	p.EventsForRange(startBeat, startBeat+2, 60, ease.Linear, func(ctx ilysa.Context) {
		ctx.RangeLightIDs(backLaser, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
			e := fx.BiasedColorSweep(ctx, intensity, sweepSpeed, grad)
			e.SetAlpha(intensity)
		})
	})

	p.ModEventsInRange(startBeat, startBeat+1, ilysa.FilterRGBLight(backLaser),
		func(ctx ilysa.Context, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
		})

	p.ModEventsInRange(startBeat+1, startBeat+2, ilysa.FilterRGBLight(backLaser),
		func(ctx ilysa.Context, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
		})

	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.Context) {
		g := append(gradient.Table{}, grad...)
		rand.Shuffle(len(g), func(i, j int) {
			g[i].Col, g[j].Col = g[j].Col, g[i].Col
		})
		p.EventsForRange(ctx.B, ctx.B+0.5, 15, ease.Linear, func(ctx ilysa.Context) {
			ctx.RangeLightIDs(leftLaser, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
				e := fx.Gradient(ctx, beatsaber.EventValueLightBlueOn, g)
				e.SetAlpha(float64(ctx.Ordinal)*0.5)
			})
			ctx.RangeLightIDs(rightLaser, lightid.AllIndividual, func(ctx ilysa.RangeLightIDContext) {
				e := fx.Gradient(ctx, beatsaber.EventValueLightRedOn, g)
				e.SetAlpha(float64(ctx.Ordinal)*0.5)
			})
		})

		p.ModEventsInRange(ctx.B, ctx.B+0.5, ilysa.FilterRGBLights(), func(ctx ilysa.Context, event ilysa.Event) {
			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.InCirc)
		})
	})
}
