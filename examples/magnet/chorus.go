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
	ctx := c.Offset(startBeat)
	var (
		grad = magnetRainbow
		l    = light.Combine(
			transform.Light(light.NewBasic(ctx, evt.LeftRotatingLasers),
				transform.DivideSingle()),
			transform.Light(light.NewBasic(ctx, evt.RightRotatingLasers),
				transform.DivideSingle()),
		)
	)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(evt.WithReset(true))
	})

	ctx.Sequence(timer.Beat(0.5), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(15),
			evt.WithRotationSpeed(2.8),
			evt.WithProp(0.9),
			evt.WithRotationDirection(chroma.Clockwise),
		)
	})

	seq := timer.NewSequencer([]float64{0, 0.5}, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
		fx.ZeroSpeedRandomizedLasers(ctx, evt.LeftLaser)
		fx.ZeroSpeedRandomizedLasers(ctx, evt.RightLaser)

		rng := timer.NewRanger(0, 0.5, 12, ease.Linear)

		ctx.Range(rng, func(ctx context.Context) {
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
	ctx := c.Offset(startBeat)

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
		rng  = timer.NewRanger(0, 4, 60, ease.Linear)
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
		WithFadeIn(fx.NewAlphaFader(0, 0.2, 0, 0.3, ease.InCubic)),
		WithFadeOut(fx.NewAlphaFader(0.2, 1.0, 0.3, 0, ease.OutCubic)),
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

func (c Chorus) Refrain1(startBeat float64) {
	ctx := c.Offset(startBeat)
	c.Sweep(ctx, 0.25, 1, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 1, false)
	c.FadeToGold(ctx, 1, []float64{0.5, 0.75, 1.25, 1.75})
}

func (c Chorus) Refrain2(startBeat float64) {
	ctx := c.Offset(startBeat)
	c.Sweep(ctx, 0, 1.25, sukoyaGradient, true)
	c.SweepSpin(ctx, 1.25, true)
	c.FadeToGold(ctx, 1.25, []float64{0.5, 1.0, 1.5, 2.0})
}

func (c Chorus) Refrain3(startBeat float64) {
	ctx := c.Offset(startBeat)

	c.Sweep(ctx, 0, 1.25, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 1.25, false)
	c.FadeToSukoya(ctx, 1.25, []float64{0, 0.5, 1.0, 1.5})
	c.HalfCollapse(ctx, 3, []float64{0}, false)

	collapseSeq := []float64{2.5, 3.0, 3.5, 4.0}

	//c.HalfCollapse(ctx, 1.0, collapseSeq, false)
	c.FadeToGold(ctx, 1.00, collapseSeq)
}

func (c Chorus) Climax(startBeat float64) {
	ctx := c.Offset(startBeat)

	bl := light.NewBasic(ctx, evt.BackLasers)

	seq := timer.NewSequencer([]float64{0, 2, 2.75, 3.5}, 4)
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
		rng := timer.NewRanger(0, ctx.SeqNextBOffset(), 18, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(bl, func(ctx context.LightContext) {
				e := ctx.NewRGBLighting(
					evt.WithColor(col),
				)
				fx.AlphaFadeEx(ctx, e, 0, 1, 2+float64(ctx.SeqOrdinal()), 0.3, ease.OutCubic)
			})
		})
	})
}

// キスをして
func (c Chorus) Refrain4(startBeat float64) {
	ctx := c.Offset(startBeat)

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

	ctx.Sequence(timer.NewSequencer([]float64{0, 2}, 0), func(ctx context.Context) {
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

	c.Rush(ctx, -1, 0, 1.5, 10, gradient.FromSet(sukoyaColors))
	// キスを
	c.Sweep(ctx, 0, 0.75, magnetGradient, false)

	// して
	c.Unsweep(ctx, 0, []float64{1.25, 1.75})
	c.FadeToGold(ctx, 0, []float64{1.25, 1.75})

	ctx.Sequence(timer.NewSequencer([]float64{1.25, 1.75}, 0), func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(0.33))
	})
}

func (c Chorus) Refrain5(startBeat float64) {
	ctx := c.Offset(startBeat)

	c.Sweep(ctx, 0, 0.50, shirayukiWhiteGradient, false)
	c.SweepSpin(ctx, 0.5, false)
	c.FadeToGold(ctx, -0.5, []float64{1.25, 1.75, 2.25, 2.75, 3.25})
}

func (c Chorus) Refrain6(startBeat float64) {
	ctx := c.Offset(startBeat)

	c.Sweep(ctx, 0, 0.75, sukoyaGradient, true)
	c.SweepSpin(ctx, 0.75, true)
	c.HalfCollapse(ctx, 2.75, []float64{0}, true)
	c.FadeToSukoya(ctx, -0.25, []float64{1.25, 2.00, 2.50, 3.00})
	c.FadeToGold(ctx, -0.50, []float64{4.0, 4.5, 5.0, 5.5})
}

