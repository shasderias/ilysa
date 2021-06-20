package chroma

import (
	"bytes"
	"encoding/json"
)

type LightID []int

func (l *LightID) UnmarshalJSON(raw []byte) error {
	if bytes.ContainsRune(raw, '[') {
		return json.Unmarshal(raw, &l)
	}
	var singleLightID int
	if err := json.Unmarshal(raw, &singleLightID); err != nil {
		return err
	}

	*l = LightID{singleLightID}
	return nil
}

func (l LightID) MarshalJSON() ([]byte, error) {
	if len(l) == 1 {
		return json.Marshal(l[0])
	}
	return json.Marshal([]int(l))
}
