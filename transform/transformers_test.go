package transform

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/lightid"
)

func newTestLight(length int) testLight {
	return testLight{
		lightid.NewSet(lightid.NewFromInterval(1, length)),
	}
}
func newTestLightFromSet(set lightid.Set) testLight {
	return testLight{set}
}

type testLight struct {
	s lightid.Set
}

func (l testLight) GenerateEvents(ctx context.LightContext) evt.Events {
	//TODO implement me
	panic("implement me")
}

func (l testLight) LightLen() int {
	return l.s.Len()
}

func (l testLight) Name() []string {
	return []string{"testLight"}
}

func (l testLight) ApplyLightIDTransform(fn func(set lightid.Set) lightid.Set) context.Light {
	return testLight{fn(l.s)}
}

func (l testLight) ApplyLightIDSequenceTransform(fn func(set lightid.Set) lightid.Set) context.Light {
	ns := fn(l.s)
	seq := light.NewSequence()

	for _, id := range ns {
		seq.Add(testLight{lightid.NewSet(id)})
	}

	return seq
}

func transformAndAssertLight(t *testing.T, light context.Light, want lightid.Set, transforms ...LightTransformer) {
	tfedLight := Light(light, transforms...)
	if diff := cmp.Diff(tfedLight.(testLight).s, want); diff != "" {
		t.Fatalf("got: %v\n%s", tfedLight.(testLight).s, diff)
	}
}

func TestDivideSingle(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		want  lightid.Set
	}{
		{
			name:  "Simple",
			light: newTestLight(4),
			want:  lightid.Set{{1}, {2}, {3}, {4}},
		},
		{
			name:  "Single",
			light: newTestLight(1),
			want:  lightid.Set{{1}},
		},
		{
			name:  "Empty",
			light: newTestLight(0),
			want:  lightid.Set{},
		},
		{
			name:  "2LightID",
			light: newTestLightFromSet(lightid.Set{{1, 2}, {3, 4}}),
			want:  lightid.Set{{1}, {2}, {3}, {4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, DivideSingle())
		})
	}
}

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		light   context.Light
		divisor int
		want    lightid.Set
	}{
		{
			name:    "Simple1",
			light:   newTestLight(4),
			divisor: 3,
			want:    lightid.Set{{1}, {2}, {3, 4}},
		},
		{
			name:    "Simple2",
			light:   newTestLight(6),
			divisor: 3,
			want:    lightid.Set{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:    "Simple3",
			light:   newTestLight(7),
			divisor: 3,
			want:    lightid.Set{{1, 2}, {3, 4}, {5, 6, 7}},
		},
		{
			name:    "DivisorGreater",
			light:   newTestLight(3),
			divisor: 4,
			want:    lightid.Set{{}, {}, {}, {1, 2, 3}},
		},
		{
			name:    "Empty",
			light:   newTestLight(0),
			divisor: 4,
			want:    lightid.Set{{}, {}, {}, {}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Divide(tt.divisor))
		})
	}
}

func TestFan(t *testing.T) {
	tests := []struct {
		name       string
		light      context.Light
		groupCount int
		want       lightid.Set
	}{
		{
			name:       "Simple1",
			light:      newTestLight(4),
			groupCount: 2,
			want:       lightid.Set{{1, 3}, {2, 4}},
		},
		{
			name:       "Simple2",
			light:      newTestLight(5),
			groupCount: 2,
			want:       lightid.Set{{1, 3, 5}, {2, 4}},
		},
		{
			name:       "GroupSizeGreater",
			light:      newTestLight(3),
			groupCount: 4,
			want:       lightid.Set{{1}, {2}, {3}, {}},
		},
		{
			name:       "Empty",
			light:      newTestLight(0),
			groupCount: 4,
			want:       lightid.Set{{}, {}, {}, {}},
		},
		{
			name:       "2IDs",
			light:      newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			groupCount: 2,
			want:       lightid.Set{{1, 3}, {2, 4}, {5, 7}, {6, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Fan(tt.groupCount))
		})
	}
}

func TestDivideIntoGroups(t *testing.T) {
	tests := []struct {
		name      string
		light     context.Light
		groupSize int
		want      lightid.Set
	}{
		{
			name:      "Simple1",
			light:     newTestLight(6),
			groupSize: 2,
			want:      lightid.Set{{1, 2}, {3, 4}, {5, 6}},
		},
		{
			name:      "Simple2",
			light:     newTestLight(5),
			groupSize: 3,
			want:      lightid.Set{{1, 2, 3}, {4, 5}},
		},
		{
			name:      "GroupSizeGreater",
			light:     newTestLight(3),
			groupSize: 4,
			want:      lightid.Set{{1, 2, 3}},
		},
		{
			name:      "Empty",
			light:     newTestLight(0),
			groupSize: 4,
			want:      lightid.Set{{}},
		},
		{
			name:      "2IDs",
			light:     newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			groupSize: 3,
			want:      lightid.Set{{1, 2, 3}, {4}, {5, 6, 7}, {8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, DivideIntoGroups(tt.groupSize))
		})
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLight(4),
			want:  lightid.Set{{4, 3, 2, 1}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			want:  lightid.Set{{4, 3, 2, 1}, {8, 7, 6, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Reverse())
		})
	}
}

func TestReverseSet(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLight(4),
			want:  lightid.Set{{1, 2, 3, 4}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			want:  lightid.Set{{5, 6, 7, 8}, {1, 2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, ReverseSet())
		})
	}
}

