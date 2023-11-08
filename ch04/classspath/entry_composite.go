package classspath

import (
	"errors"
	"strings"
)

// CompositeEntry Entry数组
type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	// 创建Entry切片数组
	var compositeEntry []Entry
	// 把参数（路径列表）按分隔符分成小路径，然后把每个小路径都转换成具体的Entry实例
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		// 依次调用每一个子路径的readClass()方法
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	// 依次调用每一个子路径的String()方法，然后使用系统路径分隔符拼接起来
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
