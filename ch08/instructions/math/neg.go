package math

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
)

// DNEG double类型数据取反
type DNEG struct{ base.NoOperandsInstruction }

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopDouble()
	stack.PushDouble(-val)
}

// FNEG float类型数据取反
type FNEG struct{ base.NoOperandsInstruction }

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopFloat()
	stack.PushFloat(-val)
}

// INEG int类型数据取反
type INEG struct{ base.NoOperandsInstruction }

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushInt(-val)
}

// LNEG long类型数据取反
type LNEG struct{ base.NoOperandsInstruction }

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushLong(-val)
}
