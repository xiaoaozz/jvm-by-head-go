package math

import (
	"jvm-by-head-go/ch08/instructions/base"
	"jvm-by-head-go/ch08/rtda"
)

// ISHL int左位移
type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	// 从操作数栈中弹出两个int变量v2和v1
	v2 := stack.PopInt() //  要位移多少比特
	v1 := stack.PopInt() // 要进行位移操作的变量
	// int变量只有32位，所以只取v2的前5个比特就足够表示位移位数
	s := uint32(v2) & 0x1f
	// 位移之后，将结果推入操作数栈，另外，Go语言位移操作符右侧必须是无符号整数，所以需要对v2进行类型转换
	res := v1 << s
	stack.PushInt(res)
}

// ISHR int算术右位移
type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1f
	result := v1 >> s
	stack.PushInt(result)
}

// IUSHR int逻辑右位移
type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	// Go语言中并没有Java语言的>>>运算符，为了能达到无符号右移的目的
	// 需要先把v1转成无符号整数，位移操作之后，再转回有符号整数
	s := uint32(v2) & 0x1f
	res := int32(uint32(v1) >> s)
	stack.PushInt(res)
}

// LSHL long左位移
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	res := v1 << s
	stack.PushLong(res)
}

// LSHR long算术右位移
type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	// long变量有64位，所以取v2的前6个比特
	s := uint32(v2) & 0x3f
	res := v1 >> s
	stack.PushLong(res)
}

// LUSHR long逻辑右位移
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3f
	res := int64(uint64(v1) >> s)
	stack.PushLong(res)
}
