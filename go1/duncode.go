package main

import bytes2 "bytes"

type Duncode struct {
	CodePoint int
	ZoneId    int
	BlockId   int
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
	if ok { // 双节
		duncode.ZoneId = 1
		duncode.BlockId = -1
		duncode.Index = idx
		return duncode
	} else {
		for _, block := range blocks {
			if block.Began <= point && point <= block.End {
				duncode.BlockId = block.BlockId
				duncode.ZoneId = block.ZoneId
				if duncode.ZoneId == 4 { //孤字
					duncode.Index = point
				} else { //8位字、7位字
					duncode.Index = point - block.Began
				}
				return duncode
			}
		}
	}
	return duncode
}

func (a *Duncode) compress(b *Duncode) (r bool) {
	if a.BlockId == b.BlockId && a.ZoneId == 2 && b.ZoneId == 2 && len(a.Symbols) < 3 && a.Index != 0 {
		if len(a.Symbols) == 0 {
			a.Symbols = []int{a.Index, b.Index}
		} else {
			a.Symbols = append(a.Symbols, b.Index)
		}
		return true
	} else if a.BlockId == b.BlockId && a.ZoneId == 3 && b.ZoneId == 3 && len(a.Symbols) < 3 && a.Index != 0 {
		if len(a.Symbols) == 0 {
			a.Symbols = []int{a.Index, b.Index}
		} else {
			a.Symbols = append(a.Symbols, b.Index)
		}
		return true
	}
	return false
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
			// z = byte(d.Index)
			// break
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
		// |       2 | 8位字  | 10nnnxxx  | 1xxxxxyy  | 1yyyyyyz | 0zzzzzzz  | x,y,z   | Greek…      |           1.33 |
		var Zone2Id = blocks[d.BlockId].Zone2Id
		var a byte = byte(0b10)<<6 | byte(Zone2Id)<<3 | x>>5
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
			// z = byte(d.Index)
			// break
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
		//|       3 | 7位字  | 110nnnnn  | 1xxxxxxx  | 1yyyyyyy | 0zzzzzzz  | x,y,z   | Devanagari… |           1.33 |
		var Zone3Id = blocks[d.BlockId].Zone3Id
		var a byte = byte(0b110)<<5 | byte(Zone3Id)
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
		if a>>6 == 0b00 { //8位字
			// |       2 | 8位字  | 10nnnxxx  | 1xxxxxyy  | 1yyyyyyz | 0zzzzzzz  | x,y,z   | Greek…      |           1.33 |
			var nn = int(((a << 2) >> 5))
			d.ZoneId = 2
			for _, block := range blocks {
				if block.Zone2Id == nn {
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
			//|       3 | 7位字  | 110nnnnn  | 1xxxxxxx  | 1yyyyyyy | 0zzzzzzz  | x,y,z   | Devanagari… |           1.33 |
			var nn = int(a & 0b11111)
			for _, block := range blocks {
				if block.Zone3Id == nn {
					d.ZoneId = 2
					d.BlockId = block.BlockId
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

func (d *Duncode) toChars() (chars []rune) {
	switch d.ZoneId {
	case 0:
		d.CodePoint = d.Index
		chars = []rune{rune(d.CodePoint)}
		return chars
	case 1:
		d.CodePoint = int(ShuangJies[d.Index])
		chars = []rune{ShuangJies[d.Index]}
		return chars
	case 2:
		if len(d.Symbols) == 0 {
			d.CodePoint = blocks[d.BlockId].Began + d.Index
			chars = []rune{rune(d.CodePoint)}
			return chars
		} else {
			var x = rune(blocks[d.BlockId].Began + d.Symbols[0])
			var y = rune(blocks[d.BlockId].Began + d.Symbols[1])
			chars = []rune{x, y}
			if len(d.Symbols) == 3 {
				var z = rune(blocks[d.BlockId].Began + d.Symbols[2])
				chars = append(chars, z)
			}
			return chars
		}
	case 3:
		if len(d.Symbols) == 0 {
			d.CodePoint = blocks[d.BlockId].Began + d.Index
			chars = []rune{rune(d.CodePoint)}
			return chars
		} else {
			var x = blocks[d.BlockId].Began + d.Symbols[0]
			var y = blocks[d.BlockId].Began + d.Symbols[1]
			chars = []rune{rune(x), rune(y)}
			if len(d.Symbols) == 3 {
				var z = rune(blocks[d.BlockId].Began + d.Symbols[2])
				chars = append(chars, rune(z))
			}
			return chars
		}
	case 4:
		d.CodePoint = d.Index
		chars = []rune{rune(d.CodePoint)}
		return chars
	}
	panic("toChars not valid Duncode Zone id")
}

func Encode(s string) (bytes []byte) {
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
		var chars = now.toChars()
		buffer.Reset()
		line += string(chars)
	}
	return line
}
