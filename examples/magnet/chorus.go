package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light2"
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

		light = light2.NewCombinedLight(
			light2.TransformLight(
				light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			),
			light2.TransformLight(
				light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c),
				ilysa.ToLightTransformer(ilysa.DivideSingle),
			),
		)
	)

	ctx := c.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(evt.WithReset(true))
	})

	ctx.EventForBeat(0.5, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(15),
			evt.WithPreciseLaserSpeed(2.8),
			evt.WithProp(0.9),
			evt.WithLaserDirection(chroma.Clockwise),
		)
	})

	ctx.EventsForSequence(0, []float64{0, 0.5}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(1), evt.WithPreciseLaserSpeed(0),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(1), evt.WithPreciseLaserSpeed(0),
		)

		ctx.EventsForRange(ctx.B(), ctx.B()+0.5, 12, ease.Linear, func(ctx ilysa.RangeContext) {
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

	ll := light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c)
	rl := light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c)

	light := light2.NewSequenceLight(ll, rl)

	ctx.EventsForBeats(0, 1, 4, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(24),
			evt.WithPreciseLaserSpeed(0),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(24),
			evt.WithPreciseLaserSpeed(0),
		)

		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := ctx.NewRGBLightingEvent(
				evt.WithColor(magnetColors.Next()),
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

	ctx.EventsForRange(rippleStart, rippleEnd, 60, ease.Linear, func(ctx ilysa.RangeContext) {
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

	bl := light2.NewBasicLight(beatsaber.EventTypeBackLasers, c)

	ctx.EventsForSequence(0, []float64{0, 2, 2.75, 3.5}, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		dir := chroma.Clockwise
		if ctx.Ordinal() == 1 {
			dir = dir.Reverse()
		}

		ctx.NewPreciseRotation(
			evt.WithRotation(45*float64(ctx.Ordinal()+1)),
			evt.WithRotationStep(float64(90/(ctx.Ordinal()+1))),
			evt.WithPreciseLaserSpeed(16),
			evt.WithProp(1.3),
			evt.WithLaserDirection(dir),
		)

		eb := ctx.B() + ctx.SequenceNextBOffset()
		if ctx.Last() {
			eb = ctx.B() + 0.5
		}

		col := crossickColors.Next()
		ctx.EventsForRange(ctx.B(), eb, 18, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(bl, func(ctx ilysa.TimeLightContext) {
				e := ctx.NewRGBLightingEvent(
					evt.WithColor(col),
				)
				fx.AlphaBlend(ctx, e, 0, 1, 2+float64(seqCtx.Ordinal()), 0.3, ease.OutCubic)
			})
		})
	})
}

// キスをして
func (c Chorus) Refrain4(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(180),
			evt.WithPreciseLaserSpeed(20),
			evt.WithProp(1.2),
			evt.WithLaserDirection(chroma.Clockwise),
		)

		ctx.NewPreciseZoom(
			evt.WithRotationStep(-0.66),
		)
	})

	ll := light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c)
	rl := light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c)

	rotLasers := light2.NewCombinedLight(ll, rl)

	ctx.EventsForSequence(0, []float64{0, 2}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser),
			evt.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(30),
			evt.WithPreciseLaserSpeed(0),
		)
		ctx.WithLight(rotLasers, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(
				evt.WithColor(allColors.Next()),
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
		ctx.NewPreciseZoom(evt.WithRotationStep(0.33))
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

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(180),
			evt.WithPreciseLaserSpeed(20),
			evt.WithProp(1.2),
			evt.WithLaserDirection(chroma.Clockwise),
		)

		ctx.NewPreciseZoom(
			evt.WithRotationStep(-0.66),
		)
	})

	light := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeRingLights, c),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventsForRange(-1, 0, 45, ease.InExpo, func(ctx ilysa.RangeContext) {
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
		ctx.NewPreciseZoom(evt.WithRotationStep(-0.33))
	})

	ctx.EventsForSequence(0, []float64{2, 2.75, 3.5}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(30*float64(ctx.Ordinal())),
			evt.WithRotationStep(7*float64(ctx.Ordinal()*5)),
			evt.WithPreciseLaserSpeed(20),
			evt.WithProp(1.2),
			evt.WithLaserDirection(chroma.Clockwise),
		)
	})

	ctx.EventForBeat(3.5, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseZoom(evt.WithRotationStep(-1))
	})

	ctx.EventForBeat(4, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(12.5),
			evt.WithPreciseLaserSpeed(4),
			evt.WithProp(1),
			evt.WithLaserDirection(chroma.CounterClockwise),
		)

		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedFade),
			evt.WithColor(sukoyaWhite),
		)

	})

	bl := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, ctx),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.EventForBeat(5, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(360),
			evt.WithRotationStep(0),
			evt.WithPreciseLaserSpeed(1.6),
			evt.WithProp(0.9),
			evt.WithLaserDirection(chroma.Clockwise),
		)
		ctx.NewPreciseZoom(evt.WithRotationStep(1.5))

		ctx.Range(ctx.B(), ctx.B()+4, 60, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(bl, func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 1.2, magnetRainbow)
				fx.AlphaBlend(ctx, e, 0, 1, 6, 0.0, ease.InOutQuad)
			})
		})
	})
}

