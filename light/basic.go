package light

import (
	"fmt"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
)

// Basic represents a light with the base game's attributes. Lighting events
// created by Basic do not set _customData._lightID.
type Basic struct {
	lightType  evt.LightType
	maxLightID int
}

func NewBasic(ctx MaxLightIDer, t evt.LightType) Basic {
	return Basic{
		lightType:  t,
		maxLightID: ctx.MaxLightID(t),
	}
}

func (l Basic) NewRGBLighting(ctx context.LightRGBLightingContext) evt.RGBLightingEvents {
	return evt.RGBLightingEvents{
		ctx.NewRGBLighting(
			evt.WithLight(l.lightType),
		),
	}
}

func (l Basic) LightIDLen() int {
	return 1
}

func (l Basic) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	return NewComposite(l.lightType, fn(lightid.NewFromInterval(1, l.maxLightID)))
}

func (l Basic) LightIDSequenceTransform(fn func(lightid.ID) lightid.Set) context.Light {
	sl := NewSequence()

	idSet := fn(lightid.NewFromInterval(1, l.maxLightID))

	for _, id := range idSet {
		sl.Add(NewComposite(l.lightType, lightid.NewSet(id)))
	}

	return sl
}

func (l Basic) Name() []string {
	return []string{fmt.Sprintf("Basic-%d", l.lightType)}
}
