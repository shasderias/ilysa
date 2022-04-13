package fx

import (
	"github.com/shasderias/ilysa/colorful/gradient"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

func Gradient(ctx context.LightContext, e evt.Events, table gradient.Table) {
	e.Apply(evt.OptColor(table.Lerp(ctx.LightT())))
}

func GradientT(ctx context.LightContext, e evt.Events, table gradient.Table) {
	e.Apply(evt.OptColor(table.Lerp(ctx.T())))
}

func GradientT2(ctx context.LightContext, table gradient.Table) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		e.Apply(evt.OptColor(table.Lerp(ctx.T())))
	})
}
