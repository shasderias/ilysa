package ilysa

import (
	"math"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/util"
)

type Light interface {
	CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent
	EventTypeSet() beatsaber.EventTypeSet
	LightIDLen() int
}

type LightIDMaxer interface {
	LightIDMax(beatsaber.EventType) int
}

// CompositeLight represents a single base game light and a set of LightIDs.
type CompositeLight struct {
	eventType beatsaber.EventType
	set       LightIDSet
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
	sl := NewSequenceLight()

	for _, lid := range cl.set {
		transformed := tfer(lid)
		transformed.Len()
	}

	return sl
}


type CombinedLight struct {
	lights []Light
}

func NewCombinedLight(lights ...Light) CombinedLight {
	return CombinedLight{
		lights: append([]Light{}, lights...),
	}
}

func (cl CombinedLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	ce := NewCompoundRGBLightingEvent()

	for _, l := range cl.lights {
		if l.LightIDLen() < cl.LightIDLen() {
			scale := util.Scale(1, float64(l.LightIDLen()), 1, float64(cl.LightIDLen()))
			for i := 1; i <= l.LightIDLen(); i++ {
				if int(math.Round(scale(float64(i)))) == ctx.LightIDCur() {
					goto createRGB
				}
			}
			continue
		}
	createRGB:
		ce.Add(*l.CreateRGBEvent(ctx)...)
	}

	return ce
}

func (cl *CombinedLight) Add(lights ...Light) {
	cl.lights = append(cl.lights, lights...)
}

func (cl CombinedLight) EventTypeSet() beatsaber.EventTypeSet {
	et := beatsaber.NewEventTypeSet()
	for _, l := range cl.lights {
		et = et.Union(l.EventTypeSet())
	}
	return et
}

func (cl CombinedLight) LightIDLen() int {
	max := 0
	for _, light := range cl.lights {
		if light.LightIDLen() > max {
			max = light.LightIDLen()
		}
	}
	return max
}

type SequenceLight struct {
	lights []Light
}

func NewSequenceLight(lights ...Light) SequenceLight {
	return SequenceLight{append([]Light{}, lights...)}
}

func (sl *SequenceLight) Add(lights ...Light) {
	sl.lights = append(sl.lights, lights...)
}

func (sl SequenceLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	light := sl.Index(ctx.Ordinal())

	return light.CreateRGBEvent(ctx)
}

func (sl SequenceLight) EventTypeSet() beatsaber.EventTypeSet {
	et := beatsaber.NewEventTypeSet()
	for _, l := range sl.lights {
		et = et.Union(l.EventTypeSet())
	}
	return et
}

func (sl SequenceLight) Index(idx int) Light {
	l := len(sl.lights)
	return sl.lights[util.WraparoundIdx(l, idx)]
}

func (sl SequenceLight) Slice(i, j int) SequenceLight {
	return SequenceLight{lights: sl.lights[i:j]}
}

func (sl SequenceLight) LightIDLen() int {
	max := 0
	for _, light := range sl.lights {
		if light.LightIDLen() > max {
			max = light.LightIDLen()
		}
	}
	return max
}

func (sl SequenceLight) ApplyLightIDTransform(tfer LightIDTransformer) SequenceLight {
	newLights := []Light{}
	for _, l := range sl.lights {
		tfl, ok := l.(LightIDTransformable)
		if !ok {
			newLights = append(newLights, l)
			continue
		}
		newLights = append(newLights, tfl.LightIDTransform(tfer))
	}
	return NewSequenceLight(newLights...)
}
