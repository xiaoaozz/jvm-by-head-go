package heap

import "jvm-by-head-go/ch06/classfile"

// ClassRef 类符号引用
type ClassRef struct {
	SymRef
}

// newClassRef 根据class文件中存储的类常量创建ClassRef实例
func newClassRef(cp *ConstantPool, classInfo *classfile.ConstantClassInfo) *ClassRef {
	ref := &ClassRef{}
	ref.cp = cp
	ref.className = classInfo.Name()
	return ref
}
