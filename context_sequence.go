package ilysa

import "github.com/shasderias/ilysa/calc"

type sequenceContext struct {
	baseContext
	sequence []float64
}

func newSequenceContext(c baseContext, s []float64) sequenceContext {
	return sequenceContext{
		baseContext: c,
		sequence:    s,
	}
}

func (c sequenceContext) SequenceIndex(idx int) float64 {
	return c.sequence[calc.WraparoundIdx(len(c.sequence), idx)]
}

func (c sequenceContext) OrdinalOffset(offset int) float64 {
	return c.sequence[calc.WraparoundIdx(len(c.sequence), c.Ordinal()+offset)]
}

func (c sequenceContext) NextB() float64 {
	return c.B() + c.NextBOffset()
}

func (c sequenceContext) NextBOffset() float64 {
	return c.OrdinalOffset(1) - c.OrdinalOffset(0)
}

func (c sequenceContext) PrevB() float64 {
	return c.B() + c.PrevBOffset()
}

func (c sequenceContext) PrevBOffset() float64 {
	return c.OrdinalOffset(0) - c.OrdinalOffset(-1)
}

func (c sequenceContext) WithLight(l Light, callback func(ctx SequenceLightContext)) {
	for i := 0; i < l.LightIDLen(); i++ {
		callback(newSequenceLightContext(c.baseContext, c, l, i))
	}
}
