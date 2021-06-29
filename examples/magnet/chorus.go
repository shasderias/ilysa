package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/fx"
)

type Chorus struct {
	ilysa.BaseContext
	offset float64
}

func NewChorus(project *ilysa.Project, startBeat float64) Chorus {
	return Chorus{
		BaseContext: project.WithBeatOffset(startBeat),
	}
}

func (c Chorus) Play() {
	c.RhythmBridge(0)

	c.Rhythm(2)
	c.Rhythm(6)
	c.Rhythm(10)

	c.Refrain1(1)
	c.Refrain2(4.5)
	c.Refrain3(8.5)
	c.Climax(14)
	c.Refrain4(18)
	c.Refrain5(20.75)
	c.Refrain6(24.50)
	c.Refrain7(30)

	c.Rhythm2(18)
	c.Rhythm2(22)
	c.Rhythm2(26)
	c.Rhythm2(30)
}

func (c Chorus) RhythmBridge(startBeat float64) {
	var (
		grad = magnetRainbow

		light = ilysa.NewCombinedLight(
			ilysa.TransformLight(
				ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			),
			ilysa.TransformLight(
				ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			),
		)
	)

	ctx := c.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(ilysa.WithReset(true))
	})

	ctx.EventForBeat(0.5, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(180),
			ilysa.WithStep(15),
			ilysa.WithSpeed(2.8),
			ilysa.WithProp(0.9),
			ilysa.WithDirection(chroma.Clockwise),
		)
	})

	ctx.EventsForSequence(0, []float64{0, 0.5}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(1), ilysa.WithSpeed(0),
		)
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(1), ilysa.WithSpeed(0),
		)

		ctx.EventsForRange(ctx.B(), ctx.B()+0.5, 12, ease.Linear, func(ctx ilysa.TimeContext) {
			if ctx.Ordinal() == 1 {
				grad = grad.Reverse()
			}
			ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
				e := fx.Gradient(ctx, grad)
				fx.AlphaBlend(ctx, e, 0, 1, 6, 0, ease.OutCubic)
			})
		})
	})
}

func (c Chorus) Rhythm(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	ll := ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c)
	rl := ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c)

	light := ilysa.NewSequenceLight(ll, rl)

	ctx.EventsForBeats(0, 1, 4, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(24),
			ilysa.WithSpeed(0),
		)
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(24),
			ilysa.WithSpeed(0),
		)

		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := ctx.NewRGBLightingEvent(
				ilysa.WithColor(magnetColors.Next()),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			)
			e.SetAlpha(0.3)
		})
	})

	var (
		rippleDuration = 4.0
		rippleStart    = 0.0
		rippleEnd      = rippleStart + rippleDuration
		rippleLights   = c.NewBasicLight(beatsaber.EventTypeRingLights).Transform(ilysa.DivideSingle)
		rippleStep     = 0.8
		grad           = gradient.Table{
			{shirayukiPurple, 0.0},
			{sukoyaWhite, 0.3},
			{sukoyaWhite, 0.7},
			{shirayukiPurple, 1.0},
		}
	)

	ctx.EventsForRange(rippleStart, rippleEnd, 60, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(rippleLights, func(ctx ilysa.TimeLightContext) {
			events := fx.ColorSweep(ctx, 1.2, grad)

			fx.Ripple(ctx, events, rippleStep,
				fx.WithAlphaBlend(0, 0.2, 0, 0.3, ease.InCubic),
				fx.WithAlphaBlend(0.2, 1.0, 0.3, 0, ease.OutCubic),
			)
		})
	})
}

func (c Chorus) Refrain1(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)
	c.Sweep(ctx, 0.25, 1, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 1, false)
	c.FadeToGold(ctx, 1, []float64{0.5, 0.75, 1.25, 1.75})
}

func (c Chorus) Refrain2(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)
	c.Sweep(ctx, 0, 1.25, sukoyaGradient, true)
	c.SweepSpin(ctx, 1.25, true)
	c.FadeToGold(ctx, 1.25, []float64{0.5, 1.0, 1.5, 2.0})
}

func (c Chorus) Refrain3(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	c.Sweep(ctx, 0, 1.25, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 1.25, false)
	c.FadeToSukoya(ctx, 1.25, []float64{0, 0.5, 1.0, 1.5})
	c.HalfCollapse(ctx, 3, []float64{0}, false)

	collapseSeq := []float64{2.5, 3.0, 3.5, 4.0}

	//c.HalfCollapse(ctx, 1.0, collapseSeq, false)
	c.FadeToGold(ctx, 1.00, collapseSeq)
}

