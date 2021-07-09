package rework

type lightContext struct {
	baseContext
	lightTimer
}

func newLightContext(c baseContext, lt lightTimer) lightContext {
	return lightContext{
		baseContext: c,
		lightTimer:  lt,
	}
}
