package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light2"
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

	light := light2.NewCombinedLight(
		light2.TransformLight(
			light2.NewBasicLight(beatsaber.EventTypeRingLights, p),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
		light2.TransformLight(
			light2.NewBasicLight(beatsaber.EventTypeBackLasers, p),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		),
	)

	ctx.EventsForSequence(0, []float64{0, 0.75, 1.25, 1.75, 2.25}, func(ctx ilysa.SequenceContext) {
		grad := gradient.New(colors.Index(ctx.Ordinal()), colors.Index(ctx.Ordinal()+1))
		ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
			fx.Gradient(ctx, grad)
		})
	})

	ctx.EventForBeat(2.25, func(ctx ilysa.RangeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			ctx.NewLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
		})
	})
}

func (p Intro) LeadinDrums(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	light := light2.NewSequenceLight(
		light2.TransformLight(
			light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, p),
			ilysa.ToLightTransformer(ilysa.Fan(2)),
		),
		light2.TransformLight(
			light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, p),
			ilysa.ToLightTransformer(ilysa.Fan(2)),
		),
	)

	ctx.EventsForSequence(0, []float64{0, 0.25, 0.75, 1, 1.5}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(1),
			evt.WithLockPosition(false), evt.WithPreciseLaserSpeed(0), evt.WithLaserDirection(chroma.Clockwise),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(1),
			evt.WithLockPosition(false), evt.WithPreciseLaserSpeed(0), evt.WithLaserDirection(chroma.CounterClockwise),
		)

		ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(evt.WithColor(crossickColors.Next()))

			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			os := ctx.SequenceNextBOffset()
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
		light = light2.TransformLight(
			p.NewBasicLight(beatsaber.EventTypeBackLasers),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		)
		grad = magnetRainbow
	)

	ctx.EventsForRange(0, 1.5, steps, ease.InOutCirc, func(ctx ilysa.RangeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, sweepSpeed, grad)
			fx.AlphaBlend(ctx, e, 0, 0.75, 0, intensity, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.75, 1, intensity, 0, ease.OutBounce)
		})
	})
}

func (p Intro) StartSplash(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(8))
		ctx.NewLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(8))

		//ctx.NewRGBLighting(
		//	ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
		//	ilysa.WithColor(sukoyaPink))
		//ctx.NewRGBLighting(
		//	ilysa.WithType(beatsaber.EventTypeRightRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightRedFlash),
		//	ilysa.WithColor(shirayukiPurple))
		//ctx.NewRGBLighting(
		//	ilysa.WithType(beatsaber.EventTypeCenterLights), ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
		//	ilysa.WithColor(magnetPurple))

		ctx.NewPreciseZoom(evt.WithRotationStep(-0.33))
	})
}

