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
		if i == 10 {
			fmt.Println(i)
		}
		var a = rune2Duncode(x)
		var b = a.toBytes()
		var now = &Duncode{}
		now.readBytes(b)
		var c = now.toChars()
		fmt.Printf("testDuncode %d %c --> %d %c \n", i, x, len(b), c[0])
	}

}

func testDuncodeCompress(s string) (bytes []byte) {
	var buffer = bytes2.Buffer{}
	var duncodes = []*Duncode{}
	var last = &Duncode{}
	for i, x := range []rune(s) {
		if i == 6 {
			fmt.Println(i)
		}
		var now = rune2Duncode(x)
		if last.compress(now) {
			continue
		}
		duncodes = append(duncodes, last)
		var b = last.toBytes()
		var decoded = &Duncode{}
		decoded.readBytes(b)
		var c = decoded.toChars()
		fmt.Printf(" encode %d %c -->%c \n", i, x, c)
		buffer.Write(b)
		last = now
	}
	duncodes = append(duncodes, last)
	var b = last.toBytes()
	var now = &Duncode{}
	now.readBytes(b)
	var c = now.toChars()
	fmt.Printf(" encode -->%c \n", c)
	buffer.Write(b)
	return buffer.Bytes()
}

func encode(s string) (bytes []byte) {
	var buffer = bytes2.Buffer{}
	var last = &Duncode{}
	for i, x := range []rune(s) {
		if i == 8 {
			fmt.Println(i)
		}
		var now = rune2Duncode(x)
		if last.compress(now) {
			continue
		}
		var b = last.toBytes()
		fmt.Printf(" encode %d %c \n", i, x)
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
	var s = "Aÿぃ好乇αβЖѰक़ꌊ걹"
	fmt.Println(s)

	// testDuncode(s)
	testDuncodeCompress(s)
	// testLine(s)
	fmt.Print("done")

}
