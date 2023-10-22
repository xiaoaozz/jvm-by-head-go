package classfile

// CodeAttribute Code属性结构
type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16                 // 操作数栈的最大深度
	maxLocals      uint16                 // 局部变量表大小
	code           []byte                 // 字节码
	exceptionTable []*ExceptionTableEntry // 异常处理表
	attributes     []AttributeInfo        // 属性表
}

// ExceptionTableEntry 异常处理表
type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

// readInfo 读取属性索引
func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

// readExceptionTable 读取异常处理表
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
