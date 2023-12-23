package heap

import (
	"fmt"
	"jvm-by-head-go/ch09/classfile"
	"jvm-by-head-go/ch09/classpath"
)

/*
class names:
    - primitive types: boolean, byte, int ...
    - primitive arrays: [Z, [B, [I ...
    - non-array classes: java/lang/Object ...
    - array classes: [Ljava/lang/Object; ...
*/

// ClassLoader 类加载器
type ClassLoader struct {
	cp          *classpath.Classpath // Classpath指针
	verboseFlag bool
	classMap    map[string]*Class // 记录已经加载的类数据。key是类的完全限定名（可以当作方法区的具体实现）
}

// NewClassLoader 创建类加载器实例
func NewClassLoader(cp *classpath.Classpath, verboseFlag bool) *ClassLoader {
	return &ClassLoader{
		cp:          cp,
		verboseFlag: verboseFlag,
		classMap:    make(map[string]*Class),
	}
}

// LoadClass 把类数据加载到方法区
func (self *ClassLoader) LoadClass(name string) *Class {
	// 查询classMap。看类是否已经被加载
	if class, ok := self.classMap[name]; ok {
		// 如果是，直接返回类数据
		return class
	}
	if name[0] == '[' {
		return self.loadArrayClass(name)
	}
	// 否则调用loadNonArrayClass加载类
	return self.loadNonArrayClass(name)
}

// loadNonArrayClass 加载非数组的类（数组类和普通类加载有很大的区别）
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	// 1. 查找class文件并把数据读取到内存
	data, entry := self.readClass(name)
	// 2. 解析class文件，生成虚拟机可以使用的类数据，并放入方法区
	class := self.defineClass(data)
	// 3. 链接
	link(class)
	if self.verboseFlag {
		fmt.Printf("[Loaded %s from %s]\n", name, entry)
	}
	return class
}

// loadArrayClass 加载数组类
func (self *ClassLoader) loadArrayClass(name string) *Class {
	class := &Class{
		accessFlags: ACC_PUBLIC, // todo
		name:        name,
		loader:      self,
		initStarted: true,
		superClass:  self.LoadClass("java/lang/Object"),
		interfaces: []*Class{
			self.LoadClass("java/lang/Cloneable"),
			self.LoadClass("java/io/Serializable"),
		},
	}
	self.classMap[name] = class
	return class
}

// readClass 读取类
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	// 调用ClassPath的ReadClass()方法
	data, entry, err := self.cp.ReadClass(name)
	// 进行错误处理
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// defineClass 解析class文件
func (self *ClassLoader) defineClass(data []byte) *Class {
	// 将class文件数据转换成Class结构体
	class := parseClass(data)
	class.loader = self
	// 根据jvm规范的5.3.5，调用以下两个函数解析类符号引用
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

// parseClass 解析Class文件
func parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		//panic("java.lang.ClassFormatError")
		panic(err)
	}
	return newClass(cf)
}

// resolveSuperClass 递归调用LoadClass方法加载class的超类
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// resolveInterfaces 递归调用LoadClass()方法加载类的每一个直接接口
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

// link 链接
func link(class *Class) {
	// 验证阶段
	verify(class)
	// 准备阶段
	prepare(class)
}

func verify(class *Class) {
	// todo jvm规范4.10节
}

// prepare 准备阶段，主要是给类变量分配空间并给予初始值
func prepare(class *Class) {
	// jvm规范5.4.2节
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// calcInstanceFieldSlotIds 计算实例字段的个数，同时给它们编号。
func calcInstanceFieldSlotIds(class *Class) {
	// 为了解决字段对应Slots中位置的问题。
	// 注意：
	// （1）静态实例和实例字段要分开编号
	// （2）对于实例字段，一定要从继承关系的最顶端，即java.lang.Object开始编号
	// （3）编号时要考虑long和double类型
	// 否则会混乱，导致错误的结果
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

// calcStaticFieldSlotIds 计算静态字段的个数，同时给它们编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	for _, field := range class.fields {
		if field.IsStatic() {
			field.slotId = slotId
			slotId++
			if field.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// allocAndInitStaticVars 给类变量分配空间，然后给它们赋予初始值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// initStaticFinalVar 从常量池中加载常量值，然后给静态变量赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstValueIndex()
	slotId := field.SlotId()

	if cpIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)
		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)
		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)
		case "Ljava/lang/String;":
			goStr := cp.GetConstant(cpIndex).(string)
			jStr := JString(class.Loader(), goStr)
			vars.SetRef(slotId, jStr)
		}
	}
}
