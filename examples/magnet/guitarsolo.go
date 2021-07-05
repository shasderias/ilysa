package main

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/fx"
	"github.com/shasderias/ilysa/light2"
)

type GuitarSolo struct {
	p *ilysa.Project
	ilysa.BaseContext
}

func NewGuitarSolo(p *ilysa.Project, startBeat float64) GuitarSolo {
	return GuitarSolo{
		p:           p,
		BaseContext: p.WithBeatOffset(startBeat),
	}
}

func (g GuitarSolo) Play() {
	g.EventForBeat(0, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser),
			ilysa.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser),
			ilysa.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
		)
	})

	g.Beat(0)
	g.Beat(4)
	g.Beat(8)

	g.Solo(0.50, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25}, false)
	g.Solo(2.25, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50}, true)
	g.Solo(4.25, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50}, false)
	g.Solo(6.25, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50, 1.75}, true)
	g.Solo(8.50, []float64{
		0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50, 1.75,
		2.00, 2.25, 2.50, 2.75, 3.00, 3.25, 3.50,
	}, false)
	g.Solo(12.50, []float64{0.00, 0.25}, true)
	g.Solo(13.25, []float64{0.00, 0.25}, false)

}

func (g GuitarSolo) Beat(startBeat float64) {
	ctx := g.WithBeatOffset(startBeat)

	bl := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeBackLasers, g),
		ilysa.ToLightTransformer(ilysa.DivideSingle),
	)

	gradSet := gradient.NewSet(
		shirayukiGradient,
		sukoyaGradient,
		shirayukiSingleGradient,
		sukoyaSingleGradient,
	)

	ctx.EventsForBeats(0, 2, 4, func(ctx ilysa.RangeContext) {
		ctx.NewPreciseRotation(
			evt.WithRotation(15),
			evt.WithRotationStep(15),
			evt.WithProp(2),
			evt.WithPreciseLaserSpeed(8),
			evt.WithLaserDirection(chroma.CounterClockwise),
		)

		step := -0.5
		if ctx.Ordinal() %2 == 0 {
			step = 0.5
		}

		ctx.NewPreciseZoom(evt.WithRotationStep(step))

		grad := gradSet.Next()

		ctx.Range(ctx.B(), ctx.B()+0.50, 12, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(bl, func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 2.4, grad)
				fx.AlphaBlend(ctx, e, 0, 1, 1.5, 0, ease.InBounce)
			})
		})

	})
}

func (g GuitarSolo) Solo(startBeat float64, sequence []float64, reverse bool) {
	ctx := g.WithBeatOffset(startBeat)

	var (
		llReverse ilysa.LightIDTransformer = ilysa.Shuffle
		rlReverse ilysa.LightIDTransformer = ilysa.Shuffle
	)

	if reverse {
		llReverse = ilysa.Reverse
		rlReverse = ilysa.Identity
	} else {
		llReverse = ilysa.Identity
		rlReverse = ilysa.Reverse
	}

	ll := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeLeftRotatingLasers, g),
		ilysa.ToLightTransformer(llReverse),
		ilysa.ToSequenceLightTransformer(ilysa.DivideSingle),
	).(light2.SequenceLight)
	rl := light2.TransformLight(
		light2.NewBasicLight(beatsaber.EventTypeRightRotatingLasers, g),
		ilysa.ToLightTransformer(rlReverse),
		ilysa.ToSequenceLightTransformer(ilysa.DivideSingle),
	).(light2.SequenceLight)
	light := light2.NewSequenceLight()

	for i := 0; i < ll.Len(); i++ {
		light.Add(light2.NewCombinedLight(ll.Index(i), rl.Index(i)))
	}

	ctx.EventForBeat(0, func(ctx ilysa.RangeContext) {
		//ctx.NewPreciseLaser(
		//	ilysa.WithDirectionalLaser(ilysa.LeftLaser),
		//	ilysa.WithPreciseLaserSpeed(0),
		//	ilysa.WithLaserSpeed(6),
		//)
		//ctx.NewPreciseLaser(
		//	ilysa.WithDirectionalLaser(ilysa.RightLaser),
		//	ilysa.WithPreciseLaserSpeed(0),
		//	ilysa.WithLaserSpeed(6),
		//)

	})

	ctx.EventsForSequence(0, sequence, func(ctx ilysa.SequenceContext) {
		seqCtx := ctx

		var (
			llLock, rlLock                           = false, true
			llSpeed, llIntValue, rlSpeed, rlIntValue = 0.0, 5, 5.0, 5
			//llAlpha, rlAlpha = 3.0, 1.0
		)

		if reverse {
			llLock, rlLock = true, false
			llSpeed, llIntValue, rlSpeed, rlIntValue = 5.0, 5, 0.0, 5
			//llAlpha, rlAlpha = 1.0, 3.0
		}

		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.LeftLaser),
			evt.WithLockPosition(llLock),
			evt.WithPreciseLaserSpeed(llSpeed),
			ilysa.WithIntValue(llIntValue),
		)
		ctx.NewPreciseLaser(
			evt.WithDirectionalLaser(ilysa.RightLaser),
			evt.WithLockPosition(rlLock),
			evt.WithPreciseLaserSpeed(rlSpeed),
			ilysa.WithIntValue(rlIntValue),
		)

		ctx.EventsForRange(ctx.B(), ctx.B()+1.25, 30, ease.Linear, func(ctx ilysa.RangeContext) {
			ctx.WithLight(light.Index(seqCtx.Ordinal()), func(ctx ilysa.TimeLightContext) {
				e := fx.ColorSweep(ctx, 1.9, magnetRainbowPale)
				//e := ctx.NewRGBLighting(ilysa.WithColor(color))
				fx.AlphaBlend(ctx, e, 0, 1, 3, 0, ease.InCubic)
			})
		})
	})
}
