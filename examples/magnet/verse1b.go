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

type Verse1b struct {
	ilysa.BaseContext
	offset float64
}

func NewVerse1b(project *ilysa.Project, offset float64) Verse1b {
	return Verse1b{
		BaseContext: project.WithBeatOffset(offset),
	}
}

func (v Verse1b) Play() {
	v.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(5),
		)
		ctx.NewLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(5),
		)
	})

	v.Rhythm(0, false)
	v.Rhythm(4, true)
	v.Rhythm(8, true)
	v.Rhythm(12, true)
	v.Rhythm(16, true)
	v.Rhythm(20, false)

	v.IntroBridge(0)
	v.PianoRoll(6)
	v.Stinger(12)
	v.Naosara(22)
	//v.Moeagaru(22.5)
	v.Moeagaru(26.5)
}

func (v Verse1b) IntroBridge(startBeat float64) {
	ctx := v.WithBeatOffset(startBeat)

	ctx.EventsForSequence(0, []float64{0, 1, 3}, func(ctx ilysa.SequenceContext) {
		var (
			step      = []float64{45, 90, 180}
			rotation  = []float64{45, 45, 180}
			speed     = []float64{9, 9, 15}
			prop      = []float64{1.2, 1.2, 6}
			direction = []chroma.SpinDirection{chroma.CounterClockwise, chroma.CounterClockwise, chroma.Clockwise}
		)

		ctx.NewPreciseRotation(
			evt.WithRotation(rotation[ctx.Ordinal()]),
			evt.WithRotationStep(step[ctx.Ordinal()]),
			evt.WithPreciseLaserSpeed(speed[ctx.Ordinal()]),
			evt.WithProp(prop[ctx.Ordinal()]),
			evt.WithLaserDirection(direction[ctx.Ordinal()]),
		)
	})

	ctx.EventsForSequence(0, []float64{0, 1, 2}, func(ctx ilysa.SequenceContext) {
		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
			evt.WithColor(crossickColors.Next()),
			evt.WithAlpha(3),
		)
	})

	ctx.EventForBeat(3, func(ctx ilysa.RangeContext) {
		ctx.NewRGBLighting(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightOff),
		)
	})
}

func (v Verse1b) Rhythm(startBeat float64, spin bool) {
	var (
		leftLaser  = light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, v)
		rightLaser = light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, v)
		beatLight  = light2.NewSequenceLight(leftLaser, rightLaser)
		color      = crossickColors
	)

	v.Sequence(startBeat, []float64{1, 3}, func(ctx ilysa.SequenceContext) {
		if spin {
			ctx.NewPreciseRotation(
				evt.WithRotation(90),
				evt.WithRotationStep(11.25),
				evt.WithPreciseLaserSpeed(7),
				evt.WithProp(3),
				evt.WithLaserDirection(chroma.CounterClockwise),
			)
		}

		ctx.WithLight(beatLight, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				evt.WithColor(color.Next()),
			)
		})
	})

	var (
		rippleDuration = 1.0
		rippleStart    = startBeat + 2
		rippleEnd      = rippleStart + rippleDuration
		rippleLights   = v.NewBasicLight(beatsaber.EventTypeRingLights).Transform(ilysa.DivideSingle)
		rippleStep     = 0.8
	)

	v.Range(rippleStart, rippleEnd, 30, ease.Linear, func(ctx ilysa.RangeContext) {
		ctx.WithLight(rippleLights, func(ctx ilysa.TimeLightContext) {
			e := fx.ColorSweep(ctx, 2, gradient.FromSet(crossickColors))

			fx.Ripple(ctx, e, rippleStep)
			fx.AlphaBlend(ctx, e, 0, 0.2, 0, 2, ease.InCubic)
			fx.AlphaBlend(ctx, e, 0.8, 1, 2, 0, ease.OutCubic)
		})
	})
}