func (c Chorus) Refrain7(startBeat float64) {
	ctx := c.Offset(startBeat)

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

	l := transform.Light(
		light.NewBasic(ctx, evt.RingLights),
		transform.Reverse(),
		transform.DivideSingle(),
	)

	rng := timer.NewRanger(-1, 0, 45, ease.InExpo)
	ctx.Range(rng, func(ctx context.Context) {
		grad := shirayukiGradient
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.BiasedColorSweep(ctx, 4, grad)
			fx.Ripple(ctx, e, 1.5)
			fx.AlphaFadeEx(ctx, e, 0, 0.6, 2, 10, ease.OutCubic)
			fx.AlphaFadeEx(ctx, e, 0.6, 1.0, 10, 0, ease.OutCubic)
		})
	})

	seq := timer.NewSequencer([]float64{1, 1.75, 2.5, 3.0, 3.5, 4.5}, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
		grad := sukoyaGradient
		c.Rush(ctx, ctx.B(), ctx.B()+0.4, 0.4, 2*float64(ctx.Ordinal()), grad)
		//c.RushNoFade(ctx, ctx.B(), ctx.B()+0.4, 0.4, 2*float64(ctx.Ordinal()), grad)
		ctx.NewPreciseZoom(evt.WithZoomStep(-0.33))
	})

	seq = timer.NewSequencer([]float64{2, 2.75, 3.5}, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(30*float64(ctx.Ordinal())),
			evt.WithRotationStep(7*float64(ctx.Ordinal()*5)),
			evt.WithRotationSpeed(20),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)
	})

	ctx.Sequence(timer.Beat(3.5), func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(-1))
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

		rng := timer.NewRanger(0, 4, 60, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(bl, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 1.2, magnetRainbow)
				fx.AlphaFadeEx(ctx, e, 0, 1, 6, 0.0, ease.InOutQuad)
			})
		})
	})
}

func (c Chorus) Rhythm2(startBeat float64) {
	ctx := c.Offset(startBeat)

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

	ctx.Range(timer.NewRanger(startBeat, endBeat, 30, ease.OutCubic), func(ctx context.Context) {
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

func (c Chorus) HalfCollapse(ctx context.Context, startBeat float64, sequence []float64, reverse bool) {
	ctx = ctx.Offset(startBeat)

	dir := chroma.Clockwise.ReverseIf(reverse)

	seq := timer.NewSequencer(sequence, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(9),
			evt.WithRotationSpeed(8),
			evt.WithProp(0.8),
			evt.WithRotationDirection(dir),
		)
	})
}

func (c Chorus) Unsweep(ctx context.Context, startBeat float64, sequence []float64) {
	ctx = ctx.Offset(startBeat)

	seq := timer.NewSequencer(sequence, 0)
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

func (c Chorus) Rush(ctx context.Context, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	l := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.Reverse(),
		transform.DivideSingle(),
	)

	rng := timer.NewRanger(startBeat, endBeat, 45, ease.InExpo)

	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, 2, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaFadeEx(ctx, e, 0, 0.6, 2, peakAlpha, ease.OutCubic)
			fx.AlphaFadeEx(ctx, e, 0.6, 1.0, peakAlpha, 0, ease.OutCubic)
		})
	})
}

func (c Chorus) RushNoFade(ctx context.Context, startBeat, endBeat, step, peakAlpha float64, grad gradient.Table) {
	l := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.Reverse(),
		transform.DivideSingle(),
	)

	rng := timer.NewRanger(startBeat, endBeat, 45, ease.InExpo)

	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.BiasedColorSweep(ctx, 4, grad)
			fx.Ripple(ctx, e, step)
			fx.AlphaFadeEx(ctx, e, 0, 0.6, 2, peakAlpha, ease.OutCubic)
		})
	})
}

func (c Chorus) FadeToSukoya(ctx context.Context, startBeat float64, sequence []float64) {
	seq := timer.NewSequencer(sequence, 0)

	lightSweepDiv := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(seq.Len()).Sequence(),
	).(light.Sequence)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.Sequence(seq, func(ctx context.Context) {
		rng := timer.NewRanger(0, 0.5, 16, ease.OutCubic)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(lightSweepDiv, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 3.6, gradient.FromSet(sukoyaColors))
				fx.AlphaFadeEx(ctx, e, 0, 1, 0.3, 3, ease.OutElastic)
			})
		})
	})
}

func (c Chorus) FadeToGold(ctx context.Context, startBeat float64, sequence []float64) {
	seq := timer.NewSequencer(sequence, 0)

	lightSweepDiv := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Rotate(3),
		transform.Divide(seq.Len()).Sequence(),
	).(light.Sequence)

	lightSweepDiv = lightSweepDiv.Shuffle()

	ctx.Sequence(seq, func(ctx context.Context) {
		grad := magnetRainbowPale.RotateRand()
		rng := timer.NewRanger(0, 0.5, 16, ease.OutCubic)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(lightSweepDiv, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 4, grad)
				fx.AlphaFadeEx(ctx, e, 0, 1, 10, 0, ease.OutElastic)
			})
		})
	})
}
