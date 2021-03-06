package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/scale"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
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

	// we are only lighting beats 116 to 140 for this showcase
	const showcaseStart = 116

	// create a new context that offsets all subsequent beat numbers, this is
	// useful when creating reusable lights
	ctx := p.Offset(showcaseStart)

	// Beats 116-128 - Left/Right Lasers
	// Starting off with a relatively simple effect, here we will:
	// (1) alternate between the left and right lasers;
	// (2) smoothly change the lasers' colors through a gradient;
	// (3) increase the rotation speed of the lasers as the music approaches the drop.
	var (
		leftLaser  = light.NewBasic(ctx, evt.LeftRotatingLasers)  // base game left laser
		rightLaser = light.NewBasic(ctx, evt.RightRotatingLasers) // base game right laser

		// this creates a new Ilysa light that alternates between the left light and right lasers
		leftRightSequence = light.NewSequence(leftLaser, rightLaser)

		// this creates a gradient that blends from blue to red to purple, uncomment the 4th line to add a yellow
		// it is possible to create gradient tables with non-linear positions
		// e.g. red at 0.0, green at 0.3 and blue at 1.0
		grad = gradient.New(
			colorful.MustParseHex("#0c71c9"),             // blue
			colorful.MustParseHex("#ff145f"),             // red
			colorful.Color{R: 1.5, G: 0.6, B: 1.2, A: 1}, // PURPLE
			//colorful.MustParseHex("#fffb0d"), // yellow
		)
	)

	// generate events every half (0.5) beat, starting at beat (0), repeat a total of 24 times ...
	ctx.Sequence(timer.Interval(0, 0.5, 24), func(ctx context.Context) {
		// ... generate Chroma precise rotation speed events for the left and right lasers,
		// setting speed to the iteration count with locked positions
		// i.e. beat = 0.0, speed = 0
		//      beat = 0.5, speed = 1
		//      beat = 1.0, speed = 2, etc
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(evt.LeftLaser),
			evt.WithIntValue(ctx.Ordinal()), evt.WithPreciseLaserSpeed(float64(ctx.Ordinal())),
			evt.WithLockPosition(true),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(evt.RightLaser),
			evt.WithIntValue(ctx.Ordinal()), evt.WithPreciseLaserSpeed(float64(ctx.Ordinal())),
			evt.WithLockPosition(true),
		)

		// alphaEase is a function that will scale a number from the unit interval ([0,1]) to the interval [0.5,6]
		// we use this later to blend the alpha of the generated events from 0.5 to 6
		alphaEase := scale.FromUnitClamp(0.5, 6)

		// ... use the light we created earlier ...
		ctx.Light(leftRightSequence, func(ctx context.LightContext) {
			// ... to create a Chroma RGB event
			// WithLight automatically sets _eventType for us to alternate between left and right rotating lights
			ctx.NewRGBLighting(
				// use the gradient we created earlier to set the color
				evt.WithColor(grad.Lerp(ctx.T())),
				// use the alphaEase function we made above here, with an in-out quadratic ease
				evt.WithAlpha(alphaEase(ease.InOutQuad(ctx.T()))),
			)

			// create a Chroma event to turn off the light ...
			oe := ctx.NewRGBLighting(evt.WithLightValue(evt.LightOff))
			// ... 0.5 beats later
			oe.Apply(evt.WithBOffset(0.5))
		})
	})

	// Beats 116-128 - Big Rings
	// Here we:
	// (1) divide the big ring's lightIDs into 3 groups (in the Nice environment, [1:13], [14:26] and [27:40]);
	// (2) flicker each group in time with the rhythm's sounds;
	// (3) for each flicker, color the lightIDs in a gradient; and
	// (4) do a precision rotation of the rings, increasing rotation, speed and step as we approach the drop.
	var (
		flickerDuration = 0.35
		rhythmSeq       = []float64{
			0, 0.5, // 116
			1.0, 1.5, 1.75, // 117
			2.0, 2.25, 2.50, // 118
			3.0,                   // 119, we underlight here to give the drum rolls leading to the drop more emphasis
			4.0,                   // 120
			5.0,                   // 121
			6.0,                   // 122
			7.0,                   // 123
			8.0, 8.25, 8.50, 8.75, // 124
			9.0, 9.25, 9.50, 9.75, // 125
			10.0, 10.5, // 126
			11.0, 11.5, // 127
		}
		bigRings = light.NewBasic(ctx, evt.RingLights)
		// take base game's ring lights
		bigRingsSplit = transform.Light(bigRings,
			// and split it into 3 lights, each with 1/3 the lightIDs of the base game's ring lights
			// i.e. [1:13], [14:26], [27:40] in the Nice environment
			transform.Divide(3).Sequence(),
			// within each group, divide the lightIDs into single lightIDs so that we can light them in a gradient
			// i.e. group1: [1], [2] ... [13], group2: [14], [15] ... [26], group3: [27], [28] ... [40]
			transform.DivideSingle(),
		).(light.Sequence)

		// colors we will be using to light the ring lights
		bigRingColors = colorful.NewSet(
			colorful.MustParseHex("#34eb4f"), // lime green
			colorful.MustParseHex("#b8f3ff"), // sky blue
			colorful.MustParseHex("#f76f3e"), // orange
			colorful.MustParseHex("#f73edc"), // pink
		)
	)

	// generate events starting at beat 0, with the rhythmSeq's offsets
	ctx.Sequence(timer.Seq(rhythmSeq, 12), func(ctx context.Context) {
		// create a function that scales a number from the unit interval ([0,1] to [0.5,6])
		// we use this to set the propagation speed of the ring spins
		propScale := scale.FromUnitClamp(0.5, 5)

		// create a Chroma precise rotation event
		re := ctx.NewPreciseRotation(
			evt.WithRotation(45+float64(ctx.Ordinal())*5),         // with rotation 45, increasing by 5 with each spin
			evt.WithRotationStep(25+(float64(ctx.Ordinal())*1.5)), // with step 25, increasing by 1.5 with each spin
			evt.WithRotationSpeed(0.5+float64(ctx.Ordinal())*0.5), // with speed 0.5, increasing by 0.5 with each spin
			evt.WithProp(propScale(ctx.T())),                      // with propagation 0.5, scaling to 6 over this sequence
		)

		// for beats [1,8), rotate counterclockwise on even spins, clockwise on odd spins
		if ctx.Ordinal()%2 == 0 && ctx.B() < 8 {
			re.Apply(evt.WithRotationDirection(chroma.CounterClockwise))
		} else {
			re.Apply(evt.WithRotationDirection(chroma.Clockwise))
		}

		seqCtx := ctx
		// get the nth light, Idx() wraparounds, so this will give us ...
		// ... on the 1st iteration, big ring lights with lightIDs [1:13]
		// ... on the 2nd iteration, big ring lights with lightIDs [14:26]
		// etc el
		bigRingsSingle := bigRingsSplit.Idx(ctx.Ordinal())

		// create:
		// - 30 evenly spaced events (ease.Linear);
		// - starting from the current beat in rhythmSeq - 0.05 beats (ctx.B() - 0.05)); and (we start a little to make the lights feel more responsive)
		// - ending flickerDuration later (ctx.B() + flickerDuration - 0.05).
		ctx.Range(timer.Rng(-0.05, flickerDuration-0.05, 30, ease.Linear), func(ctx context.Context) {
			// use the light we picked out
			ctx.Light(bigRingsSingle, func(ctx context.LightContext) {
				// generate a gradient from the color set we selected
				// i.e. on the 1st iteration, lime green to sky blue
				//      on the 2nd iteration, sky blue to orange
				// etc el
				grad := gradient.New(
					bigRingColors.Idx(seqCtx.Ordinal()),
					bigRingColors.Idx(seqCtx.Ordinal()+1),
				)

				// apply the gradient, fx.Gradient will generate the requisite events based on the light we are using and the gradient passed to it
				e := fx.Gradient(ctx, grad)
				// set the alpha of the generated events to 15
				e.Apply(evt.WithAlpha(15))
				// apply a ripple effect (stagger the starting time of each lightID), with 0.10 beats between each successive lightID
				fx.RippleT(ctx, e, flickerDuration)
				// and apply an alpha fade from 1 to 0, starting halfway (0.5) through the sequence, with the OutCirc easing
				fx.AlphaFadeEx(ctx, e, 0.5, 1.0, 15, 0, ease.OutCirc)
			})
		})
	})

	// Beats 116-128 - Center Lights/Back Lights
	// Here we:
	// - do a zoom every 4 beats;
	// - animate a (synced!) rainbow gradient over the center lights and the back lights;
	// - fade the center and back lights out.
	// This effect is rather subtle in the Nice environment due to the limited number of lightIDs. We do it anyways
	// as Ilysa takes into account the number of lightIDs available in the selected environment when generating
	// events, and this lets gets us a whole new lightshow just by changing the environment.
	var (
		centerLights = transform.Light(
			light.NewBasic(ctx, evt.CenterLights), // take the base game's center lights
			transform.DivideSingle(),              // divide into single lightIDs
		)
		backLights = transform.Light(
			light.NewBasic(ctx, evt.BackLasers), // repeat for back lasers
			transform.DivideSingle(),
		)
		combinedLights = light.Combine(centerLights, backLights) // combine them
	)

	// this is similar to the pattern we used for the previous effect, see above for commentary
	ctx.Sequence(timer.Interval(0, 4, 3), func(ctx context.Context) {
		ctx.NewZoom() // base game zoom event
		ctx.Range(timer.Rng(0, 3.9, 60, ease.Linear), func(ctx context.Context) {
			ctx.Light(combinedLights, func(ctx context.LightContext) {
				// ColorSweep is an effect that comes with Ilysa that animates a gradient moving over a set of
				// lightIDs. The "speed" of the animation is controllable using the 2nd argument (1.4 in this case).
				e := fx.ColorSweep(ctx, 1.4, gradient.Rainbow)
				fx.AlphaFadeEx(ctx, e, 0.3, 1, 1, 0, ease.OutCirc)
			})
		})
	})

	// Beats 128-140 - Drop
	const (
		dropOffset = 12
		dropLength = 12
	)

	// once the drop lands
	ctx.Sequence(timer.Beat(dropOffset), func(ctx context.Context) {
		ctx.NewPreciseRotation( // do a precision rotation event
			evt.WithRotation(720),
			evt.WithRotationStep(17),
			evt.WithProp(0.5),
			evt.WithRotationDirection(chroma.CounterClockwise),
			evt.WithRotationSpeed(3),
		)
		ctx.NewLaser( // slow down the left and right lasers
			evt.WithDirectionalLaser(evt.LeftLaser),
			evt.WithIntValue(1),
		)
		ctx.NewLaser(
			evt.WithDirectionalLaser(evt.RightLaser),
			evt.WithIntValue(1),
		)
	})

	// Beats 128-140 - Big Rings
	// This takes the ColorSweep effect introduced earlier, applies it to the whole big ring and adds a
	// shimmery effect to it.
	var (
		bigRingsWhole = transform.Light( // here we take the ring lights as a whole ...
			bigRings,
			transform.DivideSingle(), // .. and divide the lightIDs into individual units
		)
	)

	// over the length of the drop
	ctx.Range(timer.Rng(dropOffset, dropOffset+dropLength, 120, ease.Linear), func(ctx context.Context) {
		ctx.Light(bigRingsWhole, func(ctx context.LightContext) {
			// animate a gradient moving over the ring lasers
			e := fx.ColorSweep(ctx, 1.2, gradient.Rainbow)
			// add a shimmer effect by setting the alpha values of each lightID based on 1d-noise generated
			// with a bunch of sine functions
			fx.AlphaShimmer(ctx, e, 3)
			// fade to black
			fx.AlphaFadeEx(ctx, e, 0, 1, 2.5, 0, ease.InCirc)
		})
	})

	// Beats 128-136 - Left/Right Lasers
	// Reuse of the concepts introduced earlier to alternate between the left and right rotating lasers, with each
	// laser being it in a gradient with a ripple effect. The higher step value for the ripple changes the feel of
	// the effect to be less like a ripple and more like the lasers lighting up in random order.
	var (
		leftRightSequenceSplit = transform.Light(
			leftRightSequence,
			transform.DivideSingle(),
		).(light.Sequence)
		dropColors = colorful.NewSet(
			colorful.MustParseHex("#3775bd"), // shades of blue
			colorful.MustParseHex("#add4ed"),
			colorful.MustParseHex("#b0aded"),
		)
	)

	ctx.Sequence(timer.Interval(dropOffset, 1, 8), func(ctx context.Context) {
		light := leftRightSequenceSplit.Idx(ctx.Ordinal())
		seqCtx := ctx
		ctx.Range(timer.Rng(0, 0.7, 30, ease.Linear), func(ctx context.Context) {
			grad := gradient.New(
				dropColors.Idx(seqCtx.Ordinal()),
				dropColors.Idx(seqCtx.Ordinal()+2),
			)
			ctx.Light(light, func(ctx context.LightContext) {
				e := fx.Gradient(ctx, grad)
				fx.RippleT(ctx, e, 1.5)
				fx.AlphaFadeEx(ctx, e, 0, 0.3, 0, 1, ease.OutSin)
				fx.AlphaFadeEx(ctx, e, 0.3, 1, 1, 0, ease.InSin)
			})
		})
	})

	// save events back to Expert+ difficulty
	return p.Save()
}
