// Package transform provides functions for manipulating light ID based lights.
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

type lightTransformer struct {
	fn func(context.Light) context.Light
}

func newLightTransformer(fn func(l context.Light) context.Light) LightTransformer {
	return lightTransformer{fn}
}
func (tf lightTransformer) LightTransform(l context.Light) context.Light {
	return tf.fn(l)
}

type lightIDToLightIDSetTransformer struct {
	fn       func(set lightid.ID) lightid.Set
	sequence bool
}

func newLightIDToLightIDSetTransformer(fn func(set lightid.ID) lightid.Set) lightIDToLightIDSetTransformer {
	return lightIDToLightIDSetTransformer{fn, false}
}
func (tf lightIDToLightIDSetTransformer) LightTransform(light context.Light) context.Light {
	return applyLightIDTransformer(light, tf.fn, tf.sequence)
}
func (tf lightIDToLightIDSetTransformer) Sequence() {
	tf.sequence = true
}
func (tf lightIDToLightIDSetTransformer) do(id lightid.ID) lightid.Set {
	return tf.fn(id)
}

type lightIDSetToLightIDSetTransformer struct {
	fn       func(set lightid.Set) lightid.Set
	sequence bool
}

func newLightIDSetToLightIDSetTransformer(fn func(set lightid.Set) lightid.Set) lightIDSetToLightIDSetTransformer {
	return lightIDSetToLightIDSetTransformer{fn, false}
}
func (tf lightIDSetToLightIDSetTransformer) Sequence() {
	tf.sequence = true
}
func (tf lightIDSetToLightIDSetTransformer) LightTransform(light context.Light) context.Light {
	return applyLightIDSetTransformer(light, tf.fn, tf.sequence)
}
