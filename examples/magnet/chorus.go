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

type Chorus struct {
	context.Context
}

func NewChorus(project *ilysa.Project, startBeat float64) Chorus {
	return Chorus{project.Offset(startBeat)}
}

func (c Chorus) Play1() {
	// 114
	c.RhythmBridge(0)

	c.Rhythm(2)
	c.Rhythm(6)
	c.Rhythm(10)

	c.Refrain1(1.25, shirayukiWhiteGradient)

	//c.Refrain1a(1.25)
	c.Refrain2(4.75)
	c.Refrain3a(8.75)
	c.Climax(14)
	c.Refrain4(18)
	c.Refrain5(20.75)
	c.Refrain6(24.75)
	c.Refrain7a(30)
	//
	c.Rhythm2(18)
	c.Rhythm2(22)
	c.Rhythm2(26)
	c.Rhythm2(30)
}

func (c Chorus) Play3() {
	// 326
	c.Rhythm(1)
	c.Rhythm(5)
	c.Rhythm(9)

	c.Refrain1(0.25, shirayukiWhiteGradient)

	c.Refrain2(3.75)
	c.Refrain3a(7.75)
	c.Climax(13)
	c.Refrain4(17)
	c.Refrain5(19.75)
	c.Refrain6(23.75)
	c.Refrain7b(29)
	c.Refrain2(35.75)
	c.Refrain3b(39.75)
	c.ClimaxB(46)
	c.Refrain4(49)
	c.Refrain5(51.75)
	c.Refrain6(55.75)
	c.Refrain7a(61)

	c.Rhythm2(17)
	c.Rhythm2(21)
	c.Rhythm2(25)
	c.Rhythm2(29)
	c.Rhythm2(33)
	c.Rhythm2(37)
	c.Rhythm2(41)
	c.Rhythm2(48)
	c.Rhythm2(52)
	c.Rhythm2(56)
	c.Rhythm2(60)
}

func (c Chorus) Play4() {
	// 359
	//c.Rhythm(1)
	//c.Rhythm(5)
	//c.Rhythm(9)

	//c.Refrain1b(0.5)
	//c.Refrain2(2.75)
	//c.Refrain3b(6.75)
	//c.Climax(12)
	//c.Refrain4(13.5)
	//c.Refrain5(19.75)
	//c.Refrain6(23.50)
	//c.Refrain7a(29)
	//
	//c.Rhythm2(17)
	//c.Rhythm2(21)
	//c.Rhythm2(25)
	//c.Rhythm2(29)
}

func (c Chorus) RhythmBridge(startBeat float64) {
	ctx := c.BOffset(startBeat)

	var (
		grad = magnetRainbow
		l    = light.Combine(
			transform.Light(light.NewBasic(ctx, evt.LeftRotatingLasers),
				transform.DivideSingle()),
			transform.Light(light.NewBasic(ctx, evt.RightRotatingLasers),
				transform.DivideSingle()),
		)
	)

	ctx.Sequence(timer.Beat(0.5), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(12),
			evt.WithRotationSpeed(8),
			evt.WithProp(10),
			evt.WithRotationDirection(chroma.Clockwise),
		)
	})

	seq := timer.Seq([]float64{0, 0.5}, 0.99)
	ctx.Sequence(seq, func(ctx context.Context) {
		fx.ZeroSpeedRandomizedLasers(ctx, evt.LeftLaser)
		fx.ZeroSpeedRandomizedLasers(ctx, evt.RightLaser)

		if ctx.SeqLast() {
			e := fx.OffAll(ctx)
			e.Apply(evt.WithBOffset(ctx.SeqNextBOffset()))
		}

		ctx.Range(timer.Rng(0, 0.5, 12, ease.Linear), func(ctx context.Context) {
			if ctx.Ordinal() == 1 {
				grad = grad.Reverse()
			}
			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				fx.AlphaFadeEx(ctx, e, 0, 1, 6, 0, ease.OutCubic)
			})
		})
	})
}

