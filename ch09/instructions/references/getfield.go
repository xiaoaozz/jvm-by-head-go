package references

import (
	"jvm-by-head-go/ch09/instructions/base"
	"jvm-by-head-go/ch09/rtda"
	"jvm-by-head-go/ch09/rtda/heap"
)

// GET_FIELD 该指令获取对象的实例变量值，然后推入操作数栈
type GET_FIELD struct {
	// 需要两个操作数
	// 第一个操作数是uint16索引
	// 第二个操作数是对象引用
	base.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	// 字段符号引用解析
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 弹出对象引用，如果是null，则抛出NullPointerException异常
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()
	// 根据字段类型，获取相应的实例变量值，然后推入操作数栈
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
		// todo
	}
}
