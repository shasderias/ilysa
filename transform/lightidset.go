package transform

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/lightid"
)

type flatten struct{}

func Flatten() flatten {
	return flatten{}
}

func (f flatten) do(set lightid.Set) lightid.Set {
	flattenedID := lightid.ID{}

	for _, id := range set {
		flattenedID = append(flattenedID, id...)
	}

	return lightid.NewSet(flattenedID)
}

func (f flatten) LightTransform(l context.Light) context.Light {
	return applyLightIDSetTransformer(l, f.do, false)
}
