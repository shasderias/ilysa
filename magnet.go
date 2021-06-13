package main

import (
	"fmt"
	"math"
	"math/rand"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/chroma/lightid"
	"ilysa/pkg/colorful"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ease"
	"ilysa/pkg/genfx"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/modfx"
	"ilysa/pkg/util"
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
		sukoyaPink,
		sukoyaWhite,
		magnetPurple,
		magnetPink,
		magnetWhite,
	)

	magnetColors = colorful.NewSet(
		magnetPink,
		magnetWhite,
		magnetPurple,
	)

	shirayukiColors = colorful.NewSet(
		shirayukiGold,
		shirayukiPurple,
	)

	sukoyaColors = colorful.NewSet(
		sukoyaPink,
		sukoyaWhite,
	)
)

var (
	magnetGradient = gradient.Table{
		{magnetPurple, 0.0},
		{magnetWhite, 0.5},
		{magnetPink, 1.0},
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

	LeadIn(p)

	p.EventsForSequence(16, []float64{0, 0.75, 1.25, 1.75, 2.25}, func(ctx *ilysa.Context) {
		set := colorful.NewSet(magnetPurple, magnetPink, magnetWhite, colorful.Black)

		types := []beatsaber.EventType{
			beatsaber.EventTypeCenterLights,
			beatsaber.EventTypeRingLights,
			beatsaber.EventTypeBackLasers,
		}

		values := []beatsaber.EventValue{
			beatsaber.EventValueLightRedOn,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedFade,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightOff,
		}

		for _, typ := range types {
			grad := gradient.Table{
				{set.Pick(ctx.Ordinal), 0.0},
				{set.Pick(ctx.Ordinal + 1), 1.0},
			}

			min := 1
			max := ctx.ActiveDifficultyProfile().MaxLightID(typ)

			genfx.GradientProp(ctx, typ, values[ctx.Ordinal], grad, min, max)
		}
	})

	p.EventsForSequence(18.25, []float64{0, 0.25, 0.75, 1, 1.5}, func(ctx *ilysa.Context) {
		var (
			leftLaser            = beatsaber.EventTypeLeftRotatingLasers
			rightLaser           = beatsaber.EventTypeRightRotatingLasers
			leftLaserMaxLightID  = ctx.ActiveDifficultyProfile().MaxLightID(leftLaser)
			rightLaserMaxLightID = ctx.ActiveDifficultyProfile().MaxLightID(rightLaser)
			ordinal              = ctx.Ordinal
		)

		lightIDGroups := [][]int{
			lightid.EveryNthLightID(1, leftLaserMaxLightID, 2, 0),
			lightid.EveryNthLightID(1, leftLaserMaxLightID, 2, 1),
			lightid.EveryNthLightID(1, rightLaserMaxLightID, 2, 0),
			lightid.EveryNthLightID(1, rightLaserMaxLightID, 2, 1),
		}

		types := []beatsaber.EventType{leftLaser, leftLaser, rightLaser, rightLaser}

		values := []beatsaber.EventValue{
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedOn,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedOn,
		}

		ctx.NewPreciseRotationSpeedEvent(ilysa.LeftLaser, 1).PreciseLaser = chroma.PreciseLaser{
			LockPosition: false, Speed: 0, Direction: chroma.Clockwise,
		}
		ctx.NewPreciseRotationSpeedEvent(ilysa.RightLaser, 1).PreciseLaser = chroma.PreciseLaser{
			LockPosition: false, Speed: 0, Direction: chroma.CounterClockwise,
		}

		if ordinal < 4 {
			e := ctx.NewRGBLightingEvent(types[ordinal], values[ordinal])
			e.SetLightID(lightIDGroups[ordinal])
			e.SetColor(magnetColors.Pick(ordinal))
		} else {
			ctx.NewRGBLightingEvent(leftLaser, beatsaber.EventValueLightOff)
			ctx.NewRGBLightingEvent(rightLaser, beatsaber.EventValueLightOff)
		}

		if ordinal > 0 {
			ctx.NewRGBLightingEvent(types[ordinal-1], beatsaber.EventValueLightOff)
		}
	})

	BassTwang(p, 18.5)

	p.EventForBeat(20, func(ctx *ilysa.Context) {
		ctx.NewRotationSpeedEvent(ilysa.LeftLaser, 8)
		ctx.NewRotationSpeedEvent(ilysa.RightLaser, 8)

		lrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventValueLightRedFlash)
		lrl.SetColor(sukoyaPink)
		rrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeRightRotatingLasers, beatsaber.EventValueLightBlueFlash)
		rrl.SetColor(shirayukiPurple)

		cl := ctx.NewRGBLightingEvent(beatsaber.EventTypeCenterLights, beatsaber.EventValueLightRedFlash)
		cl.SetColor(magnetPurple)

		ctx.NewZoomEvent()
	})

	IntroRhythm(p, 20, 23)
	IntroRhythm(p, 24, 27)
	IntroRhythm(p, 28, 31)
	IntroMelody1(p, 20)
	IntroMelody2(p, 23.25, false)
	IntroMelody1(p, 24)
	IntroMelody2(p, 27.25, true)
	IntroMelody3(p, 28)
	IntroChorus(p, 32)
	IntroPianoRoll(p, 36.5, 6)
	IntroTrill(p, 38.5)
	IntroTrill(p, 42.5)
	IntroClimb(p, 39.5)
	IntroFall(p, 43.25)
	IntroTrill(p, 44.5)
	IntroBridge(p, 45.0)
	IntroRhythm(p, 46.0, 50)
	IntroOutro(p, 46.5)
	IntroOutroSplash(p, 50.0)

	return p.Save()
}

