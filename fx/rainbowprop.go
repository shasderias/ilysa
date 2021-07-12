package fx

//func RainbowProp(p context.Context, rng timer.Ranger, l context.Light, grad gradient.Table, opts ... RainbowPropOpt) {
//	p.Range(rng, func(ctx context.Context) {
//		ctx.Light(l, func(ctx context.LightContext) {
//			e := ctx.NewRGBLighting(
//				evt.WithColor(grad.Lerp(ctx.T())),
//			)
//			Ripple(ctx, e, step)
//			AlphaFadeEx(ctx, e, 0.0, 0.4, 0, 1, ease.InCirc)
//			AlphaFadeEx(ctx, e, 0.4, 1, 1, 0, ease.OutCirc)
//		})
//	})
//}
//
//type RainbowPropOpt interface {
//	apply()
//}
