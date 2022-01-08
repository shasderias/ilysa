package timer

type Light interface {
	LightIDT() float64 // current time in light ID sequence, 0-1
	LightIDOrdinal() int
	LightLen() int
	LightIDCur() int

	Next() bool
}

type Lighter struct {
	l LightLener
}

type LightLener interface {
	LightLen() int
}
