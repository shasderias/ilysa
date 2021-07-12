package chroma

import (
	"encoding/json"
	"image/color"
)

type RGB struct {
	LightID LightID     `json:"_lightID,omitempty"`
	Color   color.Color `json:"_color,omitempty"`
}

func (r RGB) CustomData() (json.RawMessage, error) {
	if r.LightID == nil && r.Color == nil {
		return nil, nil
	}

	cd := map[string]interface{}{}

	if r.LightID != nil {
		cd["_lightID"] = LightID(r.LightID)
	}
	if r.Color != nil {
		cd["_color"] = ColorFromColor(r.Color)
	}

	return json.Marshal(cd)
}
