package ilysa

type LightID []int

func NewLightID(ids ...int) LightID {
	return append(LightID{}, ids...)
}

func NewLightIDFromInterval(startID, endID int) LightID {
	lightID := make(LightID, 0, endID-startID+1)
	for i := startID; i <= endID; i++ {
		lightID = append(lightID, i)
	}
	return lightID
}

type LightIDSet []LightID

func NewLightIDSet(ids ...LightID) LightIDSet {
	set := append(LightIDSet{}, ids...)
	return set
}

func (s *LightIDSet) Add(id ...LightID) {
	*s = append(*s, id...)
}

func (s LightIDSet) Index(i int) LightID {
	n := len(s)
	return s[((i%n)+n)%n]
}

func (s *LightIDSet) Len() int {
	return len(*s)
}

func (s *LightIDSet) AppendToIndex(index int, id ...int) {
	(*s)[index] = append((*s)[index], id...)
}

func Identity(id LightID) LightIDSet {
	return NewLightIDSet(id)
}

func DivideSingle(id LightID) LightIDSet {
	set := NewLightIDSet()
	for _, lightID := range id {
		set.Add(LightID{lightID})
	}
	return set
}

func Divide(groups int) LightIDTransformer {
	return func(id LightID) LightIDSet {
		groupSize := len(id) / groups

		set := NewLightIDSet()
		for i := 0; i < groups; i++ {
			set.Add(id[0:groupSize])
			id = id[groupSize:]
		}
		set[groups-1] = append(set[groups-1], id...)

		return set
	}
}

func DivideIntoGroupsOf(divisor int) LightIDTransformer {
	return func(id LightID) LightIDSet {
		set := NewLightIDSet()

		for len(id) > divisor {
			set.Add(id[0:divisor])
			id = id[divisor:]
		}

		set.AppendToIndex(set.Len()-1, id...)

		return set
	}
}

func Fan(groupCount int) LightIDTransformer {
	return func(id LightID) LightIDSet {
		set := make(LightIDSet, groupCount)

		for i, lightID := range id {
			set[i%groupCount] = append(set[i%groupCount], lightID)
		}

		return set
	}
}
