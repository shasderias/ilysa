package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
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

	// -- your code goes here --
	s := timer.SequencerFromSlice([]float64{4, 8, 12, 16})

	ctx := p.WithBeatOffset(2)
	ctx.Sequence(s, func(ctx ilysa.SequenceContext) {
		ctx.NewLighting(evt.WithLight(evt.CenterLights), evt.WithLightValue(evt.LightRedOn))
	})
	ctx.Range()

	// save events back to Expert+ difficulty
	return p.Dump()
}
