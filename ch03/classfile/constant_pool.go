package classfile

// ConstantPool 常量池
type ConstantPool []ConstantInfo

// readConstantPool 读取常量池
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		// 注意，索引从1开始，有效的常量池索引是从 1-n
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
		// CONSTANT_Long_info和CONSTANT_DOUBLE_iofo各占两个位置
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}
	}
	return cp
}

// getConstantInfo 按索引查找常量
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

// getNameAndType 从常量池查找字段或者方法的名字和描述符
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

// getClassName 从常量池查找类名
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

// getUtf8 从常量池查找UTF-8字符串
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}

// ConstantInfo 表示常量信息
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

// newConstantInfo 创建具体的常量
func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_String:
		return &ConstantStringInfo{}
	case CONSTANT_Class:
		return &ConstantClassInfo{}
	case CONSTANT_Fieldref:
		return &ConstantFieldrefInfo{
			ConstantMemberrefInfo{cp: cp},
		}
	case CONSTANT_Methodref:
		return &ConstantMethodrefInfo{
			ConstantMemberrefInfo{cp: cp},
		}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodrefInfo{
			ConstantMemberrefInfo{cp: cp},
		}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		panic("java.lang.ClassFormatError: constant pool tag!")
	}

}

// readConstantInfo 读取tag值，然后创建具体的常量，并读取常量信息
func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	// 读取出tag的值
	tag := reader.readUint8()
	// 创建具体的常量
	c := newConstantInfo(tag, cp)
	// 读取常量信息
	c.readInfo(reader)
	return c
}
