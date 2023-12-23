package references

import (
	"jvm-by-head-go/ch09/instructions/base"
	"jvm-by-head-go/ch09/rtda"
	"jvm-by-head-go/ch09/rtda/heap"
)

// INVOKE_SPECIAL 对超类、私有和实例初始化方法调用的特殊处理
type INVOKE_SPECIAL struct{ base.Index16Instruction }

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	// 拿到当前类、当前常量池、方法符号引用，然后解析符号引用。拿到解析后的类和方法
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	methodRef := cp.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	// 假定从方法符号引用中解析出来的类是C，方法是M

	// 如果M是构造函数，则声明M的类必须是C，否则抛出NoSuchMethodError异常
	if resolvedMethod.Name() == "<init>" && resolvedMethod.Class() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	// 如果M是静态方法，则抛出IncompatibleClassChangeError异常
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	// 从操作数栈中弹出this引用
	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount() - 1)
	// 如果该引用是null，则抛出NullPointerException异常
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// 确保protected方法只能被声明方法的类或者子类调用，否则抛出IllegalAccessError异常
	if resolvedMethod.IsProtected() &&
		resolvedMethod.Class().IsSuperClassOf(currentClass) &&
		resolvedMethod.Class().GetPackageName() != currentClass.GetPackageName() &&
		ref.Class() != currentClass &&
		!ref.Class().IsSubClassOf(currentClass) {

		panic("java.lang.IllegalAccessError")
	}

	// 如果调用的中超类中的函数，但不是构造函数，且当前类的ACC_SUPER标志被设置，需要一个额外的过程查找最终要调用的方法
	// 否则前面从方法符号引用中解析出来的方法就是要调用的方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.Name() != "<init>" {

		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(),
			methodRef.Name(), methodRef.Descriptor())
	}
	// 如果查找过程失败，或者找到的方法是抽象的，则抛出AbstractMethodError异常
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	// 否则就一切正常调用处理方法
	base.InvokeMethod(frame, methodToBeInvoked)
}
