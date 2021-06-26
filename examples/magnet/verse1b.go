package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/fx"
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
	v.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser), ilysa.WithIntValue(5),
		)
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser), ilysa.WithIntValue(5),
		)
	})

	v.Rhythm(0)
	v.Rhythm(4)
	v.Rhythm(8)
	v.Rhythm(12)
	v.Rhythm(16)
	v.Rhythm(20)

	v.IntroBridge(0)
	v.PianoRoll(6)
	v.Stinger(12)
	v.Moeagaru(22.5)
	v.Moeagaru(26.5)
}

func (v Verse1b) IntroBridge(startBeat float64) {
	ctx := v.WithBeatOffset(startBeat)

	ctx.EventsForSequence(0, []float64{0, 1, 2}, func(ctx ilysa.SequenceContext) {
		step := []float64{90, 105}

		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(step[ctx.Ordinal()%2]),
			ilysa.WithSpeed(4),
			ilysa.WithProp(2),
			ilysa.WithDirection(chroma.CounterClockwise),
		)

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
			ilysa.WithColor(crossickColors.Next()),
			ilysa.WithAlpha(4),
		)
	})
}

func (v Verse1b) Rhythm(startBeat float64) {
	var (
		leftLaser  = ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, v)
		rightLaser = ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, v)
		beatLight  = ilysa.NewSequenceLight(leftLaser, rightLaser)
		color      = crossickColors
	)

	v.EventsForSequence(startBeat, []float64{0, 1, 2, 3}, func(ctx ilysa.SequenceContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(22.5),
			ilysa.WithSpeed(4),
			ilysa.WithProp(2),
			ilysa.WithDirection(chroma.CounterClockwise),
		)

		ctx.WithLight(beatLight, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(color.Next()),
			)
		})
	})

	var (
		rippleDuration = 1.0
		rippleStart    = startBeat + 2
		rippleEnd      = rippleStart + rippleDuration
		rippleLights   = v.NewBasicLight(beatsaber.EventTypeRingLights).Transform(ilysa.DivideSingle)
		rippleStep     = 0.8
		grad           = gradient.Table{
			{magnetColors.Index(0), 0.0},
			{magnetColors.Index(1), 0.05},
			{magnetColors.Index(2), 0.5},
			{magnetColors.Index(3), 1.0},
		}
	)

	v.EventsForRange(rippleStart, rippleEnd, 30, ease.Linear, func(ctx ilysa.TimeContext) {
		ctx.WithLight(rippleLights, func(ctx ilysa.TimeLightContext) {
			events := fx.BiasedColorSweep(ctx, 3, grad)

			fx.Ripple(ctx, events, rippleStep,
				fx.WithAlphaBlend(0, 0.2, 0, 2, ease.InCubic),
				fx.WithAlphaBlend(0.8, 1, 2, 0, ease.OutCubic),
			)
			//events.Mod(ilysa.WithAlpha(1.5))
			//for _, ee := range *events {
			//	ee.ShiftBeat(ctx.LightIDT() * rippleStep)
			//}
			//switch {
			//case ctx.T() <= 0.5:
			//	alphaScale := scale.ClampedToUnitInterval(0, 0.5)
			//	events.Mod(ilysa.WithAlpha(events.GetAlpha() * ease.InOutQuart(alphaScale(ctx.T()))))
			//case ctx.T() > 0.8:
			//	alphaScale := scale.ClampedToUnitInterval(0.8, 1)
			//	events.Mod(ilysa.WithAlpha(events.GetAlpha() * ease.InExpo(1-alphaScale(ctx.T()))))
			//}
		})
	})
}

func (v Verse1b) Stinger(startBeat float64) {
	var (
		seq1 = []float64{0.25, 0.50, 0.75, 1.00, 1.25, 1.50}
		seq2 = []float64{2.00, 2.25, 2.50, 2.75, 3.00, 3.50}

		backLasers = ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v)
		//seq1Light  = backLasers
		seq1Light = ilysa.TransformLight(backLasers,
			ilysa.ToLightTransformer(ilysa.Fan(2)),
			ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		).(ilysa.SequenceLight)
		seq2Light = seq1Light

		seq1Colors = allColors
		seq2Colors = allColors
	)

	ctx := v.WithBeatOffset(startBeat)

	ctx.EventsForSequence(0, seq1, func(ctx ilysa.SequenceContext) {
		if ctx.Ordinal() == 0 {
			ctx.NewPreciseZoomEvent(ilysa.WithStep(-0.9))
		}

		ctx.WithLight(seq1Light, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(ilysa.WithColor(seq1Colors.Rand()))

			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			if !ctx.Last() {
				oe.ShiftBeat(ctx.NextBOffset())
			} else {
				oe.ShiftBeat(0.25)
			}
		})

	})

	ctx.EventsForSequence(0, seq2, func(ctx ilysa.SequenceContext) {
		if ctx.Ordinal() == 0 {
			ctx.NewPreciseZoomEvent(ilysa.WithStep(0.9))
		}

		ctx.WithLight(seq2Light, func(ctx ilysa.SequenceLightContext) {
			ctx.NewRGBLightingEvent(ilysa.WithColor(seq2Colors.Rand()))

			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			if !ctx.Last() {
				oe.ShiftBeat(ctx.NextBOffset())
			} else {
				oe.ShiftBeat(0.25)
			}
		})
	})
}

