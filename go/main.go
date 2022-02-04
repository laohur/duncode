package main

import (
	"fmt"
)

func main() {
	var s = "Aÿぃ好乇αβЖѰक़ꌊ걹"
	// testDuncode(s)
	// testDuncodeCompress(s)
	// testLine(s)
	var bytes = encode(s)
	var t = decode(bytes)
	fmt.Printf("duncode :%s --> %s\n", s, t)
	fmt.Print("done")

}
