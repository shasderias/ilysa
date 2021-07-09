package rework

type timeLightContext struct {
	baseContext
	lightTimer
}

//func newTimeLightContext(c baseContext, l light2.Light, ordinal int) timeLightContext {
//	return timeLightContext{
//		baseContext: c,
//		lightTimer:  newLightTimer(l, ordinal),
//	}
//}
//
//func (c timeLightContext) NewLightingEvent(opts ...evt.LightingOpt) *CompoundBasicLightingEvent {
//	events := CompoundBasicLightingEvent{}
//
//	if c.LightIDOrdinal() > 0 {
//		return &events
//	}
//
//	for _, et := range c.Light.EventTypeSet() {
//		opts := append([]evt.LightingOpt{WithType(et)}, opts...)
//		events.Add(c.baseContext.NewLighting(opts...))
//	}
//
//	return &events
//}
//
//func (c timeLightContext) NewRGBLightingEvent(opts ...CompoundRGBLightingEventOpt) *CompoundRGBLightingEvent {
//	e := c.Light.CreateRGBLightingEvent(newLightContext(c.baseContext, c.lightTimer))
//	e.Mod(opts...)
//	return e
//}
