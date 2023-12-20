package rtda

import "jvm-by-head-go/ch06/rtda/heap"

// Slot 槽
type Slot struct {
	num int32        // 存放整数
	ref *heap.Object // 存放引用
}