func (c Chorus) Rhythm(startBeat float64) {
	ctx := c.BOffset(startBeat)

	l := light.NewSequence(
		light.NewBasic(ctx, evt.LeftRotatingLasers),
		light.NewBasic(ctx, evt.RightRotatingLasers),
	)

	seq := timer.Interval(0, 1, 4)
	ctx.Sequence(seq, func(ctx context.Context) {
		fx.ZeroSpeedRandomizedLasers(ctx, evt.LeftLaser)
		fx.ZeroSpeedRandomizedLasers(ctx, evt.RightLaser)

		ctx.Light(l, func(ctx context.LightContext) {
			ctx.NewRGBLighting(
				evt.WithColor(magnetColors.Next()),
				evt.WithLightValue(evt.LightBlueFade),
				evt.WithAlpha(0.3),
			)
		})
	})

	var (
		//rippleDuration = 4.0
		//rippleStart    = 0.0
		//rippleEnd      = rippleStart + rippleDuration
		//rippleLights   = c.NewBasicLight(beatsaber.EventTypeRingLights).Transform(rework.DivideSingle)
		//rippleTime     = 0.8
		rng  = timer.Rng(0, 4, 60, ease.Linear)
		grad = gradient.Table{
			{shirayukiPurple, 0.0},
			{sukoyaWhite, 0.3},
			{sukoyaWhite, 0.7},
			{shirayukiPurple, 1.0},
		}
	)

	RingRipple(ctx, rng, grad,
		WithRippleTime(0.8),
		WithSweepSpeed(1.2),
		WithFadeIn(fx.NewAlphaFader(0, 0.2, 0, 1.2, ease.InCirc)),
		WithFadeOut(fx.NewAlphaFader(0.2, 1.0, 1.2, 0, ease.OutCirc)),
	)

	//ctx.Range(rippleStart, rippleEnd, 60, ease.Linear, func(ctx context.Context) {
	//	ctx.Light(rippleLights, func(ctx context.LightContext) {
	//		events := fx.ColorSweep(ctx, 1.2, grad)
	//
	//		fx.Ripple(ctx, events, rippleTime,
	//			fx.WithAlphaBlend(0, 0.2, 0, 0.3, ease.InCubic),
	//			fx.WithAlphaBlend(0.2, 1.0, 0.3, 0, ease.OutCubic),
	//		)
	//	})
	//})
}

func (c Chorus) Refrain1(startBeat float64, windupGrad gradient.Table) {
	ctx := c.BOffset(startBeat)

	var (
		windupEndBeat = 0.75
		fadeSeq       = timer.Seq([]float64{1.25, 1.50, 2.00, 2.50}, 0)
	)

	c.Sweep(ctx, 0, windupEndBeat, windupGrad, false)
	c.SweepSpin(ctx, windupEndBeat, false)
	c.FadeToGold(ctx, fadeSeq)
}

//func (c Chorus) Refrain1a(startBeat float64) {
//	ctx := c.BOffset(startBeat)
//	c.Sweep(ctx, 0, 0.75, shirayukiWhiteGradient, false)
//	c.SweepSpin(ctx, 0.75, false)
//	c.FadeToGold(ctx, timer.Seq([]float64{1.25, 1.50, 2.00, 2.50}, 0))
//}
//
//func (c Chorus) Refrain1b(startBeat float64) {
//	ctx := c.BOffset(startBeat)
//	c.FadeToGold(ctx, timer.Seq([]float64{0.0, 0.25, 0.75, 1.25}, 0))
//}

func (c Chorus) Refrain2(startBeat float64) {
	ctx := c.BOffset(startBeat)
	c.Sweep(ctx, 0, 1.00, sukoyaGradient, false)
	c.SweepSpin(ctx, 1.00, false)
	c.FadeToGold(ctx, timer.Seq([]float64{1.50, 2.00, 2.50, 3.00}, 0))
}