func (v Verse1b) Stinger(startBeat float64) {
	var (
		seq1 = []float64{0.25, 0.50, 0.75, 1.00, 1.25, 1.50}
		seq2 = []float64{2.00, 2.25, 2.50, 2.75, 3.00, 3.50}
	)
	ctx := v.WithBeatOffset(startBeat)

	v.PianoHits(ctx, seq1, 6, 0.6)
	v.PianoHits(ctx, seq2, 6, 0.6)

	//	backLasers = ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v)
	//	//seq1Light  = backLasers
	//	seq1Light = ilysa.TransformLight(backLasers,
	//		ilysa.ToLightTransformer(ilysa.Fan(2)),
	//		ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//		ilysa.ToLightTransformer(ilysa.DivideSingle),
	//	).(ilysa.SequenceLight)
	//	seq2Light = seq1Light
	//
	//	seq1Colors = allColors
	//	seq2Colors = allColors
	//)
	//
	//ctx := v.WithBeatOffset(startBeat)
	//
	//ctx.Sequence(0, seq1, func(ctx ilysa.SequenceContext) {
	//	if ctx.Ordinal() == 0 {
	//		ctx.NewPreciseZoom(ilysa.WithRotationStep(-0.9))
	//	}
	//
	//	ctx.WithLight(seq1Light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLighting(ilysa.WithColor(seq1Colors.Rand()))
	//
	//		oe := ctx.NewRGBLighting(ilysa.WithValue(beatsaber.EventValueLightOff))
	//		if !ctx.Last() {
	//			oe.ShiftBeat(ctx.SequenceNextBOffset())
	//		} else {
	//			oe.ShiftBeat(0.25)
	//		}
	//	})
	//
	//})
	//
	//ctx.Sequence(0, seq2, func(ctx ilysa.SequenceContext) {
	//	if ctx.Ordinal() == 0 {
	//		ctx.NewPreciseZoom(ilysa.WithRotationStep(0.9))
	//	}
	//
	//	ctx.WithLight(seq2Light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLighting(ilysa.WithColor(seq2Colors.Rand()))
	//
	//		oe := ctx.NewRGBLighting(ilysa.WithValue(beatsaber.EventValueLightOff))
	//		if !ctx.Last() {
	//			oe.ShiftBeat(ctx.SequenceNextBOffset())
	//		} else {
	//			oe.ShiftBeat(0.25)
	//		}
	//	})
	//})
}

func (v Verse1b) PianoRoll(startBeat float64) {
	//ctx := v.WithBeatOffset(startBeat)
	//
	//bl := ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v)
	//
	//blForward := ilysa.TransformLight(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//).(ilysa.SequenceLight)
	//
	//blReverse := ilysa.TransformLight(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//	ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	//).(ilysa.SequenceLight)
	//
	//light := ilysa.NewSequenceLight(
	//	blForward.Index(0), blReverse.Index(1), blForward.Index(0),
	//)
	//
	//ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
	//	ctx.NewPreciseRotation(
	//		ilysa.WithRotation(90),
	//		ilysa.WithRotationStep(8),
	//		ilysa.WithPreciseLaserSpeed(8),
	//		ilysa.WithProp(6),
	//		ilysa.WithLaserDirection(chroma.Clockwise),
	//	)
	//
	//})
	//
	//ctx.Sequence(0, []float64{0, 0.25, 0.50}, func(ctx ilysa.SequenceContext) {
	//	seqCtx := ctx
	//	ctx.rangeTimer(ctx.B(), ctx.B()+0.25, 8, ease.Linear, func(ctx ilysa.RangeContext) {
	//		ctx.WithLight(light.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
	//			e := ctx.NewRGBLighting(ilysa.WithColor(magnetColors.Next()))
	//			fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.OutCubic)
	//		})
	//	})
	//})
	//
	//ctx.EventForBeat(0.75, func(ctx ilysa.RangeContext) {
	//	ctx.WithLight(bl, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewRGBLighting(
	//			ilysa.WithValue(beatsaber.EventValueLightRedFade),
	//			ilysa.WithColor(magnetWhite),
	//		)
	//	})
	//})
	var (
		sequence = []float64{0, 0.25, 0.50, 0.75}
	)

	ctx := v.WithBeatOffset(startBeat)
	v.PianoHits(ctx, sequence, 5, 1)
}

