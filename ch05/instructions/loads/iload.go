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

// ILOAD iload指令的具体实现 （iload系列指令操作int变量）
type ILOAD struct {
	base.Index8Instruction
}

// Execute iload指令的执行逻辑具体实现，iload指令的索引来自操作数
func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

// Execute iload_0指令的执行逻辑具体实现
func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

// Execute iload_1指令的执行逻辑具体实现，类似指令的索引隐含在操作码中
func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

// Execute iload_2指令的执行逻辑具体实现
func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

// Execute iload_3指令的执行逻辑具体实现
func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}

// _iload 定义一个统一的函数供load类型指令使用
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
