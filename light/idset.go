package light

type IDSet []ID

func NewIDSet(ids ...ID) IDSet {
	set := append(IDSet{}, ids...)
	return set
}

func (s *IDSet) Add(id ...ID) {
	*s = append(*s, id...)
}

func (s *IDSet) AppendToIndex(index int, id ...int) {
	(*s)[index] = append((*s)[index], id...)
}

func (s IDSet) Index(i int) ID {
	n := len(s)
	return s[((i%n)+n)%n]
}

func (s IDSet) Len() int {
	return len(s)
}
