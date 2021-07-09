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

type Verse struct {
	context.Context
	project *ilysa.Project
	offset  float64
}

func NewVerse1a(p *ilysa.Project, offset float64) Verse {
	ctx := p.Offset(offset)
	return Verse{
		Context: ctx,
		project: p,
		offset:  offset,
	}
}

func (p Verse) Play1() {
	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		fx.OffAll(ctx)
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.LeftLaser), evt.WithPreciseLaserSpeed(1.5))
		ctx.NewPreciseLaser(evt.WithDirectionalLaser(evt.RightLaser), evt.WithPreciseLaserSpeed(1.5))
	})

	p.Rhythm(0, true)
	p.Rhythm(4, false)
	p.Rhythm(8, false)
	p.Rhythm(12, false)
	p.Rhythm(16, false)
	p.Rhythm(20, false)
	p.Rhythm(24, false)
	p.Rhythm(28, true)

	p.PianoBackstep(7)
	p.PianoBackstep(15)
	p.PianoBackstep(23)
	p.RinPun(27)
}

func (p Verse) Play2() {
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
	p.Rhythm(16, false)
	p.Rhythm(20, false)
	p.Rhythm(24, false)
	p.Rhythm(28, true)

	p.PianoBackstep(7)
	//p.PianoBackstep(15)
	p.SnareDrums(14, []float64{0.25, 0.50, 1.00, 1.25, 1.75, 2.00})
	p.PianoBackstep(23)
	p.RinPun(27)
}

func (p Verse) Rhythm(startBeat float64, kickOnly bool) {
	ctx := p.Offset(startBeat)
	var (
		kickDrumLight = light.NewSequence(
			light.NewBasic(ctx, evt.LeftRotatingLasers),
			light.NewBasic(ctx, evt.RightRotatingLasers),
		)
		kickDrumSequence = timer.NewSequencer([]float64{0, 2.5}, 0)
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
		rng  = timer.NewRanger(0, 2, 30, ease.Linear)
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

func (p Verse) PianoBackstep(startBeat float64) {
	ctx := p.Offset(startBeat)

	var (
		sequence   = timer.NewSequencer([]float64{0, 0.5}, 0)
		backLasers = transform.Light(
			light.NewBasic(ctx, evt.BackLasers),
			transform.Divide(2).Sequence(),
			transform.DivideSingle(),
		)
	)

	ctx.Sequence(sequence, func(ctx context.Context) {
		rotOpt := evt.NewOpts(
			evt.WithRotation(15),
			evt.WithRotationStep(15),
			evt.WithProp(12),
			evt.WithRotationSpeed(10),
		)

		if ctx.Ordinal()%2 == 0 {
			ctx.NewPreciseRotation(evt.WithRotationDirection(chroma.Clockwise), rotOpt)
		} else {
			ctx.NewPreciseRotation(evt.WithRotationDirection(chroma.CounterClockwise), rotOpt)
		}

		rng := timer.NewRanger(-0.1, 4.5, 6, ease.Linear)

		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(backLasers, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, magnetGradient)
				fx.AlphaFadeEx(ctx, e, 0, 1, 0.9, 0, ease.OutCirc)
			})
		})
	})
}

func (p Verse) SnareDrums(startBeat float64, sequence []float64) {
	ctx := p.Offset(startBeat)
	l := transform.Light(
		light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2).Sequence(),
		transform.DivideSingle(),
	)

	gradSet := gradient.NewSet(magnetGradient, shirayukiGradient)

	seq := timer.NewSequencer(sequence, 0)
	p.Sequence(seq, func(ctx context.Context) {
		grad := gradSet.Next()

		if ctx.Ordinal()%2 == 0 {
			ctx.NewPreciseRotation(
				evt.WithRotation(15),
				evt.WithRotationStep(15),
				evt.WithProp(20),
				evt.WithRotationSpeed(8),
				evt.WithRotationDirection(chroma.CounterClockwise),
			)
			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				e.Apply(evt.WithAlpha(2))
				oe := ctx.NewRGBLighting(evt.WithLightValue(evt.LightOff))
				oe.Apply(evt.WithBeatOffset(0.15))
			})
		} else {
			ctx.NewPreciseRotation(
				evt.WithRotation(30),
				evt.WithRotationStep(15),
				evt.WithProp(1.2),
				evt.WithRotationSpeed(8),
				evt.WithRotationDirection(chroma.CounterClockwise),
			)
			seq := timer.Interval(0, 0.5, 8)
			ctx.Sequence(seq, func(ctx context.Context) {
				ctx.Light(l, func(ctx context.LightContext) {
					e := fx.Gradient(ctx, grad)
					fx.AlphaFadeEx(ctx, e, 0, 1, 2, 0, ease.InCubic)
				})
			})
		}
	})
}

func (p Verse) RinPun(startBeat float64) {
	ctx := p.Offset(startBeat)
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

	ctx.NewZoom()
	ctx.NewPreciseRotation(
		evt.WithRotation(45),
		evt.WithRotationStep(12),
		evt.WithRotationSpeed(2),
		evt.WithProp(1.2),
		evt.WithRotationDirection(chroma.CounterClockwise),
	)

	seq := timer.NewSequencer([]float64{0, 0.5}, 0)
	ctx.Sequence(seq, func(ctx context.Context) {
		rng := timer.NewRanger(-0.3, 1.2, 30, ease.Linear)
		ctx.Range(rng, func(ctx context.Context) {
			ctx.Light(sl, func(ctx context.LightContext) {
				e := fx.ColorSweep(ctx, 0.4, magnetRainbow)
				fx.RippleT(ctx, e, 0.30)
				fx.AlphaFadeEx(ctx, e, 0, 0.5, 0, 1, ease.InCubic)
				fx.AlphaFadeEx(ctx, e, 0.5, 1.0, 1, 0.7, ease.OutCirc)
			})
		})
	})

	ctx.Sequence(timer.NewSequencer([]float64{1, 2}, 0), func(ctx context.Context) {
		ctx.NewZoom()
		ctx.NewPreciseRotation(
			evt.WithRotation(45),
			evt.WithRotationStep(12),
			evt.WithRotationSpeed(2),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)

		//seqOrdinal := ctx.Ordinal()
		ctx.Range(timer.NewRanger(-0.2, ctx.B()+0.7, 30, ease.InOutCirc), func(ctx context.Context) {
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
			transform.Divide(3).Sequence(),
			transform.DivideSingle(),
		).(light.Sequence)

		sl = light.NewSequence(
			sl.Idx(1),
			sl.Idx(0),
			sl.Idx(2),
		)

		p.Sequence(timer.NewSequencer([]float64{3, 3.5, 3.75}, 0), func(ctx context.Context) {
			ctx.NewZoom()
			ctx.NewPreciseRotation(
				evt.WithRotation(22.5),
				evt.WithRotationStep(12),
				evt.WithRotationSpeed(0.5),
				evt.WithProp(2),
				evt.WithRotationDirection(chroma.CounterClockwise),
			)
			ctx.Range(timer.NewRanger(0, 0.6, 30, ease.InOutCirc), func(ctx context.Context) {
				ctx.Light(sl, func(ctx context.LightContext) {
					e := fx.ColorSweep(ctx, 0.8, magnetRainbow)
					fx.Ripple(ctx, e, 0.65)
					fx.AlphaFadeEx(ctx, e, 0, 0.5, 0, 2, ease.InCubic)
					fx.AlphaFadeEx(ctx, e, 0.5, 1.0, 2, 0, ease.OutCirc)
				})
			})
		})
	}
}
