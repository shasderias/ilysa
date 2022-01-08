package light

import (
	"fmt"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
)

type Custom struct {
	evtType evt.Type
	id      lightid.ID
}

func NewCustom(t evt.Type, len, offset int) Custom {
	return Custom{t, lightid.NewFromInterval(1+offset, len+offset)}
}

func NewCustomFromRange(t evt.Type, startID, endID int) Custom {
	return Custom{t, lightid.NewFromInterval(startID, endID)}
}

func NewCustomFromLightIDs(t evt.Type, ids ...int) Custom {
	return Custom{t, lightid.New(ids...)}
}

func (c Custom) GenerateEvents(ctx context.LightContext) evt.Events {
	return evt.NewEvents(
		evt.NewChromaLighting(ctx,
			evt.OptType(c.evtType), evt.OptLightID(c.id),
		),
	)
}

func (c Custom) LightLen() int {
	return 1
}

func (c Custom) Name() []string {
	return []string{fmt.Sprintf("Custom-%d-%v", c.evtType, c.id)}
}

func (c Custom) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	return NewComposite(c.evtType, fn(c.id))
}

func (c Custom) LightIDSequenceTransform(fn func(lightid.ID) lightid.Set) context.Light {
	sl := NewSequence()

	idSet := fn(c.id)

	for _, id := range idSet {
		sl.Add(NewComposite(c.evtType, lightid.NewSet(id)))
	}

	return sl
}