func (p Intro) Rhythm(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	light := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeRingLights, p),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(0),
			evt.WithProp(1),
			evt.WithPreciseLaserSpeed(24),
			evt.WithLaserDirection(chroma.Clockwise),
		)

		grad := magnetGradient.RotateRand()
		ctx.Range(ctx.B(), ctx.B()+0.5, 12, ease.InCubic, func(ctx ilysa.RangeContext) {
			ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
				e := fx.Gradient(ctx, grad)
				fx.Ripple(ctx, e, 1.2)
			})
		})
	})

	ctx.EventsForSequence(0, []float64{1, 3}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(12.5),
			evt.WithProp(20),
			evt.WithPreciseLaserSpeed(20),
			evt.WithLaserDirection(chroma.CounterClockwise),
		)
	})

	ctx.EventForBeat(1, func(ctx ilysa.RangeContext) {
		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedFade),
			evt.WithColor(sukoyaPink),
		)

		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			evt.WithColor(sukoyaWhite),
		)
	})

	ctx.EventForBeat(3, func(ctx ilysa.RangeContext) {
		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeRingLights),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			evt.WithColor(magnetColors.Next()),
		)

		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			evt.WithColor(shirayukiPurple),
		)

		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedFade),
			evt.WithColor(shirayukiGold),
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

	//p.Sequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
	//	ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
	//		//if ctx.Ordinal()%ctx.LightIDLen() != ctx.LightIDOrdinal() {
	//		//	return
	//		//}
	//		//fx.Gradient(ctx, beatsaber.EventValueLightBlueOn, magnetGradient)
	//		ctx.NewRGBLighting(
	//			ilysa.WithValue(beatsaber.EventValueLightBlueOn),
	//			ilysa.WithColor(magnetPurple),
	//		)
	//
	//		e := ctx.NewRGBLighting(
	//			ilysa.WithValue(beatsaber.EventValueLightOff),
	//		)
	//		if !ctx.Last() {
	//			e.ShiftBeat(ctx.SequenceNextBOffset())
	//		} else {
	//			e.ShiftBeat(0.249)
	//		}
	//		//ilysa.WithLightID(ctx.LightIDCur()),
	//
	//		//if ctx.Ordinal() > 0 {
	//		//	ctx.NewRGBLighting().SetValue(beatsaber.EventValueLightOff).
	//		//		SetTypeID(ctx.PreLightID)
	//		//}
	//	})
	//})
	//
	////p.EventForBeat(startBeat+2.999, func(ctx ilysa.RangeContext) {
	////	ctx.NewRGBLighting().SetType(light).SetValue(beatsaber.EventValueLightOff)
	////})
}

