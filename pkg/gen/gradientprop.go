package gen

//func GradientProp(ctx ilysa.Context, typ beatsaber.EventType, val beatsaber.EventValue, table gradient.Table, LightIDMin, LightIDMax int) {
//	scale := util.Scale(float64(LightIDMin), float64(LightIDMax), 0, 1)
//	for i := LightIDMin; i < LightIDMax; i++ {
//		e := ctx.NewRGBLightingEvent(typ, val)
//		e.SetSingleLightID(i)
//		e.SetColor(table.GetInterpolatedColorFor(scale(float64(i))))
//	}
//}
