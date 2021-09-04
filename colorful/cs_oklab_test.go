package colorful_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/shasderias/ilysa/colorful"
)

type oklab struct {
	L, A, B float64
}

func TestOklab(t *testing.T) {
	testCases := []struct {
		RGB     colorful.Color
		L, A, B float64
	}{
		{
			colorful.Color{1, 0, 0, 1},
			0.6279182221983763, 0.22483566700277774, 0.12578875827934327,
		},
	}
	for _, tt := range testCases {
		l, a, b := tt.RGB.Oklab()
		if diff := cmp.Diff(oklab{l, a, b}, oklab{tt.L, tt.A, tt.B},
			cmpopts.EquateApprox(0.001, 0),
		); diff != "" {
			t.Fatal(diff)
		}
	}
}
