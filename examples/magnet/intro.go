package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/fx"
)

func NewIntro(p *ilysa.Project, startBeat float64) Intro {
	return Intro{
		p:           p,
		BaseContext: p.WithBeatOffset(startBeat),
	}

}

type Intro struct {
	p *ilysa.Project
	ilysa.BaseContext
}

func (p Intro) Play() {
	p.PianoDoubles(0)
	p.LeadinDrums(2.25)
	p.BassTwang(2.5)
	p.StartSplash(4)
	p.Rhythm(4)
	p.Rhythm(8)
	p.Rhythm(12)
	p.Melody1(4)
	p.Melody2(7.25, false)
	p.Melody1(8)
	p.Melody2(11.25, true)
	p.Melody3(12)
	p.Chorus(16)
	p.PianoRoll(20.5, 6)
	p.Trill(22.5)
	p.Climb(23.5)
	p.TrillNoFade(26.5)
	p.Fall(27.25)
	p.Trill(28.5)
	p.Bridge(29.0)
	p.Rhythm(30)
	p.Outro(30.5)
	p.OutroSplash(34.0)
}

func (p Intro) PianoDoubles(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	colors := colorful.NewSet(magnetPurple, magnetPink, magnetWhite, colorful.Black)

	light := ilysa.NewCombinedLight(
		ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeRingLights, p),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
		ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
	)

	ctx.EventsForSequence(0, []float64{0, 0.75, 1.25, 1.75, 2.25}, func(ctx ilysa.SequenceContext) {
		grad := gradient.New(colors.Index(ctx.Ordinal()), colors.Index(ctx.Ordinal()+1))
		ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
			fx.Gradient(ctx, grad)
		})
	})

	ctx.EventForBeat(2.25, func(ctx ilysa.TimeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
		})
	})
}

func (p Intro) LeadinDrums(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	light := ilysa.NewSequenceLight(
		ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, p),
			ilysa.ToLightTransformer(ilysa.Fan(2)),
		),
		ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, p),
			ilysa.ToLightTransformer(ilysa.Fan(2)),
		),
	)

	ctx.EventsForSequence(0, []float64{0, 0.25, 0.75, 1, 1.5}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(1),
			ilysa.WithLockPosition(false), ilysa.WithSpeed(0), ilysa.WithDirection(chroma.Clockwise),
		)
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(1),
			ilysa.WithLockPosition(false), ilysa.WithSpeed(0), ilysa.WithDirection(chroma.CounterClockwise),
		)

		ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(ilysa.WithColor(crossickColors.Next()))

			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			os := ctx.NextBOffset()
			if ctx.Last() {
				os = 0.25
			}
			oe.ShiftBeat(os)
		})
	})
}

func (p Intro) BassTwang(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	const (
		steps      = 60
		intensity  = 1
		sweepSpeed = 1.8
	)

	var (
		light = ilysa.TransformLight(
			p.NewBasicLight(beatsaber.EventTypeBackLasers),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		)
		grad = magnetRainbow
	)

	ctx.EventsForRange(0, 1.5, steps, ease.InOutCirc, func(ctx ilysa.TimeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, sweepSpeed, grad)
			fx.AlphaBlend(ctx, e, 0, 0.75, 0, intensity, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.75, 1, intensity, 0, ease.OutBounce)
		})
	})
}

func (p Intro) StartSplash(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(8))
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(8))

		//ctx.NewRGBLightingEvent(
		//	ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
		//	ilysa.WithColor(sukoyaPink))
		//ctx.NewRGBLightingEvent(
		//	ilysa.WithType(beatsaber.EventTypeRightRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightRedFlash),
		//	ilysa.WithColor(shirayukiPurple))
		//ctx.NewRGBLightingEvent(
		//	ilysa.WithType(beatsaber.EventTypeCenterLights), ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
		//	ilysa.WithColor(magnetPurple))

		ctx.NewPreciseZoomEvent(ilysa.WithStep(-0.33))
	})
}

