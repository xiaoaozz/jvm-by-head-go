package lang

import (
	"jvm-by-head-go/ch09/native"
	"jvm-by-head-go/ch09/rtda"
	"jvm-by-head-go/ch09/rtda/heap"
)

const jlString = "java/lang/String"

func init() {
	native.Register(jlString, "intern", "()Ljava/lang/String;", intern)
}

// public native String intern();
// ()Ljava/lang/String;
func intern(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	interned := heap.InternString(this)
	frame.OperandStack().PushRef(interned)
}
