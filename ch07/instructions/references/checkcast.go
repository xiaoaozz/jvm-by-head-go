package references

import (
	"jvm-by-head-go/ch07/instructions/base"
	"jvm-by-head-go/ch07/rtda"
	"jvm-by-head-go/ch07/rtda/heap"
)

// CHECK_CAST
type CHECK_CAST struct{ base.Index16Instruction }

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	// 先从操作数栈中弹出对象引用，再推回去，不会改变操作数栈的状态
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	// 如果引用是null，则指令执行结束
	if ref == nil {
		return
	}
	// 否则，解析类符号引用，判断对象是否是类的实例
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	// 如果是，则指令执行结束
	// 如果不是，则抛出ClassCastException异常
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}

}
