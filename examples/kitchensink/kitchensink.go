package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
)

var (
	LaneRight1      = light.NewCustom(evt.BackLasers, 20, 4)
	LaneLeft1       = light.NewCustom(evt.BackLasers, 20, 24)
	LaneHorizLeft1  = light.NewCustom(evt.BackLasers, 10, 44)
	LaneHorizRight1 = light.NewCustom(evt.BackLasers, 10, 54)
	LaneHorizLeft2  = light.NewCustom(evt.BackLasers, 10, 64)
	LaneHorizRight2 = light.NewCustom(evt.BackLasers, 10, 74)
	CloseBarTop     = light.NewCustom(evt.BackLasers, 1, 84)
	CloseBarBottom  = light.NewCustom(evt.BackLasers, 1, 85)
	FarBarTop       = light.NewCustom(evt.BackLasers, 1, 86)
	FarBarBottom    = light.NewCustom(evt.BackLasers, 1, 87)
	LeftBarTop      = light.NewCustom(evt.BackLasers, 1, 88)
	LeftBarBottom   = light.NewCustom(evt.BackLasers, 1, 89)
	RightBarTop     = light.NewCustom(evt.BackLasers, 1, 90)
	RightBarBottom  = light.NewCustom(evt.BackLasers, 1, 91)
)

var (
	Red         = colorful.MustParseHex("#FF0000")
	Blue        = colorful.MustParseHex("#0000FF")
	RedBlueGrad = gradient.New(Red, Blue)
)

// set mapPath to the directory containing your beatmap
const mapPath = `D:\Beat Saber Data\CustomWIPLevels\Ilysa`

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

	ctx := p.Offset(0)

	// back lasers red on at beat 0
	ctx.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.NewLighting(
			evt.WithLight(evt.BackLasers),
			evt.WithLightValue(evt.LightRedOn),
		)
	})

	// back lasers blue fade events at beats 0, 0.5, 1.0, then off at 1.5
	ctx.Sequence(timer.Seq([]float64{0, 0.5, 1.0}, 1.5), func(ctx context.Context) {
		ctx.NewLighting(
			evt.WithLight(evt.BackLasers),
			evt.WithLightValue(evt.LightBlueFade),
		)

		if ctx.SeqLast() { // if this is the last step of the sequence
			ctx.NewLighting( // create an off event ...
				evt.WithLight(evt.BackLasers),
				evt.WithLightValue(evt.LightOff),
				evt.WithBOffset(ctx.SeqNextBOffset()), // ... offset it by 0.5 beats (difference between ghostBeat and last beat)
			)
		}
	})

	// 11 ring blue flash events from beat 0 to 1 (i.e. 0, 0.1, 0.2 ... 1.0)
	ctx.Range(timer.Rng(0, 1, 11, ease.Linear), func(ctx context.Context) {
		ctx.NewLighting(
			evt.WithLight(evt.RingLights),
			evt.WithLightValue(evt.LightBlueFlash),
		)
	})

	// events at beats 0.0, 0.5, 1.0, 2.0, 2.5, 3.0
	ctx.Sequence(timer.Seq([]float64{0, 2}, 2), func(ctx context.Context) {
		ctx.Range(timer.Rng(0, 1, 3, ease.Linear), func(ctx context.Context) {
			//  events here
		})
	})

	ctx.NewLighting(
		evt.WithLight(evt.BackLasers),
		evt.WithLightValue(evt.LightBlueFade),
		evt.WithBOffset(0),
	)

	evt.NewRGBLighting(
		evt.WithLight(evt.BackLasers),
		evt.WithLightValue(evt.LightBlueFade),
		evt.WithBeat(0),
		evt.WithBOffset(0),

		evt.WithColor(Red),
		evt.WithAlpha(2),
	)

	// print events
	return p.Save()
}
