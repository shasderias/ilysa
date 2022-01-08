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
