package ilysa

type sequenceLightContext struct {
	baseContext
	sequenceContext
	lightTimer
}

//func newSequenceLightContext(bc baseContext, sc sequenceContext, l light2.Light, lightIDOrdinal int) sequenceLightContext {
//	return sequenceLightContext{
//		baseContext:     bc,
//		sequenceContext: sc,
//		lightTimer:      newLightTimer(l, lightIDOrdinal),
//	}
//}
//
//func (c sequenceLightContext) NewLightingEvent(opts ...evt.LightingOpt) *CompoundBasicLightingEvent {
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
//func (c sequenceLightContext) NewRGBLightingEvent(opts ...CompoundRGBLightingEventOpt) *CompoundRGBLightingEvent {
//	e := c.Light.CreateRGBLightingEvent(newLightContext(c.baseContext, c.lightTimer))
//	e.Mod(opts...)
//	return e
//}
