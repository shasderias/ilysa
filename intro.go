package main

import (
	"math/rand"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	light2 "ilysa/pkg/chroma/lightid"
	"ilysa/pkg/colorful"
	"ilysa/pkg/colorful/gradient"
	"ilysa/pkg/ease"
	"ilysa/pkg/ilysa"
	"ilysa/pkg/util"
)

func IntroRhythm(p *ilysa.Project, startBeat, endBeat float64) {
	var (
		steps = int(endBeat - startBeat)
	)
	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx *ilysa.Context) {
		set := magnetColors

		switch {
		case ctx.Ordinal == 0:
			e := ctx.NewPreciseRotationEvent()
			e.Rotation = 180
			e.Step = 0
			e.Prop = 1
			e.Speed = 24
			br := ctx.NewRGBLightingEvent(beatsaber.EventTypeRingLights, beatsaber.EventValueLightRedFade)
			br.SetColor(set.Pick(ctx.Ordinal))
		case ctx.Ordinal%2 == 1:
			e := ctx.NewPreciseRotationEvent()
			e.Rotation = 12.5
			e.Step = 10 + float64(ctx.Ordinal)
			e.Prop = 20
			e.Speed = 20
			e.Direction = chroma.CounterClockwise

			lrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventValueLightRedFade)
			lrl.SetColor(magnetPurple)

			rrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeRightRotatingLasers, beatsaber.EventValueLightBlueFade)
			rrl.SetColor(magnetPink)
		case ctx.Ordinal%2 == 0:
			br := ctx.NewRGBLightingEvent(beatsaber.EventTypeRingLights, beatsaber.EventValueLightBlueFade)
			br.SetColor(set.Pick(ctx.Ordinal))

			lrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventValueLightBlueFade)
			lrl.SetColor(magnetPink)

			rrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeRightRotatingLasers, beatsaber.EventValueLightRedFade)
			rrl.SetColor(magnetPurple)
		}
	})
}

func IntroRhythmSplash(p *ilysa.Project, startBeat, endBeat float64) {
	var (
		steps = int(endBeat - startBeat)
	)
	p.EventsForRange(startBeat, endBeat, steps, ease.Linear, func(ctx *ilysa.Context) {
		set := magnetColors

		br := ctx.NewRGBLightingEvent(beatsaber.EventTypeRingLights, beatsaber.EventValueLightRedFlash)
		br.SetColor(set.Pick(ctx.Ordinal))

	})
}

func IntroMelody1(p *ilysa.Project, startBeat float64) {
	var (
		sequence    = []float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75}
		light       = beatsaber.EventTypeBackLasers
		maxLightID  = p.ActiveDifficultyProfile().MaxLightID(light)
		allLightIDs = light2.FromInterval(1, maxLightID)
		lightIDSet  = light2.Divide(allLightIDs, 3)
	)

	p.EventsForSequence(startBeat, sequence, func(ctx *ilysa.Context) {
		e := ctx.NewRGBLightingEvent(beatsaber.EventTypeBackLasers, beatsaber.EventValueLightBlueOn)
		e.SetColor(magnetPurple)
		e.SetLightID(lightIDSet.Pick(ctx.Ordinal))

		if ctx.Ordinal > 0 {
			oe := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
			oe.SetLightID(lightIDSet.Pick(ctx.Ordinal - 1))
		}
	})

	p.EventForBeat(startBeat+2.999, func(ctx *ilysa.Context) {
		ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
	})
}

func IntroMelody2(p *ilysa.Project, startBeat float64, reverseZoom bool) {
	var (
		sequence    = []float64{0, 0.25, 0.50}
		light       = beatsaber.EventTypeBackLasers
		maxLightID  = p.ActiveDifficultyProfile().MaxLightID(light)
		allLightIDs = light2.FromInterval(1, maxLightID)
		lightIDSet  = light2.Fan(allLightIDs, 3)
	)

	p.EventForBeat(startBeat-0.001, func(ctx *ilysa.Context) {
		ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
	})

	p.EventsForSequence(startBeat, sequence, func(ctx *ilysa.Context) {
		e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightBlueOn)
		e.SetColor(magnetPink)
		e.SetLightID(lightIDSet.Pick(ctx.Ordinal))

		ze := ctx.NewPreciseZoomEvent()
		if reverseZoom {
			ze.Step = 0.3
		} else {
			ze.Step = -0.3
		}

		if ctx.Ordinal > 0 {
			oe := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
			oe.SetLightID(lightIDSet.Pick(ctx.Ordinal - 1))
		}
	})

	p.EventForBeat(startBeat+0.749, func(ctx *ilysa.Context) {
		ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
	})
}

