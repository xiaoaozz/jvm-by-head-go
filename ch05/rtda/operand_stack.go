package rtda

import (
	"math"
)

// OperandStack 操作数栈
type OperandStack struct {
	size  uint // 记录栈顶位置
	slots []Slot
}

// newOperandStack 构造函数
func newOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}
	return nil
}

// PushInt 往栈顶放一个int变量，然后把size++
func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++
}

// PopInt 先把size--，然后弹出一个int变量
func (self *OperandStack) PopInt() int32 {
	self.size--
	return self.slots[self.size].num
}

// PushFloat float变量先转成int变量，然后再处理
func (self *OperandStack) PushFloat(val float32) {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

// PopFloat float变量先转成int变量，然后再处理
func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return math.Float32frombits(bits)
}

// PushLong long类型变量拆成2个int变量处理
func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size+1].num = int32(val >> 32)
	self.size += 2
}

// PopLong 将2个int变量组合成一个long变量
func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size+1].num)
	return int64(high)<<32 | int64(low)
}

// PushDouble double变量先转化成long变量
func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

// PopDouble double变量先转化成long变量
func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

// PushRef 引用类型变量处理
func (self *OperandStack) PushRef(ref *Object) {
	self.slots[self.size].ref = ref
	self.size++
}

// PopRef 引用类型变量处理
func (self *OperandStack) PopRef() *Object {
	self.size--
	ref := self.slots[self.size].ref
	// 帮助垃圾收集器回收结构体实例
	self.slots[self.size].ref = nil
	return ref
}

// PushSlot 栈指令处理
func (self *OperandStack) PushSlot(slot Slot) {
	self.slots[self.size] = slot
	self.size++
}

// PopSlot 栈指令处理
func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}
