package light

import "ilysa/pkg/util"

type ID []int

func NewLightID(ids ...int) ID {
	return append(ID{}, ids...)
}

func NewLightIDFromInterval(startID, endID int) ID {
	lightID := make(ID, 0, endID-startID+1)
	for i := startID; i <= endID; i++ {
		lightID = append(lightID, i)
	}
	return lightID
}

type IDSet []ID

func NewIDSet(ids ...ID) IDSet {
	return append(IDSet{}, ids...)
}

func (s *IDSet) Add(id ...ID) {
	*s = append(*s, id...)
}

func (s *IDSet) AppendToIndex(idx int, id ...int) {
	(*s)[idx] = append((*s)[idx], id...)
}

func (s IDSet) Index(i int) ID {
	n := len(s)
	return s[util.WraparoundIdx(n, i)]
}

func (s IDSet) Len() int {
	return len(s)
}
