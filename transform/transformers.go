package transform

import (
	"math/rand"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/internal/calc"
	"github.com/shasderias/ilysa/lightid"
)

func Identity() LightTransformer {
	return newLightTransformer(func(l context.Light) context.Light {
		return l
	})
}

// DivideSingle divides each light ID into individual light IDs.
// e.g. [1,2,3,4] -> DivideSingle() -> [1], [2], [3], [4]
// e.g. [1,2], [3,4] -> DivideSingle() -> [1], [2], [3], [4]
func DivideSingle() lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightIDs := range set {
			for _, lightID := range lightIDs {
				s.Add(lightid.New(lightID))
			}
		}

		return s
	})
}

// Divide divides each light ID into divisor groups.
// e.g. [1,2,3,4,5] -> Divide(3) -> [1,2], [3,4], [5]
// e.g. [1,2,3,4], [5,6,7,8] -> Divide(2) -> [1,2], [3,4], [5,6], [7,8]
func Divide(divisor int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightIDs := range set {
			subSet := lightid.NewSet()

			groupSize := len(lightIDs) / divisor

			for i := 0; i < divisor; i++ {
				subSet.Add(lightIDs[0:groupSize])
				lightIDs = lightIDs[groupSize:]
			}
			subSet.AppendToIndex(subSet.Len()-1, lightIDs...)

			s.Add(subSet...)
		}

		return s
	})
}

// Fan divides each light ID into groupCount groups, the first id is allocated
// to the first group, the second id is allocated to the second group and so on,
// until all ids are allocated to a group. If Fan reaches the last group, it
// wraps around to the first group.
// e.g. [1,2,3,4,5,6,7,8] -> Fan(2) -> [1,3,5,7], [2,4,6,8]
// e.g. [1,2,3,4,5,6,7,8] -> Fan(3) -> [1,4,7], [2,5,8], [3,6]
// e.g. [1,2,3,4], [5,6,7,8] -> Fan(2) -> [1,3], [2,4], [5,7], [6,8]
func Fan(groupCount int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightIDs := range set {
			subSet := make(lightid.Set, groupCount)
			for i := range subSet {
				subSet[i] = lightid.New()
			}
			for i, lightID := range lightIDs {
				subSet[i%groupCount].Add(lightID)
			}
			s.Add(subSet...)
		}

		return s
	})
}

// DivideIntoGroups divides each light ID into groups such that each group has
// groupSize ids.
// e.g. [1,2,3,4,5,6,7,8] -> DivideIntoGroups(4) -> [1,2,3,4], [5,6,7,8]
// e.g. [1,2,3,4,5,6,7,8] -> DivideIntoGroups(3) -> [1,2,3], [4,5,6], [7,8]
// e.g. [1,2,3,4], [5,6,7,8] -> DivideIntoGroups(3) -> [1,2,3], [4], [5,6,7], [8]
func DivideIntoGroups(groupSize int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, id := range set {
			subSet := lightid.NewSet()

			for len(id) > groupSize {
				subSet.Add(id[0:groupSize])
				id = id[groupSize:]
			}

			subSet.Add(id)

			s.Add(subSet...)
		}

		return s
	})
}

func reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Reverse reverses the order of ids in each light ID.
// e.g. [1,2,3,4] -> Reverse() -> [4,3,2,1]
// e.g. [1,2], [3,4] -> Reverse() -> [2,1], [4,3]
func Reverse() lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightID := range set {
			newLightID := lightID.Clone()
			reverse(newLightID)
			s.Add(newLightID)
		}

		return s
	})
}

// ReverseSet reverses the order of the light's light IDs.
// e.g. [1,2,3,4] -> ReverseSet() -> [1,2,3,4]
// e.g. [1,2], [3,4], [5,6] -> ReverseSet() -> [5,6], [3,4], [1,2]
func ReverseSet() lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		newSet := set.Clone()
		reverse(newSet)
		return newSet
	})
}

// Shuffle randomizes the order of ids in each light ID.
// e.g. [1,2,3,4] -> Shuffle() -> [2,3,4,1] (random)
// e.g. [1,2,3], [4,5,6] -> Shuffle() -> [2,3,1], [5,6,4] (random)
func Shuffle(seed ...int64) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		shuffleFunc := rand.Shuffle

		if len(seed) > 0 {
			shuffleFunc = rand.New(rand.NewSource(seed[0])).Shuffle
		}

		s := lightid.NewSet()

		for _, lightID := range set {
			newLightID := lightID.Clone()
			shuffleFunc(len(newLightID), func(i, j int) {
				newLightID[i], newLightID[j] = newLightID[j], newLightID[i]
			})
			s.Add(newLightID)
		}

		return s
	})
}