func IntroMelody3(p *ilysa.Project, startBeat float64) {
	var (
		sequence    = []float64{0, 0.5, 1, 1.25, 1.75, 2.25, 2.75, 3.00, 3.25, 3.50}
		light       = beatsaber.EventTypeBackLasers
		maxLightID  = p.ActiveDifficultyProfile().MaxLightID(light)
		allLightIDs = light2.FromInterval(1, maxLightID)
		lightIDSet  = light2.Divide(allLightIDs, 3)
	)

	p.EventsForSequence(startBeat, sequence, func(ctx *ilysa.Context) {
		e := ctx.NewRGBLightingEvent(beatsaber.EventTypeBackLasers, beatsaber.EventValueLightBlueOn)
		e.SetColor(magnetPurple)
		e.SetLightID(lightIDSet.Pick(ctx.Ordinal))

		if ctx.Ordinal > 0 {
			oe := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
			oe.SetLightID(lightIDSet.Pick(ctx.Ordinal - 1))
		}
	})

	p.EventForBeat(startBeat+3.999, func(ctx *ilysa.Context) {
		ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
	})
}

func IntroChorus(p *ilysa.Project, startBeat float64) {
	var (
		sequence   = []float64{0, 1, 2, 2.75, 3.5, 4}
		light      = beatsaber.EventTypeBackLasers
		maxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
		colorGrad  = allColorsGradient
	)

	p.EventForBeat(startBeat, func(ctx *ilysa.Context) {
		ctx.NewRGBLightingEvent(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventValueLightOff)
		ctx.NewRGBLightingEvent(beatsaber.EventTypeRightRotatingLasers, beatsaber.EventValueLightOff)
	})

	p.EventsForSequence(startBeat, sequence, func(ctx *ilysa.Context) {
		ze := ctx.NewPreciseZoomEvent()
		ze.Step = 0.2

		ctx.NewPreciseRotationSpeedEvent(ilysa.LeftLaser, 1).PreciseLaser = chroma.PreciseLaser{
			LockPosition: false, Speed: 0, Direction: chroma.Clockwise,
		}

		ctx.NewPreciseRotationSpeedEvent(ilysa.RightLaser, 1).PreciseLaser = chroma.PreciseLaser{
			LockPosition: false, Speed: 0, Direction: chroma.CounterClockwise,
		}

		grad := append(gradient.Table{}, colorGrad...)
		rand.Shuffle(len(colorGrad), func(i, j int) {
			grad[i].Col, grad[j].Col = grad[j].Col, grad[i].Col
		})

		for i := 1; i <= maxLightID; i++ {
			gradientPos := util.Scale(1, float64(maxLightID), 0, 1)
			color := grad.GetInterpolatedColorFor(gradientPos(float64(i)))

			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
			e.SetSingleLightID(i)
			e.SetColor(color)
			e.Beat += 1.0 / 64.0
		}

		e := ctx.NewPreciseRotationEvent()
		e.Rotation = 45
		e.Step = 5 + (1.5 * float64(ctx.Ordinal))
		e.Prop = 20
		e.Speed = 4
		if ctx.Ordinal%2 == 0 {
			e.Direction = chroma.Clockwise
		} else {
			e.Direction = chroma.CounterClockwise
		}

		if ctx.Ordinal == 5 {
			e.Rotation = 360
		}
	})
}

func IntroPianoRoll(p *ilysa.Project, startBeat float64, count int) {
	var (
		light      = beatsaber.EventTypeBackLasers
		maxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
		lightIDs   = light2.FromInterval(1, maxLightID)
		lightIDSet = light2.Divide(lightIDs, count)
	)
	p.EventsForBeats(startBeat, 0.25, count, func(ctx *ilysa.Context) {
		e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
		e.SetLightID(lightIDSet[ctx.Ordinal])
	})
}

