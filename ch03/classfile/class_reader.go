package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

// readUint8 读取u1类型数据
func (self *ClassReader) readUint8() uint8 {
	// 从字节数组中读取一个字节（8位），返回改字节的无符号整数值
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// readUint16 读取u2类型数据
func (self *ClassReader) readUint16() uint16 {
	// 从字节数组中读取两个字节（16位），以大端字节序解析为无符号整数值，然后返回该值
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// readUint32 读取u4类型数据
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

// readUit64 读取uint64类型数据，Java虚拟机规范并没有定义u8类型
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

// readUint16s 读取uint16的集合数据
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

// readBytes 读取指定数量的字节数据
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
