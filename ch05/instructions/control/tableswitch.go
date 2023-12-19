package control

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

type TABLE_SWITCH struct {
	defaultOffset int32   // 默认情况下执行跳转所需的字节码偏移量
	low           int32   // case的取值范围，下界
	high          int32   // case的取值范围，上界
	jumpOffsets   []int32 // 索引表，里面存放high-low+1个int值，对应各种case情况下，执行跳转所需要的字节码偏移量
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	// tableswitch指令操作码后面有0-3字节的padding，保证defaultOffset在字节码中的地址是4的倍数
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetsCount := self.high - self.low + 1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	// 先从操作数栈中弹出一个int变量
	index := frame.OperandStack().PopInt()
	var offset int
	// 判断是否在low和high给定的范围之内
	if index >= self.low && index <= self.high {
		// 如果在，从jumpOffsets表中查出偏移量进行跳转
		offset = int(self.jumpOffsets[index-self.low])
	} else {
		// 如果不在，按照defaultOffset跳转
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}
