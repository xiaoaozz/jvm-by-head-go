package native

import "jvm-by-head-go/ch09/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}

// Register 注册本地方法
func Register(className, methodName, methodDescriptor string, method NativeMethod) {
	// 类名、方法名和方法描述符加在一起才能唯一确定一个方法，所以将它们组合作为本地方法注册表的键
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// FindNativeMethod 查找本地方法
func FindNativeMethod(className, methodName, methodDescriptor string) NativeMethod {
	// 根据类名、方法名和方法描述符查找本地方法的实现
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return emptyNativeMethod
	}
	// 如果找不到，则返回nil
	return nil
}
