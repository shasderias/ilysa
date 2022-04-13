package transform

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/shasderias/ilysa/lightid"
)

func TestDivideSingle(t *testing.T) {
	tests := []struct {
		name    string
		lightID lightid.ID
		want    lightid.Set
	}{
		{
			name:    "Simple",
			lightID: lightid.New(1, 2, 3, 4),
			want: lightid.NewSet(
				lightid.New(1),
				lightid.New(2),
				lightid.New(3),
				lightid.New(4),
			),
		},
		{
			name:    "Single",
			lightID: lightid.New(1),
			want: lightid.NewSet(
				lightid.New(1),
			),
		},
		{
			name:    "Empty",
			lightID: lightid.New(),
			want:    lightid.NewSet(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(DivideSingle().do(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		lightID lightid.ID
		divisor int
		want    lightid.Set
	}{
		{
			name:    "Simple - 1",
			lightID: lightid.New(1, 2, 3, 4),
			divisor: 3,
			want: lightid.NewSet(
				lightid.New(1),
				lightid.New(2),
				lightid.New(3, 4),
			),
		},
		{
			name:    "Simple - 2",
			lightID: lightid.New(1, 2, 3, 4, 5, 6),
			divisor: 3,
			want: lightid.NewSet(
				lightid.New(1, 2),
				lightid.New(3, 4),
				lightid.New(5, 6),
			),
		},
		{
			name:    "Simple - 3",
			lightID: lightid.New(1, 2, 3, 4, 5, 6, 7),
			divisor: 3,
			want: lightid.NewSet(
				lightid.New(1, 2),
				lightid.New(3, 4),
				lightid.New(5, 6, 7),
			),
		},
		{
			name:    "Divisor Greater",
			lightID: lightid.New(1, 2, 3),
			divisor: 4,
			want: lightid.NewSet(
				lightid.New(),
				lightid.New(),
				lightid.New(),
				lightid.New(1, 2, 3),
			),
		},
		{
			name:    "Empty",
			lightID: lightid.New(),
			divisor: 4,
			want: lightid.NewSet(
				lightid.New(),
				lightid.New(),
				lightid.New(),
				lightid.New(),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(Divide(tt.divisor).do(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}

func TestFan(t *testing.T) {
	tests := []struct {
		name       string
		lightID    lightid.ID
		groupCount int
		want       lightid.Set
	}{
		{
			name:       "Simple - 1",
			lightID:    lightid.New(1, 2, 3, 4),
			groupCount: 2,
			want: lightid.NewSet(
				lightid.New(1, 3),
				lightid.New(2, 4),
			),
		},
		{
			name:       "Simple - 2",
			lightID:    lightid.New(1, 2, 3, 4, 5),
			groupCount: 2,
			want: lightid.NewSet(
				lightid.New(1, 3, 5),
				lightid.New(2, 4),
			),
		},
		{
			name:       "groupSize greater",
			lightID:    lightid.New(1, 2, 3),
			groupCount: 4,
			want: lightid.NewSet(
				lightid.New(1),
				lightid.New(2),
				lightid.New(3),
				lightid.New(),
			),
		},
		{
			name:       "Empty",
			lightID:    lightid.New(),
			groupCount: 4,
			want: lightid.NewSet(
				lightid.New(),
				lightid.New(),
				lightid.New(),
				lightid.New(),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(Fan(tt.groupCount).do(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}

func TestDivideIntoGroups(t *testing.T) {
	tests := []struct {
		name      string
		lightID   lightid.ID
		groupSize int
		want      lightid.Set
	}{
		{
			name:      "Simple - 1",
			lightID:   lightid.New(1, 2, 3, 4, 5, 6),
			groupSize: 2,
			want: lightid.NewSet(
				lightid.New(1, 2),
				lightid.New(3, 4),
				lightid.New(5, 6),
			),
		},
		{
			name:      "Simple - 2",
			lightID:   lightid.New(1, 2, 3, 4, 5),
			groupSize: 3,
			want: lightid.NewSet(
				lightid.New(1, 2, 3),
				lightid.New(4, 5),
			),
		},
		{
			name:      "groupSize greater",
			lightID:   lightid.New(1, 2, 3),
			groupSize: 4,
			want: lightid.NewSet(
				lightid.New(1, 2, 3),
			),
		},
		{
			name:      "Empty",
			lightID:   lightid.New(),
			groupSize: 4,
			want: lightid.NewSet(
				lightid.New(),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(DivideIntoGroups(tt.groupSize).do(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}
