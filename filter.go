package ilysa

import "github.com/shasderias/ilysa/beatsaber"

var FilterAllLightingEvents EventFilter = func(event Event) bool {
	switch event.(type) {
	case *BasicLightingEvent:
		return true
	case *RGBLightingEvent:
		return true
	}
	return false
}

func FilterLightingEvents(targetType beatsaber.EventType) EventFilter {
	return func(event Event) bool {
		e, ok := event.(*RGBLightingEvent)
		if !ok {
			return false
		}

		return e.Type() == targetType
	}
}

func FilterRGBLights() EventFilter {
	return func(event Event) bool {
		_, ok := event.(*RGBLightingEvent)
		return ok
	}
}

func FilterRGBLight(light Light) EventFilter {
	return func(event Event) bool {
		_, ok := event.(*RGBLightingEvent)
		if !ok {
			return false
		}

		return light.EventTypeSet().Has(event.Type())
	}
}
