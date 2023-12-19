package main

import (
	"fmt"
	"jvm-by-head-go/ch05/classfile"
	"jvm-by-head-go/ch05/instructions"
	"jvm-by-head-go/ch05/instructions/base"
	"jvm-by-head-go/ch05/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute() // 获取Code属性
	maxLocals := codeAttr.MaxLocals()      // 获取执行方法所需的局部变量表
	maxStack := codeAttr.MaxStack()        // 获取执行方法所需的操作数栈空间
	bytecode := codeAttr.Code()            // 获取执行方法的字节码
	// 先创建一个Thread实例
	thread := rtda.NewThread()
	// 然后创建一个帧并把它推入Java虚拟机栈顶
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	// 执行方法
	defer catchErr(frame)
	loop(thread, bytecode)
}

// catchErr 暂时先用这种方式让解释器结束运行
func catchErr(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("LocalVars:%v\n", frame.LocalVars())
		fmt.Printf("OperandStack:%v\n", frame.OperandStack())
		panic(r)
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := &base.BytecodeReader{}
	for {
		// 计算pc
		pc := frame.NextPC()
		thread.SetPc(pc)

		// 解码指令
		reader.Reset(bytecode, pc)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.SetNextPC(reader.PC())

		// 执行指令
		fmt.Printf("pc:%2d inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