func (p Intro) Melody2(startBeat float64, reverseZoom bool) {
	var (
		sequence = []float64{0, 0.25, 0.50}
		//light    = p.NewBasicLight(beatsaber.EventTypeBackLasers).TransformToSequence(ilysa.Fan(3))
	)

	ctx := p.WithBeatOffset(startBeat)

	p.PianoGlow(ctx, sequence, 3, true)

	//p.EventForBeat(startBeat-0.01, func(ctx ilysa.RangeContext) {
	//	ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewLighting(ilysa.WithValue(beatsaber.EventValueLightOff))
	//	})
	//})
	//
	p.Sequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		//ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
		//	ctx.NewRGBLighting(
		//		ilysa.WithColor(magnetPink),
		//	)
		//
		//	oe := ctx.NewRGBLighting(
		//		ilysa.WithValue(beatsaber.EventValueLightOff),
		//	)
		//	if !ctx.Last() {
		//		oe.ShiftBeat(ctx.SequenceNextBOffset())
		//	} else {
		//		oe.ShiftBeat(0.245)
		//	}
		//})

		ze := ctx.NewPreciseZoom()
		if reverseZoom {
			ze.Apply(evt.WithRotationStep(0.3))
		} else {
			ze.Apply(evt.WithRotationStep(-0.3))
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

	//p.Sequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
	//	ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLighting(
	//			ilysa.WithValue(beatsaber.EventValueLightBlueOn),
	//			ilysa.WithColor(magnetPurple),
	//		)
	//
	//		oe := ctx.NewRGBLighting(ilysa.WithValue(beatsaber.EventValueLightOff))
	//		if !ctx.Last() {
	//			oe.ShiftBeat(ctx.SequenceNextBOffset())
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

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightOff))

		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeRightRotatingLasers), ilysa.WithValue(beatsaber.EventValueLightOff))
	})

	ctx.EventsForSequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.NewPreciseZoom(evt.WithRotationStep(0.2))

		re := ctx.NewPreciseRotation(
			evt.WithRotation(45+15*float64(ctx.Ordinal())),
			evt.WithRotationStep(5+(1.5*float64(ctx.Ordinal()))),
			evt.WithProp(20),
			evt.WithPreciseLaserSpeed(4+float64(ctx.Ordinal())*2),
		)

		if ctx.Ordinal()%2 == 0 {
			re.Apply(evt.WithLaserDirection(chroma.Clockwise))
		} else {
			re.Apply(evt.WithLaserDirection(chroma.CounterClockwise))
		}

		//if ctx.Ordinal() == 5 {
		//	re.Apply(ilysa.WithRotation(360))
		//}

		//grad := append(gradient.Table{}, colorGrad...)
		//rand.Shuffle(len(colorGrad), func(i, j int) {
		//	grad[i].Col, grad[j].Col = grad[j].Col, grad[i].Col
		//})

		var light light2.Light

		if ctx.Ordinal()%2 == 0 {
			light = light2.TransformLight(
				p.NewBasicLight(beatsaber.EventTypeBackLasers),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			)
		} else {
			light = light2.TransformLight(
				p.NewBasicLight(beatsaber.EventTypeBackLasers),
				ilysa.ToLightTransformer(ilysa.Reverse),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			)
		}

		ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
			nbo := seqCtx.SequenceNextBOffset()
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
		//	e := ctx.NewRGBLighting(light, beatsaber.EventValueLightRedOn)
		//	e.SetSingleLightID(i)
		//	e.SetColor(color)
		//	e.beat += 1.0 / 64.0
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

	l := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Shuffle),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(0, 0.5, 12, ease.Linear, func(ctx ilysa.RangeContext) {
		ctx.WithLight(l, func(ctx ilysa.TimeLightContext) {
			e := fx.Gradient(ctx, magnetGradient)
			fx.Ripple(ctx, e, 0.15,
				fx.WithAlphaBlend(0, 0.3, 0, 1, ease.OutSine),
				fx.WithAlphaBlend(0.3, 1, 1, 0, ease.InSine),
			)
		})
	})

	//p.EventsForBeats(startBeat, step, count, func(ctx ilysa.RangeContext) {
	//	ctx.WithLight(backLasers, func(ctx ilysa.TimeLightContext) {
	//		if rand.Float64() > ratio {
	//			return
	//		}
	//
	//		ctx.NewRGBLighting(ilysa.WithColor(
	//			allColorsGradient.Ierp(rand.Float64()),
	//		))
	//
	//		oe := ctx.NewRGBLighting(ilysa.WithValue(beatsaber.EventValueLightOff))
	//		oe.ShiftBeat(step / 2)
	//	})
	//	//for i := 0; i < lightCount; i++ {
	//	//	e := ctx.NewRGBLighting(backLasers, beatsaber.EventValueLightRedOn)
	//	//	e.SetSingleLightID(rand.Intn(LightIDMax) + 1)
	//	//	e.SetColor(allColorsGradient.Ierp(rand.Float64()))
	//	//
	//	//	oe := ctx.NewRGBLighting(backLasers, beatsaber.EventValueLightOff)
	//	//	oe.beat += step / 2
	//	//
	//	//	if !ctx.Last {
	//	//		continue
	//	//	}
	//	//
	//	//}
	//})
	//p.rangeTimer(startBeat+0.5, startBeat+0.5+1.2, 30, ease.Linear, func(ctx ilysa.RangeContext) {
	//	ctx.WithLight(ringLasers, func(ctx ilysa.TimeLightContext) {
	//		fx.ColorSweep(ctx, 0.4, magnetRainbow)
	//	})
	//})
	//
	//p.ModEventsInRange(startBeat+0.5, startBeat+0.5+0.6-0.001, ilysa.FilterRGBLight(ringLasers),
	//	func(ctx ilysa.RangeContext, event ilysa.Event) {
	//		fx.RGBAlphaBlend(ctx, event, 0, 1, ease.InCubic)
	//	})
	//
	//p.ModEventsInRange(startBeat+0.5+0.6, startBeat+0.5+1.2, ilysa.FilterRGBLight(ringLasers),
	//	func(ctx ilysa.RangeContext, event ilysa.Event) {
	//		fx.RGBAlphaBlend(ctx, event, 1, 0, ease.OutCirc)
	//	})
}

