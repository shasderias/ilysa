package fx

import (
	"math"

	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/scale"
)

var (
	pi  = math.Pi
	abs = math.Abs
	sin = math.Sin
)

func absSinSweepLightID(sweepSpeed, offset float64) func(ctx context.LightContext) float64 {
	s := scale.ToUnitClamp(-1, 1)
	return func(ctx context.LightContext) float64 {
		a := s(sin(ctx.T()*sweepSpeed*ctx.Duration() + ctx.LightT()*pi + offset*sweepSpeed*pi))
		return a
	}
}

func ColorSweep(ctx context.LightContext, sweepSpeed float64, grad gradient.Table) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		gradPos := absSinSweepLightID(sweepSpeed, ctx.FixedRand())
		e.Apply(evt.OptColor(grad.Lerp(gradPos(ctx))))
	})
}

//func SinSweepLightID(sweepSpeed, offset float64) func(ctx context.LightContext) float64 {
//	return func(ctx context.LightContext) float64 {
//		a := sin(ctx.T()*sweepSpeed/ctx.Duration() + ctx.LightT()*pi + offset)
//		return a
//	}
//}
//
//
//func BiasedColorSweep(ctx context.LightContext, sweepSpeed float64, grad gradient.Table) evt.RGBLightingEvents {
//	gradPos := SinSweepLightID(sweepSpeed, ctx.FixedRand())
//	e := ctx.NewRGBLighting(evt.WithColor(grad.Lerp(gradPos(ctx))))
//	return e
//}
//

//
//func AlphaShimmer(ctx context.LightContext, events evt.RGBLightingEvents, shimmerSpeed float64) {
//	for _, e := range events {
//		e.Apply(opt.Alpha(e.Alpha() * noise.DefaultWave(shimmerSpeed*ctx.T()+ctx.LightT())))
//	}
//}
//
////func SweepLightID(light beatsaber.EventTypeSet, lightIDPicker lightid.Picker) func(ctx ilysa.rng) {
////	return func(ctx ilysa.rng) {
////		lidSet := lightIDPicker(ctx, light)
////
////		for i := 1; i < MaxLightID; i++ {
////			e := ctx.NewChromaLighting(light, beatsaber.EventValueLightRedOn)
////			e.SetSingleLightID(i)
////			e.SetColor(gradient.Rainbow.Lerp(
////				sin(ctx.t*sweepSpeed + (float64(i)/float64(MaxLightID))*pi + offset),
////			))
////		}
////	}
////}
////
////func Shimmer(p *ilysa.Project, startBeat, endBeat float64, steps int, light beatsaber.EventTypeSet, intensity, colorSweepSpeed float64) {
////	var (
////		sin        = math.Sin
////		pi         = math.Pi
////		duration   = endBeat - startBeat
////		offset     = rand.Float64() * pi
////		MaxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
////	)
////
////	colorSweepSpeed *= duration
////
////	p.rangeTimer(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.rng) {
////		for i := 1; i <= MaxLightID; i++ {
////			e := ctx.NewChromaLighting(light, beatsaber.EventValueLightRedOn)
////			e.SetSingleLightID(i)
////			e.SetColor(gradient.Rainbow.Lerp(
////				sin(ctx.t*colorSweepSpeed + (float64(i)/float64(MaxLightID))*pi + offset),
////			))
////			e.SetAlpha(intensity)
////		}
////	})
////
////	//p.ModEventsInRange(startBeat, endBeat,
////	//	ilysa.FilterLightingEvents(light),
////	//	func(ctx ilysa.rng, event ilysa.Event) {
////	//		e := event.(*ilysa.RGBLightingEvent)
////	//		lightID := float64(e.FirstLightID())
////	//		e.SetAlpha(e.Alpha() * util.DefaultNoise(ctx.t*shimmerSweepSpeed+lightID/float64(MaxLightID)*pi+offset))
////	//	})
////
////	//fadeScale := util.Clamp(startBeat, endBeat, 0, 1)
////	//
////	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.2), 0, 1, ease.InCubic)
////	//modfx.RGBAlphaFade(p, light, fadeScale(0.9), endBeat, 1, 0, ease.OutCubic)
////	//fadeScale := util.Clamp(startBeat, endBeat, 0, 1)
////	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
////	//modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutCubic)
////}
