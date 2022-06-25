package light

import (
	"math/rand"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/internal/calc"
	"github.com/shasderias/ilysa/lightid"
)

type Sequence []context.Light

func NewSequence(lights ...context.Light) Sequence {
	return lights
}

func (s *Sequence) Add(lights ...context.Light) {
	*s = append(*s, lights...)
}

func (s Sequence) GenerateEvents(ctx context.LightContext) evt.Events {
	l := s.Idx(ctx.SeqOrdinal())

	return l.GenerateEvents(ctx)
}

func (s Sequence) LightLen() int {
	max := 0
	for _, l := range s {
		if l.LightLen() > max {
			max = l.LightLen()
		}
	}
	return max
}

func (s Sequence) Idx(i int) context.Light {
	return s[calc.WraparoundIdx(len(s), i)]
}

func (s Sequence) Len() int {
	return len(s)
}

func (s Sequence) Lights() []context.Light {
	return s
}

func (s Sequence) Shuffle() Sequence {
	clone := make(Sequence, len(s))
	copy(clone, s)
	rand.Shuffle(len(clone), func(i, j int) {
		clone[i], clone[j] = clone[j], clone[i]
	})
	return clone
}

func (s Sequence) Slice(i, j int) Sequence {
	return s[i:j]
}

func (s Sequence) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	seq := NewSequence()
	for _, l := range s {
		tfl, ok := l.(LightIDTransformableLight)
		if !ok {
			seq.Add(l)
		} else {
			seq.Add(tfl.LightIDTransform(fn))
		}
	}
	return seq
}

func (s Sequence) LightIDSetTransform(fn func(lightid.Set) lightid.Set) context.Light {
	seq := NewSequence()
	for _, l := range s {
		tfl, ok := l.(LightIDSetTransformableLight)
		if !ok {
			seq.Add(l)
		} else {
			seq.Add(tfl.LightIDSetTransform(fn))
		}
	}
	return seq
}

func (s Sequence) LightIDSequenceTransform(fn func(lightid.ID) lightid.Set) context.Light {
	seq := NewSequence()
	for _, l := range s {
		tfl, ok := l.(LightIDSequenceTransformableLight)
		if !ok {
			seq.Add(l)
		} else {
			subSeq := tfl.LightIDSequenceTransform(fn).(Sequence)
			seq.Add(subSeq...)
		}
	}
	return seq
}

func (s Sequence) LightIDSetSequenceTransform(fn func(lightid.Set) lightid.Set) context.Light {
	seq := NewSequence()
	for _, l := range s {
		tfl, ok := l.(LightIDSetSequenceTransformableLight)
		if !ok {
			seq.Add(l)
		} else {
			subSeq := tfl.LightIDSetSequenceTransform(fn).(Sequence)
			seq.Add(subSeq...)
		}
	}
	return seq
}

func (s Sequence) Name() []string {
	name := []string{}
	for _, l := range s {
		name = append(name, l.Name()...)
	}
	return name
}
