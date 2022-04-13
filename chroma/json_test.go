package chroma

import (
	"testing"

	"github.com/shasderias/ilysa/internal/null"
)

func TestMarshalToCustomData(t *testing.T) {
	testCases := []struct {
		name         string
		v            interface{}
		expectedJSON string
	}{
		{
			name: "InvalidByOmission",
			v: LaserSpeed{
				LockPosition: null.New(false),
				Direction:    null.New(Clockwise),
			},
			expectedJSON: `{"_direction":1,"_lockPosition":false}`,
		},
		{
			name: "Invalid",
			v: LaserSpeed{
				LockPosition: null.New(false),
				Speed:        null.New(0.0),
				Direction:    null.New(Clockwise),
			},
			expectedJSON: `{"_direction":1,"_lockPosition":false,"_speed":0}`,
		},
		{
			name: "Rounding5DP",
			v: LaserSpeed{
				Speed: null.New(0.12345),
			},
			expectedJSON: `{"_speed":0.123}`,
		},
		{
			name: "Rounding2DP",
			v: LaserSpeed{
				Speed: null.New(0.12),
			},
			expectedJSON: `{"_speed":0.12}`,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			b, err := marshalToCustomData(tt.v)
			if err != nil {
				t.Fatal(err)
			}
			if string(b) != tt.expectedJSON {
				t.Errorf("want %s; got %s", tt.expectedJSON, string(b))
			}
		})
	}
}
