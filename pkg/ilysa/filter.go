package ilysa

import "ilysa/pkg/beatsaber"

var FilterAllLightingEvents EventFilter = func(event Event) bool {
	_, ok := event.(*RGBLightingEvent)
	return ok
}

func FilterLightingEvents(targetType beatsaber.EventType) EventFilter {
	return func(event Event) bool {
		e, ok := event.(*RGBLightingEvent)
		if !ok {
			return false
		}

		return e.Base().Type == targetType
	}
}
