package ilysa

import (
	"ilysa/pkg/beatsaber"
)

// BasicLight represents a light with the base game's attributes. Lighting events
// created by BasicLight do not set _customData._lightID,
type BasicLight struct {
	eventType  beatsaber.EventType
	maxLightID int
}

func (p *Project) NewBasicLight(typ beatsaber.EventType) BasicLight {
	return BasicLight{
		eventType:  typ,
		maxLightID: p.ActiveDifficultyProfile().LightIDMax(typ),
	}
}

func NewBasicLight(typ beatsaber.EventType, m LightIDMaxer) BasicLight {
	maxLightID := m.LightIDMax(typ)
	return BasicLight{
		eventType:  typ,
		maxLightID: maxLightID,
	}
}

func (bl BasicLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLightingEvent().SetLight(bl.eventType),
	)
}

func (bl BasicLight) EventType() beatsaber.EventType {
	return bl.eventType
}

func (bl BasicLight) EventTypeSet() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(bl.eventType)
}

func (bl BasicLight) LightIDLen() int {
	return 1
}

func (bl BasicLight) LightIDSet() LightIDSet {
	return NewLightIDSet(
		NewLightIDFromInterval(1, bl.maxLightID),
	)
}

func (bl BasicLight) LightIDTransform(tfer LightIDTransformer) Light {
	return CompositeLight{
		eventType: bl.eventType,
		set:       tfer(NewLightIDFromInterval(1, bl.maxLightID)),
	}
}

func (bl BasicLight) LightIDSequenceTransform(tfer LightIDTransformer) Light {
	sl := NewSequenceLight()
	set := tfer(NewLightIDFromInterval(1, bl.maxLightID))

	for _, lightID := range set {
		sl.Add(CompositeLight{
			eventType: bl.eventType,
			set:       LightIDSet{lightID},
		})

	}
	return sl
}

func (bl BasicLight) LightIDSetTransform(tfer LightIDSetTransformer) Light {
	return CompositeLight{
		eventType: bl.eventType,
		set:       tfer(bl.LightIDSet()),
	}
}

func (bl BasicLight) LightIDSetSequenceTransform(tfer LightIDSetTransformer) Light {
	sl := NewSequenceLight()
	set := tfer(bl.LightIDSet())

	for _, lightID := range set {
		sl.Add(CompositeLight{
			eventType: bl.eventType,
			set:       LightIDSet{lightID},
		})

	}
	return sl
}

func (bl BasicLight) Transform(tfer LightIDTransformer) CompositeLight {
	return CompositeLight{
		eventType: bl.eventType,
		set:       tfer(NewLightIDFromInterval(1, bl.maxLightID)),
	}
}

func (bl BasicLight) TransformToSequence(tfer LightIDTransformer) SequenceLight {
	sl := NewSequenceLight()
	set := tfer(NewLightIDFromInterval(1, bl.maxLightID))

	for _, lightID := range set {
		sl.Add(CompositeLight{
			eventType: bl.eventType,
			set:       LightIDSet{lightID},
		})

	}
	return sl
}
