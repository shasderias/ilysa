package gradient

import (
	"math/rand"

	"github.com/shasderias/ilysa/internal/calc"
)

// A Set is an ordered set of gradients. Set provides convenience methods such as
// picking a random gradient from a Set (Set.Rand()) and picking successive
// gradients from a Set (Set.Next()).
type Set struct {
	gradients []Table
	i         *int
}

// NewSet returns a new Set containing gradients.
func NewSet(gradients ...Table) Set {
	i := 0
	return Set{
		gradients: append([]Table{}, gradients...),
		i:         &i,
	}
}

// Index returns the gradient at the ith position. If i is out of range, i is
// "wrapped around" to bring it back in range.
func (s Set) Index(i int) Table {
	i = calc.WraparoundIdx(len(s.gradients), i)
	return s.gradients[i]
}

// Idx returns the gradient at the ith position. Idx is syntax sugar for Index.
func (s Set) Idx(i int) Table {
	return s.Index(i)
}

// Next return successive gradients from the Set each time it is called,
// starting with the first gradient.
func (s Set) Next() Table {
	t := s.Index(*s.i)
	*s.i++
	if *s.i == len(s.gradients) {
		*s.i = 0
	}
	return t
}

// Rand returns a random gradient from the Set.
func (s Set) Rand() Table {
	return s.gradients[rand.Intn(len(s.gradients))]
}
