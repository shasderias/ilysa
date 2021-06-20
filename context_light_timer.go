package ilysa

type lightTimer struct {
	Light
	lightIDOrdinal int
}

func newLightTimer(l Light, lightIDOrdinal int) lightTimer {
	return lightTimer{
		Light:          l,
		lightIDOrdinal: lightIDOrdinal,
	}
}

func (c lightTimer) LightIDOrdinal() int {
	return c.lightIDOrdinal
}

func (c lightTimer) LightIDCur() int {
	return c.lightIDOrdinal + 1
}

func (c lightTimer) LightIDT() float64 {
	return float64(c.LightIDOrdinal()) / float64(c.LightIDLen())
}
