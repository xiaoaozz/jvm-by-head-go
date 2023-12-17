package constants

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

// ACONST_NULL aconst_null指令
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

// Execute aconst_null指令把null引用推入操作数栈顶
func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// DCONST_0 dconst_0指令
type DCONST_0 struct {
	base.NoOperandsInstruction
}

// Execute dconst_0指令把double类型0推入操作数栈顶
func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_0 struct {
	base.NoOperandsInstruction
}
type FCONST_1 struct {
	base.NoOperandsInstruction
}
type FCONST_2 struct {
	base.NoOperandsInstruction
}

// ICONST_M1 iconst_m1指令
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

// Execute iconst_m1指令把int类型的-1推入操作数栈顶
func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}
type ICONST_1 struct {
	base.NoOperandsInstruction
}
type ICONST_2 struct {
	base.NoOperandsInstruction
}
type ICONST_3 struct {
	base.NoOperandsInstruction
}
type ICONST_4 struct {
	base.NoOperandsInstruction
}
type ICONST_5 struct {
	base.NoOperandsInstruction
}
type LCONST_0 struct {
	base.NoOperandsInstruction
}
type LCONST_1 struct {
	base.NoOperandsInstruction
}
