package duncode

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
	fmt.Printf("duncode :%s --> %s\n", s, t)
	fmt.Print("done")

}
