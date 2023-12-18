package comparisons

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

// LCMP long类型数据比较
type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 弹出栈顶的两个long变量
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	// 进行比较，将结果推入栈顶
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
