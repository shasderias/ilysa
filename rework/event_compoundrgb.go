package rework

//type CompoundRGBLightingEvent []*evt.RGBLighting
//type CompoundRGBLightingEventOpt interface {
//	applyCompoundRGBLightingEvent(*CompoundRGBLightingEvent)
//}
//
//func NewCompoundRGBLightingEvent(events ...*evt.RGBLighting) *CompoundRGBLightingEvent {
//	compoundEvent := CompoundRGBLightingEvent{}
//	compoundEvent = append(compoundEvent, events...)
//	return &compoundEvent
//}
//
//func (e *CompoundRGBLightingEvent) Add(events ...*evt.RGBLighting) {
//	*e = append(*e, events...)
//}
//
//func (e *CompoundRGBLightingEvent) ShiftBeat(offset float64) {
//	for i := range *e {
//		(*e)[i].beat += offset
//	}
//}
//
//func (e *CompoundRGBLightingEvent) SetValue(val beatsaber.EventValue) {
//	for i := range *e {
//		(*e)[i].SetValue(val)
//	}
//}
//
//func (e *CompoundRGBLightingEvent) GetColor() color.Color {
//	return (*e)[0].Color()
//}
//
//func (e *CompoundRGBLightingEvent) SetColor(c color.Color) {
//	for i := range *e {
//		(*e)[i].SetColor(c)
//	}
//}
//
//func (e *CompoundRGBLightingEvent) Alpha() float64 {
//	return (*e)[0].Alpha()
//}
//
//func (e *CompoundRGBLightingEvent) SetAlpha(a float64) {
//	for i := range *e {
//		(*e)[i].SetAlpha(a)
//	}
//}
//
//func (e *CompoundRGBLightingEvent) Mod(opts ...CompoundRGBLightingEventOpt) {
//	for _, opt := range opts {
//		opt.applyCompoundRGBLightingEvent(e)
//	}
//}
