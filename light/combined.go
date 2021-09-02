package light

import (
	"math"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
	"github.com/shasderias/ilysa/scale"
	"github.com/shasderias/ilysa/timer"
)

type Combined struct {
	lights []context.Light
}

func Combine(lights ...context.Light) Combined {
	return Combined{lights}
}

func (c Combined) NewRGBLighting(ctx context.LightRGBLightingContext) evt.RGBLightingEvents {
	events := evt.RGBLightingEvents{}

	maxLightIDLen := c.LightIDLen()

	for _, l := range c.lights {
		if l.LightIDLen() < c.LightIDLen() {
			lightIDScale := scale.Clamp(1, float64(l.LightIDLen()), 1, float64(maxLightIDLen))
			for i := 1; i <= l.LightIDLen(); i++ {
				if int(math.RoundToEven(lightIDScale(float64(i)))) == ctx.LightIDCur() {
					lt := timer.NewLighter(l)
					lctx := context.WithLightTimer(ctx, lt.IterateFrom(i-1))
					events.Add(l.NewRGBLighting(lctx)...)
					goto end
				}
			}
			continue
		}
		events.Add(l.NewRGBLighting(ctx)...)
	end:
	}

	return events
}

func (c *Combined) Add(lights ...context.Light) {
	c.lights = append(c.lights, lights...)
}

func (c Combined) LightIDLen() int {
	max := 0
	for _, l := range c.lights {
		if l.LightIDLen() > max {
			max = l.LightIDLen()
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
