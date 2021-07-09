package rework

//var FilterAllLightingEvents EventFilter = func(event Event) bool {
//	switch event.(type) {
//	case *evt.Lighting:
//		return true
//	case *evt.RGBLighting:
//		return true
//	}
//	return false
//}
//
//func FilterLightingEvents(targetType beatsaber.EventType) EventFilter {
//	return func(event Event) bool {
//		e, ok := event.(*evt.RGBLighting)
//		if !ok {
//			return false
//		}
//
//		return e.Type() == targetType
//	}
//}
//
//func FilterRGBLights() EventFilter {
//	return func(event Event) bool {
//		_, ok := event.(*evt.RGBLighting)
//		return ok
//	}
//}
//
//func FilterRGBLight(light light2.Light) EventFilter {
//	return func(event Event) bool {
//		_, ok := event.(*evt.RGBLighting)
//		if !ok {
//			return false
//		}
//
//		return light.EventTypeSet().Has(event.Type())
//	}
//}