func (p Intro) Rhythm(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	light := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeRingLights, p),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(180),
			ilysa.WithStep(0),
			ilysa.WithProp(1),
			ilysa.WithSpeed(24),
			ilysa.WithDirection(chroma.Clockwise),
		)

		grad := magnetGradient.RotateRand()
		ctx.EventsForRange(ctx.B(), ctx.B()+0.5, 12, ease.InCubic, func(ctx ilysa.TimeContext) {
			ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
				e := fx.Gradient(ctx, grad)
				fx.Ripple(ctx, e, 1.2)
			})
		})
	})

	ctx.EventsForSequence(0, []float64{1, 3}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45),
			ilysa.WithStep(12.5),
			ilysa.WithProp(20),
			ilysa.WithSpeed(20),
			ilysa.WithDirection(chroma.CounterClockwise),
		)
	})

	ctx.EventForBeat(1, func(ctx ilysa.TimeContext) {
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedFade),
			ilysa.WithColor(sukoyaPink),
		)

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			ilysa.WithColor(sukoyaWhite),
		)
	})

	ctx.EventForBeat(3, func(ctx ilysa.TimeContext) {
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeRingLights),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			ilysa.WithColor(magnetColors.Next()),
		)

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			ilysa.WithColor(shirayukiPurple),
		)

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedFade),
			ilysa.WithColor(shirayukiGold),
		)
	})
}

func (p Intro) Melody1(startBeat float64) {
	var (
		sequence = []float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75}
		//light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Divide(3))
	)

	ctx := p.WithBeatOffset(startBeat)

	p.PianoGlow(ctx, sequence, 5, false)

	//p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
	//	ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
	//		//if ctx.Ordinal()%ctx.LightIDLen() != ctx.LightIDOrdinal() {
	//		//	return
	//		//}
	//		//fx.Gradient(ctx, beatsaber.EventValueLightBlueOn, magnetGradient)
	//		ctx.NewRGBLightingEvent(
	//			ilysa.WithValue(beatsaber.EventValueLightBlueOn),
	//			ilysa.WithColor(magnetPurple),
	//		)
	//
	//		e := ctx.NewRGBLightingEvent(
	//			ilysa.WithValue(beatsaber.EventValueLightOff),
	//		)
	//		if !ctx.Last() {
	//			e.ShiftBeat(ctx.NextBOffset())
	//		} else {
	//			e.ShiftBeat(0.249)
	//		}
	//		//ilysa.WithLightID(ctx.LightIDCur()),
	//
	//		//if ctx.Ordinal() > 0 {
	//		//	ctx.NewRGBLightingEvent().SetValue(beatsaber.EventValueLightOff).
	//		//		SetTypeID(ctx.PreLightID)
	//		//}
	//	})
	//})
	//
	////p.EventForBeat(startBeat+2.999, func(ctx ilysa.TimeContext) {
	////	ctx.NewRGBLightingEvent().SetType(light).SetValue(beatsaber.EventValueLightOff)
	////})
}

func (p Intro) Melody2(startBeat float64, reverseZoom bool) {
	var (
		sequence = []float64{0, 0.25, 0.50}
		//light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Fan(3))
	)

	ctx := p.WithBeatOffset(startBeat)

	p.PianoGlow(ctx, sequence, 3, true)

	//p.EventForBeat(startBeat-0.01, func(ctx ilysa.TimeContext) {
	//	ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
	//	})
	//})
	//
	p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		//ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
		//	ctx.NewRGBLightingEvent(
		//		ilysa.WithColor(magnetPink),
		//	)
		//
		//	oe := ctx.NewRGBLightingEvent(
		//		ilysa.WithValue(beatsaber.EventValueLightOff),
		//	)
		//	if !ctx.Last() {
		//		oe.ShiftBeat(ctx.NextBOffset())
		//	} else {
		//		oe.ShiftBeat(0.245)
		//	}
		//})

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
		//light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Divide(3))
	)
	ctx := p.WithBeatOffset(startBeat)

	p.PianoGlow(ctx, sequence, 5, false)

	//p.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
	//	ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLightingEvent(
	//			ilysa.WithValue(beatsaber.EventValueLightBlueOn),
	//			ilysa.WithColor(magnetPurple),
	//		)
	//
	//		oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
	//		if !ctx.Last() {
	//			oe.ShiftBeat(ctx.NextBOffset())
	//		} else {
	//			oe.ShiftBeat(0.45)
	//		}
	//	})
	//})
}

