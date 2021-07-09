package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

func ZeroSpeedRandomizedLasers(ctx context.Context, l evt.DirectionalLaser) *evt.PreciseLaser {
	return ctx.NewPreciseLaser(
		evt.WithDirectionalLaser(l),
		evt.WithIntValue(1),
		evt.WithPreciseLaserSpeed(0),
	)
}
