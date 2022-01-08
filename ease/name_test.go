package ease_test

import (
	"testing"

	"github.com/shasderias/ilysa/ease"
)

func TestEaseName(t *testing.T) {
	if ease.Linear.EaseName() != "easeLinear" {
		t.Errorf("got %s; want %s", ease.Linear.EaseName(), "easeLinear")
	}
}
