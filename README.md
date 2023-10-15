# jvm-by-head-go

手写JVM的Go语言版本

## 项目介绍

### 组织架构

```
jvm-by-head-go
    │  go.mod  依赖包管理 
    │  README.md 项目描述
    ├─ch01 简单的命令行工具
    │      cmd.go
    │      main.go
    │
    └─ch02 寻找classpath文件
        │  cmd.go
        │  main.go
        │
        └─classspath 
                classpath.go
                entry.go 组合模式实现 
                entry_composite.go
                entry_dir.go
                entry_wildcard.go
                entry_zip.go
```

### 功能架构

#### 系统架构

#### 功能模块

#### 开发进度
- 2023-09-25 项目环境搭建，文档编写，以及准备go、java双语言开发
- 2023-10-14 命令行工具开发完成（简易版），测试通过
- 2023-10-15 搜索class文件模块完成，并使用组合模式优化，测试通过
-
### 知识架构

