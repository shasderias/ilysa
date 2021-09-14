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

type GuitarSolo struct {
	context.Context
	p *ilysa.Project
	gradient.Set
}

func NewGuitarSolo(p *ilysa.Project, startBeat float64) GuitarSolo {
	gradSet := gradient.NewSet(
		magnetGradient,
		magnetRainbow,
	)
	return GuitarSolo{p.BOffset(startBeat), p, gradSet}
}

func (g GuitarSolo) Play() {
	g.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser),
			evt.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
		)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser),
			evt.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
		)
	})

	g.Beat(0, true)
	g.Beat(8, false)
	g.SyncopatedDrums(14, timer.Seq([]float64{0.5, 1.0, 1.25, 1.75}, 0))
	//g.Beat(16, true)

	ctx := g.BOffset(0)
	g.Solo(ctx, timer.Interval(0.50, 0.25, 6), false)
	g.Solo(ctx, timer.Interval(2.25, 0.25, 7), true)
	g.Solo(ctx, timer.Interval(4.25, 0.25, 7), false)
	g.Solo(ctx, timer.Interval(6.25, 0.25, 8), true)
	g.Solo(ctx, timer.Interval(8.50, 0.25, 15), false)
	g.Solo(ctx, timer.Interval(12.50, 0.25, 2), true)
	g.Solo(ctx, timer.Interval(13.25, 0.25, 2), false)
	g.Solo(ctx, timer.Interval(13.75, 0.125, 2), true)
	g.Solo(ctx, timer.Interval(14.00, 0.50, 3), false)
	g.Solo(ctx, timer.Interval(15.50, 0.25, 2), false)

	g.Beat2(16)
	g.Beat2(24)

	g.PianoDoubles(ctx, 16,
		timer.Seq([]float64{0.25, 1.00, 1.75, 2.50, 3.25, 4.00, 4.50, 4.75, 5.25, 5.75}, 6.00),
		[]float64{6, 5, 4, 3, 2, 1, 2, 3, 3, 3},
	)

	g.BrokenChords(ctx, 22,
		timer.Seq([]float64{0.25, 0.50, 0.75}, 0),
		timer.Seq([]float64{1.00}, 0),
		timer.Seq([]float64{1.25, 1.50, 1.75}, 0),
	)

	g.PianoDoubles(ctx, 24,
		timer.Seq([]float64{0.00, 0.50, 0.75, 1.50, 2.25}, 2.50),
		[]float64{4, 3, 5, 6, 7},
	)

	g.BrokenChords(ctx, 27,
		timer.Seq([]float64{0.25, 0.50, 0.75}, 1.00),
		timer.Seq([]float64{1.00}, 1.25),
		timer.Seq([]float64{1.50, 1.75}, 2.00),
	)

	g.Finale(ctx, 29,
		timer.Seq([]float64{0.25, 0.50, 0.75, 1.00}, 1.25),
		timer.Seq([]float64{1.50, 2.00, 2.50}, 3.00),
	)

	//g.PianoDoubles(ctx, 28,
	//	timer.Seq([]float64{0.00}, 2.50),
	//	[]float64{5},
	//)
}

func (g GuitarSolo) Beat(startBeat float64, transition bool) {
	ctx := g.BOffset(startBeat)

	rl := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.DivideSingle(),
	)

	gradSet := gradient.NewSet(shirayukiWhiteGradient, sukoyaWhiteGradient)

	ctx.Sequence(timer.Interval(0, 2, 4), func(ctx context.Context) {
		centerOn(ctx, magnetColors.Next())
		ctx.NewPreciseZoom(evt.WithZoomStep(float64(ctx.SeqOrdinal())))

		grad := gradSet.Next()
		ctx.Range(timer.Rng(0, 1.75, 15, ease.Linear), func(ctx context.Context) {
			ctx.Light(rl, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				fx.AlphaFadeEx(ctx, e, 0, 1, 1, 0, ease.InOutQuad)
			})
		})

		if !transition || ctx.SeqOrdinal() > 2 {
			ctx.NewPreciseRotation(
				evt.WithRotation(15),
				evt.WithRotationStep(15*float64(ctx.Ordinal())),
				evt.WithRotationSpeed(4),
				evt.WithProp(0.8),
				evt.WithRotationDirection(chroma.CounterClockwise),
			)
		}
	})
}

