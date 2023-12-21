package math

import (
	"jvm-by-head-go/ch06/instructions/base"
	"jvm-by-head-go/ch06/rtda"
)

// IINC iinc指令给局部变量表中的int变量增加常量值
type IINC struct {
	Index uint
	Const int32
}

// FetchOperands 从字节码读取操作数
func (self *IINC) FetchOperands(reader *base.BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
	self.Const = int32(reader.ReadInt8())
}

// Execute iinc指令执行逻辑
func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	// 从常量池中读取常量
	val := localVars.GetInt(self.Index)
	// 然后再加上常量
	val += self.Const
	// 将结果写入局部变量表
	localVars.SetInt(self.Index, val)
}
