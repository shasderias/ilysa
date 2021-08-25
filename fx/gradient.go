package fx

import (
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

func Gradient(ctx context.LightContext, table gradient.Table) evt.RGBLightingEvents {
	return ctx.NewRGBLighting(
		evt.WithColor(table.Lerp(ctx.LightIDT())),
	)
}

func GradientT(ctx context.LightContext, table gradient.Table) evt.RGBLightingEvents {
	return ctx.NewRGBLighting(
		evt.WithColor(table.Lerp(ctx.T())),
	)
}
