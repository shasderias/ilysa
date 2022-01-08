package fx

//// OffAll generates events to turn all lights/lasers off.
//func OffAll(ctx context.Context) evt.Events {
//	var (
//		lights = []evt.LightType{
//			evt.BackLasers,
//			evt.RingLights,
//			evt.LeftRotatingLasers,
//			evt.RightRotatingLasers,
//			evt.CenterLights,
//		}
//	)
//
//	events := evt.Events{}
//
//	for _, l := range lights {
//		events = append(events, ctx.NewLighting(evt.WithLight(l), evt.WithLightValue(evt.LightOff)))
//	}
//
//	return events
//}
