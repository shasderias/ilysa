package ilysa

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDivideSingle(t *testing.T) {
	tests := []struct {
		name    string
		lightID LightID
		want    LightIDSet
	}{
		{
			name:    "Simple",
			lightID: NewLightID(1, 2, 3, 4),
			want: NewLightIDSet(
				NewLightID(1),
				NewLightID(2),
				NewLightID(3),
				NewLightID(4),
			),
		},
		{
			name:    "Single",
			lightID: NewLightID(1),
			want: NewLightIDSet(
				NewLightID(1),
			),
		},
		{
			name:    "Empty",
			lightID: NewLightID(),
			want:    NewLightIDSet(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(DivideSingle(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		lightID LightID
		divisor int
		want    LightIDSet
	}{
		{
			name:    "Simple - 1",
			lightID: NewLightID(1, 2, 3, 4),
			divisor: 3,
			want: NewLightIDSet(
				NewLightID(1),
				NewLightID(2),
				NewLightID(3, 4),
			),
		},
		{
			name:    "Simple - 2",
			lightID: NewLightID(1, 2, 3, 4, 5, 6),
			divisor: 3,
			want: NewLightIDSet(
				NewLightID(1, 2),
				NewLightID(3, 4),
				NewLightID(5, 6),
			),
		},
		{
			name:    "Simple - 3",
			lightID: NewLightID(1, 2, 3, 4, 5, 6, 7),
			divisor: 3,
			want: NewLightIDSet(
				NewLightID(1, 2),
				NewLightID(3, 4),
				NewLightID(5, 6, 7),
			),
		},
		{
			name:    "Divisor Greater",
			lightID: NewLightID(1, 2, 3),
			divisor: 4,
			want: NewLightIDSet(
				NewLightID(),
				NewLightID(),
				NewLightID(),
				NewLightID(1, 2, 3),
			),
		},
		{
			name:    "Empty",
			lightID: NewLightID(),
			divisor: 4,
			want: NewLightIDSet(
				NewLightID(),
				NewLightID(),
				NewLightID(),
				NewLightID(),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(Divide(tt.divisor)(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}

func TestFan(t *testing.T) {
	tests := []struct {
		name       string
		lightID    LightID
		groupCount int
		want       LightIDSet
	}{
		{
			name:       "Simple - 1",
			lightID:    NewLightID(1, 2, 3, 4),
			groupCount: 2,
			want: NewLightIDSet(
				NewLightID(1, 3),
				NewLightID(2, 4),
			),
		},
		{
			name:       "Simple - 2",
			lightID:    NewLightID(1, 2, 3, 4, 5),
			groupCount: 2,
			want: NewLightIDSet(
				NewLightID(1, 3, 5),
				NewLightID(2, 4),
			),
		},
		{
			name:       "groupSize greater",
			lightID:    NewLightID(1, 2, 3),
			groupCount: 4,
			want: NewLightIDSet(
				NewLightID(1),
				NewLightID(2),
				NewLightID(3),
				NewLightID(),
			),
		},
		{
			name:       "Empty",
			lightID:    NewLightID(),
			groupCount: 4,
			want: NewLightIDSet(
				NewLightID(),
				NewLightID(),
				NewLightID(),
				NewLightID(),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(Fan(tt.groupCount)(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}

func TestDivideIntoGroupsOf(t *testing.T) {
	tests := []struct {
		name      string
		lightID   LightID
		groupSize int
		want      LightIDSet
	}{
		{
			name:      "Simple - 1",
			lightID:   NewLightID(1, 2, 3, 4, 5, 6),
			groupSize: 2,
			want: NewLightIDSet(
				NewLightID(1, 2),
				NewLightID(3, 4),
				NewLightID(5, 6),
			),
		},
		{
			name:      "Simple - 2",
			lightID:   NewLightID(1, 2, 3, 4, 5),
			groupSize: 3,
			want: NewLightIDSet(
				NewLightID(1, 2, 3),
				NewLightID(4, 5),
			),
		},
		{
			name:      "groupSize greater",
			lightID:   NewLightID(1, 2, 3),
			groupSize: 4,
			want: NewLightIDSet(
				NewLightID(1, 2, 3),
			),
		},
		{
			name:      "Empty",
			lightID:   NewLightID(),
			groupSize: 4,
			want: NewLightIDSet(
				NewLightID(),
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if diff := cmp.Diff(DivideIntoGroupsOf(tt.groupSize)(tt.lightID), tt.want); diff != "" {
				t.Fatalf("%s: %s", tt.name, diff)
			}
		})
	}
}
