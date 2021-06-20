package fx

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
)

func OffAll(ctx ilysa.TimeContext) {
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
		ctx.NewLightingEvent(ilysa.WithType(l), ilysa.WithValue(beatsaber.EventValueLightOff))
	}
}
