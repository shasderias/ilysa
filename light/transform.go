package light

import (
	"github.com/shasderias/ilysa/context"
)

type Transformer interface {
	Transform(context.Light) context.Light
}

//
//func abc() {
//	light.Transform(
//		backLasers,
//		transform.DivideSingle,
//		transform.DivideSingle.Sequence(),
//	)
//}
