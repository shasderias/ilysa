package chroma

import "ilysa/pkg/beatsaber"

type LightProfile interface {
	MinLightID() int
	MaxLightID() int
}

func ProfileFor(eventType beatsaber.EventType) {

}
