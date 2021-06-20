package fx

import (
	"math"

	ilysa2 "github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/noise"
	"github.com/shasderias/ilysa/colorful/gradient"
)

var (
	pi  = math.Pi
	abs = math.Abs
	sin = math.Sin
)

func SinSweepLightID(sweepSpeed, offset float64) func(ctx ilysa2.TimeLightContext) float64 {
	return func(ctx ilysa2.TimeLightContext) float64 {
		return sin(ctx.T()*ctx.Duration()*sweepSpeed + ctx.LightIDT()*pi + offset)
	}
}

func AbsSinSweepLightID(sweepSpeed, offset float64) func(ctx ilysa2.TimeLightContext) float64 {
	return func(ctx ilysa2.TimeLightContext) float64 {
		return abs(sin(ctx.T()*ctx.Duration()*sweepSpeed + ctx.LightIDT()*pi + offset))
	}
}

func BiasedColorSweep(ctx ilysa2.TimeLightContext, sweepSpeed float64, grad gradient.Table) *ilysa2.CompoundRGBLightingEvent {
	gradPos := SinSweepLightID(sweepSpeed, ctx.FixedRand())
	return ctx.NewRGBLightingEvent(
		ilysa2.WithColor(grad.Ierp(gradPos(ctx))),
	)
}

func ColorSweep(ctx ilysa2.TimeLightContext, sweepSpeed float64, grad gradient.Table, opts...ColorSweepOpt) *ilysa2.CompoundRGBLightingEvent {
	gradPos := AbsSinSweepLightID(sweepSpeed, ctx.FixedRand())

	e :=  ctx.NewRGBLightingEvent(
		ilysa2.WithColor(grad.Ierp(gradPos(ctx))),
	)

	for _, opt := range opts {
		opt.applyColorSweep(ctx, e)
	}

	return e
}

type ColorSweepOpt interface {
	applyColorSweep(light ilysa2.TimeLightContext, event *ilysa2.CompoundRGBLightingEvent)
}

func AlphaShimmer(ctx ilysa2.TimeLightContext, e ilysa2.EventWithAlpha, shimmerSpeed float64) {
	e.SetAlpha(e.GetAlpha() * noise.DefaultWave(shimmerSpeed*ctx.T()+ctx.LightIDT()))
}

//func SweepLightID(light beatsaber.EventTypeSet, lightIDPicker lightid.Picker) func(ctx ilysa.Timer) {
//	return func(ctx ilysa.Timer) {
//		lidSet := lightIDPicker(ctx, light)
//
//		for i := 1; i < LightIDMax; i++ {
//			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
//			e.SetSingleLightID(i)
//			e.SetColor(gradient.Rainbow.Ierp(
//				sin(ctx.t*sweepSpeed + (float64(i)/float64(LightIDMax))*pi + offset),
//			))
//		}
//	}
//}
//
//func Shimmer(p *ilysa.Project, startBeat, endBeat float64, steps int, light beatsaber.EventTypeSet, intensity, colorSweepSpeed float64) {
//	var (
//		sin        = math.Sin
//		pi         = math.Pi
//		duration   = endBeat - startBeat
//		offset     = rand.Float64() * pi
//		LightIDMax = p.ActiveDifficultyProfile().LightIDMax(light)
//	)
//
//	colorSweepSpeed *= duration
//
//	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Timer) {
//		for i := 1; i <= LightIDMax; i++ {
//			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
//			e.SetSingleLightID(i)
//			e.SetColor(gradient.Rainbow.Ierp(
//				sin(ctx.t*colorSweepSpeed + (float64(i)/float64(LightIDMax))*pi + offset),
//			))
//			e.SetAlpha(intensity)
//		}
//	})
//
//	//p.ModEventsInRange(startBeat, endBeat,
//	//	ilysa.FilterLightingEvents(light),
//	//	func(ctx ilysa.Timer, event ilysa.Event) {
//	//		e := event.(*ilysa.RGBLightingEvent)
//	//		lightID := float64(e.FirstLightID())
//	//		e.SetAlpha(e.GetAlpha() * util.DefaultNoise(ctx.t*shimmerSweepSpeed+lightID/float64(LightIDMax)*pi+offset))
//	//	})
//
//	//fadeScale := util.Clamped(startBeat, endBeat, 0, 1)
//	//
//	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.2), 0, 1, ease.InCubic)
//	//modfx.RGBAlphaFade(p, light, fadeScale(0.9), endBeat, 1, 0, ease.OutCubic)
//	//fadeScale := util.Clamped(startBeat, endBeat, 0, 1)
//	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
//	//modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutCubic)
//}
