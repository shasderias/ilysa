package timer

import (
	"github.com/shasderias/ilysa/internal/calc"
	"github.com/shasderias/ilysa/internal/errors"
)

type Sequence interface {
	// SeqT is the current time for current sequence on a 0-1 scale. As a special
	// case, SeqT returns 1 when the sequence only has one beat.
	SeqT() float64
	SeqOrdinal() int
	SeqLen() int
	SeqNextB() float64
	SeqNextBOffset() float64
	SeqPrevB() float64
	SeqPrevBOffset() float64
	SeqFirst() bool
	SeqLast() bool

	Next() bool
	ToRange() Range
	ToSequence() Sequence
}

type Sequencer struct {
	s []float64
	g float64
}

// Seq specifies a sequence of beats.
func Seq(beats ...float64) Sequencer {
	if len(beats) == 0 {
		panic(errors.Errorf("sequence must contain at least one beat"))
	}
	return Sequencer{beats, beats[len(beats)-1] + 1}
}

// SeqInterval specifies a beat sequence that starts on startB, ends on endB,
// and has a beat every 1/interval beats.
func SeqInterval(startB, endB, interval float64) Sequencer {
	beats := []float64{}
	for b := startB; b <= endB; b += 1 / interval {
		beats = append(beats, b)
	}
	beats = append(beats, endB+1/interval)
	return Seq(beats...)
}

func (s Sequencer) Idx(i int) float64 { return s.s[calc.WraparoundIdx(s.Len(), i)] }
func (s Sequencer) Len() int          { return len(s.s) }
func (s Sequencer) First() float64    { return s.s[0] }
func (s Sequencer) Last() float64     { return s.s[s.Len()-1] }
func (s Sequencer) G(gb float64) Sequencer {
	s.g = gb
	return s
}

func (s Sequencer) Iterate() Sequence {
	return &SequenceIterator{s, -1}
}

type SequenceIterator struct {
	Sequencer
	ordinal int
}

func (i *SequenceIterator) Next() bool {
	i.ordinal++
	if i.ordinal == i.Len() {
		return false
	}
	return true
}

func (i *SequenceIterator) SeqT() float64 {
	if i.Len() == 1 {
		return 1
	}
	return float64(i.ordinal) / float64(i.Len()-1)
}
func (i *SequenceIterator) SeqOrdinal() int   { return i.ordinal }
func (i *SequenceIterator) SeqLen() int       { return i.Len() }
func (i *SequenceIterator) SeqNextB() float64 { return i.Idx(i.ordinal + 1) }
func (i *SequenceIterator) SeqNextBOffset() float64 {
	if i.SeqLast() {
		return i.g - i.Idx(i.ordinal)
	}
	return i.Idx(i.ordinal+1) - i.Idx(i.ordinal)
}
func (i *SequenceIterator) SeqPrevB() float64       { return i.Idx(i.ordinal - 1) }
func (i *SequenceIterator) SeqPrevBOffset() float64 { return i.Idx(i.ordinal) - i.Idx(i.ordinal-1) }
func (i *SequenceIterator) SeqFirst() bool          { return i.ordinal == 0 }
func (i *SequenceIterator) SeqLast() bool           { return i.ordinal == i.Len()-1 }

func (i *SequenceIterator) B() float64           { return i.Idx(i.ordinal) }
func (i *SequenceIterator) T() float64           { return i.SeqT() }
func (i *SequenceIterator) Ordinal() int         { return i.ordinal }
func (i *SequenceIterator) StartB() float64      { return i.Idx(0) }
func (i *SequenceIterator) EndB() float64        { return i.Idx(i.Len() - 1) }
func (i *SequenceIterator) Duration() float64    { return i.EndB() - i.StartB() }
func (i *SequenceIterator) First() bool          { return i.SeqFirst() }
func (i *SequenceIterator) Last() bool           { return i.SeqLast() }
func (i *SequenceIterator) ToRange() Range       { return i }
func (i *SequenceIterator) ToSequence() Sequence { return i }
