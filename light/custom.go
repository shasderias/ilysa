package light

import (
	"fmt"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
)

type Custom struct {
	lightType evt.LightType
	id        lightid.ID
}

func NewCustom(t evt.LightType, len, offset int) Custom {
	return Custom{t, lightid.NewFromInterval(1+offset, len+offset)}
}

func NewCustomFromRange(t evt.LightType, startID, endID int) Custom {
	return Custom{t, lightid.NewFromInterval(startID, endID)}
}

func NewCustomFromLightIDs(t evt.LightType, ids ...int) Custom {
	return Custom{t, lightid.New(ids...)}
}

func (c Custom) NewRGBLighting(ctx context.LightRGBLightingContext) evt.RGBLightingEvents {
	return evt.RGBLightingEvents{
		ctx.NewRGBLighting(
			evt.WithLight(c.lightType),
			evt.WithLightID(c.id),
		),
	}
}

func (c Custom) LightIDLen() int {
	return 1
}

func (c Custom) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	return NewComposite(c.lightType, fn(c.id))
}

func (c Custom) LightIDSequenceTransform(fn func(lightid.ID) lightid.Set) context.Light {
	sl := NewSequence()

	idSet := fn(c.id)

	for _, id := range idSet {
		sl.Add(NewComposite(c.lightType, lightid.NewSet(id)))
	}

	return sl
}

func (c Custom) Name() []string {
	return []string{fmt.Sprintf("Custom-%d-%v", c.lightType, c.id)}
}
