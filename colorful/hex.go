package colorful

import "fmt"

// Hex returns the hex "html" representation of the color, as in #ff0080.
func (col Color) Hex() string {
	// Add 0.5 for rounding
	return fmt.Sprintf("#%02x%02x%02x", uint8(col.R*255.0+0.5), uint8(col.G*255.0+0.5), uint8(col.B*255.0+0.5))
}

// ParseHex parses and returns the Color for a hex triplet or quadruplet from s. The
// string must begin with a leading '#'.
func ParseHex(scol string) (Color, error) {
	var r, g, b uint8
	var a uint8 = 255

	var (
		format   string
		factor   float64
		scanArgs []interface{}
	)

	switch len(scol) {
	case 9:
		format = "#%02x%02x%02x%02x"
		factor = 1.0 / 255.0
		scanArgs = []interface{}{&r, &g, &b, &a}
	case 7:
		format = "#%02x%02x%02x"
		factor = 1.0 / 255.0
		scanArgs = []interface{}{&r, &g, &b}
	case 4:
		format = "#%1x%1x%1x"
		a = 0xf
		factor = 1.0 / 15.0
		scanArgs = []interface{}{&r, &g, &b}
	default:
		return Color{}, fmt.Errorf("color: %v is not a hex-color", scol)
	}

	_, err := fmt.Sscanf(scol, format, scanArgs...)
	if err != nil {
		return Color{}, err
	}

	return Color{float64(r) * factor, float64(g) * factor, float64(b) * factor, float64(a) * factor}, nil
}

// Hex parses and returns the Color for a hex triplet or quadruplet from s. The
// string must begin with a leading '#'. Hex panics if s cannot be parsed.
func Hex(s string) Color {
	c, err := ParseHex(s)
	if err != nil {
		panic("MustParseHex: " + err.Error())
	}
	return c
}

// MustParseHex parses and returns the Color for a hex triplet or quadruplet
// from s. The string must begin with a leading '#'. MustParseHex panics if
// s cannot be parsed.
//
// Deprecated: use Hex() instead.
func MustParseHex(s string) Color {
	return Hex(s)
}
