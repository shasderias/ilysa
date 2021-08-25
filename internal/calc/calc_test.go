package calc

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		name string
		n    int64
		want int64
	}{
		{name: "Zero", n: 0, want: 0},
		{name: "Positive", n: 1, want: 1},
		{name: "Negative", n: -1, want: 1},
		{name: "Big Positive", n: 10000, want: 10000},
		{name: "Big Negative", n: -10000, want: 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.n); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWraparoundIdx(t *testing.T) {
	tests := []struct {
		name     string
		len      int
		subTests []struct {
			idx  int
			want int
		}
	}{
		{name: "One", len: 1, subTests: []struct {
			idx  int
			want int
		}{{0, 0}, {1, 0}, {-1, 0}, {50, 0}}},
		{name: "Four", len: 4, subTests: []struct {
			idx  int
			want int
		}{{0, 0}, {1, 1}, {4, 0}, {-1, 3}, {8, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, st := range tt.subTests {
				if got := WraparoundIdx(tt.len, st.idx); got != st.want {
					t.Errorf("WraparoundIdx(%d, %d) = %d, want %d", tt.len, st.idx, got, st.want)
				}
			}
		})
	}
}
