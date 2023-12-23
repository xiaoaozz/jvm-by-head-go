package stack

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
)

// DUP dup指令，复制栈顶一个或者两个数值并将复制值或双份的复制值重新压入栈顶
type DUP struct {
	base.NoOperandsInstruction
}

// Execute dup指令执行逻辑的具体实现，复制栈顶的单个变量
func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
}

type DUP_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
          __/
         |
         V
[...][c][a][b][a]
*/

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	s1 := stack.PopSlot()
	s2 := stack.PopSlot()
	stack.PushSlot(s1)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}

type DUP_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _____/
      |
      V
[...][a][c][b][a]
*/

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	s1 := stack.PopSlot()
	s2 := stack.PopSlot()
	s3 := stack.PopSlot()
	stack.PushSlot(s1)
	stack.PushSlot(s3)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}

type DUP2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]____
          \____   |
               |  |
               V  V
[...][c][b][a][b][a]
*/

func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	s1 := stack.PopSlot()
	s2 := stack.PopSlot()
	stack.PushSlot(s2)
	stack.PushSlot(s1)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}

type DUP2_X1 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][c][b][a]
       _/ __/
      |  |
      V  V
[...][b][a][c][b][a]
*/

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	s1 := stack.PopSlot()
	s2 := stack.PopSlot()
	s3 := stack.PopSlot()
	stack.PushSlot(s2)
	stack.PushSlot(s1)
	stack.PushSlot(s3)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}

type DUP2_X2 struct {
	base.NoOperandsInstruction
}

/*
bottom -> top
[...][d][c][b][a]
       ____/ __/
      |   __/
      V  V
[...][b][a][d][c][b][a]
*/

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	s1 := stack.PopSlot()
	s2 := stack.PopSlot()
	s3 := stack.PopSlot()
	s4 := stack.PopSlot()
	stack.PushSlot(s2)
	stack.PushSlot(s1)
	stack.PushSlot(s4)
	stack.PushSlot(s3)
	stack.PushSlot(s2)
	stack.PushSlot(s1)
}
