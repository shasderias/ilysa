package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

// Ripple offsets each successive lightID by step beats
func Ripple(ctx context.LightContext, e evt.Events, step float64) {
	e.Apply(evt.OptShiftB(float64(ctx.LightOrdinal()) * step))
}

// RippleT offsets each successive lightID by a number of beats such that the last
// lightID triggers delay beats after the first
func RippleT(ctx context.LightContext, e evt.Events, delay float64) {
	e.Apply(evt.OptShiftB(ctx.LightT() * delay))
}
