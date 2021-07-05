package light2

import (
	"math/rand"

	"github.com/shasderias/ilysa"
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/calc"
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

func (sl SequenceLight) CreateRGBLightingEvent(ctx LightContext) *CompoundRGBLightingEvent {
	light := sl.Index(ctx.Ordinal())

	return light.CreateRGBLightingEvent(ctx)
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
	return sl.lights[calc.WraparoundIdx(l, idx)]
}

func (sl SequenceLight) Len() int {
	return len(sl.lights)
}

func (sl SequenceLight) Slice(i, j int) SequenceLight {
	return SequenceLight{lights: sl.lights[i:j]}
}

func (sl SequenceLight) Shuffle() SequenceLight {
	rand.Shuffle(len(sl.lights), func(i, j int) {
		sl.lights[i], sl.lights[j] = sl.lights[j], sl.lights[i]
	})
	return SequenceLight{
		lights: sl.lights,
	}
}

func (sl SequenceLight) LightIDTransform(tfer ilysa.LightIDTransformer) Light {
	tfedLights := []Light{}
	for _, l := range sl.lights {
		tfl, ok := l.(ilysa.LightIDTransformable)
		if !ok {
			tfedLights = append(tfedLights, l)
			continue
		}
		tfedLights = append(tfedLights, tfl.LightIDTransform(tfer))
	}
	return NewSequenceLight(tfedLights...)
}

func (sl SequenceLight) LightIDSetTransform(tfer ilysa.LightIDSetTransformer) Light {
	tfedLights := []Light{}
	for _, l := range sl.lights {
		tfl, ok := l.(ilysa.LightIDSetTransformable)
		if !ok {
			tfedLights = append(tfedLights, l)
			continue
		}
		tfedLights = append(tfedLights, tfl.LightIDSetTransform(tfer))
	}
	return NewSequenceLight(tfedLights...)
}
