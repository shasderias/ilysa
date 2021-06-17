package ilysa

import (
	"ilysa/pkg/ease"
	"ilysa/pkg/util"
)

func (c baseContext) eventForBeat(beat float64, callback func(TimingContext)) {
	callback(c.withTiming(beat, beat, beat, 0))
}

func (c baseContext) eventsForBeats(startBeat, duration float64, count int, callback func(TimingContext)) {
	startBeat += c.beatOffset

	endBeat := startBeat + (duration * float64(count-1))

	for i := 0; i < count; i++ {
		callback(c.withTiming(startBeat+duration*float64(i), startBeat, endBeat, i))
	}
}

func (c baseContext) eventsForRange(startBeat, endBeat float64, steps int, easeFunc ease.Func,
	callback func(TimingContext)) {

	startBeat += c.beatOffset

	tScaler := util.ScaleToUnitInterval(0, float64(steps-1))

	for i := 0; i < steps; i++ {
		beat := Ierp(startBeat, endBeat, tScaler(float64(i)), easeFunc)
		callback(c.withTiming(beat, startBeat, endBeat, i))
	}
}

func (c baseContext) eventsForSequence(startBeat float64, sequence []float64, callback func(ctx SequenceContext)) {
	if len(sequence) == 0 {
		panic("EventsForSequence: sequence must contain at least 1 beat")
	}

	startBeat += c.beatOffset

	endBeat := startBeat + sequence[len(sequence)-1]

	for i, offset := range sequence {
		beat := startBeat + offset
		callback(c.withTiming(beat, startBeat, endBeat, i).withSequence(sequence))
	}
}