func (c Chorus) Refrain3a(startBeat float64) {
	ctx := c.BOffset(startBeat)

	c.Sweep(ctx, 0, 0.50, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 0.50, false)
	c.FadeToSukoya(ctx, timer.Seq([]float64{1.25, 1.75, 2.25, 2.75}, 0))
	c.HalfCollapse(ctx, timer.Beat(3), false)

	collapseSeq := timer.Seq([]float64{3.75, 4.25, 4.75}, 0)

	c.HalfCollapse(ctx, collapseSeq, false)
	c.FadeToGold(ctx, collapseSeq)
}

func (c Chorus) Refrain3b(startBeat float64) {
	ctx := c.BOffset(startBeat)

	c.Sweep(ctx, 0, 0.50, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 0.50, false)
	c.FadeToSukoya(ctx, timer.Seq([]float64{1.00, 1.50, 2.00, 2.75}, 0))
	c.HalfCollapse(ctx, timer.Beat(3), false)

	collapseSeq := timer.Seq([]float64{3.75, 4.25, 4.75, 5.25}, 0)

	c.HalfCollapse(ctx, collapseSeq, false)
	c.FadeToGold(ctx, collapseSeq)
}

func (c Chorus) Climax(startBeat float64) {
	ctx := c.BOffset(startBeat)

	bl := light.NewBasic(ctx, evt.BackLasers)

	seq := timer.Seq([]float64{0, 2, 2.75, 3.5}, 4)
	ctx.Sequence(seq, func(ctx context.Context) {
		dir := chroma.Clockwise
		if ctx.Ordinal() == 1 {
			dir = dir.Reverse()
		}

		ctx.NewPreciseRotation(
			evt.WithRotation(45*float64(ctx.Ordinal()+1)),
			evt.WithRotationStep(float64(90/(ctx.Ordinal()+1))),
			evt.WithRotationSpeed(16),
			evt.WithProp(1.3),
			evt.WithRotationDirection(dir),
		)

		col := crossickColors.Next()
		ctx.Range(timer.Rng(0, ctx.SeqNextBOffset(), 18, ease.Linear), func(ctx context.Context) {
			ctx.Light(bl, func(ctx context.LightContext) {
				e := ctx.NewRGBLighting(evt.WithColor(col))
				fx.AlphaFadeEx(ctx, e, 0, 1, 2+float64(ctx.SeqOrdinal()), 0.3, ease.OutCubic)
			})
		})
	})
}

func (c Chorus) ClimaxB(startBeat float64) {
	ctx := c.BOffset(startBeat)

	bl := light.NewBasic(ctx, evt.BackLasers)

	seq := timer.Seq([]float64{0, 0.5, 1, 1.75, 2.5}, 3)
	ctx.Sequence(seq, func(ctx context.Context) {
		dir := chroma.Clockwise
		if ctx.Ordinal() == 1 {
			dir = dir.Reverse()
		}

		ctx.NewPreciseRotation(
			evt.WithRotation(45*float64(ctx.Ordinal()+1)),
			evt.WithRotationStep(float64(90/(ctx.Ordinal()+1))),
			evt.WithRotationSpeed(16),
			evt.WithProp(1.3),
			evt.WithRotationDirection(dir),
		)

		col := crossickColors.Next()
		ctx.Range(timer.Rng(0, ctx.SeqNextBOffset(), 18, ease.Linear), func(ctx context.Context) {
			ctx.Light(bl, func(ctx context.LightContext) {
				e := ctx.NewRGBLighting(evt.WithColor(col))
				fx.AlphaFadeEx(ctx, e, 0, 1, 2+float64(ctx.SeqOrdinal()), 0.3, ease.OutCubic)
			})
		})
	})
}

