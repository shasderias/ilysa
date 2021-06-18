package ilysa

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/util"
)

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

func (sl SequenceLight) LightIDLen() int {
	max := 0
	for _, light := range sl.lights {
		if light.LightIDLen() > max {
			max = light.LightIDLen()
		}
	}
	return max
}

func (sl SequenceLight) Index(idx int) Light {
	l := len(sl.lights)
	return sl.lights[util.WraparoundIdx(l, idx)]
}

func (sl SequenceLight) Slice(i, j int) SequenceLight {
	return SequenceLight{lights: sl.lights[i:j]}
}

func (sl SequenceLight) LightIDTransform(tfer LightIDTransformer) Light {
	tfedLights := []Light{}
	for _, l := range sl.lights {
		tfl, ok := l.(LightIDTransformable)
		if !ok {
			tfedLights = append(tfedLights, l)
			continue
		}
		tfedLights = append(tfedLights, tfl.LightIDTransform(tfer))
	}
	return NewSequenceLight(tfedLights...)
}
