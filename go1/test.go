package main

import (
	bytes2 "bytes"
	"fmt"
	"log"
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
		var a = Rune2Duncode(x)
		var b = a.toBytes()
		var now = &Duncode{}
		now.readBytes(b)
		var c = now.toChars()
		if x != c[0] {
			log.Fatalf("i:%d %c != %c", i, x, c)
		}
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
		var now = Rune2Duncode(x)
		if i == 0 {
			last = now
			continue
		} else if last.compress(now) {
			continue
		}
		duncodes = append(duncodes, last)
		var b = last.toBytes()
		var decoded = &Duncode{}
		decoded.readBytes(b)
		var c = decoded.toChars()
		fmt.Printf(" encode %d %c -->%c  bytes %d\n", i, x, c, b)
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

func string2bytes(s string) (bytes []byte) {
	var buffer = bytes2.Buffer{}
	// var duncodes = []*Duncode{}
	var last = &Duncode{}
	for i, x := range []rune(s) {
		if i == 0 {
			fmt.Printf(" encode %d ...\n", i)
		}
		var now = Rune2Duncode(x)
		if i == 0 {
			last = now
			continue
		} else if last.compress(now) {
			continue
		}
		// duncodes = append(duncodes, last)
		var b = last.toBytes()
		// var decoded = &Duncode{}
		// decoded.readBytes(b)
		// var c = decoded.toChars()
		// fmt.Printf(" encode %d %c -->%c \n", i, x, c)
		buffer.Write(b)
		last = now
	}
	// duncodes = append(duncodes, last)
	var b = last.toBytes()
	// var now = &Duncode{}
	// now.readBytes(b)
	// var c = now.toChars()
	// fmt.Printf(" encode -->%c \n", c)
	buffer.Write(b)
	return buffer.Bytes()
}

func bytes2string(bytes []byte) (s string) {
	var line = ""
	var buffer = bytes2.Buffer{}
	// var charArray = []rune{}
	for i, x := range bytes {
		if i == 0 {
			fmt.Printf(" decode %d ...\n", i)
		}
		buffer.WriteByte(x)
		if x >= 0x80 {
			continue
		}
		// fmt.Printf(" decode %d %d\n", i, x)
		var now = Duncode{}
		now.readBytes(buffer.Bytes())
		var chars = now.toChars()
		buffer.Reset()
		line += string(chars)
		// charArray = append(charArray, chars)
		// for _,c:=range chars{
		// charArray=append(charArray, c)
		// }
	}
	// var t= string(charArray)
	return line
}

func testLine(s string) {
	var bytes = string2bytes(s)
	var t = bytes2string(bytes)
	var l1 = []rune(s)
	var l2 = []rune(t)
	for i, x := range l1 {
		if x != l2[i] {
			panic(i)
		}
	}
	if len(l1) < len(l2) {
		panic(len(l1))
	}

	fmt.Printf("duncode :%s %d  --> %s %d\n", s, len(s), t, len(bytes))
}

func test() {
	var s = "Aÿぃ好乇αβζЖѰѾѾक़ऄळ४ॐꌊ걹"
	fmt.Println(s)
	// testDuncode(s)
	// testDuncodeCompress(s)
	testLine(s)
	fmt.Print("done")

}
