package gradient_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shasderias/ilysa/colorful"
	"github.com/shasderias/ilysa/colorful/gradient"
)

var (
	R = colorful.Red
	G = colorful.Green
	B = colorful.Blue
)

func TestNewPingPong(t *testing.T) {

	tests := []struct {
		name   string
		count  int
		colors []colorful.Color
		want   gradient.Table
	}{
		{
			name:   "Simple1",
			count:  1,
			colors: []colorful.Color{R, G, B},
			want:   gradient.New(R, G, B, G, R),
		},
		{
			name:   "Simple2",
			count:  2,
			colors: []colorful.Color{R, G, B},
			want:   gradient.New(R, G, B, G, R, G, B),
		},
		{
			name:   "Simple3",
			count:  3,
			colors: []colorful.Color{R, G, B},
			want:   gradient.New(R, G, B, G, R, G, B, G, R),
		},
		{
			name:   "Short",
			count:  1,
			colors: []colorful.Color{R, B},
			want:   gradient.New(R, B, R),
		},
		{
			name:   "Short2",
			count:  2,
			colors: []colorful.Color{R, B},
			want:   gradient.New(R, B, R, B),
		},
		{
			name:   "Long",
			count:  1,
			colors: []colorful.Color{R, B, B, R},
			want:   gradient.New(R, B, B, R, B, B, R),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := gradient.NewPingPong(tt.count, tt.colors...)

			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		Grad gradient.Table
		N    int
		Want gradient.Table
	}{
		{
			gradient.New(R, G, B),
			1,
			gradient.New(G, B, R),
		},
		{
			gradient.New(R, G, B),
			2,
			gradient.New(B, R, G),
		},
		{
			gradient.New(R, G, B),
			3,
			gradient.New(R, G, B),
		},
		{
			gradient.New(R, G, B),
			4,
			gradient.New(G, B, R),
		},
	}

	for _, tt := range testCases {
		if diff := cmp.Diff(tt.Grad.Rotate(tt.N), tt.Want); diff != "" {
			t.Fatal(diff)
		}
	}
}

func TestBoost(t *testing.T) {
	var (
		red   = colorful.Color{1, 0, 0, 0.5}
		green = colorful.Color{0, 1, 0, 0.2}
		blue  = colorful.Color{0, 0, 1, 1.0}
		table = gradient.Table{
			{red, 0.0},
			{green, 0.3},
			{blue, 1.0},
		}
		want = gradient.Table{
			{colorful.Color{1.5, 0, 0, 0.5}, 0.0},
			{colorful.Color{0, 1.5, 0, 0.2}, 0.3},
			{colorful.Color{0, 0, 1.5, 1.0}, 1.0},
		}
	)

	boostedTable := table.Boost(1.5)

	if diff := cmp.Diff(boostedTable, want); diff != "" {
		t.Fatal(diff)
	}
}
