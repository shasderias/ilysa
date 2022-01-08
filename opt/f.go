package opt

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

func FScaleT(ctx context.Context, a, b float64, easeFn ease.Func) float64 {
	return scale.FromUnitClamp(a, b)(easeFn(ctx.T()))
}

func FScaleSeqOrdinal(ctx interface {
	SeqLen() int
	SeqOrdinal() int
}, a, b float64) float64 {
	scaler := scale.Clamp(1, float64(ctx.SeqLen()), a, b)
	return scaler(float64(ctx.SeqOrdinal()))
}