func (c Chorus) Climax(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	bl := ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, c)

	ctx.EventsForSequence(0, []float64{0, 2, 2.75, 3.5}, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		dir := chroma.Clockwise
		if ctx.Ordinal() == 1 {
			dir = dir.Reverse()
		}

		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45*float64(ctx.Ordinal()+1)),
			ilysa.WithStep(float64(90/(ctx.Ordinal()+1))),
			ilysa.WithSpeed(16),
			ilysa.WithProp(1.3),
			ilysa.WithDirection(dir),
		)

		eb := ctx.B() + ctx.NextBOffset()
		if ctx.Last() {
			eb = ctx.B() + 0.5
		}

		col := crossickColors.Next()
		ctx.EventsForRange(ctx.B(), eb, 18, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(bl, func(ctx ilysa.TimeLightContext) {
				e := ctx.NewRGBLightingEvent(
					ilysa.WithColor(col),
				)
				fx.AlphaBlend(ctx, e, 0, 1, 2+float64(seqCtx.Ordinal()), 0.3, ease.OutCubic)
			})
		})
	})
}

// キスをして
func (c Chorus) Refrain4(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(180),
			ilysa.WithStep(180),
			ilysa.WithSpeed(20),
			ilysa.WithProp(1.2),
			ilysa.WithDirection(chroma.Clockwise),
		)

		ctx.NewPreciseZoomEvent(
			ilysa.WithStep(-0.66),
		)
	})

	ll := ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c)
	rl := ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c)

	rotLasers := ilysa.NewCombinedLight(ll, rl)

	ctx.EventsForSequence(0, []float64{0, 2}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(30),
			ilysa.WithSpeed(0),
		)
		ctx.WithLight(rotLasers, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithColor(allColors.Next()),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
			)
		})
	})

	c.Rush(ctx, -1, 0, 1.5, 10, gradient.FromSet(sukoyaColors))
	// キスを
	c.Sweep(ctx, 0, 0.75, magnetGradient, false)

	// して
	c.Unsweep(ctx, 0, []float64{1.25, 1.75})
	c.FadeToGold(ctx, 0, []float64{1.25, 1.75})
	ctx.EventsForSequence(0, []float64{1.25, 1.75}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseZoomEvent(ilysa.WithStep(0.33))
	})
}

func (c Chorus) Refrain5(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	c.Sweep(ctx, 0, 0.50, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 0.5, false)
	c.FadeToGold(ctx, -0.5, []float64{1.25, 1.75, 2.25, 2.75, 3.25})
}

func (c Chorus) Refrain6(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	c.Sweep(ctx, 0, 0.75, sukoyaGradient, true)
	c.SweepSpin(ctx, 0.75, true)
	c.HalfCollapse(ctx, 2.75, []float64{0}, true)
	c.FadeToSukoya(ctx, -0.25, []float64{1.25, 2.00, 2.50, 3.00})
	c.FadeToGold(ctx, -0.50, []float64{4.0, 4.5, 5.0, 5.5})
}

func (c Chorus) Refrain7(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(180),
			ilysa.WithStep(180),
			ilysa.WithSpeed(20),
			ilysa.WithProp(1.2),
			ilysa.WithDirection(chroma.Clockwise),
		)

		ctx.NewPreciseZoomEvent(
			ilysa.WithStep(-0.66),
		)
	})

	light := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeRingLights, c),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(-1, 0, 45, ease.InExpo, func(ctx ilysa.TimeContext) {
		grad := shirayukiGradient
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.BiasedColorSweep(ctx, 4, grad)
			fx.Ripple(ctx, e, 1.5)
			fx.AlphaBlend(ctx, e, 0, 0.6, 2, 10, ease.OutCubic)
			fx.AlphaBlend(ctx, e, 0.6, 1.0, 10, 0, ease.OutCubic)
		})
	})

	ctx.EventsForSequence(-1, []float64{2, 2.75, 3.5, 4.0, 4.5, 5.5}, func(ctx ilysa.SequenceContext) {
		grad := sukoyaGradient
		c.Rush(ctx, ctx.B(), ctx.B()+0.4, 0.4, 2*float64(ctx.Ordinal()), grad)
		//c.RushNoFade(ctx, ctx.B(), ctx.B()+0.4, 0.4, 2*float64(ctx.Ordinal()), grad)
		ctx.NewPreciseZoomEvent(ilysa.WithStep(-0.33))
	})

	ctx.EventsForSequence(0, []float64{2, 2.75, 3.5}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(30*float64(ctx.Ordinal())),
			ilysa.WithStep(7*float64(ctx.Ordinal()*5)),
			ilysa.WithSpeed(20),
			ilysa.WithProp(1.2),
			ilysa.WithDirection(chroma.Clockwise),
		)
	})

	ctx.EventForBeat(3.5, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseZoomEvent(ilysa.WithStep(-1))
	})

	ctx.EventForBeat(4, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(12.5),
			ilysa.WithSpeed(4),
			ilysa.WithProp(1),
			ilysa.WithDirection(chroma.CounterClockwise),
		)

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedFade),
			ilysa.WithColor(sukoyaWhite),
		)

	})

	bl := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, ctx),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventForBeat(5, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(360),
			ilysa.WithStep(0),
			ilysa.WithSpeed(1.6),
			ilysa.WithProp(0.9),
			ilysa.WithDirection(chroma.Clockwise),
		)
		ctx.NewPreciseZoomEvent(ilysa.WithStep(1.5))

		ctx.EventsForRange(ctx.B(), ctx.B()+4, 60, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(bl, func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 1.2, magnetRainbow)
				fx.AlphaBlend(ctx, e, 0, 1, 6, 0.0, ease.InOutQuad)
			})
		})
	})
}

