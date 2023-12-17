package classspath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	// 去掉末尾的*号，得到baseDir
	baseDir := path[:len(path)-1]
	compositeEntry := []Entry{}
	// 调用filepath包的Walk()函数遍历baseDir创建ZipEntry
	// walkFn适用于遍历目录
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// 如果是目录并且不是baseDir本身，则跳过它
		if info.IsDir() && path != baseDir {
			// 返回SkipDir跳过子目录
			return filepath.SkipDir
		}
		// 如果文件是以“.jar”或者“.JAR”结尾的，则创建ZipEntry
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	}
	// 使用filepath.Walk遍历目录
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
