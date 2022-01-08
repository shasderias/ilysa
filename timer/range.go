package timer

import (
	"math"

	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

type Range interface {
	B() float64 // current beat
	// T is the current time in the current range on a 0-1 scale. As a special case,
	// T returns 1 when the range only has 1 step.
	T() float64
	Ordinal() int      // ordinal number of the current iteration, starting from 0
	StartB() float64   // first beat of the current sequence
	EndB() float64     // last beat of the current sequence
	Duration() float64 // duration of the current sequence, in beats
	First() bool       // true if this is the first iteration
	Last() bool        // true if this is the last iteration

	Next() bool
	ToRange() Range
	ToSequence() Sequence
}

type Ranger struct {
	startBeat float64
	endBeat   float64
	steps     int
	easeFn    func(float64) float64
	tToBeat   func(float64) float64
}

// Rng specifies a range of beats that starts on startB, ends on endB, has a total of
// n beats, eased using fn.
func Rng(startB, endB float64, n int, fn ease.Func) Ranger {
	return Ranger{
		startBeat: startB,
		endBeat:   endB,
		steps:     n,
		easeFn:    fn,
		tToBeat:   scale.FromUnitClamp(startB, endB),
	}
}

// RngInterval specifies a range of beats that starts on startB, ends on endB, has a
// beat every 1/interval beats. The range is then eased using fn.
func RngInterval(startB, endB, interval float64, fn ease.Func) Ranger {
	n := int(math.RoundToEven(endB-startB)*interval) + 1
	if n < 1 {
		n = 1
	}
	return Rng(startB, endB, n, fn)
}

func (r Ranger) Iterate() Range {
	return &RangeIterator{
		Ranger:  r,
		ordinal: -1,
		beatScaler: func(m float64) float64 {
			m = r.easeFn(m)
			return scale.FromUnitClamp(r.startBeat, r.endBeat)(m)
		},
	}
}

func (r Ranger) Idx(i int) float64 {
	t := float64(i) / float64(r.steps-1)
	return r.tToBeat(r.easeFn(t))
}

func (r Ranger) Len() int {
	return r.steps
}

type RangeIterator struct {
	Ranger
	ordinal    int
	beatScaler scale.Fn
}

func (i *RangeIterator) Next() bool {
	i.ordinal++
	if i.ordinal == i.steps {
		return false
	}
	return true
}

func (i *RangeIterator) B() float64        { return i.beatScaler(i.T()) }
func (i *RangeIterator) T() float64        { return float64(i.ordinal) / float64(i.steps-1) }
func (i *RangeIterator) Ordinal() int      { return i.ordinal }
func (i *RangeIterator) StartB() float64   { return i.startBeat }
func (i *RangeIterator) EndB() float64     { return i.endBeat }
func (i *RangeIterator) Duration() float64 { return i.endBeat - i.startBeat }
func (i *RangeIterator) First() bool       { return i.B() == i.StartB() }
func (i *RangeIterator) Last() bool        { return i.B() == i.EndB() }

// Sequence Methods
func (i *RangeIterator) SeqT() float64     { return float64(i.ordinal) / float64(i.Len()-1) }
func (i *RangeIterator) SeqOrdinal() int   { return i.ordinal }
func (i *RangeIterator) SeqLen() int       { return i.Len() }
func (i *RangeIterator) SeqNextB() float64 { return i.Idx(i.ordinal + 1) }
func (i *RangeIterator) SeqNextBOffset() float64 {
	if i.SeqLast() {
		// approximation
		return i.Idx(i.ordinal) - i.Idx(i.ordinal-1)
	}
	return i.Idx(i.ordinal+1) - i.Idx(i.ordinal)
}
func (i *RangeIterator) SeqPrevB() float64       { return i.Idx(i.ordinal - 1) }
func (i *RangeIterator) SeqPrevBOffset() float64 { return i.Idx(i.ordinal) - i.Idx(i.ordinal-1) }
func (i *RangeIterator) SeqFirst() bool          { return i.ordinal == 0 }
func (i *RangeIterator) SeqLast() bool           { return i.ordinal == i.Len()-1 }
func (i *RangeIterator) ToRange() Range          { return i }
func (i *RangeIterator) ToSequence() Sequence    { return i }
