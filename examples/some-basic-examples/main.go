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

	p.EventForBeat(4, func(ctx ilysa.TimeContext) {
		// beat 4, back lasers, blue flash
		ctx.NewLightingEvent(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
		)

		// beat 4, raise hydrualics, all cars
		ctx.NewRotationEvent(
			ilysa.WithType(beatsaber.EventTypeInterscopeRaiseHydraulics),
			ilysa.WithValue(1),
		)

		// beat 4, zoom event
		ctx.NewZoomEvent()

		// beat 4, left laser, speed 3
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(3),
		)
	})

	p.EventsForBeats(6, 2, 5, func(ctx ilysa.TimeContext) {
		// beats 6, 8, 10, 12, 14 ...

		// ... back lasers, red on event with Chroma, color #123123, alpha 0.3, lightIDs 1, 2, 3
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedOn),
			ilysa.WithColor(colorful.MustParseHex("#123123")),
			ilysa.WithAlpha(0.3),
			ilysa.WithLightID(ilysa.NewLightID(1, 2, 3)),
		)

		// ... Chroma precision rotation speed event, lock positions, value 3, precise speed 0, clockwise direction
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithLockPosition(true),
			ilysa.WithIntValue(1),
			ilysa.WithSpeed(0),
			ilysa.WithDirection(chroma.Clockwise),
		)

		// ... Chroma precise rotation event ...
		ctx.NewPreciseRotationEvent(
			ilysa.WithNameFilter("BigTrackLaneRings"), // ... nameFilter "BigTrackLaneRings"
			ilysa.WithReset(false),                    // no reset
			ilysa.WithRotation(45),                    // rotation 45
			ilysa.WithStep(15.0),                      // step 15
			ilysa.WithProp(0.5),                       // prop 0.5
			ilysa.WithSpeed(3),                        // speed 3
			ilysa.WithDirection(chroma.Clockwise),     // clockwise spin
			ilysa.WithCounterSpin(true),               // inner rings counter spin
		)
		// ... Chroma precise zoom event, step 4
		ctx.NewPreciseZoomEvent(
			ilysa.WithStep(4),
		)
	})

	// generate events on beats 0, 0.25, 0.75 and 1.25, starting from beat 4
	// i.e. 4.00, 4.25, 4.75, 5.25
	p.EventsForSequence(4, []float64{0, 0.25, 0.75, 1.25}, func(ctx ilysa.SequenceContext) {
		ctx.NewLightingEvent(
			ilysa.WithType(beatsaber.EventTypeRingLights),
			ilysa.WithValue(beatsaber.EventValueLightBlueFade),
		)
	})

	// save events back to Expert+ difficulty
	return p.Save()
}
