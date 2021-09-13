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

type takeSet struct {
	idx      []int
	sequence bool
}

// TakeSet returns a transformer that takes (returns only) the light ID sets at
// indexes.
//
// [1,2,3], [4,5,6] ... [28,29,30]
// TakeSet(0, 9)
// [1,2,3], [28,29,30]
func TakeSet(indexes ...int) takeSet {
	return takeSet{indexes, false}
}

func (t takeSet) do(set lightid.Set) lightid.Set {
	newSet := lightid.NewSet()
	for _, idx := range t.idx {
		newSet.Add(set.Index(idx))
	}
	return newSet
}

func (t takeSet) LightTransform(l context.Light) context.Light {
	return applyLightIDSetTransformer(l, t.do, t.sequence)
}

func (t takeSet) Sequence() takeSet {
	return takeSet{t.idx, true}
}
