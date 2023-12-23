package heap

import "jvm-by-head-go/ch08/classfile"

type FieldRef struct {
	MemberRef
	field *Field // 缓存解析后的字段指针
}

// newFieldRef 创建FieldRef实例
func newFieldRef(cp *ConstantPool, refInfo *classfile.ConstantFieldrefInfo) *FieldRef {
	ref := &FieldRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// resolveFieldRef 字段符号引用解析 jvm规范5.4.3.2节
func (self *FieldRef) resolveFieldRef() {
	// 如果类D想通过字段符号引用访问C的某个字段
	d := self.cp.class
	// 首先要解析符号引用得到类C
	c := self.ResolvedClass()
	// 然后根据字段名和描述符查找字段
	field := lookupField(c, self.name, self.descriptor)
	// 如果字段查找失败，则虚拟机抛出NoSuchFieldError
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	// 如果D没有足够的权限访问该字段，则虚拟机抛出IllegalAccessError
	if !field.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

// lookupField 字段查找步骤
func lookupField(c *Class, name, descriptor string) *Field {
	// 首先在C的字段中查找
	for _, field := range c.fields {
		if field.name == name && field.descriptor == descriptor {
			return field
		}
	}
	// 如果找不到，则在C的直接接口递归应用这个查找过程
	for _, iface := range c.interfaces {
		if field := lookupField(iface, name, descriptor); field != nil {
			return field
		}
	}
	// 如果仍然找不到，则在C的超类中递归应用这个查找过程
	if c.superClass != nil {
		return lookupField(c.superClass, name, descriptor)
	}
	// 如果仍然找不到，则查找失败
	return nil
}
