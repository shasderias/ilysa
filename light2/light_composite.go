package light2

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/rework"
)

// CompositeLight represents a single base game light and a set of LightIDs.
type CompositeLight struct {
	eventType beatsaber.EventType
	set       rework.LightIDSet
}

func NewCompositeLight(typ beatsaber.EventType, set rework.LightIDSet) CompositeLight {
	return CompositeLight{
		eventType: typ,
		set:       set,
	}
}

func (cl CompositeLight) CreateRGBLightingEvent(ctx LightContext) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLighting(
			WithType(cl.eventType),
			evt.WithLightID(cl.set.Index(ctx.LightIDOrdinal())),
		),
	)
}

func (cl CompositeLight) EventTypeSet() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(cl.eventType)
}

func (cl CompositeLight) LightIDLen() int {
	return cl.set.Len()
}

func (cl CompositeLight) LightIDSet() rework.LightIDSet {
	return cl.set
}

func (cl CompositeLight) LightIDTransform(tfer rework.LightIDTransformer) Light {
	newSet := rework.LightIDSet{}

	for _, lid := range cl.set {
		newSet.Add(tfer(lid)...)
	}

	return CompositeLight{
		eventType: cl.eventType,
		set:       newSet,
	}
}

func (cl CompositeLight) LightIDSequenceTransform(tfer rework.LightIDTransformer) Light {
	max := 0
	for _, lid := range cl.set {
		set := tfer(lid)
		if set.Len() > max {
			max = set.Len()
		}
	}

	compositeLights := make([]CompositeLight, max)

	for i := 0; i < max; i++ {
		compositeLights[i] = NewCompositeLight(cl.eventType, rework.NewLightIDSet())
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

func (cl CompositeLight) LightIDSetTransform(tfer rework.LightIDSetTransformer) Light {
	return CompositeLight{
		eventType: cl.eventType,
		set:       tfer(cl.set),
	}
}

func (cl CompositeLight) LightIDSetSequenceTransform(tfer rework.LightIDSetTransformer) Light {
	newSet := tfer(cl.set)

	seqLight := NewSequenceLight()

	for _, lightID := range newSet {
		seqLight.Add(NewCompositeLight(cl.eventType, rework.NewLightIDSet(lightID)))
	}

	return seqLight
}
