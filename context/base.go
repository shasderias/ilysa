package context

import (
	"math/rand"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func Base(project Project) Context {
	ctx := base{
		project:   project,
		fixedRand: rand.Float64(),
		config:    config{},
	}
	return &ctx
}

type base struct {
	project                   Project
	fixedRand                 float64
	cumulativePreciseRotation float64
	config                    config
}

func (b *base) base() *base   { return b }
func (b *base) baseCtx() bool { return true }

func (b *base) Apply(e evt.Event) {
	e.SetBeat(b.B() + b.BOffset())
	b.AddEvents(e)
}
func (b *base) FixedRand() float64 { return b.fixedRand }
func (b *base) AddEvents(events ...evt.Event) {
	*(b.project.Events()) = append(*(b.project.Events()), events...)
}
func (b *base) Events() *evt.Events { return b.project.Events() }

func (b *base) WSeq(s Sequencer, callback func(ctx Context)) { newSeqCtx(b, s, callback) }
func (b *base) WRng(r Ranger, callback func(ctx Context))    { newRngCtx(b, r, callback) }
func (b *base) WLight(l Light, callback func(ctx LightContext, e evt.Events)) {
	newLightCtx(b, l, callback)
}
func (b *base) WBOffset(o float64) Context { return newBOffsetCtx(b, o) }

func (b *base) BOffset() float64  { return 0 }
func (b *base) B() float64        { return 0 }
func (b *base) T() float64        { panic("implement me") }
func (b *base) Ordinal() int      { panic("implement me") }
func (b *base) StartB() float64   { panic("implement me") }
func (b *base) EndB() float64     { panic("implement me") }
func (b *base) Duration() float64 { panic("implement me") }
func (b *base) First() bool       { panic("implement me") }
func (b *base) Last() bool        { panic("implement me") }

func (b *base) SeqT() float64              { panic("implement me") }
func (b *base) SeqOrdinal() int            { panic("implement me") }
func (b *base) SeqLen() int                { panic("implement me") }
func (b *base) SeqNextB() float64          { panic("implement me") }
func (b *base) SeqNextBOffset() float64    { panic("implement me") }
func (b *base) SeqPrevB() float64          { panic("implement me") }
func (b *base) SeqPrevBOffset() float64    { panic("implement me") }
func (b *base) SeqFirst() bool             { panic("implement me") }
func (b *base) SeqLast() bool              { panic("implement me") }
func (b *base) ToRange() timer.Range       { return b }
func (b *base) ToSequence() timer.Sequence { return b }
func (b *base) Next() bool                 { return false }

func (b *base) SetMapVersion(v beatsaber.DifficultyVersion) { b.config.mapVersion = v }
func (b *base) GetMapVersion() beatsaber.DifficultyVersion  { return b.config.mapVersion }
