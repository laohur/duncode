package main

import (
	"bufio"
	bytes2 "bytes"
	"log"
	"os"
	"strings"
)

type Duncode struct {
	CodePoint int
	ZoneId    int
	BlockId   int
	MotherId  int
	Index     int
	Symbols   []int
}

func Rune2Duncode(char rune) (d *Duncode) {
	var duncode = &Duncode{}
	var point = int(char)
	duncode.CodePoint = point
	if point < 128 { // ascii
		duncode.ZoneId = 0
		duncode.BlockId = 0
		duncode.Index = point
		return duncode
	}
	var idx, ok = ShuangJieIndex[char]
	if ok && point > 0x07FF { // 双节
		duncode.ZoneId = 1
		duncode.BlockId = -1
		duncode.Index = idx
		return duncode
	}

	for _, block := range blocks {
		if block.Began <= point && point <= block.End && block.ZoneId >= 2 {
			duncode.BlockId = block.BlockId
			duncode.ZoneId = block.ZoneId
			if duncode.ZoneId == 4 { //孤字
				duncode.Index = point
			} else { //8位字、7位字
				duncode.MotherId = block.MotherId
				if block.Mother == block.Chinese {
					duncode.Index = point - block.Began
				} else {
					duncode.Index = point - block.Began + block.Offset //始终母块为基
				}
			}
			return duncode
		}
	}
	if ok { // 双节
		duncode.ZoneId = 1
		duncode.BlockId = -1
		duncode.Index = idx
		return duncode
	}

	// log.Fatalf("error 无效字符" + string(char))
	duncode.Index = point
	duncode.ZoneId = 4
	duncode.BlockId = -1
	return duncode
}

func (a *Duncode) compress(b *Duncode) (r bool) {
	if a.ZoneId == 2 || a.ZoneId == 3 {
		if a.ZoneId == b.ZoneId && a.MotherId == b.MotherId && len(a.Symbols) < 3 && a.Index != 0 {
			a.BlockId = a.MotherId
			if len(a.Symbols) == 0 {
				a.Symbols = []int{a.Index, b.Index}
			} else {
				a.Symbols = append(a.Symbols, b.Index)
			}
			return true
		}
	}
	return false
}

func (d *Duncode) findSon(symbol int) *Block {
	var nowBlock = blocks[d.BlockId]
	if nowBlock.Offset <= symbol && symbol < nowBlock.Offset+nowBlock.Size {
		return nowBlock
	}
	var sons = Children[nowBlock.Mother]
	for _, name := range *sons {
		var now = blocks[BlockName2Id[name]]
		if now.Offset <= symbol && symbol < now.Offset+now.Size {
			return now
		}
	}
	panic("找不到母块")
}

func (a *Duncode) decompress() (array []*Duncode) {
	if len(a.Symbols) > 0 {
		array = make([]*Duncode, len(a.Symbols))
		for i, symbol := range a.Symbols {
			var block = a.findSon(symbol)
			var x = Duncode{
				CodePoint: a.CodePoint,
				ZoneId:    a.ZoneId,
				BlockId:   block.BlockId,
				MotherId:  a.MotherId,
				Index:     symbol,
			}
			// array=append(array,&x)
			array[i] = &x
		}
	} else {
		if len(blocks[a.MotherId].Child) > 0 {
			var block = a.findSon(a.Index)
			a.BlockId = block.BlockId
		}
		array = []*Duncode{a}
	}
	return array
}

