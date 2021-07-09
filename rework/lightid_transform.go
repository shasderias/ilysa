package rework

import (
	"math"
	"math/rand"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/light"
)

type LightIDTransformer func(id LightID) LightIDSet
type LightIDSetTransformer func(set LightIDSet) LightIDSet

func Fan(groupCount int) LightIDTransformer {
	return func(id LightID) LightIDSet {
		set := make(LightIDSet, groupCount)

		for i := range set {
			set[i] = NewLightID()
		}

		for i, lightID := range id {
			set[i%groupCount] = append(set[i%groupCount], lightID)
		}

		return set
	}
}

func Identity(lightID LightID) LightIDSet {
	return NewLightIDSet(lightID)
}

func Even(lightID LightID) LightIDSet {
	evenIDs := NewLightID()

	for _, id := range lightID {
		if id%2 == 0 {
			evenIDs = append(evenIDs, id)
		}
	}

	return NewLightIDSet(evenIDs)
}

func Odd(lightID LightID) LightIDSet {
	evenIDs := NewLightID()

	for _, id := range lightID {
		if id%2 == 1 {
			evenIDs = append(evenIDs, id)
		}
	}

	return NewLightIDSet(evenIDs)
}

func ToLightIDSetTransformer(tfer LightIDTransformer) LightIDSetTransformer {
	return func(set LightIDSet) LightIDSet {
		newSet := NewLightIDSet()

		for _, id := range set {
			newSet.Add(tfer(id)...)
		}

		return newSet
	}
}

func DivideSingle(id LightID) LightIDSet {
	set := NewLightIDSet()
	for _, lightID := range id {
		set.Add(LightID{lightID})
	}
	return set
}

func Shuffle(id LightID) LightIDSet {
	rand.Shuffle(len(id), func(i, j int) {
		id[i], id[j] = id[j], id[i]
	})
	return NewLightIDSet(id)
}

// Divide returns a LightIDTransformer that divides a light ID into groupSize equal
// groups. If the light ID cannot be divided into groupSize equal groups, the
// remainder light IDs are placed in the last group.
func Divide(divisor int) LightIDTransformer {
	return func(id LightID) LightIDSet {
		groupSize := len(id) / divisor

		set := NewLightIDSet()
		for i := 0; i < divisor; i++ {
			set.Add(id[0:groupSize])
			id = id[groupSize:]
		}
		set[divisor-1] = append(set[divisor-1], id...)

		return set
	}
}

func DivideIntoGroupsOf(groupSize int) LightIDTransformer {
	return func(id LightID) LightIDSet {
		set := NewLightIDSet()

		for len(id) > groupSize {
			set.Add(id[0:groupSize])
			id = id[groupSize:]
		}

		set.Add(NewLightID(id...))

		return set
	}
}

func Rotate(steps int) LightIDTransformer {
	return func(id LightID) LightIDSet {
		if steps > len(id) {
			steps = steps % len(id)
		}
		newID := append(LightID{}, id[steps:]...)
		newID = append(newID, id[:steps]...)

		return NewLightIDSet(newID)
	}
}

func Reverse(id LightID) LightIDSet {
	for i := len(id)/2 - 1; i >= 0; i-- {
		opp := len(id) - 1 - i
		id[i], id[opp] = id[opp], id[i]
	}

	return NewLightIDSet(id)
}
func ReverseSet(set LightIDSet) LightIDSet {
	for i := len(set)/2 - 1; i >= 0; i-- {
		opp := len(set) - 1 - i
		set[i], set[opp] = set[opp], set[i]
	}
	return set
}

func Slice(ti, tj float64) LightIDTransformer {
	return func(id LightID) LightIDSet {
		i := int(math.Round(ti * float64(len(id))))
		j := int(math.Round(tj * float64(len(id))))

		return NewLightIDSet(id[i:j])
	}
}

func Flatten(set LightIDSet) LightIDSet {
	flattenedID := LightID{}

	for _, id := range set {
		flattenedID = append(flattenedID, id...)
	}

	return NewLightIDSet(flattenedID)
}

type SingleLight interface {
	EventType() beatsaber.EventType
	LightIDSet() LightIDSet
}

type LightIDTransformable interface {
	LightIDTransform(LightIDTransformer) light.Light
}

func ToLightTransformer(tfer LightIDTransformer) light.LightTransformer {
	return func(l light.Light) light.Light {
		tfl, ok := l.(LightIDTransformable)
		if !ok {
			return l
		}
		return tfl.LightIDTransform(tfer)
	}
}

type LightIDSequenceTransformable interface {
	LightIDSequenceTransform(LightIDTransformer) light.Light
}

func ToSequenceLightTransformer(tfer LightIDTransformer) light.LightTransformer {
	return func(l light.Light) light.Light {
		tfl, ok := l.(LightIDSequenceTransformable)
		if !ok {
			return l
		}
		return tfl.LightIDSequenceTransform(tfer)
	}
}

type LightIDSetTransformable interface {
	LightIDSetTransform(LightIDSetTransformer) light.Light
}

func LightIDSetTransformerToLightTransformer(tfer LightIDSetTransformer) light.LightTransformer {
	return func(l light.Light) light.Light {
		tfl, ok := l.(LightIDSetTransformable)
		if !ok {
			return l
		}
		return tfl.LightIDSetTransform(tfer)
	}
}
