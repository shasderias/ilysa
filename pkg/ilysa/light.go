package ilysa

import (
	"ilysa/pkg/beatsaber"
)

type Light interface {
	EventType() beatsaber.EventType
	MinLightID() int
	MaxLightID() int
}

// BasicLight represents a light with the base game + OOB Chroma attributes
type BasicLight struct {
	project   *Project
	eventType beatsaber.EventType
}

func (p *Project) NewBasicLight(typ beatsaber.EventType) Light {
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
