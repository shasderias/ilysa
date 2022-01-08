package null

type Null[T any] struct {
	Value T
	Valid bool
}

func (n Null[T]) Get() T {
	return n.Value
}

func (n Null[T]) GetInterface() any {
	return n.Value
}

func (n *Null[T]) Set(v T) {
	n.Value = v
	n.Valid = true
}

func (n Null[T]) IsValid() bool {
	return n.Valid
}

func New[T any](v T) Null[T] {
	return Null[T]{v, true}
}

func NewInvalid[T any]() Null[T] {
	return Null[T]{*new(T), false}
}
