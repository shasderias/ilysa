package chroma

import (
	"encoding/json"

	"github.com/shasderias/ilysa/internal/imath"
)

type LaserSpeed struct {
	LockPosition bool          `json:"_lockPosition"`
	Speed        float64       `json:"_speed"`
	Direction    SpinDirection `json:"_direction"`
}

func (pl *LaserSpeed) CustomData() (json.RawMessage, error) {
	pl.Speed = imath.Round(pl.Speed, 2)
	return json.Marshal(pl)
}
