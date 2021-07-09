package transform

import (
	"math/rand"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/lightid"
)

type identity struct{}

func Identity() identity {
	return identity{}
}

func (i identity) LightTransform(l context.Light) context.Light { return l }

func DivideSingle() divideSingle {
	return divideSingle{false}
}

type divideSingle struct {
	sequence bool
}

func (d divideSingle) do(id lightid.ID) lightid.Set {
	set := lightid.NewSet()
	for _, i := range id {
		set.Add(lightid.New(i))
	}
	return set
}

func (d divideSingle) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, d.do, d.sequence)
}

func (d divideSingle) Sequence() LightTransformer {
	return divideSingle{true}
}

// Divide returns a Transformer that divides a light ID into groupSize equal
// groups. If the light ID cannot be divided into groupSize equal groups, the
// remainder light IDs are placed in the last group.
func Divide(divisor int) divide {
	return divide{divisor, false}
}

type divide struct {
	divisor  int
	sequence bool
}

func (d divide) do(id lightid.ID) lightid.Set {
	divisor := d.divisor

	groupSize := len(id) / divisor

	set := lightid.NewSet()
	for i := 0; i < divisor; i++ {
		set.Add(id[0:groupSize])
		id = id[groupSize:]
	}
	set[divisor-1] = append(set[divisor-1], id...)

	return set
}

func (d divide) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, d.do, d.sequence)
}

func (d divide) Sequence() divide {
	return divide{d.divisor, true}
}

type fan struct {
	groupCount int
	sequence   bool
}

func Fan(groupCount int) fan {
	return fan{groupCount, false}
}

func (f fan) do(id lightid.ID) lightid.Set {
	groupCount := f.groupCount

	set := make(lightid.Set, groupCount)

	for i := range set {
		set[i] = lightid.New()
	}

	for i, lightID := range id {
		set[i%groupCount] = append(set[i%groupCount], lightID)
	}

	return set
}

func (f fan) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, f.do, f.sequence)
}

func (f fan) Sequence() fan {
	return fan{f.groupCount, true}
}

func DivideIntoGroups(groupSize int) divideIntoGroups {
	return divideIntoGroups{groupSize, false}
}

type divideIntoGroups struct {
	groupSize int
	sequence  bool
}

func (d divideIntoGroups) do(id lightid.ID) lightid.Set {
	groupSize := d.groupSize

	set := lightid.NewSet()

	for len(id) > groupSize {
		set.Add(id[0:groupSize])
		id = id[groupSize:]
	}

	set.Add(lightid.New(id...))

	return set
}

func (d divideIntoGroups) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, d.do, d.sequence)
}

func (d divideIntoGroups) Sequence() divideIntoGroups {
	return divideIntoGroups{d.groupSize, true}
}

type reverse struct {
	sequence bool
}

func Reverse() reverse {
	return reverse{false}
}

func (r reverse) do(id lightid.ID) lightid.Set {
	for i := len(id)/2 - 1; i >= 0; i-- {
		opp := len(id) - 1 - i
		id[i], id[opp] = id[opp], id[i]
	}

	return lightid.NewSet(id)
}

func (r reverse) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, r.do, r.sequence)
}

func (r reverse) Sequence() reverse {
	return reverse{true}
}

type reverseSet struct {
	sequence bool
}

func ReverseSet() reverseSet {
	return reverseSet{false}
}

func (r reverseSet) do(set lightid.Set) lightid.Set {
	for i := len(set)/2 - 1; i >= 0; i-- {
		opp := len(set) - 1 - i
		set[i], set[opp] = set[opp], set[i]
	}
	return set
}

func (r reverseSet) LightTransform(l context.Light) context.Light {
	return applyLightIDSetTransformer(l, r.do, r.sequence)
}

func (r reverseSet) Sequence() reverseSet {
	return reverseSet{true}
}

type shuffle struct {
	sequence bool
}

func Shuffle() shuffle {
	return shuffle{}
}

func (s shuffle) do(id lightid.ID) lightid.Set {
	rand.Shuffle(len(id), func(i, j int) {
		id[i], id[j] = id[j], id[i]
	})
	return lightid.NewSet(id)
}

func (s shuffle) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, s.do, s.sequence)
}

func (s shuffle) Sequence() shuffle {
	return shuffle{true}
}

//func Even(lightID ID) Set {
//	evenIDs := New()
//
//	for _, id := range lightID {
//		if id%2 == 0 {
//			evenIDs = append(evenIDs, id)
//		}
//	}
//
//	return NewSet(evenIDs)
//}

//func Odd(lightID lightid.ID) lightid.Set {
//	evenIDs := lightid.New()
//
//	for _, id := range lightID {
//		if id%2 == 1 {
//			evenIDs = append(evenIDs, id)
//		}
//	}
//
//	return lightid.NewSet(evenIDs)
//}
//
//
//func ToLightIDSetTransformer(tfer Transformer) SetTransformer {
//	return func(set lightid.Set) lightid.Set {
//		newSet := lightid.NewSet()
//
//		for _, id := range set {
//			newSet.Add(tfer(id)...)
//		}
//
//		return newSet
//	}
//}
//
////func divideSingle(id lightid.ID) lightid.Set {
////	set := lightid.NewSet()
////	for _, lightID := range id {
////		set.Add(lightid.ID{lightID})
////	}
////	return set
////}
//
//
//
type rotate struct {
	steps    int
	sequence bool
}

func Rotate(steps int) rotate {
	return rotate{steps, false}
}

func (r rotate) do(id lightid.ID) lightid.Set {
	steps := r.steps

	if steps > len(id) {
		steps = steps % len(id)
	}
	newID := append(lightid.ID{}, id[steps:]...)
	newID = append(newID, id[:steps]...)

	return lightid.NewSet(newID)
}

func (r rotate) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, r.do, r.sequence)
}

func (r rotate) Sequence() rotate {
	return rotate{r.steps, true}
}

//
//
//func Slice(ti, tj float64) Transformer {
//	return func(id lightid.ID) lightid.Set {
//		i := int(math.Round(ti * float64(len(id))))
//		j := int(math.Round(tj * float64(len(id))))
//
//		return lightid.NewSet(id[i:j])
//	}
//}
//
//
////type LightIDTransformableLight interface {
////	LightIDTransform(Transformer) ilysa.Light
////}
////
////func ToLightTransformer(tfer Transformer) ilysa.LightTransformer {
////	return func(l ilysa.Light) ilysa.Light {
////		tfl, ok := l.(LightIDTransformableLight)
////		if !ok {
////			return l
////		}
////		return tfl.LightIDTransform(tfer)
////	}
////}
////
////type LightIDSequenceTransformable interface {
////	LightIDSequenceTransform(Transformer) ilysa.Light
////}
////
////func ToSequenceLightTransformer(tfer Transformer) ilysa.LightTransformer {
////	return func(l ilysa.Light) ilysa.Light {
////		tfl, ok := l.(LightIDSequenceTransformable)
////		if !ok {
////			return l
////		}
////		return tfl.LightIDSequenceTransform(tfer)
////	}
////}
////
////type LightIDSetTransformable interface {
////	LightIDSetTransform(SetTransformer) ilysa.Light
////}
////
////func LightIDSetTransformerToLightTransformer(tfer SetTransformer) ilysa.LightTransformer {
////	return func(l ilysa.Light) ilysa.Light {
////		tfl, ok := l.(LightIDSetTransformable)
////		if !ok {
////			return l
////		}
////		return tfl.LightIDSetTransform(tfer)
////	}
////}
