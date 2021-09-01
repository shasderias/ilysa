package light

import (
	"math"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/scale"
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
			for i := 1; i < l.LightIDLen(); i++ {
				if int(math.Round(lightIDScale(float64(i)))) == ctx.LightIDCur() {
					goto create
				}
			}
			continue
		}
	create:
		events.Add(l.NewRGBLighting(ctx)...)
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

func (c Combined) Name() []string {
	name := []string{}
	for _, l := range c.lights {
		name = append(name, l.Name()...)
	}
	return name
}
