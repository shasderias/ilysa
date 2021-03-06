package lightid

// ID represents a set of light IDs that will be collectively acted upon.
// Contrast with a Set which represents a set of IDs that will be acted upon
// in sequence.
type ID []int

func New(ids ...int) ID {
	return append(ID{}, ids...)
}

func NewFromInterval(from, to int) ID {
	id := make(ID, 0, to-from+1)
	for i := from; i <= to; i++ {
		id = append(id, i)
	}
	return id
}
