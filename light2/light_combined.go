package light2

import (
	"math"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/scale"
)

type CombinedLight struct {
	lights []Light
}

func NewCombinedLight(lights ...Light) CombinedLight {
	return CombinedLight{
		lights: append([]Light{}, lights...),
	}
}

func (cl CombinedLight) CreateRGBLightingEvent(ctx LightContext) *CompoundRGBLightingEvent {
	ce := NewCompoundRGBLightingEvent()

	for _, l := range cl.lights {
		if l.LightIDLen() < cl.LightIDLen() {
			lightIDScale := scale.Clamped(1, float64(l.LightIDLen()), 1, float64(cl.LightIDLen()))
			for i := 1; i <= l.LightIDLen(); i++ {
				if int(math.Round(lightIDScale(float64(i)))) == ctx.LightIDCur() {
					goto createRGB
				}
			}
			continue
		}
	createRGB:
		ce.Add(*l.CreateRGBLightingEvent(ctx)...)
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
