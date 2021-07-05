package ilysa

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/ease"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type Project struct {
	*beatsaber.Map

	events []evt.Event
}

func New(bsMap *beatsaber.Map) *Project {
	return &Project{
		Map:    bsMap,
		events: []evt.Event{},
	}
}

func (p *Project) WithBeatOffset(offset float64) BaseContext {
	ctx := newBaseContext(p)
	return ctx.withBeatOffset(offset)
}

func (p *Project) Range(startBeat, endBeat float64, steps int, easeFunc ease.Func, callback func(RangeContext)) {
	ctx := newBaseContext(p)
	ctx.Range(startBeat, endBeat, steps, easeFunc, callback)
}

func (p *Project) Sequence(sequence timer.Sequencer, callback func(ctx SequenceContext)) {
	ctx := newBaseContext(p)
	ctx.Sequence(sequence, callback)
}

//func (p *Project) ModEventsInRange(startBeat, endBeat float64, filter EventFilter, modder func(ctx RangeContext, event Event)) {
//	ctx := newBaseContext(p)
//	ctx.ModEventsInRange(startBeat, endBeat, filter, modder)
//}

func (p *Project) LightIDMax(typ beatsaber.EventType) int {
	return p.Map.ActiveDifficultyProfile().LightIDMax(typ)
}

func (p *Project) sortEvents() {
	sort.Slice(p.events, func(i, j int) bool {
		return p.events[i].Beat() < p.events[j].Beat()
	})
}

func (p *Project) generateBeatSaberEvents() ([]beatsaber.Event, error) {
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
			return nil, err
		}
		event.CustomData = cd

		events = append(events, event)
	}

	return events, nil
}

func (p *Project) Save() error {
	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(events))

	return p.Map.SaveEvents(events)
}

func (p *Project) Dump() error {
	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	eventsJSON, err := json.Marshal(events)
	if err != nil {
		return err
	}

	fmt.Print(string(eventsJSON))

	return nil
}
