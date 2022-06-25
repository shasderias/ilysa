package fx

//func ZeroSpeedRandomizedLasers(ctx context.Context, l evt.DirectionalLaser) *evt.ChromaLaserSpeed {
//	return ctx.NewPreciseLaser(
//		evt.WithDirectionalLaser(l),
//		evt.WithIntValue(1),
//		evt.WithPreciseLaserSpeed(0),
//	)
//}
//
//
//type slowMotionLaserConfig struct {
//	lockFirst bool
//}
//
//type SlowMotionLaserOpt interface {
//	apply(*slowMotionLaserConfig)
//}
//
//func OptSlowMotionLasersLockFirst() SlowMotionLaserOpt {
//	return slowMotionLaserLockFirst{}
//}
//
//type slowMotionLaserLockFirst struct{}
//
//func (o slowMotionLaserLockFirst) apply(c *slowMotionLaserConfig) {
//	c.lockFirst = true
//}
//
//func SlowMotionLasers(ctx context.Context,
//	rng timer.Ranger, t evt.Type, startSpeed, endSpeed float64, easeFn ease.Func,
//	opts ...SlowMotionLaserOpt) evt.Events {
//	conf := slowMotionLaserConfig{}
//	for _, o := range opts {
//		o.apply(&conf)
//	}
//
//	events := evt.NewEvents()
//
//	speedScaler := scale.FromUnitClamp(startSpeed, endSpeed)
//
//	ctx.WRng(rng, func(ctx context.Context) {
//		speed := speedScaler(easeFn(ctx.T()))
//		intSpeed := int(math.Round(speed))
//
//		optSet := opt.NewSet(
//			ctx,
//			evt.OptType(t),
//			evt.OptIntValue(intSpeed),
//			evt.OptChromaLaserSpeed(speed),
//		)
//
//		if conf.lockFirst && ctx.First() {
//			optSet.Add(evt.OptChromaLaserSpeedLockPosition(true))
//		}
//		if !ctx.First() {
//			optSet.Add(evt.OptChromaLaserSpeedLockPosition(true))
//		}
//		e := evt.NewChromaLaserSpeed(optSet)
//		events.Add(e)
//	})
//
//	return events
//}
//
//// n.b. does not work on nice and possibly other environments
//func SlowMotionLasers2(ctx context.Context,
//	rng timer.Ranger, startSpeed, endSpeed float64, easeFn ease.Func,
//	opts ...SlowMotionLaserOpt) evt.Events {
//
//	events := evt.NewEvents()
//
//	speedScaler := scale.FromUnitClamp(startSpeed, endSpeed)
//
//	ctx.WRng(rng, func(ctx context.Context) {
//		speed := speedScaler(easeFn(ctx.T()))
//		intSpeed := int(math.Round(speed))
//
//		e1 := evt.NewChromaLaserSpeed(ctx,
//			evt.OptType(evt.TypeLeftLaserSpeed),
//			evt.OptChromaLaserSpeed(speed),
//			evt.OptIntValue(intSpeed),
//			opt.AllExceptFirst(ctx, evt.OptChromaLaserSpeedLockPosition(true)),
//		)
//
//		e2 := evt.NewChromaLaserSpeed(ctx,
//			evt.OptType(evt.TypeRightLaserSpeed),
//			evt.OptChromaLaserSpeed(speed),
//			evt.OptIntValue(intSpeed),
//			opt.AllExceptFirst(ctx, evt.OptChromaLaserSpeedLockPosition(true)),
//		)
//
//		events.Add(e1, e2)
//	})
//
//	return events
//}
//
//
//type slowMotionLasersOpts struct {
//	preciseLaserOpts Options
//}
//
//type slowMotionLasersOpt interface {
//	apply(opt *slowMotionLasersOpts)
//}
//
//type withPreciseLaserOpt struct {
//	opts Options
//}
//
//func WithPreciseLaserOpts(opts ...Option) slowMotionLasersOpt {
//	return withPreciseLaserOpt{opts}
//}
//
//func (o withPreciseLaserOpt) apply(opt *slowMotionLasersOpts) {
//	opt.preciseLaserOpts = o.opts
//}
