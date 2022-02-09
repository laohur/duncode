package main

import (
	"fmt"
	"log"
	"os"
)

func main1() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	var s = "𰀦Aÿぃ好乇αβζβЖѰѾѾक़ऄळ४ॐꌊ걹"
	s = "\u3000♪リンゴ可愛いや可愛いやリンゴ。半世紀も前に流行した「リンゴの歌」がぴったりするかもしれない。米アップルコンピュータ社のパソコン「マック（マッキントッシュ）」を、こよなく愛する人たちのことだ。「アップル信者」なんて言い方まである。"
	s = string([]rune(s)[:2])
	fmt.Print(s)
	testDuncode(s)
	// testDuncodeCompress(s)
	testLine(s)

	var src = "C:/data/sentences.csv"
	src = "C:/data/wiki-1m/kv.txt"
	var tgt = src + ".duncode1"
	var decoded = src + ".decoded"
	EncodeFile(src, tgt, true)
	DecodeFile(tgt, decoded, true)
	fmt.Print(" demo done")
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
