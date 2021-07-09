package fx

//func RainbowProp(p rework.Context, light light2.Light, grad gradient.Table, startBeat, duration, step float64, frames int) {
//	p.Range(startBeat, startBeat+duration, frames, ease.Linear, func(ctx context.Context) {
//		ctx.Light(light, func(ctx context.LightContext) {
//			e := ctx.NewRGBLightingEvent(
//				evt.WithColor(grad.Lerp(ctx.T())),
//			)
//			Ripple(ctx, e, step)
//			AlphaFadeEx(ctx, e, 0.0, 0.4, 0, 1, ease.InCirc)
//			AlphaFadeEx(ctx, e, 0.4, 1, 1, 0, ease.OutCirc)
//		})
//	})
//}
