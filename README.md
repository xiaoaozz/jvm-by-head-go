# jvm-by-head-go

## 项目介绍

### 概述

手写JVM的Go语言版本，代码参照张秀宏老师的《自己动手写Java虚拟机》一书，并且会加上自己的思考，对代码结构进行一定的优化。
源代码地址：[传送门](https://github.com/zxh0/jvmgo-book)

### 组织架构

```               
jvm-by-head-go
    │  go.mod 依赖包管理 
    │  README.md 项目描述
    ├─ch01 简单的命令行工具
    │      cmd.go
    │      main.go
    │
    ├─ch02
    │  └─classspath 寻找classpath文件
    │          classpath.go
    │          entry.go 组合模式实现 
    │          entry_composite.go
    │          entry_dir.go
    │          entry_wildcard.go
    │          entry_zip.go
    │
    └─ch03
        └─classfile 解析class文件 
              attribute_info.go  attr属性相关定义
              attr_code.go
              attr_constant_value.go
              attr_exceptions.go
              attr_line_number_table.go
              attr_local_variable_table.go
              attr_markers.go
              attr_source_file.go
              attr_unparsed.go
              class_file.go class文件读取相关定义
              class_reader.go
              constant_info.go constant常量池相关定义
              constant_pool.go
              cp_class.go 
              cp_invoke_dynamic.go
              cp_member_ref.go
              cp_name_and_type.go
              cp_numberic.go
              cp_string.go
              cp_utf8.go
              member_info.go



```

## 功能架构

### 系统架构

### 功能模块

* 简易命令行
* 解析class文件

## 开发进度

- 2023-09-25 项目环境搭建，文档编写，以及准备go、java双语言开发
- 2023-10-14 命令行工具开发完成（简易版），测试通过
- 2023-10-15 搜索class文件模块完成，并使用组合模式优化，测试通过
- 2023-10-21 class文件结构定义、常量池结构定义完成
- 2023-10-22 class文件解析开发完成，测试通过

## 知识架构

### 五、Java虚拟机指令集和解释器

#### 1.字节码与数据类型

在Java虚拟机的指令集中，大多数指令都包含其操作所对应的数据类型信息。
比如：

* i 代表对 int 类型的数据操作
* l 代表对 long 类型的数据操作
* s 代表 short 类型的数据操作
* b 代表 byte 类型的数据操作
* c 代表 char 类型的数据操作
* f 代表 float 类型的数据操作
* d 代表 double 类型的数据操作
* a 代表 reference 类型的数据操作

也有一些指令的助记符中没有明确指明操作类型的字母，如array length指令。

#### 2.加载和存储指令

加载和存储指令用于将数据在**栈帧**中的**局部变量表和操作数栈**之间来回传输。

* 将一个局部变量加载到操作栈

```java
iload、iload_<n>、lload、lload_<n>、fload、fload_<n>、dload、dload_<n>、aload、aload_<n>
```

* 将一个数值从操作数栈存储到局部变量表

```java
istore、istore_<n>、lstore、lstore_<n>、fstore、fstore_<n>、dstore、dstore_<n>、astore、astore_<n>
```

* 将一个常量加载到操作数栈

```java 
    bipush、sipush、ldc、ldc_w、ldc2_w、aconst_null、iconst_m1、iconst_<i>、lconst_<l>、fconst_<f>、dconst_<d>
```

* 扩充局部变量表的访问索引的指令

```java 
    wide
```

#### 6.操作数栈管理指令

Java虚拟机提供了一些直接用于操作操作数栈的指令。

* 将操作数栈的栈顶的一个或者两个元素出栈

```java 
    pop、pop2
```

* 复制栈顶一个或两个数值并将复制值或双份的复制值重新压入栈顶

```java 
    dup、dup2、dup_x1、dup2_x1、dup_x2、dup2_x2
```

* 将栈最顶端的两个数值互换

```java 
    swap
```

