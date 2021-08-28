package evt

import (
	"fmt"
	"image/color"
	"reflect"
	"runtime"
	"strings"

	"github.com/shasderias/ilysa/ease"
)

func (o withLightOpt) applyChromaGradient(g *ChromaGradient) {
	g.SetType(o.l)
}

type withDurationOpt struct {
	d float64
}

func WithDuration(beats float64) withDurationOpt {
	return withDurationOpt{d: beats}
}

func (o withDurationOpt) applyChromaGradient(g *ChromaGradient) {
	g.Gradient.LightGradient.Duration = o.d
}

type withStartColorOpt struct {
	c color.Color
}

func WithStartColor(c color.Color) withStartColorOpt {
	return withStartColorOpt{c}
}

func (o withStartColorOpt) applyChromaGradient(g *ChromaGradient) {
	g.Gradient.LightGradient.StartColor = o.c
}

type withEndColorOpt struct {
	c color.Color
}

func WithEndColor(c color.Color) withEndColorOpt {
	return withEndColorOpt{c}
}

func (o withEndColorOpt) applyChromaGradient(g *ChromaGradient) {
	g.Gradient.LightGradient.EndColor = o.c
}

type withEasingOpt struct {
	e string
}

func nameOf(f interface{}) string {
	v := reflect.ValueOf(f)
	if v.Kind() == reflect.Func {
		if rf := runtime.FuncForPC(v.Pointer()); rf != nil {
			return rf.Name()
		}
	}
	return v.String()
}

func WithEasing(easeFn ease.Func) withEasingOpt {
	easeFnName := nameOf(easeFn)

	switch {
	case strings.Contains(easeFnName, "ease.Linear"):
		easeFnName = "easeLinear"
	case strings.Contains(easeFnName, "ease.InQuad"):
		easeFnName = "easeInQuad"
	case strings.Contains(easeFnName, "ease.OutQuad"):
		easeFnName = "easeOutQuad"
	case strings.Contains(easeFnName, "ease.InOutQuad"):
		easeFnName = "easeInOutQuad"
	case strings.Contains(easeFnName, "ease.InCubic"):
		easeFnName = "easeInCubic"
	case strings.Contains(easeFnName, "ease.OutCubic"):
		easeFnName = "easeOutCubic"
	case strings.Contains(easeFnName, "ease.InOutCubic"):
		easeFnName = "easeInOutCubic"
	case strings.Contains(easeFnName, "ease.InQuart"):
		easeFnName = "easeInQuart"
	case strings.Contains(easeFnName, "ease.OutQuart"):
		easeFnName = "easeOutQuart"
	case strings.Contains(easeFnName, "ease.InOutQuart"):
		easeFnName = "easeInOutQuart"
	case strings.Contains(easeFnName, "ease.InQuint"):
		easeFnName = "easeInQuint"
	case strings.Contains(easeFnName, "ease.OutQuint"):
		easeFnName = "easeOutQuint"
	case strings.Contains(easeFnName, "ease.InOutQuint"):
		easeFnName = "easeInOutQuint"
	case strings.Contains(easeFnName, "ease.InSin"):
		easeFnName = "easeInSine"
	case strings.Contains(easeFnName, "ease.OutSin"):
		easeFnName = "easeOutSine"
	case strings.Contains(easeFnName, "ease.InOutSin"):
		easeFnName = "easeInOutSine"
	case strings.Contains(easeFnName, "ease.InExpo"):
		easeFnName = "easeInExpo"
	case strings.Contains(easeFnName, "ease.OutExpo"):
		easeFnName = "easeOutExpo"
	case strings.Contains(easeFnName, "ease.InOutExpo"):
		easeFnName = "easeInOutExpo"
	case strings.Contains(easeFnName, "ease.InCirc"):
		easeFnName = "easeInCirc"
	case strings.Contains(easeFnName, "ease.OutCirc"):
		easeFnName = "easeOutCirc"
	case strings.Contains(easeFnName, "ease.InOutCirc"):
		easeFnName = "easeInOutCirc"
	case strings.Contains(easeFnName, "ease.InBack"):
		easeFnName = "easeInBack"
	case strings.Contains(easeFnName, "ease.OutBack"):
		easeFnName = "easeOutBack"
	case strings.Contains(easeFnName, "ease.InOutBack"):
		easeFnName = "easeInOutBack"
	case strings.Contains(easeFnName, "ease.InElastic"):
		easeFnName = "easeInElastic"
	case strings.Contains(easeFnName, "ease.OutElastic"):
		easeFnName = "easeOutElastic"
	case strings.Contains(easeFnName, "ease.InOutElastic"):
		easeFnName = "easeInOutElastic"
	case strings.Contains(easeFnName, "ease.InBounce"):
		easeFnName = "easeInBounce"
	case strings.Contains(easeFnName, "ease.OutBounce"):
		easeFnName = "easeOutBounce"
	case strings.Contains(easeFnName, "ease.InOutBounce"):
		easeFnName = "easeInOutBounce"
	default:
		panic(fmt.Sprintf("WithEasing: unsupported ease: %v", easeFn))
	}

	return withEasingOpt{e: easeFnName}
}

func (o withEasingOpt) applyChromaGradient(g *ChromaGradient) {
	g.Gradient.LightGradient.Easing = o.e
}
