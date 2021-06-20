package fx

import (
	"github.com/shasderias/ilysa/pkg/beatsaber"
	"github.com/shasderias/ilysa/pkg/ilysa"
)

func OffAll(ctx ilysa.TimingContext) {
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
