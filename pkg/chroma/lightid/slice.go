package lightid

type Slice struct {
	Start int
	End   int
}

func NewSlice(start, end int) Slice {
	return Slice{start, end}
}
