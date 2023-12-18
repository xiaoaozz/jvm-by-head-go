package comparisons

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

/**
ifcond指令：
	把操作数栈顶的int变量弹出，然后跟0进行比较，满足条件则跳转。
	假设从栈顶弹出的变量是x，则指令执行跳转操作的条件如下：
	（1）ifeq: x == 0
	（2）ifne: x != 0
	（3）iflt: x < 0
	（4）ifle: x <= 0
	（5）ifgt: x > 0
	（6）ifge: x >= 0
*/

// IFEQ ifeq指令
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame, self.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}
type IFLT struct {
	base.BranchInstruction
}
type IFLE struct {
	base.BranchInstruction
}
type IFGT struct {
	base.BranchInstruction
}
type IFGE struct {
	base.BranchInstruction
}
