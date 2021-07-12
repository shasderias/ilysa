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
