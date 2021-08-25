package light

import (
	"math/rand"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/internal/calc"
	"github.com/shasderias/ilysa/lightid"
)

type Sequence struct {
	lights []context.Light
}

func NewSequence(lights ...context.Light) Sequence {
	return Sequence{append([]context.Light{}, lights...)}
}

func (s *Sequence) Add(lights ...context.Light) {
	s.lights = append(s.lights, lights...)
}

func (s Sequence) NewRGBLighting(ctx context.LightRGBLightingContext) evt.RGBLightingEvents {
	return s.Idx(ctx.SeqOrdinal()).NewRGBLighting(ctx)
}

func (s Sequence) LightIDLen() int {
	max := 0
	for _, l := range s.lights {
		if l.LightIDLen() > max {
			max = l.LightIDLen()
		}
	}
	return max
}

func (s Sequence) Idx(i int) context.Light {
	return s.lights[calc.WraparoundIdx(len(s.lights), i)]
}

func (s Sequence) Len() int {
	return len(s.lights)
}

func (s Sequence) Lights() []context.Light {
	return s.lights
}

func (s Sequence) Shuffle() Sequence {
	rand.Shuffle(len(s.lights), func(i, j int) {
		s.lights[i], s.lights[j] = s.lights[j], s.lights[i]
	})
	return Sequence{
		lights: s.lights,
	}
}

func (s Sequence) Slice(i, j int) Sequence {
	return Sequence{lights: s.lights[i:j]}
}

func (s Sequence) LightIDTransform(fn func(lightid.ID) lightid.Set) context.Light {
	seq := NewSequence()
	for _, l := range s.lights {
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
	for _, l := range s.lights {
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
	for _, l := range s.lights {
		tfl, ok := l.(LightIDSequenceTransformableLight)
		if !ok {
			seq.Add(l)
		} else {
			subSeq := tfl.LightIDSequenceTransform(fn).(Sequence)
			seq.Add(subSeq.lights...)
		}
	}
	return seq
}

func (s Sequence) LightIDSetSequenceTransform(fn func(lightid.Set) lightid.Set) context.Light {
	seq := NewSequence()
	for _, l := range s.lights {
		tfl, ok := l.(LightIDSetSequenceTransformableLight)
		if !ok {
			seq.Add(l)
		} else {
			subSeq := tfl.LightIDSetSequenceTransform(fn).(Sequence)
			seq.Add(subSeq.lights...)
		}
	}
	return seq
}
