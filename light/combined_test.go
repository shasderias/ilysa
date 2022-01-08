package light_test

//func TestCombine(t *testing.T) {
//	l6 := light.NewCustom(evt.RingLights, 6, 0)
//	l2 := light.NewCustom(evt.CenterLights, 4, 0)
//
//	fmt.Println(l6, l2, l6.lightNameLen(), l2.lightNameLen())
//
//	cl := light.Combine(l6, l2)
//
//	fmt.Println(cl.lightNameLen())
//
//	l := transform.Light(cl,
//		transform.DivideSingle(),
//	)
//
//	fmt.Println(l)
//
//	m, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
//
//	p := ilysa.New(m)
//
//	ctx := p.BOffset(0)
//
//	ctx.Beat(1, func(ctx context.Context) {
//		ctx.Light(l, func(ctx context.LightContext) {
//			ctx.NewRGBLighting()
//		})
//	})
//
//	for _, me := range *p.Events() {
//		e := me.(*evt.ChromaLighting)
//		fmt.Println(e.Beat(), e.Type(), e.LightID)
//	}
//}
