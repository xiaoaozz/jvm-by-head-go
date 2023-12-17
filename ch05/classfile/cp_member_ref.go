package classfile

type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberrefInfo) className() string {
	return self.cp.getClassName(self.classIndex)
}
func (self *ConstantMemberrefInfo) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

// ConstantFieldrefInfo 表示字段符号引用
type ConstantFieldrefInfo struct {
	ConstantMemberrefInfo
}

// ConstantMethodrefInfo 表示普通方法（非接口）符号引用
type ConstantMethodrefInfo struct {
	ConstantMemberrefInfo
}

// ConstantInterfaceMethodrefInfo 表示接口方法符号引用
type ConstantInterfaceMethodrefInfo struct {
	ConstantMemberrefInfo
}