func (g GuitarSolo) Beat2(startBeat float64) {
	ctx := g.BOffset(startBeat)

	rl := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.DivideSingle(),
	)

	gradSet := gradient.NewSet(shirayukiWhiteGradient, sukoyaWhiteGradient)

	ctx.Sequence(timer.Interval(0, 2, 4), func(ctx context.Context) {
		centerOn(ctx, magnetColors.Next())
		ctx.NewPreciseZoom(evt.WithZoomStep(float64(ctx.SeqOrdinal())))

		grad := gradSet.Next()
		ctx.Range(timer.Rng(0, 1.75, 15, ease.Linear), func(ctx context.Context) {
			ctx.Light(rl, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				fx.AlphaFadeEx(ctx, e, 0, 1, 1, 0, ease.InOutQuad)
			})
		})
	})
}

func (g GuitarSolo) SyncopatedDrums(startBeat float64, seq timer.Sequencer) {
	ctx := g.BOffset(startBeat)

	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.DivideSingle(),
	)

	ctx.Sequence(seq, func(ctx context.Context) {
		dir := chroma.Clockwise
		if ctx.T() > 0.5 {
			dir = chroma.CounterClockwise
		}

		if ctx.First() {
			ctx.Range(timer.Rng(0, 1.75, 15, ease.Linear), func(ctx context.Context) {
				ctx.Light(l, func(ctx context.LightContext) {
					e := fx.Gradient(ctx, magnetGradient)
					fx.AlphaFadeEx(ctx, e, 0, 1, 3, 0, ease.InCirc)
				})
			})
		}

		ctx.NewPreciseRotation(
			evt.WithRotation(15),
			evt.WithRotationStep(15*(1-ctx.T())),
			evt.WithRotationSpeed(20),
			evt.WithProp(20),
			evt.WithRotationDirection(dir),
		)
	})
}

func (g GuitarSolo) Solo(ctx context.Context, seq timer.Sequencer, reverse bool) {
	var (
		llReverse transform.LightTransformer = transform.Identity()
		rlReverse transform.LightTransformer = transform.Identity()
	)

	if reverse {
		llReverse = transform.Reverse()
	} else {
		rlReverse = transform.Reverse()
	}

	ll := transform.Light(light.NewBasic(ctx, evt.LeftRotatingLasers),
		llReverse,
		transform.DivideSingle().Sequence(),
	).(light.Sequence)

	rl := transform.Light(light.NewBasic(ctx, evt.RightRotatingLasers),
		rlReverse,
		transform.DivideSingle().Sequence(),
	).(light.Sequence)

	l := light.NewSequence()

	for i := 0; i < ll.Len(); i++ {
		l.Add(light.Combine(ll.Idx(i), rl.Idx(i)))
	}

	rng := timer.Rng(seq.Idx(0), seq.Idx(seq.Len()-1), 31, ease.InOutCirc)
	fx.SlowMotionLasers(ctx, rng, evt.LeftLaser, 32, 1,
		fx.WithPreciseLaserOpts(evt.WithLaserDirection(chroma.Clockwise)))
	fx.SlowMotionLasers(ctx, rng, evt.RightLaser, 32, 1,
		fx.WithPreciseLaserOpts(evt.WithLaserDirection(chroma.Clockwise)))

	grad1 := g.Set.Next()
	grad2 := g.Set.Next()
	g.Set.Next()

	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.Range(timer.Rng(0, 1.25, 30, ease.Linear), func(ctx context.Context) {
			ctx.Light(ll, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 1.9, grad1)
				fx.AlphaFadeEx(ctx, e, 0, 1, 3, 0, ease.InCubic)
			})
			ctx.Light(rl, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 1.9, grad2)
				fx.AlphaFadeEx(ctx, e, 0, 1, 3, 0, ease.InCubic)
			})
		})
	})
}