func (d *Duncode) toBytes() (bytes []byte) {
	switch d.ZoneId {
	case 0: //ascii
		return []byte{byte(d.Index)}
	case 1: //双节
		var idx uint16 = uint16(d.Index)
		var a byte = byte(0x80) + byte(idx>>7)
		var b byte = byte(idx & 0x7f)
		return []byte{a, b}
	case 2: //8位字
		var x = byte(0)
		var y = byte(0)
		var z = byte(0)
		switch len(d.Symbols) {
		case 0:
			var index, ok = ShuangJieIndex[rune(d.CodePoint)]
			if ok { // 双节
				var idx uint16 = uint16(index)
				var a byte = byte(0x80) + byte(idx>>7)
				var b byte = byte(idx & 0x7f)
				return []byte{a, b}
			}
			// 孤字
			var a = byte(0x80) + byte(d.CodePoint>>14)&byte(0x7f)
			var b = byte(0x80) + byte(d.CodePoint>>7)&byte(0x7f)
			var c = byte(d.CodePoint) & byte(0x7f)
			return []byte{a, b, c}

		case 2:
			y = byte(d.Symbols[0])
			z = byte(d.Symbols[1])
			break
		case 3:
			x = byte(d.Symbols[0])
			y = byte(d.Symbols[1])
			z = byte(d.Symbols[2])
			break
		}
		// |       2 | 8位字 | 111nnxxx  | 1xxxxxyy  | 1yyyyyyz | 0zzzzzzz | x,y,z   | Greek…      |           1.33 |
		var Zone2Id = blocks[d.MotherId].Zone2Id
		var a byte = byte(0b111)<<5 | byte(Zone2Id)<<3 | x>>5
		var b byte = byte(0b1)<<7 | x<<2 | y>>6
		var c byte = byte(0b1)<<7 | y<<1 | z>>7
		var D byte = byte(0x7f) & z
		return []byte{a, b, c, D}
	case 3: //7位字
		var x = byte(0)
		var y = byte(0)
		var z = byte(0)
		switch len(d.Symbols) {
		case 0:
			var index, ok = ShuangJieIndex[rune(d.CodePoint)]
			if ok { // 双节
				var idx uint16 = uint16(index)
				var a byte = byte(0x80) + byte(idx>>7)
				var b byte = byte(idx & 0x7f)
				return []byte{a, b}
			}
			// 孤字
			var a = byte(0x80) + byte(d.CodePoint>>14)&byte(0x7f)
			var b = byte(0x80) + byte(d.CodePoint>>7)&byte(0x7f)
			var c = byte(d.CodePoint) & byte(0x7f)
			return []byte{a, b, c}
		case 2:
			y = byte(d.Symbols[0])
			z = byte(d.Symbols[1])
			break
		case 3:
			x = byte(d.Symbols[0])
			y = byte(d.Symbols[1])
			z = byte(d.Symbols[2])
			break
		}
		// |       3 | 7位字 | 1nnnnnnn  | 1xxxxxxx  | 1yyyyyyy | 0zzzzzzz | x,y,z   | Devanagari… |           1.33 |
		var Zone3Id = blocks[d.MotherId].Zone3Id
		var a byte = byte(0b1)<<7 | byte(Zone3Id)
		var b byte = byte(0b1)<<7 | x
		var c byte = byte(0b1)<<7 | y
		var D byte = z
		return []byte{a, b, c, D}

	case 4: //孤字
		//|       4 | 孤字   |           | 1xxxxxxx  | 1xxxxxxx | 0xxxxxxx  | x       | rare        |              3 |
		var a = byte(0x80) + byte(d.Index>>14)&byte(0x7f)
		var b = byte(0x80) + byte(d.Index>>7)&byte(0x7f)
		var c = byte(d.Index) & byte(0x7f)
		return []byte{a, b, c}
	}
	panic("toBytes not valid Duncode Zone id")
}

func (d *Duncode) readBytes(bytes []byte) {
	switch len(bytes) {
	case 1: //ascii
		var b = bytes[0]
		d.Index = int(b)
		d.ZoneId = 0
		d.BlockId = 0
		return
	case 2: //双节
		var a = byte(0x7f) & bytes[0]
		var b = byte(0x7f) & bytes[1]
		d.Index = int(a)<<7 + int(b)
		d.ZoneId = 1
		d.BlockId = -1
		return
	case 3: //孤字
		//|       4 | 孤字   |           | 1xxxxxxx  | 1xxxxxxx | 0xxxxxxx  | x       | rare        |              3 |
		var a = byte(0x7f) & bytes[0]
		var b = byte(0x7f) & bytes[1]
		var c = byte(0x7f) & bytes[2]
		//var d = byte(0x7f) & bytes[3]
		var idx = 0
		idx += int(a) << 14
		idx += int(b) << 7
		idx += int(c)
		d.Index = idx
		d.ZoneId = 4
		d.BlockId = -1
		return
	case 4:
		var a = byte(0x7f) & bytes[0]
		if a>>5 == 0b11 { //8位字
			// |       2 | 8位字 | 111nnxxx  | 1xxxxxyy  | 1yyyyyyz | 0zzzzzzz | x,y,z   | Greek…      |           1.33 |
			var nn = int((a >> 3) & 0b11)
			d.ZoneId = 2
			for _, block := range blocks {
				if block.Zone2Id == nn && block.Mother == block.Chinese {
					d.MotherId = block.MotherId
					d.BlockId = block.BlockId
					break
				}
			}
			var b = (byte(0x7f) & bytes[1]) >> 2
			var x = int(a<<5 | b)
			var y = int(bytes[1]<<6 | (bytes[2]<<1)>>2)
			var z = int(bytes[2]<<7 | bytes[3])
			if x == 0 && y == 0 {
				d.Index = z
			} else if x == 0 {
				d.Symbols = []int{y, z}
			} else {
				d.Symbols = []int{x, y, z}
			}
			return
		} else { //7位字
			// |       3 | 7位字 | 1nnnnnnn  | 1xxxxxxx  | 1yyyyyyy | 0zzzzzzz | x,y,z   | Devanagari… |           1.33 |
			d.ZoneId = 3
			var nn = int(a & 0x7f)
			for _, block := range blocks {
				if block.Zone3Id == nn && block.Mother == block.Chinese {
					d.BlockId = block.BlockId
					d.MotherId = block.MotherId
					break
				}
			}
			var x = int(0x7f & bytes[1])
			var y = int(0x7f & bytes[2])
			var z = int(0x7f & bytes[3])
			if x == 0 && y == 0 {
				d.Index = z
			} else if x == 0 {
				d.Symbols = []int{y, z}
			} else {
				d.Symbols = []int{x, y, z}
			}
			return
		}
	}
	panic("readBytes not valid Duncode Zone id")
}

