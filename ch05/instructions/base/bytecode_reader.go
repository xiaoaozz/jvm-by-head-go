package base

/**
用于辅助指令解码的结构体
*/

// BytecodeReader 字节码读取器
type BytecodeReader struct {
	code []byte // 存放字节码
	pc   int    // 记录读取到了哪个字节
}

// Reset 为了避免每次解析指令都创建一个新的Reader，所以定义一个Reset方法
func (self *BytecodeReader) Reset(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

func (self *BytecodeReader) PC() int {
	return self.pc
}

// ReadUint8 读取一个字节的数据
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

// ReadInt8 将读取到的一个字节的数据转成int8返回
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

// ReadUint16 连续读取两个字节的数据
func (self *BytecodeReader) ReadUint16() uint16 {
	b1 := uint16(self.ReadUint8())
	b2 := uint16(self.ReadUint8())
	return (b1 << 8) | b2
}

// ReadInt16 将读取到的两个字节数据转为int16返回
func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

// ReadInt32 连续读取四个字节的数据
func (self *BytecodeReader) ReadInt32() int32 {
	b1 := int32(self.ReadUint8())
	b2 := int32(self.ReadUint8())
	b3 := int32(self.ReadUint8())
	b4 := int32(self.ReadUint8())
	return (b1 << 24) | (b2 << 16) | (b3 << 8) | b4

}

// SkipPadding 保证defaultOffset在字节码中的地址是4的倍数
func (self *BytecodeReader) SkipPadding() {
	for self.pc%4 != 0 {
		self.ReadUint8()
	}
}

func (self *BytecodeReader) ReadInt32s(n int32) []int32 {
	ints := make([]int32, n)
	for i := range ints {
		ints[i] = self.ReadInt32()
	}
	return ints
}
