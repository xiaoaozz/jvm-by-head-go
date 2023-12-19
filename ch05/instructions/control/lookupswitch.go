package control

import (
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

type LOOKUP_SWITCH struct {
	defaultOffset int32 // 默认偏移量
	npairs        int32
	matchOffsets  []int32
}

func (self *LOOKUP_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.npairs = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.npairs * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	// 先从操作数栈中弹出一个int变量
	key := frame.OperandStack().PopInt()
	// 然后去查找matchOffsets，看是否能找到匹配的key
	for i := int32(0); i < self.npairs*2; i += 2 {
		if self.matchOffsets[i] == key {
			// 如果能匹配上，则按照value给出的偏移量进行跳转
			offset := self.matchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	// 否则，就按照defaultOffset进行跳转
	base.Branch(frame, int(self.defaultOffset))
}
