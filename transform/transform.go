// Package transform implements functions to manipulate light ID based lights.
//
// Beat Saber lights are made up of individual lights.
package transform

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/lightid"
)

type LightTransformer interface {
	LightTransform(context.Light) context.Light
}
type lightIDTransformableLight interface {
	ApplyLightIDTransform(fn func(set lightid.Set) lightid.Set) context.Light
}
type lightIDSequenceTransformableLight interface {
	ApplyLightIDSequenceTransform(fn func(set lightid.Set) lightid.Set) context.Light
}

func Light(l context.Light, tfers ...LightTransformer) context.Light {
	for _, tf := range tfers {
		l = tf.LightTransform(l)
	}
	return l
}

type lightIDTransformer struct {
	fn       func(set lightid.Set) lightid.Set
	sequence bool
}

func newLightIDTransformer(fn func(set lightid.Set) lightid.Set) lightIDTransformer {
	return lightIDTransformer{fn, false}
}
func (tf lightIDTransformer) Sequence() LightTransformer {
	return lightIDTransformer{tf.fn, true}
}
func (tf lightIDTransformer) LightTransform(l context.Light) context.Light {
	if !tf.sequence {
		transformableLight, ok := l.(lightIDTransformableLight)
		if !ok {
			return l
		}
		return transformableLight.ApplyLightIDTransform(tf.fn)
	} else {
		transformableLight, ok := l.(lightIDSequenceTransformableLight)
		if !ok {
			return l
		}
		return transformableLight.ApplyLightIDSequenceTransform(tf.fn)
	}
}

type lightTransformer struct {
	fn func(l context.Light) context.Light
}

func newLightTransformer(fn func(l context.Light) context.Light) lightTransformer {
	return lightTransformer{fn}
}
func (tf lightTransformer) LightTransform(l context.Light) context.Light {
	return tf.fn(l)
}