func (p Intro) Chorus(startBeat float64) {
	var (
		sequence = []float64{0, 1, 2, 2.75, 3.5, 4}
		//colorGrad = allColorsGradient
	)

	ctx := p.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightOff))

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightOff))
	})

	ctx.EventsForSequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.NewPreciseZoomEvent(ilysa.WithStep(0.2))

		re := ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45+15*float64(ctx.Ordinal())),
			ilysa.WithStep(5+(1.5*float64(ctx.Ordinal()))),
			ilysa.WithProp(20),
			ilysa.WithSpeed(4+float64(ctx.Ordinal())*2),
		)

		if ctx.Ordinal()%2 == 0 {
			re.Mod(ilysa.WithDirection(chroma.Clockwise))
		} else {
			re.Mod(ilysa.WithDirection(chroma.CounterClockwise))
		}

		//if ctx.Ordinal() == 5 {
		//	re.Mod(ilysa.WithRotation(360))
		//}

		//grad := append(gradient.Table{}, colorGrad...)
		//rand.Shuffle(len(colorGrad), func(i, j int) {
		//	grad[i].Col, grad[j].Col = grad[j].Col, grad[i].Col
		//})

		var light ilysa.Light

		if ctx.Ordinal()%2 == 0 {
			light = ilysa.TransformLight(
				p.NewBasicLight(beatsaber.EventTypeBackLasers),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			)
		} else {
			light = ilysa.TransformLight(
				p.NewBasicLight(beatsaber.EventTypeBackLasers),
				ilysa.ToLightTransformer(ilysa.Reverse),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			)
		}

		ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
			nbo := seqCtx.NextBOffset()
			if seqCtx.Last() {
				nbo = 0.5
			}

			e := fx.Gradient(ctx, magnetRainbowPale)
			fx.Ripple(ctx, e, 1*nbo)

			if seqCtx.Last() {
				return
			}

			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			oe.ShiftBeat(nbo)
		})

		//for i := 1; i <= LightIDMax; i++ {
		//	gradientPos := scale.Clamped(1, float64(LightIDMax), 0, 1)
		//	color := grad.Ierp(gradientPos(float64(i)))
		//
		//	e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
		//	e.SetSingleLightID(i)
		//	e.SetColor(color)
		//	e.Beat += 1.0 / 64.0
		//}
	})
}

func (p Intro) PianoRoll(startBeat float64, count int) {
	ctx := p.WithBeatOffset(startBeat)

	seq := []float64{0, 0.25, 0.50, 0.75, 1.00, 1.25}

	p.PianoGlow(ctx, seq, 6, false)
}

