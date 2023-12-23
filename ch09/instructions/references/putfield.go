package references

import (
	"jvm-by-head-go/ch09/instructions/base"
	"jvm-by-head-go/ch09/rtda"
	"jvm-by-head-go/ch09/rtda/heap"
)

// PUT_FIELD putfield指令给实例变量赋值，需要三个操作数
type PUT_FIELD struct {
	// 前两个操作数是常量池索引和变量值
	// 第三个操作数是对象引用，从操作数栈中弹出
	base.Index16Instruction
}

func (self *PUT_FIELD) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()

	// 解析后的字段必须是实例字段，否则抛出IncompatibleClassChangeError异常
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 如果是final字段，则只能在构造函数中初始化，否则抛出IllegalAccessError异常
	if field.IsFinal() {
		if currentClass != field.Class() || currentMethod.Name() != "<init>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	stack := frame.OperandStack()
	// 根据字段类型从操作数栈中弹出对应的变量值，然后弹出对象引用
	// 如果引用是null，则需要抛出NullPointerException异常
	// 否则通过引用给实例变量赋值
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		val := stack.PopInt()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetInt(slotId, val)
	case 'F':
		val := stack.PopFloat()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetFloat(slotId, val)
	case 'J':
		val := stack.PopLong()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetLong(slotId, val)
	case 'D':
		val := stack.PopDouble()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetDouble(slotId, val)
	case 'L', '[':
		val := stack.PopRef()
		ref := stack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.Fields().SetRef(slotId, val)
	default:
		// todo
	}
}