func IntroTrill(p *ilysa.Project, startBeat float64) {
	var (
		light      = beatsaber.EventTypeBackLasers
		maxLightID = p.ActiveDifficultyProfile().MaxLightID(light)
		step       = 0.125
		count      = 5
		ratio      = 0.666
		lightCount = int(ratio * float64(maxLightID))
	)

	p.EventsForBeats(startBeat, step, count, func(ctx *ilysa.Context) {
		for i := 0; i < lightCount; i++ {
			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
			e.SetSingleLightID(rand.Intn(maxLightID) + 1)
			e.SetColor(allColorsGradient.GetInterpolatedColorFor(rand.Float64()))

			oe := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
			oe.Beat += step / 2

			if !ctx.Last {
				continue
			}

		}

	})
	p.EventForBeat(startBeat+0.5, func(ctx *ilysa.Context) {
		Shimmer(p, ctx.B, ctx.B+1.2, 30, beatsaber.EventTypeRingLights, 0.4, 0.7)
	})
}

func IntroClimb(p *ilysa.Project, startBeat float64) {
	var (
		light           = beatsaber.EventTypeBackLasers
		maxLightID      = p.ActiveDifficultyProfile().MaxLightID(light)
		step            = 0.25
		count           = 12
		lightIDs        = light2.FromInterval(1, maxLightID)
		lightIDSets     = light2.Divide(lightIDs, maxLightID/2)
		lightIDSequence = []int{6, 7, 5, 8, 4, 9, 3, 10, 2, 11, 1, 12}
		backGrad        = gradient.Table{
			{magnetPink, 0.0},
			{magnetWhite, 1.0},
		}
		sideGrad = gradient.Table{
			{magnetWhite, 0.0},
			{magnetPurple, 1.0},
		}
	)

	p.EventForBeat(startBeat, func(ctx *ilysa.Context) {
		e := ctx.NewPreciseRotationEvent()
		e.Rotation = 360
		e.Step = 15
		e.Speed = 1.3
		e.Prop = 13

		ctx.NewZoomEvent()
	})

	p.EventsForBeats(startBeat, step, count, func(ctx *ilysa.Context) {
		e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedOn)
		e.SetColor(backGrad.GetInterpolatedColorFor(ctx.Pos))
		e.SetLightID(lightIDSets.Pick(lightIDSequence[ctx.Ordinal]))

		switch {
		case ctx.Last:
			const exitValue = 3
			lrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventValueLightBlueFade)
			lrl.SetColor(magnetPurple)
			ctx.NewPreciseRotationSpeedEvent(ilysa.LeftLaser, exitValue).PreciseLaser = chroma.PreciseLaser{
				LockPosition: true, Speed: exitValue, Direction: chroma.CounterClockwise,
			}

			rrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeRightRotatingLasers, beatsaber.EventValueLightRedFade)
			rrl.SetColor(magnetPurple)
			ctx.NewPreciseRotationSpeedEvent(ilysa.RightLaser, exitValue).PreciseLaser = chroma.PreciseLaser{
				LockPosition: false, Speed: exitValue, Direction: chroma.Clockwise,
			}

		case ctx.Ordinal%2 == 0:
			lrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventValueLightBlueFlash)
			lrl.SetColor(sideGrad.GetInterpolatedColorFor(ctx.Pos))
			ctx.NewPreciseRotationSpeedEvent(ilysa.LeftLaser, ctx.Ordinal).PreciseLaser = chroma.PreciseLaser{
				LockPosition: true, Speed: float64(ctx.Ordinal), Direction: chroma.Clockwise,
			}

			ctx.NewRGBLightingEvent(beatsaber.EventTypeRightRotatingLasers, beatsaber.EventValueLightOff)
		case ctx.Ordinal%2 == 1:
			ctx.NewRGBLightingEvent(beatsaber.EventTypeLeftRotatingLasers, beatsaber.EventValueLightOff)

			rrl := ctx.NewRGBLightingEvent(beatsaber.EventTypeRightRotatingLasers, beatsaber.EventValueLightRedFlash)
			rrl.SetColor(sideGrad.GetInterpolatedColorFor(ctx.Pos))
			ctx.NewPreciseRotationSpeedEvent(ilysa.RightLaser, ctx.Ordinal).PreciseLaser = chroma.PreciseLaser{
				LockPosition: false, Speed: float64(ctx.Ordinal), Direction: chroma.CounterClockwise,
			}
		}
	})
}

