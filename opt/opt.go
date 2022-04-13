// Package opt provides handy helper methods for options.
package opt

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

type Set []evt.Option

func NewSet(opts ...evt.Option) *Set {
	set := Set{}
	set = append(Set{}, opts...)
	return &set
}

func (s *Set) Add(opts ...evt.Option) *Set {
	*s = append(*s, opts...)
	return s
}

func (s *Set) Apply(e evt.Event) {
	for _, o := range *s {
		o.Apply(e)
	}
}

func Successive(ctx context.Context, opts ...evt.Option) evt.Option {
	return opts[ctx.Ordinal()%len(opts)]
}

type Laster interface {
	Last() bool
}

// LastOnly returns an option that applies opts only if the current beat is the last beat in the sequence.
func LastOnly(ctx Laster, opts ...evt.Option) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if ctx.Last() {
			e.Apply(opts...)
		}
	})
}

// BIfLast returns b if the current beat is the last beat in the sequence, otherwise it returns a.
func BIfLast[T any](ctx Laster, a, b T) T {
	if ctx.Last() {
		return b
	}
	return a
}
