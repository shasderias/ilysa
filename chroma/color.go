package chroma

import (
	"encoding/json"
	"image/color"

	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/imath"
)

type Color []float64

func ColorFromColor(c color.Color) Color {
	if cf, ok := c.(colorful.Color); ok {
		return Color{cf.R, cf.G, cf.B, cf.A}
	}
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
	r := imath.Round(c[0], 3)
	g := imath.Round(c[1], 3)
	b := imath.Round(c[2], 3)
	a := imath.Round(c[3], 3)
	return json.Marshal([4]float64{r, g, b, a})
}