// キスをして
func (c Chorus) Refrain4(startBeat float64) {
	ctx := c.BOffset(startBeat)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(180),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)

		ctx.NewPreciseZoom(evt.WithZoomStep(-0.66))
	})

	rotLasers := light.Combine(
		light.NewBasic(ctx, evt.LeftRotatingLasers),
		light.NewBasic(ctx, evt.RightRotatingLasers),
	)

	ctx.Sequence(timer.Seq([]float64{0, 2}, 0), func(ctx context.Context) {
		opt := evt.NewOpts(
			evt.WithLaserSpeed(30),
			evt.WithPreciseLaserSpeed(0),
		)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser), opt)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser), opt)

		ctx.Light(rotLasers, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithColor(allColors.Next()), evt.WithLightValue(evt.LightRedFade))
		})
	})

	c.Rush(ctx, -0.5, 0.8, 0.9, 10, gradient.FromSet(sukoyaColors))
	// キスを
	c.Sweep(ctx, 0, 0.75, magnetGradient, false)

	// して
	c.Unsweep(ctx, 0, []float64{1.25, 1.75})
	//c.FadeToGold(ctx, 0, []float64{1.25, 1.75})

	ctx.Sequence(timer.Seq([]float64{1.25, 1.75}, 0), func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(0.33 * (1 - float64(ctx.SeqT()))))
	})
}

func (c Chorus) Refrain5(startBeat float64) {
	ctx := c.BOffset(startBeat)

	c.Sweep(ctx, 0, 0.50, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 0.5, false)
	c.FadeToGold(ctx, timer.Seq([]float64{1.00, 1.50, 2.00, 2.50, 3.00}, 0))
	//c.FadeToGold(ctx, -0.5, []float64{1.25, 1.75, 2.25, 2.75, 3.25})
}

func (c Chorus) Refrain6(startBeat float64) {
	ctx := c.BOffset(startBeat)

	c.Sweep(ctx, 0, 0.50, sukoyaGradient, true)
	c.SweepSpin(ctx, 0.50, true)
	c.HalfCollapse(ctx, timer.Beat(2.50), true)
	c.FadeToSukoya(ctx, timer.Seq([]float64{1.00, 1.75, 2.25, 2.75}, 0))
	c.FadeToGold(ctx, timer.Seq([]float64{3.75, 4.25, 4.75, 5.25}, 0))
}

func (c Chorus) Refrain7a(startBeat float64) {
	ctx := c.BOffset(startBeat)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(180),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)
		ctx.NewPreciseZoom(evt.WithZoomStep(-0.66))
	})

	RingRipple(ctx, timer.Rng(-1, 0, 40, ease.OutExpo), shirayukiGradient,
		WithSweepSpeed(4),
		WithRippleTime(0.8),
		WithReverse(true),
		WithFadeIn(fx.NewAlphaFader(0.0, 0.6, 2, 10, ease.InSin)),
		WithFadeOut(fx.NewAlphaFader(0.6, 1.0, 10, 0, ease.InCirc)),
	)

	ctx.Sequence(timer.Seq([]float64{1.0, 2.0, 2.75, 3.5, 4.0, 5.0}, 6.0), func(ctx context.Context) {
		grad := sukoyaGradient
		offset := -0.50
		c.Rush(ctx, 0+offset, ctx.SeqNextBOffset()+offset-0.10, 0.70, 2*float64(ctx.Ordinal()), grad)
	})

	ctx.Sequence(timer.Seq([]float64{2, 2.75, 3.5}, 0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(30*float64(ctx.Ordinal())),
			evt.WithRotationStep(7*float64(ctx.Ordinal()*5)),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)
	})

	ctx.Sequence(timer.Beat(4), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(12.5),
			evt.WithRotationSpeed(4),
			evt.WithProp(1),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)

		ctx.NewRGBLighting(evt.WithLight(evt.BackLasers), evt.WithLightValue(evt.LightRedFade),
			evt.WithColor(sukoyaWhite),
		)
	})

	bl := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.DivideSingle(),
	)

	ctx.Sequence(timer.Beat(5), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(360),
			evt.WithRotationStep(0),
			evt.WithRotationSpeed(1.6),
			evt.WithProp(0.9),
			evt.WithRotationDirection(chroma.Clockwise),
		)
		ctx.NewPreciseZoom(evt.WithZoomStep(1.5))

		rng := timer.Rng(0, 4, 60, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(bl, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 1.2, magnetRainbow)
				fx.AlphaFadeEx(ctx, e, 0, 1, 6, 0.0, ease.InOutQuad)
			})
		})
	})
}

