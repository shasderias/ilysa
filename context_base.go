package ilysa

import (
	"math/rand"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/scale"
)

type baseContext struct {
	*Project
	timer

	beatOffset float64
	fixedRand  float64
	modifiers  []EventModifier
}

func newBaseContext(p *Project) baseContext {
	return baseContext{
		Project:   p,
		fixedRand: rand.Float64(),
	}
}

func (c baseContext) FixedRand() float64 {
	return c.fixedRand
}

func (c baseContext) LightIDMax(typ beatsaber.EventType) int {
	return c.Project.ActiveDifficultyProfile().LightIDMax(typ)
}

func (c baseContext) withTimer(beat, startBeat, endBeat float64, ordinal int) baseContext {
	return baseContext{
		Project: c.Project,
		timer:   newTimer(beat, startBeat, endBeat, ordinal),

		beatOffset: c.beatOffset,
		fixedRand:  c.fixedRand,
		modifiers:  c.modifiers,
	}
}

func (c baseContext) withBeatOffset(o float64) baseContext {
	return baseContext{
		Project: c.Project,
		timer:   c.timer,

		beatOffset: o,
		fixedRand:  c.fixedRand,
		modifiers:  c.modifiers,
	}
}

func (c baseContext) WithBeatOffset(o float64) BaseContext {
	return baseContext{
		Project: c.Project,
		timer:   c.timer,

		beatOffset: c.beatOffset + o,
		fixedRand:  c.fixedRand,
		modifiers:  c.modifiers,
	}
}

func (c baseContext) WithModifier(modifiers ...EventModifier) baseContext {
	return baseContext{
		Project: c.Project,
		timer:   c.timer,

		beatOffset: c.beatOffset,
		fixedRand:  c.fixedRand,
		modifiers:  append([]EventModifier{}, modifiers...),
	}
}

func (c baseContext) WithLight(l Light, callback func(ctx TimeLightContext)) {
	for i := 0; i < l.LightIDLen(); i++ {
		callback(newTimeLightContext(c, l, i))
	}
}

func (c baseContext) addEvent(e Event) {
	c.applyModifiers(e)
	c.Project.events = append(c.Project.events, e)
}

func (c baseContext) applyModifiers(e Event) {
	if len(c.modifiers) == 0 {
		return
	}

	for _, m := range c.modifiers {
		m(e)
	}
}

func (c baseContext) EventForBeat(beat float64, callback func(TimeContext)) {
	beat += c.beatOffset
	callback(c.withTimer(beat, beat, beat, 0).withBeatOffset(0))
}

func (c baseContext) EventsForBeats(startBeat, duration float64, count int, callback func(TimeContext)) {
	startBeat += c.beatOffset

	endBeat := startBeat + (duration * float64(count-1))

	for i := 0; i < count; i++ {
		callback(c.withTimer(startBeat+duration*float64(i), startBeat, endBeat, i).withBeatOffset(0))
	}
}

func (c baseContext) EventsForRange(startBeat, endBeat float64, steps int, easeFunc ease.Func, callback func(TimeContext)) {
	startBeat += c.beatOffset
	endBeat += c.beatOffset

	tScaler := scale.ToUnitIntervalClamped(0, float64(steps-1))

	for i := 0; i < steps; i++ {
		beat := Ierp(startBeat, endBeat, tScaler(float64(i)), easeFunc)
		callback(c.withTimer(beat, startBeat, endBeat, i).withBeatOffset(0))
	}
}

func (c baseContext) EventsForSequence(startBeat float64, sequence []float64, callback func(ctx SequenceContext)) {
	if len(sequence) == 0 {
		panic("EventsForSequence: sequence must contain at least 1 beat")
	}

	startBeat += c.beatOffset

	endBeat := startBeat + sequence[len(sequence)-1]

	for i, offset := range sequence {
		beat := startBeat + offset
		callback(newSequenceContext(
			c.
				withTimer(beat, startBeat, endBeat, i).
				withBeatOffset(0),
			sequence,
		))
	}
}

func (c baseContext) ModEventsInRange(startBeat, endBeat float64, filter EventFilter, modder func(ctx TimeContext, event Event)) {
	p := c.Project
	p.sortEvents()

	startBeat += c.beatOffset
	endBeat += c.beatOffset

	startIdx, endIdx := 0, len(p.events)

	for i := 0; i < len(p.events); i++ {
		if p.events[i].Base().Beat >= startBeat {
			startIdx = i
			goto startFound
		}
	}
	// past last event
	return
startFound:

	for i := len(p.events) - 1; i >= startIdx; i-- {
		if p.events[i].Base().Beat <= endBeat {
			endIdx = i
			break
		}
	}

	events := p.events[startIdx : endIdx+1]

	for i := range events {
		if !filter(events[i]) {
			continue
		}
		modder(c.withTimer(events[i].Base().Beat, startBeat, endBeat, i).withBeatOffset(0), events[i])
	}
}

func (c baseContext) DeleteEvents(startBeat float64, filter EventFilter) {
	p := c.Project
	p.sortEvents()

	startBeat += c.beatOffset

	startIdx := 0

	for i := 0; i < len(p.events); i++ {
		if p.events[i].Base().Beat >= startBeat {
			startIdx = i
			goto startFound
		}
	}
	// past last event
	return
startFound:

	events := p.events[:startIdx]

	for _, e := range p.events[startIdx:] {
		if filter(e) {
			events = append(events, e)
		}
	}

	p.events = events
}
