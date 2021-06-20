package main

import (
	"fmt"

	"github.com/shasderias/ilysa/pkg/beatsaber"
	"github.com/shasderias/ilysa/pkg/colorful"
	"github.com/shasderias/ilysa/pkg/colorful/gradient"
	"github.com/shasderias/ilysa/pkg/ease"
	"github.com/shasderias/ilysa/pkg/ilysa"
	"github.com/shasderias/ilysa/pkg/ilysa/fx"
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

	ringLights := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeRingLights, p),
		ilysa.ToSequenceLightTransformer(ilysa.Fan(2)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	).(ilysa.SequenceLight)

	light := ilysa.NewCombinedLight(
		ringLights.Index(0),
		ringLights.Index(1),
	)

	ringLightsReverse := ilysa.TransformLight(
		ilysa.NewBasicLight(beatsaber.EventTypeRingLights, p),
		ilysa.ToSequenceLightTransformer(ilysa.Fan(2)),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
		ilysa.LightIDSetTransformerToLightTransformer(ilysa.Reverse),
	).(ilysa.SequenceLight)

	lightReverse := ilysa.NewCombinedLight(
		ringLightsReverse.Index(0),
		ringLightsReverse.Index(1),
	)

	grad := gradient.New(
		colorful.MustParseHex("#95bddc"),
		colorful.MustParseHex("#fbc6d0"),
		colorful.MustParseHex("#0c71c9"),
		colorful.MustParseHex("#ff145f"),
	)

	// grad2 := gradient.Table{
	// 	{Col: colorful.MustParseHex("#fbc6d0"), Pos: 0.0},
	// 	{Col: colorful.MustParseHex("#95bddc"), Pos: 0.2},
	// 	{Col: colorful.MustParseHex("#0c71c9"), Pos: 0.8},
	// 	{Col: colorful.MustParseHex("#ff145f"), Pos: 1.0},
	// }

	// grad2.Ierp(0.3)

	// set := colorful.NewSet(
	// 	colorful.MustParseHex("#fbc6d0"),
	// 	colorful.MustParseHex("#95bddc"),
	// )

	// set.Index(0)

	// // p.EventsForRange(4, 4.5, 8, ease.Linear, func(ctx ilysa.TimingContext) {
	// // 	ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
	// // 		e := ctx.NewRGBLightingEvent(
	// // 			ilysa.WithColor(grad.Ierp(ctx.T())),
	// // 		)
	// // 		fx.Ripple(ctx, e, 1)
	// // 		fx.AlphaBlend(ctx, e, 0.3, 1, 1, 0, ease.OutCirc)
	// // 	})
	// // })

	RainbowProp(p, light, grad, 4, 0.5, 1, 1)
	RainbowProp(p, lightReverse, gradient.Rainbow, 4, 0.25, 1, 1)
	RainbowProp(p, light, grad, 6, 0.25, 1, 1)
	RainbowProp(p, lightReverse, gradient.Rainbow, 6, 0.5, 1, 1)

	// save events back to Expert+ difficulty
	return p.Save()
}

func RainbowProp(p ilysa.BaseContext, light ilysa.Light, grad gradient.Table, startBeat, duration, step float64, frames int) {
	p.EventsForRange(startBeat, startBeat+duration, frames, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
			e := ctx.NewRGBLightingEvent(
				ilysa.WithColor(grad.Ierp(ctx.T())),
			)
			fx.Ripple(ctx, e, step)
			fx.AlphaBlend(ctx, e, 0.3, 1, 1, 0, ease.OutCirc)
		})
	})
}
