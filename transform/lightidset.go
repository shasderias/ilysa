package transform

//
//import (
//	"github.com/shasderias/ilysa/context"
//	"github.com/shasderias/ilysa/lightid"
//)
//
//type flatten struct{}
//
//func Flatten() flatten {
//	return flatten{}
//}
//
//func (f flatten) do(set lightid.Set) lightid.Set {
//	flattenedID := lightid.ID{}
//
//	for _, id := range set {
//		flattenedID = append(flattenedID, id...)
//	}
//
//	return lightid.NewSet(flattenedID)
//}
//
//func (f flatten) LightTransform(l context.Light) context.Light {
//	return applyLightIDSetTransformer(l, f.do, false)
//}
//
//type takeSet struct {
//	idx      []int
//	sequence bool
//}
//
//// TakeSet returns a transformer that takes (returns only) the light ID sets at
//// indexes.
////
//// [1,2,3], [4,5,6] ... [28,29,30]
//// TakeSet(0, 9)
//// [1,2,3], [28,29,30]
//func TakeSet(indexes ...int) takeSet {
//	return takeSet{indexes, false}
//}
//
//func (t takeSet) do(set lightid.Set) lightid.Set {
//	newSet := lightid.NewSet()
//	for _, idx := range t.idx {
//		newSet.Add(set.Index(idx))
//	}
//	return newSet
//}
//
//func (t takeSet) LightTransform(l context.Light) context.Light {
//	return applyLightIDSetTransformer(l, t.do, t.sequence)
//}
//
//func (t takeSet) Sequence() takeSet {
//	return takeSet{t.idx, true}
//}
//
//func TakeEvery(divsor, offset int) lightIDSetToLightIDSetTransformer {
//	return newLightIDSetToLightIDSetTransformer(func(set lightid.Set) lightid.Set {
//		newSet := lightid.NewSet()
//		for i, s := range set {
//			if i%divsor == offset {
//				newSet.Add(s)
//			}
//		}
//		return newSet
//	})
//}
//
//type sliceSet struct {
//	i, j     int
//	sequence bool
//}
//
//// SliceSet is a transformer that returns light IDs in the range [i:j).
////
//// [1,2,3], [4,5,6], [7,8,0]
//// SliceSet(0, 2)
//// [1,2,3], [4,5,6]
//func SliceSet(i, j int) sliceSet {
//	if i < 0 {
//		panic("transform>.SliceSet(): i must be 0 or greater")
//	}
//	if i > j {
//		panic("transform>.SliceSet(): i must be greater than j")
//	}
//	return sliceSet{i, j, false}
//
//}
//
//func (s sliceSet) do(set lightid.Set) lightid.Set {
//	j := s.j
//	if j > set.Len() {
//		j = set.Len()
//	}
//
//	return set.Slice(s.i, j)
//}
//
//func (s sliceSet) LightTransform(l context.Light) context.Light {
//	return applyLightIDSetTransformer(l, s.do, s.sequence)
//}
//
//func (s sliceSet) Sequence() sliceSet {
//	return sliceSet{s.i, s.j, true}
//}
