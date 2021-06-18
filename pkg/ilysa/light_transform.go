package ilysa

type LightTransformer func(Light) Light

func TransformLight(l Light, tfers ...LightTransformer) Light {
	for _, tf := range tfers {
		l = tf(l)
	}
	return l
}

//func FanLight(groupSize int) LightTransformer {
//	return func(l Light) Light {
//		bl, ok := l.(BasicLight)
//		if !ok {
//			return l
//		}
//
//	}
//}
//