func (c Chorus) Rhythm2(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	ll := ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c)
	rl := ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c)

	light := ilysa.NewSequenceLight(ll, rl)

	ctx.EventsForBeats(0, 1, 4, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(24),
			ilysa.WithSpeed(0),
		)
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(24),
			ilysa.WithSpeed(0),
		)

		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := ctx.NewRGBLightingEvent(
				ilysa.WithColor(magnetColors.Next()),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			)
			e.SetAlpha(0.3)
		})
	})
}

func (c Chorus) Sweep(ctx ilysa.BaseContext, startBeat, endBeat float64, grad gradient.Table, reverse bool) {
	backLasers := ilysa.TransformLight(ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, c),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
	)

	if reverse {
		backLasers = ilysa.TransformLight(backLasers,
			ilysa.ToLightTransformer(ilysa.Reverse),
		)
	}

	lightSweep := ilysa.TransformLight(
		backLasers,
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	lightSweepSeq := ilysa.TransformLight(
		backLasers,
		ilysa.ToSequenceLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(startBeat, endBeat, 30, ease.OutCubic, func(ctx ilysa.TimeContext) {
		ctx.WithLight(lightSweepSeq, func(ctx ilysa.TimeLightContext) {
			ctx.NewRGBLightingEvent(ilysa.WithColor(grad.Ierp(ctx.T())))
		})
	})

	ctx.EventForBeat(endBeat+0.01, func(ctx ilysa.TimeContext) {
		ctx.WithLight(lightSweep, func(ctx ilysa.TimeLightContext) {
			e := fx.Gradient(ctx, grad.Reverse())
			e.Mod(ilysa.WithAlpha(2))
		})
	})
}

func (c Chorus) SweepSpin(ctx ilysa.BaseContext, startBeat float64, reverse bool) {
	dir := chroma.CounterClockwise.ReverseIf(reverse)

	ctx.EventForBeat(startBeat, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(135),
			ilysa.WithStep(12),
			ilysa.WithSpeed(8),  // 2.1
			ilysa.WithProp(0.8), // 0.9
			ilysa.WithDirection(dir),
		)
	})
}

func (c Chorus) HalfCollapse(ctx ilysa.BaseContext, startBeat float64, sequence []float64, reverse bool) {
	dir := chroma.Clockwise.ReverseIf(reverse)

	ctx.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(9),
			ilysa.WithSpeed(8),
			ilysa.WithProp(0.8),
			ilysa.WithDirection(dir),
		)
	})

}

func (c Chorus) Unsweep(ctx ilysa.BaseContext, startBeat float64, sequence []float64) {
	ctx.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45),
			ilysa.WithStep(8),
			ilysa.WithSpeed(20),
			ilysa.WithProp(1.2),
			ilysa.WithDirection(chroma.CounterClockwise),
		)
	})
}

func (c Chorus) Rush(ctx ilysa.BaseContext, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	light := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeRingLights, c),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(startBeat, endBeat, 45, ease.InExpo, func(ctx ilysa.TimeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, 2, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaBlend(ctx, e, 0, 0.6, 2, peakAlpha, ease.OutCubic)
			fx.AlphaBlend(ctx, e, 0.6, 1.0, peakAlpha, 0, ease.OutCubic)
		})
	})
}

func (c Chorus) RushNoFade(ctx ilysa.BaseContext, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	light := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeRingLights, c),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(startBeat, endBeat, 45, ease.InExpo, func(ctx ilysa.TimeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.BiasedColorSweep(ctx, 4, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaBlend(ctx, e, 0, 0.6, 2, peakAlpha, ease.OutCubic)
		})
	})
}

func (c Chorus) FadeToSukoya(ctx ilysa.BaseContext, startBeat float64, sequence []float64) {
	lightSweepDiv := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, c),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(len(sequence))),
	).(ilysa.SequenceLight)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.EventsForRange(ctx.B(), ctx.B()+0.5, 16, ease.OutCubic, func(ctx ilysa.TimeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 3.6, gradient.FromSet(sukoyaColors))
				fx.AlphaBlend(ctx, e, 0, 1, 0.3, 3, ease.OutElastic)
			})
		})
	})
}

func (c Chorus) FadeToGold(ctx ilysa.BaseContext, startBeat float64, sequence []float64) {
	lightSweepDiv := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, c),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToLightTransformer(ilysa.Rotate(3)),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(len(sequence))),
	).(ilysa.SequenceLight)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.EventsForSequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		grad := magnetRainbowPale.RotateRand()
		ctx.EventsForRange(ctx.B(), ctx.B()+0.5, 16, ease.OutCubic, func(ctx ilysa.TimeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaBlend(ctx, e, 0, 1, 10, 0, ease.OutElastic)
			})
		})
	})
}
