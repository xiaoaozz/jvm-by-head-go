package references

import (
	"jvm-by-head-go/ch09/instructions/base"
	"jvm-by-head-go/ch09/rtda"
	"jvm-by-head-go/ch09/rtda/heap"
)

// NEW new指令，创建类实例
type NEW struct {
	base.Index16Instruction // uint16索引，来自字节码
}

func (self *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 从当前类的运行时常量池中找到一个类符号引用
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	// 解析该类符号引用，拿到类数据
	class := classRef.ResolvedClass()
	// 如果类还没有初始化，则需要先初始化
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}
	// 接口和抽象类都不能实例化，所以如果解析后的类是接口或者抽象类
	// 按照Java虚拟机规定，需要抛出InstantiationError异常
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	// 创建对象
	ref := class.NewObject()
	// 将对象引用推入栈顶
	frame.OperandStack().PushRef(ref)
}
