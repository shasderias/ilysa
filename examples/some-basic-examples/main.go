package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
)

// set mapPath to the directory containing your beatmap
const mapPath = `D:\Beat Saber Data\CustomWIPLevels\MagnetLights`

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

	p.EventForBeat(2, func(ctx ilysa.TimeContext) { // generate events for beat 2:
		ctx.NewLightingEvent( // generate a new base game (non-Chroma) lighting event
			ilysa.WithType(beatsaber.EventTypeBackLasers),   // back lasers
			ilysa.WithValue(beatsaber.EventValueLightRedOn), // red on
		)

		ctx.NewLightingEvent( // generate a new base game (non-Chroma) lighting event
			ilysa.WithType(beatsaber.EventTypeLeftRotatingLasers), // left lasers
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),    // blue fade
		)
	})

	p.EventForBeat(2, func(ctx ilysa.TimeContext) {
		ctx.NewLightingEvent(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
		)

		ctx.NewRotationEvent(
			ilysa.WithType(beatsaber.EventTypeInterscopeRaiseHydraulics),
			ilysa.WithValue(1),
		)

		ctx.NewZoomEvent()

		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(3),
		)

		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedOn),
			ilysa.WithColor(colorful.MustParseHex("#123123")),
			ilysa.WithAlpha(0.3),
			ilysa.WithLightID(ilysa.NewLightID(1, 2, 3)),
		)

		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithLockPosition(true),
			ilysa.WithIntValue(1),
			ilysa.WithSpeed(0),
			ilysa.WithDirection(chroma.Clockwise),
		)

		ctx.NewPreciseRotationEvent(
			ilysa.WithNameFilter("BigTrackLaneRings"),
			ilysa.WithReset(false),
			ilysa.WithRotation(45),
			ilysa.WithStep(15.0),
			ilysa.WithProp(0.5),
			ilysa.WithSpeed(3),
			ilysa.WithDirection(chroma.Clockwise),
			ilysa.WithCounterSpin(true),
		)
		ctx.NewPreciseZoomEvent(
			ilysa.WithStep(4),
		)
	})

	// save events back to Expert+ difficulty
	return p.Save()
}
