package timer

import "testing"

func TestSeq_Idx(t *testing.T) {
	tests := []struct {
		name     string
		seq      []float64
		subTests []struct {
			idx  int
			want float64
		}
	}{
		{"Sanity", []float64{0, 1}, []struct {
			idx  int
			want float64
		}{{0, 0}, {1, 0}, {-1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := SequencerFromSlice(tt.seq)
			for _, st := range tt.subTests {
				if got := seq.Idx(st.idx); got != st.want {
					t.Errorf("Idx(%d) = %v, want %v", st.idx, got, st.want)
				}
			}
		})
	}
}

func TestSeq_NextBOffset(t *testing.T) {
	tests := []struct {
		name     string
		seq      []float64
		subTests []struct {
			idx  int
			want float64
		}
	}{
		{"Sanity", []float64{0, 1}, []struct {
			idx  int
			want float64
		}{{0, 1}, {1, 1}, {-1, 1}}},
		{"TwoBeats", []float64{0, 1, 2}, []struct {
			idx  int
			want float64
		}{{0, 1}, {1, 1}, {-1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			seq := SequencerFromSlice(tt.seq)
			for _, st := range tt.subTests {
				if got := seq.NextBOffset(st.idx); got != st.want {
					t.Errorf("Idx(%d) = %v, want %v", st.idx, got, st.want)
				}
			}
		})
	}
}
