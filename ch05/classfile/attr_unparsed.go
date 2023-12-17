package classfile

// UnparsedAttribute 未解析的属性，用于其他未定义的属性
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

// readInfo 读取属性内容
func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