func (p Intro) Trill(startBeat float64) {
	var (
	//backLasers = p.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.DivideSingle)
	//ringLasers = p.NewBasicLight(beatsaber.EventTypeRingLights).Transform(ilysa.DivideSingle)
	//step       = 0.125
	//count      = 5
	//ratio      = 0.666
	//lightCount = int(ratio * float64(LightIDMax))
	)

	ctx := p.WithBeatOffset(startBeat)

	l := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Shuffle),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(0, 0.5, 12, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(l, func(ctx ilysa.TimeLightContext) {
			e := fx.Gradient(ctx, magnetGradient)
			fx.Ripple(ctx, e, 0.15,
				fx.WithAlphaBlend(0, 0.3, 0, 1, ease.OutSine),
				fx.WithAlphaBlend(0.3, 1, 1, 0, ease.InSine),
			)
		})
	})

	//p.EventsForBeats(startBeat, step, count, func(ctx ilysa.TimeContext) {
	//	ctx.WithLight(backLasers, func(ctx ilysa.TimeLightContext) {
	//		if rand.Float64() > ratio {
	//			return
	//		}
	//
	//		ctx.NewRGBLightingEvent(ilysa.WithColor(
	//			allColorsGradient.Ierp(rand.Float64()),
	//		))
	//
	//		oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
	//		oe.ShiftBeat(step / 2)
	//	})
	//	//for i := 0; i < lightCount; i++ {
	//	//	e := ctx.NewRGBLightingEvent(backLasers, beatsaber.EventValueLightRedOn)
	//	//	e.SetSingleLightID(rand.Intn(LightIDMax) + 1)
	//	//	e.SetColor(allColorsGradient.Ierp(rand.Float64()))
	//	//
	//	//	oe := ctx.NewRGBLightingEvent(backLasers, beatsaber.EventValueLightOff)
	//	//	oe.Beat += step / 2
	//	//
	//	//	if !ctx.Last {
	//	//		continue
	//	//	}
	//	//
	//	//}
	//})
	//p.EventsForRange(startBeat+0.5, startBeat+0.5+1.2, 30, ease.Linear, func(ctx ilysa.TimeContext) {
	//	ctx.WithLight(ringLasers, func(ctx ilysa.TimeLightContext) {
	//		fx.ColorSweep(ctx, 0.4, magnetRainbow)
	//	})
	//})
	//
	//p.ModEventsInRange(startBeat+0.5, startBeat+0.5+0.6-0.001, ilysa.FilterRGBLight(ringLasers),
	//	func(ctx ilysa.TimeContext, event ilysa.Event) {
	//		fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
	//	})
	//
	//p.ModEventsInRange(startBeat+0.5+0.6, startBeat+0.5+1.2, ilysa.FilterRGBLight(ringLasers),
	//	func(ctx ilysa.TimeContext, event ilysa.Event) {
	//		fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
	//	})
}

func (p Intro) TrillNoFade(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	l := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Shuffle),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(0, 0.5, 12, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(l, func(ctx ilysa.TimeLightContext) {
			e := fx.Gradient(ctx, magnetGradient)
			fx.Ripple(ctx, e, 0.15,
				fx.WithAlphaBlend(0, 1, 0, 1, ease.OutSine),
			)
		})
	})
}

func (p Intro) Climb(startBeat float64) {
	const rotatingLasersExitSpeed = 3
	var (
		step  = 0.25
		count = 12

		blGrad = gradient.FromSet(crossickColors)
		rlGrad = magnetGradient

		light = ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p),
			ilysa.ToSequenceLightTransformer(ilysa.Divide(7)),
		)
	)

	ctx := p.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(360),
			ilysa.WithStep(15),
			ilysa.WithSpeed(1.3),
			ilysa.WithProp(13),
		)
		ctx.NewZoomEvent()
	})

	ctx.EventsForBeats(0, step, count, func(ctx ilysa.TimeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithColor(blGrad.Ierp(ctx.T())),
			)
		})

		switch {
		case ctx.Last():
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
				ilysa.WithColor(magnetPurple),
			)
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(magnetPurple),
			)

			exitArgs := []ilysa.PreciseRotationSpeedEventOpt{
				ilysa.WithIntValue(rotatingLasersExitSpeed),
				ilysa.WithLockPosition(true),
				ilysa.WithSpeed(rotatingLasersExitSpeed),
			}

			ctx.NewPreciseRotationSpeedEvent(
				append([]ilysa.PreciseRotationSpeedEventOpt{
					ilysa.WithDirectionalLaser(ilysa.LeftLaser),
					ilysa.WithDirection(chroma.CounterClockwise),
				}, exitArgs...)...,
			)
			ctx.NewPreciseRotationSpeedEvent(
				append([]ilysa.PreciseRotationSpeedEventOpt{
					ilysa.WithDirectionalLaser(ilysa.RightLaser),
					ilysa.WithDirection(chroma.Clockwise),
				}, exitArgs...)...,
			)

		case ctx.Ordinal()%2 == 0:
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
				ilysa.WithColor(rlGrad.Ierp(ctx.T())),
			)

			ctx.NewPreciseRotationSpeedEvent(
				ilysa.WithDirectionalLaser(ilysa.LeftLaser),
				ilysa.WithIntValue(ctx.Ordinal()),
				ilysa.WithLockPosition(true),
				ilysa.WithSpeed(float64(ctx.Ordinal())*2),
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
				ilysa.WithColor(rlGrad.Ierp(ctx.T())),
			)

			ctx.NewPreciseRotationSpeedEvent(
				ilysa.WithDirectionalLaser(ilysa.RightLaser),
				ilysa.WithIntValue(ctx.Ordinal()),
				ilysa.WithLockPosition(false),
				ilysa.WithSpeed(float64(ctx.Ordinal())*2),
				ilysa.WithDirection(chroma.CounterClockwise),
			)
		}
	})
}

