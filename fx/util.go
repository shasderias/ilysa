package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

// OffAll generates events to turn all lights/lasers off.
func OffAll(ctx context.Context) {
	var (
		lights = []evt.LightType{
			evt.BackLasers,
			evt.RingLights,
			evt.LeftRotatingLasers,
			evt.RightRotatingLasers,
			evt.CenterLights,
		}
	)

	for _, l := range lights {
		ctx.NewLighting(evt.WithLight(l), evt.WithLightValue(evt.LightOff))
	}
}