// ShuffleSet randomizes the order of the light's light IDs.
// e.g. [1,2], [3,4], [5,6], [7,8] -> ShuffleSet() -> [3,4], [1,2], [7,8], [5,6] (random)
func ShuffleSet(seed ...int64) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		shuffleFunc := rand.Shuffle

		if len(seed) > 0 {
			shuffleFunc = rand.New(rand.NewSource(seed[0])).Shuffle
		}

		newSet := set.Clone()

		shuffleFunc(len(newSet), func(i, j int) {
			newSet[i], newSet[j] = newSet[j], newSet[i]
		})

		return newSet
	})
}

// Rotate rotates each id to the right by indices steps. Use negative indices to rotate left.
// e.g. [1,2,3,4] -> Rotate(2) -> [3,4,1,2]
// e.g. [1,2,3], [4,5,6] -> Rotate(2) -> [3,1,2], [6,4,5]
func Rotate(n int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightID := range set {
			normalizedN := len(lightID) - n%len(lightID)
			if normalizedN >= len(lightID) {
				normalizedN -= len(lightID)
			}

			newLightID := lightID.Clone()
			newLightID = append(newLightID[normalizedN:], newLightID[:normalizedN]...)
			s.Add(newLightID)
		}

		return s
	})
}

// RotateSet rotates the light's light IDs to the right by indices steps. Use negative indices to rotate left.
// e.g. [1,2], [3,4], [5,6], [7,8] -> RotateSet(2) -> [5,6], [7,8], [1,2], [3,4]
// e.g. [1,2], [3,4], [5,6], [7,8] -> RotateSet(-1) -> [3,4], [5,6], [7,8], [1,2]
func RotateSet(n int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		normalizedN := len(set) - n%len(set)
		if normalizedN >= len(set) {
			normalizedN -= len(set)
		}

		newSet := set.Clone()
		newSet = append(newSet[normalizedN:], newSet[:normalizedN]...)
		return newSet
	})
}

// Take, for each set, takes the ids at indicies.
// e.g. [1,2,3,4] -> Take(0,1,2) -> [1,2,3]
// e.g. [1,2], [3,4] -> Take(0) -> [1], [3]
func Take(indices ...int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightID := range set {
			takenIDs := lightid.New()
			for _, index := range indices {
				takenIDs.Add(calc.IndexWraparound(lightID, index))
			}
			s.Add(takenIDs)
		}

		return s
	})
}

// TakeSet takes the light IDs at indicies.
// e.g. [1,2], [3,4], [5,6], [7,8] -> TakeSet(0,1,2) -> [1,2], [3,4], [5,6]
// e.g. [1,2,3,4] -> TakeSet(0,1) -> [1,2,3,4], [1,2,3,4]
func TakeSet(indices ...int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, index := range indices {
			s.Add(set.Index(index).Clone())
		}

		return s
	})
}

// TakeEvery takes every nth id in each light ID, starting at offset.
// e.g. [1,2,3,4] -> TakeEvery(2,0) -> [1,3]
// e.g. [1,2,3,4] -> TakeEvery(2,1) -> [2,4]
func TakeEvery(n, offset int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightID := range set {
			takenIDs := lightid.New()
			for i, id := range lightID {
				if i%n == offset {
					takenIDs.Add(id)
				}
			}
			s.Add(takenIDs)
		}

		return s
	})
}

// TakeEverySet takes every nth light ID, starting at offset.
// e.g. [1,2], [3,4], [5,6], [7,8] -> TakeEverySet(2,0) -> [1,2], [5,6]
// e.g. [1,2], [3,4], [5,6], [7,8] -> TakeEverySet(2,1) -> [3,4], [7,8]
func TakeEverySet(n, offset int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for i, lightID := range set {
			if i%n == offset {
				s.Add(lightID)
			}
		}

		return s
	})
}

// Slice returns, for each light ID, the ids in the range [i:j).
// e.g. [1,2,3,4] -> Slice(1,3) -> [2,3]
// e.g. [1,2,3,4], [5,6,7,8] -> Slice(1,3) -> [2,3], [6,7]
func Slice(i, j int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		s := lightid.NewSet()

		for _, lightID := range set {
			s.Add(lightID.Clone()[i:j])
		}

		return s
	})
}

// SliceSet returns the light IDs in the range [i:j).
// e.g. [1,2], [3,4], [5,6], [7,8] -> SliceSet(1,3) -> [3,4], [5,6]
func SliceSet(i, j int) lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		return set.Clone()[i:j]
	})
}

// Flatten returns all the ids in one light ID.
// e.g. [1,2], [3,4] -> Flatten() -> [1,2,3,4]
func Flatten() lightIDTransformer {
	return newLightIDTransformer(func(set lightid.Set) lightid.Set {
		newID := lightid.New()

		for _, lightID := range set {
			for _, id := range lightID {
				newID.Add(id)
			}
		}

		return lightid.NewSet(newID)
	})
}
