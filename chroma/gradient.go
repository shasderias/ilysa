package chroma

import (
	"encoding/json"
	"image/color"

	"github.com/shasderias/ilysa/internal/imath"
)

type Gradient struct {
	LightGradient struct {
		Duration   float64     `json:"_duration"`
		StartColor color.Color `json:"_startColor"`
		EndColor   color.Color `json:"_endColor"`
		Easing     string      `json:"_easing"`
	} `json:"_lightGradient"`
}

func (g Gradient) CustomData() (json.RawMessage, error) {
	cd := map[string]interface{}{
		"_lightGradient": map[string]interface{}{
			"_duration":   imath.Round(g.LightGradient.Duration, 5),
			"_startColor": ColorFromColor(g.LightGradient.StartColor),
			"_endColor":   ColorFromColor(g.LightGradient.EndColor),
			"_easing":     g.LightGradient.Easing,
		},
	}

	return json.Marshal(cd)
}
