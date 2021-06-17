package ilysa

import (
	"fmt"
	"sort"

	"ilysa/pkg/beatsaber"
	"ilysa/pkg/ease"
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

func (p *Project) EventsForRange(startBeat, endBeat float64, steps int, easeFunc ease.Func, callback func(TimingContext)) {
	ctx := newBaseContext(p)
	ctx.eventsForRange(startBeat, endBeat, steps, easeFunc, callback)
}

func (p *Project) EventForBeat(beat float64, callback func(ctx TimingContext)) {
	ctx := newBaseContext(p)
	ctx.eventForBeat(beat, callback)
}

func (p *Project) EventsForBeats(startBeat, duration float64, count int, callback func(ctx TimingContext)) {
	ctx := newBaseContext(p)
	ctx.eventsForBeats(startBeat, duration, count, callback)
}

func (p *Project) EventsForSequence(startBeat float64, sequence []float64, callback func(ctx SequenceContext)) {
	ctx := newBaseContext(p)
	ctx.eventsForSequence(startBeat, sequence, callback)
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

	ctx := newBaseContext(p)

	events := p.events[startIdx : endIdx+1]

	for i := range events {
		if !filter(events[i]) {
			continue
		}
		modder(ctx.withTiming(events[i].Base().Beat, startBeat, endBeat, i), events[i])
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
