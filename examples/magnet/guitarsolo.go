package main

//type GuitarSolo struct {
//	p *ilysa.Project
//	rework.Context
//}
//
//func NewGuitarSolo(p *ilysa.Project, startBeat float64) GuitarSolo {
//	return GuitarSolo{
//		p:           p,
//		Context: p.Offset(startBeat),
//	}
//}
//
//func (g GuitarSolo) Play() {
//	g.EventForBeat(0, func(ctx context.Context) {
//		ctx.NewPreciseLaser(
//			evt.WithDirectionalLaser(evt.LeftLaser),
//			evt.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
//		)
//		ctx.NewPreciseLaser(
//			evt.WithDirectionalLaser(evt.RightLaser),
//			evt.WithIntValue(3), evt.WithPreciseLaserSpeed(4.5),
//		)
//	})
//
//	g.Beat(0)
//	g.Beat(4)
//	g.Beat(8)
//
//	g.Solo(0.50, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25}, false)
//	g.Solo(2.25, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50}, true)
//	g.Solo(4.25, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50}, false)
//	g.Solo(6.25, []float64{0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50, 1.75}, true)
//	g.Solo(8.50, []float64{
//		0.00, 0.25, 0.50, 0.75, 1.00, 1.25, 1.50, 1.75,
//		2.00, 2.25, 2.50, 2.75, 3.00, 3.25, 3.50,
//	}, false)
//	g.Solo(12.50, []float64{0.00, 0.25}, true)
//	g.Solo(13.25, []float64{0.00, 0.25}, false)
//
//}
//
//func (g GuitarSolo) Beat(startBeat float64) {
//	ctx := g.Offset(startBeat)
//
//	bl := transform.Light(
//		light.NewBasic(beatsaber.EventTypeBackLasers, g),
//		rework.ToLightTransformer(rework.DivideSingle),
//	)
//
//	gradSet := gradient.NewSet(
//		shirayukiGradient,
//		sukoyaGradient,
//		shirayukiSingleGradient,
//		sukoyaSingleGradient,
//	)
//
//	ctx.EventsForBeats(0, 2, 4, func(ctx context.Context) {
//		ctx.NewPreciseRotation(
//			evt.WithRotation(15),
//			evt.WithRotationStep(15),
//			evt.WithProp(2),
//			evt.WithPreciseLaserSpeed(8),
//			evt.WithLaserDirection(chroma.CounterClockwise),
//		)
//
//		step := -0.5
//		if ctx.Ordinal() %2 == 0 {
//			step = 0.5
//		}
//
//		ctx.NewPreciseZoom(evt.WithRotationStep(step))
//
//		grad := gradSet.Next()
//
//		ctx.Range(ctx.B(), ctx.B()+0.50, 12, ease.Linear, func(ctx context.Context) {
//			ctx.Light(bl, func(ctx context.LightContext) {
//				e := fx.ColorSweep(ctx, 2.4, grad)
//				fx.AlphaBlend(ctx, e, 0, 1, 1.5, 0, ease.InBounce)
//			})
//		})
//
//	})
//}
//
//func (g GuitarSolo) Solo(startBeat float64, sequence []float64, reverse bool) {
//	ctx := g.Offset(startBeat)
//
//	var (
//		llReverse rework.LightIDTransformer = rework.Shuffle
//		rlReverse rework.LightIDTransformer = rework.Shuffle
//	)
//
//	if reverse {
//		llReverse = rework.Reverse
//		rlReverse = rework.Identity
//	} else {
//		llReverse = rework.Identity
//		rlReverse = rework.Reverse
//	}
//
//	ll := transform.Light(
//		light.NewBasic(beatsaber.EventTypeLeftRotatingLasers, g),
//		rework.ToLightTransformer(llReverse),
//		rework.ToSequenceLightTransformer(rework.DivideSingle),
//	).(light2.SequenceLight)
//	rl := transform.Light(
//		light.NewBasic(beatsaber.EventTypeRightRotatingLasers, g),
//		rework.ToLightTransformer(rlReverse),
//		rework.ToSequenceLightTransformer(rework.DivideSingle),
//	).(light2.SequenceLight)
//	light := light2.NewSequenceLight()
//
//	for i := 0; i < ll.Len(); i++ {
//		light.Add(light2.NewCombinedLight(ll.Idx(i), rl.Idx(i)))
//	}
//
//	ctx.Sequence(timer.Beat(0, func(ctx context.Context) {
//		//ctx.NewPreciseLaser(
//		//	ilysa.WithDirectionalLaser(evt.LeftLaser),
//		//	ilysa.WithPreciseLaserSpeed(0),
//		//	ilysa.WithLaserSpeed(6),
//		//)
//		//ctx.NewPreciseLaser(
//		//	ilysa.WithDirectionalLaser(evt.RightLaser),
//		//	ilysa.WithPreciseLaserSpeed(0),
//		//	ilysa.WithLaserSpeed(6),
//		//)
//
//	})
//
//	ctx.EventsForSequence(0, sequence, func(ctx rework.SequenceContext) {
//		seqCtx := ctx
//
//		var (
//			llLock, rlLock                           = false, true
//			llSpeed, llIntValue, rlSpeed, rlIntValue = 0.0, 5, 5.0, 5
//			//llAlpha, rlAlpha = 3.0, 1.0
//		)
//
//		if reverse {
//			llLock, rlLock = true, false
//			llSpeed, llIntValue, rlSpeed, rlIntValue = 5.0, 5, 0.0, 5
//			//llAlpha, rlAlpha = 1.0, 3.0
//		}
//
//		ctx.NewPreciseLaser(
//			evt.WithDirectionalLaser(evt.LeftLaser),
//			evt.WithLockPosition(llLock),
//			evt.WithPreciseLaserSpeed(llSpeed),
//			evt.WithIntValue(llIntValue),
//		)
//		ctx.NewPreciseLaser(
//			evt.WithDirectionalLaser(evt.RightLaser),
//			evt.WithLockPosition(rlLock),
//			evt.WithPreciseLaserSpeed(rlSpeed),
//			evt.WithIntValue(rlIntValue),
//		)
//
//		ctx.EventsForRange(ctx.B(), ctx.B()+1.25, 30, ease.Linear, func(ctx context.Context) {
//			ctx.Light(light.Idx(seqCtx.Ordinal()), func(ctx context.LightContext) {
//				e := fx.ColorSweep(ctx, 1.9, magnetRainbowPale)
//				//e := ctx.NewRGBLighting(evt.WithColor(color))
//				fx.AlphaBlend(ctx, e, 0, 1, 3, 0, ease.InCubic)
//			})
//		})
//	})
//}
