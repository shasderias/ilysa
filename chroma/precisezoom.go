package chroma

import "encoding/json"

type Zoom struct {
	Step  float64 `json:"_step"`
	Speed float64 `json:"_speed"`
}

func (e *Zoom) CustomData() (json.RawMessage, error) { return json.Marshal(e) }
