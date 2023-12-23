package stores

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
)

/**
存储指令：
	存储指令把变量从操作数栈顶弹出，然后存入局部变量表。
*/

// LSTORE 存储long类型数据的指令
type LSTORE struct {
	base.Index8Instruction
}

// Execute lstore指令执行逻辑的具体实现，该指令的索引来自操作数
func (self *LSTORE) Execute(frame *rtda.Frame) {
	_lstore(frame, uint(self.Index))
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

// Execute lstore_0指令执行逻辑的具体实现
func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_lstore(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

// Execute lstore_1指令执行逻辑的具体实现
func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_lstore(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

// Execute lstore_2指令执行逻辑的具体实现
func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_lstore(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

// Execute lstore_3指令执行逻辑的具体实现
func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_lstore(frame, 3)
}

// _lstore 定义一个统一的函数供store指令使用
func _lstore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
