package classspath

import (
	"io/ioutil"
	"path/filepath"
)

// DirEntry 表示目录形式的类路径
type DirEntry struct {
	absDir string // 存放目录的绝对路径
}

// 构造函数
func newDirEntry(path string) *DirEntry {
	// 将参数path转化为文件的绝对路径，如果出现错误，则调用panic函数种植程序执行
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// 否则创建Entry实例并返回
	return &DirEntry{absDir}
}

// DirEntry的readClass()方法实现
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 将目录和class文件名拼成一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	// 调用ioutil包提供的ReadFile()函数读取class文件内容
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	// 直接返回目录
	return self.absDir
}
