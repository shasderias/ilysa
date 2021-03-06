package chroma

type SpinDirection int

const (
	CounterClockwise SpinDirection = 0
	Clockwise        SpinDirection = 1
)

func (d SpinDirection) Reverse() SpinDirection {
	if d == CounterClockwise {
		return Clockwise
	}
	return CounterClockwise
}

func (d SpinDirection) ReverseIf(b bool) SpinDirection {
	if !b {
		return d
	}
	if d == CounterClockwise {
		return Clockwise
	}
	return CounterClockwise
}
