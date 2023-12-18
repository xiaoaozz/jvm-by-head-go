package constants

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

// BIPUSH bipush指令从操作数中获取一个byte类型整数，扩展成int类型，然后推入栈顶
type BIPUSH struct {
	val int8
}

// FetchOperands bipush指令从字节码中提取操作数的具体实现
func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

// Execute bipush指令执行逻辑具体实现
func (self *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}

// SIPUSH sipush指令从操作数中获取一个short类型整数，扩展成int类型，然后推入栈顶
type SIPUSH struct {
	val int16
}

// FetchOperands sipush指令从字节码中提取操作数的具体实现
func (self *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt16()
}

// Execute sipush指令执行逻辑的具体实现
func (self *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
