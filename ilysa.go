package ilysa

import (
	"fmt"
	"sort"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/ease"
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

func (p *Project) WithBeatOffset(offset float64) BaseContext {
	ctx := newBaseContext(p)
	return ctx.withBeatOffset(offset)
}

func (p *Project) EventsForRange(startBeat, endBeat float64, steps int, easeFunc ease.Func, callback func(TimeContext)) {
	ctx := newBaseContext(p)
	ctx.EventsForRange(startBeat, endBeat, steps, easeFunc, callback)
}

func (p *Project) EventForBeat(beat float64, callback func(ctx TimeContext)) {
	ctx := newBaseContext(p)
	ctx.EventForBeat(beat, callback)
}

func (p *Project) EventsForBeats(startBeat, duration float64, count int, callback func(ctx TimeContext)) {
	ctx := newBaseContext(p)
	ctx.EventsForBeats(startBeat, duration, count, callback)
}

func (p *Project) EventsForSequence(startBeat float64, sequence []float64, callback func(ctx SequenceContext)) {
	ctx := newBaseContext(p)
	ctx.EventsForSequence(startBeat, sequence, callback)
}

func (p *Project) ModEventsInRange(startBeat, endBeat float64, filter EventFilter, modder func(ctx TimeContext, event Event)) {
	ctx := newBaseContext(p)
	ctx.ModEventsInRange(startBeat, endBeat, filter, modder)
}

func (p *Project) LightIDMax(typ beatsaber.EventType) int {
	return p.Map.ActiveDifficultyProfile().LightIDMax(typ)
}

func (p *Project) Save() error {
	events := []beatsaber.Event{}

	p.sortEvents()

	for _, e := range p.events {
		event := beatsaber.Event{
			Time:  p.Map.UnscaleTime(e.Beat()),
			Type:  e.Type(),
			Value: e.Value(),
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
		return p.events[i].Beat() < p.events[j].Beat()
	})
}
