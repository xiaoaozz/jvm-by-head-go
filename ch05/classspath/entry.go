package classspath

import (
	"os"
	"strings"
)

// 系统路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 负责寻找和加载class文件
	readClass(ClassName string) ([]byte, Entry, error)
	// String 相当于Java中的toString()方法，返回变量的字符串表示
	String() string
}

// 构造函数，用于实现不同的classpath路径的寻找方式，使用组合模式来设计
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		// 如果path含有分隔符，
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	// 返回目录形式的类路径
	return newDirEntry(path)
}
