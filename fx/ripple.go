package fx

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

// Ripple offsets each successive lightID by step beats.
func Ripple(ctx context.LightContext, step float64) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		e.Apply(evt.OptShiftB(float64(ctx.LightOrdinal()) * step))
	})
}

// RippleT offsets each successive lightID such that the last lightID triggers
// delay beats after the first.
func RippleT(ctx context.LightContext, delay float64) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		e.Apply(evt.OptShiftB(ctx.LightT() * delay))
	})
}
