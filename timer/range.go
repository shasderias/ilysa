package timer

import (
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

type Range interface {
	B() float64        // current beat
	T() float64        // current time in the current sequence, on a 0-1 scale
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

func NewRanger(startBeat, endBeat float64, steps int, fn ease.Func) Ranger {
	return Ranger{
		startBeat: startBeat,
		endBeat:   endBeat,
		steps:     steps,
		easeFn:    fn,
		tToBeat:   scale.FromUnitClamp(startBeat, endBeat),
	}
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
	t := float64(i) / float64(r.steps)
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
