# Ilysa

Ilysa is a Go library that helps you create your favorite lighting patterns with speed and ease(probably)!

## Sell Me

Ilysa lets you generate lightshows like this:

[![Ilysa Showcase](https://img.youtube.com/vi/MohFQiz8tAU/0.jpg)](https://www.youtube.com/watch?v=MohFQiz8tAU)

<details>
  <summary>from this</summary>

```go
package main

import (
	"fmt"

	"github.com/shasderias/ilysa/pkg/beatsaber"
	"github.com/shasderias/ilysa/pkg/chroma"
	"github.com/shasderias/ilysa/pkg/colorful"
	"github.com/shasderias/ilysa/pkg/colorful/gradient"
	"github.com/shasderias/ilysa/pkg/ease"
	"github.com/shasderias/ilysa/pkg/ilysa"
	"github.com/shasderias/ilysa/pkg/ilysa/fx"
	"github.com/shasderias/ilysa/pkg/util"
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
	ctx := p.WithBeatOffset(showcaseStart)

	// Beats 116-128 - Left/Right Lasers
	// Starting off with a relatively simple effect, here we will:
	// (1) alternate between the left and right lasers;
	// (2) smoothly change the lasers' colors through a gradient;
	// (3) increase the rotation speed of the lasers as the music approaches the drop.
	var (
		leftLaser  = ilysa.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, p)  // base game left laser
		rightLaser = ilysa.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, p) // base game right laser

		// this creates a new Ilysa light that alternates between the left light and right lasers
		leftRightSequence = ilysa.NewSequenceLight(leftLaser, rightLaser)

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
	ctx.EventsForBeats(0, 0.5, 24, func(ctx ilysa.TimingContext) {
		// ... generate Chroma precise rotation speed events for the left and right lasers,
		// setting speed to the iteration count with locked positions
		// i.e. beat = 0.0, speed = 0
		//      beat = 0.5, speed = 1
		//      beat = 1.0, speed = 2, etc
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(ctx.Ordinal()), ilysa.WithSpeed(float64(ctx.Ordinal())),
			ilysa.WithLockPosition(true),
		)
		ctx.NewPreciseRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(ctx.Ordinal()), ilysa.WithSpeed(float64(ctx.Ordinal())),
			ilysa.WithLockPosition(true),
		)

		// alphaEase is a function that will scale a number from the unit interval ([0,1]) to the interval [0.5,6]
		// we use this later to blend the alpha of the generated events from 0.5 to 6
		alphaEase := util.ScaleFromUnitInterval(0.5, 6)

		// ... use the light we created earlier ...
		ctx.UseLight(leftRightSequence, func(ctx ilysa.TimingContextWithLight) {
			// ... to create a Chroma RGB event
			// UseLight automatically sets _eventType for us to alternate between left and right rotating lights
			ctx.NewRGBLightingEvent(
				// use the gradient we created earlier to set the color
				ilysa.WithColor(grad.Ierp(ctx.T())),
				// use the alphaEase function we made above here, with an in-out quadratic ease
				ilysa.WithAlpha(alphaEase(ease.InOutQuad(ctx.T()))),
			)

			// create a Chroma event to turn off the light ...
			oe := ctx.NewRGBLightingEvent(ilysa.WithValue(beatsaber.EventValueLightOff))
			// ... 0.5 beats later
			oe.ShiftBeat(0.5)
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
		bigRings = ilysa.NewBasicLight(beatsaber.EventTypeRingLights, p)
		// take base game's ring lights
		bigRingsSplit = ilysa.TransformLight(bigRings,
			// and split it into 3 lights, each with 1/3 the lightIDs of the base game's ring lights
			// i.e. [1:13], [14:26], [27:40] in the Nice environment
			ilysa.ToSequenceLightTransformer(ilysa.Divide(3)),
			// within each group, divide the lightIDs into single lightIDs so that we can light them in a gradient
			// i.e. group1: [1], [2] ... [13], group2: [14], [15] ... [26], group3: [27], [28] ... [40]
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		).(ilysa.SequenceLight)
		// colors we will be using to light the ring lights
		bigRingColors = colorful.NewSet(
			colorful.MustParseHex("#34eb4f"), // lime green
			colorful.MustParseHex("#b8f3ff"), // sky blue
			colorful.MustParseHex("#f76f3e"), // orange
			colorful.MustParseHex("#f73edc"), // pink
		)
	)

	// generate events starting at beat 0, with the rhythmSeq's offsets
	ctx.EventsForSequence(0, rhythmSeq, func(ctx ilysa.SequenceContext) {
		// create a function that scales a number from the unit interval ([0,1] to [0.5,6])
		// we use this to set the propagation speed of the ring spins
		propScale := util.ScaleFromUnitInterval(0.5, 5)

		// create a Chroma precise rotation event
		re := ctx.NewPreciseRotationEvent(
			ilysa.WithRotation(45+float64(ctx.Ordinal())*5), // with rotation 45, increasing by 5 with each spin
			ilysa.WithStep(25+(float64(ctx.Ordinal())*1.5)), // with step 25, increasing by 1.5 with each spin
			ilysa.WithSpeed(0.5+float64(ctx.Ordinal())*0.5), // with speed 0.5, increasing by 0.5 with each spin
			ilysa.WithProp(propScale(ctx.T())),              // with propagation 0.5, scaling to 6 over this sequence
		)

		// for beats [1,8), rotate counterclockwise on even spins, clockwise on odd spins
		if ctx.Ordinal()%2 == 0 && ctx.B() < 8 {
			re.Mod(ilysa.WithDirection(chroma.CounterClockwise))
		} else {
			re.Mod(ilysa.WithDirection(chroma.Clockwise))
		}

		seqCtx := ctx
		// get the nth light, Index() wraparounds, so this will give us ...
		// ... on the 1st iteration, big ring lights with lightIDs [1:13]
		// ... on the 2nd iteration, big ring lights with lightIDs [14:26]
		// etc el
		light := bigRingsSplit.Index(ctx.Ordinal())

		// create:
		// - 30 evenly spaced events (ease.Linear);
		// - starting from the current beat in rhythmSeq - 0.05 beats (ctx.B() - 0.05)); and (we start a little to make the lights feel more responsive)
		// - ending flickerDuration later (ctx.B() + flickerDuration - 0.05).
		ctx.EventsForRange(ctx.B()-0.05, ctx.B()+flickerDuration-0.05, 30, ease.Linear, func(ctx ilysa.TimingContext) {
			// use the light we picked out
			ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
				// generate a gradient from the color set we selected
				// i.e. on the 1st iteration, lime green to sky blue
				//      on the 2nd iteration, sky blue to orange
				// etc el
				grad := gradient.New(
					bigRingColors.Index(seqCtx.Ordinal()),
					bigRingColors.Index(seqCtx.Ordinal()+1),
				)

				// apply the gradient, fx.Gradient will generate the requisite events based on the light we are using and the gradient passed to it
				e := fx.Gradient(ctx, grad)
				// set the alpha of the generated events to 15
				e.SetAlpha(15)
				// apply a ripple effect (stagger the starting time of each lightID), with 0.10 beats between each successive lightID
				fx.Ripple(ctx, e, 0.10,
					// and apply an alpha fade from 1 to 0, starting halfway (0.5) through the sequence, with the OutCirc easing
					fx.WithAlphaBlend(0.5, 1.0, 1, 0, ease.OutCirc))
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
		centerLights = ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeCenterLights, p), // take the base game's center lights
			ilysa.ToLightTransformer(ilysa.DivideSingle),            // divide into single lightIDs
		)
		backLights = ilysa.TransformLight(
			ilysa.NewBasicLight(beatsaber.EventTypeBackLasers, p), // repeat for back lasers
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		)
		combinedLights = ilysa.NewCombinedLight(centerLights, backLights) // combine them
	)

	// this is similar to the pattern we used for the previous effect, see above for commentary
	ctx.EventsForBeats(0, 4, 3, func(ctx ilysa.TimingContext) {
		ctx.NewZoomEvent() // base game zoom event
		ctx.EventsForRange(ctx.B(), ctx.B()+3.9, 60, ease.Linear, func(ctx ilysa.TimingContext) {
			ctx.UseLight(combinedLights, func(ctx ilysa.TimingContextWithLight) {
				// ColorSweep is an effect that comes with Ilysa that animates a gradient moving over a set of
				// lightIDs. The "speed" of the animation is controllable using the 2nd argument (1.4 in this case).
				fx.ColorSweep(ctx, 1.4, gradient.Rainbow,
					fx.WithAlphaBlend(0.3, 1, 1, 0, ease.OutCirc),
				)
			})
		})
	})

	// Beats 128-140 - Drop
	const (
		dropOffset = 12
		dropLength = 12
	)

	// once the drop lands
	ctx.EventForBeat(dropOffset, func(ctx ilysa.TimingContext) {
		ctx.NewPreciseRotationEvent( // do a precision rotation event
			ilysa.WithRotation(720),
			ilysa.WithStep(17),
			ilysa.WithProp(0.5),
			ilysa.WithDirection(chroma.CounterClockwise),
			ilysa.WithSpeed(3),
		)
		ctx.NewRotationSpeedEvent( // slow down the left and right lasers
			ilysa.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(1),
		)
		ctx.NewRotationSpeedEvent(
			ilysa.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(1),
		)
	})

	// Beats 128-140 - Big Rings
	// This takes the ColorSweep effect introduced earlier, applies it to the whole big ring and adds a
	// shimmery effect to it.
	var (
		bigRingsWhole = ilysa.TransformLight( // here we take the ring lights as a whole ...
			bigRings,
			ilysa.ToLightTransformer(ilysa.DivideSingle), // .. and divide the lightIDs into individual units
		)
	)

	// over the length of the drop
	ctx.EventsForRange(dropOffset, dropOffset+dropLength, 120, ease.Linear, func(ctx ilysa.TimingContext) {
		ctx.UseLight(bigRingsWhole, func(ctx ilysa.TimingContextWithLight) {
			// animate a gradient moving over the ring lasers
			e := fx.ColorSweep(ctx, 0.6, gradient.Rainbow)
			// add a shimmer effect by setting the alpha values of each lightID based on 1d-noise generated
			// with a bunch of sine functions
			fx.AlphaShimmer(ctx, e, 3)
			// fade to black
			fx.AlphaBlend(ctx, e, 0.6, 1, 1, 0, ease.OutSine)
		})
	})

	// Beats 128-136 - Left/Right Lasers
	// Reuse of the concepts introduced earlier to alternate between the left and right rotating lasers, with each
	// laser being it in a gradient with a ripple effect. The higher step value for the ripple changes the feel of
	// the effect to be less like a ripple and more like the lasers lighting up in random order.
	var (
		leftRightSequenceSplit = ilysa.TransformLight(
			leftRightSequence,
			ilysa.ToLightTransformer(ilysa.DivideSingle),
		).(ilysa.SequenceLight)
		dropColors = colorful.NewSet(
			colorful.MustParseHex("#3775bd"), // shades of blue
			colorful.MustParseHex("#add4ed"),
			colorful.MustParseHex("#b0aded"),
		)
	)

	ctx.EventsForBeats(dropOffset, 1, 8, func(ctx ilysa.TimingContext) {
		light := leftRightSequenceSplit.Index(ctx.Ordinal())
		seqCtx := ctx
		ctx.EventsForRange(ctx.B(), ctx.B()+0.75, 30, ease.Linear, func(ctx ilysa.TimingContext) {
			grad := gradient.New(
				dropColors.Index(seqCtx.Ordinal()),
				dropColors.Index(seqCtx.Ordinal()+2),
			)
			ctx.UseLight(light, func(ctx ilysa.TimingContextWithLight) {
				e := fx.Gradient(ctx, grad)
				fx.Ripple(ctx, e, 1.2,
					fx.WithAlphaBlend(0, 0.3, 0, 1, ease.InSine),
					fx.WithAlphaBlend(0.3, 1, 1, 0, ease.OutSine),
				)
			})
		})
	})

	// save events back to Expert+ difficulty
	return p.Save()
}
```

</details>

## Is Ilysa for me?

Beat Saber lighting knowledge and intermediate computer skills required. Ilysa is not for you if:

- you have never placed a Chroma event in ChroMapper;
- ```_eventType``` and ```_eventValue``` don't mean anything to you; or
- a command prompt scares you.

## Ilysa ...

* may not give better results than hand lighting, it only makes actualizing complicated effects easy;
* may not be easier than handlighting (the converse is probably true);
* is not an autolighter - results directly proportionate to user's skill at lighting and technical art;
* does not generate any other beatmap elements (for walls, you probably want
  spookyGh0st's [Beatwalls](https://github.com/spookyGh0st/beatwalls#readme])).

## Do I need to know Go? Programming?

Ilysa was designed to be somewhat usable by a non-programmer (actual results may vary). You can probably get somewhere
just by copy/pasting code and tweaking values.

If you want to achieve novel effects, you'll need at least a rudimentary understanding of Go.

If you can already program in another language, Go should be a snap. Take a stroll
through [A Tour of Go](https://tour.golang.org/welcome/1) and carry on.

If you have never programmed, Go is an easy language to learn.
Completing [A Tour of Go](https://tour.golang.org/welcome/1)
up to the section on Methods should give you enough understanding of Go.

# Getting Started

## Requirements

* a working Go environment
* a working Git installation
* a code editor (these instructions are tested with Visual Studio Code and the Ilysa's author uses Goland)
* a beatmap with *all* requisite BPM blocks placed

### Go Environment

Follow the instructions at:

* [Download and install](https://golang.org/doc/install)
* [Tutorial: Get started with Go](https://golang.org/doc/tutorial/getting-started)

If successfully complete all the instructions
at [Tutorial: Get started with Go](https://golang.org/doc/tutorial/getting-started), you have a working Go environment.

### Git Installation

If you are on:

* Windows - download and install from https://git-scm.com/downloads
* macOS - run `git` from Terminal to start the installer
* Linux - you don't need me

### Code Editor

#### Visual Studio Code

Install the Go extension from the Extensions menu and follow its Quick Start instructions. Be sure to install the
command line tools when prompted to do so for code assistance support.

#### Goland

You know what you're doing.

### Beatmap

Ilysa works in BPM adjusted beats, i.e. the beat numbers displayed in MMA2 or ChroMapper. If you do not place all the
required BPM blocks before starting, and add BPM blocks after writing Ilysa code, you will probably have to retime your
code.

# Walkthrough

**⚠️ Ilysa will replace all lighting events in the selected difficulty. Please dedicate a copy of your map for use with
Ilysa and make backups (you should be making backups regardless)!⚠️**

## Preliminaries

Create a new directory to hold your Ilysa project, then in that directory, execute the following commands:

Initialize the directory as the root of a Go project.

```
go mod init projectName
```

Download Ilysa. (TODO: Include instructions for alpha testers to configure Github account.)

```
go get -u github.com/shasderias/ilysa
```

## Boilerplate

Experienced Go programmer? Ilysa imposes no structure on your code, copy and paste the lines in the `do()` function and
you're off to the races.

New to Go? Create a main.go in your project directory and copy and paste the following:

main.go (from ```examples/getting-started```)

```go
package main

import (
	"fmt"

	"github.com/shasderias/ilysa/pkg/beatsaber"
	"github.com/shasderias/ilysa/pkg/ilysa"
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

	// save events back to Expert+ difficulty
	return p.Save()
}
```

Compile and run your code by executing:

```shell
go run .
```

This will remove all lighting events from your map and add all events generated by Ilysa to it. As the code above does
not generate any events, it will simply remove all existing lighting events.

## Your First Ilysa Event

Add the following lines to your program:

```go
package main

// snip
func do() error {
	// snip
	// -- your code goes here --
	p.EventForBeat(2, func(ctx ilysa.TimingContext) { // generate events for beat 2:
		ctx.NewLightingEvent( // generate a new base game (non-Chroma) lighting event
			ilysa.WithType(beatsaber.EventTypeBackLasers),   // back lasers
			ilysa.WithValue(beatsaber.EventValueLightRedOn), // red on
		)
	})
	// snip
}
```

The value returned by ```ilysa.New()``` (```p``` in the snippet above) represents an Ilysa project and is your entry
point for working with the library.

There are a few methods defined on ```p``` that can be used to generate events, of which ```EventForBeat``` is the
simplest. ```EventForBeat``` accepts two arguments, a beat number (2 in the example above) and a callback function. The
callback function has one argument, a context value. You call methods on this context value to generate lighting events.

The signature of the callback function changes based on the method. So it is easiest to let your editor's code
assistance do the work. In Visual Studio Code, type `p.EventForBeat(2, `, hit `Ctrl-Space`, select the first option and
the editor will fill in the rest.

The method ```NewLightingEvent``` generates a base game lighting event. It accepts 0 or more functional arguments. In
the example above:

* passing ```ilysa.WithType(beatsaber.EventTypeBackLasers)``` to ```NewLightingEvent``` tells Ilysa to create an event
  that controls the lights in the Back Lasers group; and
* passing ```ilysa.WithValue(beatsaber.EventValueLightRedOn)``` to ```NewLightingEvent``` tells Ilysa to create an event
  that changes the back laser lights to red and turns them on.

For generating base game events, ```ctx``` has the following methods:

```go
ctx.NewRotationEvent()
// accepts ilysa.WithValue() for controlling the hydraulics in the Interscope environment

ctx.NewZoomEvent()
// does not accept any arguments

ctx.NewRotationSpeedEvent()
// accepts:
// - ilysa.WithDirectionalLaser() for selecting which laser to control
// e.g. ilysa.WithDirectionalLaser(ilysa.LeftLaser)
// - ilysa.WithIntValue() for setting the rotation speed of the laser selected
// e.g. ilysa.WithIntValue(3)
```

See ```examples/some-basic-examples```  for a few more basic examples.

## Something a Little More Fancy

```p```  has other methods that can be used to generate events. Let's take a look at
```EventsForBeats```:

```go
// generate events every quarter beat (0.25), starting at beat 3, do this a total of 16 times ...
// i.e. 3.00, 3.25, 3.50, 3.75, 4.00 ... 6.50, 6.75
p.EventsForBeats(3, 0.25, 16, func(ctx ilysa.TimingContext) {  
    // ... each time, generate a rotation speed event ...
    ctx.NewRotationSpeedEvent(
        // ... that controls the left laser's rotation speed
        ilysa.WithDirectionalLaser(ilysa.LeftLaser),
        // ctx.Ordinal() returns the iteration number, starting at 0
        // i.e. for beat 3.00, ctx.Ordinal is 0, for beat 3.25, ctx.Ordinal is 1
        // the following line will therefore increase the left laser's rotation speed from 0 to 15 over 3.75 beats
        ilysa.WithValue(beatsaber.EventValue(ctx.Ordinal())),
    )
})
```

### Context

```ctx``` has a number of methods that return values useful for varying events properties to time:

```go
B() float64         // current beat
T() float64         // current time in the current sequence, on a 0-1 scale
Ordinal() int       // ordinal number of the current iteration, starting from 0
StartBeat() float64 // first beat of the current sequence
EndBeat() float64   // last beat of the current sequence
Duration() float64  // duration of the current sequence, in beats
First() bool        // true if this is the first iteration
Last() bool         // true if this is the last iteration
FixedRand() float64 // a number from 0-1, fixed for the current sequence, but different for every sequence
```

In some cases, `ctx` may have additional methods relevant to the current scope.

In the above example, we used ```ctx.Ordinal()``` to vary the left laser's rotation speed with time.

Finally, to create events that recur at specific beats, we have ```EventsForSequence``` which accepts a sequence of
beats and creates events based on that sequence.

```go
// generate events on beats 0, 0.25, 0.75 and 1.25, starting from beat 4
// i.e. 4.00, 4.25, 4.75, 5.25
p.EventsForSequence(4, []float64{0, 0.25, 0.75, 1.25}, func(ctx ilysa.SequenceContext) {
    ctx.NewLightingEvent(
        ilysa.WithType(beatsaber.EventTypeRingLights),
        ilysa.WithValue(beatsaber.EventValueLightBlueFade),
    )
})
```

The above snippet turns the ring lights on and fades them to black at beats 4.00, 4.25, 4.75 and 5.25.

# Gimme cut and paste! I want to ...

## ... generate events ...

<details>
  <summary>... on a specific beat!</summary>
Use

```go
EventForBeat(beat float64, callback func(TimingContext))
```

e.g. beat 24.5

```go
p.EventForBeat(24.5, func(ctx ilysa.TimingContext) {
  // use ctx to generate events here
})
```

</details>


<details>
  <summary>... at regular beats!</summary>

Use

```go
EventsForBeats(startBeat, duration float64, count int, callback func(TimingContext))
```

e.g. beats 0, 4, 8, 12, 16

```go
p.EventsForBeats(0, 4, 5, func(ctx ilysa.TimingContext) {
  // use ctx to generate events here
})
```

</details>

<details>
  <summary>... whenever I want to!</summary>

Use

```go
EventsForSequence(startBeat float64, sequence []float64, callback func(ctx SequenceContext))
```

e.g. beats 2.0, 2.25, 2.75, 3.0, 3.75

```go
p.EventsForSequence(0, []float64{2,2.25,2.75,3.0,3.75}, func(ctx ilysa.SequenceContext) {
  // use ctx to generate events here
})
```

</details>

<details>
  <summary>..., lots of events (in a range)!</summary>

Use

```go
EventsForRange(startBeat, endBeat float64, steps int, easeFunc ease.Func, callback func(TimingContext))
```

e.g. beats 2.0, 2.1, 2.2 ... 3.0

```go
p.EventsForRange(2, 3, 11, ease.Linear, func(tc ilysa.TimingContext) {
   // use ctx to generate events here
})
```

</details>

## ... generate the base game event for ...

<details>
  <summary>... lights!</summary>

Use

```go
NewLightingEvent(opts ...BasicLightingEventOpt) *BasicLightingEvent 
```

e.g. beat 2, back lasers, blue flash

```go
p.EventForBeat(2, func(ctx ilysa.TimingContext) {
  ctx.NewLightingEvent(
    ilysa.WithType(beatsaber.EventTypeBackLasers),
    ilysa.WithValue(beatsaber.EventValueLightBlueFlash),
  )
})
```

</details>

<details>
  <summary>... ring spins!</summary>

Use

```go
NewLightingEvent(opts ...BasicLightingEventOpt) *BasicLightingEvent 
```

e.g. beat 2, spin

```go
p.EventForBeat(2, func(ctx ilysa.TimingContext) {
  ctx.NewRotationEvent()
}
```

e.g. beat 2, raise hydraulics, all cars

```go
p.EventForBeat(2, func(ctx ilysa.TimingContext) {
  ctx.NewRotationEvent(
    ilysa.WithType(beatsaber.EventTypeInterscopeRaiseHydraulics),
    ilysa.WithValue(1), // TODO: add Interscope environment enums
  )
}
```

</details>

<details>
  <summary>... ring zooms!</summary>

Use

```go
NewZoomEvent()
```

e.g. beat 2, zoom pls

```go
p.EventForBeat(2, func(ctx ilysa.TimingContext) {
  ctx.NewZoomEvent()
}
```

</details>

<details>
  <summary>... light rotation speed!</summary>

Use

```go
NewRotationSpeedEvent(opts ...RotationSpeedEventOpt) *RotationSpeedEvent
```

e.g. beat 2, left laser, zooooooooom

```go
p.EventForBeat(2, func(ctx ilysa.TimingContext) {
  ctx.NewRotationSpeedEvent(
    ilysa.WithDirectionalLaser(ilysa.LeftLaser), 
    ilysa.WithIntValue(50),
  )
}
```

</details>

## ... generate the Chroma event ...

<details>
  <summary>... RGB lights!</summary>

Use

```go
NewRGBLightingEvent(options...)
```

e.g. fully loaded

```go
ctx.NewRGBLightingEvent(
    ilysa.WithType(beatsaber.EventTypeBackLasers),
    ilysa.WithValue(beatsaber.EventValueLightRedOn),
    ilysa.WithColor(colorful.MustParseHex("#123123")),
    ilysa.WithAlpha(0.3),
    ilysa.WithLightID(ilysa.NewLightID(1, 2, 3)),
)
```

</details>

<details>
  <summary>... precise laser!</summary>

Use

```go
NewPreciseRotationSpeedEvent(options...)
```

e.g. fully loaded

```go
ctx.NewPreciseRotationSpeedEvent(
    ilysa.WithLockPosition(true),
    ilysa.WithIntValue(1),
    ilysa.WithSpeed(0),
    ilysa.WithDirection(chroma.Clockwise),
)
```

</details>

<details>
<summary>... precise rotation!</summary>

Use

```go
NewPreciseRotationEvent(opts ...PreciseRotationEventOpt) *PreciseRotationEvent 
```

e.g. fully loaded

```go

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
```

</details>

<details>
<summary>... precise zoom!</summary>

Use

```go
NewPreciseZoomEvent(opts ...PreciseZoomEventOpt) *PreciseZoomEvent
```

e.g. fully loaded

```go
ctx.NewPreciseZoomEvent(
    ilysa.WithStep(4),
)
```

</details>

## ... generate something fancy!

<details>
  <summary>Rainbow Prop?</summary>

Wat dis? Light that runs down a sequence of lightIDs, changing color as it moves.

*- video goes here -*

```go
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
```

</details>

# Lights

*TODO: Sob. Somebody write this for me.*

# Conventions

By convention:

* T or t - is always in the range [0:1]
* B - always refers to BPM scaled beats
* LightID or lightID - is always 1-indexed
* Ordinal - is always 0-indexed

# Utility Packages

## Ease

The `ease` package implements all the easing functions at https://easings.net/.

## Util

The `util` package implements a few functions handy for scaling numbers:

```go
// returns a function that scales a number from [rMin,rMax] to [tMin,tMax]
func Scale(rMin, rMax, tMin, tMax float64) func(m float64) float64
// returns a function that scales a number from [0,1] to [tMin,tMax]
func ScaleFromUnitInterval(tMin, tMax float64) func(m float64) float64 
// returns a function that scales a number from [rMin,rMax] to [0,1]
func ScaleToUnitInterval(rMin, rMax float64) func(m float64) float64 
```

## Colors

The `colorful` package (adapted from [go-colorful](https://github.com/lucasb-eyer/go-colorful)) implements a bunch of
functions handy for working with colors.

You will mostly be working with:

### Defining Colors

```go
// hopefully self-explanatory
black := colorful.MustParseHex("#000000")
white := colorful.Color{R: 1, G: 1, B: 1, A: 1}
```

### Sets

You can put colors in sets:

```go
set := colorful.NewSet(
  colorful.MustParseHex("#fbc6d0"),
  colorful.MustParseHex("#95bddc"),
  colorful.MustParseHex("#3a2b1c"),
  colorful.MustParseHex("#451234"),
)
```

This makes it convenient to do several things.

Index. Useful when combined with `ctx.Ordinal()`.

```go
set.Index(0) // returns the 1st color in the set
set.Index(3) // returns the 4th color in the set
set.Index(4) // returns the 1st color in the set (wraparound)

// try using this in a p.EventsForSequence() to cycle through the colors in the set
set.Index(ctx.Ordinal())
```

Iterate. Useful when `ctx.Ordinal()` doesn't have sufficient range to cycle through all the colors.

```go

// returns the next color in the set, starting with the 1st one
// the set maintains internal state keeping track of the last color returned
set.Next() 
```

Random.

```go
set.Rand() // returns a random color from the set
```

## Gradients

You can quickly define a gradient with all the colors equidistant from each other by using `gradient.New()`.

```go
grad := gradient.New(
  colorful.MustParseHex("#fbc6d0"),
  colorful.MustParseHex("#95bddc"),
  colorful.MustParseHex("#0c71c9"),
  colorful.MustParseHex("#ff145f"),
)
```

Gradients with custom color positions are defined like so.

```go
// Pos MUST be sorted and range from 0.0 to 1.0
grad := gradient.Table{
  {Col: colorful.MustParseHex("#fbc6d0"), Pos: 0.0},
  {Col: colorful.MustParseHex("#95bddc"), Pos: 0.2},
  {Col: colorful.MustParseHex("#0c71c9"), Pos: 0.8},
  {Col: colorful.MustParseHex("#ff145f"), Pos: 1.0},
}
```

Once you have a gradient, get the color at position `t` by calling `Ierp()`.

```go
grad := gradient.Table{
  {Col: colorful.MustParseHex("#fbc6d0"), Pos: 0.0},
  {Col: colorful.MustParseHex("#95bddc"), Pos: 0.2},
  {Col: colorful.MustParseHex("#0c71c9"), Pos: 0.8},
  {Col: colorful.MustParseHex("#ff145f"), Pos: 1.0},
}

grad.Ierp(0.35)

// most commonly used with ctx.T()
grad.Ierp(ctx.T())
```

# Tips and Tricks

## Visual Studio Code Keyboard Shortcuts

* Ctrl-Shift-Space - display function arguments
* Ctrl-Space - autocomplete

# Resources

* https://www.desmos.com/calculator

# Credits

* Alice (Alice#5792) for lighting advice and being a sounding board. Check her out
  on [Twitch](https://www.twitch.tv/alicexiv)!
* Lucas Beyer, Bastien Dejean (@baskerville), Phil Kulak (@pkulak), Christian Muehlhaeuser (@muesli), makeworld (
  @makeworld-the-better-one) and the other contributors to the [go-colorful](https://github.com/lucasb-eyer/go-colorful)
  library, used under the MIT License.
* Pennock Tech, LLC for the [swallowjson](https://github.com/PennockTech/swallowjson) library, used under the MIT
  License.
* Top_Cat (Top_Cat#1961) for the Beat Saber environment definition files at https://github.com/Top-Cat/bs-env.