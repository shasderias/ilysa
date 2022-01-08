package context

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type Project interface {
	Events() *evt.Events
}

type Sequencer interface {
	Iterate() timer.Sequence
}

type Ranger interface {
	Iterate() timer.Range
}

type Light interface {
	GenerateEvents(LightContext) evt.Events
	LightLen() int
	Name() []string
}

type Context interface {
	ConfigContext

	timer.Sequence
	timer.Range
	BOffset() float64
	FixedRand() float64

	base() *base
	baseCtx() bool

	WSeq(s Sequencer, callback func(ctx Context))
	WRng(r Ranger, callback func(ctx Context))
	WLight(l Light, callback func(ctx LightContext, e evt.Events))
	WBOffset(float64) Context

	Apply(e evt.Event)
	AddEvents(events ...evt.Event)
	Events() *evt.Events
}

type LightContext interface {
	timer.Sequence
	timer.Range
	BOffset() float64

	LightT() float64
	LightOrdinal() int
	LightLen() int
	LightCur() int

	Apply(e evt.Event)
	FixedRand() float64
	AddEvents(events ...evt.Event)
}

type ConfigContext interface {
	SetMapVersion(v beatsaber.DifficultyVersion)
	GetMapVersion() beatsaber.DifficultyVersion
}
