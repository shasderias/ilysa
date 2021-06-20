package ilysa

type timeLightContext struct {
	baseContext
	lightTimer
}

func newTimeLightContext(c baseContext, l Light, ordinal int) timeLightContext {
	return timeLightContext{
		baseContext: c,
		lightTimer:  newLightTimer(l, ordinal),
	}
}

func (c timeLightContext) NewLightingEvent(opts ...BasicLightingEventOpt) *CompoundBasicLightingEvent {
	events := CompoundBasicLightingEvent{}

	if c.LightIDOrdinal() > 0 {
		return &events
	}

	for _, et := range c.Light.EventTypeSet() {
		opts := append([]BasicLightingEventOpt{WithType(et)}, opts...)
		events.Add(c.baseContext.NewLightingEvent(opts...))
	}

	return &events
}

func (c timeLightContext) NewRGBLightingEvent(opts ...CompoundRGBLightingEventOpt) *CompoundRGBLightingEvent {
	e := c.Light.CreateRGBLightingEvent(newLightContext(c.baseContext, c.lightTimer))
	e.Mod(opts...)
	return e
}
