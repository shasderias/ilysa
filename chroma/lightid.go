package chroma

import (
	"bytes"
	"encoding/json"
)

type LightID []int

func (l *LightID) UnmarshalJSON(raw []byte) error {
	if bytes.ContainsRune(raw, '[') {
		var lightIDAry []int
		if err := json.Unmarshal(raw, &lightIDAry); err != nil {
			return err
		}
		*l = lightIDAry
		return nil
	} else {
		var lightIDNumber int
		if err := json.Unmarshal(raw, &lightIDNumber); err != nil {
			return err
		}

		*l = LightID{lightIDNumber}
		return nil
	}
}

func (l LightID) MarshalJSON() ([]byte, error) {
	if len(l) == 1 {
		return json.Marshal(l[0])
	}
	return json.Marshal([]int(l))
}

func (l LightID) Has(t ...int) bool {
	for _, id := range l {
		for _, targetID := range t {
			if id == targetID {
				return true
			}
		}
	}
	return false
}
