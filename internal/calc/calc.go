package calc

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func WraparoundIdx(len, idx int) int {
	return ((idx % len) + len) % len
}

func IndexWraparound[T any](slice []T, index int) T {
	return slice[WraparoundIdx(len(slice), index)]
}
