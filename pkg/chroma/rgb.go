package chroma

import (
	"encoding/json"
	"image/color"

	"ilysa/pkg/colorful"
)

type RGB struct {
	LightID LightID     `json:"_lightID,omitempty"`
	Color   color.Color `json:"_color,omitempty"`
}

func (r *RGB) CustomData() (json.RawMessage, error) {
	if r.LightID == nil && r.Color == nil {
		return nil, nil
	}

	cd := map[string]interface{}{}

	if r.LightID != nil {
		cd["_lightID"] = r.LightID
	}
	if r.Color != nil {
		cd["_color"] = ColorFromColor(r.Color)
	}

	return json.Marshal(cd)
}

func (r *RGB) SetColor(c color.Color) {
	r.Color = c
}

func (r *RGB) GetAlpha() float64 {
	c := colorful.FromColor(r.Color)
	return c.A
}

func (r *RGB) SetAlpha(a float64) {
	c := colorful.FromColor(r.Color)
	c.A = a
	r.Color = c
}

func (r *RGB) FirstLightID() int {
	if len(r.LightID) == 0 {
		panic("RGB.FirstLightID: lightID is nil or contains no lightIDs")
	}
	return r.LightID[0]
}

func (r *RGB) SetSingleLightID(id int) {
	r.LightID = []int{id}
}

func (r *RGB) SetLightID(ids LightID) {
	r.LightID = ids
}
