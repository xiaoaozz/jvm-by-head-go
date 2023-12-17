package rtda

// Thread 线程
type Thread struct {
	pc    int    // pc寄存器指针
	stack *Stack // 虚拟机栈
}

// NewThread 构造函数
func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}

func (self *Thread) PC() int {
	return self.pc
}

func (self *Thread) SetPc(pc int) {
	self.pc = pc
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// CurrentFrame 返回当前栈帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