func (p Intro) Fall(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	p.PianoGlow(ctx, []float64{0, 0.25, 0.5, 0.75}, 5, true)
}

func (p Intro) Bridge(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(180),
			ilysa.WithStep(12.5),
			ilysa.WithDirection(chroma.CounterClockwise),
			ilysa.WithSpeed(3),
			ilysa.WithProp(5),
			ilysa.WithCounterSpin(true),
		)
	})

	light := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Divide(7)),
	)

	ctx.EventsForRange(0, 1, 30, ease.OutCubic, func(ctx ilysa.TimeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.Gradient(ctx, shirayukiGradient)
			fx.AlphaBlend(ctx, e, 0, 1, 0, 1, ease.Linear)
		})
	})
}

func (p Intro) Outro(startBeat float64) {
	var (
		seq1 = []float64{0, 0.25, 0.50}
		seq2 = []float64{1.0, 1.25, 1.50}
		seq3 = []float64{2.0, 2.25, 2.50, 2.75, 3.25}
	)

	ctx := p.WithBeatOffset(startBeat)

	p.PianoTransmute(ctx, seq1, 3, true, shirayukiSingleGradient)
	p.PianoTransmute(ctx, seq2, 3, true, sukoyaSingleGradient)
	p.PianoGlow(ctx, seq3, 5, true)

	//var (
	//	sequence = []float64{0, 0.25, 0.50, 1.0, 1.25, 1.50, 2.0, 2.25, 2.50, 2.75, 3.25}
	//	light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Fan(len(sequence)))
	//)
	//
	//ctx := p.WithBeatOffset(startBeat)
	//
	//ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
	//	ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewRGBLightingEvent(
	//			ilysa.WithColor(allColorsGradient.Ierp(ctx.LightIDT())),
	//		)
	//	})
	//})
	//
	//ctx.EventsForSequence(0, sequence[:len(sequence)-1], func(ctx ilysa.SequenceContext) {
	//	ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
	//	})
	//})
	//
	//ctx.EventForBeat(sequence[len(sequence)-1], func(ctx ilysa.TimeContext) {
	//	ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
	//	})
	//})
}

