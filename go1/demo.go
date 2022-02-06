package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func EncodeFile(src, tgt string, debug bool) {
	log.Printf("EncodeFile src:%s --> tgt:%s started \n", src, tgt)
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
		if debug {
			var decoded = Decode(bytes)
			if line != decoded {
				var d = strings.Compare(line, decoded)
				log.Fatalf("n_line:%d %s != %s %d", n_line, line, decoded, d)
				os.Exit(n_line)
			}
		}
		_, err := w.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}
		n_line += 1
		n_src += len(line)
		n_tgt += len(bytes)
		if n_line%1000000 == 0 {
			log.Printf("line %d src %d tgt %d", n_line, len(line), len(bytes))
		}
	}
	w.Flush()
	log.Printf("line %d src %d tgt %d done \n", n_line, n_src, n_tgt)
	fi, err := os.Stat(src)
	var size0 = fi.Size()
	fi, err = os.Stat(tgt)
	var size1 = fi.Size()
	log.Printf(" src %d tgt %d done \n", size0, size1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main1() {
	var s = "Aÿぃ好乇αβζβЖѰѾѾक़ऄळ४ॐꌊ걹"
	// s = "꧋ꦲꦗꦤꦼꦒꦺꦴꦂꦫꦶꦮꦶꦠ꧀ꦰꦶꦁꦲꦶꦱꦶꦃꦲꦼꦤꦺꦴꦩ꧀"
	s = string([]rune(s)[:])
	fmt.Print(s)
	testDuncode(s)
	// testDuncodeCompress(s)
	testLine(s)

	var src = "C:/data/sentences.csv"
	var tgt = src + ".duncode1"
	EncodeFile(src, tgt, true)
	fmt.Print("done")

}

func main() {
	// var s = "Aÿぃ好乇αβζЖѰѾѾक़ऄळ४ॐꌊ걹"
	// testDuncode(s)
	// testDuncodeCompress(s)
	// testLine(s)
	// var bytes = Encode(s)
	// var t = Decode(bytes)
	// fmt.Printf("duncode :%s %d  --> %s %d\n", s, len(s), t, len(bytes))
	args := os.Args
	var src = args[1]
	var tgt = args[2]
	var debug = false
	if len(args) > 3 {
		debug = true
	}
	// var src = "C:/data/sentences.csv"
	// var tgt = src + ".duncode1"
	// log.Printf("done")
	EncodeFile(src, tgt, debug)
	// log.Printf("done")

}
