package rtda

// Frame 栈帧结构
type Frame struct {
	lower        *Frame        // 用来实现链表数据结构
	localVars    LocalVars     // 保存局部变量表指针
	operandStack *OperandStack // 保存操作数栈指针
	thread       *Thread
	nextPC       int
}

// NewFrame 构造函数
func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

// LocalVars 获取局部变量表
func (f Frame) LocalVars() LocalVars {
	return f.localVars
}

// OperandStack 获取操作数栈
func (f Frame) OperandStack() *OperandStack {
	return f.operandStack
}

// Thread 获取当前线程
func (f Frame) Thread() *Thread {
	return f.thread
}

// NextPC 获取下一个pc指针
func (f Frame) NextPC() int {
	return f.nextPC
}

// SetNextPC 设置下一个pc指针的值
func (f Frame) SetNextPC(nextPC int) {
	f.nextPC = nextPC
}
