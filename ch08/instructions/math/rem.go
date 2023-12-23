package math

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
	"math"
)

/**
求余指令
*/

type DREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	// Go语言没有给浮点数类型定义求余操作符，所以需要使用math包
	// 因为浮点类型因为有 Infinity 无穷大，所以即使是除零，也不会导致ArithmeticException异常
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	// Go语言没有给浮点数类型定义求余操作符，所以需要使用math包
	// 因为浮点类型因为有 Infinity 无穷大，所以即使是除零，也不会导致ArithmeticException异常
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

type IREM struct {
	base.NoOperandsInstruction
}

// Execute irem指令执行逻辑的具体实现
func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 先从操作数栈中弹出两个int变量
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	// 求余，将结果推入操作数栈
	result := v1 % v2
	stack.PushInt(result)
}

type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}