func (v Verse1b) Naosara(startBeat float64) {
	ctx := v.WithBeatOffset(startBeat)

	//ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
	//	ctx.NewPreciseRotation(ilysa.WithReset(true))
	//})

	ctx.EventsForSequence(0, []float64{0.4, 0.9, 1.4}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(6),
			evt.WithPreciseLaserSpeed(3),
			evt.WithProp(20),
			evt.WithLaserDirection(chroma.CounterClockwise),
		)
	})
	ctx.EventForBeat(1.9, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(11.5),
			evt.WithPreciseLaserSpeed(0.9),
			evt.WithProp(0.8),
			evt.WithLaserDirection(chroma.Clockwise),
		)
	})

	v.Burn(ctx, 0.0, 0.50, magnetGradient, 0, false)
	v.Burn(ctx, 0.5, 1.00, magnetGradient, 1, false)
	v.Burn(ctx, 1.0, 1.50, sukoyaGradient, 0, false)
	v.Burn(ctx, 1.5, 2.00, shirayukiGradient, 1, false)

}

func (v Verse1b) Moeagaru(startBeat float64) {
	ctx := v.WithBeatOffset(startBeat)

	v.PianoHits(ctx, []float64{0, 0.5, 1.0, 1.5}, 4, 0.5)
	//v.PianoHits(ctx, []float64{3}, 1, 0.25)

	lightSweepDiv := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, v),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		//ilysa.ToSequenceLightTransformer(ilysa.Divide(divisor)),
	)

	//lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.EventsForRange(3, 3.25, 12, ease.OutCubic, func(ctx ilysa.RangeContext) {
		ctx.WithLight(lightSweepDiv, func(ctx ilysa.TimeLightContext) {
			fx.ColorSweep(ctx, 12, magnetRainbowPale)
		})
	})

	ctx.EventForBeat(3.25, func(ctx ilysa.RangeContext) {
		ctx.NewRGBLighting(ilysa.WithType(beatsaber.EventTypeBackLasers), ilysa.WithValue(beatsaber.EventValueLightOff))
	})

	//bl := ilysa.TransformLight(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//).(ilysa.SequenceLight)
	//
	//blReverse := ilysa.TransformLight(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//	ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	//).(ilysa.SequenceLight)
	//
	////lightReverse:= ilysa.TransformLight(
	////	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	////	ilysa.ToLightTransformer(ilysa.DivideSingle),
	////	ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	////)
	//
	////grad := gradient.New(
	////	magnetColors.Index(0),
	////	magnetColors.Index(1),
	////)
	//
	//ctx.EventsForBeats(0, 0.5, 4, func(ctx ilysa.RangeContext) {
	//	var direction chroma.SpinDirection
	//	if ctx.Ordinal()%2 == 0 {
	//		direction = chroma.CounterClockwise
	//	} else {
	//		direction = chroma.Clockwise
	//	}
	//	ctx.NewPreciseRotation(
	//		ilysa.WithRotation(15+float64(ctx.Ordinal())*5),
	//		ilysa.WithRotationStep(12.5),
	//		ilysa.WithPreciseLaserSpeed(8),
	//		ilysa.WithProp(4),
	//		ilysa.WithLaserDirection(direction),
	//	)
	//})
	//
	//grad := gradient.Table{
	//	{Col: magnetColors.Index(0), Pos: 0.0},
	//	{Col: magnetColors.Index(1), Pos: 0.6},
	//	{Col: magnetColors.Index(1), Pos: 1.0},
	//}
	//
	//var (
	//	duration = 0.5
	//	frames   = 10
	//	step     = 0.3
	//)
	//
	//ctx.EventsForBeats(0, 0.5, 4, func(ctx ilysa.RangeContext) {
	//	seqCtx := ctx
	//
	//	var light ilysa.SequenceLight
	//
	//	if ctx.Ordinal()%2 == 0 {
	//		light = bl
	//	} else {
	//		light = blReverse
	//	}
	//	ctx.rangeTimer(ctx.B(), ctx.B()+duration, frames, ease.Linear, func(ctx ilysa.RangeContext) {
	//		ctx.WithLight(light.Index(seqCtx.Ordinal()+1), func(ctx ilysa.TimeLightContext) {
	//			e := fx.Gradient(ctx, magnetGradient)
	//			e.SetAlpha(1 + float64(seqCtx.Ordinal())*4)
	//			fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.OutSine)
	//
	//		})
	//		ctx.WithLight(light.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
	//			e := ctx.NewRGBLighting(
	//				ilysa.WithColor(grad.Ierp(ctx.T())),
	//			)
	//			fx.Ripple(ctx, e, step)
	//			fx.AlphaBlend(ctx, e, 0.0, 0.4, 0, 1, ease.InCirc)
	//			fx.AlphaBlend(ctx, e, 0.4, 1, 1, 0, ease.OutSine)
	//		})
	//	})
	//})

	//fx.RainbowProp(ctx, light, grad, 0, 0.2, 0.3, 10)
	//fx.RainbowProp(ctx, light, grad, 0.2, 0.2, 0.8, 10)
	//fx.RainbowProp(ctx, light, grad, 1, 0.2, 0.3, 10)
	//fx.RainbowProp(ctx, light, grad, 1.5, 0.2, 0.3, 10)
	//fx.RainbowProp(ctx, light, magnetRainbow, 0.5, 0.25, 2.4, 8)
	//fx.RainbowProp(ctx, light, magnetRainbow, 1.0, 0.25, 2.4, 8)
	//fx.RainbowProp(ctx, light, magnetRainbow, 1.5, 0.25, 2.4, 8)

}

