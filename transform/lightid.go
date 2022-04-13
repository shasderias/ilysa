package transform

import (
	"math/rand"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/internal/calc"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/lightid"
)

func Identity() LightTransformer {
	return newLightTransformer(func(l context.Light) context.Light {
		return l
	})
}

// DivideSingle divides each set of light IDs by the number of light IDs that set has.
// e.g. [1,2,3,4] -> [1], [2], [3], [4]
func DivideSingle() lightIDToLightIDSetTransformer {
	return newLightIDToLightIDSetTransformer(func(id lightid.ID) lightid.Set {
		set := lightid.NewSet()
		for _, i := range id {
			set.Add(lightid.New(i))
		}
		return set
	})
}

func divideSingle[T any](elems []T) [][]T {
	set := make([][]T, 0)
	for _, elem := range elems {
		set = append(set, []T{elem})
	}
	return set
}

// Divide divides each set of light IDs into divisor groups.
// e.g. [1,2,3,4,5] -> Divide(3) -> [1,2], [3,4], [5]
func Divide(divisor int) lightIDToLightIDSetTransformer {
	return newLightIDToLightIDSetTransformer(func(id lightid.ID) lightid.Set {
		divisor := divisor

		groupSize := len(id) / divisor

		set := lightid.NewSet()
		for i := 0; i < divisor; i++ {
			set.Add(id[0:groupSize])
			id = id[groupSize:]
		}
		set[divisor-1] = append(set[divisor-1], id...)

		return set
	})
}

// Fan divides each set of light IDs into groupCount groups with successive light IDs being allocated to
// successive groups (wrapping around to the first group when the last group is reached) until all light IDs
// have been allocated.
// e.g. [1,2,3,4,5,6,7,8] -> Fan(2) -> [1,3,5,7], [2,4,6,8]
// e.g. [1,2,3,4,5,6,7,8] -> Fan(3) -> [1,4,7], [2,5,8], [3,6]
func Fan(groupCount int) lightIDToLightIDSetTransformer {
	return newLightIDToLightIDSetTransformer(func(id lightid.ID) lightid.Set {
		set := make(lightid.Set, groupCount)

		for i := range set {
			set[i] = lightid.New()
		}

		for i, lightID := range id {
			set[i%groupCount] = append(set[i%groupCount], lightID)
		}

		return set
	})
}

func DivideIntoGroups(groupSize int) lightIDToLightIDSetTransformer {
	return newLightIDToLightIDSetTransformer(func(id lightid.ID) lightid.Set {
		set := lightid.NewSet()

		for len(id) > groupSize {
			set.Add(id[0:groupSize])
			id = id[groupSize:]
		}

		set.Add(lightid.New(id...))

		return set
	})
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

type shuffleSet struct {
	sequence bool
}

func ShuffleSet() shuffleSet {
	return shuffleSet{}
}

func (s shuffleSet) do(oldSet lightid.Set) lightid.Set {
	set := lightid.NewSet(oldSet...)

	rand.Shuffle(len(set), func(i, j int) {
		set[i], set[j] = set[j], set[i]
	})
	return set
}

func (s shuffleSet) LightTransform(l context.Light) context.Light {
	return applyLightIDSetTransformer(l, s.do, s.sequence)
}

func (s shuffleSet) Sequence() shuffleSet {
	return shuffleSet{true}
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
