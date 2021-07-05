package context

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

func Base() Context {
	ctx := base{}
	ctx.eventer = newEventer(ctx)
	return ctx
}

type base struct {
	eventer
}

func (b base) addEvents(events ...evt.Event) {
	spew.Dump(events)
}

func (b base) Offset() float64 {
	return 0
}

func (b base) B() float64 {
	return 0
}

func (b base) NoOffsetB() float64 {
	return 0
}

func (b base) T() float64 {
	panic("implement me")
}

func (b base) Ordinal() int {
	panic("implement me")
}

func (b base) StartB() float64 {
	panic("implement me")
}

func (b base) EndB() float64 {
	panic("implement me")
}

func (b base) Duration() float64 {
	panic("implement me")
}

func (b base) First() bool {
	panic("implement me")
}

func (b base) Last() bool {
	panic("implement me")
}

func (b base) SeqT() float64 {
	panic("implement me")
}

func (b base) SeqOrdinal() int {
	panic("implement me")
}

func (b base) SeqLen() int {
	panic("implement me")
}

func (b base) SeqNextB() float64 {
	panic("implement me")
}

func (b base) SeqNextBOffset() float64 {
	panic("implement me")
}

func (b base) SeqPrevB() float64 {
	panic("implement me")
}

func (b base) SeqPrevBOffset() float64 {
	panic("implement me")
}

func (b base) SeqFirst() bool {
	panic("implement me")
}

func (b base) SeqLast() bool {
	panic("implement me")
}

func (b base) baseTimer() bool {
	return true
}

func (b base) ToRange() timer.Range {
	return b
}

func (b base) ToSequence() timer.Sequence {
	return b
}

func (b base) Next() bool {
	return false
}
