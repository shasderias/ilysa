package fx

import (
	"math"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/util"
)

var (
	pi  = math.Pi
	abs = math.Abs
	sin = math.Sin
)

func SinSweepLightID(sweepSpeed, offset float64) func(ctx ilysa.RangeLightIDContext) float64 {
	return func(ctx ilysa.RangeLightIDContext) float64 {
		return sin(ctx.Pos*ctx.Duration*sweepSpeed + float64(ctx.CurLightID[0])/float64(ctx.MaxLightID)*pi + offset)
	}
}

func AbsSinSweepLightID(sweepSpeed, offset float64) func(ctx ilysa.RangeLightIDContext) float64 {
	return func(ctx ilysa.RangeLightIDContext) float64 {
		return abs(sin(ctx.Pos*ctx.Duration*sweepSpeed + float64(ctx.CurLightID[0])/float64(ctx.MaxLightID)*pi + offset))
	}
}

func BiasedColorSweep(ctx ilysa.RangeLightIDContext, intensity, sweepSpeed float64, grad gradient.Table) *ilysa.RGBLightingEvent {
	gradPos := SinSweepLightID(sweepSpeed, ctx.RandFloat64)
	return ctx.NewRGBLightingEvent().
		SetValue(beatsaber.EventValueLightRedOn).
		SetColor(grad.GetInterpolatedColorFor(gradPos(ctx))).
		SetAlpha(intensity)
}

func ColorSweep(ctx ilysa.RangeLightIDContext, intensity, sweepSpeed float64, grad gradient.Table) *ilysa.RGBLightingEvent {
	gradPos := AbsSinSweepLightID(sweepSpeed, ctx.RandFloat64)
	return ctx.NewRGBLightingEvent().
		SetValue(beatsaber.EventValueLightRedOn).
		SetColor(grad.GetInterpolatedColorFor(gradPos(ctx))).
		SetAlpha(intensity)
}

func AlphaShimmer(ctx ilysa.RangeLightIDContext, e *ilysa.RGBLightingEvent, shimmerSpeed float64) {
	e.SetAlpha(e.GetAlpha() * util.DefaultWave(shimmerSpeed*ctx.Pos+ctx.LightIDPos))
}

//func SweepLightID(light beatsaber.EventType, lightIDPicker lightid.Picker) func(ctx ilysa.Context) {
//	return func(ctx ilysa.Context) {
//		lidSet := lightIDPicker(ctx, light)
//
//		for i := 1; i < maxLightID; i++ {
//			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
//			e.SetSingleLightID(i)
//			e.SetColor(gradient.Rainbow.GetInterpolatedColorFor(
//				sin(ctx.Pos*sweepSpeed + (float64(i)/float64(maxLightID))*pi + offset),
//			))
//		}
//	}
//}
//
//func Shimmer(p *ilysa.Project, startBeat, endBeat float64, steps int, light beatsaber.EventType, intensity, colorSweepSpeed float64) {
//	var (
//		sin        = math.Sin
//		pi         = math.Pi
//		duration   = endBeat - startBeat
//		offset     = rand.Float64() * pi
//		maxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
//	)
//
//	colorSweepSpeed *= duration
//
//	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Context) {
//		for i := 1; i <= maxLightID; i++ {
//			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
//			e.SetSingleLightID(i)
//			e.SetColor(gradient.Rainbow.GetInterpolatedColorFor(
//				sin(ctx.Pos*colorSweepSpeed + (float64(i)/float64(maxLightID))*pi + offset),
//			))
//			e.SetAlpha(intensity)
//		}
//	})
//
//	//p.ModEventsInRange(startBeat, endBeat,
//	//	ilysa.FilterLightingEvents(light),
//	//	func(ctx ilysa.Context, event ilysa.Event) {
//	//		e := event.(*ilysa.RGBLightingEvent)
//	//		lightID := float64(e.FirstLightID())
//	//		e.SetAlpha(e.GetAlpha() * util.DefaultNoise(ctx.Pos*shimmerSweepSpeed+lightID/float64(maxLightID)*pi+offset))
//	//	})
//
//	//fadeScale := util.Scale(startBeat, endBeat, 0, 1)
//	//
//	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.2), 0, 1, ease.InCubic)
//	//modfx.RGBAlphaFade(p, light, fadeScale(0.9), endBeat, 1, 0, ease.OutCubic)
//	//fadeScale := util.Scale(startBeat, endBeat, 0, 1)
//	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
//	//modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutCubic)
//}
