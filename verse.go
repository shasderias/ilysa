package main

//
//type Verse struct {
//	*ilysa.Project
//	offset float64
//}
//
//func (p Verse) PlayVerse1(startBeat float64) {
//	p.offset = startBeat
//
//	p.EventForBeat(startBeat-0.001, func(ctx ilysa.Context) {
//		gen.OffAll(ctx)
//		ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.LeftLaser).SetSpeed(1.5)
//		ctx.NewPreciseRotationSpeedEvent().SetLaser(ilysa.RightLaser).SetSpeed(1.5)
//	})
//
//	p.Rhythm(0)
//	p.Rhythm(4)
//	p.Rhythm(8)
//	p.Rhythm(12)
//	p.Rhythm(16)
//	p.Rhythm(20)
//	p.Rhythm(24)
//	p.Rhythm(28)
//
//	//p.Lyrics(0)
//	p.PianoBackstep(7)
//
//	//p.Lyrics(8)
//	//p.PianoBackstep(15)
//
//}
//
//func (p Verse) Rhythm(startBeat float64) {
//	var (
//		kickDrumLights = []ilysa.Light{
//			p.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers),
//			p.NewBasicLight(beatsaber.EventTypeRightRotatingLasers),
//		}
//		kickDrumSequence = []float64{0, 2.5}
//		kickDrumColors   = shirayukiColors
//	)
//	p.EventsForSequence(startBeat+p.offset, kickDrumSequence, func(ctx ilysa.Context) {
//		for _, light := range kickDrumLights {
//			e := ctx.NewRGBLightingEvent().SetLight(light).
//				SetValue(beatsaber.EventValueLightRedFade).
//				SetColor(kickDrumColors.Next())
//			if startBeat == 0 && ctx.ordinal == 0 {
//				e.SetAlpha(10)
//			} else {
//				e.SetAlpha(0.7)
//			}
//		}
//	})
//
//	const (
//		rippleDuration = 2
//	)
//
//	var (
//		rippleStart  = startBeat + p.offset
//		rippleEnd    = rippleStart + rippleDuration
//		rippleLights = p.NewBasicLight(beatsaber.EventTypeRingLights)
//		rippleStep   = 0.6
//		grad         = gradient.Table{
//			{shirayukiPurple, 0.0},
//			{shirayukiGold, 0.3},
//			{shirayukiGold, 0.7},
//			{shirayukiPurple, 1.0},
//		}
//	)
//	p.EventForBeat(rippleStart, func(ctx ilysa.Context) {
//		ctx.NewPreciseRotationEvent().
//			SetRotation(90).
//			SetStep(22.5).
//			SetSpeed(2).
//			SetProp(0.3)
//	})
//
//	p.EventsForRange(rippleStart, rippleEnd, 30, ease.Linear, func(ctx ilysa.Context) {
//		ctx.UseLight(rippleLights, lightid.AllIndividual, func(ctx ilysa.ContextWithLight) {
//			e := fx.ColorSweep(ctx, 1.5, 1.4, grad)
//			e.Beat += ctx.LightIDT * rippleStep
//			switch {
//			case ctx.t <= 0.5:
//				alphaScale := util.ScaleToUnitInterval(0, 0.5)
//				e.SetAlpha(e.GetAlpha() * ease.InOutQuart(alphaScale(ctx.t)))
//			case ctx.t > 0.8:
//				alphaScale := util.ScaleToUnitInterval(0.8, 1)
//				e.SetAlpha(e.GetAlpha() * ease.InExpo(1-alphaScale(ctx.t)))
//			}
//		})
//	})
//}
//
//func (p Verse) Lyrics(startBeat float64) {
//	var (
//		// 52-58.5
//		sequence = []float64{0, 0.5, 0.75, 1.25, 1.75, 2.25, 2.75, 3.0, 3.5, 4.0, 4.5, 4.75, 5.25, 5.75, 6.25, 6.5}
//		light    = p.NewBasicLight(beatsaber.EventTypeBackLasers)
//	)
//
//	p.EventsForSequence(startBeat+p.offset, sequence, func(ctx ilysa.Context) {
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
//func (p Verse) PianoBackstep(startBeat float64) {
//	var (
//		sequence  = []float64{0, 0.5}
//		backLight = p.NewBasicLight(beatsaber.EventTypeBackLasers)
//	)
//
//	p.EventsForSequence(startBeat+p.offset, sequence, func(ctx ilysa.Context) {
//		e := ctx.NewPreciseRotationEvent().
//			SetRotation(5).
//			SetStep(5).
//			SetProp(12).
//			SetSpeed(10)
//
//		if ctx.ordinal%2 == 0 {
//			e.SetDirection(chroma.Clockwise)
//		} else {
//			e.SetDirection(chroma.CounterClockwise)
//		}
//
//		p.EventsForRange(ctx.b, ctx.b+0.25, 6, ease.Linear, func(ctx ilysa.Context) {
//			ctx.UseLight(backLight, lightid.GroupDivide(2), func(ctx ilysa.ContextWithLight) {
//				fx.Gradient(ctx, beatsaber.EventValueLightRedOn, magnetGradient)
//			})
//		})
//
//		//ctx.ModEventsInRange(ctx.b, ctx.b+0.10, ilysa.FilterRGBLight(backLight), func(ctx ilysa.Context, event ilysa.Event) {
//		//	fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCirc)
//		//})
//		ctx.ModEventsInRange(ctx.b, ctx.b+0.25, ilysa.FilterRGBLight(backLight), func(ctx ilysa.Context, event ilysa.Event) {
//			fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
//		})
//
//		//ctx.UseLight(backLight, lightid.GroupDivide(2), func(ctx ilysa.ContextWithLight) {
//		//	if ctx.ordinal != ctx.LightIDOrdinal {
//		//		return
//		//	}
//		//
//		//	fx.BiasedColorSweep(ctx, 2, 0.8, gradient.Rainbow)
//		//})
//	})
//}