func (v Verse1b) PianoHits(ctx ilysa.BaseContext, sequence []float64, divisor int, duration float64) {
	lightSweepDiv := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, v),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(divisor)),
	).(light2.SequenceLight)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.Sequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.EventsForRange(ctx.B(), ctx.B()+duration, 16, ease.OutCubic, func(ctx ilysa.RangeContext) {
			ctx.WithLight(lightSweepDiv.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 3.6, magnetRainbowPale)
				alpha := 3.0
				if seqCtx.Last() {
					alpha = 8
				}
				fx.AlphaBlend(ctx, e, 0, 1, alpha, 0, ease.OutElastic)
			})
		})
	})
}

func (v Verse1b) Burn(ctx ilysa.BaseContext, startBeat, endBeat float64, grad gradient.Table, ordinal int, reverse bool) {
	reverseTransform := ilysa.ToLightTransformer(ilysa.Identity)
	if ordinal == 1 {
		reverseTransform = ilysa.ToLightTransformer(ilysa.Reverse)
	}

	backLasers := light2.TransformLight(light2.NewBasicLight(beatsaber.EventTypeBackLasers, v),
		ilysa.ToLightTransformer(ilysa.Fan(2)),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Flatten),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(2)),
		reverseTransform,
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(light2.SequenceLight)

	ctx.Range(startBeat, endBeat, 15, ease.InExpo, func(ctx ilysa.RangeContext) {
		ctx.WithLight(backLasers.Index(ordinal), func(ctx ilysa.TimeLightContext) {
			e := ctx.NewRGBLightingEvent(evt.WithColor(grad.Ierp(ctx.T())))
			fx.Ripple(ctx, e, 0.5)
		})
	})
}
