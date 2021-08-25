package colorful

import (
	"math/rand"
)

type Set struct {
	colors []Color
	i      *int
}

func NewSet(colors ...Color) Set {
	c := make([]Color, 0, len(colors))
	c = append(c, colors...)
	i := 0
	return Set{
		colors: c,
		i:      &i,
	}
}

func (s Set) Idx(ordinal int) Color {
	return s.colors[ordinal%len(s.colors)]
}

func (s Set) Next() Color {
	c := s.colors[*s.i]
	*s.i++
	if *s.i == len(s.colors) {
		*s.i = 0
	}
	return c
}

func (s Set) Rand() Color {
	return s.colors[rand.Intn(len(s.colors))]
}

func (s Set) Colors() []Color {
	return s.colors
}
