package fx

import (
	"github.com/shasderias/ilysa/pkg/colorful/gradient"
	"github.com/shasderias/ilysa/pkg/ilysa"
)

func Gradient(ctx ilysa.TimingContextWithLight, table gradient.Table) *ilysa.CompoundRGBLightingEvent {
	return ctx.NewRGBLightingEvent(
		ilysa.WithColor(table.GetInterpolatedColorFor(ctx.LightIDT())),
	)
}
