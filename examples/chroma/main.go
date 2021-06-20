package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful"
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

	// -- your code goes here --
	p.EventForBeat(2, func(ctx ilysa.TimeContext) {
		ctx.NewRGBLightingEvent(
			ilysa.WithType(beatsaber.EventTypeBackLasers),
			ilysa.WithValue(beatsaber.EventValueLightRedOn),
			ilysa.WithColor(colorful.MustParseHex("#123123")),
			ilysa.WithAlpha(0.3),
			ilysa.WithLightID(ilysa.NewLightID(1, 2, 3)),
		)
	})

	// save events back to Expert+ difficulty
	return p.Save()
}
