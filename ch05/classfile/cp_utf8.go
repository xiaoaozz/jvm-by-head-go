package classfile

import (
	"fmt"
	"unicode/utf16"
)

// ConstantUtf8Info MUTF-8编码字符串
type ConstantUtf8Info struct {
	str string
}

// readInfo 将数字转化成MUTF-8编码的字符串
func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// decodeMUTF8 将字符数组转化为MUTF8字符串，根据java.io.DataInputStream.readUTF(DataInput)方法改写
func decodeMUTF8(bytearr []byte) string {
	// 获取字节数组的长度
	utflen := len(bytearr)
	// 创建一个数组，用于存储字符
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	// 第一阶段：处理单字节字符
	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}
	// 第二阶段：处理多字节字符
	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// 可能产生的字符数小于utflen，因此截取chararr数组以适应实际的字符数
	chararr = chararr[0:chararr_count]
	// 使用utf16.Decode函数将chararr转换为UTF-16编码的字符数组
	runes := utf16.Decode(chararr)
	// 将UTF-16编码字符数组转换为字符串
	return string(runes)
}
