package chroma

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLightID_MarshalJSON(t *testing.T) {
	testCases := []struct {
		lightID      LightID
		expectedJSON string
	}{
		{[]int{1}, "1"},
		{[]int{1, 2}, "[1,2]"},
	}

	for _, tt := range testCases {
		j, err := json.Marshal(tt.lightID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(tt.expectedJSON, string(j)); diff != "" {
			t.Fatal(diff)
		}
	}
}

func TestLightID_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		json            string
		expectedLightID LightID
	}{
		{`1`, []int{1}},
		{`[1,2]`, []int{1, 2}},
	}

	for _, tt := range testCases {
		var lightID LightID
		err := json.Unmarshal([]byte(tt.json), &lightID)
		if err != nil {
			t.Fatal(err)
		}
		if diff := cmp.Diff(tt.expectedLightID, lightID); diff != "" {
			t.Fatal(diff)
		}
	}
}