func (p Intro) OutroSplash(startBeat float64) {
	var (
		sweepSpeed = 1.5
		grad       = magnetRainbow
		sequence   = []float64{0, 0.75, 1.5}
		//leftLaser  = p.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers).Transform(ilysa.DivideSingle)
		//rightLaser = p.NewBasicLight(beatsaber.EventTypeRightRotatingLasers).Transform(ilysa.DivideSingle)
		backLaser = p.NewBasicLight(beatsaber.EventTypeBackLasers).Transform(ilysa.DivideSingle)
	)

	ctx := p.WithBeatOffset(startBeat)

	ctx.EventsForRange(0, 4, 60, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(backLaser, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, sweepSpeed, grad)
			fx.AlphaBlend(ctx, e, 0, 0.5, 0, 2, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.5, 1, 2, 0, ease.OutCubic)
		})
	})

	ctx.EventsForSequence(0, sequence, func(ctx ilysa.SequenceContext) {
		if ctx.Last() {
			ctx.NewPreciseRotationEvent(
				ilysa.WithRotation(360),
				ilysa.WithStep(0),
				ilysa.WithDirection(chroma.Clockwise),
				ilysa.WithSpeed(7),
				ilysa.WithProp(0.8),
				ilysa.WithCounterSpin(true),
			)
		} else {
			ctx.NewPreciseRotationEvent(
				ilysa.WithRotation(45),
				ilysa.WithStep(12.5),
				ilysa.WithSpeed(26),
				ilysa.WithProp(8),
				ilysa.WithDirection(chroma.CounterClockwise),
			)
		}
		//nbo := ctx.NextBOffset()
		//if ctx.Last() {
		//	nbo = 0.75
		//}
		p.Rush(ctx, ctx.B()-0.25, ctx.B()+0.25, 0.75, float64(ctx.Ordinal()), magnetGradient)

		//g := grad.RotateRand()
		//ctx.EventsForRange(ctx.B(), ctx.B()+splashDuration, 15, ease.Linear, func(ctx ilysa.TimeContext) {
		//	ctx.WithLight(leftLaser, func(ctx ilysa.TimeLightContext) {
		//		e := fx.Gradient(ctx, g)
		//		e.Mod(ilysa.WithAlpha(float64(ctx.Ordinal()) * 0.5))
		//	})
		//	ctx.WithLight(rightLaser, func(ctx ilysa.TimeLightContext) {
		//		e := fx.Gradient(ctx, g)
		//		e.Mod(ilysa.WithAlpha(float64(ctx.Ordinal()) * 0.5))
		//	})
		//})
		//
		//ctx.ModEventsInRange(ctx.B(), ctx.B()+splashDuration, ilysa.FilterRGBLights(), func(ctx ilysa.TimeContext, event ilysa.Event) {
		//	fx.RGBAlphaBlend(ctx, event, 1, 0, ease.InCirc)
		//})
	})
}

func (p Intro) PianoGlow(ctx ilysa.BaseContext, sequence []float64, divisor int, shuffle bool) {
	lightSweepDiv := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(divisor)),
	).(ilysa.SequenceLight)

	if shuffle {
		lightSweepDiv = lightSweepDiv.Shuffle()
	}

	ctx.EventsForSequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		grad := magnetRainbowPale.RotateRand()
		ctx.EventsForRange(ctx.B(), ctx.B()+0.435, 12, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaBlend(ctx, e, 0, 0.15, 0, 1, ease.OutCubic)
				fx.AlphaBlend(ctx, e, 0.3, 1, 3, 0, ease.InCubic)
			})
		})
	})
}

func (p Intro) PianoTransmute(ctx ilysa.BaseContext, sequence []float64, divisor int, shuffle bool, grad gradient.Table) {
	lightSweepDiv := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(divisor)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(ilysa.SequenceLight)

	if shuffle {
		lightSweepDiv = lightSweepDiv.Shuffle()
	}

	ctx.EventsForSequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.EventsForRange(ctx.B(), ctx.B()+0.435, 12, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaBlend(ctx, e, 0, 0.3, 1, 3, ease.OutCubic)
				fx.AlphaBlend(ctx, e, 0.3, 1, 3, 1, ease.InCubic)
			})
		})
	})
}

func (p Intro) Rush(ctx ilysa.BaseContext, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	light := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeRingLights, p),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(startBeat, endBeat, 45, ease.InExpo, func(ctx ilysa.TimeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, 2, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaBlend(ctx, e, 0, 0.6, 1, peakAlpha, ease.OutCubic)
			fx.AlphaBlend(ctx, e, 0.6, 1.0, peakAlpha, 0, ease.InCubic)
		})
	})
}
