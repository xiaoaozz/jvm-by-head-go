package references

import (
	"jvm-by-head-go/ch07/instructions/base"
	"jvm-by-head-go/ch07/rtda"
)

// INVOKE_SPECIAL 对超类、私有和实例初始化方法调用的特殊处理
type INVOKE_SPECIAL struct{ base.Index16Instruction }

// Execute hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
