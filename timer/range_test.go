package timer

import (
	"fmt"

	"github.com/shasderias/ilysa/ease"
)

func ExampleRng() {
	rng := Rng(0, 1, 5, ease.Linear)
	iter := rng.Iterate()
	for iter.Next() {
		fmt.Println(iter.B())
	}
	// Output:
	// 0
	// 0.25
	// 0.5
	// 0.75
	// 1
}

func ExampleRngInterval() {
	rng := RngInterval(1, 2, 4, ease.Linear)
	iter := rng.Iterate()
	for iter.Next() {
		fmt.Println(iter.B())
	}
	// Output:
	// 1
	// 1.25
	// 1.5
	// 1.75
	// 2
}
