package chroma

import (
	"reflect"
	"testing"

	"ilysa/pkg/chroma/lightid"
)

func TestLightIDDiv(t *testing.T) {
	type args struct {
		min       int
		max       int
		div       int
		remainder int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "AllIndividual",
			args: args{
				min:       1,
				max:       4,
				div:       1,
				remainder: 0,
			},
			want: []int{1, 2, 3, 4},
		},
		{
			name: "Even",
			args: args{
				min:       1,
				max:       4,
				div:       2,
				remainder: 0,
			},
			want: []int{2, 4},
		},
		{
			name: "Odd",
			args: args{
				min:       1,
				max:       4,
				div:       2,
				remainder: 1,
			},
			want: []int{1, 3},
		},
		{
			name: "Every 3rd ID",
			args: args{
				min:       1,
				max:       4,
				div:       3,
				remainder: 0,
			},
			want: []int{3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := light.EveryNthLightID(tt.args.min, tt.args.max, tt.args.div, tt.args.remainder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EveryNthLightID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLightIDRange(t *testing.T) {
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1-4",
			args: args{
				start: 1,
				end:   4,
			},
			want: []int{1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := light.MakeInterval(tt.args.start, tt.args.end); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		lightIDs []int
		divisor  int
	}
	tests := []struct {
		name string
		args args
		want light.LightIDSet
	}{
		{
			name: "Simple",
			args: args{
				lightIDs: []int{1, 2, 3},
				divisor:  3,
			},
			want: light.LightIDSet{[]int{1}, []int{2}, []int{3}},
		},
		{
			name: "Simple - Overflow",
			args: args{
				lightIDs: []int{1, 2, 3, 4},
				divisor:  3,
			},
			want: light.LightIDSet{[]int{1}, []int{2}, []int{3}},
		},
		{
			name: "Double",
			args: args{
				lightIDs: []int{1, 2, 3, 4, 5, 6},
				divisor:  3,
			},
			want: light.LightIDSet{[]int{1, 2}, []int{3, 4}, []int{5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := light.Divide(tt.args.lightIDs, tt.args.divisor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAlternate(t *testing.T) {
	type args struct {
		lightIDs []int
		groups   int
	}
	tests := []struct {
		name string
		args args
		want light.LightIDSet
	}{
		{
			name: "Simple",
			args: args{
				lightIDs: []int{1, 2, 3},
				groups:   3,
			},
			want: [][]int{{1}, {2}, {3}},
		},
		{
			name: "Simple - Overflow",
			args: args{
				lightIDs: []int{1, 2, 3, 4},
				groups:   3,
			},
			want: [][]int{{1, 4}, {2}, {3}},
		},
		{
			name: "Two A-Piece",
			args: args{
				lightIDs: []int{1, 2, 3, 4, 5, 6},
				groups:   3,
			},
			want: [][]int{{1, 4}, {2, 5}, {3, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := light.Fan(tt.args.lightIDs, tt.args.groups); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Fan() = %v, want %v", got, tt.want)
			}
		})
	}
}
