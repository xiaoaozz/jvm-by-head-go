package rtda

import "math"

// LocalVars 局部变量表
type LocalVars []Slot

// newLocalVars 构造函数
func newLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

// SetInt int变量 setter方法
func (self LocalVars) SetInt(index uint, val int32) {
	self[index].num = val
}

// GetInt int变量 getter方法
func (self LocalVars) GetInt(index uint) int32 {
	return self[index].num
}

// SetFloat float变量可以先转成int变量，然后按照int变量处理 setter方法
func (self LocalVars) SetFloat(index uint, val float32) {
	bits := math.Float32bits(val)
	self[index].num = int32(bits)
}

// GetFloat float变量可以先转成int变量，然后按照int变量处理 setter方法
func (self LocalVars) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

// SetLong long变量可以拆成两个int变量
func (self LocalVars) SetLong(index uint, val int64) {
	self[index].num = int32(val)
	self[index+1].num = int32(val >> 32)
}

// GetLong 将两个int变量组合成一个long变量
func (self LocalVars) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index+1].num)
	return int64(high)<<32 | int64(low)
}

// SetDouble double先转化成long，然后按照long处理
func (self LocalVars) SetDouble(index uint, val float64) {
	bits := math.Float64bits(val)
	self.SetLong(index, int64(bits))
}

// GetDouble 获取double变量
func (self LocalVars) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

// SetRef 设置引用类型
func (self LocalVars) SetRef(index uint, ref *Object) {
	self[index].ref = ref
}

// GetRef 获取引用类型
func (self LocalVars) GetRef(index uint) *Object {
	return self[index].ref
}
