package classfile

// DeprecatedAttribute 用于指出类、接口、字段或者方法（已经不建议使用）
type DeprecatedAttribute struct {
	MarkerAttribute
}

// SyntheticAttribute 用来标记源文件种不存在、由编译器生成的类成员（支持嵌套类和嵌套接口）
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct{}

// readInfo 因为MarkerAttribute只是标记作用，所以方法为空
func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// 此方法为空
}
