package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
)

// set mapPath to the directory containing your beatmap
const mapPath = `C:\directory\containing\your\beatmap\goes\here`

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

	// generate events every quarter beat (0.25), starting at beat 3, do this a total of 16 times
	// i.e. 3.00, 3.25, 3.50, 3.75, 4.00 ... 6.50, 6.75
	p.EventsForBeats(3, 0.25, 16, func(ctx ilysa.TimeContext) {
		// each time, generate a rotation speed event
		ctx.NewRotationSpeedEvent(
			// that controls the left laser's rotation speed
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			// ctx.Ordinal() returns the iteration number, starting with 0
			// i.e. for beat 3.00, ctx.Ordinal is 0, for beat 3.25, ctx.Ordinal is 1
			// the following line will therefore increase the left laser's rotation speed from 0 to 15 over 3.75 beats
			ilysa.WithIntValue(ctx.Ordinal()),
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
