package fx

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light2"
)

func RainbowProp(p ilysa.BaseContext, light light2.Light, grad gradient.Table, startBeat, duration, step float64, frames int) {
	p.Range(startBeat, startBeat+duration, frames, ease.Linear, func(ctx ilysa.RangeContext) {
		ctx.WithLight(light, func(ctx ilysa.TimeLightContext) {
			e := ctx.NewRGBLightingEvent(
				evt.WithColor(grad.Ierp(ctx.T())),
			)
			Ripple(ctx, e, step)
			AlphaBlend(ctx, e, 0.0, 0.4, 0, 1, ease.InCirc)
			AlphaBlend(ctx, e, 0.4, 1, 1, 0, ease.OutCirc)
		})
	})
}
