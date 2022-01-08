package fx

import (
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

func Gradient(ctx context.LightContext, grad gradient.Table) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		e.Apply(evt.OptColor(grad.Lerp(ctx.LightT())))
	})
}

func GradientT(ctx context.LightContext, grad gradient.Table) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		e.Apply(evt.OptColor(grad.Lerp(ctx.T())))
	})
}
