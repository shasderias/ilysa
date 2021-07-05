package light

type ID []int

func NewID(ids ...int) ID {
	return append(ID{}, ids...)
}

func NewIDFromInterval(from, to int) ID {
	id := make(ID, 0, to-from+1)
	for i := from; i <= to; i++ {
		id = append(id, i)
	}
	return id
}
