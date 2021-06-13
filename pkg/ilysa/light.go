package ilysa

import (
	"ilysa/pkg/beatsaber"
)

type Light struct {
	ctx *Context
	beatsaber.EventType
}

func NewLight(typ beatsaber.EventType) *Light {
	return &Light{
		EventType: typ,
	}
}
