package fx

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/evt"
)

func Gradient(ctx ilysa.TimeLightContext, table gradient.Table) *ilysa.CompoundRGBLightingEvent {
	return ctx.NewRGBLightingEvent(
		evt.WithColor(table.Ierp(ctx.LightIDT())),
	)
}
