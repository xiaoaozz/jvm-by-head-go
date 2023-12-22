package heap

// jvm8规范 6.5.instanceof
// jvm8规范 6.5.checkcast
func (self *Class) isAssignableFrom(other *Class) bool {
	// 有三种情况下，S类型的引用值可以赋值给T类型
	// （1）S和T是同一类型
	// （2）T是类且S是T的子类
	// （3）T是接口且S实现了T接口
	s, t := other, self

	if s == t {
		return true
	}

	if !t.IsInterface() {
		return s.IsSubClassOf(t)
	} else {
		return s.IsImplements(t)
	}
}

// IsSubClassOf 判断self是否继承了other
func (self *Class) IsSubClassOf(other *Class) bool {
	// 实际上就是判断T是否是S的（直接或间接）超类
	for c := self.superClass; c != nil; c = c.superClass {
		if c == other {
			return true
		}
	}
	return false
}

// IsImplements 判断self是否实现了iface接口
func (self *Class) IsImplements(iface *Class) bool {
	// 实际上就是看S或者S的（直接或间接）超类是否实现了某个接口T‘，T'要么是T，要么是T的子接口
	for c := self; c != nil; c = c.superClass {
		for _, i := range c.interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

// self extends iface
func (self *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (self *Class) IsSuperClassOf(other *Class) bool {
	return other.IsSubClassOf(self)
}
