package heap

// SymRef 代表符号引用，用于复合使用代码，供其他四个引用继承
type SymRef struct {
	cp        *ConstantPool // 符号引用所在的运行时常量池指针
	className string        // 类的完全限定类名
	class     *Class        // 缓存解析后的类结构体指针
}

// ResolvedClass 类解析
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		// 调用resolveClassRef()方法解析
		self.resolveClassRef()
	}
	// 如果类符号引用已经解析，直接返回类指针
	return self.class
}

// resolveClassRef 解析类符号
func (self *SymRef) resolveClassRef() {
	// jvm虚拟机第5.4.3.1节
	// 如果类D通过符号引用N引用类C的话，则需要先解析N，先用D的类加载器加载C
	// 然后检查D是否有权限访问C，如果没有，则抛出IllegalAccessError异常
	d := self.cp.class
	c := d.loader.LoadClass(self.className)
	if !c.isAccessibleTo(d) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = c
}
