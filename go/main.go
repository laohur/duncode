package main

import (
	bytes2 "bytes"
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func testDuncode(s string) {
	for i, x := range []rune(s) {
		var a = rune2Duncode(x)
		var b = a.toBytes()
		var now = &Duncode{}
		now.readBytes(b)
		var c = now.toChars()
		fmt.Printf("testDuncode %d :" ,i)
		fmt.Println(string(c))
	}

}
func testDuncodeBytes(s string) {
	var buffer = bytes2.Buffer{}
	var last = &Duncode{}
	for _, x := range []rune(s) {
		fmt.Println(x)
		var now = rune2Duncode(x)
		if !last.compress(now) {
			var b = last.toBytes()
			buffer.Write(b)
			last = now
		}
		// else continue
	}
	var b = last.toBytes()
	buffer.Write(b)
}

func encode(s string) (bytes []byte) {
	var buffer = bytes2.Buffer{}
	var last = &Duncode{}
	for i, x := range []rune(s) {
		var now = rune2Duncode(x)
		if last.compress(now) {
			continue
		}
		var b = last.toBytes()
		fmt.Printf(" encode %d %d \n", i, x)
		buffer.Write(b)
		last = now
	}
	var b = last.toBytes()
	buffer.Write(b)
	return buffer.Bytes()
}
func decode(bytes []byte) (s string) {
	var line = ""
	var buffer = bytes2.Buffer{}
	for i, x := range bytes {
		buffer.WriteByte(x)
		if x >= 0x80 {
			continue
		}
		fmt.Printf(" decode %d %d\n", i, x)
		var now = Duncode{}
		now.readBytes(buffer.Bytes())
		buffer.Reset()
		var chars = now.toChars()
		line += string(chars)
	}
	return line
}

func testLine(s string) {
	var bytes = encode(s)
	var t = decode(bytes)
	fmt.Println(t)
}

func main() {
	fmt.Println("Hello, 世界")
	var s = "Aÿぃ好乇αβЖΘक़ꌊ걹"

	//testDuncodeBytes(s)
	testDuncode(s)
	//testLine(s)
	fmt.Print("done")

}
