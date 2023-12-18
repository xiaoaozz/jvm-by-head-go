package loads

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

/**
加载指令：
	加载指令从局部变量表获取变量，然后推入操作数栈顶。
	（1）aload系列指令操作引用类型变量
	（2）dload系列指令操作double类型变量
	（3）fload系统指令操作float类型变量
	（4）iload系统指令操作int类型变量
	（5）lload系统指令操作long类型变量
	（6）xaload系统指令操作数组
*/

type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	_dload(frame, self.Index)
}

type DLOAD_0 struct{ base.NoOperandsInstruction }

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	_dload(frame, 0)
}

type DLOAD_1 struct{ base.NoOperandsInstruction }

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	_dload(frame, 1)
}

type DLOAD_2 struct{ base.NoOperandsInstruction }

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	_dload(frame, 2)
}

type DLOAD_3 struct{ base.NoOperandsInstruction }

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	_dload(frame, 3)
}

func _dload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}
