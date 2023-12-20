package stack

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

// SWAP swap指令，将栈顶的两个值互换
type SWAP struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
          \/
          /\
         V  V
[...][c][a][b]
*/

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	s1 := stack.PopSlot()
	s2 := stack.PopSlot()
	stack.PushSlot(s1)
	stack.PushSlot(s2)
}
