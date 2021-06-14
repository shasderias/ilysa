package ilysa

import "ilysa/pkg/beatsaber"

func DefaultType(typ beatsaber.EventType) func(e Event) {
	return func(e Event) {
		e.Base().Type = typ
	}
}

func DefaultValue(val beatsaber.EventValue) func(e Event) {
	return func(e Event) {
		e.Base().Value = val
	}
}
