package main

import (
	"fmt"
)

func demo() {
	var s = "Aÿぃ好乇αβЖѰक़ꌊ걹"
	// testDuncode(s)
	// testDuncodeCompress(s)
	// testLine(s)
	var bytes = Encode(s)
	var t = Decode(bytes)
	fmt.Printf("duncode :%s %d  --> %s %d\n", s, len(s), t, len(bytes))
	fmt.Print("done")

}
