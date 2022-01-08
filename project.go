package ilysa

import (
	"github.com/shasderias/ilysa/beatsaber"
	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
)

type Project struct {
	Map
	context.Context

	events evt.Events
}

type Map interface {
	Save() error
	SetEvents(events interface{})
	UnscaleTime(beat float64) beatsaber.Time
	DifficultyVersion() beatsaber.DifficultyVersion
}

func New(bsMap Map) *Project {
	p := &Project{
		Map:    bsMap,
		events: evt.NewEvents(),
	}

	p.Context = context.Base(p)

	return p
}

func (p *Project) AddEvents(events ...evt.Event) { p.events = append(p.events, events...) }
func (p *Project) Events() *evt.Events           { return &p.events }

func (p *Project) Save() error {
	p.Map.SetEvents(p.events)
	return p.Map.Save()
}
