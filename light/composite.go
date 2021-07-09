package light

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/lightid"
)

type Composite struct {
	t   evt.LightType
	set lightid.Set
}

func NewComposite(t evt.LightType, set lightid.Set) Composite {
	return Composite{t, set}
}

func (l Composite) NewRGBLighting(ctx context.LightRGBLightingContext) evt.RGBLightingEvents {
	return evt.RGBLightingEvents{
		ctx.NewRGBLighting(
			evt.WithLight(l.t),
			evt.WithLightID(l.set.Index(ctx.LightIDOrdinal())),
		),
	}
}

func (l Composite) LightIDLen() int {
	return l.set.Len()
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
