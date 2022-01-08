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
	evtType evt.Type
	len     int
}

func NewBasic(t evt.Type, len int) Basic {
	return Basic{t, len}
}

func (l Basic) GenerateEvents(ctx context.LightContext) evt.Events {
	return evt.NewEvents(
		evt.NewChromaLighting(ctx,
			evt.OptType(l.evtType),
		),
	)
}

func (l Basic) LightLen() int {
	return 1
}

func (l Basic) Name() []string {
	return []string{fmt.Sprintf("Basic-%d", l.evtType)}
}

func (l Basic) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	return NewComposite(l.evtType, fn(lightid.NewFromInterval(1, l.len)))
}

func (l Basic) LightIDSequenceTransform(fn func(lightid.ID) lightid.Set) context.Light {
	sl := NewSequence()

	idSet := fn(lightid.NewFromInterval(1, l.len))

	for _, id := range idSet {
		sl.Add(NewComposite(l.evtType, lightid.NewSet(id)))
	}

	return sl
}
