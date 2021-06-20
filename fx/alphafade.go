package fx

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

func RGBAlphaBlend(ctx ilysa.TimeContext, event ilysa.Event, startAlpha, endAlpha float64, easeFn ease.Func) {
	alphaScale := scale.FromUnitIntervalClamped(startAlpha, endAlpha)

	e, ok := event.(ilysa.EventWithAlpha)
	if !ok {
		return
	}
	e.SetAlpha(e.GetAlpha() * alphaScale(easeFn(ctx.T())))
}
