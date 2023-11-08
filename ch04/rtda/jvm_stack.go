package rtda

// Stack Java虚拟机栈结构体
type Stack struct {
	maxSize uint   // 栈的容量，最多可以容纳多少栈帧
	size    uint   // 保存栈的当前大小
	_top    *Frame // 保存栈顶指针
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

// push 将栈帧推入栈顶
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

// pop 把栈顶栈帧弹出
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

// top 返回栈顶栈帧
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
