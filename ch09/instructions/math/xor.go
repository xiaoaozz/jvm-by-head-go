package math

import (
	"jvm-by-head-go/ch09/instructions/base"
	"jvm-by-head-go/ch09/rtda"
)

// IXOR int类型异或
type IXOR struct{ base.NoOperandsInstruction }

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopInt()
	v2 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

// LXOR long类型异或
type LXOR struct{ base.NoOperandsInstruction }

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopLong()
	v2 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}
