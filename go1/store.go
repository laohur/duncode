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

func Int(num string, base int) (n int) {
	a, err := strconv.ParseInt(num, base, 64)
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
	Mother   string
	MotherId int
	Offset   int
	Child    []string
	ZoneId   int
	Zone2Id  int
	Zone3Id  int
}

var blocks = make([]*Block, 0)

var ZoneName2Id = map[string]int{
	"ascii": 0,
	"双节":    1,
	"八位字":   2,
	"七位字":   3,
	"孤字":    4,
}
var BlockName2Id = make(map[string]int)
var Children = make(map[string]*[]string)

func Split(s string,sep rune)([]string){
	splitFn := func(c rune) bool {
		return c == sep
    }
    // fmt.Printf("Fields are: %q\n", s.FieldsFunc("a,,b,c", splitFn))
	var t= strings.FieldsFunc(s, splitFn)
	return t
}

// fields = [began, end, size, en, ch, zone, mother, offset, child]
func loadBlocks(path string) {
	dat, err := os.ReadFile(path)
	check(err)
	var text = string(dat)
	var doc = strings.Split(text, "\n")
	var Zone2Id = -1 // zone2
	var Zone3Id = -1 //zone3
	for BlockId, t := range doc {
		var row = strings.Split(t, "\t")
		if len(row) < 8 {
			continue
		}
		for i := 0; i < len(row); i += 1 {
			row[i] = strings.TrimSpace(row[i])
		}
		var ZoneId = ZoneName2Id[row[5]]
		
		var child=Split(row[8], ';')
		var block = Block{
			BlockId:  BlockId,
			Began:    Int(row[0], 16),
			End:      Int(row[1], 16),
			Size:     -1,
			English:  row[3],
			Chinese:  row[4],
			ZoneName: row[5],
			ZoneId:   ZoneId,
			Mother:   row[6],
			MotherId: -1,
			Offset:   Int(row[7], 10),
			Child:    child,
			Zone2Id:  -1,
			Zone3Id:  -1,
		}
		BlockName2Id[block.Chinese] = BlockId
		block.Size = block.End - block.Began + 1
		if ZoneId == 2 {
			block.MotherId=BlockName2Id[block.Mother]
			if block.Mother==block.Chinese  {
				Zone2Id += 1
				block.Zone2Id = Zone2Id
			} 
		} else if ZoneId == 3 {
			block.MotherId=BlockName2Id[block.Mother]
			if block.Mother==block.Chinese  {
				Zone3Id += 1
				block.Zone3Id = Zone3Id
			}
		}
		if len(block.Child)>0{
			Children[block.Chinese]=&block.Child
		}
		blocks = append(blocks,& block)
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

func getPath(name string) string {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	// fmt.Println(GetCurrPath())
	// var block_path = "Blocks.txt"
	var path = path.Join(exPath, name)
	return path
}

func init() {
	fmt.Println("\nHello, store")
	// fmt.Println(GetCurrPath())
	var block_path = "Blocks.txt"
	// block_path = getPath(block_path)
	loadBlocks(block_path)
	// fmt.Println(len(blocks))
	var shuangjie_path = "ShuangJie.txt"
	// shuangjie_path = getPath(shuangjie_path)
	loadShuangJie(shuangjie_path)
	// fmt.Println(len(ShuangJieIndex))
	fmt.Println("store loaded\n")

}
