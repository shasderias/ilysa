package main

import (
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
)

func centerOn(ctx context.Context, color colorful.Color) {
	centerLights := light.NewBasic(ctx, evt.CenterLights)

	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.Light(centerLights, func(ctx context.LightContext) {
			ctx.NewRGBLighting(
				evt.WithColor(color),
				evt.WithAlpha(1.25),
			)
		})
	})
}
