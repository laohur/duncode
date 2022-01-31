package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Println("Hello, 世界")

	// dat, err := os.ReadFile("Blocks.txt")

	fmt.Print("done")

}
