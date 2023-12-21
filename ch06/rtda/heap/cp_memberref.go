package heap

import "jvm-by-head-go/ch06/classfile"

type MemberRef struct {
	SymRef
	name       string
	descriptor string
}

// copyMemberRefInfo 从class文件内存储的字段或者方法常量中提取数据
func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.ConstantMemberrefInfo) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}
func (self *MemberRef) Descriptor() string {
	return self.descriptor
}
