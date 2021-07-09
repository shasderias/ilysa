package main

import (
	"math"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

func NewIntro(p *ilysa.Project, startBeat float64) Intro {
	return Intro{
		p:       p,
		Context: p.Offset(startBeat),
	}

}

type Intro struct {
	p *ilysa.Project
	context.Context
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
	ctx := p.Offset(startBeat)

	colors := colorful.NewSet(magnetPurple, magnetWhite, colorful.Black, magnetPink)

	ringBackCombined := light.Combine(
		transform.Light(light.NewBasic(ctx, evt.RingLights),
			transform.DivideSingle(),
		),
		transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.DivideSingle(),
		),
	)

	seq := timer.NewSequencer([]float64{0, 0.75, 1.25, 1.75}, 2.25)

	ctx.Sequence(seq, func(ctx context.Context) {
		grad := gradient.New(colors.Idx(ctx.Ordinal()), colors.Idx(ctx.Ordinal()+1))
		ctx.Light(ringBackCombined, func(ctx context.LightContext) {
			fx.Gradient(ctx, grad)
		})
		if ctx.Last() {
			opts := evt.NewOpts(evt.WithBeatOffset(ctx.SeqNextBOffset()), evt.WithLightValue(evt.LightOff))
			ctx.NewRGBLighting(evt.WithLight(evt.RingLights), opts)
			ctx.NewRGBLighting(evt.WithLight(evt.BackLasers), opts)
		}
	})
}

func (p Intro) LeadinDrums(startBeat float64) {
	ctx := p.Offset(startBeat)

	l := light.NewSequence(
		transform.Light(light.NewBasic(ctx, evt.LeftRotatingLasers),
			transform.Fan(2),
		),
		transform.Light(light.NewBasic(ctx, evt.RightRotatingLasers),
			transform.Fan(2),
		),
	)

	seq := timer.NewSequencer([]float64{0, 0.25, 0.75, 1, 1.5}, 1.75)

	ctx.Sequence(seq, func(ctx context.Context) {
		fx.ZeroSpeedRandomizedLasers(ctx, evt.LeftLaser)
		fx.ZeroSpeedRandomizedLasers(ctx, evt.RightLaser)

		ctx.Light(l, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithColor(crossickColors.Next()))
			ctx.NewRGBLighting(evt.WithBeatOffset(ctx.SeqNextBOffset()), evt.WithLightValue(evt.LightOff))
		})
	})
}

func (p Intro) BassTwang(startBeat float64) {
	ctx := p.Offset(startBeat)

	var (
		backLasers = transform.Light(
			light.NewBasic(ctx, evt.BackLasers),
			transform.DivideSingle(),
		)
		duration   = 1.5
		steps      = 45
		sweepSpeed = 2.2
		intensity  = 4.0
		grad       = magnetRainbowPale
	)

	ctx.NewPreciseRotation(
		evt.WithRotation(135),
		evt.WithRotationStep(13.5),
		evt.WithProp(20),
		evt.WithRotationSpeed(2.4),
		evt.WithRotationDirection(chroma.CounterClockwise),
		evt.WithCounterSpin(false),
	)

	ctx.NewPreciseZoom(
		evt.WithZoomStep(-1),
	)

	ctx.Range(timer.NewRanger(0, duration, steps, ease.Linear), func(ctx context.Context) {
		ctx.Light(backLasers, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, sweepSpeed, grad)
			fx.AlphaFadeEx(ctx, e, 0, 0.35, 0, intensity, ease.OutCubic)
			fx.AlphaFadeEx(ctx, e, 0.35, 1, intensity, 0, ease.InSin)
		})
	})
}

func (p Intro) StartSplash(startBeat float64) {
	ctx := p.Offset(startBeat)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewLaser(evt.WithDirectionalLaser(evt.LeftLaser), evt.WithIntValue(8))
		ctx.NewLaser(evt.WithDirectionalLaser(evt.RightLaser), evt.WithIntValue(8))

		ctx.NewRGBLighting(
			evt.WithLight(evt.LeftRotatingLasers), evt.WithLightValue(evt.LightBlueFlash),
			evt.WithColor(sukoyaPink))
		ctx.NewRGBLighting(
			evt.WithLight(evt.RightRotatingLasers), evt.WithLightValue(evt.LightRedFlash),
			evt.WithColor(shirayukiPurple))
		ctx.NewRGBLighting(
			evt.WithLight(evt.CenterLights), evt.WithLightValue(evt.LightBlueFlash),
			evt.WithColor(magnetPurple))

		ctx.NewPreciseZoom(evt.WithZoomStep(0))
	})
}

