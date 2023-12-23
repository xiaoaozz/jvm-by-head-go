package main

import (
	"flag"
	"fmt"
	"os"
)

// Cmd 命令行结构体
type Cmd struct {
	helpFlag         bool     // 获取帮助标志
	versionFlag      bool     // 版本标志
	verboseClassFlag bool     // -verbose:class选项，控制是否把类加载信息输出到控制台
	verboseInstFlag  bool     // -verbose:inst选项，控制是否把指令执行信息输出到控制台
	XjreOption       string   // 非标准选项-Xjre
	cpOption         string   // classpath
	class            string   // 类名
	args             []string // 参数数组
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printUsage() {
	fmt.Printf("Usage:%s [-options] class [args...] \n", os.Args[0])
}
