package classfile

// SourceFileAttribute 用于指出源文件名
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

// readInfo 读取属性索引，并赋值给sourceFileIndex
func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

// FileName 获取class源文件名
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
