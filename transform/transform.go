package transform

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/lightid"
)

type LightTransformer interface {
	LightTransform(context.Light) context.Light
}

type LightIDTransformableLight interface {
	LightIDTransform(func(lightid.ID) lightid.Set) context.Light
}

type LightIDSetTransformableLight interface {
	LightIDSetTransform(func(lightid.Set) lightid.Set) context.Light
}

type LightIDSequenceTransformableLight interface {
	LightIDSequenceTransform(func(lightid.ID) lightid.Set) context.Light
}

type LightIDSetSequenceTransformableLight interface {
	LightIDSetSequenceTransform(func(lightid.Set) lightid.Set) context.Light
}

func Light(l context.Light, tfers ...LightTransformer) context.Light {
	for _, tf := range tfers {
		l = tf.LightTransform(l)
	}
	return l
}
func applyLightIDTransformer(l context.Light, fn func(id lightid.ID) lightid.Set, sequence bool) context.Light {
	if !sequence {
		transformableLight, ok := l.(LightIDTransformableLight)
		if !ok {
			return l
		}
		return transformableLight.LightIDTransform(fn)
	}

	transformableLight, ok := l.(LightIDSequenceTransformableLight)
	if !ok {
		return l
	}
	return transformableLight.LightIDSequenceTransform(fn)
}

func applyLightIDSetTransformer(l context.Light, fn func(set lightid.Set) lightid.Set, sequence bool) context.Light {
	if !sequence {
		transformableLight, ok := l.(LightIDSetTransformableLight)
		if !ok {
			return l
		}
		return transformableLight.LightIDSetTransform(fn)
	}

	transformableLight, ok := l.(LightIDSetSequenceTransformableLight)
	if !ok {
		return l
	}
	return transformableLight.LightIDSetSequenceTransform(fn)
}
