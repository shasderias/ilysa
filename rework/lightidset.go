package rework

type LightIDSet []LightID

func NewLightIDSet(ids ...LightID) LightIDSet {
	set := append(LightIDSet{}, ids...)
	return set
}

func (s *LightIDSet) Add(id ...LightID) {
	*s = append(*s, id...)
}

func (s *LightIDSet) AppendToIndex(index int, id ...int) {
	(*s)[index] = append((*s)[index], id...)
}

func (s LightIDSet) Index(i int) LightID {
	n := len(s)
	return s[((i%n)+n)%n]
}

func (s LightIDSet) Len() int {
	return len(s)
}
