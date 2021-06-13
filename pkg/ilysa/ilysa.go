package ilysa

import (
	"fmt"
	"sort"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/ease"
	"ilysa/pkg/util"
)

type Project struct {
	*beatsaber.Map

	events []Event
}

func New(bsMap *beatsaber.Map) *Project {
	return &Project{
		Map:    bsMap,
		events: []Event{},
	}
}

func (p *Project) EventsForRange(startBeat, endBeat float64, steps int, easeFunc ease.Func, generator EventGenerator) {
	ctx := Context{Project: p}
	posScaler := util.ScaleToUnitInterval(0, float64(steps-1))

	for i := 0; i < steps; i++ {
		pos := posScaler(float64(i))
		beat := Ierp(startBeat, endBeat, pos, easeFunc)
		ctx.Timing = NewTiming(startBeat, endBeat, beat, i)
		generator(&ctx)
	}
}

func (p *Project) EventForBeat(targetBeat float64, generator EventGenerator) {
	ctx := Context{Project: p}
	ctx.Timing = NewTiming(targetBeat, targetBeat, targetBeat, 0)
	generator(&ctx)
}

func (p *Project) EventsForBeats(startBeat, duration float64, count int, generator EventGenerator) {
	ctx := Context{Project: p}
	endBeat := startBeat + (duration * float64(count-1))

	for i := 0; i < count; i++ {
		ctx.Timing = NewTiming(startBeat, endBeat, startBeat+duration*float64(i), i)
		generator(&ctx)
	}
}

func (p *Project) EventsForSequence(startBeat float64, sequence []float64, generator EventGenerator) {
	ctx := Context{Project: p}

	if len(sequence) == 0 {
		panic("EventsForSequence: sequence must contain at least 1 beat")
	}

	endBeat := startBeat + sequence[len(sequence)-1]

	for i, offset := range sequence {
		beat := startBeat + offset
		ctx.Timing = NewTiming(startBeat, endBeat, beat, i)
		generator(&ctx)
	}
}

func (p *Project) ModEventsInRange(startBeat, endBeat float64, filter EventFilter, modder EventModder) {
	p.sortEvents()

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

	ctx := Context{Project: p}

	events := p.events[startIdx : endIdx+1]

	for i := range events {
		if !filter(events[i]) {
			continue
		}
		ctx.Timing = NewTiming(startBeat, endBeat, events[i].Base().Beat, i)
		modder(&ctx, events[i])
	}
}

func (p *Project) Save() error {
	events := []beatsaber.Event{}

	p.sortEvents()

	for _, e := range p.events {
		be := e.Base()
		event := beatsaber.Event{
			Time:  p.Map.UnscaleTime(be.Beat),
			Type:  be.Type,
			Value: be.Value,
		}

		cd, err := e.CustomData()
		if err != nil {
			return err
		}
		event.CustomData = cd

		events = append(events, event)
	}

	fmt.Printf("generated %d events\n", len(events))

	return p.Map.SaveEvents(events)
}

func (p *Project) sortEvents() {
	sort.Slice(p.events, func(i, j int) bool {
		return p.events[i].Base().Beat < p.events[j].Base().Beat
	})
}
