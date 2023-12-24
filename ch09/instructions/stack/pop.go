package stack

import (
	"jvm-by-head-go/ch09/instructions/base"
	"jvm-by-head-go/ch09/rtda"
)

// POP pop指令把栈顶变量弹出
type POP struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
            |
            V
[...][c][b]
*/

// Execute pop指令执行逻辑的具体实现
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

// POP2 pop2指令，double和long变量在操作数栈中占两个位置，专门为其设置
type POP2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
         |  |
         V  V
[...][c]
*/

// Execute pop2指令的执行逻辑，弹出两个slot
func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
