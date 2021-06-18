package modfx

import (
	"github.com/shasderias/ilysa/pkg/beatsaber"
	"github.com/shasderias/ilysa/pkg/ease"
	"github.com/shasderias/ilysa/pkg/ilysa"
	"github.com/shasderias/ilysa/pkg/util"
)

func RGBAlphaFade(p *ilysa.Project, target beatsaber.EventType,
	startBeat, endBeat, startAlpha, endAlpha float64, fadeEase ease.Func) {

	alphaScale := util.Scale(0, 1, startAlpha, endAlpha)

	p.ModEventsInRange(startBeat, endBeat,
		ilysa.FilterLightingEvents(target),
		func(ctx ilysa.Timing, event ilysa.Event) {
			e := event.(*ilysa.RGBLightingEvent)

			alphaMut := alphaScale(fadeEase(ctx.Pos))


			e.SetAlpha(e.GetAlpha() * alphaMut)
		})
}
