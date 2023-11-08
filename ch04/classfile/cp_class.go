package classfile

// ConstantClassInfo 表示类或者接口的符号引用
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

// readInfo 读取常量池索引
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

// Name 按索引从常量池中查找Name
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
