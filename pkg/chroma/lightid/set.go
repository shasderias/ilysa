package lightid

import (
	"github.com/shasderias/ilysa/pkg/chroma"
)

type Set []chroma.LightID

func NewSet(lightIDs ...chroma.LightID) *Set {
	set := make(Set, 0)
	set = append(set, lightIDs...)
	return &set
}

func (s Set) Pick(i int) chroma.LightID {
	n := len(s)
	return s[((i%n)+n)%n]
}

func (s *Set) Len() int {
	return len(*s)
}

func (s *Set) Add(lid chroma.LightID) {
	*s = append(*s, lid)
}