func TestShuffle(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLight(6),
			want:  lightid.Set{{5, 6, 2, 3, 1, 4}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			want:  lightid.Set{{4, 2, 1, 3}, {8, 5, 7, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Shuffle(37))
		})
	}
}

func TestShuffleSet(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLight(6),
			want:  lightid.Set{{1, 2, 3, 4, 5, 6}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2}, {3, 4}, {5, 6}, {7, 8}}),
			want:  lightid.Set{{7, 8}, {3, 4}, {1, 2}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, ShuffleSet(37))
		})
	}
}

func TestRotate(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		n     int
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLight(6),
			n:     1,
			want:  lightid.Set{{6, 1, 2, 3, 4, 5}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			n:     -1,
			want:  lightid.Set{{2, 3, 4, 1}, {6, 7, 8, 5}},
		},
		{
			name:  "Simple3",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			n:     -4,
			want:  lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Rotate(tt.n))
		})
	}
}

func TestRotateSet(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		n     int
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLight(6),
			n:     1,
			want:  lightid.Set{{1, 2, 3, 4, 5, 6}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}),
			n:     -1,
			want:  lightid.Set{{4, 5, 6}, {7, 8, 9}, {1, 2, 3}},
		},
		{
			name:  "Simple3",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			n:     -4,
			want:  lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, RotateSet(tt.n))
		})
	}
}

func TestTake(t *testing.T) {
	tests := []struct {
		name    string
		light   context.Light
		indices []int
		want    lightid.Set
	}{
		{
			name:    "Simple1",
			light:   newTestLight(6),
			indices: []int{0, 2, 4},
			want:    lightid.Set{{1, 3, 5}},
		},
		{
			name:    "Simple2",
			light:   newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			indices: []int{0, 3},
			want:    lightid.Set{{1, 4}, {5, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Take(tt.indices...))
		})
	}
}

func TestTakeSet(t *testing.T) {
	tests := []struct {
		name    string
		light   context.Light
		indices []int
		want    lightid.Set
	}{
		{
			name:    "Simple1",
			light:   newTestLight(6),
			indices: []int{0},
			want:    lightid.Set{{1, 2, 3, 4, 5, 6}},
		},
		{
			name:    "Simple2",
			light:   newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			indices: []int{1},
			want:    lightid.Set{{5, 6, 7, 8}},
		},
		{
			name:    "MultiTake",
			light:   newTestLightFromSet(lightid.Set{{1, 2, 3, 4}}),
			indices: []int{0, 2},
			want:    lightid.Set{{1, 2, 3, 4}, {1, 2, 3, 4}},
		},
		{
			name:    "WrapAround",
			light:   newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			indices: []int{3},
			want:    lightid.Set{{5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, TakeSet(tt.indices...))
		})
	}
}

func TestTakeEvery(t *testing.T) {
	tests := []struct {
		name            string
		light           context.Light
		divisor, offset int

		want lightid.Set
	}{
		{
			name:    "Simple1",
			light:   newTestLight(4),
			divisor: 2, offset: 0,

			want: lightid.Set{{1, 3}},
		},
		{
			name:    "Simple2",
			light:   newTestLight(4),
			divisor: 2, offset: 1,

			want: lightid.Set{{2, 4}},
		},
		{
			name:    "Simple3",
			light:   newTestLight(4),
			divisor: 3, offset: 0,

			want: lightid.Set{{1, 4}},
		},
		{
			name:    "Simple4",
			light:   newTestLight(4),
			divisor: 3, offset: 1,

			want: lightid.Set{{2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, TakeEvery(tt.divisor, tt.offset))
		})
	}
}

func TestTakeEverySet(t *testing.T) {
	tests := []struct {
		name            string
		light           context.Light
		divisor, offset int

		want lightid.Set
	}{
		{
			name:    "Simple1",
			light:   newTestLightFromSet(lightid.Set{{1, 2}, {3, 4}, {5, 6}, {7, 8}}),
			divisor: 2, offset: 0,

			want: lightid.Set{{1, 2}, {5, 6}},
		},
		{
			name:    "Simple2",
			light:   newTestLightFromSet(lightid.Set{{1, 2}, {3, 4}, {5, 6}, {7, 8}}),
			divisor: 2, offset: 1,

			want: lightid.Set{{3, 4}, {7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, TakeEverySet(tt.divisor, tt.offset))
		})
	}
}
func TestSlice(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		i, j  int
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLight(6),
			i:     3, j: 5,

			want: lightid.Set{{4, 5}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			i:     1, j: 3,

			want: lightid.Set{{2, 3}, {6, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Slice(tt.i, tt.j))
		})
	}
}

func TestSliceSet(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		i, j  int
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLightFromSet(lightid.Set{{1, 2}, {3, 4}, {5, 6}}),
			i:     1, j: 3,

			want: lightid.Set{{3, 4}, {5, 6}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			i:     0, j: 2,

			want: lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, SliceSet(tt.i, tt.j))
		})
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		name  string
		light context.Light
		want  lightid.Set
	}{
		{
			name:  "Simple1",
			light: newTestLightFromSet(lightid.Set{{1, 2}, {3, 4}, {5, 6}}),
			want:  lightid.Set{{1, 2, 3, 4, 5, 6}},
		},
		{
			name:  "Simple2",
			light: newTestLightFromSet(lightid.Set{{1, 2, 3, 4}, {5, 6, 7, 8}}),
			want:  lightid.Set{{1, 2, 3, 4, 5, 6, 7, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			transformAndAssertLight(t, tt.light, tt.want, Flatten())
		})
	}
}
