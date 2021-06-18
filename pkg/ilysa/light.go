package ilysa

import (
	"math"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/util"
)

type Light interface {
	CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent
	EventType() beatsaber.EventTypeSet
	LightIDLen() int
}

// BasicLight represents a light with the base game's attributes
type BasicLight struct {
	eventType  beatsaber.EventType
	lightIDMax int
}

func (p *Project) NewBasicLight(typ beatsaber.EventType) BasicLight {
	return BasicLight{
		eventType:  typ,
		lightIDMax: p.ActiveDifficultyProfile().LightIDMax(typ),
	}
}

func (l BasicLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLightingEvent().SetLight(l.eventType),
	)
}

func (l BasicLight) EventType() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(l.eventType)
}

func (l BasicLight) LightIDLen() int {
	return 1
}

func (l BasicLight) LightIDMax() int {
	return l.lightIDMax
}

func (l BasicLight) LightIDTransform(tfer LightIDTransformer) Light {
	return CompoundLight{
		eventType: l.eventType,
		set:       tfer(NewLightIDFromInterval(1, l.lightIDMax)),
	}
}

func (l BasicLight) LightIDTransformSequence(tfer LightIDTransformer) SequenceLight {
	sl := NewSequenceLight()
	set := tfer(NewLightIDFromInterval(1, l.lightIDMax))

	for _, lightID := range set {
		sl.Add(CompoundLight{
			eventType: l.eventType,
			set:       LightIDSet{lightID},
		})

	}
	return sl
}

func (l BasicLight) Transform(tfer LightIDTransformer) CompoundLight {
	return CompoundLight{
		eventType: l.eventType,
		set:       tfer(NewLightIDFromInterval(1, l.lightIDMax)),
	}
}

// [1,2,3,4] => [1], [2], [3], [4]
func (l BasicLight) TransformToSequence(tfer LightIDTransformer) SequenceLight {
	sl := NewSequenceLight()
	set := tfer(NewLightIDFromInterval(1, l.lightIDMax))

	for _, lightID := range set {
		sl.Add(CompoundLight{
			eventType: l.eventType,
			set:       LightIDSet{lightID},
		})

	}
	return sl
}

type CompoundLight struct {
	eventType beatsaber.EventType
	set       LightIDSet
}

func (cl CompoundLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLightingEvent(
			WithType(cl.eventType),
			WithLightID(cl.set.Index(ctx.LightIDOrdinal())),
		),
	)
}

func (cl CompoundLight) EventType() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(cl.eventType)
}

func (cl CompoundLight) LightIDLen() int {
	return cl.set.Len()
}

func (cl CompoundLight) LightIDTransform(tfer LightIDTransformer) Light {
	newSet := LightIDSet{}

	for _, lid := range cl.set {
		newSet.Add(tfer(lid)...)
	}

	return CompoundLight{
		eventType: cl.eventType,
		set:       newSet,
	}
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

func (cl CombinedLight) EventType() beatsaber.EventTypeSet {
	et := beatsaber.NewEventTypeSet()
	for _, l := range cl.lights {
		et = et.Union(l.EventType())
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

func (sl SequenceLight) EventType() beatsaber.EventTypeSet {
	et := beatsaber.NewEventTypeSet()
	for _, l := range sl.lights {
		et = et.Union(l.EventType())
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
