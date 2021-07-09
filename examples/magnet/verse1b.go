package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

type Verse1b struct {
	context.Context
}

func NewVerse1b(p *ilysa.Project, offset float64) Verse1b {
	return Verse1b{
		Context: p.Offset(offset),
	}
}

func (v Verse1b) Play() {
	v.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewLaser(evt.WithDirectionalLaser(evt.LeftLaser), evt.WithIntValue(5))
		ctx.NewLaser(evt.WithDirectionalLaser(evt.RightLaser), evt.WithIntValue(5))
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
	////v.Moeagaru(22.5)
	v.Moeagaru(26.5)
}

func (v Verse1b) IntroBridge(startBeat float64) {
	ctx := v.Offset(startBeat)

	seq := timer.NewSequencer([]float64{0, 1, 3}, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
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
			evt.WithRotationSpeed(speed[ctx.Ordinal()]),
			evt.WithProp(prop[ctx.Ordinal()]),
			evt.WithRotationDirection(direction[ctx.Ordinal()]),
		)
	})

	ctx.Sequence(timer.NewSequencer([]float64{0, 1, 2}, 4), func(ctx context.Context) {
		ctx.NewRGBLighting(evt.WithLight(evt.BackLasers), evt.WithLightValue(evt.LightBlueFlash),
			evt.WithColor(crossickColors.Next()),
			evt.WithAlpha(3),
		)

		if ctx.Last() {
			ctx.NewRGBLighting(evt.WithLight(evt.BackLasers), evt.WithLightValue(evt.LightOff),
				evt.WithBeatOffset(ctx.SeqNextBOffset()),
			)
		}
	})
}

func (v Verse1b) Rhythm(startBeat float64, spin bool) {
	ctx := v.Offset(startBeat)
	var (
		seq               = timer.NewSequencer([]float64{1, 3}, 0)
		leftRightSequence = light.NewSequence(
			light.NewBasic(ctx, evt.LeftRotatingLasers),
			light.NewBasic(ctx, evt.RightRotatingLasers),
		)
		color = crossickColors
	)

	ctx.Sequence(seq, func(ctx context.Context) {
		if spin {
			ctx.NewPreciseRotation(
				evt.WithRotation(90),
				evt.WithRotationStep(11.25),
				evt.WithRotationSpeed(7),
				evt.WithProp(3),
				evt.WithRotationDirection(chroma.CounterClockwise),
			)
		}

		ctx.Light(leftRightSequence, func(ctx context.LightContext) {
			ctx.NewRGBLighting(
				evt.WithLightValue(evt.LightRedFade),
				evt.WithColor(color.Next()),
			)
		})
	})

	rng := timer.NewRanger(1, 3, 30, ease.Linear)
	RingRipple(ctx, rng, gradient.FromSet(crossickColors),
		WithRippleTime(0.8),
		WithSweepSpeed(2),
		WithFadeIn(fx.NewAlphaFader(0, 0.2, 0, 2, ease.InCubic)),
		WithFadeOut(fx.NewAlphaFader(0.7, 1, 2, 0, ease.OutCubic)),
	)

	//var (
	//	rippleDuration = 1.0
	//	rippleStart    = startBeat + 2
	//	rippleEnd      = rippleStart + rippleDuration
	//	rippleLights   = v.NewBasicLight(beatsaber.EventTypeRingLights).Transform(rework.DivideSingle)
	//	rippleStep     = 0.8
	//)
	//
	//v.Range(rippleStart, rippleEnd, 30, ease.Linear, func(ctx context.Context) {
	//	ctx.Light(rippleLights, func(ctx context.LightContext) {
	//		e := fx.ColorSweep(ctx, 2, gradient.FromSet(crossickColors))
	//
	//		fx.Ripple(ctx, e, rippleStep)
	//		fx.AlphaBlend(ctx, e, 0, 0.2, 0, 2, ease.InCubic)
	//		fx.AlphaBlend(ctx, e, 0.8, 1, 2, 0, ease.OutCubic)
	//	})
	//})
}