func (p Intro) Rhythm(startBeat float64) {
	ctx := p.Offset(startBeat)

	l := transform.Light(
		light.NewBasic(ctx, evt.RingLights),
		transform.DivideSingle(),
	)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(2.5),
			evt.WithProp(1.5),
			evt.WithRotationSpeed(24),
			evt.WithRotationDirection(chroma.Clockwise),
		)

		grad := gradient.New(sukoyaColors.Rand(), shirayukiColors.Rand())

		ctx.Range(timer.NewRanger(ctx.B(), ctx.B()+0.1, 12, ease.InCubic), func(ctx context.Context) {
			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				fx.RippleT(ctx, e, 0.9, fx.EaseT(ease.InExpo))
			})
		})
	})

	ctx.Sequence(timer.NewSequencer([]float64{1, 3}, 0), func(ctx context.Context) {
		ctx.Range(timer.NewRanger(0, 0.8, 12, ease.InSin), func(ctx context.Context) {
			speed := (1-ctx.T())*12 + 1
			intSpeed := int(math.Round(speed))

			opts := evt.NewOpts(evt.WithLaserSpeed(intSpeed), evt.WithPreciseLaserSpeed(speed))
			if !ctx.First() {
				opts.Add(evt.WithLockPosition(true))
			}

			ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser), opts)
			ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser), opts)
		})

		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(12.5),
			evt.WithProp(20),
			evt.WithRotationSpeed(20),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)
	})

	ctx.Sequence(timer.Beat(1), func(ctx context.Context) {
		ctx.NewRGBLighting(
			evt.WithLight(evt.LeftRotatingLasers),
			evt.WithLightValue(evt.LightRedFade),
			evt.WithColor(sukoyaPink),
		)

		ctx.NewRGBLighting(
			evt.WithLight(evt.RightRotatingLasers),
			evt.WithLightValue(evt.LightBlueFade),
			evt.WithColor(sukoyaWhite),
		)
	})

	ctx.Sequence(timer.Beat(3), func(ctx context.Context) {
		ctx.NewRGBLighting(
			evt.WithLight(evt.RingLights),
			evt.WithLightValue(evt.LightBlueFade),
			evt.WithColor(magnetColors.Next()),
		)

		ctx.NewRGBLighting(
			evt.WithLight(evt.LeftRotatingLasers),
			evt.WithLightValue(evt.LightBlueFade),
			evt.WithColor(shirayukiPurple),
			evt.WithAlpha(3),
		)

		ctx.NewRGBLighting(
			evt.WithLight(evt.RightRotatingLasers),
			evt.WithLightValue(evt.LightRedFade),
			evt.WithColor(shirayukiGold),
			evt.WithAlpha(3),
		)
	})
}

func (p Intro) Melody1(startBeat float64) {
	ctx := p.Offset(startBeat)

	seq := timer.NewSequencer([]float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75}, 0)
	p.PianoGlow(ctx, seq, 5, 0.435, false, false)
}

func (p Intro) Melody2(startBeat float64, reverseZoom bool) {
	ctx := p.Offset(startBeat)

	seq := timer.NewSequencer([]float64{0, 0.25, 0.50}, 0)
	p.PianoGlow(ctx, seq, 3, 0.2, false, true)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		if reverseZoom {
			ctx.NewPreciseZoom(evt.WithZoomStep(0))
		} else {
			ctx.NewPreciseZoom(evt.WithZoomStep(-0.5))
		}
	})
}

func (p Intro) Melody3(startBeat float64) {
	ctx := p.Offset(startBeat)

	seq := timer.NewSequencer([]float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75, 3.00, 3.25, 3.50}, 0)
	p.PianoGlow(ctx, seq, 5, 0.435, false, false)
}

func (p Intro) Chorus(startBeat float64) {
	ctx := p.Offset(startBeat)

	sequence := timer.NewSequencer([]float64{0, 1, 2, 2.75, 3.5, 4}, 4.5)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewRGBLighting(evt.WithLight(evt.LeftRotatingLasers), evt.WithLightValue(evt.LightOff))
		ctx.NewRGBLighting(evt.WithLight(evt.RightRotatingLasers), evt.WithLightValue(evt.LightOff))
	})

	ctx.Sequence(sequence, func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(-0.1 * float64(ctx.SeqOrdinal())))

		re := ctx.NewPreciseRotation(
			evt.WithRotation(45+15*float64(ctx.Ordinal())),
			evt.WithRotationStep(5+(1.5*float64(ctx.Ordinal()))),
			evt.WithProp(20),
			evt.WithRotationSpeed(4+float64(ctx.Ordinal())*2),
		)
		if ctx.Ordinal()%2 == 0 {
			re.Apply(evt.WithRotationDirection(chroma.Clockwise))
		} else {
			re.Apply(evt.WithRotationDirection(chroma.CounterClockwise))
		}

		var backLasers context.Light

		if ctx.Ordinal()%2 == 0 {
			backLasers = transform.Light(light.NewBasic(ctx, evt.BackLasers),
				transform.DivideSingle(),
			)
		} else {
			backLasers = transform.Light(light.NewBasic(ctx, evt.BackLasers),
				transform.Reverse(),
				transform.DivideSingle(),
			)
		}

		ctx.Light(backLasers, func(ctx context.LightContext) {
			e := fx.Gradient(ctx, magnetRainbowPale)
			fx.RippleT(ctx, e, ctx.SeqNextBOffset())
			fx.AlphaFadeEx(ctx, e, 0, 1, 4, 0.8, ease.OutCubic)

			if !ctx.Last() {
				oe := ctx.NewRGBLighting(evt.WithLightValue(evt.LightOff))
				oe.Apply(evt.WithBeatOffset(ctx.SeqNextBOffset()))
			}
		})
	})
}

