package lightid

import (
	"reflect"
	"testing"

	"ilysa/pkg/chroma"
)

func TestSetSanity(t *testing.T) {
	s := NewSet()
	if s.Len() != 0 {
		t.Fatalf("got %d; want %d", s.Len(), 0)
	}

	s.Add(chroma.LightID{1})

	if s.Len() != 1 {
		t.Fatalf("got %d; want %d", s.Len(), 1)
	}

	p := s.Pick(0)

	if p[0] != 1 && len(p) != 1 {
		t.Fatal()
	}
}

func TestSetPick(t *testing.T) {
	tests := []struct {
		name string
		s    Set
		n    []int
		want []chroma.LightID
	}{
		{
			name: "Basic",
			s:    Set{{0}, {1}, {2}},
			n:    []int{0, 1, 2, 3},
			want: []chroma.LightID{
				{0}, {1}, {2}, {0},
			},
		},
		{
			name: "Negative",
			s:    Set{{0}, {1}, {2}},
			n:    []int{0, -1, -2},
			want: []chroma.LightID{
				{0}, {2}, {1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < len(tt.n); i++ {
				if got := tt.s.Pick(tt.n[i]); !reflect.DeepEqual(got, tt.want[i]) {
					t.Errorf("Index() = %v, want %v", got, tt.want[i])
				}
			}
		})
	}
}
