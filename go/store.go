package duncode

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Int(num string) (n int) {
	a, err := strconv.ParseInt(num, 16, 64)
	// check(err)
	if err != nil {
		panic(err)
	}
	return int(a)
}

type Block struct {
	BlockId  int
	Began    int
	End      int
	Size     int
	English  string
	Chinese  string
	ZoneName string
	ZoneId   int
	Zone2Id  int
	Zone3Id  int
}

var blocks = make([]Block, 0)

func loadBlocks() {
	dat, err := os.ReadFile("Blocks.txt")
	check(err)
	var text = string(dat)
	var doc = strings.Split(text, "\n")
	var Zone2Id = -1 // zone2
	var Zone3Id = -1 //zone3
	for BlockId, t := range doc {
		var row = strings.Split(t, "\t")
		if len(row) != 7 {
			continue
		}
		for i := 0; i < 7; i += 1 {
			row[i] = strings.TrimSpace(row[i])
		}
		var ZoneId = Int(row[6])
		var block = Block{Began: Int(row[0]),
			BlockId:  BlockId,
			End:      Int(row[1]),
			Size:     Int(row[2]),
			English:  row[3],
			Chinese:  row[4],
			ZoneName: row[5],
			ZoneId:   ZoneId,
			Zone2Id:  -1,
			Zone3Id:  -1,
		}
		if ZoneId == 2 {
			Zone2Id += 1
			block.Zone2Id = Zone2Id
		} else if ZoneId == 3 {
			Zone3Id += 1
			block.Zone3Id = Zone3Id
		}
		blocks = append(blocks, block)
	}
	fmt.Printf("loadBlocks %d done \n", len(blocks))
}

var ShuangJies = make([]rune, 0)
var ShuangJieIndex = make(map[rune]int)

func loadShuangJie() {
	dat, err := os.ReadFile("ShuangJie.txt")
	check(err)
	var text = string(dat)
	var doc = strings.Split(text, "\n")
	for i, s := range doc {
		if len(s) == 0 {
			break
		}
		var t = []rune(s)[0]
		ShuangJies = append(ShuangJies, t)
		ShuangJieIndex[t] = i
	}
	fmt.Printf("loadShuangJie %d done  \n", len(ShuangJies))
}

func init() {
	fmt.Println("\nHello, store")
	loadBlocks()
	// fmt.Println(len(blocks))
	loadShuangJie()
	// fmt.Println(len(ShuangJieIndex))
	fmt.Println("store loaded\n")

}
