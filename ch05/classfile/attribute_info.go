package classfile

// AttributeInfo 属性信息
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

// readAttributes 读取属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

// readAttribute 读取单个属性
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	// 先读取属性名索引，从常量池中找到属性名
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	// 读取属性长度，调用newAttributeInfo创建具体的属性实例
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

// newAttributeInfo 创建具体的属性实例，这里先解析8种属性
func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionsAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
