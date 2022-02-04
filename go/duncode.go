package main

import bytes2 "bytes"

// ZoneIdMap = {
//     "ascii": 0,  # BlockId 0
//     "双节": 1,  # lanid 1
//     "8位字": 2,
//     "7位字": 3,
//     "独字": 4
// }

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
	if point < 128 {
		duncode.ZoneId = 0
		duncode.BlockId = 0
		duncode.Index = point
		return duncode
	}
	var idx, ok = ShuangJieIndex[char]
	if ok {
		duncode.ZoneId = 1
		duncode.BlockId = -1
		duncode.Index = idx
		return duncode
	} else {
		for i, block := range blocks {
			if block.Began <= point && point <= block.End {
				duncode.BlockId = i
				duncode.ZoneId = block.ZoneId
				if duncode.ZoneId == 4 {
					duncode.Index = point
				} else {
					duncode.Index = point - block.Began
				}
				return duncode
			}
		}
	}
	return duncode
}
func (a *Duncode) compress(b *Duncode) (r bool) {
	if a.ZoneId == 2 && b.ZoneId == 2 && len(a.Symbols) <= 1 && a.Index != 0 {
		if len(a.Symbols) == 0 {
			a.Symbols = []int{a.Index, b.Index}
			return true
		}
	} else if a.ZoneId == 3 && b.ZoneId == 3 && len(a.Symbols) <= 1 && a.Index != 0 {
		if len(a.Symbols) == 0 {
			a.Symbols = []int{a.Index, b.Index}
			return true
		}
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
		if len(d.Symbols) < 2 {
			var a byte = byte(0xf0) + (byte(blocks[d.BlockId].Zone2Id)&byte(0x3))<<2
			var b byte = byte(0x80) + byte(d.Index>>7)
			var c byte = byte(0x7f) & byte(d.Index)
			return []byte{a, b, c}
		}
		var x = byte(d.Symbols[0])
		var y = byte(d.Symbols[1])
		var a byte = byte(0xf0) + (byte(blocks[d.BlockId].Zone2Id)&byte(0x3))<<2 + x>>6
		var b byte = byte(0x80) | x<<1 + y>>7
		var c byte = byte(0x7f) & y
		return []byte{a, b, c}
	case 3: //7位字
		if len(d.Symbols) < 2 {
			var a = byte(0x80) + byte(blocks[d.BlockId].Zone3Id)
			var b = byte(0x80)
			var c = byte(d.Index)
			return []byte{a, b, c}
		}
		var x = byte(d.Symbols[0])
		var y = byte(d.Symbols[1])
		var a = byte(0x80) + byte(blocks[d.BlockId].Zone3Id)
		var b = byte(0x80) + x
		var c = y
		return []byte{a, b, c}

	case 4: //独字
		var a = byte(0x80) + byte(d.Index>>21)&byte(0x7f)
		var b = byte(0x80) + byte(d.Index>>14)&byte(0x7f)
		var c = byte(0x80) + byte(d.Index>>7)&byte(0x7f)
		var d = byte(d.Index) & byte(0x7f)
		return []byte{a, b, c, d}
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
	case 3:
		var a = byte(0x7f) & bytes[0]
		if a>>4 == 0b111 { //8位字
			var nn = int(((a << 4) >> 6))
			d.ZoneId = 2
			for _, block := range blocks {
				if block.Zone2Id == nn {
					d.BlockId = block.BlockId
					break
				}
			}
			var b = (byte(0x7f) & bytes[1]) >> 1
			var x = int(a<<6 + b)
			var y = int((0b1&bytes[1])<<7 + bytes[2])
			if x == 0 {
				d.Index = y
			} else {
				d.Symbols = []int{x, y}
			}
			return
		} else { //7位字
			var nn = int(a & (0x7f))
			for _, block := range blocks {
				if block.Zone3Id == nn {
					d.ZoneId = 2
					d.BlockId = block.BlockId
					break
				}
			}
			var x = int(0x7f & bytes[1])
			var y = int(0x7f & bytes[2])
			if x == 0 {
				d.Index = y
			} else {
				d.Symbols = []int{x, y}
			}
			return
		}
	case 4: //独字
		var a = byte(0x7f) & bytes[0]
		var b = byte(0x7f) & bytes[1]
		var c = byte(0x7f) & bytes[2]
		//var d = byte(0x7f) & bytes[3]
		var idx = 0
		idx += int(a) << 21
		idx += int(b) << 14
		idx += int(c) << 7
		idx += int(bytes[3])
		d.Index = idx
		d.ZoneId = 4
		d.BlockId = -1
		return
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
		if len(d.Symbols) < 2 {
			d.CodePoint = blocks[d.BlockId].Began + d.Index
			chars = []rune{rune(d.CodePoint)}
			return chars
		} else {
			var x = rune(blocks[d.BlockId].Began + d.Symbols[0])
			var y = rune(blocks[d.BlockId].Began + d.Symbols[1])
			chars = []rune{x, y}
			return chars
		}
	case 3:
		if len(d.Symbols) < 2 {
			d.CodePoint = blocks[d.BlockId].Began + d.Index
			chars = []rune{rune(d.CodePoint)}
			return chars
		} else {
			var x = blocks[d.BlockId].Began + d.Symbols[0]
			var y = blocks[d.BlockId].Began + d.Symbols[1]
			chars = []rune{rune(x), rune(y)}
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
