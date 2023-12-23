package constants

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
)

type NOP struct{ base.NoOperandsInstruction }

func (self *NOP) Execute(frame *rtda.Frame) {
	// really do nothing
}
