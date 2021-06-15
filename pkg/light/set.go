package light

type Set []ID

func NewSet(ids ...ID) *Set {
	set := append(Set{}, ids...)
	return &set
}

func (s Set) Index(i int) ID {
	n := len(s)
	return s[((i%n)+n)%n]
}

func (s *Set) Len() int {
	return len(*s)
}

func (s *Set) Add(id ID) {
	*s = append(*s, id)
}

func (s *Set) AppendToIndex(index int, id ...int) {
	(*s)[index] = append((*s)[index], id...)
}
