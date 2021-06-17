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

func NewSet(ids ...LightID) *LightIDSet {
	set := append(LightIDSet{}, ids...)
	return &set
}

func (s LightIDSet) Index(i int) LightID {
	n := len(s)
	return s[((i%n)+n)%n]
}

func (s *LightIDSet) Len() int {
	return len(*s)
}

func (s *LightIDSet) Add(id LightID) {
	*s = append(*s, id)
}

func (s *LightIDSet) AppendToIndex(index int, id ...int) {
	(*s)[index] = append((*s)[index], id...)
}

type LightIDSplitter func(id LightID) *LightIDSet

func Identity(id LightID) *LightIDSet {
	return NewSet(id)
}

func DivideSingle(id LightID) *LightIDSet {
	set := NewSet()
	for _, lightID := range id {
		set.Add(LightID{lightID})
	}
	return set
}

func Divide(divisor int) LightIDSplitter {
	return func(id LightID) *LightIDSet {
		set := NewSet()

		for len(id) > divisor {
			set.Add(id[0:divisor])
			id = id[divisor:]
		}

		set.AppendToIndex(set.Len()-1, id...)

		return set
	}
}

func Fan(groupCount int) LightIDSplitter {
	return func(id LightID) *LightIDSet {
		set := make(LightIDSet, groupCount)

		for i, lightID := range id {
			set[i%groupCount] = append(set[i%groupCount], lightID)
		}

		return &set
	}
}
