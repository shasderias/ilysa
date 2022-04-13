package chroma

import (
	"encoding/json"

	. "github.com/shasderias/ilysa/internal/null"
)

type RingRotation struct {
	NameFilter  Null[string]        `json:"_nameFilter"`
	Reset       Null[bool]          `json:"_reset"`
	Rotation    Null[float64]       `json:"_rotation"`
	Step        Null[float64]       `json:"_step"`
	Prop        Null[float64]       `json:"_prop"`
	Speed       Null[float64]       `json:"_speed"`
	Direction   Null[SpinDirection] `json:"_direction"`
	CounterSpin Null[bool]          `json:"_counterSpin"`
}

func (r RingRotation) CustomData() (json.RawMessage, error) {
	return marshalToCustomData(r)
}