func (v Verse1b) Stinger(startBeat float64) {
	ctx := v.Offset(startBeat)
	v.PianoHits(ctx, timer.NewSequencer([]float64{0.25, 0.50, 0.75, 1.00, 1.25, 1.50}, 0), 6, 0.6)
	v.PianoHits(ctx, timer.NewSequencer([]float64{2.00, 2.25, 2.50, 2.75, 3.00, 3.50}, 0), 6, 0.6)

	//	backLasers = ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v)
	//	//seq1Light  = backLasers
	//	seq1Light = ilysa.LightTransform(backLasers,
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
	//ctx := v.Offset(startBeat)
	//
	//ctx.seq(0, seq1, func(ctx ilysa.SequenceContext) {
	//	if ctx.Ordinal() == 0 {
	//		ctx.NewPreciseZoom(ilysa.WithRotationStep(-0.9))
	//	}
	//
	//	ctx.Light(seq1Light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLighting(evt.WithColor(seq1Colors.Rand()))
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
	//ctx.seq(0, seq2, func(ctx ilysa.SequenceContext) {
	//	if ctx.Ordinal() == 0 {
	//		ctx.NewPreciseZoom(ilysa.WithRotationStep(0.9))
	//	}
	//
	//	ctx.Light(seq2Light, func(ctx ilysa.SequenceLightContext) {
	//		ctx.NewRGBLighting(evt.WithColor(seq2Colors.Rand()))
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

	//ctx := v.Offset(startBeat)
	//
	//bl := ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v)
	//
	//blForward := ilysa.LightTransform(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//).(ilysa.SequenceLight)
	//
	//blReverse := ilysa.LightTransform(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//	ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	//).(ilysa.SequenceLight)
	//
	//light := ilysa.NewSequenceLight(
	//	blForward.Idx(0), blReverse.Idx(1), blForward.Idx(0),
	//)
	//
	//ctx.Sequence(timer.Beat(0, func(ctx ilysa.RangeContext) {
	//	ctx.NewPreciseRotation(
	//		ilysa.WithRotation(90),
	//		ilysa.WithRotationStep(8),
	//		ilysa.WithRotationSpeed(8),
	//		ilysa.WithProp(6),
	//		ilysa.WithRotationDirection(chroma.Clockwise),
	//	)
	//
	//})
	//
	//ctx.seq(0, []float64{0, 0.25, 0.50}, func(ctx ilysa.SequenceContext) {
	//	seqCtx := ctx
	//	ctx.rangeTimer(ctx.B(), ctx.B()+0.25, 8, ease.Linear, func(ctx ilysa.RangeContext) {
	//		ctx.Light(light.Idx(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
	//			e := ctx.NewRGBLighting(evt.WithColor(magnetColors.Next()))
	//			fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.OutCubic)
	//		})
	//	})
	//})
	//
	//ctx.Sequence(timer.Beat(0.75, func(ctx ilysa.RangeContext) {
	//	ctx.Light(bl, func(ctx ilysa.TimeLightContext) {
	//		ctx.NewRGBLighting(
	//			ilysa.WithValue(beatsaber.EventValueLightRedFade),
	//			evt.WithColor(magnetWhite),
	//		)
	//	})
	//})

	ctx := v.Offset(startBeat)
	v.PianoHits(ctx, timer.NewSequencer([]float64{0, 0.25, 0.50, 0.75}, 0), 5, 1)
}

func (v Verse1b) Naosara(startBeat float64) {
	ctx := v.Offset(startBeat)

	seq := timer.NewSequencer([]float64{0.4, 0.9, 1.4}, 1.9)

	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(6),
			evt.WithRotationSpeed(3),
			evt.WithProp(20),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)

		if ctx.SeqLast() {
			ctx.NewPreciseRotation(
				evt.WithBeatOffset(ctx.SeqNextBOffset()),
				evt.WithRotation(90),
				evt.WithRotationStep(11.5),
				evt.WithRotationSpeed(0.9),
				evt.WithProp(0.8),
				evt.WithRotationDirection(chroma.Clockwise),
			)
		}
	})

	v.Burn(ctx, timer.NewRanger(0.0, 0.5, 15, ease.InExpo), magnetGradient, 0)
	v.Burn(ctx, timer.NewRanger(0.5, 1.0, 15, ease.InExpo), magnetGradient, 1)
	v.Burn(ctx, timer.NewRanger(1.0, 1.5, 15, ease.InExpo), sukoyaGradient, 0)
	v.Burn(ctx, timer.NewRanger(1.5, 2.0, 15, ease.InExpo), shirayukiGradient, 1)
	//v.Burn(ctx, 0.5, 1.00, magnetGradient, 1)
	//v.Burn(ctx, 1.0, 1.50, sukoyaGradient, 0)
	//v.Burn(ctx, 1.5, 2.00, shirayukiGradient, 1)
}

