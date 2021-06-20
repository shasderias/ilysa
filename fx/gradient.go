package fx

import (
	ilysa2 "github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/colorful/gradient"
)

func Gradient(ctx ilysa2.TimeLightContext, table gradient.Table) *ilysa2.CompoundRGBLightingEvent {
	return ctx.NewRGBLightingEvent(
		ilysa2.WithColor(table.Ierp(ctx.LightIDT())),
	)
}