func (v Verse1b) PianoRoll(startBeat float64) {
	ctx := v.WithBeatOffset(startBeat)

	bl := ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v)

	blForward := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(ilysa.SequenceLight)

	blReverse := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	).(ilysa.SequenceLight)

	light := ilysa.NewSequenceLight(
		blForward.Index(0), blReverse.Index(1), blForward.Index(0),
	)

	ctx.EventForBeat(0, func(ctx ilysa.TimeContext) {
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(90),
			ilysa.WithStep(8),
			ilysa.WithSpeed(8),
			ilysa.WithProp(6),
			ilysa.WithDirection(chroma.Clockwise),
		)

	})

	ctx.EventsForSequence(0, []float64{0, 0.25, 0.50}, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx
		ctx.EventsForRange(ctx.B(), ctx.B()+0.25, 8, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(light.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := ctx.NewRGBLightingEvent(ilysa.WithColor(magnetColors.Next()))
				fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.OutCubic)
			})
		})
	})

	ctx.EventForBeat(0.75, func(ctx ilysa.TimeContext) {
		ctx.WithLight(bl, func(ctx ilysa.TimeLightContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(magnetWhite),
			)
		})
	})
}

func (v Verse1b) Moeagaru(startBeat float64) {
	ctx := v.WithBeatOffset(startBeat)

	bl := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(ilysa.SequenceLight)

	blReverse := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
		ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	).(ilysa.SequenceLight)

	//lightReverse:= ilysa.TransformLight(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//	ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	//)

	//grad := gradient.New(
	//	magnetColors.Index(0),
	//	magnetColors.Index(1),
	//)

	ctx.EventsForBeats(0, 0.5, 4, func(ctx ilysa.TimeContext) {
		var direction chroma.SpinDirection
		if ctx.Ordinal()%2 == 0 {
			direction = chroma.CounterClockwise
		} else {
			direction = chroma.Clockwise
		}
		ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(15+float64(ctx.Ordinal())*5),
			ilysa.WithStep(12.5),
			ilysa.WithSpeed(8),
			ilysa.WithProp(4),
			ilysa.WithDirection(direction),
		)
	})

	grad := gradient.Table{
		{Col: magnetColors.Index(0), Pos: 0.0},
		{Col: magnetColors.Index(1), Pos: 0.6},
		{Col: magnetColors.Index(1), Pos: 1.0},
	}

	var (
		duration = 0.5
		frames   = 10
		step     = 0.3
	)

	ctx.EventsForBeats(0, 0.5, 4, func(ctx ilysa.TimeContext) {
		seqCtx := ctx

		var light ilysa.SequenceLight

		if ctx.Ordinal()%2 == 0 {
			light = bl
		} else {
			light = blReverse
		}
		ctx.EventsForRange(ctx.B(), ctx.B()+duration, frames, ease.Linear, func(ctx ilysa.TimeContext) {
			ctx.WithLight(light.Index(seqCtx.Ordinal()+1), func(ctx ilysa.TimeLightContext) {
				e := fx.Gradient(ctx, magnetGradient)
				e.SetAlpha(1 + float64(seqCtx.Ordinal())*4)
				fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.OutSine)

			})
			ctx.WithLight(light.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := ctx.NewRGBLightingEvent(
					ilysa.WithColor(grad.Ierp(ctx.T())),
				)
				fx.Ripple(ctx, e, step)
				fx.AlphaBlend(ctx, e, 0.0, 0.4, 0, 1, ease.InCirc)
				fx.AlphaBlend(ctx, e, 0.4, 1, 1, 0, ease.OutSine)
			})
		})
	})

	//fx.RainbowProp(ctx, light, grad, 0, 0.2, 0.3, 10)
	//fx.RainbowProp(ctx, light, grad, 0.2, 0.2, 0.8, 10)
	//fx.RainbowProp(ctx, light, grad, 1, 0.2, 0.3, 10)
	//fx.RainbowProp(ctx, light, grad, 1.5, 0.2, 0.3, 10)
	//fx.RainbowProp(ctx, light, gradient.Rainbow, 0.5, 0.25, 2.4, 8)
	//fx.RainbowProp(ctx, light, gradient.Rainbow, 1.0, 0.25, 2.4, 8)
	//fx.RainbowProp(ctx, light, gradient.Rainbow, 1.5, 0.25, 2.4, 8)

}
