package math

import (
	"jvm-by-head-go/ch07/instructions/base"
	"jvm-by-head-go/ch07/rtda"
)

// DADD double类型数据相加
type DADD struct{ base.NoOperandsInstruction }

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v1 := stack.PopDouble()
	v2 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

// FADD float类型数据相加
type FADD struct{ base.NoOperandsInstruction }

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

// IADD int类型数据相加
type IADD struct{ base.NoOperandsInstruction }

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

// LADD long类型数据相加
type LADD struct{ base.NoOperandsInstruction }

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}
