package chroma

import "encoding/json"

type RingRotation struct {
	NameFilter  string        `json:"_nameFilter,omitempty"`
	Reset       bool          `json:"_reset,omitempty"`
	Rotation    float64       `json:"_rotation"`
	Step        float64       `json:"_step"`
	Prop        float64       `json:"_prop,omitempty"`
	Speed       float64       `json:"_speed,omitempty"`
	Direction   SpinDirection `json:"_direction"`
	CounterSpin bool          `json:"_counterSpin,omitempty"`
}

func (r RingRotation) CustomData() (json.RawMessage, error) {
	return json.Marshal(r)
}
