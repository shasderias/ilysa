package ilysa

import (
	"math"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/util"
)

type Light interface {
	CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent
	EventType() beatsaber.EventTypeSet
	LightIDMin() int
	LightIDMax() int
	LightIDLen() int
}

// BasicLight represents c light with the base game + OOB Chroma attributes
type BasicLight struct {
	project   *Project
	eventType beatsaber.EventType
}

func (p *Project) NewBasicLight(typ beatsaber.EventType) BasicLight {
	return BasicLight{
		project:   p,
		eventType: typ,
	}
}

func (l BasicLight) CreateRGBEvent(ctx TimingContext) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLightingEvent().SetLight(l.eventType),
	)
}

func (l BasicLight) EventType() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(l.eventType)
}

func (l BasicLight) LightIDMin() int {
	return 1
}

func (l BasicLight) LightIDMax() int {
	return 1
}

func (l BasicLight) LightIDLen() int {
	return 1
}

func (l BasicLight) Split(splitter LightIDSplitter) SplitLight {
	maxLightID := l.project.ActiveDifficultyProfile().LightIDMax(l.eventType)
	return SplitLight{
		project:   l.project,
		eventType: l.eventType,
		set:       splitter(NewLightIDFromInterval(1, maxLightID)),
	}
}

func (l BasicLight) SplitToSequence(splitter LightIDSplitter) *SequenceLight {
	maxLightID := l.project.ActiveDifficultyProfile().LightIDMax(l.eventType)

	sl := NewSequenceLight()
	set := splitter(NewLightIDFromInterval(1, maxLightID))

	for _, lightID := range *set {
		sl.Add(SplitLight{
			project:   l.project,
			eventType: l.eventType,
			set:       &LightIDSet{lightID},
		})

	}

	return sl
}

type SplitLight struct {
	project   *Project
	eventType beatsaber.EventType
	set       *LightIDSet
}

func (l SplitLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	return NewCompoundRGBLightingEvent(
		ctx.NewRGBLightingEvent(
			WithType(l.eventType),
			WithLightID(l.set.Index(ctx.LightIDOrdinal())),
		),
	)
}

func (l SplitLight) EventType() beatsaber.EventTypeSet {
	return beatsaber.NewEventTypeSet(l.eventType)
}

func (l SplitLight) LightIDMin() int {
	return 1
}

func (l SplitLight) LightIDMax() int {
	return l.set.Len()
}

func (l SplitLight) LightIDLen() int {
	return l.set.Len()
}

type CombinedLight struct {
	lights []Light
}

func NewCombinedLight(lights ...Light) *CombinedLight {
	return &CombinedLight{
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

func (cl CombinedLight) LightIDMin() int {
	return 1
}

func (cl CombinedLight) LightIDMax() int {
	max := 0
	for _, light := range cl.lights {
		if light.LightIDMax() > max {
			max = light.LightIDMax()
		}
	}
	return max
}

func (cl CombinedLight) LightIDLen() int {
	return cl.LightIDMax()
}

type SequenceLight struct {
	lights []Light
}

func NewSequenceLight(lights ...Light) *SequenceLight {
	return &SequenceLight{append([]Light{}, lights...)}
}

func (sl *SequenceLight) Add(lights ...Light) {
	sl.lights = append(sl.lights, lights...)
}

func (sl *SequenceLight) CreateRGBEvent(ctx TimingContextForLight) *CompoundRGBLightingEvent {
	light := sl.Index(ctx.Ordinal())

	return light.CreateRGBEvent(ctx)
}

func (sl *SequenceLight) EventType() beatsaber.EventTypeSet {
	et := beatsaber.NewEventTypeSet()
	for _, l := range sl.lights {
		et = et.Union(l.EventType())
	}
	return et
}

func (sl *SequenceLight) Index(idx int) Light {
	l := len(sl.lights)
	return sl.lights[util.WraparoundIdx(l, idx)]
}

func (sl *SequenceLight) LightIDMin() int {
	return 1
}

func (sl *SequenceLight) LightIDMax() int {
	max := 0
	for _, light := range sl.lights {
		if light.LightIDMax() > max {
			max = light.LightIDMax()
		}
	}
	return max
}

func (sl *SequenceLight) LightIDLen() int {
	return sl.LightIDMax()
}