func (c Chorus) Rhythm2(startBeat float64) {
	ctx := c.WithBeatOffset(startBeat)

	ll := light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, c)
	rl := light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, c)

	light := light2.NewSequenceLight(ll, rl)

	ctx.EventsForBeats(0, 1, 4, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(24),
			evt.WithPreciseLaserSpeed(0),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(24),
			evt.WithPreciseLaserSpeed(0),
		)

		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := ctx.NewRGBLightingEvent(
				evt.WithColor(magnetColors.Next()),
				ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			)
			e.SetAlpha(0.3)
		})
	})
}

func (c Chorus) Sweep(ctx ilysa.BaseContext, startBeat, endBeat float64, grad gradient.Table, reverse bool) {
	backLasers := light2.TransformLight(light2.NewBasicLight(beatsaber.EventTypeBackLasers, c),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
	)

	if reverse {
		backLasers = light2.TransformLight(backLasers,
			ilysa.ToLightTransformer(ilysa.Reverse),
		)
	}

	lightSweep := light2.TransformLight(
		backLasers,
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	lightSweepSeq := light2.TransformLight(
		backLasers,
		ilysa.ToSequenceLightTransformer(ilysa.DivideSingle),
	)

	ctx.Range(startBeat, endBeat, 30, ease.OutCubic, func(ctx ilysa.RangeContext) {
		ctx.WithLight(lightSweepSeq, func(ctx ilysa.TimeLightContext) {
			ctx.NewRGBLightingEvent(evt.WithColor(grad.Ierp(ctx.T())))
		})
	})

	ctx.EventForBeat(endBeat+0.01, func(ctx ilysa.RangeContext) {
		ctx.WithLight(lightSweep, func(ctx ilysa.TimeLightContext) {
			e := fx.Gradient(ctx, grad.Reverse())
			e.Mod(evt.WithAlpha(2))
		})
	})
}

func (c Chorus) SweepSpin(ctx ilysa.BaseContext, startBeat float64, reverse bool) {
	dir := chroma.CounterClockwise.ReverseIf(reverse)

	ctx.EventForBeat(startBeat, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(135),
			evt.WithRotationStep(12),
			evt.WithPreciseLaserSpeed(8), // 2.1
			evt.WithProp(0.8),            // 0.9
			evt.WithLaserDirection(dir),
		)
	})
}

func (c Chorus) HalfCollapse(ctx ilysa.BaseContext, startBeat float64, sequence []float64, reverse bool) {
	dir := chroma.Clockwise.ReverseIf(reverse)

	ctx.Sequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(9),
			evt.WithPreciseLaserSpeed(8),
			evt.WithProp(0.8),
			evt.WithLaserDirection(dir),
		)
	})

}

func (c Chorus) Unsweep(ctx ilysa.BaseContext, startBeat float64, sequence []float64) {
	ctx.Sequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(8),
			evt.WithPreciseLaserSpeed(20),
			evt.WithProp(1.2),
			evt.WithLaserDirection(chroma.CounterClockwise),
		)
	})
}

func (c Chorus) Rush(ctx ilysa.BaseContext, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	light := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeRingLights, c),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.Range(startBeat, endBeat, 45, ease.InExpo, func(ctx ilysa.RangeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, 2, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaBlend(ctx, e, 0, 0.6, 2, peakAlpha, ease.OutCubic)
			fx.AlphaBlend(ctx, e, 0.6, 1.0, peakAlpha, 0, ease.OutCubic)
		})
	})
}

func (c Chorus) RushNoFade(ctx ilysa.BaseContext, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	light := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeRingLights, c),
		ilysa.ToLightTransformer(ilysa.Reverse),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	ctx.Range(startBeat, endBeat, 45, ease.InExpo, func(ctx ilysa.RangeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := fx.BiasedColorSweep(ctx, 4, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaBlend(ctx, e, 0, 0.6, 2, peakAlpha, ease.OutCubic)
		})
	})
}

func (c Chorus) FadeToSukoya(ctx ilysa.BaseContext, startBeat float64, sequence []float64) {
	lightSweepDiv := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, c),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(len(sequence))),
	).(light2.SequenceLight)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.Sequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.EventsForRange(ctx.B(), ctx.B()+0.5, 16, ease.OutCubic, func(ctx ilysa.RangeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 3.6, gradient.FromSet(sukoyaColors))
				fx.AlphaBlend(ctx, e, 0, 1, 0.3, 3, ease.OutElastic)
			})
		})
	})
}

func (c Chorus) FadeToGold(ctx ilysa.BaseContext, startBeat float64, sequence []float64) {
	lightSweepDiv := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, c),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToLightTransformer(ilysa.Rotate(3)),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(len(sequence))),
	).(light2.SequenceLight)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.Sequence(startBeat, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		grad := magnetRainbowPale.RotateRand()
		ctx.EventsForRange(ctx.B(), ctx.B()+0.5, 16, ease.OutCubic, func(ctx ilysa.RangeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaBlend(ctx, e, 0, 1, 10, 0, ease.OutElastic)
			})
		})
	})
}
