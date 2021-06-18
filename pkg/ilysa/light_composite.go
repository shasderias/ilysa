package ilysa

import (
	"ilysa/pkg/beatsaber"
)

// CompositeLight represents a single base game light and a set of LightIDs.
type CompositeLight struct {
	eventType beatsaber.EventType
	set       LightIDSet
}

func NewCompositeLight(typ beatsaber.EventType, set LightIDSet) CompositeLight {
	return CompositeLight{
		eventType: typ,
		set:       set,
	}
}

func (cl CompositeLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLightingEvent(
			WithType(cl.eventType),
			WithLightID(cl.set.Index(ctx.LightIDOrdinal())),
		),
	)
}

func (cl CompositeLight) EventTypeSet() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(cl.eventType)
}

func (cl CompositeLight) LightIDLen() int {
	return cl.set.Len()
}

func (cl CompositeLight) LightIDSet() LightIDSet {
	return cl.set
}

func (cl CompositeLight) LightIDTransform(tfer LightIDTransformer) Light {
	newSet := LightIDSet{}

	for _, lid := range cl.set {
		newSet.Add(tfer(lid)...)
	}

	return CompositeLight{
		eventType: cl.eventType,
		set:       newSet,
	}
}

func (cl CompositeLight) LightIDTransformSequence(tfer LightIDTransformer) Light {
	max := 0
	for _, lid := range cl.set {
		set := tfer(lid)
		if set.Len() > max {
			max = set.Len()
		}
	}

	compositeLights := make([]CompositeLight, 0, max)

	for i := 0; i < max; i++ {
		compositeLights[i] = NewCompositeLight(cl.eventType, NewLightIDSet())
	}

	for _, lid := range cl.set {
		set := tfer(lid)
		for i, lightID := range set {
			compositeLights[i].set.Add(lightID)
		}
	}

	seqLight := NewSequenceLight()
	for _, l := range compositeLights {
		seqLight.Add(l)
	}

	return seqLight
}
