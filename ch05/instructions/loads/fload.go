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

type FLOAD struct{ base.Index8Instruction }

func (self *FLOAD) Execute(frame *rtda.Frame) {
	_fload(frame, self.Index)
}

type FLOAD_0 struct{ base.NoOperandsInstruction }

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	_fload(frame, 0)
}

type FLOAD_1 struct{ base.NoOperandsInstruction }

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	_fload(frame, 1)
}

type FLOAD_2 struct{ base.NoOperandsInstruction }

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	_fload(frame, 2)
}

type FLOAD_3 struct{ base.NoOperandsInstruction }

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	_fload(frame, 3)
}

func _fload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}