func (p Intro) TrillNoFade(startBeat float64) {
	ctx := p.WithBeatOffset(startBeat)

	l := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Shuffle),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(0, 0.5, 12, ease.Linear, func(ctx ilysa.RangeContext) {
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

		light = light2.TransformLight(
			light2.NewBasicLight(beatsaber.EventTypeBackLasers, p),
			ilysa.ToSequenceLightTransformer(ilysa.Divide(7)),
		)
	)

	ctx := p.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(360),
			evt.WithRotationStep(15),
			evt.WithPreciseLaserSpeed(1.3),
			evt.WithProp(13),
		)
		ctx.NewZoom()
	})

	ctx.EventsForBeats(0, step, count, func(ctx ilysa.RangeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			ctx.NewRGBLightingEvent(
				evt.WithColor(blGrad.Ierp(ctx.T())),
			)
		})

		switch {
		case ctx.Last():
			ctx.NewRGBLighting(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
				evt.WithColor(magnetPurple),
			)
			ctx.NewRGBLighting(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				evt.WithColor(magnetPurple),
			)

			exitArgs := []evt.PreciseLaserOpt{
				ilysa.WithIntValue(rotatingLasersExitSpeed),
				evt.WithLockPosition(true),
				evt.WithPreciseLaserSpeed(rotatingLasersExitSpeed),
			}

			ctx.NewPreciseLaser(
				append([]evt.PreciseLaserOpt{
					evt.WithDirectionalLaser(ilysa.LeftLaser),
					evt.WithLaserDirection(chroma.CounterClockwise),
				}, exitArgs...)...,
			)
			ctx.NewPreciseLaser(
				append([]evt.PreciseLaserOpt{
					evt.WithDirectionalLaser(ilysa.RightLaser),
					evt.WithLaserDirection(chroma.Clockwise),
				}, exitArgs...)...,
			)

		case ctx.Ordinal()%2 == 0:
			ctx.NewRGBLighting(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
				evt.WithColor(rlGrad.Ierp(ctx.T())),
			)

			ctx.NewPreciseLaser(
				evt.WithDirectionalLaser(ilysa.LeftLaser),
				ilysa.WithIntValue(ctx.Ordinal()),
				evt.WithLockPosition(true),
				evt.WithPreciseLaserSpeed(float64(ctx.Ordinal())*2),
				evt.WithLaserDirection(chroma.Clockwise),
			)

			ctx.NewRGBLighting(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightOff),
			)

		case ctx.Ordinal()%2 == 1:
			ctx.NewRGBLighting(
				ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightOff),
			)

			ctx.NewRGBLighting(
				ilysa.WithType(beatsaber.EventTypeRightRotatingLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFlash),
				evt.WithColor(rlGrad.Ierp(ctx.T())),
			)

			ctx.NewPreciseLaser(
				evt.WithDirectionalLaser(ilysa.RightLaser),
				ilysa.WithIntValue(ctx.Ordinal()),
				evt.WithLockPosition(false),
				evt.WithPreciseLaserSpeed(float64(ctx.Ordinal())*2),
				evt.WithLaserDirection(chroma.CounterClockwise),
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

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(12.5),
			evt.WithLaserDirection(chroma.CounterClockwise),
			evt.WithPreciseLaserSpeed(3),
			evt.WithProp(5),
			evt.WithCounterSpin(true),
		)
	})

	light := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Divide(7)),
	)

	ctx.EventsForRange(0, 1, 30, ease.OutCubic, func(ctx ilysa.RangeContext) {
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
	//ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
	//	ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewRGBLighting(
	//			ilysa.WithColor(allColorsGradient.Ierp(ctx.LightIDT())),
	//		)
	//	})
	//})
	//
	//ctx.Sequence(0, sequence[:len(sequence)-1], func(ctx ilysa.SequenceContext) {
	//	ctx.WithLight(light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLighting(ilysa.WithValue(beatsaber.EventValueLightOff))
	//	})
	//})
	//
	//ctx.EventForBeat(sequence[len(sequence)-1], func(ctx ilysa.RangeContext) {
	//	ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewLighting(ilysa.WithValue(beatsaber.EventValueLightOff))
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

	ctx.EventsForRange(0, 4, 60, ease.Linear, func(ctx ilysa.RangeContext) {
		ctx.WithLight(backLaser, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, sweepSpeed, grad)
			fx.AlphaBlend(ctx, e, 0, 0.5, 0, 2, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.5, 1, 2, 0, ease.OutCubic)
		})
	})

	ctx.EventsForSequence(0, sequence, func(ctx ilysa.SequenceContext) {
		if ctx.Last() {
			ctx.NewPreciseRotation(
				evt.WithRotation(360),
				evt.WithRotationStep(0),
				evt.WithLaserDirection(chroma.Clockwise),
				evt.WithPreciseLaserSpeed(7),
				evt.WithProp(0.8),
				evt.WithCounterSpin(true),
			)
		} else {
			ctx.NewPreciseRotation(
				evt.WithRotation(45),
				evt.WithRotationStep(12.5),
				evt.WithPreciseLaserSpeed(26),
				evt.WithProp(8),
				evt.WithLaserDirection(chroma.CounterClockwise),
			)
		}
		//nbo := ctx.SequenceNextBOffset()
		//if ctx.Last() {
		//	nbo = 0.75
		//}
		p.Rush(ctx, ctx.B()-0.25, ctx.B()+0.25, 0.75, float64(ctx.Ordinal()), magnetGradient)

		//g := grad.RotateRand()
		//ctx.rangeTimer(ctx.B(), ctx.B()+splashDuration, 15, ease.Linear, func(ctx ilysa.RangeContext) {
		//	ctx.WithLight(leftLaser, func(ctx ilysa.TimeLightContext) {
		//		e := fx.Gradient(ctx, g)
		//		e.Apply(ilysa.WithAlpha(float64(ctx.Ordinal()) * 0.5))
		//	})
		//	ctx.WithLight(rightLaser, func(ctx ilysa.TimeLightContext) {
		//		e := fx.Gradient(ctx, g)
		//		e.Apply(ilysa.WithAlpha(float64(ctx.Ordinal()) * 0.5))
		//	})
		//})
		//
		//ctx.ModEventsInRange(ctx.B(), ctx.B()+splashDuration, ilysa.FilterRGBLights(), func(ctx ilysa.RangeContext, event ilysa.Event) {
		//	fx.RGBAlphaBlend(ctx, event, 1, 0, ease.InCirc)
		//})
	})
}

