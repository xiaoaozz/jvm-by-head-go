package classspath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZipEntry 表示zip或者jar文件形式的类路径
type ZipEntry struct {
	absPath string          // 存放目录的绝对路径
	zipRC   *zip.ReadCloser // 提供打开zip文件的工具
}

// 构造函数
func newZipEntry(path string) *ZipEntry {
	// 将参数path转化为文件的绝对路径，如果出现错误，则调用panic函数种植程序执行
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	// 否则创建Entry实例并返回
	return &ZipEntry{absPath, nil}
}

// ZipEntry的readClass()方法实现
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 尝试打开zip文件，如果失败，直接返回
	if self.zipRC == nil {
		err := self.openJar()
		if err != nil {
			return nil, nil, err
		}
	}
	// 获取class文件
	classFile := self.findClass(className)
	if classFile == nil {
		return nil, nil, errors.New("class not found: " + className)
	}
	// 读取文件内容
	data, err := readClass(classFile)
	return data, self, err
}

func (self *ZipEntry) String() string {
	// 直接返回目录
	return self.absPath
}

// 获取打开压缩文件的工具类
func (self *ZipEntry) openJar() error {
	// 尝试打开zip文件，如果没有报错，则将该实例保存在ZipEntry中
	r, err := zip.OpenReader(self.absPath)
	if err == nil {
		self.zipRC = r
	}
	// 否则返回错误
	return err
}

// 遍历已经打开的压缩包文件中寻找class文件
func (self *ZipEntry) findClass(ClassName string) *zip.File {
	for _, f := range self.zipRC.File {
		if f.Name == ClassName {
			return f
		}
	}
	return nil
}

// 读取class文件内容
func readClass(classFile *zip.File) ([]byte, error) {
	rc, err := classFile.Open()
	if err != nil {
		return nil, err
	}
	// 读取文件内容
	data, err := ioutil.ReadAll(rc)
	rc.Close()
	if err != nil {
		return nil, err
	}
	return data, nil
}
