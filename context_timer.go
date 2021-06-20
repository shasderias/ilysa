package ilysa

import "github.com/shasderias/ilysa/scale"

type timer struct {
	b         float64
	startBeat float64
	endBeat   float64
	ordinal   int
}

func newTimer(beat, startBeat, endBeat float64, ordinal int) timer {
	return timer{
		b:         beat,
		startBeat: startBeat,
		endBeat:   endBeat,
		ordinal:   ordinal,
	}
}

func (t timer) B() float64 {
	return t.b
}

func (t timer) T() float64 {
	return scale.ToUnitIntervalClamped(t.startBeat, t.endBeat)(t.b)
}

func (t timer) StartBeat() float64 {
	return t.startBeat
}

func (t timer) EndBeat() float64 {
	return t.endBeat
}

func (t timer) Duration() float64 {
	return t.endBeat - t.startBeat
}

func (t timer) Ordinal() int {
	return t.ordinal
}

func (t timer) First() bool {
	return t.b == t.startBeat
}

func (t timer) Last() bool {
	return t.b == t.endBeat
}
