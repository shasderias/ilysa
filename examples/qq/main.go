package main

import (
	"fmt"
	"math/rand"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/fx"
)

// set mapPath to the directory containing your beatmap
const mapPath = `D:\Beat Saber Data\CustomWIPLevels\Ilysa`

// please use a working copy dedicated to Ilysa (and make backups!) as Ilysa
// WILL OVERWRITE ALL LIGHTING EVENTS IN THE SELECTED DIFFICULTY

func main() {
	if err := do(); err != nil {
		fmt.Println("error:", err)
	}
}

func do() error {
	// open the beatmap at mapPath
	bsMap, err := beatsaber.Open(mapPath)
	if err != nil {
		return err
	}

	// create a new Ilysa project
	p := ilysa.New(bsMap)

	// load the Expert+ difficulty with the standard characteristic
	err = p.Map.SetActiveDifficulty(beatsaber.CharacteristicStandard, beatsaber.BeatmapDifficultyExpertPlus)
	if err != nil {
		return err
	}

	// get an Ilysa light representing a base game back laser (i.e. only "1" lightID)
	backLasers := ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p)

	// transform the light into a back laser light with 1 lightID for each lightID it has in the beatmap's environment
	// i.e. make it work like ChroMapper in LightID mode
	backLasersSplit := ilysa.TransformLight(backLasers,
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)
	const qqLevel = 1

	switch qqLevel {
	case 1:
		p.EventsForBeats(0, 1, 50, func(ctx ilysa.TimeContext) {
			ctx.NewLightingEvent(
				ilysa.WithType(beatsaber.EventTypeBackLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
			)
		})
	case 2:
		p.EventsForBeats(0, 1, 50, func(ctx ilysa.TimeContext) {
			ctx.NewRGBLightingEvent(
				ilysa.WithType(beatsaber.EventTypeBackLasers),
				ilysa.WithValue(beatsaber.EventValueLightRedFade),
				ilysa.WithColor(gradient.Rainbow.Ierp(rand.Float64())),
			)
		})
	case 3:
		// get an Ilysa light representing a base game back laser (i.e. only "1" lightID)
		backLasers := ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p)

		// transform the light into a back laser light with 1 lightID for each lightID it has in the beatmap's environment
		// i.e. make it work like ChroMapper in LightID mode
		backLasersSplit := ilysa.TransformLight(backLasers,
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		)

		p.EventsForBeats(0, 1, 50, func(ctx ilysa.TimeContext) {
			ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) { // use the light we created
				ctx.NewRGBLightingEvent(
					ilysa.WithValue(beatsaber.EventValueLightRedFade),
					ilysa.WithColor(gradient.Rainbow.Ierp(rand.Float64())),
				)
			})
		})
	case 4:
		// light creation code omitted for brevity
		p.EventsForBeats(0, 1, 50, func(ctx ilysa.TimeContext) {
			ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) {
				e := ctx.NewRGBLightingEvent( // save the event we created to the variable e
					ilysa.WithValue(beatsaber.EventValueLightRedFade),
					ilysa.WithColor(gradient.Rainbow.Ierp(rand.Float64())),
				)
				// shift each event forward by 0.05 beats * ordinal number of the lightID
				// i.e. 1st lightID is shifted forward by 0 beats
				//      2nd lightID is shifted forward by 0.05 beats, etc
				e.ShiftBeat(float64(ctx.LightIDOrdinal()) * 0.05)
			})
		})
	case 5:
		// light creation code omitted for brevity
		p.EventsForBeats(0, 1, 50, func(ctx ilysa.TimeContext) {
			ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) {
				e := ctx.NewRGBLightingEvent(
					ilysa.WithValue(beatsaber.EventValueLightRedFade),
					ilysa.WithColor(gradient.Rainbow.Ierp(rand.Float64())),
				)
				e.ShiftBeat(float64(ctx.LightIDOrdinal()) * 0.05)

				oe := ctx.NewRGBLightingEvent( // create an off event
					ilysa.WithValue(beatsaber.EventValueLightOff),
				)
				oe.ShiftBeat(float64(ctx.LightIDOrdinal())*0.05 + 0.1) // shift it forward by an additional 0.1 beat
			})
		})
	case 6:
		// light creation code omitted for brevity
		p.EventsForBeats(0, 1, 1, func(ctx ilysa.TimeContext) {
			// for each beat, create events at regular intervals from beat to beat + 0.5 beats, for a total of 8 beats
			ctx.EventsForRange(ctx.T(), ctx.T()+0.5, 8, ease.Linear, func(ctx ilysa.TimeContext) {
				ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) {
					e := ctx.NewRGBLightingEvent(
						// ilysa.WithValue(beatsaber.EventValueLightRedFade), // we never needed this
						ilysa.WithColor(gradient.Rainbow.Ierp(rand.Float64())),
					)
					e.ShiftBeat(float64(ctx.LightIDOrdinal()) * 0.05)
					e.SetAlpha(1 - ctx.T()) // linear alpha fade to 0
				})
			})
		})
	case 7:
		// light creation code omitted for brevity
		p.EventsForBeats(0, 1, 1, func(ctx ilysa.TimeContext) {
			ctx.EventsForRange(ctx.T(), ctx.T()+0.5, 8, ease.Linear, func(ctx ilysa.TimeContext) {
				color := gradient.Rainbow.Ierp(rand.Float64())
				ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) {
					e := ctx.NewRGBLightingEvent(
						ilysa.WithColor(color),
					)
					e.ShiftBeat(float64(ctx.LightIDOrdinal()) * 0.05)
					e.SetAlpha(1 - ctx.T())
				})
			})
		})
	case 8:
		// light creation code omitted for brevity
		p.EventsForBeats(0, 1, 1, func(ctx ilysa.TimeContext) {
			ctx.EventsForRange(ctx.T(), ctx.T()+0.5, 8, ease.Linear, func(ctx ilysa.TimeContext) {
				ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) {
					// the fx package contains a suite of building blocks you can use to build more complicated effects
					// the Gradient function generates events and colors them based on the gradient passed to it
					e := fx.Gradient(ctx, gradient.Rainbow)

					e.ShiftBeat(float64(ctx.LightIDOrdinal()) * 0.05)
					e.SetAlpha(1 - ctx.T())
				})
			})
		})
	case 9:
		//light creation code omitted for brevity
		p.EventsForBeats(0, 1, 1, func(ctx ilysa.TimeContext) {
			ctx.EventsForRange(ctx.T(), ctx.T()+0.5, 8, ease.Linear, func(ctx ilysa.TimeContext) {
				ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) {
					// ColorSweep is a more advanced Gradient that shifts the gradient's position with time
					// the 2nd argument (1.2 below) controls the speed at which the gradient "moves"
					e := fx.ColorSweep(ctx, 1.2, gradient.Rainbow)

					// we then use fx.Ripple to stagger the start time of each lightID
					fx.Ripple(ctx, e, 0.2)

					e.SetAlpha(1 - ctx.T())
				})
			})
		})
	case 10:
		// light creation code omitted for brevity
		p.EventsForBeats(0, 1, 1, func(ctx ilysa.TimeContext) {
			ctx.EventsForRange(ctx.T(), ctx.T()+0.5, 8, ease.Linear, func(ctx ilysa.TimeContext) {
				ctx.WithLight(backLasersSplit, func(ctx ilysa.TimeLightContext) {
					e := fx.ColorSweep(ctx, 1.2, gradient.Rainbow)

					fx.Ripple(ctx, e, 0.2)

					// fx.AlphaBlend does what it says on the tin, it accepts in order:
					// startT, endT, startAlpha, endAlpha and a ease function
					fx.AlphaBlend(ctx, e, 0, 1, 1, 0, ease.OutCirc)
				})
			})
		})
	}

	// save events back to Expert+ difficulty
	return p.Save()
}
