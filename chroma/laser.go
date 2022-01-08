package chroma

import (
	"encoding/json"

	. "github.com/shasderias/ilysa/internal/null"
)

type LaserSpeed struct {
	LockPosition Null[bool]          `json:"_lockPosition"`
	Speed        Null[float64]       `json:"_speed"`
	Direction    Null[SpinDirection] `json:"_direction"`
}

func (ls LaserSpeed) CustomData() (json.RawMessage, error) { return marshalToCustomData(ls) }
