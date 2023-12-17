package base

import (
	"jvm-by-head-go/ch05/rtda"
)

/**
下面是一些指令接口与抽象指令的实现
*/

// Instruction 指令接口
type Instruction interface {
	FetchOperands(reader *BytecodeReader) // FetchOperands()从字节码中提取操作数
	Execute(frame *rtda.Frame)            // Execute()方法执行指令逻辑
}

// NoOperandsInstruction 表示没有操作数的指令，不定义任何字段
type NoOperandsInstruction struct {
}

// FetchOperands NoOperandsInstruction指令的具体实现
func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {
	// 这里什么也不做
}

// BranchInstruction 表示跳转指令
type BranchInstruction struct {
	Offset int // 存放跳转偏移量
}

// FetchOperands BranchInstruction指令的具体实现
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	// 从字节码读取一个uint16整数，转成int后赋值给Offset字段
	self.Offset = int(reader.ReadInt16())
}

// Index8Instruction 存储和加载指令需要根据索引存取局部变量表，索引由单字节操作数给出，这类指令被抽象为Index8Instruction
type Index8Instruction struct {
	Index uint // 局部变量表索引
}

// FetchOperands Index8Instruction指令的具体实现
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	self.Index = uint(reader.ReadUint8())
}

// Index16Instruction 一些指令需要访问运行时常量池，常量池索引由两字节操作数给出，这类指令抽象成Index16Instruction
type Index16Instruction struct {
	Index uint // 常量池索引
}

// FetchOperands Index16Instruction指令的具体实现
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	// 从字节码中读取一个uint16整，转成uint后复制给Index字段
	self.Index = uint(reader.ReadUint16())
}
