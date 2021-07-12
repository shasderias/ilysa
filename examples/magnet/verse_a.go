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

type VerseA struct {
	context.Context
	project *ilysa.Project
	offset  float64
}

func NewVerseA(p *ilysa.Project, offset float64) VerseA {
	ctx := p.Offset(offset)
	return VerseA{
		Context: ctx,
		project: p,
		offset:  offset,
	}
}

func (p VerseA) Play1() {
	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		fx.OffAll(ctx)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser), evt.WithPreciseLaserSpeed(1.5))
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser), evt.WithPreciseLaserSpeed(1.5))
		ctx.NewPreciseZoom(evt.WithZoomStep(0))
	})

	p.Rhythm(0, true)
	p.Rhythm(4, false)
	p.Rhythm(8, false)
	p.Rhythm(12, false)
	p.Rhythm(16, false)
	p.Rhythm(20, false)
	p.Rhythm(24, false)
	p.Rhythm(28, true)

	p.PianoBackstep(7, shirayukiWing, shirayukiWing)
	p.PianoBackstep(15, shirayukiWing, shirayukiWing)
	p.PianoBackstep(23, shirayukiWing, shirayukiWing)
	p.RinPun(27)
}

func (p VerseA) Play2() {
	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		fx.OffAll(ctx)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser),
			evt.WithPreciseLaserSpeed(3.5))
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser),
			evt.WithPreciseLaserSpeed(3.5))
	})

	p.Rhythm(0, true)
	p.Rhythm(4, false)
	p.Rhythm(8, false)
	p.Rhythm(12, false)
	p.Rhythm2(16, false)
	p.Rhythm2(20, false)
	p.Rhythm2(24, false)
	p.Rhythm2(28, true)

	p.PianoBackstep(7, shirayukiWing, shirayukiWing)
	//p.PianoBackstep(15)
	p.SnareDrums(14, timer.Seq([]float64{0.25, 0.50, 1.00, 1.25, 1.75, 2.00}, 2.25))
	p.PianoBackstep(23, shirayukiWing, shirayukiWing)
	p.RinPun(27)
}

func (p VerseA) Rhythm(startBeat float64, kickOnly bool) {
	ctx := p.BOffset(startBeat)
	var (
		kickDrumLight = light.NewSequence(
			light.NewBasic(ctx, evt.LeftRotatingLasers),
			light.NewBasic(ctx, evt.RightRotatingLasers),
		)
		kickDrumSequence = timer.Seq([]float64{0, 2.5}, 0)
		kickDrumColors   = magnetColors
	)

	ctx.Sequence(kickDrumSequence, func(ctx context.Context) {
		ctx.Light(kickDrumLight, func(ctx context.LightContext) {
			ctx.NewRGBLighting(
				evt.WithLightValue(evt.LightRedFade),
				evt.WithColor(kickDrumColors.Next()),
				evt.WithAlpha(0.7),
			)
		})
	})

	if kickOnly {
		return
	}

	var (
		rng  = timer.Rng(0, 2, 30, ease.Linear)
		grad = gradient.Table{
			{shirayukiPurple, 0.0},
			{shirayukiGold, 0.3},
			{shirayukiGold, 0.7},
			{shirayukiPurple, 1.0},
		}
	)

	RingRipple(ctx, rng, grad,
		WithRippleTime(0.6),
		WithSweepSpeed(1.8),
		WithFadeIn(fx.NewAlphaFader(0, 0.3, 0, 0.7, ease.OutQuart)),
		WithFadeOut(fx.NewAlphaFader(0.3, 1, 0.7, 0, ease.InCirc)),
	)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(22.5),
			evt.WithRotationSpeed(2),
			evt.WithProp(0.5),
		)
	})
}

