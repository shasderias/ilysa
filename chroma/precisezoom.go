package chroma

import "encoding/json"

type PreciseZoom struct {
	Step  float64 `json:"_step"`
	Speed float64 `json:"_speed,omitempty"`
}

func (e *PreciseZoom) CustomData() (json.RawMessage, error) { return json.Marshal(e) }
