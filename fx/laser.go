package fx

import (
	"math"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/opt"
	"github.com/shasderias/ilysa/scale"
	"github.com/shasderias/ilysa/timer"
)

//func ZeroSpeedRandomizedLasers(ctx context.Context, l evt.DirectionalLaser) *evt.ChromaLaserSpeed {
//	return ctx.NewPreciseLaser(
//		evt.WithDirectionalLaser(l),
//		evt.WithIntValue(1),
//		evt.WithPreciseLaserSpeed(0),
//	)
//}
//

type slowMotionLaserConfig struct {
	lockFirst bool
}

type SlowMotionLaserOpt interface {
	apply(*slowMotionLaserConfig)
}

func OptSlowMotionLasersLockFirst() SlowMotionLaserOpt {
	return slowMotionLaserLockFirst{}
}

type slowMotionLaserLockFirst struct{}

func (o slowMotionLaserLockFirst) apply(c *slowMotionLaserConfig) {
	c.lockFirst = true
}

func SlowMotionLasers(ctx context.Context,
	rng timer.Ranger, t evt.Type, startSpeed, endSpeed float64,
	opts ...SlowMotionLaserOpt) evt.Events {
	conf := slowMotionLaserConfig{}
	for _, o := range opts {
		o.apply(&conf)
	}

	events := evt.NewEvents()

	speedScaler := scale.FromUnitClamp(startSpeed, endSpeed)

	ctx.WRng(rng, func(ctx context.Context) {
		speed := speedScaler(ctx.T())
		intSpeed := int(math.Round(speed))

		optSet := opt.NewSet(
			ctx,
			evt.OptType(t),
			evt.OptIntValue(intSpeed),
			evt.OptChromaLaserSpeed(speed),
		)

		if conf.lockFirst && ctx.First() {
			optSet.Add(evt.OptChromaLaserSpeedLockPosition(true))
		}
		if !ctx.First() {
			optSet.Add(evt.OptChromaLaserSpeedLockPosition(true))
		}
		e := evt.NewChromaLaserSpeed(optSet)
		events.Add(e)
	})

	return events
}

//
//type slowMotionLasersOpts struct {
//	preciseLaserOpts Options
//}
//
//type slowMotionLasersOpt interface {
//	apply(opt *slowMotionLasersOpts)
//}
//
//type withPreciseLaserOpt struct {
//	opts Options
//}
//
//func WithPreciseLaserOpts(opts ...Option) slowMotionLasersOpt {
//	return withPreciseLaserOpt{opts}
//}
//
//func (o withPreciseLaserOpt) apply(opt *slowMotionLasersOpts) {
//	opt.preciseLaserOpts = o.opts
//}
