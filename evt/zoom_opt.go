package evt

// WithZoomStep dictates how much position offset is added between each ring
func WithZoomStep(s float64) withZoomStepOpt {
	return withZoomStepOpt{s}
}

type withZoomStepOpt struct {
	s float64
}

func (o withZoomStepOpt) applyPreciseZoom(e *PreciseZoom) {
	e.Step = o.s
}

// WithZoomSpeed dictates how quickly each ring will move to its new position
func WithZoomSpeed(s float64) withZoomSpeedOpt {
	return withZoomSpeedOpt{s}
}

type withZoomSpeedOpt struct {
	s float64
}

func (o withZoomSpeedOpt) applyPreciseZoom(e *PreciseZoom) {
	e.Speed = o.s
}
