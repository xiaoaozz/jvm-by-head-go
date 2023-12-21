package references

import (
	"jvm-by-head-go/ch06/instructions/base"
	"jvm-by-head-go/ch06/rtda"
	"jvm-by-head-go/ch06/rtda/heap"
)

// PUT_STATIC putstatic指令给类的某个静态变量赋值
type PUT_STATIC struct {
	// putstatic指令需要两个操作数
	// 第一个操作数是uint16索引，来自字节码，通过这个索引可以从当前类的运行时常量池中找到第一个字段符号引用
	// 解析这个符号引用就可以知道要给哪个静态变量赋值
	// 第二个操作数是要赋给静态变量的值，从操作数栈中弹出
	base.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	// 获取当前方法
	currentMethod := frame.Method()
	// 获取当前类
	currentClass := currentMethod.Class()
	// 获取当前常量池
	cp := currentClass.ConstantPool()
	fieldRef := cp.GetConstant(self.Index).(*heap.FieldRef)
	// 解析字段符号引用
	field := fieldRef.ResolvedField()
	class := field.Class()
	// todo 如果声明字段的类还没有实现，则需要先初始化类

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	// 根据字段类型从操作数栈中弹出相应的值，然后赋给静态变量
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	default:
		// todo
	}
}
