package stdlib

import (
	"github.com/shasderias/ilysa/chroma"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func LaserSpeed(ctx context.Context, b float64, intSpeed int, lasers ...evt.Type) evt.Events {
	events := evt.NewEvents()
	ctx.WSeq(timer.Seq(b), func(ctx context.Context) {
		for _, laser := range lasers {
			events.Add(evt.NewLaserSpeed(ctx, evt.OptType(laser), evt.OptIntValue(intSpeed)))
		}
	})
	return events
}

func ChromaLaserSpeed(ctx context.Context, b float64, dir chroma.SpinDirection, floatSpeed float64, intSpeed int, lasers ...evt.Type) evt.Events {
	events := evt.NewEvents()
	ctx.WSeq(timer.Seq(b), func(ctx context.Context) {
		for _, laser := range lasers {
			events.Add(
				evt.NewChromaLaserSpeed(ctx,
					evt.OptType(laser),
					evt.OptChromaLaserSpeed(floatSpeed), evt.OptIntValue(intSpeed),
					evt.OptChromaLaserSpeedSpinDirection(dir)),
			)
		}
	})
	return events
}

func Off(ctx context.Context, b float64, types ...evt.Type) evt.Events {
	events := evt.NewEvents()
	ctx.WSeq(timer.Seq(b), func(ctx context.Context) {
		for _, typ := range types {
			events.Add(evt.NewLighting(ctx, evt.OptType(typ), evt.OptValue(evt.ValueLightOff)))
		}
	})
	return events
}

func ChromaZoom(ctx context.Context, b float64, step, speed float64) evt.Events {
	events := evt.NewEvents()
	ctx.WSeq(timer.Seq(b), func(ctx context.Context) {
		events.Add(evt.NewChromaZoom(ctx, evt.OptChromaZoomStep(step), evt.OptChromaZoomSpeed(speed)))
	})
	return events
}

func ChromaRotate(ctx context.Context, b float64, rotation, step, prop, speed float64) evt.Events {
	events := evt.NewEvents()
	ctx.WSeq(timer.Seq(b), func(ctx context.Context) {
		events.Add(
			evt.NewChromaRingRotation(ctx,
				evt.OptChromaRingRotation(rotation),
				evt.OptChromaRingRotationStep(step),
				evt.OptChromaRingRotationProp(prop),
				evt.OptChromaRingRotationSpeed(speed)),
		)
	})
	return events
}
