package colorful

// Boost returns a new color with col's RGB values multiplied by m.
func (col Color) Boost(m float64) Color {
	return Color{col.R * m, col.G * m, col.B * m, col.A}
}
