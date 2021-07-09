package light2

import (
	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/rework"
)

// BasicLight represents a light with the base game's attributes. Lighting events
// created by BasicLight do not set _customData._lightID,
type BasicLight struct {
	eventType  beatsaber.EventType
	maxLightID int
}

func (p *ilysa.Project) NewBasicLight(typ beatsaber.EventType) BasicLight {
	return BasicLight{
		eventType:  typ,
		maxLightID: p.ActiveDifficultyProfile().LightIDMax(typ),
	}
}

func NewBasicLight(typ beatsaber.EventType, m LightIDMaxer) BasicLight {
	maxLightID := m.MaxLightID(typ)
	return BasicLight{
		eventType:  typ,
		maxLightID: maxLightID,
	}
}

func (bl BasicLight) CreateRGBLightingEvent(ctx LightContext) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLighting().SetLight(bl.eventType),
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

func (bl BasicLight) LightIDSet() ilysa.LightIDSet {
	return rework.NewLightIDSet(
		rework.NewLightIDFromInterval(1, bl.maxLightID),
	)
}

func (bl BasicLight) LightIDTransform(tfer ilysa.LightIDTransformer) Light {
	return CompositeLight{
		eventType: bl.eventType,
		set:       tfer(rework.NewLightIDFromInterval(1, bl.maxLightID)),
	}
}

func (bl BasicLight) LightIDSequenceTransform(tfer ilysa.LightIDTransformer) Light {
	sl := NewSequenceLight()
	set := tfer(rework.NewLightIDFromInterval(1, bl.maxLightID))

	for _, lightID := range set {
		sl.Add(CompositeLight{
			eventType: bl.eventType,
			set:       rework.LightIDSet{lightID},
		})

	}
	return sl
}

func (bl BasicLight) LightIDSetTransform(tfer ilysa.LightIDSetTransformer) Light {
	return CompositeLight{
		eventType: bl.eventType,
		set:       tfer(bl.LightIDSet()),
	}
}

func (bl BasicLight) LightIDSetSequenceTransform(tfer ilysa.LightIDSetTransformer) Light {
	sl := NewSequenceLight()
	set := tfer(bl.LightIDSet())

	for _, lightID := range set {
		sl.Add(CompositeLight{
			eventType: bl.eventType,
			set:       rework.LightIDSet{lightID},
		})

	}
	return sl
}

func (bl BasicLight) Transform(tfer ilysa.LightIDTransformer) CompositeLight {
	return CompositeLight{
		eventType: bl.eventType,
		set:       tfer(rework.NewLightIDFromInterval(1, bl.maxLightID)),
	}
}

func (bl BasicLight) TransformToSequence(tfer ilysa.LightIDTransformer) SequenceLight {
	sl := NewSequenceLight()
	set := tfer(rework.NewLightIDFromInterval(1, bl.maxLightID))

	for _, lightID := range set {
		sl.Add(CompositeLight{
			eventType: bl.eventType,
			set:       rework.LightIDSet{lightID},
		})

	}
	return sl
}