func (v Verse1b) Moeagaru(startBeat float64) {
	ctx := v.Offset(startBeat)

	v.PianoHits(ctx, timer.NewSequencer([]float64{0, 0.5, 1.0, 1.5}, 0), 4, 0.5)

	lightSweepDiv := transform.Light(
		light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		//ilysa.ToSequenceLightTransformer(ilysa.Divide(divisor)),
	)

	//lightSweepDiv = lightSweepDiv.Shuffle()

	rng := timer.NewRanger(3, 3.25, 12, ease.OutCubic)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(lightSweepDiv, func(ctx context.LightContext) {
			fx.ColorSweep(ctx, 12, magnetRainbowPale)
		})
	})

	ctx.Sequence(timer.Beat(3.25), func(ctx context.Context) {
		ctx.NewRGBLighting(evt.WithLight(evt.BackLasers), evt.WithLightValue(evt.LightOff))
	})

	//bl := ilysa.LightTransform(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//).(ilysa.SequenceLight)
	//
	//blReverse := ilysa.LightTransform(
	//	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	//	ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
	//	ilysa.ToLightTransformer(ilysa.DivideSingle),
	//	ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	//).(ilysa.SequenceLight)
	//
	////lightReverse:= ilysa.LightTransform(
	////	ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, v),
	////	ilysa.ToLightTransformer(ilysa.DivideSingle),
	////	ilysa.LightIDSetTransformerToLightTransformer(ilysa.ReverseSet),
	////)
	//
	////grad := gradient.New(
	////	magnetColors.Idx(0),
	////	magnetColors.Idx(1),
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
	//		ilysa.WithRotationSpeed(8),
	//		ilysa.WithProp(4),
	//		ilysa.WithRotationDirection(direction),
	//	)
	//})
	//
	//grad := gradient.Table{
	//	{Col: magnetColors.Idx(0), Pos: 0.0},
	//	{Col: magnetColors.Idx(1), Pos: 0.6},
	//	{Col: magnetColors.Idx(1), Pos: 1.0},
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
	//		ctx.Light(light.Idx(seqCtx.Ordinal()+1), func(ctx ilysa.TimeLightContext) {
	//			e := fx.Gradient(ctx, magnetGradient)
	//			e.SetAlpha(1 + float64(seqCtx.Ordinal())*4)
	//			fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.OutSine)
	//
	//		})
	//		ctx.Light(light.Idx(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
	//			e := ctx.NewRGBLighting(
	//				evt.WithColor(grad.Lerp(ctx.T())),
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

func (v Verse1b) PianoHits(ctx context.Context, seq timer.Sequencer, divisor int, duration float64) {
	l := transform.Light(
		light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(divisor).Sequence(),
	).(light.Sequence)
	//lightSweepDiv := transform.Light(
	//	light.NewBasic(beatsaber.EventTypeBackLasers, v),
	//	rework.ToLightTransformer(rework.Fan(2)),
	//	rework.LightIDSetTransformerToLightTransformer(rework.Flatten),
	//	rework.ToSequenceLightTransformer(rework.Divide(divisor)),
	//).(light2.SequenceLight)

	l = l.Shuffle()

	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.Range(timer.NewRanger(0, duration, 16, ease.OutCubic), func(ctx context.Context) {
			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 3.6, magnetRainbowPale)
				alpha := 3.0
				if ctx.SeqLast() {
					alpha = 8
				}
				fx.AlphaFadeEx(ctx, e, 0, 1, alpha, 0, ease.OutElastic)
			})
		})
	})

	//ctx.Sequence(0, sequence, func(ctx rework.SequenceContext) {
	//	seqCtx := ctx
	//	ctx.EventsForRange(ctx.B(), ctx.B()+duration, 16, ease.OutCubic, func(ctx context.Context) {
	//		ctx.Light(lightSweepDiv.Idx(seqCtx.Ordinal()), func(ctx context.LightContext) {
	//		})
	//	})
	//})
}

func (v Verse1b) Burn(ctx context.Context, rng timer.Ranger, grad gradient.Table, ordinal int) {
	var reverseTransform transform.LightTransformer = transform.Identity()

	if ordinal == 1 {
		reverseTransform = transform.Reverse()
	}

	backLasers := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(2).Sequence(),
		reverseTransform,
		transform.DivideSingle(),
	).(light.Sequence)

	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(backLasers.Idx(ordinal), func(ctx context.LightContext) {
			e := ctx.NewRGBLighting(evt.WithColor(grad.Lerp(ctx.T())))
			fx.Ripple(ctx, e, 0.5)
		})
	})

	//backLasers := transform.Light(light.NewBasic(beatsaber.EventTypeBackLasers, v),
	//	rework.ToLightTransformer(rework.Fan(2)),
	//	rework.LightIDSetTransformerToLightTransformer(rework.Flatten),
	//	rework.ToSequenceLightTransformer(rework.Divide(2)),
	//	reverseTransform,
	//	rework.ToLightTransformer(rework.DivideSingle),
	//).(light2.SequenceLight)

	//ctx.Range(startBeat, endBeat, 15, ease.InExpo, func(ctx context.Context) {
	//	ctx.Light(backLasers.Idx(ordinal), func(ctx context.LightContext) {
	//		e := ctx.NewRGBLightingEvent(evt.WithColor(grad.Lerp(ctx.T())))
	//		fx.Ripple(ctx, e, 0.5)
	//	})
	//})
}