func (c Chorus) Refrain7b(startBeat float64) {
	ctx := c.BOffset(startBeat)

	ctx.Sequence(timer.Seq([]float64{0, 1}, 0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(180),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)
		ctx.NewPreciseZoom(evt.WithZoomStep(-0.66))

		RingRipple(ctx, timer.Rng(-1, 0, 30, ease.OutExpo), shirayukiGradient,
			WithSweepSpeed(4),
			WithRippleTime(0.8),
			WithReverse(true),
			WithFadeIn(fx.NewAlphaFader(0.0, 0.6, 2, 10, ease.InSin)),
			WithFadeOut(fx.NewAlphaFader(0.6, 1.0, 10, 0, ease.InCirc)),
		)
	})

	ctx.Sequence(timer.Seq([]float64{1.5, 2.0, 2.75, 3.5}, 4.0), func(ctx context.Context) {
		grad := sukoyaGradient
		offset := -0.35
		c.Rush(ctx, 0+offset, ctx.SeqNextBOffset()+offset-0.10, 0.70, 2*float64(ctx.Ordinal()), grad)
	})

	ctx.Sequence(timer.Seq([]float64{2, 2.75, 3.5}, 0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(30*float64(ctx.Ordinal())),
			evt.WithRotationStep(7*float64(ctx.Ordinal()*5)),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)
	})

	ctx.Sequence(timer.Beat(4), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(12.5),
			evt.WithRotationSpeed(4),
			evt.WithProp(1),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)

		//c.Rush(ctx, 0+offset, ctx.SeqNextBOffset()+offset-0.10, 0.70, 2*float64(ctx.Ordinal()), grad)

		ctx.NewRGBLighting(evt.WithLight(evt.BackLasers), evt.WithLightValue(evt.LightRedFlash),
			evt.WithColor(sukoyaWhite),
		)
		ctx.NewPreciseZoom(evt.WithZoomStep(0))
	})
	c.FadeToGold(ctx, timer.Seq([]float64{4.50, 4.75, 5.25, 5.75}, 6.00))
	//ctx.Sequence(, func(ctx context.Context) {
	//
	//
	//})

	//bl := transform.Light(light.NewBasic(ctx, evt.BackLasers),
	//	transform.DivideSingle(),
	//)

	//ctx.Sequence(timer.Beat(5), func(ctx context.Context) {
	//	ctx.NewPreciseRotation(
	//		evt.WithRotation(90),
	//		evt.WithRotationStep(0),
	//		evt.WithRotationSpeed(1.6),
	//		evt.WithProp(0.9),
	//		evt.WithRotationDirection(chroma.Clockwise),
	//	)
	//	ctx.NewPreciseZoom(evt.WithZoomStep(0))
	//
	//	//rng := timer.Rng(0, 0.5, 20, ease.Linear)
	//	//ctx.Range(rng, func(ctx context.Context) {
	//	//	ctx.Light(bl, func(ctx context.LightContext) {
	//	//		e := fx.ColorSweep(ctx, 1.2, magnetRainbow)
	//	//		fx.AlphaFadeEx(ctx, e, 0, 1, 6, 0.0, ease.InOutQuad)
	//	//	})
	//	//})
	//})
}
func (c Chorus) Rhythm2(startBeat float64) {
	ctx := c.BOffset(startBeat)

	l := light.NewSequence(
		light.NewBasic(ctx, evt.LeftRotatingLasers),
		light.NewBasic(ctx, evt.RightRotatingLasers),
	)

	seq := timer.Interval(0, 1, 4)
	ctx.Sequence(seq, func(ctx context.Context) {
		fx.ZeroSpeedRandomizedLasers(ctx, evt.LeftLaser)
		fx.ZeroSpeedRandomizedLasers(ctx, evt.RightLaser)

		ctx.Light(l, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithLightValue(evt.LightBlueFade),
				evt.WithColor(magnetColors.Next()),
				evt.WithAlpha(0.3),
			)
		})
	})
}

