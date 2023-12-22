package heap

import (
	"jvm-by-head-go/ch07/classfile"
)

type MethodRef struct {
	MemberRef
	method *Method
}

func newMethodRef(cp *ConstantPool, refInfo *classfile.ConstantMethodrefInfo) *MethodRef {
	ref := &MethodRef{}
	ref.cp = cp
	ref.copyMemberRefInfo(&refInfo.ConstantMemberrefInfo)
	return ref
}

func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		// 如果还没有解析过符号引用，调用resolveMethodRef方法进行解析
		self.resolveMethodRef()
	}
	// 否则直接返回方法指针
	return self.method
}

// resolveMethodRef jvm8规范 5.4.3.3 解析方法引用
func (self *MethodRef) resolveMethodRef() {
	// 如果类D想通过方法符号引用访问类C的某个方法，先要解析符号引用得到类C
	d := self.cp.class
	c := self.ResolvedClass()
	// 如果C是接口，则抛出IncompatibleClassChangeError异常
	if c.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 否则根据方法名和描述符查找方法
	method := lookupMethod(c, self.name, self.descriptor)
	// 如果找不到对应的方法，则抛出NoSuchMethodError异常
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	// 否则检查类D是否有权限访问该方法
	if !method.isAccessibleTo(d) {
		// 如果没有，则抛出IllegalAccessError异常
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

// lookupMethod 查找方法
func lookupMethod(class *Class, name, descriptor string) *Method {
	// 先从class的继承层次中查找，如果找不到，就去class的接口中去找
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}
