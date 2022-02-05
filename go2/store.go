package main

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
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

var ZoneName2Id = map[string]int{
	"ascii": 0,
	"双节":    1,
	"8位字":   2,
	"7位字":   3,
	"孤字":    4,
}

func loadBlocks(path string) {
	dat, err := os.ReadFile(path)
	check(err)
	var text = string(dat)
	var doc = strings.Split(text, "\n")
	var Zone2Id = -1 // zone2
	var Zone3Id = -1 //zone3
	for BlockId, t := range doc {
		var row = strings.Split(t, "\t")
		if len(row) < 6 {
			continue
		}
		for i := 0; i < 6; i += 1 {
			row[i] = strings.TrimSpace(row[i])
		}
		// var ZoneId = Int(row[6])
		var ZoneId = ZoneName2Id[row[5]]
		// if ZoneId != zoneid {
		// 	panic("ZoneId!=zoneid")
		// }
		var block = Block{
			BlockId: BlockId,
			Began:   Int(row[0]),
			End:     Int(row[1]),
			// Size:     Int(row[2]),
			English:  row[3],
			Chinese:  row[4],
			ZoneName: row[5],
			ZoneId:   ZoneId,
			Zone2Id:  -1,
			Zone3Id:  -1,
		}
		block.Size = block.End - block.Began + 1
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

func loadShuangJie(path string) {
	dat, err := os.ReadFile(path)
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

func GetCurrPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

func init() {
	fmt.Println("\nHello, store")
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	// fmt.Println(GetCurrPath())
	// var block_path = "Blocks.txt"
	var block_path = path.Join(exPath, "Blocks.txt")
	loadBlocks(block_path)
	// fmt.Println(len(blocks))
	// var shuangjie_path = "ShuangJie.txt"
	var shuangjie_path = path.Join(exPath, "ShuangJie.txt")
	loadShuangJie(shuangjie_path)
	// fmt.Println(len(ShuangJieIndex))
	fmt.Println("store loaded\n")

}
