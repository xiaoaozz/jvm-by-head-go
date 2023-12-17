package classfile

// ExceptionAttribute 记录方法抛出的异常表
type ExceptionsAttribute struct {
	exceptionsIndexTable []uint16
}

// readInfo 读取属性索引
func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionsIndexTable = reader.readUint16s()
}

// ExceptionIndexTable 获取异常处理表
func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionsIndexTable
}