func (p Intro) PianoRoll(startBeat float64, count int) {
	ctx := p.Offset(startBeat)

	seq := timer.NewSequencer([]float64{0, 0.25, 0.50, 0.75, 1.00, 1.25}, 0)
	p.PianoGlow(ctx, seq, 6, 0.380, false, false)
}

func (p Intro) Trill(startBeat float64) {
	ctx := p.Offset(startBeat)

	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Shuffle(),
		transform.DivideSingle(),
	)

	ctx.NewPreciseRotation(
		evt.WithRotation(180),
		evt.WithRotationStep(12.5),
		evt.WithRotationSpeed(12),
		evt.WithProp(6),
		evt.WithRotationDirection(chroma.Clockwise),
	)

	rng := timer.NewRanger(0, 1.0, 12, ease.Linear)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.Gradient(ctx, magnetRainbow)
			fx.RippleT(ctx, e, 0.30)
			fx.AlphaFadeEx(ctx, e, 0, 1, 0.8, 0, ease.InSin)
		})
	})
}

func (p Intro) TrillNoFade(startBeat float64) {
	ctx := p.Offset(startBeat)

	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.DivideSingle(),
	)

	ctx.NewPreciseRotation(
		evt.WithRotation(90),
		evt.WithRotationStep(5),
		evt.WithRotationSpeed(12),
		evt.WithProp(4),
		evt.WithRotationDirection(chroma.CounterClockwise),
	)

	rng := timer.NewRanger(0, 0.3, 12, ease.Linear)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.Gradient(ctx, magnetRainbow)
			fx.RippleT(ctx, e, 0.30)
			fx.AlphaFadeEx(ctx, e, 0, 1, 0, 0.8, ease.InSin)
		})
	})
}

func (p Intro) Climb(startBeat float64) {
	ctx := p.Offset(startBeat)

	var (
		step                    = 0.25
		count                   = 12
		rotatingLasersExitSpeed = 3.0

		blGrad = gradient.FromSet(crossickColors)
		rlGrad = magnetGradient

		l = transform.Light(
			light.NewBasic(ctx, evt.BackLasers),
			transform.Divide(7).Sequence(),
		)
	)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(360),
			evt.WithRotationStep(15),
			evt.WithRotationSpeed(1.3),
			evt.WithProp(13),
		)
		ctx.NewZoom()
	})

	interval := timer.Interval(0, step, count)

	ctx.Sequence(interval, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithColor(blGrad.Lerp(ctx.T())))
		})

		switch {
		case ctx.Last():
			ctx.NewRGBLighting(evt.WithLight(evt.LeftRotatingLasers), evt.WithLightValue(evt.LightBlueFade),
				evt.WithColor(magnetPurple),
			)
			ctx.NewRGBLighting(evt.WithLight(evt.RightRotatingLasers), evt.WithLightValue(evt.LightRedFade),
				evt.WithColor(magnetPurple),
			)

			exitArgs := []evt.PreciseLaserOpt{
				evt.WithLaserSpeed(int(rotatingLasersExitSpeed)),
				evt.WithPreciseLaserSpeed(rotatingLasersExitSpeed),
				evt.WithLockPosition(true),
			}

			ctx.NewPreciseLaser(
				append([]evt.PreciseLaserOpt{
					evt.WithDirectionalLaser(evt.LeftLaser),
					evt.WithLaserDirection(chroma.CounterClockwise),
				}, exitArgs...)...,
			)
			ctx.NewPreciseLaser(
				append([]evt.PreciseLaserOpt{
					evt.WithDirectionalLaser(evt.RightLaser),
					evt.WithLaserDirection(chroma.Clockwise),
				}, exitArgs...)...,
			)

		case ctx.Ordinal()%2 == 0:
			ctx.NewRGBLighting(
				evt.WithLight(evt.LeftRotatingLasers),
				evt.WithLightValue(evt.LightBlueFlash),
				evt.WithColor(rlGrad.Lerp(ctx.T())),
			)

			ctx.NewPreciseLaser(
				evt.WithDirectionalLaser(evt.LeftLaser),
				evt.WithIntValue(ctx.Ordinal()),
				evt.WithLockPosition(true),
				evt.WithPreciseLaserSpeed(float64(ctx.Ordinal())*2),
				evt.WithLaserDirection(chroma.Clockwise),
			)

			ctx.NewRGBLighting(
				evt.WithLight(evt.RightRotatingLasers),
				evt.WithLightValue(evt.LightOff),
			)

		case ctx.Ordinal()%2 == 1:
			ctx.NewRGBLighting(
				evt.WithLight(evt.LeftRotatingLasers),
				evt.WithLightValue(evt.LightOff),
			)

			ctx.NewRGBLighting(
				evt.WithLight(evt.RightRotatingLasers),
				evt.WithLightValue(evt.LightRedFlash),
				evt.WithColor(rlGrad.Lerp(ctx.T())),
			)

			ctx.NewPreciseLaser(
				evt.WithDirectionalLaser(evt.RightLaser),
				evt.WithIntValue(ctx.Ordinal()),
				evt.WithLockPosition(false),
				evt.WithPreciseLaserSpeed(float64(ctx.Ordinal())*2),
				evt.WithLaserDirection(chroma.CounterClockwise),
			)
		}
	})
}

