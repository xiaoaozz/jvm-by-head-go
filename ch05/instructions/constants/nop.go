package constants

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

// NOP nop指令
type NOP struct {
	base.NoOperandsInstruction
}

// Execute nop指令执行的具体实现
func (self *NOP) Execute(frame *rtda.Frame) {
	// 什么也不用做
}