func (d *Duncode) toChar() (char rune) {
	switch d.ZoneId {
	case 0:
		d.CodePoint = d.Index
		char = rune(d.CodePoint)
		return char
	case 1:
		d.CodePoint = int(ShuangJies[d.Index])
		char = ShuangJies[d.Index]
		return char
	case 2:
		var offset = d.Index - blocks[d.BlockId].Offset
		d.CodePoint = blocks[d.BlockId].Began + offset
		char = rune(d.CodePoint)
		return char
	case 3:
		var offset = d.Index - blocks[d.BlockId].Offset
		d.CodePoint = blocks[d.BlockId].Began + offset
		char = rune(d.CodePoint)
		return char
	case 4:
		d.CodePoint = d.Index
		char = rune(d.CodePoint)
		return char
	}
	panic("toChar not valid Duncode Zone id")
}

func Encode(s string) (bytes []byte) {
	if len(s) == 0 {
		return
	}
	var buffer = bytes2.Buffer{}
	var last = &Duncode{}
	for i, x := range []rune(s) {
		var now = Rune2Duncode(x)
		if i == 0 {
			last = now
			continue
		} else if last.compress(now) {
			continue
		}
		var b = last.toBytes()
		buffer.Write(b)
		last = now
	}
	var b = last.toBytes()
	buffer.Write(b)
	return buffer.Bytes()
}

func Decode(bytes []byte) (s string) {
	var line = ""
	var buffer = bytes2.Buffer{}
	for _, x := range bytes {
		buffer.WriteByte(x)
		if x >= 0x80 {
			continue
		}
		var now = Duncode{}
		now.readBytes(buffer.Bytes())
		var decompressed = now.decompress()
		for _, d := range decompressed {
			var char = d.toChar()
			line += string(char)
		}
		buffer.Reset()
	}
	return line
}

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
	var n_char = 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	// scanner.Split(bufio.ScanWords) // use scanwords
	const maxCapacity = 1024 * 1024 * 1024 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	// scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		var line = scanner.Text() + "\n"
		n_char += len([]rune(line))
		var bytes = Encode(line)
		_, err := w.Write(bytes)
		if err != nil {
			log.Fatal(err)
		}
		if debug {
			w.Flush()
			var decoded = Decode(bytes)
			if line != decoded {
				var d = strings.Compare(line, decoded)
				log.Fatalf("n_line:%d %s != %s %d", n_line, line, decoded, d)
				os.Exit(n_line)
			}
		}
		n_line += 1
		n_src += len(line)
		n_tgt += len(bytes)
		if n_line%1000000 == 0 {
			log.Printf("line %d n_char%d src %d tgt %d", n_line, n_char, len(line), len(bytes))
		}
	}
	w.Flush()
	log.Printf("line %d n_char %d src %d tgt %d done \n", n_line, n_char, n_src, n_tgt)
	fi, err := os.Stat(src)
	var size0 = fi.Size()
	fi, err = os.Stat(tgt)
	var size1 = fi.Size()
	log.Printf(" src %d tgt %d done \n", size0, size1)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func DecodeFile(src, tgt string, debug bool) {
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
	var n_char = 0
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	// scanner.Split(bufio.ScanWords) // use scanwords
	const maxCapacity = 1024 * 1024 * 1024 // your required line length
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)
	// scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// var line = scanner.Text()
		var bytes = scanner.Bytes()
		bytes = append(bytes, '\n')
		// n_char += len([]rune(line))
		// var bytes = Encode(line)
		var line = Decode(bytes)
		if debug {
			w.Flush()
			// var decoded = Encode(bytes)
			// if line != decoded {
			// 	var d = strings.Compare(line, decoded)
			// 	log.Fatalf("n_line:%d %s != %s %d", n_line, line, decoded, d)
			// 	os.Exit(n_line)
			// }
		}
		var encoded = []byte(line)
		_, err := w.Write(encoded)
		if err != nil {
			log.Fatal(err)
		}
		n_line += 1
		n_src += len(line)
		n_tgt += len(bytes)
		if n_line%1000000 == 0 {
			log.Printf("line %d n_char%d src %d tgt %d", n_line, n_char, len(line), len(bytes))
		}
	}
	w.Flush()
	log.Printf("line %d n_char %d src %d tgt %d done \n", n_line, n_char, n_src, n_tgt)
	if debug {
		fi, _ := os.Stat(src)
		var size0 = fi.Size()
		fi, err = os.Stat(tgt)
		var size1 = fi.Size()
		log.Printf(" src %d tgt %d done \n", size0, size1)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
