package light

import (
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

//func (l Basic) LightIDSeqTransform(tfer transform.Transformer) context.Light {
//
//}

//func (l Basic) LightIDSet() ilysa.LightIDSet {
//	return ilysa.NewLightIDSet(
//		ilysa.NewLightIDFromInterval(1, l.maxLightID),
//	)
//}
//
//func (l Basic) LightIDTransform(tfer ilysa.LightIDTransformer) Light {
//	return CompositeLight{
//		eventType: l.eventType,
//		set:       tfer(ilysa.NewLightIDFromInterval(1, l.maxLightID)),
//	}
//}
//
//func (l Basic) LightIDSequenceTransform(tfer ilysa.LightIDTransformer) Light {
//	sl := NewSequenceLight()
//	set := tfer(ilysa.NewLightIDFromInterval(1, l.maxLightID))
//
//	for _, lightID := range set {
//		sl.Add(CompositeLight{
//			eventType: l.eventType,
//			set:       ilysa.LightIDSet{lightID},
//		})
//
//	}
//	return sl
//}
//
//func (l Basic) LightIDSetTransform(tfer ilysa.LightIDSetTransformer) Light {
//	return CompositeLight{
//		eventType: l.eventType,
//		set:       tfer(l.LightIDSet()),
//	}
//}
//
//func (l Basic) LightIDSetSequenceTransform(tfer ilysa.LightIDSetTransformer) Light {
//	sl := NewSequenceLight()
//	set := tfer(l.LightIDSet())
//
//	for _, lightID := range set {
//		sl.Add(CompositeLight{
//			eventType: l.eventType,
//			set:       ilysa.LightIDSet{lightID},
//		})
//
//	}
//	return sl
//}
//
//func (l Basic) Transform(tfer ilysa.LightIDTransformer) CompositeLight {
//	return CompositeLight{
//		eventType: l.eventType,
//		set:       tfer(ilysa.NewLightIDFromInterval(1, l.maxLightID)),
//	}
//}
//
//func (l Basic) TransformToSequence(tfer ilysa.LightIDTransformer) SequenceLight {
//	sl := NewSequenceLight()
//	set := tfer(ilysa.NewLightIDFromInterval(1, l.maxLightID))
//
//	for _, lightID := range set {
//		sl.Add(CompositeLight{
//			eventType: l.eventType,
//			set:       ilysa.LightIDSet{lightID},
//		})
//
//	}
//	return sl
//}
