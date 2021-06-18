package ilysa

import "ilysa/pkg/beatsaber"

type LightIDTransformer func(id LightID) LightIDSet

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

func DivideSingle(id LightID) LightIDSet {
	set := NewLightIDSet()
	for _, lightID := range id {
		set.Add(LightID{lightID})
	}
	return set
}

// Divide returns a LightIDTransformer that divides a light ID into groupSize equal
// groups. If the light ID cannot be divide into groupSize equal groups, the
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

type SingleLight interface {
	EventType() beatsaber.EventType
	LightIDSet() LightIDSet
}

type LightIDTransformable interface {
	LightIDTransform(LightIDTransformer) Light
}

func ToLightTransformer(tfer LightIDTransformer) LightTransformer {
	return func(l Light) Light {
		tfl, ok := l.(LightIDTransformable)
		if !ok { return l}
		return tfl.LightIDTransform(tfer)
	}
}

type LightIDSequenceTransformable interface {
	LightIDSequenceTransform(LightIDTransformer) Light
}

func ToSequenceLightTransformer(tfer LightIDTransformer) LightTransformer {
	return func(l Light) Light {
		tfl, ok := l.(LightIDSequenceTransformable)
		if !ok { return l}
		return tfl.LightIDSequenceTransform(tfer)
		//sl, ok := l.(SingleLight)
		//if !ok {
		//	return l
		//}
		//
		//seqLight := NewSequenceLight()
		//for _, l := range sl.LightIDSet() {
		//	set := tfer(l)
		//	for _, id := range set {
		//		seqLight.Add(NewCompositeLight(sl.EventType(), NewLightIDSet(id)))
		//	}
		//}
		//return seqLight
	}
}

//type LightIDTransformable interface {
//	LightIDTransform(LightIDTransformer) Light
//}

func FanBasicLight(groupCount int) func(bl BasicLight) []CompositeLight {
	return func(bl BasicLight) []CompositeLight {
		set := Fan(groupCount)(bl.LightIDSet()[0])

		lights := make([]CompositeLight, groupCount)

		for i := 0; i < groupCount; i++ {
			lights[i] = NewCompositeLight(bl.EventType(), NewLightIDSet(set[i]))
		}

		return lights
	}
}
