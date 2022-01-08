package transform

import (
	"math/rand"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/internal/calc"
	"github.com/shasderias/ilysa/light"
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
	id = lightid.New(id...)

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

func (r reverseSet) do(oldSet lightid.Set) lightid.Set {
	set := lightid.NewSet(oldSet...)

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

type shuffleSeed struct {
	sequence bool
	seed     int64
	randSrc  *rand.Rand
}

func ShuffleSeed(seed int64) shuffleSeed {
	return shuffleSeed{seed: seed, randSrc: rand.New(rand.NewSource(seed))}
}

func (s shuffleSeed) do(id lightid.ID) lightid.Set {
	s.randSrc.Shuffle(len(id), func(i, j int) {
		id[i], id[j] = id[j], id[i]
	})
	return lightid.NewSet(id)
}

func (s shuffleSeed) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, s.do, s.sequence)
}

func (s shuffleSeed) Sequence() shuffleSeed {
	return shuffleSeed{true, s.seed, s.randSrc}
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

type rotateSet struct {
	steps    int
	sequence bool
}

func RotateSet(steps int) rotateSet {
	return rotateSet{steps, false}
}

func (r rotateSet) do(set lightid.Set) lightid.Set {
	steps := r.steps

	if steps > len(set) {
		steps = steps % len(set)
	}
	newSet := append(lightid.Set{}, set[steps:]...)
	newSet = append(newSet, set[:steps]...)

	return newSet
}

func (r rotateSet) LightTransform(l context.Light) context.Light {
	return applyLightIDSetTransformer(l, r.do, r.sequence)
}

func (r rotateSet) Sequence() rotateSet {
	return rotateSet{r.steps, true}
}

type take struct {
	indices  []int
	sequence bool
}

// Take is a transformer that for each light ID, returns the value at indices
// idx.
//
// [1,2,3,4,5]
// Take(1,3,5)
// [1,3,5]
func Take(idx ...int) take {
	return take{idx, false}
}

func (t take) do(id lightid.ID) lightid.Set {
	l := len(id)
	newIDs := []int{}
	for _, idx := range id {
		newIDs = append(newIDs, id[calc.WraparoundIdx(l, idx)])
	}

	return lightid.NewSet(lightid.New(newIDs...))
}

func (t take) Sequence() take {
	return take{t.indices, true}

}

func (t take) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, t.do, t.sequence)
}

type slice struct {
	i, j     int
	sequence bool
}

// Slice is a transformer that for each light ID of the light, returns light IDs
// in the range [i:j).
//
// [1,2,3 ... 60]
// Slice(0, 10)
// [1,2,3 ... 10]
func Slice(i, j int) slice {
	if i < 0 {
		panic("transform.Slice(): i must be 0 or greater")
	}
	if i > j {
		panic("transform.Slice(): i must be smaller than j")
	}
	return slice{i, j, false}
}

func (s slice) do(id lightid.ID) lightid.Set {
	j := s.j
	if j > len(id) {
		j = len(id)
	}

	return lightid.NewSet(lightid.New(id[s.i:j]...))
}

func (s slice) Sequence() slice {
	return slice{s.i, s.j, true}
}

func (s slice) LightTransform(l context.Light) context.Light {
	return applyLightIDTransformer(l, s.do, s.sequence)
}

type sequenceIdx struct {
	i int
}

func SequenceIdx(i int) sequenceIdx {
	return sequenceIdx{i}
}

func (si sequenceIdx) LightTransform(l context.Light) context.Light {
	sl, ok := l.(light.Sequence)
	if !ok {
		return l
	}

	return sl.Idx(si.i)
}
