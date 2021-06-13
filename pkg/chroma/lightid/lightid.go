package lightid

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/ilysa"
)

func For(ctx *ilysa.Context, typ beatsaber.EventType) chroma.LightID {
	max := ctx.ActiveDifficultyProfile().MaxLightID(typ)

	return FromInterval(1, max)
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

func Fan(lightIDs []int, groups int) LightIDSet {
	if len(lightIDs) < groups {
		panic("Fan: not enough lights to fan")
	}

	set := make(LightIDSet, groups)

	for i := 0; i < len(lightIDs); i++ {
		set[i%groups] = append(set[i%groups], lightIDs[i])
	}

	return set
}

func Divide(lightIDs []int, divisor int) LightIDSet {
	if len(lightIDs) < divisor {
		panic("Divide: not enough lights to divide")
	}

	setCount := len(lightIDs) / divisor
	set := LightIDSet{}

	for i := 0; i < divisor; i++ {
		set = append(set, lightIDs[setCount*i:setCount*(i+1)])
	}
	return set
}

type LightIDSet [][]int

func (s LightIDSet) Pick(n int) []int {
	return s[n%len(s)]
}
