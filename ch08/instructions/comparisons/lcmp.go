package comparisons

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
)

// LCMP long类型数据比较
type LCMP struct{ base.NoOperandsInstruction }

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}
}