func (p VerseA) Rhythm2(startBeat float64, kickOnly bool) {
	ctx := p.BOffset(startBeat)
	var (
		kickDrumLight = light.NewSequence(
			light.NewBasic(ctx, evt.LeftRotatingLasers),
			light.NewBasic(ctx, evt.RightRotatingLasers),
		)
		kickDrumSequence = timer.Seq([]float64{0, 1, 2, 2.5, 3}, 0)
		kickDrumColors   = magnetColors
	)

	ctx.Sequence(kickDrumSequence, func(ctx context.Context) {
		ctx.Light(kickDrumLight, func(ctx context.LightContext) {
			ctx.NewRGBLighting(
				evt.WithLightValue(evt.LightRedFade),
				evt.WithColor(kickDrumColors.Next()),
				evt.WithAlpha(0.7),
			)
		})
	})

	if kickOnly {
		return
	}

	var (
		rng  = timer.Rng(0, 2, 30, ease.Linear)
		grad = gradient.Table{
			{shirayukiPurple, 0.0},
			{shirayukiGold, 0.3},
			{shirayukiGold, 0.7},
			{shirayukiPurple, 1.0},
		}
	)

	RingRipple(ctx, rng, grad,
		WithRippleTime(0.6),
		WithSweepSpeed(1.8),
		WithFadeIn(fx.NewAlphaFader(0, 0.3, 0, 0.7, ease.OutQuart)),
		WithFadeOut(fx.NewAlphaFader(0.3, 1, 0.7, 0, ease.InCirc)),
	)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(22.5),
			evt.WithRotationSpeed(2),
			evt.WithProp(0.5),
		)
	})
}
func (p VerseA) PianoBackstep(startBeat float64, lWingGrad, rWingGrad gradient.Table) {
	ctx := p.BOffset(startBeat)

	var (
		sequence = timer.Seq([]float64{0, 0.5}, 2)
		bl       = transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.Fan(2),
			transform.Flatten(),
			transform.Divide(2).Sequence(),
		).(light.Sequence)
		lWing = transform.Light(bl.Idx(0),
			transform.Reverse(),
			transform.DivideSingle(),
		)

		rWing = transform.Light(bl.Idx(1),
			transform.DivideSingle(),
		)
	)

	ctx.Sequence(sequence, func(ctx context.Context) {
		rotOpt := evt.NewOpts(
			evt.WithRotation(15),
			evt.WithRotationStep(7.5),
			evt.WithRotationSpeed(10),
			evt.WithProp(20),
		)

		if ctx.Ordinal()%2 == 0 {
			ctx.NewPreciseRotation(evt.WithRotationDirection(chroma.Clockwise), rotOpt)
		} else {
			ctx.NewPreciseRotation(evt.WithRotationDirection(chroma.CounterClockwise), rotOpt)
		}

		if ctx.SeqFirst() {
			ctx.Light(lWing, func(ctx context.LightContext) {
				fx.Gradient(ctx, lWingGrad)
			})
			ctx.Light(rWing, func(ctx context.LightContext) {
				fx.Gradient(ctx, rWingGrad)
			})
		} else {
			ctx.Range(timer.Rng(0, ctx.SeqNextBOffset(), 18, ease.Linear), func(ctx context.Context) {
				ctx.Light(lWing, func(ctx context.LightContext) {
					e := fx.Gradient(ctx, lWingGrad)
					fx.AlphaFadeEx(ctx, e, 0, 1, 0.9, 0, ease.OutCubic)
				})
				ctx.Light(rWing, func(ctx context.LightContext) {
					e := fx.Gradient(ctx, rWingGrad)
					fx.AlphaFadeEx(ctx, e, 0, 1, 0.9, 0, ease.OutCubic)
				})
			})
		}
	})
}

func (p VerseA) SnareDrums(startBeat float64, seq timer.Sequencer) {
	ctx := p.BOffset(startBeat)

	l := light.Combine(
		transform.Light(light.NewBasic(ctx, evt.LeftRotatingLasers),
			transform.DivideSingle(),
		),
		transform.Light(light.NewBasic(ctx, evt.RightRotatingLasers),
			transform.DivideSingle(),
		),
	)

	gradSet := gradient.NewSet(
		magnetGradient,
		magnetRainbowPale,
		shirayukiGradient,
		magnetRainbow,
		shirayukiWhiteGradient,
		magnetRainbowPale,
	)

	ctx.Sequence(seq, func(ctx context.Context) {
		fx.SlowMotionLasers(ctx, timer.Rng(0, 0.25, 8, ease.InSin),
			evt.LeftLaser, 15, 1)
		fx.SlowMotionLasers(ctx, timer.Rng(0, 0.25, 8, ease.InSin),
			evt.RightLaser, 15, 1)

		grad := gradSet.Index(ctx.SeqOrdinal())

		ctx.Light(l, func(ctx context.LightContext) {
			fx.Gradient(ctx, grad)
		})

		if ctx.Last() {
			ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser),
				evt.WithBOffset(0.26),
				evt.WithLaserSpeed(8), evt.WithPreciseLaserSpeed(8))
			ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser),
				evt.WithBOffset(0.26),
				evt.WithLaserSpeed(8), evt.WithPreciseLaserSpeed(8))
		}
	})

	//lightFn := func(steps int) context.Light {
	//	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
	//		transform.Fan(2),
	//		transform.DivideSingle(),
	//		transform.RotateSet(steps),
	//		transform.Flatten(),
	//		transform.Divide(3).Sequence(),
	//		transform.DivideSingle(),
	//	).(light.Sequence).Idx(0)
	//	fmt.Println(l)
	//	return l
	//}
	//
	//gradSet := gradient.NewSet(magnetGradient, shirayukiGradient)
	//
	//rotateSteps := 0
	//ctx.Sequence(timer.Seq(sequence, 2.25), func(ctx context.Context) {
	//	grad := gradSet.Next()
	//
	//	if ctx.Ordinal()%2 == 0 {
	//		//ctx.NewPreciseRotation(
	//		//	evt.WithRotation(15),
	//		//	evt.WithRotationStep(15),
	//		//	evt.WithProp(20),
	//		//	evt.WithRotationSpeed(20),
	//		//	evt.WithRotationDirection(chroma.CounterClockwise),
	//		//)
	//		ctx.Light(lightFn(rotateSteps), func(ctx context.LightContext) {
	//			e := fx.Gradient(ctx, grad)
	//			e.Apply(evt.WithAlpha(2))
	//			oe := ctx.NewRGBLighting(evt.WithLightValue(evt.LightOff))
	//			oe.Apply(evt.WithBOffset(ctx.SeqNextBOffset() - 0.05))
	//		})
	//		rotateSteps += 2
	//	} else {
	//		//ctx.NewPreciseRotation(
	//		//	evt.WithRotation(15),
	//		//	evt.WithRotationStep(15),
	//		//	evt.WithProp(20),
	//		//	evt.WithRotationSpeed(20),
	//		//	evt.WithRotationDirection(chroma.CounterClockwise),
	//		//)
	//		ctx.Light(lightFn(rotateSteps), func(ctx context.LightContext) {
	//			fx.Gradient(ctx, grad)
	//			oe := ctx.NewRGBLighting(evt.WithLightValue(evt.LightOff))
	//			oe.Apply(evt.WithBOffset(ctx.SeqNextBOffset()))
	//		})
	//	}
	//})
}

