package constants

import (
	"jvm-by-head-go/ch06/instructions/base"
	"jvm-by-head-go/ch06/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
