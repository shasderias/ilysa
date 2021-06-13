package ilysa

import "ilysa/pkg/util"

type Context struct {
	*Project
	Timing
}

type Timing struct {
	StartBeat float64
	EndBeat   float64
	Duration  float64
	B         float64
	Pos       float64
	Ordinal   int
	Last      bool
}

func NewTiming(startBeat, endBeat, beat float64, ordinal int) Timing {
	return Timing{
		StartBeat: startBeat,
		EndBeat:   endBeat,
		B:         beat,
		Duration:  endBeat - startBeat,
		Pos:       util.ScaleToUnitInterval(startBeat, endBeat)(beat),
		Ordinal:   ordinal,
		Last:      endBeat == beat,
	}
}

type EventGenerator func(ctx *Context)
type EventModder func(ctx *Context, events Event)
type EventFilter func(event Event) bool
