package ilysa

//func BenchmarkFilterEvents(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		filterEvents()
//	}
//}
//
//func filterEvents() {
//	mockMap, _ := beatsaber.NewMockMap(beatsaber.EnvironmentOrigins, 120, "[]")
//	p := New(mockMap)
//	ctx := p.BOffset(0)
//
//	ctx.BeatRange(0, 100, 1000, ease.Linear, func(ctx context.Context) {
//		ctx.NewRGBLighting()
//	})
//
//	for i := 0; i <= 10; i++ {
//		bos := float64(i * 10)
//
//		p.FilterEvents(func(e evt.Event) bool {
//			if e.Beat() < 90+bos+5 {
//				return true
//			}
//			return false
//		})
//
//		ctx.BeatRange(90+bos, 100+bos, 100, ease.Linear, func(ctx context.Context) {
//			ctx.NewRGBLighting()
//		})
//	}
//}
//
