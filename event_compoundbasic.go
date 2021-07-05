package ilysa

//type CompoundBasicLightingEvent []*evt.Lighting
//type CompoundBasicLightingEventOpt interface {
//	applyCompoundBasicLightingEvent(*CompoundBasicLightingEvent)
//}
//
//func NewCompoundBasicLightingEvent(events ...*evt.Lighting) *CompoundBasicLightingEvent {
//	compoundEvent := CompoundBasicLightingEvent{}
//	compoundEvent = append(compoundEvent, events...)
//	return &compoundEvent
//}
//
//func (e *CompoundBasicLightingEvent) Add(events ...*evt.Lighting) {
//	*e = append(*e, events...)
//}
//
//func (e *CompoundBasicLightingEvent) SetValue(val beatsaber.EventValue) {
//	for i := range *e {
//		(*e)[i].SetValue(val)
//	}
//}
//
//func (e *CompoundBasicLightingEvent) Mod(opts ...CompoundBasicLightingEventOpt) {
//	for _, opt := range opts {
//		opt.applyCompoundBasicLightingEvent(e)
//	}
//}
