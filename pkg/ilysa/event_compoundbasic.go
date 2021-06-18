package ilysa

import "github.com/shasderias/ilysa/pkg/beatsaber"

type CompoundBasicLightingEvent []*BasicLightingEvent
type CompoundBasicLightingEventOpt interface {
	applyCompoundBasicLightingEvent(*CompoundBasicLightingEvent)
}

func NewCompoundBasicLightingEvent(events ...*BasicLightingEvent) *CompoundBasicLightingEvent {
	compoundEvent := CompoundBasicLightingEvent{}
	compoundEvent = append(compoundEvent, events...)
	return &compoundEvent
}

func (e *CompoundBasicLightingEvent) Add(events ...*BasicLightingEvent) {
	*e = append(*e, events...)
}

func (e *CompoundBasicLightingEvent) SetValue(val beatsaber.EventValue) {
	for i := range *e {
		(*e)[i].SetValue(val)
	}
}

func (e *CompoundBasicLightingEvent) Mod(opts ...CompoundBasicLightingEventOpt) {
	for _, opt := range opts {
		opt.applyCompoundBasicLightingEvent(e)
	}
}
