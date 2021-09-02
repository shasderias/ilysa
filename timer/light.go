package timer

type Light interface {
	LightIDT() float64 // current time in light ID sequence, 0-1
	LightIDOrdinal() int
	LightIDLen() int
	LightIDCur() int

	Next() bool
}

type Lighter struct {
	l LightIDLener
}

type LightIDLener interface {
	LightIDLen() int
}

func NewLighter(l LightIDLener) Lighter {
	return Lighter{l}
}

func (l Lighter) Iterate() Light {
	return &LightIterator{l, -1}
}

func (l Lighter) IterateFrom(ordinal int) Light {
	return &LightIterator{l, ordinal}
}

type LightIterator struct {
	Lighter
	ordinal int
}

func (i *LightIterator) Next() bool {
	i.ordinal++
	if i.ordinal == i.LightIDLen() {
		return false
	}
	return true
}

func (i *LightIterator) LightIDT() float64 {
	if i.LightIDLen() == 1 {
		return 1
	}
	return float64(i.ordinal) / float64(i.LightIDLen()-1)
}
func (i *LightIterator) LightIDOrdinal() int { return i.ordinal }
func (i *LightIterator) LightIDLen() int     { return i.Lighter.l.LightIDLen() }
func (i *LightIterator) LightIDCur() int     { return i.ordinal + 1 }
