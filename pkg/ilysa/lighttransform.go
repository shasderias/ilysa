package ilysa

type LightTransformer func(Light) Light
type LightIDTransformer func(id LightID) LightIDSet

type LightIDSetter interface {
	LightIDSet() LightIDSet
}

type LightIDTransformable interface {
	LightIDTransform(LightIDTransformer) Light
}

func DivideBasicLight(divisor int) func(bl BasicLight) []CompositeLight {
	return func(bl BasicLight) []CompositeLight {
		set := Divide(divisor)(NewLightIDFromInterval(1, bl.maxLightID))

		lights := make([]CompositeLight, divisor)

		for i := 0; i < divisor; i++ {
			lights[i] = CompositeLight{
				eventType: bl.eventType,
				set:       LightIDSet{set[i]},
			}
		}

		return lights
	}
}

func FanBasicLight(groupCount int) func(bl BasicLight) []CompositeLight {
	return func(bl BasicLight) []CompositeLight {
		set := Fan(groupCount)(NewLightIDFromInterval(1, bl.maxLightID))

		lights := make([]CompositeLight, groupCount)

		for i := 0; i < groupCount; i++ {
			lights[i] = CompositeLight{
				eventType: bl.eventType,
				set:       LightIDSet{set[i]},
			}
		}

		return lights
	}
}

//func TransformLight(l Light, tfers ...LightTransformer) Light {
//	for _, tf := range tfers {
//		l = tf(l)
//	}
//	return l
//}
//
//func FanLight(groupCount int) LightTransformer {
//	return func(l Light) Light {
//		bl, ok := l.(BasicLight)
//		if !ok {
//			return l
//		}
//
//	}
//}
//
