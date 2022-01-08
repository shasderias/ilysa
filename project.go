package ilysa

import (
	"fmt"

	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/internal/imath"
)

type Project struct {
	Map
	context.Context

	events evt.Events
}

type Map interface {
	Save() error
	GetEvents() *[]beatsaber.Event
	UnscaleTime(beat float64) beatsaber.Time
}

func New(bsMap Map) *Project {
	p := &Project{
		Map:    bsMap,
		events: evt.NewEvents(),
	}

	p.Context = context.Base(p)

	return p
}

func (p *Project) AddEvents(events ...evt.Event) {
	p.events = append(p.events, events...)
}

func (p *Project) Events() *evt.Events {
	return &p.events
}

//func (p *Project) FilterEvents(f func(e evt.Event) bool) *[]evt.Event {
//	filteredEvents := make([]evt.Event, 0)
//
//	for _, e := range p.events {
//		if f(e) {
//			filteredEvents = append(filteredEvents, e)
//		}
//	}
//
//	p.events = filteredEvents
//
//	return &filteredEvents
//}
//
//func (p *Project) FilterEventsFast(f func(e evt.Event) bool) *[]evt.Event {
//	n := 0
//	for _, e := range p.events {
//		if f(e) {
//			p.events[n] = e
//			n++
//		}
//	}
//
//	p.events = p.events[:n]
//
//	return &p.events
//}
//
//func (p *Project) MapEvents(f func(e evt.Event) evt.Event) {
//	for i := range p.events {
//		p.events[i] = f(p.events[i])
//	}
//}
//
//func (p *Project) sortEvents() {
//	sort.Slice(p.events, func(i, j int) bool {
//		return p.events[i].Beat() < p.events[j].Beat()
//	})
//}

func (p *Project) generateBeatSaberEvents() ([]beatsaber.Event, error) {
	//cumulativeRotation := 0.0

	events := []beatsaber.Event{}

	//prevRot := 0.0
	for _, e := range p.events {
		roundedTime := beatsaber.Time(imath.Round(float64(p.Map.UnscaleTime(e.Beat())), 5))
		event := beatsaber.Event{
			Time:       roundedTime,
			Type:       int(e.Type()),
			Value:      int(e.Value()),
			FloatValue: e.FloatValue(),
		}

		//if rotEvent, ok := e.(*evt.ChromaRingRotation); ok {
		//	if rotEvent.Beat()-prevRot < 0.1 {
		//		fmt.Printf("warning: two rotation events less than 0.1 beats apart at beat %f\n", rotEvent.Beat())
		//	}
		//	prevRot = rotEvent.Beat()

		//if e.HasTag(evt.IlysaRotationResetTag) {
		//	rot := -(math.Remainder(cumulativeRotation, 360))
		//	if rotEvent.Direction == chroma.Clockwise {
		//		rotEvent.Rotation -= rot
		//	} else {
		//		rotEvent.Rotation += -rot
		//	}
		//}
		//if rotEvent.Direction == chroma.Clockwise {
		//	cumulativeRotation += rotEvent.Rotation
		//} else {
		//	cumulativeRotation -= rotEvent.Rotation
		//}
		//}

		if cder, ok := e.(evt.CustomDataer); ok {
			var err error
			event.CustomData, err = cder.CustomData()
			if err != nil {
				return nil, err
			}
		}

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

	*(p.GetEvents()) = events
	return p.Map.Save()
}

//type EventCustomData struct {
//	IlysaDev bool                   `json:"ilysaDev"`
//	Rest     map[string]interface{} `json:"-"`
//}
//
//func (p *Project) getNonIlysaEvents() ([]beatsaber.Event, error) {
//	mapEvents := p.Map.Events()
//
//	fmt.Printf("map has %d events\n", len(mapEvents))
//
//	niEvents := []beatsaber.Event{}
//
//	for _, e := range mapEvents {
//		if len(e.CustomData) == 0 {
//			niEvents = append(niEvents, e)
//			continue
//		}
//
//		cd := EventCustomData{}
//		if err := json.Unmarshal(e.CustomData, &cd); err != nil {
//			return nil, err
//		}
//		if !cd.IlysaDev {
//			niEvents = append(niEvents, e)
//		}
//	}
//
//	fmt.Printf("map has %d non-Ilysa events\n", len(niEvents))
//
//	return niEvents, nil
//}
//
//func (p *Project) SaveDev() error {
//	newEvents, err := p.getNonIlysaEvents()
//	if err != nil {
//		return err
//	}
//
//	generatedEvents, err := p.generateBeatSaberEvents()
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("generated %d events\n", len(generatedEvents))
//
//	for i, e := range generatedEvents {
//		cd := EventCustomData{}
//		if err := swallowjson.UnmarshalWith(&cd, "Rest", e.CustomData); err != nil {
//			return err
//		}
//		cd.IlysaDev = true
//		newRawMessage, err := swallowjson.MarshalWith(cd, "Rest")
//		if err != nil {
//			return err
//		}
//
//		generatedEvents[i].CustomData = newRawMessage
//	}
//
//	newEvents = append(newEvents, generatedEvents...)
//
//	fmt.Printf("saving %d events\n", len(newEvents))
//
//	return p.Map.SaveEvents(newEvents)
//}
//
//func (p *Project) SaveProd() error {
//	niEvents, err := p.getNonIlysaEvents()
//	if err != nil {
//		return err
//	}
//
//	events, err := p.generateBeatSaberEvents()
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("generated %d events\n", len(events))
//
//	newEvents := append(niEvents, events...)
//
//	fmt.Printf("saving %d events\n", len(newEvents))
//
//	return p.Map.SaveEvents(newEvents)
//}
//
//func (p *Project) Append() error {
//	events, err := p.generateBeatSaberEvents()
//	if err != nil {
//		return err
//	}
//
//	fmt.Printf("generated %d events\n", len(events))
//
//	return p.Map.AppendEvents(events)
//}
//
//func (p *Project) Dump() error {
//	events, err := p.generateBeatSaberEvents()
//	if err != nil {
//		return err
//	}
//
//	eventsJSON, err := json.Marshal(events)
//	if err != nil {
//		return err
//	}
//
//	fmt.Print(string(eventsJSON))
//
//	return nil
//}