func (p Intro) Fall(startBeat float64) {
	ctx := p.Offset(startBeat)

	seq := timer.NewSequencer([]float64{0, 0.25, 0.5, 0.75}, 0)
	p.PianoGlow(ctx, seq, 4, 0.20, false, false)

}

func (p Intro) Bridge(startBeat float64) {
	ctx := p.Offset(startBeat)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseRotation(
			evt.WithRotation(180),
			evt.WithRotationStep(12.5),
			evt.WithRotationDirection(chroma.CounterClockwise),
			evt.WithRotationSpeed(3),
			evt.WithProp(5),
			evt.WithCounterSpin(true),
		)
	})

	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Divide(2),
	)

	rng := timer.NewRanger(0, 1, 30, ease.OutCubic)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := fx.Gradient(ctx, shirayukiGradient)
			fx.AlphaFadeEx(ctx, e, 0, 1, 0, 1, ease.Linear)
		})
	})
}

func (p Intro) Outro(startBeat float64) {
	var (
		seq1 = timer.NewSequencer([]float64{0, 0.25, 0.50}, 0)
		seq2 = timer.NewSequencer([]float64{1.0, 1.25, 1.50}, 0)
		seq3 = timer.NewSequencer([]float64{2.0, 2.25, 2.50, 2.75, 3.25}, 0)
	)

	ctx := p.Offset(startBeat)

	p.PianoTransmute(ctx, seq1, 3, true, shirayukiSingleGradient)
	p.PianoTransmute(ctx, seq2, 3, true, sukoyaSingleGradient)
	p.PianoGlow(ctx, seq3, 5, 0.2, false, true)
}

func (p Intro) OutroSplash(startBeat float64) {
	ctx := p.Offset(startBeat)

	var (
		sweepSpeed = 1.5
		grad       = magnetRainbow
		backLaser  = transform.Light(light.NewBasic(ctx, evt.BackLasers),
			transform.DivideSingle(),
		)
	)

	rng := timer.NewRanger(0, 4, 60, ease.Linear)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(backLaser, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, sweepSpeed, grad)
			fx.AlphaFadeEx(ctx, e, 0, 0.5, 0, 2, ease.InCubic)
			fx.AlphaFadeEx(ctx, e, 0.5, 1, 2, 0, ease.OutCubic)
		})
	})

	seq := timer.NewSequencer([]float64{0, 0.75, 1.5}, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
		if ctx.Last() {
			ctx.NewPreciseRotation(
				evt.WithRotation(360),
				evt.WithRotationStep(0),
				evt.WithRotationDirection(chroma.Clockwise),
				evt.WithRotationSpeed(7),
				evt.WithProp(0.8),
				evt.WithCounterSpin(true),
			)
		} else {
			ctx.NewPreciseRotation(
				evt.WithRotation(45),
				evt.WithRotationStep(12.5),
				evt.WithRotationSpeed(26),
				evt.WithProp(8),
				evt.WithRotationDirection(chroma.CounterClockwise),
			)
		}
		p.Rush(ctx, -0.40, 0, 0.5, float64(ctx.SeqOrdinal())*0.5+1, magnetGradient)
	})
}
