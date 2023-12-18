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

type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	_aload(frame, self.Index)
}

type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	_aload(frame, 0)
}

type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	_aload(frame, 1)
}

type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	_aload(frame, 2)
}

type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	_aload(frame, 3)
}

// _aload 定义一个函数，供aload指令使用
func _aload(frame *rtda.Frame, index uint) {
	ref := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(ref)
}
