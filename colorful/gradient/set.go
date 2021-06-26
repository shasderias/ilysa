package gradient

import "math/rand"

type Set struct {
	gradients []Table
	i         int
}

func NewSet(gradients ...Table) Set {
	return Set{
		gradients: append([]Table{}, gradients...),
		i:         0,
	}
}

func (s Set) Index(ordinal int) Table {
	return s.gradients[ordinal]
}

func (s *Set) Next() Table {
	t := s.Index(s.i)
	s.i++
	if s.i == len(s.gradients) {
		s.i = 0
	}
	return t
}

func (s Set) Rand() Table {
	return s.gradients[rand.Intn(len(s.gradients))]
}
