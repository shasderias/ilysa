package transform_test

import (
	"fmt"
	"testing"

	"github.com/shasderias/ilysa/context"
	"github.com/shasderias/ilysa/evt"
	"github.com/shasderias/ilysa/light"
	"github.com/shasderias/ilysa/transform"
)

func TestSanity(t *testing.T) {
	proj := context.NewMockProject(t, 3)

	bl := light.NewBasic(proj, evt.BackLasers)

	blSingle := transform.Light(bl,
		transform.DivideSingle(),
	)

	fmt.Println(blSingle)

	blSequence := transform.Light(bl,
		transform.Identity(),
		transform.DivideSingle().Sequence(),
	)

	fmt.Println(blSequence)
}
