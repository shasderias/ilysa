package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
)

const mapPath = `D:\Beat Saber Data\CustomWIPLevels\MagnetLights`

func main() {
	if err := do(); err != nil {
		fmt.Println(err)
	}
}

var (
	shirayukiGold   = colorful.MustParseHex("#CFB96A")
	shirayukiPurple = colorful.MustParseHex("#6F37E3")
	sukoyaPink      = colorful.MustParseHex("#E92EB1")
	sukoyaWhite     = colorful.MustParseHex("#FFFFFF")
	magnetPurple    = colorful.MustParseHex("#7F7BEB")
	magnetPink      = colorful.MustParseHex("#B63A8C")
	magnetWhite     = colorful.MustParseHex("#FFFFFF")
)

var (
	allColors = colorful.NewSet(
		shirayukiGold,
		shirayukiPurple,
		sukoyaWhite,
		sukoyaPink,
		magnetPurple,
		magnetPink,
		magnetWhite,
	)

	magnetColors = colorful.NewSet(
		magnetPink,
		magnetPurple,
		magnetWhite,
	)

	shirayukiColors = colorful.NewSet(
		shirayukiGold,
		shirayukiPurple,
	)

	sukoyaColors = colorful.NewSet(
		sukoyaPink,
		sukoyaWhite,
	)

	crossickColors = colorful.NewSet(
		shirayukiGold,
		shirayukiPurple,
		sukoyaWhite,
		sukoyaPink,
	)
)

var (
	magnetGradient = gradient.Table{
		{magnetPurple, 0.0},
		{magnetPink, 0.25},
		{magnetWhite, 0.50},
		{magnetPurple, 1.00},
	}

	allColorsGradient = gradient.Table{
		{shirayukiGold, 0.0},
		{shirayukiPurple, 0.167},
		{sukoyaPink, 0.167 * 2},
		{sukoyaWhite, 0.167 * 3},
		{magnetPurple, 0.167 * 4},
		{magnetPink, 0.167 * 5},
		{magnetWhite, 1.0},
	}
)

func do() error {
	magnet, err := beatsaber.Open(mapPath)
	if err != nil {
		return err
	}

	p := ilysa.New(magnet)

	err = p.Map.SetActiveDifficulty(beatsaber.CharacteristicStandard, beatsaber.BeatmapDifficultyExpertPlus)
	if err != nil {
		return err
	}

	//LeadIn(p)
	//
	//Intro{
	//	Project:   p,
	//	startBeat: 0,
	//}.Play()
	//
	//verse1 := NewVerse1a(p, 52)
	//verse1.Play()
	//
	//verse2 := NewVerse1b(p, 84)
	//verse2.Play()

	chorus := NewChorus(p, 114)
	chorus.Play()

	//v := Verse{Project: p}
	//v.Play(52)

	//
	//BassTwang(p, 18.5)
	//
	//

	return p.Save()
}

//func Shimmer(p *ilysa.Project, startBeat, endBeat float64, steps int, light beatsaber.EventTypeSet, colorSweepSpeed, shimmerSweepSpeed float64) {
//	var (
//		duration   = endBeat - startBeat
//		offset     = rand.Float64() * math.Pi * 2
//		LightIDMax = p.ActiveDifficultyProfile().LightIDMax(light)
//		sin        = math.Sin
//		pi         = math.Pi
//	)
//
//	colorSweepSpeed *= duration
//	shimmerSweepSpeed *= duration
//
//	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Timer) {
//		for i := 1; i <= LightIDMax; i++ {
//			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
//			e.SetSingleLightID(i)
//			e.SetColor(gradient.Rainbow.Ierp(
//				sin(ctx.t*colorSweepSpeed + (float64(i)/float64(LightIDMax))*pi + offset),
//			))
//			e.SetAlpha(5)
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
//	//fadeScale := scale.Clamped(startBeat, endBeat, 0, 1)
//	//
//	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.2), 0, 1, ease.InCubic)
//	//modfx.RGBAlphaFade(p, light, fadeScale(0.9), endBeat, 1, 0, ease.OutCubic)
//	fadeScale := scale.Clamped(startBeat, endBeat, 0, 1)
//	modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
//	modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutCubic)
//}
//
//
//func SimpleShimmer(p *ilysa.Project, startBeat, endBeat float64) {
//	const steps = 60
//
//	var (
//		sin        = math.Sin
//		light      = beatsaber.EventTypeBackLasers
//		LightIDMax = p.ActiveDifficultyProfile().LightIDMax(light)
//		pi         = math.Pi
//	)
//
//	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx ilysa.Timer) {
//		for i := 1; i <= LightIDMax; i++ {
//			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightBlueOn)
//			e.SetSingleLightID(i)
//			e.SetColor(magnetGradient.Ierp(
//				sin(ctx.t*3 + (float64(i)/float64(LightIDMax))*pi + 4),
//			))
//			e.SetAlpha(5)
//		}
//	})
//
//	fadeScale := scale.Clamped(startBeat, endBeat, 0, 1)
//	modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
//	modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutBounce)
//}