package colorful

import (
	"math/rand"

	"github.com/shasderias/ilysa/internal/calc"
)

// A Set is an ordered set of Colors. Set provides convenience methods when
// working with colors such returning a random color (Set.Rand()) and returning
// successive colors from a Set (Set.Next()).
type Set struct {
	colors []Color
	i      *int
}

// NewSet creates a Set containing colors.
func NewSet(colors ...Color) Set {
	c := make([]Color, 0, len(colors))
	c = append(c, colors...)
	i := 0
	return Set{
		colors: c,
		i:      &i,
	}
}

// Index returns the ith color. If i is out of range, i is "wrapped around" to
// bring it back in range.
func (s Set) Index(i int) Color {
	i = calc.WraparoundIdx(len(s.colors), i)
	return s.colors[i]
}

// Idx returns the ith color. Idx is syntax sugar for Index.
func (s Set) Idx(i int) Color {
	return s.Index(i)
}

// Next returns successive colors from Set, starting with the first color and
// wrapping back to the first color after the last color.
func (s Set) Next() Color {
	c := s.colors[*s.i]
	*s.i++
	if *s.i == len(s.colors) {
		*s.i = 0
	}
	return c
}

// Rand returns as random color from Set.
func (s Set) Rand() Color {
	return s.colors[rand.Intn(len(s.colors))]
}

// Colors returns a slice of the colors in the Set.
func (s Set) Colors() []Color {
	return s.colors
}
