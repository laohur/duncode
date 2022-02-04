package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func EncodeFile(src, tgt string) {
	file, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, err := os.Create(tgt)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	w := bufio.NewWriter(f)

	var n_line = 0
	var n_src = 0
	var n_tgt = 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		var line = scanner.Text()
		var bytes = Encode(line)
		_, err := w.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}
		n_line += 1
		n_src += len(line)
		n_tgt += len(bytes)
		if n_line%1000000 == 0 {
			log.Printf("line %d src %d tgt %d", n_line, n_src, n_tgt)
		}
	}
	w.Flush()
	log.Printf("line %d src %d tgt %d done \n", n_line, n_src, n_tgt)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func demo() {
	// var s = "Aÿぃ好乇αβЖѰक़ꌊ걹"
	// testDuncode(s)
	// testDuncodeCompress(s)
	// testLine(s)
	// var bytes = Encode(s)
	// var t = Decode(bytes)
	// fmt.Printf("duncode :%s %d  --> %s %d\n", s, len(s), t, len(bytes))
	var src = "C:/data/sentences.csv"
	var tgt = src + ".duncode1"
	EncodeFile(src, tgt)
	fmt.Print("done")

}


func main() {
	// var s = "Aÿぃ好乇αβЖѰक़ꌊ걹"
	// testDuncode(s)
	// testDuncodeCompress(s)
	// testLine(s)
	// var bytes = Encode(s)
	// var t = Decode(bytes)
	// fmt.Printf("duncode :%s %d  --> %s %d\n", s, len(s), t, len(bytes))
	args := os.Args
	var src=args[1]
	var tgt=args[2]
	// var src = "C:/data/sentences.csv"
	// var tgt = src + ".duncode1"
	EncodeFile(src, tgt)
	fmt.Print("done")

}
