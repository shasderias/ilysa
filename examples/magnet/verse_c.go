package main

import (
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

type VerseC struct {
	context.Context
}

func NewVerseC(p *ilysa.Project, offset float64) VerseC {
	return VerseC{p.BOffset(offset)}
}

func (v VerseC) Play() {
	v.Bridge(0)

	v.Lyrics(0, timer.Seq([]float64{
		0.00, 1.00, 1.75, 2.50, 3.00, 3.50, 4.00, 4.50, 4.75, 5.25, 5.75,
		7.50, 7.75, 8.00, 9.50, 10.00, 10.75, 11.50, 12.00, 12.50, 12.75,
	}, 13.00))

	v.ShirayukiBridge(16)

	v.Lyrics(16, timer.Seq([]float64{
		0.00, 1.00, 1.75, 2.25, 3.00, 3.50, 4.00, 4.50, 4.75, 5.25, 5.75, 6.50, 7.00, 7.50, 8.00,
		10.50, 11.00, 11.50, 12.00, 13.50, 13.75,
	}, 14.00))

	v.DunDun(30)

	v.Breathe(33)
}

func (v VerseC) Bridge(startBeat float64) {
	ctx := v.BOffset(startBeat)

	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.DivideSingle(),
	)

	ctx.Sequence(timer.Beat(-0.05), func(ctx context.Context) {
		fx.OffAll(ctx)
	})

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(0))
		ctx.NewPreciseRotation(
			evt.WithRotation(450),
			evt.WithRotationStep(12),
			evt.WithRotationSpeed(0.008),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.Clockwise),
		)

		ctx.Range(timer.Rng(0, 2, 30, ease.Linear), func(ctx context.Context) {
			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, sukoyaGradient)
				fx.AlphaFadeEx(ctx, e, 0, 1, 0, 0.40, ease.InCirc)
			})
		})
	})
}

func (v VerseC) Lyrics(startBeat float64, seq timer.Sequencer) {
	ctx := v.BOffset(startBeat)

	l := transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.DivideSingle(),
	)

	ctx.Sequence(seq, func(ctx context.Context) {
		grad := magnetRainbowPale.RotateRand()

		startAlpha := 3*ease.InCirc(ctx.SeqT()) + 0.2
		endAlpha := 1*ease.InCirc(ctx.SeqT()) + 0.2

		ctx.Range(timer.Rng(0, ctx.SeqNextBOffset(), 12, ease.Linear), func(ctx context.Context) {

			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				fx.RippleT(ctx, e, 0.8)
				fx.AlphaFadeEx(ctx, e, 0, 1, startAlpha, endAlpha, ease.OutCirc)
			})
		})
	})
}

func (v VerseC) ShirayukiBridge(startBeat float64) {
	ctx := v.BOffset(startBeat)

	l := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.DivideSingle(),
	)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(0))
		ctx.NewPreciseRotation(
			evt.WithRotation(450),
			evt.WithRotationStep(12),
			evt.WithRotationSpeed(0.05),
			evt.WithProp(1.2),
			evt.WithRotationDirection(chroma.CounterClockwise),
		)

		ctx.Range(timer.Rng(0, 2, 30, ease.Linear), func(ctx context.Context) {
			ctx.Light(l, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, shirayukiGradient)
				fx.AlphaFadeEx(ctx, e, 0, 1, 0, 0.40, ease.InCirc)
			})
		})
	})
}

func (v VerseC) DunDun(startBeat float64) {
	ctx := v.BOffset(startBeat)

	ll := transform.Light(light.NewBasic(ctx, evt.LeftRotatingLasers),
		transform.DivideSingle())
	rl := transform.Light(light.NewBasic(ctx, evt.RightRotatingLasers),
		transform.DivideSingle())
	bl := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.DivideSingle())

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(-1))

		ctx.Light(ll, func(ctx context.LightContext) {
			fx.Gradient(ctx, sukoyaSingleGradient)
		})

		ctx.Light(rl, func(ctx context.LightContext) {
			fx.Gradient(ctx, shirayukiSingleGradient)
		})

		ctx.Light(bl, func(ctx context.LightContext) {
			fx.Gradient(ctx, magnetRainbow)
		})
	})

	ctx.Sequence(timer.Beat(0.5), func(ctx context.Context) {
		ctx.NewPreciseZoom(evt.WithZoomStep(0))
		fx.OffAll(ctx)
	})
	ctx.Sequence(timer.Beat(0.51), func(ctx context.Context) {
		centerOn(ctx, magnetWhite)
	})
}

func (v VerseC) Breathe(startBeat float64) {
	ctx := v.BOffset(startBeat)

	lights := []evt.LightType{
		//evt.BackLasers,
		evt.RingLights,
		evt.LeftRotatingLasers,
		evt.RightRotatingLasers,
		//evt.CenterLights,
	}

	cl := light.Combine()

	for _, l := range lights {
		cl.Add(transform.Light(light.NewBasic(ctx, l),
			transform.DivideSingle(),
		))
	}

	blackGrad := gradient.New(colorful.Black, colorful.MustParseHex("#FFFFFF"))

	ringRng := timer.Rng(-0.75, -0.25, 15, ease.Linear)
	RingRipple(ctx, ringRng, blackGrad,
		WithSweepSpeed(1.2),
		WithRippleTime(1.0),
		WithFadeIn(fx.NewAlphaFader(0, 1, 2, 0, ease.InCirc)),
		WithReverse(true),
	)

	rng := timer.Rng(0, 0.75, 15, ease.Linear)
	ctx.Range(rng, func(ctx context.Context) {
		ctx.Light(cl, func(ctx context.LightContext) {
			e := fx.ColorSweep(ctx, 1.2, blackGrad)
			fx.AlphaFadeEx(ctx, e, 0.0, 0.5, 0, 2, ease.InCirc)
			fx.AlphaFadeEx(ctx, e, 0.5, 1.0, 2, 0, ease.OutCirc)
		})
	})
}
