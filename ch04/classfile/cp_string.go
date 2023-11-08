package classfile

// ConstantStringInfo string类型常量
type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

// readInfo 读取常量池索引
func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// String 按索引从常量池中查找字符串
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