func (c Chorus) Sweep(ctx context.Context, startBeat, endBeat float64, grad gradient.Table, reverse bool) {
	backLasers := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
	)

	if reverse {
		backLasers = transform.Light(backLasers, transform.Reverse())
	}

	lightSweep := transform.Light(backLasers, transform.DivideSingle())
	lightSweepSeq := transform.Light(backLasers, transform.DivideSingle().Sequence())

	ctx.Range(timer.Rng(startBeat, endBeat, 30, ease.OutCubic), func(ctx context.Context) {
		ctx.Light(lightSweepSeq, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithColor(grad.Lerp(ctx.T())))
		})
	})

	ctx.Sequence(timer.Beat(endBeat+0.01), func(ctx context.Context) {
		ctx.Light(lightSweep, func(ctx context.LightContext) {
			e := fx.Gradient(ctx, grad.Reverse())
			e.Apply(evt.WithAlpha(2))
		})
	})
}

func (c Chorus) SweepSpin(ctx context.Context, startBeat float64, reverse bool) {
	dir := chroma.CounterClockwise.ReverseIf(reverse)

	ctx.Sequence(timer.Beat(startBeat), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(135),
			evt.WithRotationStep(12),
			evt.WithRotationSpeed(8), // 2.1
			evt.WithProp(0.8),        // 0.9
			evt.WithRotationDirection(dir),
		)
	})
}

func (c Chorus) HalfCollapse(ctx context.Context, seq timer.Sequencer, reverse bool) {
	dir := chroma.Clockwise.ReverseIf(reverse)

	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(9),
			evt.WithRotationSpeed(8),
			evt.WithProp(20),
			evt.WithRotationDirection(dir),
		)
	})
}

func (c Chorus) Unsweep(ctx context.Context, startBeat float64, sequence []float64) {
	ctx = ctx.BOffset(startBeat)

	seq := timer.Seq(sequence, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(8),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)
	})
}

func (c Chorus) Rush(ctx context.Context, startBeat, endBeat, rippleDuration, peakAlpha float64, grad gradient.Table) {
	l := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.Reverse(),
		transform.DivideSingle(),
	)

	ctx.Range(timer.Rng(startBeat, endBeat, 30, ease.InExpo), func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, 2.4, grad)
			fx.RippleT(ctx, e, rippleDuration)
			//fx.AlphaFadeEx(ctx, e, 0, 0.3, 2, peakAlpha, ease.Linear)
			fx.AlphaFadeEx(ctx, e, 0.0, 1.0, peakAlpha, 0, ease.InQuart)
		})
	})
}

func (c Chorus) RushNoFade(ctx context.Context, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	l := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.Reverse(),
		transform.DivideSingle(),
	)

	rng := timer.Rng(startBeat, endBeat, 45, ease.InExpo)

	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.BiasedColorSweep(ctx, 4, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaFadeEx(ctx, e, 0, 0.6, 2, peakAlpha, ease.OutCubic)
		})
	})
}

func (c Chorus) FadeToSukoya(ctx context.Context, seq timer.Sequencer) {
	lightSweepDiv := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(seq.Len()).Sequence(),
	).(light.Sequence)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.Range(timer.Rng(0, 0.5, 16, ease.OutCubic), func(ctx context.Context) {
			ctx.Light(lightSweepDiv, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 3.6, gradient.FromSet(sukoyaColors))
				fx.AlphaFadeEx(ctx, e, 0, 1, 0.3, 3, ease.OutElastic)
			})
		})
	})
}

func (c Chorus) FadeToGold(ctx context.Context, seq timer.Sequencer) {
	lightSweepDiv := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Rotate(3),
		transform.Divide(seq.Len()).Sequence(),
		transform.DivideSingle(),
	).(light.Sequence)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.Sequence(seq, func(ctx context.Context) {
		grad := magnetRainbowPale.RotateRand()
		ctx.Range(timer.Rng(0, 0.5, 16, ease.OutCubic), func(ctx context.Context) {
			ctx.Light(lightSweepDiv, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 1, grad)
				fx.AlphaFadeEx(ctx, e, 0, 1, 10, 0, ease.OutCirc)
			})
		})
	})
}
