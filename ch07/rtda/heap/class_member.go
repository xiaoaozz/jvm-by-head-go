package heap

import "jvm-by-head-go/ch07/classfile"

// ClassMember 类成员信息，如字段、方法
type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class // Class结构体指针，可以通过字段或者方法访问它所属的类
}

// copyMemberInfo 从class文件中复制数据
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.AccessFlags()
	self.name = memberInfo.Name()
	self.descriptor = memberInfo.Descriptor()
}

func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags&ACC_PRIVATE
}
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags&ACC_PROTECTED
}
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags&ACC_STATIC
}
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}

/**
get方法
*/

func (self *ClassMember) Name() string {
	return self.name
}
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}
func (self *ClassMember) Class() *Class {
	return self.class
}

// isAccessibleTo 字段访问规则 jvm规范5.4.4节
func (self *ClassMember) isAccessibleTo(d *Class) bool {
	if self.IsPublic() {
		// 如果字段是public，则任何类都可以访问
		return true
	}
	c := self.class
	if self.IsProtected() {
		// 如果字段是protected，则只有子类和同一个包下的类可以访问
		return d == c || d.isSubClassOf(c) ||
			c.getPackageName() == d.getPackageName()
	}
	if !self.IsPrivate() {
		return c.getPackageName() == d.getPackageName()
	}
	// 如果字段有默认访问权限（非public、非protected、非private），则只有同一个包下的类可以访问
	return d == c
}
