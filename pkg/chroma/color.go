package chroma

import (
	"encoding/json"
	"image/color"
)

type Color []float64

func ColorFromColor(c color.Color) Color {
	r, g, b, a := c.RGBA()
	return Color{float64(r) / 0xffff, float64(g) / 0xffff, float64(b) / 0xffff, float64(a) / 0xffff}
}

func (c *Color) UnmarshalJSON(raw []byte) error {
	bsColor := new([4]float64)
	err := json.Unmarshal(raw, &bsColor)
	if err != nil {
		return err
	}

	*c = Color{bsColor[0], bsColor[1], bsColor[2], bsColor[3]}

	return nil
}

func (c Color) MarshalJSON() ([]byte, error) {
	return json.Marshal([4]float64{c[0], c[1], c[2], c[3]})
}
