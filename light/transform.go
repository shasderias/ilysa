package light

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/lightid"
)

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
