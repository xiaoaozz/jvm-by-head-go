package rtda

// Frame 栈帧结构
type Frame struct {
	lower        *Frame        // 用来实现链表数据结构
	localVars    LocalVars     // 保存局部变量表指针
	operandStack *OperandStack // 保存操作数栈指针
}

// NewFrame 构造函数
func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
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
