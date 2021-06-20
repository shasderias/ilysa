package fx

import (
	ilysa2 "github.com/shasderias/ilysa"
	ease2 "github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

func RGBAlphaBlend(ctx ilysa2.TimeContext, event ilysa2.Event, startAlpha, endAlpha float64, easeFn ease2.Func) {
	alphaScale := scale.FromUnitIntervalClamped(startAlpha, endAlpha)

	e, ok := event.(ilysa2.EventWithAlpha)
	if !ok {
		return
	}
	e.SetAlpha(e.GetAlpha() * alphaScale(easeFn(ctx.T())))
}
