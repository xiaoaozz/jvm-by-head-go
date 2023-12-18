package comparisons

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

// FCMPG 比较long类型变量
type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

// FCMPL 比较long类型变量
type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}

// _fcmp 定义一个统一的函数，供两条指令使用
func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	// 当两个float变量中至少有一个是NaN时
	// 用fcmpg指令比较的结果是1
	// 用fcmpl指令比较的结果是-1
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(-1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		// 浮点数计算可能产生NaN值，所以比较两个浮点数时，第四种结果是无法比较
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}
