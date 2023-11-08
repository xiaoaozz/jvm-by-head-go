package classfile

// ConstantValueAttribute 表示常量表达式的值
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

// readInfo 读取属性索引
func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

// ConstantValueIndex 获取索引索引
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
