package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
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

const BPMChangesJSON = `[]`

func main() {
	var (
		env   = beatsaber.EnvironmentPanic
		bpm   = 130.0
		laser = evt.RingLights
		grad  = gradient.New(
			colorful.MustParseHex("#12c2e9"),
			colorful.MustParseHex("#c471ed"),
			colorful.MustParseHex("#f64f59"),
		)
	)

	var (
		startBeat = 116.0
		endBeat   = 117.0
		steps     = 30
		easing    = ease.Linear
	)

	var (
		rippleTime = 1.0
	)

	var (
		fadeInStartT = 0.0
		fadeInEndT   = 0.0
		fadeInStartA = 1.0
		fadeInEndA   = 1.0
		fadeInEase   = ease.InCirc
	)

	var (
		fadeOutStartT = 0.0
		fadeOutEndT   = 1.0
		fadeOutStartA = 3.0
		fadeOutEndA   = 0.0
		fadeOutEase   = ease.OutCirc
	)

	m, err := beatsaber.NewMockMap(env, bpm, BPMChangesJSON)
	if err != nil {
		panic(err)
	}

	p := ilysa.New(m)

	rng := timer.Rng(startBeat, endBeat, steps, easing)

	l := transform.Light(light.NewBasic(p, laser),
		transform.DivideSingle())

	p.Range(rng, func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			e := ctx.NewRGBLighting(
				evt.WithColor(grad.Lerp(ctx.T())),
			)
			fx.RippleT(ctx, e, rippleTime)
			fx.AlphaFadeEx(ctx, e, fadeInStartT, fadeInEndT, fadeInStartA, fadeInEndA, fadeInEase)
			fx.AlphaFadeEx(ctx, e, fadeOutStartT, fadeOutEndT, fadeOutStartA, fadeOutEndA, fadeOutEase)
		})
	})

	p.Dump()
}
