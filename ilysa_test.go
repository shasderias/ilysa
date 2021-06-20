package ilysa

//func TestEventsForRange(t *testing.T) {
//	type args struct {
//		startBeat float64
//		endBeat   float64
//		steps     int
//		easeFunc  ease.Func
//	}
//	tests := []struct {
//		name           string
//		args           args
//		callbackOnBeat []float64
//	}{
//		{
//			name: "Basic",
//			args: args{
//				startBeat: 1,
//				endBeat:   3,
//				steps:     3,
//				easeFunc:  ease.Linear,
//			},
//			callbackOnBeat: []float64{1, 2, 3},
//		},
//		{
//			name: "Decimal",
//			args: args{
//				startBeat: 1,
//				endBeat:   2,
//				steps:     3,
//				easeFunc:  ease.Linear,
//			},
//			callbackOnBeat: []float64{1, 1.5, 2},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &Project{}
//			callbackCount := 0
//			p.EventsForRange(tt.args.startBeat, tt.args.endBeat, tt.args.steps, tt.args.easeFunc, func(ctx Timer) {
//				if ctx.B != tt.callbackOnBeat[callbackCount] {
//					t.Fatalf("got %dth callback on beat %f; want callback on beat %f instead",
//						callbackCount, ctx.B, tt.callbackOnBeat[callbackCount],
//					)
//				}
//				callbackCount++
//			})
//		})
//	}
//}
//
//func TestEventsForBeats(t *testing.T) {
//	type args struct {
//		startBeat float64
//		duration  float64
//		count     int
//	}
//	tests := []struct {
//		name           string
//		args           args
//		callbackOnBeat []float64
//	}{
//		{
//			name: "Basic",
//			args: args{
//				startBeat: 1,
//				duration:  1,
//				count:     4,
//			},
//			callbackOnBeat: []float64{1, 2, 3, 4},
//		},
//		{
//			name: "Basic Decimal",
//			args: args{
//				startBeat: 1,
//				duration:  0.25,
//				count:     4,
//			},
//			callbackOnBeat: []float64{1, 1.25, 1.50, 1.75},
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &Project{}
//			callbackCount := 0
//			p.EventsForBeats(tt.args.startBeat, tt.args.duration, tt.args.count, func(ctx Timer) {
//				if ctx.B != tt.callbackOnBeat[callbackCount] {
//					t.Fatalf("got %dth callback on beat %f; want callback on beat %f instead",
//						callbackCount, ctx.B, tt.callbackOnBeat[callbackCount],
//					)
//				}
//				callbackCount++
//			})
//		})
//	}
//}
