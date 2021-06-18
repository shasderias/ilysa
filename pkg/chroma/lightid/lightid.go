package lightid

import (
	"github.com/shasderias/ilysa/pkg/beatsaber"
	"github.com/shasderias/ilysa/pkg/chroma"
)

type LightProfile interface {
	LightIDMin() int
	LightIDMax() int
}

type Picker func(profile LightProfile) Set

func AllIndividual(profile LightProfile) Set {
	set := Set{}

	for i := profile.LightIDMin(); i < profile.LightIDMax(); i++ {
		set.Add(chroma.LightID{i})
	}

	return set
}

func All(profile LightProfile) Set {
	set := Set{}

	set.Add(FromInterval(profile.LightIDMin(), profile.LightIDMax()))

	return set
}

func GroupDivide(divisor int) func(profile LightProfile) Set {
	return func(profile LightProfile) Set {
		allIDs := FromInterval(profile.LightIDMin(), profile.LightIDMax())

		return Divide(allIDs, divisor)
	}
}

func For(LightIDMax int, typ beatsaber.EventType) chroma.LightID {
	return FromInterval(1, LightIDMax)
}

func FromInterval(a, b int) chroma.LightID {
	lightIDs := make(chroma.LightID, b-a+1)
	for i := 0; i < len(lightIDs); i++ {
		lightIDs[i] = i + a
	}
	return lightIDs
}

func EveryNthLightID(min, max, div, remainder int) []int {
	lightIDs := make([]int, 0, (max-min)/div)
	for i := min; i <= max; i++ {
		if i%div == remainder {
			lightIDs = append(lightIDs, i)
		}
	}
	return lightIDs
}

func Fan(groups int) func(profile LightProfile) Set {
	return func(profile LightProfile) Set {
		min := profile.LightIDMin()
		max := profile.LightIDMax()
		if (max - min) < groups {
			panic("Fan: not enough lights to fan")
		}

		set := make(Set, groups)

		for i := min; i <= max; i++ {
			set[i%groups] = append(set[i%groups], i)
		}

		return set
	}
}

func Divide(lightIDs []int, divisor int) Set {
	if len(lightIDs) < divisor {
		panic("Divide: not enough lights to divide")
	}

	setCount := len(lightIDs) / divisor
	set := Set{}

	for i := 0; i < divisor; i++ {
		set = append(set, lightIDs[setCount*i:setCount*(i+1)])
	}
	return set
}
