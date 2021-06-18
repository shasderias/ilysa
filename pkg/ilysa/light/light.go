package light

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/ilysa"
)

type IDMaxer interface {
	LightIDMax(beatsaber.EventType) int
}

// Basic represents a light with the base game's attributes. Lighting events
// created by Basic do not set _customData._lightID,
type Basic struct {
	eventType  beatsaber.EventType
	lightIDMax int
}

func NewBasic(typ beatsaber.EventType, m IDMaxer) Basic {
	maxLightID := m.LightIDMax(typ)
	return Basic{
		eventType:  typ,
		lightIDMax: maxLightID,
	}
}

func (bl Basic) CreateRGBEvent(ctx ilysa.TimingContextForLight) *ilysa.CompoundRGBLightingEvent {
	return ilysa.NewCompoundRGBLightingEvent(
		ctx.NewRGBLightingEvent().SetLight(bl.eventType),
	)
}

func (bl Basic) EventType() beatsaber.EventType {
	return bl.eventType
}

func (bl Basic) EventTypeSet() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(bl.eventType)
}

func (bl Basic) LightIDLen() int {
	return 1
}

func (bl Basic) LightIDSet() IDSet{
	return NewLightIDSet(
		NewLightIDFromInterval(1, bl.lightIDMax),
	)
}

//func (bl Basic) LightIDTransform(tfer LightIDTransformer) Light {
//	return CompositeLight{
//		eventType: bl.eventType,
//		set:       tfer(NewLightIDFromInterval(1, bl.lightIDMax)),
//	}
//}
//
//func (bl Basic) LightIDTransformSequence(tfer LightIDTransformer) SequenceLight {
//	sl := NewSequenceLight()
//	set := tfer(NewLightIDFromInterval(1, bl.lightIDMax))
//
//	for _, lightID := range set {
//		sl.Add(CompositeLight{
//			eventType: bl.eventType,
//			set:       LightIDSet{lightID},
//		})
//
//	}
//	return sl
//}
//
//func (bl Basic) Transform(tfer LightIDTransformer) CompositeLight {
//	return CompositeLight{
//		eventType: bl.eventType,
//		set:       tfer(NewLightIDFromInterval(1, bl.lightIDMax)),
//	}
//}
//
//func (bl Basic) TransformToSequence(tfer LightIDTransformer) SequenceLight {
//	sl := NewSequenceLight()
//	set := tfer(NewLightIDFromInterval(1, bl.lightIDMax))
//
//	for _, lightID := range set {
//		sl.Add(CompositeLight{
//			eventType: bl.eventType,
//			set:       LightIDSet{lightID},
//		})
//
//	}
//	return sl
//}
