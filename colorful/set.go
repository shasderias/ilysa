package colorful

import (
	"math/rand"
)

type Set struct {
	colors []Color
	i      *int
}

//func (s *Set) Lerp(t float64, easeFunc ...ease.Func) float64 {
//	e := ease.Linear
//
//	switch len(easeFunc) {
//	case 0:
//		// do nothing
//	case 1:
//		e = easeFunc[0]
//	default:
//		panic("colorful.Set.Lerp: requires 0 or 1 easeFuncs")
//	}
//
//	return
//}

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
