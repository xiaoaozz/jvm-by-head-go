package classfile

import (
	"math"
)

// ConstantIntegerInfo 整数常量
type ConstantIntegerInfo struct {
	val int32
}

// readInfo 读取一个int类型的整数
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

// ConstantFloatInfo 单精度浮点数常量
type ConstantFloatInfo struct {
	val float32
}

// readInfo 读取一个int类型的整数，并转换成float类型
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}

// ConstantLongInfo 长整型常量
type ConstantLongInfo struct {
	val int64
}

// readInfo 读取一个long类型的整数
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

// ConstantDoubleInfo 双精度浮点数常量
type ConstantDoubleInfo struct {
	val float64
}

// readInfo 读取一个double类型的整数
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = math.Float64frombits(bytes)
}
