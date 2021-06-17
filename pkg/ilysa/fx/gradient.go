package fx

import (
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ilysa"
)

func Gradient(ctx ilysa.TimingContextWithLight, table gradient.Table) *ilysa.CompoundRGBLightingEvent {
	return ctx.NewRGBLightingEvent(
		ilysa.WithColor(table.GetInterpolatedColorFor(ctx.LightIDT())),
	)
}