func Shimmer(p *ilysa.Project, startBeat, endBeat float64, steps int, light beatsaber.EventType, colorSweepSpeed, shimmerSweepSpeed float64) {
	var (
		duration   = endBeat - startBeat
		offset     = rand.Float64() * math.Pi * 2
		maxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
		sin        = math.Sin
		pi         = math.Pi
	)

	colorSweepSpeed *= duration
	shimmerSweepSpeed *= duration

	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx *ilysa.Context) {
		for i := 1; i <= maxLightID; i++ {
			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
			e.SetSingleLightID(i)
			e.SetColor(gradient.Rainbow.GetInterpolatedColorFor(
				sin(ctx.Pos*colorSweepSpeed + (float64(i)/float64(maxLightID))*pi + offset),
			))
			e.SetAlpha(5)
		}
	})

	//p.ModEventsInRange(startBeat, endBeat,
	//	ilysa.FilterLightingEvents(light),
	//	func(ctx *ilysa.Context, event ilysa.Event) {
	//		e := event.(*ilysa.RGBLightingEvent)
	//		lightID := float64(e.FirstLightID())
	//		e.SetAlpha(e.GetAlpha() * util.DefaultNoise(ctx.Pos*shimmerSweepSpeed+lightID/float64(maxLightID)*pi+offset))
	//	})

	//fadeScale := util.Scale(startBeat, endBeat, 0, 1)
	//
	//modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.2), 0, 1, ease.InCubic)
	//modfx.RGBAlphaFade(p, light, fadeScale(0.9), endBeat, 1, 0, ease.OutCubic)
	fadeScale := util.Scale(startBeat, endBeat, 0, 1)
	modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
	modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutCubic)
}

func BassTwang(p *ilysa.Project, startBeat float64) {
	const steps = 60

	var (
		sin        = math.Sin
		endBeat    = startBeat + 1.495
		light      = beatsaber.EventTypeBackLasers
		maxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
		pi         = math.Pi
	)

	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx *ilysa.Context) {
		for i := 1; i <= maxLightID; i++ {
			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightBlueOn)
			e.SetSingleLightID(i)
			e.SetColor(magnetGradient.GetInterpolatedColorFor(
				sin(ctx.Pos*3 + (float64(i)/float64(maxLightID))*pi + 4),
			))
			e.SetAlpha(3)
		}
	})

	fadeScale := util.Scale(startBeat, endBeat, 0, 1)
	modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
	modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutBounce)
}

func SimpleShimmer(p *ilysa.Project, startBeat, endBeat float64) {
	const steps = 60

	var (
		sin        = math.Sin
		light      = beatsaber.EventTypeBackLasers
		maxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
		pi         = math.Pi
	)

	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx *ilysa.Context) {
		for i := 1; i <= maxLightID; i++ {
			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightBlueOn)
			e.SetSingleLightID(i)
			e.SetColor(magnetGradient.GetInterpolatedColorFor(
				sin(ctx.Pos*3 + (float64(i)/float64(maxLightID))*pi + 4),
			))
			e.SetAlpha(5)
		}
	})

	fadeScale := util.Scale(startBeat, endBeat, 0, 1)
	modfx.RGBAlphaFade(p, light, startBeat, fadeScale(0.5), 0, 1, ease.InCubic)
	modfx.RGBAlphaFade(p, light, fadeScale(0.501), endBeat, 1, 0, ease.OutBounce)
}
