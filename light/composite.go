package light

import (
	"fmt"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
)

type Composite struct {
	t   evt.Type
	set lightid.Set
	uid string
}

func NewComposite(t evt.Type, set lightid.Set) Composite {
	return Composite{t, set, getLightUID()}
}

func (l Composite) GenerateEvents(ctx context.LightContext) evt.Events {
	return evt.NewEvents(
		evt.NewChromaLighting(ctx,
			evt.OptType(l.t),
			evt.OptLightID(l.set.Index(ctx.LightOrdinal())),
		),
	)
}

func (l Composite) LightLen() int {
	return l.set.Len()
}

func (l Composite) Name() []string {
	return []string{fmt.Sprintf("Composite-%d-%s", l.t, l.uid)}
}

func (l Composite) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	set := lightid.NewSet()

	for _, id := range l.set {
		set.Add(fn(id)...)
	}

	return NewComposite(l.t, set)
}

func (l Composite) LightIDSequenceTransform(fn func(lightid.ID) lightid.Set) context.Light {
	set := lightid.NewSet()

	for _, id := range l.set {
		set.Add(fn(id)...)
	}

	seqLight := NewSequence()
	for _, id := range set {
		seqLight.Add(NewComposite(l.t, lightid.NewSet(id)))
	}
	return seqLight
}

func (l Composite) LightIDSetTransform(fn func(lightid.Set) lightid.Set) context.Light {
	return NewComposite(l.t, fn(l.set))
}

func (l Composite) LightIDSetSequenceTransform(fn func(lightid.Set) lightid.Set) context.Light {
	set := fn(l.set)

	seqLight := NewSequence()
	for _, id := range set {
		seqLight.Add(NewComposite(l.t, lightid.NewSet(id)))
	}

	return seqLight
}
