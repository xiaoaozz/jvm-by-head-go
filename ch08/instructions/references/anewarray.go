package references

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
	"jvm-by-head-go/ch08/rtda/heap"
)

// ANEW_ARRAY 创建引用类型数组
type ANEW_ARRAY struct{ base.Index16Instruction }

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()

	// if componentClass.InitializationNotStarted() {
	// 	thread := frame.Thread()
	// 	frame.SetNextPC(thread.PC()) // undo anewarray
	// 	thread.InitClass(componentClass)
	// 	return
	// }

	stack := frame.OperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrClass := componentClass.ArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}
