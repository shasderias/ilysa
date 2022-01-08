package light

import (
	"math"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
	"github.com/shasderias/ilysa/scale"
)

type Combined struct {
	lights []context.Light
}

func Combine(lights ...context.Light) Combined {
	return Combined{lights}
}

func (c Combined) GenerateEvents(ctx context.LightContext) evt.Events {
	events := evt.NewEvents()

	maxLightLen := c.LightLen()

	for _, l := range c.lights {
		if l.LightLen() < c.LightLen() {
			lightIDScale := scale.Clamp(1, float64(l.LightLen()), 1, float64(maxLightLen))
			for i := 1; i <= l.LightLen(); i++ {
				if int(math.RoundToEven(lightIDScale(float64(i)))) == ctx.LightCur() {
					lctx := context.LightContextAtOrdinal(ctx, l, i-1)
					events.Add(l.GenerateEvents(lctx)...)
					goto end
				}
			}
			continue
		}
		events.Add(l.GenerateEvents(ctx)...)
	end:
	}

	return events
}

func (c *Combined) Add(lights ...context.Light) {
	c.lights = append(c.lights, lights...)
}

func (c Combined) LightLen() int {
	max := 0
	for _, l := range c.lights {
		if l.LightLen() > max {
			max = l.LightLen()
		}
	}
	return max
}

func (c Combined) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	comb := Combine()

	for _, l := range c.lights {
		tfl, ok := l.(LightIDTransformableLight)
		if !ok {
			comb.Add(l)
		} else {
			comb.Add(tfl.LightIDTransform(fn))
		}
	}

	return comb
}

func (c Combined) LightIDSetTransform(fn func(lightid.Set) lightid.Set) context.Light {
	comb := Combine()

	for _, l := range c.lights {
		tfl, ok := l.(LightIDSetTransformableLight)
		if !ok {
			comb.Add(l)
		} else {
			comb.Add(tfl.LightIDSetTransform(fn))
		}
	}

	return comb
}

func (c Combined) Name() []string {
	name := []string{}
	for _, l := range c.lights {
		name = append(name, l.Name()...)
	}
	return name
}
