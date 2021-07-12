package fx

import (
	"math"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/scale"
	"github.com/shasderias/ilysa/timer"
)

func ZeroSpeedRandomizedLasers(ctx context.Context, l evt.DirectionalLaser) *evt.PreciseLaser {
	return ctx.NewPreciseLaser(
		evt.WithDirectionalLaser(l),
		evt.WithIntValue(1),
		evt.WithPreciseLaserSpeed(0),
	)
}

func SlowMotionLasers(ctx context.Context, rng timer.Ranger, l evt.DirectionalLaser, startSpeed, endSpeed float64,
	opts ...slowMotionLasersOpt) evt.Events {

	defaultOpts := slowMotionLasersOpts{
		preciseLaserOpts: evt.Opts{},
	}
	for _, o := range opts {
		o.apply(&defaultOpts)
	}

	events := evt.NewEvents()

	speedScaler := scale.FromUnitClamp(startSpeed, endSpeed)

	ctx.Range(rng, func(ctx context.Context) {
		speed := speedScaler(ctx.T())
		intSpeed := int(math.Round(speed))

		preciseLaserOpts := evt.NewOpts(evt.WithLaserSpeed(intSpeed), evt.WithPreciseLaserSpeed(speed))
		preciseLaserOpts.Add(defaultOpts.preciseLaserOpts...)
		if !ctx.First() {
			preciseLaserOpts.Add(evt.WithLockPosition(true))
		}

		events.Add(ctx.NewPreciseLaser(evt.WithDirectionalLaser(l), preciseLaserOpts))
	})

	return events
}

type slowMotionLasersOpts struct {
	preciseLaserOpts evt.Opts
}

type slowMotionLasersOpt interface {
	apply(opt *slowMotionLasersOpts)
}

type withPreciseLaserOpt struct {
	opts evt.Opts
}

func WithPreciseLaserOpts(opts ...evt.Opt) slowMotionLasersOpt {
	return withPreciseLaserOpt{opts}
}

func (o withPreciseLaserOpt) apply(opt *slowMotionLasersOpts) {
	opt.preciseLaserOpts = o.opts
}
