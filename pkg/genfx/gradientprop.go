package genfx

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/util"
)

func GradientProp(ctx ilysa.Context, typ beatsaber.EventType, val beatsaber.EventValue, table gradient.Table, minLightID, maxLightID int) {
	scale := util.Scale(float64(minLightID), float64(maxLightID), 0, 1)
	for i := minLightID; i < maxLightID; i++ {
		e := ctx.NewRGBLightingEvent(typ, val)
		e.SetSingleLightID(i)
		e.SetColor(table.GetInterpolatedColorFor(scale(float64(i))))
	}
}
