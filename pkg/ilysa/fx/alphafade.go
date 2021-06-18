package fx

import (
	"ilysa/pkg/ease"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/util"
)

//func RGBAlphaFade(p *ilysa.Project, light ilysa.Light,
//	startBeat, endBeat, startAlpha, endAlpha float64,
//	easeFn ease.Func) {
//
//	alphaScale := util.ScaleFromUnitInterval(startAlpha, endAlpha)
//
//	p.ModEventsInRange(startBeat, endBeat, ilysa.FilterRGBLight(light),
//		func(ctx ilysa.Timing, event ilysa.Event) {
//			e := event.(*ilysa.RGBLightingEvent)
//			e.SetAlpha(e.GetAlpha() * alphaScale(easeFn(ctx.t)))
//		})
//}

func RGBAlphaBlend(ctx ilysa.TimingContext, event ilysa.Event, startAlpha, endAlpha float64, easeFn ease.Func) {
	alphaScale := util.ScaleFromUnitInterval(startAlpha, endAlpha)

	e, ok := event.(ilysa.EventWithAlpha)
	if !ok { return }
	e.SetAlpha(e.GetAlpha() * alphaScale(easeFn(ctx.T())))
}
