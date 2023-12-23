package constants

import (
	"jvm-by-head-go/ch09/instructions/base"
	"jvm-by-head-go/ch09/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
