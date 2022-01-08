package timer

import "fmt"

func ExampleSeq() {
	seq := Seq(1, 2, 3, 4, 4+1)
	iter := seq.Iterate()
	for iter.Next() {
		fmt.Println(iter.ToRange().B())
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
}

func ExampleSeqInterval() {
	seq := SeqInterval(1, 3, 4)
	iter := seq.Iterate()
	for iter.Next() {
		fmt.Println(iter.ToRange().B())
	}
	// Output:
	// 1
	// 1.25
	// 1.5
	// 1.75
	// 2
	// 2.25
	// 2.5
	// 2.75
	// 3
}
