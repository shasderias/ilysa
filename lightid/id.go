package lightid

// ID represents a set of light IDs that will be collectively acted upon.
// Contrast with a Set which represents a set of IDs that will be acted upon
// in sequence.
type ID []int

func New(ids ...int) ID {
	return append(ID{}, ids...)
}

func (lightIDs ID) Clone() ID {
	return append(ID{}, lightIDs...)
}

func NewFromInterval(from, to int) ID {
	id := make(ID, 0, to-from+1)
	for i := from; i <= to; i++ {
		id = append(id, i)
	}
	return id
}

func (lightIDs *ID) Add(id ...int) {
	*lightIDs = append(*lightIDs, id...)
}

func (lightIDs ID) Has(wantIDs ...int) bool {
	for _, wantID := range wantIDs {
		for _, lid := range lightIDs {
			if wantID == lid {
				return true
			}
		}
	}
	return false
}
