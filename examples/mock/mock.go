package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

const BPMChangesJSON = `
[
	{
		"_time": 323.103,
		"_BPM": 108,
		"_beatsPerBar": 4,
		"_metronomeOffset": 4
	},
	{
		"_time": 325.103,
		"_BPM": 108,
		"_beatsPerBar": 4,
		"_metronomeOffset": 4
	},
	{
		"_time": 326.103,
		"_BPM": 108,
		"_beatsPerBar": 4,
		"_metronomeOffset": 4
	}
]`

func main() {
	m, err := beatsaber.NewMockMap(beatsaber.EnvironmentFitBeat, 108, BPMChangesJSON)
	if err != nil {
		panic(err)
	}

	p := ilysa.New(m)

	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewLighting(evt.WithLight(evt.LeftRotatingLasers), evt.WithLightValue(evt.LightBlueFlash))
	})

	p.Dump()
}
