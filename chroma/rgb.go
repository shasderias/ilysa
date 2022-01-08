package chroma

import (
	"encoding/json"
	"image/color"

	"github.com/shasderias/ilysa/ease"
)

type LerpType string

const (
	LerpTypeHSV LerpType = "HSV"
	LerpTypeRGB          = "Lighting"
)

type Lighting struct {
	LightID  LightID
	Color    color.Color
	Easing   ease.Func
	LerpType LerpType
}

func (r Lighting) CustomData() (json.RawMessage, error) {
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
	if r.Easing != nil {
		cd["_easing"] = r.Easing.EaseName()
	}
	if r.LerpType != "" {
		cd["_lerpType"] = string(r.LerpType)
	}

	return json.Marshal(cd)
}

func (r Lighting) Copy() Lighting {
	return Lighting{
		LightID:  r.LightID,
		Color:    r.Color,
		Easing:   r.Easing,
		LerpType: r.LerpType,
	}
}
