package light

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
)

type Custom struct {
	lightType evt.LightType
	len       int
	offset    int
}

func NewCustom(t evt.LightType, len, offset int) Custom {
	return Custom{
		t, len, offset,
	}
}

func (c Custom) NewRGBLighting(ctx context.LightRGBLightingContext) evt.RGBLightingEvents {
	return evt.RGBLightingEvents{
		ctx.NewRGBLighting(
			evt.WithLight(c.lightType),
			evt.WithLightID(lightid.NewFromInterval(1+c.offset, c.len+c.offset)),
		),
	}
}

func (c Custom) LightIDLen() int {
	return c.len
}

func (c Custom) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	return NewComposite(c.lightType, fn(lightid.NewFromInterval(1+c.offset, c.len+c.offset)))
}

func (c Custom) LightIDSequenceTransform(fn func(lightid.ID) lightid.Set) context.Light {
	sl := NewSequence()

	idSet := fn(lightid.NewFromInterval(1+c.offset, c.len+c.offset))

	for _, id := range idSet {
		sl.Add(NewComposite(c.lightType, lightid.NewSet(id)))
	}

	return sl
}
