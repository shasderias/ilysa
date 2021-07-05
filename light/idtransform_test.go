package light

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDivideSingle(t *testing.T) {
	tests := []struct {
		name    string
		lightID ID
		want    IDSet
	}{
		{
			name:    "Simple",
			lightID: NewID(1, 2, 3, 4),
			want: NewIDSet(
				NewID(1),
				NewID(2),
				NewID(3),
				NewID(4),
			),
		},
		{
			name:    "Single",
			lightID: NewID(1),
			want: NewIDSet(
				NewID(1),
			),
		},
		{
			name:    "Empty",
			lightID: NewID(),
			want:    NewIDSet(),
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
		lightID ID
		divisor int
		want    IDSet
	}{
		{
			name:    "Simple - 1",
			lightID: NewID(1, 2, 3, 4),
			divisor: 3,
			want: NewIDSet(
				NewID(1),
				NewID(2),
				NewID(3, 4),
			),
		},
		{
			name:    "Simple - 2",
			lightID: NewID(1, 2, 3, 4, 5, 6),
			divisor: 3,
			want: NewIDSet(
				NewID(1, 2),
				NewID(3, 4),
				NewID(5, 6),
			),
		},
		{
			name:    "Simple - 3",
			lightID: NewID(1, 2, 3, 4, 5, 6, 7),
			divisor: 3,
			want: NewIDSet(
				NewID(1, 2),
				NewID(3, 4),
				NewID(5, 6, 7),
			),
		},
		{
			name:    "Divisor Greater",
			lightID: NewID(1, 2, 3),
			divisor: 4,
			want: NewIDSet(
				NewID(),
				NewID(),
				NewID(),
				NewID(1, 2, 3),
			),
		},
		{
			name:    "Empty",
			lightID: NewID(),
			divisor: 4,
			want: NewIDSet(
				NewID(),
				NewID(),
				NewID(),
				NewID(),
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
		lightID    ID
		groupCount int
		want       IDSet
	}{
		{
			name:       "Simple - 1",
			lightID:    NewID(1, 2, 3, 4),
			groupCount: 2,
			want: NewIDSet(
				NewID(1, 3),
				NewID(2, 4),
			),
		},
		{
			name:       "Simple - 2",
			lightID:    NewID(1, 2, 3, 4, 5),
			groupCount: 2,
			want: NewIDSet(
				NewID(1, 3, 5),
				NewID(2, 4),
			),
		},
		{
			name:       "groupSize greater",
			lightID:    NewID(1, 2, 3),
			groupCount: 4,
			want: NewIDSet(
				NewID(1),
				NewID(2),
				NewID(3),
				NewID(),
			),
		},
		{
			name:       "Empty",
			lightID:    NewID(),
			groupCount: 4,
			want: NewIDSet(
				NewID(),
				NewID(),
				NewID(),
				NewID(),
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
		lightID   ID
		groupSize int
		want      IDSet
	}{
		{
			name:      "Simple - 1",
			lightID:   NewID(1, 2, 3, 4, 5, 6),
			groupSize: 2,
			want: NewIDSet(
				NewID(1, 2),
				NewID(3, 4),
				NewID(5, 6),
			),
		},
		{
			name:      "Simple - 2",
			lightID:   NewID(1, 2, 3, 4, 5),
			groupSize: 3,
			want: NewIDSet(
				NewID(1, 2, 3),
				NewID(4, 5),
			),
		},
		{
			name:      "groupSize greater",
			lightID:   NewID(1, 2, 3),
			groupSize: 4,
			want: NewIDSet(
				NewID(1, 2, 3),
			),
		},
		{
			name:      "Empty",
			lightID:   NewID(),
			groupSize: 4,
			want: NewIDSet(
				NewID(),
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
