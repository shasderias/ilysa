package ilysa

import (
	"ilysa/pkg/beatsaber"
	"ilysa/pkg/chroma"
	"ilysa/pkg/chroma/lightid"
)

type Light interface {
	EventType() beatsaber.EventType
	MinLightID() int
	MaxLightID() int
	LightID() chroma.LightID
}

// BasicLight represents a light with the base game + OOB Chroma attributes
type BasicLight struct {
	project   *Project
	eventType beatsaber.EventType
}

func (p *Project) NewBasicLight(typ beatsaber.EventType) BasicLight {
	return BasicLight{
		project:   p,
		eventType: typ,
	}
}

func (l BasicLight) EventType() beatsaber.EventType {
	return l.eventType
}

func (l BasicLight) MinLightID() int {
	return 1
}

func (l BasicLight) MaxLightID() int {
	return l.project.ActiveDifficultyProfile().MaxLightID(l.eventType)
}

type LightSlice struct {
	eventType beatsaber.EventType
	slice     lightid.Slice
}

func (p *Project) NewSlicedLight(typ beatsaber.EventType, s lightid.Slice) LightSlice {
	return LightSlice{
		eventType: typ,
		slice:     s,
	}
}

func (l LightSlice) EventType() beatsaber.EventType {
	return l.eventType
}

func (l LightSlice) MinLightID() int {
	return l.slice.Start
}

func (l LightSlice) MaxLightID() int {
	return l.slice.End
}
