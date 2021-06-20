package fx

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/colorful/gradient"
)

func Gradient(ctx ilysa.TimeLightContext, table gradient.Table) *ilysa.CompoundRGBLightingEvent {
	return ctx.NewRGBLightingEvent(
		ilysa.WithColor(table.Ierp(ctx.LightIDT())),
	)
}
