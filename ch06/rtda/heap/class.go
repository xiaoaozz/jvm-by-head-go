package heap

import (
	"jvm-by-head-go/ch06/classfile"
	"strings"
)

// Class 类结构信息
type Class struct {
	accessFlags       uint16        // 类的访问标志，共16位
	name              string        // 类名
	superClassName    string        // 超类名
	interfaceNames    []string      // 接口名
	constantPool      *ConstantPool // 运行时常量池指针
	fields            []*Field      // 字段表
	methods           []*Method     // 方法表
	loader            *ClassLoader  // 类加载指针
	superClass        *Class        // 超类指针
	interfaces        []*Class      // 接口指针
	instanceSlotCount uint          // 实例变量占据空间大小
	staticSlotCount   uint          // 类变量占据空间大小
	staticVars        Slots         // 静态变量
}

// newClass 将classFile转换成Class结构体
func newClass(cf *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = cf.AccessFlags()
	class.name = cf.ClassName()
	class.superClassName = cf.SuperClassName()
	class.interfaceNames = cf.InterfaceNames()
	class.constantPool = newConstantPool(class, cf.ConstantPool())
	class.fields = newFields(class, cf.Fields())
	class.methods = newMethods(class, cf.Methods())
	return class
}

func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags&ACC_PUBLIC
}
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags&ACC_FINAL
}
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags&ACC_SUPER
}
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags&ACC_INTERFACE
}
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags&ACC_SYNTHETIC
}
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags&ACC_ANNOTATION
}
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags&ACC_ENUM
}

// getters
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// isAccessibleTo jvm虚拟机规范5.4.4节给出类的访问权限控制规则
func (self *Class) isAccessibleTo(other *Class) bool {
	// 如果类D像访问类C，需要满足两个条件之一：
	// （1）C是public
	// （2）C和D在同一个运行时包内
	return self.IsPublic() ||
		self.getPackageName() == other.getPackageName()
}

// getPackageName 获取包名
func (self *Class) getPackageName() string {
	// 比如类名是 java/lang/Object，则包名就是java.lang
	if i := strings.LastIndex(self.name, "/"); i >= 0 {
		return self.name[:i]
	}
	// 类定义在默认包中，包名为空字符串
	return ""
}

func (self *Class) GetMainMethod() *Method {
	return self.getStaticMethod("main", "([Ljava/lang/String;)V")
}

func (self *Class) getStaticMethod(name, descriptor string) *Method {
	for _, method := range self.methods {
		if method.IsStatic() &&
			method.name == name &&
			method.descriptor == descriptor {

			return method
		}
	}
	return nil
}

// NewObject 初始化实例
func (self *Class) NewObject() *Object {
	return newObject(self)
}
