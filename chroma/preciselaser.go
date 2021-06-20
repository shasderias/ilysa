package chroma

import "encoding/json"

type PreciseLaser struct {
	LockPosition bool          `json:"_lockPosition"`
	Speed     float64       `json:"_speed"`
	Direction SpinDirection `json:"_direction"`
}

func (pl *PreciseLaser) CustomData() (json.RawMessage, error) {
	return json.Marshal(pl)
}
