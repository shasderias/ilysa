// Ilysa supports gradients.
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
	"github.com/shasderias/ilysa/transform"
)

// Define a few colors.
var (
	Red   = colorful.MustParseHex("#FF0000")
	Green = colorful.MustParseHex("#00FF00")
	Blue  = colorful.MustParseHex("#0000FF")
)

// `gradient.New()` defines a gradient with the specified colors equidistant
// from each other.
var (
	RGBGrad = gradient.New(Red, Blue, Green)
	RBGrad  = gradient.New(Red, Blue)
)

// Gradients with specific positions can be defined like so. The positions
// must be sorted, the first position must be 0.0 and the last position
// must be 1.0.
var RGBSkewedGrad = gradient.Table{
	{Col: Red, Pos: 0.0},
	{Col: Green, Pos: 0.3},
	{Col: Blue, Pos: 1.0},
}

func main() {
	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
	p := ilysa.New(m)

	// Once you have a gradient, you can get the color at a position by calling
	// the `Lerp` method.
	fmt.Println(RBGrad.Lerp(0.5))

	// From here, changing the color of a light through a gradient is trivial.
	p.Range(timer.Rng(0, 1, 10, ease.Linear), func(ctx context.Context) {
		ctx.NewRGBLighting(
			evt.WithLight(evt.BackLasers),
			evt.WithColor(RGBSkewedGrad.Lerp(ctx.T())),
		)
		fmt.Println(ctx.T(), RGBSkewedGrad.Lerp(ctx.T()))
	})

	// As is displaying a gradient over a light's light IDs.
	l := transform.Light(light.NewBasic(p, evt.BackLasers),
		transform.DivideSingle(),
	)
	p.Sequence(timer.Beat(0), func(ctx context.Context) {
		ctx.Light(l, func(ctx context.LightContext) {
			ctx.NewRGBLighting(
				// `ctx.LightIDT()` is only defined within a `ctx.Light()` context,
				// it is like `ctx.T()` but for light ID sequences.
				evt.WithColor(RGBGrad.Lerp(ctx.LightIDT())),
			)
		})
	})

	p.Dump()
}
