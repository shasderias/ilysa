package ilysa

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/internal/imath"
	"github.com/shasderias/ilysa/internal/swallowjson"
)

type Project struct {
	Map
	context.Context

	events []evt.Event
}

type Map interface {
	ActiveDifficultyProfile() *beatsaber.EnvProfile
	AppendEvents([]beatsaber.Event) error
	Events() []beatsaber.Event
	SaveEvents([]beatsaber.Event) error
	SetActiveDifficulty(c beatsaber.Characteristic, difficulty beatsaber.BeatmapDifficulty) error
	UnscaleTime(beat float64) beatsaber.Time
}

func New(bsMap Map) *Project {
	p := &Project{
		Map:    bsMap,
		events: []evt.Event{},
	}

	p.Context = context.Base(p)

	return p
}

func (p *Project) BOffset(offset float64) context.Context {
	return context.Base(p).BOffset(offset)
}

func (p *Project) MaxLightID(t evt.LightType) int {
	return p.Map.ActiveDifficultyProfile().MaxLightID(beatsaber.EventType(t))
}

func (p *Project) AddEvents(events ...evt.Event) {
	p.events = append(p.events, events...)
}

func (p *Project) Events() *[]evt.Event {
	return &p.events
}

func (p *Project) sortEvents() {
	sort.Slice(p.events, func(i, j int) bool {
		return p.events[i].Beat() < p.events[j].Beat()
	})
}

func (p *Project) generateBeatSaberEvents() ([]beatsaber.Event, error) {
	events := []beatsaber.Event{}

	for _, e := range p.events {
		roundedTime := beatsaber.Time(imath.Round(float64(p.Map.UnscaleTime(e.Beat())), 5))
		event := beatsaber.Event{
			Time:  roundedTime,
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

type EventCustomData struct {
	IlysaDev bool                   `json:"ilysaDev"`
	Rest     map[string]interface{} `json:"-"`
}

func (p *Project) getNonIlysaEvents() ([]beatsaber.Event, error) {
	mapEvents := p.Map.Events()

	fmt.Printf("map has %d events\n", len(mapEvents))

	niEvents := []beatsaber.Event{}

	for _, e := range mapEvents {
		if len(e.CustomData) == 0 {
			niEvents = append(niEvents, e)
			continue
		}

		cd := EventCustomData{}
		if err := json.Unmarshal(e.CustomData, &cd); err != nil {
			return nil, err
		}
		if !cd.IlysaDev {
			niEvents = append(niEvents, e)
		}
	}

	fmt.Printf("map has %d non-Ilysa events\n", len(niEvents))

	return niEvents, nil
}

func (p *Project) SaveDev() error {
	newEvents, err := p.getNonIlysaEvents()
	if err != nil {
		return err
	}

	generatedEvents, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(generatedEvents))

	for i, e := range generatedEvents {
		cd := EventCustomData{}
		if err := swallowjson.UnmarshalWith(&cd, "Rest", e.CustomData); err != nil {
			return err
		}
		cd.IlysaDev = true
		newRawMessage, err := swallowjson.MarshalWith(cd, "Rest")
		if err != nil {
			return err
		}

		generatedEvents[i].CustomData = newRawMessage
	}

	newEvents = append(newEvents, generatedEvents...)

	fmt.Printf("saving %d events\n", len(newEvents))

	return p.Map.SaveEvents(newEvents)
}

func (p *Project) SaveProd() error {
	niEvents, err := p.getNonIlysaEvents()
	if err != nil {
		return err
	}

	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(events))

	newEvents := append(niEvents, events...)

	fmt.Printf("saving %d events\n", len(newEvents))

	return p.Map.SaveEvents(newEvents)
}

func (p *Project) Append() error {
	events, err := p.generateBeatSaberEvents()
	if err != nil {
		return err
	}

	fmt.Printf("generated %d events\n", len(events))

	return p.Map.AppendEvents(events)
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
