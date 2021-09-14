package main

import (
	"fmt"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/timer"
	"github.com/shasderias/ilysa/transform"
)

type Sandbox struct {
	context.Context
}

func NewSandbox(p *ilysa.Project, startBeat float64) Sandbox {
	return Sandbox{p.BOffset(startBeat)}
}

func (s Sandbox) Play() {
	//s.ButterflySweep(0)
	fx.OffAll(s.BOffset(4))
	s.ButterflySweepWhole(4)
}

func (s Sandbox) ButterflySweep(startBeat float64) {
	ctx := s.BOffset(startBeat)

	transform.Light(light.NewBasic(ctx, evt.RingLights),
		transform.DivideIntoGroups(4),
		transform.TakeSet(1, 2, 5, 8, 11),
		transform.Flatten(),
		transform.DivideSingle(),
	)

	bl := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(2).Sequence(),
	).(light.Sequence)

	lWing := transform.Light(bl.Idx(0),
		transform.Reverse(),
		transform.DivideSingle().Sequence(),
	)

	rWing := transform.Light(bl.Idx(1),
		transform.DivideSingle().Sequence(),
	)

	fmt.Println(lWing, rWing)
	butterflyLight := light.Combine(lWing, rWing)

	//grad := gradient.New(sukoyaPink, sukoyaWhite, sukoyaWhite, sukoyaPink)
	grad := gradient.New(shirayukiPurple, shirayukiGold, shirayukiPurple)

	ctx.NewPreciseZoom(evt.WithZoomStep(0))
	ctx.NewPreciseRotation(
		evt.WithRotation(60),
		evt.WithRotationStep(9),
		evt.WithRotationSpeed(8),
		evt.WithProp(10),
		evt.WithRotationDirection(chroma.Clockwise),
	)

	ctx.Range(timer.Rng(0, 1, 15, ease.InOutSin), func(ctx context.Context) {
		ctx.Light(butterflyLight, func(ctx context.LightContext) {
			ctx.NewRGBLighting(evt.WithColor(grad.Lerp(ctx.T())))
		})
	})
}

func (s Sandbox) ButterflySweepWhole(startBeat float64) {
	ctx := s.BOffset(startBeat)

	bl := transform.Light(light.NewBasic(ctx, evt.BackLasers),
		transform.Fan(2),
		transform.Flatten(),
		transform.Divide(2).Sequence(),
	).(light.Sequence)

	lWing := transform.Light(bl.Idx(0),
		//transform.Reverse(),
		transform.DivideSingle().Sequence(),
	).(light.Sequence)

	rWing := transform.Light(bl.Idx(1),
		transform.DivideSingle().Sequence(),
	).(light.Sequence)

	butterflyLight := light.NewSequence()
	butterflyLight.Add(lWing.Lights()...)
	butterflyLight.Add(rWing.Lights()...)

	sukoyaWing := gradient.New(sukoyaPink, sukoyaWhite, sukoyaWhite, sukoyaPink)
	shirayukiWing := gradient.New(shirayukiPurple, shirayukiGold, shirayukiPurple)

	ctx.NewPreciseZoom(evt.WithZoomStep(0))
	ctx.NewPreciseRotation(
		evt.WithRotation(60),
		evt.WithRotationStep(12),
		evt.WithRotationSpeed(8),
		evt.WithProp(10),
		evt.WithRotationDirection(chroma.Clockwise),
	)

	ctx.Range(timer.Rng(0, 3, butterflyLight.Len(), ease.Linear), func(ctx context.Context) {
		ctx.Light(butterflyLight, func(ctx context.LightContext) {
			var grad gradient.Table
			var t float64
			if ctx.T() < 0.5 {
				grad = sukoyaWing
				t = ctx.T() * 2
			} else {
				grad = shirayukiWing
				t = (ctx.T() - 0.5) * 2
			}
			fmt.Println(ctx.T(), t)
			ctx.NewRGBLighting(evt.WithColor(grad.Lerp(t)))
		})
	})
}
