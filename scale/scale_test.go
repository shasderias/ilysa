package scale

import (
	"testing"
)

func TestClamp(t *testing.T) {
	tests := []struct {
		name string
		t    float64
		min  float64
		max  float64
		want float64
	}{
		{name: "Min", t: -0.5, min: 0, max: 1, want: 0},
		{name: "Max", t: 1.5, min: 0, max: 1, want: 1},
		{name: "NotRequired", t: 0.5, min: 0, max: 1, want: 0.5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clamp(tt.t, tt.min, tt.max); got != tt.want {
				t.Errorf("clamp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUnclamp(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		r    float64
		rMin float64
		rMax float64
		tMin float64
		tMax float64
		want float64
	}{
		{name: "Sanity", r: 0.5, rMin: 0, rMax: 1, tMin: 0, tMax: 2, want: 1},
		{name: "Max", r: 1.5, rMin: 0, rMax: 1, tMin: 0, tMax: 2, want: 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unclamp(tt.rMin, tt.rMax, tt.tMin, tt.tMax)(tt.r); got != tt.want {
				t.Errorf("Unclamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
