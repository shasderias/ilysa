package chroma

import (
	"encoding/json"

	. "github.com/shasderias/ilysa/internal/null"
)

type Zoom struct {
	Step  Null[float64] `json:"_step"`
	Speed Null[float64] `json:"_speed"`
}

func (z Zoom) CustomData() (json.RawMessage, error) {
	return marshalToCustomData(z)
}
