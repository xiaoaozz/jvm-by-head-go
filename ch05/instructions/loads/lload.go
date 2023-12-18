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

type LLOAD struct{ base.Index8Instruction }

func (self *LLOAD) Execute(frame *rtda.Frame) {
	_lload(frame, self.Index)
}

type LLOAD_0 struct{ base.NoOperandsInstruction }

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	_lload(frame, 0)
}

type LLOAD_1 struct{ base.NoOperandsInstruction }

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	_lload(frame, 1)
}

type LLOAD_2 struct{ base.NoOperandsInstruction }

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	_lload(frame, 2)
}

type LLOAD_3 struct{ base.NoOperandsInstruction }

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	_lload(frame, 3)
}

func _lload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}
