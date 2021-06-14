package fx

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ilysa"
)

func Gradient(ctx ilysa.RangeLightIDContext, val beatsaber.EventValue, table gradient.Table) *ilysa.RGBLightingEvent {
	return ctx.NewRGBLightingEvent().
		SetValue(val).
		SetLightID(ctx.CurLightID).
		SetColor(table.GetInterpolatedColorFor(ctx.LightIDPos))
}
