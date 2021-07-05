package light

import (
	"math"
	"math/rand"

	"github.com/shasderias/ilysa/beatsaber"
)

type LightIDTransformer func(id ID) IDSet
type LightIDSetTransformer func(set IDSet) IDSet

func Fan(groupCount int) LightIDTransformer {
	return func(id ID) IDSet {
		set := make(IDSet, groupCount)

		for i := range set {
			set[i] = NewID()
		}

		for i, lightID := range id {
			set[i%groupCount] = append(set[i%groupCount], lightID)
		}

		return set
	}
}

func Identity(lightID ID) IDSet {
	return NewIDSet(lightID)
}

func Even(lightID ID) IDSet {
	evenIDs := NewID()

	for _, id := range lightID {
		if id%2 == 0 {
			evenIDs = append(evenIDs, id)
		}
	}

	return NewIDSet(evenIDs)
}

func Odd(lightID ID) IDSet {
	evenIDs := NewID()

	for _, id := range lightID {
		if id%2 == 1 {
			evenIDs = append(evenIDs, id)
		}
	}

	return NewIDSet(evenIDs)
}

func ToLightIDSetTransformer(tfer LightIDTransformer) LightIDSetTransformer {
	return func(set IDSet) IDSet {
		newSet := NewIDSet()

		for _, id := range set {
			newSet.Add(tfer(id)...)
		}

		return newSet
	}
}

func DivideSingle(id ID) IDSet {
	set := NewIDSet()
	for _, lightID := range id {
		set.Add(ID{lightID})
	}
	return set
}

func Shuffle(id ID) IDSet {
	rand.Shuffle(len(id), func(i, j int) {
		id[i], id[j] = id[j], id[i]
	})
	return NewIDSet(id)
}

// Divide returns a LightIDTransformer that divides a light ID into groupSize equal
// groups. If the light ID cannot be divided into groupSize equal groups, the
// remainder light IDs are placed in the last group.
func Divide(divisor int) LightIDTransformer {
	return func(id ID) IDSet {
		groupSize := len(id) / divisor

		set := NewIDSet()
		for i := 0; i < divisor; i++ {
			set.Add(id[0:groupSize])
			id = id[groupSize:]
		}
		set[divisor-1] = append(set[divisor-1], id...)

		return set
	}
}

func DivideIntoGroupsOf(groupSize int) LightIDTransformer {
	return func(id ID) IDSet {
		set := NewIDSet()

		for len(id) > groupSize {
			set.Add(id[0:groupSize])
			id = id[groupSize:]
		}

		set.Add(NewID(id...))

		return set
	}
}

func Rotate(steps int) LightIDTransformer {
	return func(id ID) IDSet {
		if steps > len(id) {
			steps = steps % len(id)
		}
		newID := append(ID{}, id[steps:]...)
		newID = append(newID, id[:steps]...)

		return NewIDSet(newID)
	}
}

func Reverse(id ID) IDSet {
	for i := len(id)/2 - 1; i >= 0; i-- {
		opp := len(id) - 1 - i
		id[i], id[opp] = id[opp], id[i]
	}

	return NewIDSet(id)
}

func ReverseSet(set IDSet) IDSet {
	for i := len(set)/2 - 1; i >= 0; i-- {
		opp := len(set) - 1 - i
		set[i], set[opp] = set[opp], set[i]
	}
	return set
}

func Slice(ti, tj float64) LightIDTransformer {
	return func(id ID) IDSet {
		i := int(math.Round(ti * float64(len(id))))
		j := int(math.Round(tj * float64(len(id))))

		return NewIDSet(id[i:j])
	}
}

func Flatten(set IDSet) IDSet {
	flattenedID := ID{}

	for _, id := range set {
		flattenedID = append(flattenedID, id...)
	}

	return NewIDSet(flattenedID)
}

type SingleLight interface {
	EventType() beatsaber.EventType
	LightIDSet() IDSet
}

//type LightIDTransformable interface {
//	LightIDTransform(LightIDTransformer) ilysa.Light
//}
//
//func ToLightTransformer(tfer LightIDTransformer) ilysa.LightTransformer {
//	return func(l ilysa.Light) ilysa.Light {
//		tfl, ok := l.(LightIDTransformable)
//		if !ok {
//			return l
//		}
//		return tfl.LightIDTransform(tfer)
//	}
//}
//
//type LightIDSequenceTransformable interface {
//	LightIDSequenceTransform(LightIDTransformer) ilysa.Light
//}
//
//func ToSequenceLightTransformer(tfer LightIDTransformer) ilysa.LightTransformer {
//	return func(l ilysa.Light) ilysa.Light {
//		tfl, ok := l.(LightIDSequenceTransformable)
//		if !ok {
//			return l
//		}
//		return tfl.LightIDSequenceTransform(tfer)
//	}
//}
//
//type LightIDSetTransformable interface {
//	LightIDSetTransform(LightIDSetTransformer) ilysa.Light
//}
//
//func LightIDSetTransformerToLightTransformer(tfer LightIDSetTransformer) ilysa.LightTransformer {
//	return func(l ilysa.Light) ilysa.Light {
//		tfl, ok := l.(LightIDSetTransformable)
//		if !ok {
//			return l
//		}
//		return tfl.LightIDSetTransform(tfer)
//	}
//}