func (g GuitarSolo) PianoDoubles(ctx context.Context, startBeat float64, seq timer.Sequencer, steps []float64) {
	ctx = ctx.BOffset(startBeat)

	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.DivideSingle(),
		transform.Flatten(),
		transform.Divide(3).Sequence(),
	).(light.Sequence).Shuffle()

	ctx.Sequence(seq, func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(15),
			evt.WithRotationStep(steps[ctx.Ordinal()]*3),
			evt.WithRotationSpeed(20),
			evt.WithProp(20),
			evt.WithRotationDirection(chroma.Clockwise),
		)

		ctx.Light(l, func(ctx context.LightContext) {
			fx.Gradient(ctx, magnetGradient)
			ctx.NewRGBLighting(
				evt.WithLightValue(evt.LightOff),
			).Apply(evt.WithBOffset(ctx.SeqNextBOffset()))
		})
	})
}

func (g GuitarSolo) BrokenChords(ctx context.Context, startBeat float64, seqUp, seqMid, seqDown timer.Sequencer) {
	ctx = ctx.BOffset(startBeat)

	upDownLights := func(divisor int) context.Light {
		return transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.Fan(2),
			transform.DivideSingle(),
			transform.Flatten(),
			transform.Divide(divisor).Sequence(),
		)
	}

	midLight := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.DivideSingle(),
	)

	ctx.Sequence(timer.Seq([]float64{seqUp.Idx(0)}, 0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(12),
			evt.WithRotationSpeed(20),
			evt.WithProp(20),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)
	})

	ctx.Sequence(seqUp, func(ctx context.Context) {
		if ctx.First() {
			ctx.NewPreciseZoom(evt.WithZoomStep(-1))
		}
		ctx.Light(upDownLights(seqUp.Len()), func(ctx context.LightContext) {
			fx.Gradient(ctx, magnetGradient)
		})
	})

	ctx.Sequence(seqMid, func(ctx context.Context) {
		ctx.Light(midLight, func(ctx context.LightContext) {
			fx.Gradient(ctx, magnetRainbowPale)
		})
	})

	ctx.Sequence(seqDown, func(ctx context.Context) {
		if ctx.First() {
			ctx.NewPreciseZoom(evt.WithZoomStep(0))
		}
		ctx.Light(upDownLights(seqDown.Len()), func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithLightValue(evt.LightOff))
		})
	})
}
func (g GuitarSolo) Finale(ctx context.Context, startBeat float64, seqUp, seqDown timer.Sequencer) {
	ctx = ctx.BOffset(startBeat)

	upDownLights := func(divisor int) context.Light {
		return transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.Fan(2),
			transform.DivideSingle(),
			transform.Flatten(),
			transform.Divide(divisor).Sequence(),
			transform.DivideSingle(),
		)
	}

	ctx.Sequence(timer.Seq([]float64{seqUp.Idx(0)}, 0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(90),
			evt.WithRotationStep(12),
			evt.WithRotationSpeed(20),
			evt.WithProp(20),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)
	})

	ctx.Sequence(seqUp, func(ctx context.Context) {
		ctx.Light(upDownLights(seqUp.Len()), func(ctx context.LightContext) {
			fx.Gradient(ctx, magnetGradient)
		})
	})

	ctx.Sequence(seqDown, func(ctx context.Context) {
		ctx.Light(upDownLights(seqDown.Len()), func(ctx context.LightContext) {
			fx.Gradient(ctx, magnetRainbowPale)
		})
	})
}
