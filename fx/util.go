package fx

import (
	ilysa2 "github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber")

func OffAll(ctx ilysa2.TimeContext) {
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
		ctx.NewLightingEvent(ilysa2.WithType(l), ilysa2.WithValue(beatsaber.EventValueLightOff))
	}
}
