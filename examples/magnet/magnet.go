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
	magnetRainbowPale = gradient.New(
		colorful.MustParseHex("#F48DB4"),
		colorful.MustParseHex("#BCA2D8"),
		colorful.MustParseHex("#70B5D8"),
		colorful.MustParseHex("#44BFB4"),
		colorful.MustParseHex("#6DBE81"),
		colorful.MustParseHex("#A5B559"),
		colorful.MustParseHex("#D6A454"),
		colorful.MustParseHex("#F49472"),
	)

	magnetRainbow = gradient.New(
		colorful.MustParseHex("#FF0000"),
		colorful.MustParseHex("#FF8000"),
		colorful.MustParseHex("#FFFF00"),
		colorful.MustParseHex("#00FF00"),
		colorful.MustParseHex("#00FFFF"),
		colorful.MustParseHex("#0000FF"),
		colorful.MustParseHex("#8000FF"),
		colorful.MustParseHex("#FF00FF"),
	)
)

var (
	shirayukiGold   = colorful.MustParseHex("#F5CA1C")
	shirayukiPurple = colorful.MustParseHex("#711FCF")
	sukoyaPink      = colorful.MustParseHex("#F521CF")
	sukoyaWhite     = colorful.MustParseHex("#FFFCFF")
)

var (
	magnetRed       = colorful.MustParseHex("#600F45")
	magnetPurpleRed = colorful.MustParseHex("#8A317C")
	magnetPurple    = colorful.MustParseHex("#B241BA")
	magnetPink      = colorful.MustParseHex("#C856D9")
	magnetWhite     = colorful.MustParseHex("#FFBEFF")
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
		magnetRed,
		magnetPurpleRed,
		magnetPurple,
		magnetPink,
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
		{magnetRed, 0.0},
		{magnetPurpleRed, 0.25},
		{magnetPurple, 0.50},
		{magnetPink, 0.75},
		{magnetWhite, 1.00},
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

	//shirayukiGradient = gradient.Table{
	//	{shirayukiPurple, 0},
	//	{shirayukiGold, 0.33},
	//	{shirayukiPurple, 0.5},
	//	{shirayukiGold, 0.66},
	//	{shirayukiPurple, 1},
	//}

	shirayukiGradient = gradient.New(
		shirayukiPurple,
		shirayukiGold,
		shirayukiGold,
		shirayukiPurple,
	)

	shirayukiSingleGradient = gradient.New(
		shirayukiPurple,
		shirayukiGold,
	)

	sukoyaGradient = gradient.New(
		sukoyaPink,
		sukoyaWhite,
		sukoyaWhite,
		sukoyaPink,
	)

	shirayukiWhiteGradient = gradient.New(
		shirayukiPurple,
		magnetWhite,
	)

	sukoyaSingleGradient = gradient.New(
		sukoyaPink,
		sukoyaWhite,
	)

	sukoyaWhiteGradient = gradient.New(
		sukoyaPink,
		magnetWhite,
	)
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

	leadIn1 := NewLeadIn(p, 4)
	leadIn1.Play()

	intro1 := NewIntro(p, 16)
	intro1.Play()

	verse1 := NewVerse1a(p, 52)
	verse1.Play()

	verse2 := NewVerse1b(p, 84)
	verse2.Play()

	chorus := NewChorus(p, 114)
	chorus.Play()

	breakdown := NewBreakdown(p, 149)
	breakdown.Play()

	verse3 := NewVerse1a(p, 164)
	verse3.Play()

	verse4 := NewVerse1b(p, 196)
	verse4.Play()

	chorus2 := NewChorus(p, 226)
	chorus2.Play()

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
//			e.SetColor(magnetRainbow.Ierp(
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
