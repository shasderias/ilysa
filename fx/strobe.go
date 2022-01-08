package fx

//func Strobe(ctx context.LightContext, grad gradient.Table, onAlpha, offAlpha float64) {
//	a := onAlpha
//	if ctx.Ordinal()%2 == 1 {
//		a = offAlpha
//	}
//	ctx.NewRGBLighting(
//		evt.WithColor(grad.Lerp(ctx.T())),
//		opt.Alpha(a),
//	)
//}
//
//func ABColorStrobe(ctx context.LightContext, aColorSet, bColorSet colorful.Set, onAlpha, offAlpha float64) {
//	if ctx.Ordinal()%2 == 0 {
//		ctx.NewRGBLighting(
//			evt.WithColor(aColorSet.Next()),
//			opt.Alpha(onAlpha),
//		)
//	} else {
//		ctx.NewRGBLighting(
//			evt.WithColor(bColorSet.Next()),
//			opt.Alpha(offAlpha),
//		)
//	}
//}
