package main

import (
	"fmt"
	"log"
	"os"
)

func main1() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var s = "𰀦Aÿぃ好乇αβζβЖѰѾѾक़ऄळ४ॐꌊ걹"
	s="Mark Abhisit Vejjajiva (toktok Tae: อภิสิทธิ์ เวชชาชีวะ ; bon long 3 Ogis 1964 long Newcastle upon Tyne long Yunaeted Kingdom‎), em i wan man blong politik blong Taelan. Em i bin praem minista blong Taelan, stat long 2008, kasem 2011.\n"
	s = string([]rune(s)[:])
	fmt.Print(s)
	testDuncode(s)
	// testDuncodeCompress(s)
	testLine(s)

	var src = "C:/data/sentences.csv"
	src = "C:/data/wiki-1m/bi.txt"
	var tgt = src + ".duncode2"
	var decoded = src + ".decoded"
	EncodeFile(src, tgt, true)
	DecodeFile(tgt,decoded, true)
	fmt.Print("done")
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	args := os.Args
	var method = args[1]
	var src = args[2]
	var tgt = args[3]
	var debug = false
	if len(args) > 4 {
		debug = true
	}
	if method == "encode" {
		EncodeFile(src, tgt, debug)
	} else if method == "decode" {
		DecodeFile(src, tgt, debug)
	}
}