func (p Intro) PianoGlow(ctx ilysa.BaseContext, sequence []float64, divisor int, shuffle bool) {
	lightSweepDiv := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(divisor)),
	).(light2.SequenceLight)

	if shuffle {
		lightSweepDiv = lightSweepDiv.Shuffle()
	}

	ctx.Sequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		grad := magnetRainbowPale.RotateRand()
		ctx.EventsForRange(ctx.B(), ctx.B()+0.435, 12, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaBlend(ctx, e, 0, 0.15, 0, 1, ease.OutCubic)
				fx.AlphaBlend(ctx, e, 0.3, 1, 3, 0, ease.InCubic)
			})
		})
	})
}

func (p Intro) PianoTransmute(ctx ilysa.BaseContext, sequence []float64, divisor int, shuffle bool, grad gradient.Table) {
	lightSweepDiv := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, p),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(divisor)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(light2.SequenceLight)

	if shuffle {
		lightSweepDiv = lightSweepDiv.Shuffle()
	}

	ctx.Sequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.EventsForRange(ctx.B(), ctx.B()+0.435, 12, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaBlend(ctx, e, 0, 0.3, 1, 3, ease.OutCubic)
				fx.AlphaBlend(ctx, e, 0.3, 1, 3, 1, ease.InCubic)
			})
		})
	})
}

func (p Intro) Rush(ctx ilysa.BaseContext, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	light := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeRingLights, p),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.Range(startBeat, endBeat, 45, ease.InExpo, func(ctx ilysa.RangeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, 2, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaBlend(ctx, e, 0, 0.6, 1, peakAlpha, ease.OutCubic)
			fx.AlphaBlend(ctx, e, 0.6, 1.0, peakAlpha, 0, ease.InCubic)
		})
	})
}
