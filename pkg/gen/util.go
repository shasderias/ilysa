package gen

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/ilysa"
)

func OffAll(ctx ilysa.Timing) {
	var (
		lights = beatsaber.NewEventTypeSet(
			beatsaber.EventTypeBackLasers,
			beatsaber.EventTypeRingLights,
			beatsaber.EventTypeLeftRotatingLasers,
			beatsaber.EventTypeRightRotatingLasers,
			beatsaber.EventTypeCenterLights,
		)
	)

	for _, l := range lights {
		ctx.NewLightingEvent(l, beatsaber.EventValueLightOff)
	}
}
