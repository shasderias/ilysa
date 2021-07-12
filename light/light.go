package light

import "github.com/shasderias/ilysa/evt"

type MaxLightIDer interface {
	MaxLightID(t evt.LightType) int
}
