package classspath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry // 启动类路径
	extClasspath  Entry // 扩展类路径
	userClasspath Entry // 用户类路径
}

// Parse 解析路径
func Parse(jreOpetion, cpOption string) *Classpath {
	cp := &Classpath{}
	// 使用-Xjre选项解析启动类路径和扩展类路径
	cp.parseBootAndExtClasspath(jreOpetion)
	// 使用-classpath/-cp选项解析用户类路径
	cp.parseUserClasspath(cpOption)
	return cp
}

// ReadClass 依次从启动类路径、扩展类路径和用户类路径中搜索class文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

// String 返回用户类路径的字符串表示
func (self *Classpath) String() string {
	return self.userClasspath.String()
}

// parseBootAndExtClasspath 解析启动类加载器和扩展类加载器的路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 获取jre目录
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// parseUserClasspath 解析用户类加载器的路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	// 如果用户没有提供 -classpath/-cp选项，则使用当前目录作为用户类路径
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// getJreDir 获取jre的目录路径
func getJreDir(jreOption string) string {
	// 如果传入的jreOption非空并且目录存在，则返回jreOption
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果当前目录下存在名为“jre”的目录，则返回该目录的路径
	if exists("./jre") {
		return "./jre"
	}
	// 如果环境变量中有JAVA_HOME，并且JAVA_HOME目录下存在名为“jre”的子目录
	// 则返回JAVA_HOME目录下的“jre”的路径
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	// 以上条件都不满足，则抛出异常，表示无法找打jre文件夹
	panic("Can not find jre folder!")
}

// exists 判断目录是否存在
func exists(path string) bool {
	// 使用os.Stat检查文件或者目录是否存在，如果出现错误，说明文件或目录不存在
	if _, err := os.Stat(path); err != nil {
		// 如果错误是文件或者目录不存在，则返回false，否则返回true
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
