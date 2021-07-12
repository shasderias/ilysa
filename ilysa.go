package ilysa

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/timer"
)

type Project struct {
	Map

	events []evt.Event
}

type Map interface {
	ActiveDifficultyProfile() *beatsaber.EnvProfile
	SaveEvents([]beatsaber.Event) error
	SetActiveDifficulty(c beatsaber.Characteristic, difficulty beatsaber.BeatmapDifficulty) error
	UnscaleTime(beat float64) beatsaber.Time
}

func New(bsMap Map) *Project {
	return &Project{
		Map:    bsMap,
		events: []evt.Event{},
	}
}

func (p *Project) Offset(offset float64) context.Context {
	return context.Base(p).BOffset(offset)
}

func (p *Project) Range(r timer.Ranger, callback func(context.Context)) {
	context.Base(p).Range(r, callback)
}

func (p *Project) Sequence(s timer.Sequencer, callback func(ctx context.Context)) {
	context.Base(p).Sequence(s, callback)
}

func (p *Project) MaxLightID(t evt.LightType) int {
	return p.Map.ActiveDifficultyProfile().MaxLightID(beatsaber.EventType(t))
}

func (p *Project) AddEvents(events ...evt.Event) {
	p.events = append(p.events, events...)
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
