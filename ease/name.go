package ease

import (
	"fmt"
)

func (f *Func) EaseName() string {
	switch {
	case f == &Linear:
		return "easeLinear"
	case f == &Step:
		return "easeStep"
	case f == &InQuad:
		return "easeInQuad"
	case f == &OutQuad:
		return "easeOutQuad"
	case f == &InOutQuad:
		return "easeInOutQuad"
	case f == &InCubic:
		return "easeInCubic"
	case f == &OutCubic:
		return "easeOutCubic"
	case f == &InOutCubic:
		return "easeInOutCubic"
	case f == &InQuart:
		return "easeInQuart"
	case f == &OutQuart:
		return "easeOutQuart"
	case f == &InOutQuart:
		return "easeInOutQuart"
	case f == &InQuint:
		return "easeInQuint"
	case f == &OutQuint:
		return "easeOutQuint"
	case f == &InOutQuint:
		return "easeInOutQuint"
	case f == &InSin:
		return "easeInSine"
	case f == &OutSin:
		return "easeOutSine"
	case f == &InOutSin:
		return "easeInOutSine"
	case f == &InExpo:
		return "easeInExpo"
	case f == &OutExpo:
		return "easeOutExpo"
	case f == &InOutExpo:
		return "easeInOutExpo"
	case f == &InCirc:
		return "easeInCirc"
	case f == &OutCirc:
		return "easeOutCirc"
	case f == &InOutCirc:
		return "easeInOutCirc"
	case f == &InBack:
		return "easeInBack"
	case f == &OutBack:
		return "easeOutBack"
	case f == &InOutBack:
		return "easeInOutBack"
	case f == &InElastic:
		return "easeInElastic"
	case f == &OutElastic:
		return "easeOutElastic"
	case f == &InOutElastic:
		return "easeInOutElastic"
	case f == &InBounce:
		return "easeInBounce"
	case f == &OutBounce:
		return "easeOutBounce"
	case f == &InOutBounce:
		return "easeInOutBounce"
	default:
		panic(fmt.Sprintf("EaseName(): unsupported ease: %v", f))
	}
}