func (p VerseA) RinPun(startBeat float64) {
	ctx := p.BOffset(startBeat)
	backLasers := light.NewBasic(ctx, evt.BackLasers)

	sl := transform.Light(backLasers,
		transform.Fan(2).Sequence(),
		transform.DivideSingle(),
	).(light.Sequence)

	sl = light.NewSequence(
		sl.Idx(0),
		transform.Light(sl.Idx(1),
			transform.ReverseSet(),
		))

	ctx.Sequence(timer.Seq([]float64{0, 0.5}, 0), func(ctx context.Context) {
		if ctx.SeqFirst() {
			ctx.NewPreciseRotation(
				evt.WithRotation(45),
				evt.WithRotationStep(11.5),
				evt.WithRotationSpeed(2),
				evt.WithProp(20),
				evt.WithRotationDirection(chroma.CounterClockwise),
			)
		}
		ctx.Range(timer.Rng(-0.3, 0.7, 30, ease.Linear), func(ctx context.Context) {
			ctx.Light(sl, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 0.4, magnetRainbow)
				fx.RippleT(ctx, e, 0.30)
				fx.AlphaFadeEx(ctx, e, 0, 0.5, 0, 1, ease.InCubic)
				fx.AlphaFadeEx(ctx, e, 0.5, 1.0, 1, 0.7, ease.OutCirc)
			})
		})
	})

	ctx.Sequence(timer.Seq([]float64{1, 2}, 0), func(ctx context.Context) {
		if ctx.SeqFirst() {
			ctx.NewPreciseRotation(
				evt.WithRotation(45),
				evt.WithRotationStep(11.5),
				evt.WithRotationSpeed(2),
				evt.WithProp(20),
				evt.WithRotationDirection(chroma.Clockwise),
			)
		}

		ctx.Range(timer.Rng(-0.2, 0.5, 30, ease.InOutCirc), func(ctx context.Context) {
			ctx.Light(sl, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 0.3, magnetRainbow)
				fx.RippleT(ctx, e, 0.65)
				fx.AlphaFadeEx(ctx, e, 0, 0.5, 0, 2, ease.InCubic)
				fx.AlphaFadeEx(ctx, e, 0.8, 1.0, 2, 0, ease.OutCirc)
			})
		})
	})

	// tsuketa
	{
		sl := transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.Fan(2),
			transform.Flatten(),
			transform.Divide(3).Sequence(),
			transform.DivideSingle(),
		).(light.Sequence)

		sl = light.NewSequence(
			sl.Idx(1),
			sl.Idx(0),
			sl.Idx(2),
		)

		ctx.Sequence(timer.Seq([]float64{3, 3.5, 3.75}, 0), func(ctx context.Context) {
			if ctx.SeqFirst() {
				ctx.NewPreciseRotation(
					evt.WithRotation(22.5),
					evt.WithRotationStep(11.5),
					evt.WithRotationSpeed(2),
					evt.WithProp(20),
					evt.WithRotationDirection(chroma.CounterClockwise),
				)
			}
			ctx.Range(timer.Rng(-0.1, 0.2, 30, ease.InOutCirc), func(ctx context.Context) {
				ctx.Light(sl, func(ctx context.LightContext) {
					e := fx.ColorSweep(ctx, 0.8, magnetRainbow)
					fx.RippleT(ctx, e, 0.5)
					fx.AlphaFadeEx(ctx, e, 0, 0.5, 0, 2, ease.OutCirc)
					fx.AlphaFadeEx(ctx, e, 0.5, 1.0, 2, 0, ease.InCirc)
					fx.AlphaShimmer(ctx, e, 2)
				})
			})
		})
	}
}