func IntroFall(p *ilysa.Project, startBeat float64) {
	var (
		light       = beatsaber.EventTypeBackLasers
		maxLightID  = p.ActiveDifficultyProfile().MaxLightID(light)
		step        = 0.25
		count       = 4
		lightIDs    = light2.FromInterval(1, maxLightID)
		lightIDSets = light2.Divide(lightIDs, count)
		colorSet    = colorful.NewSet(magnetPurple, magnetPink)
		values      = []beatsaber.EventValue{
			beatsaber.EventValueLightRedOn,
			beatsaber.EventValueLightOff,
			beatsaber.EventValueLightBlueOn,
			beatsaber.EventValueLightRedFlash,
		}
	)

	p.EventsForBeats(startBeat, step, count, func(ctx *ilysa.Context) {
		e := ctx.NewRGBLightingEvent(light, values[ctx.Ordinal])
		e.SetColor(colorSet.Next())
		if ctx.Ordinal <= 3 {
			e.SetLightID(lightIDSets[ctx.Ordinal])
		}
	})
}

func IntroBridge(p *ilysa.Project, startBeat float64) {
	p.EventForBeat(startBeat, func(ctx *ilysa.Context) {
		re := ctx.NewPreciseRotationEvent()
		re.Rotation = 180
		re.Step = 12.5
		re.Direction = chroma.CounterClockwise
		re.Speed = 3
		re.Prop = 5
		re.CounterSpin = true
	})

	p.EventsForRange(startBeat, startBeat+1, 30, ease.OutCubic, func(ctx *ilysa.Context) {
		if !ctx.Last {
			e := ctx.NewRGBLightingEvent(beatsaber.EventTypeBackLasers, beatsaber.EventValueLightBlueOn)
			e.SetColor(magnetPurple)
			e.SetAlpha(1 - ctx.Pos)
		} else {
			e := ctx.NewRGBLightingEvent(beatsaber.EventTypeBackLasers, beatsaber.EventValueLightRedFade)
			e.SetColor(magnetWhite)
		}
	})
}

func IntroOutro(p *ilysa.Project, startBeat float64) {
	var (
		light      = beatsaber.EventTypeBackLasers
		maxLaserID = p.ActiveDifficultyProfile().MaxLightID(light)
		sequence   = []float64{0, 0.25, 0.50, 1.0, 1.25, 1.50, 2.0, 2.25, 2.50, 2.75, 3.25}
		lightIDSet = light2.Fan(light2.FromInterval(1, maxLaserID), len(sequence))
	)

	p.EventForBeat(startBeat-0.001, func(ctx *ilysa.Context) {
		for i := 1; i <= maxLaserID; i++ {
			e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightRedFlash)
			e.SetSingleLightID(i)
			e.SetColor(allColorsGradient.GetInterpolatedColorFor(float64(i) / float64(maxLaserID)))
		}
	})

	p.EventsForSequence(startBeat, sequence, func(ctx *ilysa.Context) {
		e := ctx.NewRGBLightingEvent(light, beatsaber.EventValueLightOff)
		e.SetLightID(lightIDSet[ctx.Ordinal])
	})
}

func Impact(ctx *ilysa.Context, light beatsaber.EventType, value beatsaber.EventValue, grad gradient.Table) {
	maxID := ctx.ActiveDifficultyProfile().MaxLightID(light)

	for i := 1; i <= maxID; i++ {
		e := ctx.NewRGBLightingEvent(light, value)
		e.SetSingleLightID(i)
		e.SetColor(grad.GetInterpolatedColorFor(float64(i) / float64(maxID)))
	}
}

func IntroOutroSplash(p *ilysa.Project, startBeat float64) {
	var (
		sequence   = []float64{0, 0.75, 1.5}
		leftLaser  = beatsaber.EventTypeLeftRotatingLasers
		rightLaser = beatsaber.EventTypeRightRotatingLasers
	)

	Shimmer(p, startBeat, startBeat+2, 60, beatsaber.EventTypeBackLasers, 0.6, 0.9)

	p.EventsForSequence(startBeat, sequence, func(ctx *ilysa.Context) {
		Impact(ctx, leftLaser, beatsaber.EventValueLightBlueFade, allColorsGradient)
		Impact(ctx, rightLaser, beatsaber.EventValueLightRedFade, allColorsGradient)
	})
}
