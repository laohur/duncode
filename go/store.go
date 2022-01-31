package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Int(num string) (n int64) {
	n, err := strconv.ParseInt(num, 16, 64)
	// check(err)
	if err != nil {
		panic(err)
	}
	return n
}

type Block struct {
	Began    int64
	End      int64
	Size     int64
	English  string
	Chinese  string
	ZoneName string
	ZoneId   int64
}

var blocks = make([]Block, 0)

func loadBlocks() {
	dat, err := os.ReadFile("Blocks.txt")
	check(err)
	var text = string(dat)
	var doc = strings.Split(text, "\n")
	for _, t := range doc {
		var row = strings.Split(t, "\t")
		if len(row) != 7 {
			continue
		}
		for i := 0; i < 7; i += 1 {
			row[i] = strings.TrimSpace(row[i])
		}
		var block = Block{Began: Int(row[0]),
			End:      Int(row[1]),
			Size:     Int(row[2]),
			English:  row[3],
			Chinese:  row[4],
			ZoneName: row[5],
			ZoneId:   Int(row[6]),
		}
		blocks = append(blocks, block)
	}
	fmt.Printf("loadBlocks %d done \n", len(blocks))
}

var ShuangJies = make([]string, 0)
var ShuangJieIndex = make(map[string]int)

func loadShuangJie() {
	dat, err := os.ReadFile("ShuangJie.txt")
	check(err)
	var text = string(dat)
	var doc = strings.Split(text, "\n")
	for i, t := range doc {
		s := strings.TrimSpace(t)
		if len(s) == 0 {
			continue
		}
		ShuangJies = append(ShuangJies, s)
		ShuangJieIndex[s] = i
	}
	fmt.Printf("loadShuangJie %d done \n", len(ShuangJies))
}

func init() {
	fmt.Println("Hello, 世界")
	loadBlocks()
	fmt.Println(len(blocks))
	loadShuangJie()
	fmt.Println(len(ShuangJieIndex))

}
