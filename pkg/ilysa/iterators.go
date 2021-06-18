package ilysa

import (
	"github.com/shasderias/ilysa/pkg/ease"
	"github.com/shasderias/ilysa/pkg/util"
)

func (c baseContext) EventForBeat(beat float64, callback func(TimingContext)) {
	beat += c.beatOffset
	callback(c.withTiming(beat, beat, beat, 0).WithBeatOffset(0))
}

func (c baseContext) EventsForBeats(startBeat, duration float64, count int, callback func(TimingContext)) {
	startBeat += c.beatOffset

	endBeat := startBeat + (duration * float64(count-1))

	for i := 0; i < count; i++ {
		callback(c.withTiming(startBeat+duration*float64(i), startBeat, endBeat, i).WithBeatOffset(0))
	}
}

func (c baseContext) EventsForRange(startBeat, endBeat float64, steps int, easeFunc ease.Func, callback func(TimingContext)) {
	startBeat += c.beatOffset
	endBeat += c.beatOffset


	tScaler := util.ScaleToUnitInterval(0, float64(steps-1))

	for i := 0; i < steps; i++ {
		beat := Ierp(startBeat, endBeat, tScaler(float64(i)), easeFunc)
		callback(c.withTiming(beat, startBeat, endBeat, i).WithBeatOffset(0))
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
		callback(c.withTiming(beat, startBeat, endBeat, i).WithBeatOffset(0).withSequence(sequence))
	}
}

func (c baseContext) ModEventsInRange(startBeat, endBeat float64, filter EventFilter, modder EventModder) {
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
		modder(c.withTiming(events[i].Base().Beat, startBeat, endBeat, i).WithBeatOffset(0), events[i])
	}
}
