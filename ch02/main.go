package main

import (
	"fmt"
	"jvm-by-head-go/ch02/classspath"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Printf("version 0.0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classspath.Parse(cmd.XjreOption, cmd.cpOption)
	// 打印命令行参数
	fmt.Printf("claspath: %v class: %v args:%v\n", cp, cmd.class, cmd.args)
	// 读取主类数据
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	// 将数据打印到控制台
	if err != nil {
		fmt.Printf("Could not find or load main class %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
