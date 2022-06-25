// Package opt provides handy helper methods for options.
package opt

import (
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

type Context interface {
	First() bool
	Last() bool

	SeqFirst() bool
	SeqLast() bool
}

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

func Successive[T any](ctx context.Context, a ...T) T {
	return a[ctx.Ordinal()%len(a)]
}

// LastOnly returns an option that applies opts only if the current beat is the
// last beat in the sequence/range.
func LastOnly(ctx Context, opts ...evt.Option) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if ctx.Last() {
			e.Apply(opts...)
		}
	})
}

// SeqLastOnly returns an option that applies opts only if the current beat is
// the last beat in the sequence.
func SeqLastOnly(ctx Context, opts ...evt.Option) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if ctx.SeqLast() {
			e.Apply(opts...)
		}
	})
}

// FirstOnly returns an option that applies opts only if the current beat is
// the first beat in the sequence/range.
func FirstOnly(ctx Context, opts ...evt.Option) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if ctx.First() {
			e.Apply(opts...)
		}
	})
}

// AllExceptLast returns an option that applies opts only if the current beat
// is not the last beat in the sequence/range.
func AllExceptLast(ctx Context, opts ...evt.Option) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if !ctx.Last() {
			e.Apply(opts...)
		}
	})
}

// AllExceptFirst returns an option that applies opts only if the current beat
// is not the first beat in the sequence/range.
func AllExceptFirst(ctx Context, opts ...evt.Option) evt.Option {
	return evt.NewFuncOpt(func(e evt.Event) {
		if !ctx.First() {
			e.Apply(opts...)
		}
	})
}

// BIfLast returns b if the current beat is the last beat in the sequence/range,
// otherwise it returns a.
func BIfLast[T any](ctx Context, a, b T) T {
	if ctx.Last() {
		return b
	}
	return a
}

// BIfSeqLast returns b if the current beat is the last beat in the sequence,
// otherwise it returns a.
func BIfSeqLast[T any](ctx interface{ SeqLast() bool }, a, b T) T {
	if ctx.SeqLast() {
		return b
	}
	return a
}
