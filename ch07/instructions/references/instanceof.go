package references

import (
	"jvm-by-head-go/ch07/instructions/base"
	"jvm-by-head-go/ch07/rtda"
	"jvm-by-head-go/ch07/rtda/heap"
)

// INSTANCE_OF 该指令判断对象是否是某个类的实例（或者对象的类是否实现了某个接口）
type INSTANCE_OF struct {
	// 该指令需要两个操作数
	// 第一个操作数是uint16索引，从方法的字节码中获取，通过这个索引可以从当前类的运行时常量池中找到第一个类符号引用
	// 第二个操作数是对象引用，从操作数栈中弹出
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	// 先弹出对象引用，如果是null，则把0推入操作数栈
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	// 如果对象引用不是null，则解析类符号引用，判断对象是否是类的实例，然后把判断结果推入操作数栈
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}
