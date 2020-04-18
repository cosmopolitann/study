# GO语言初识

### 认识GO语言

Go 是一个开源的编程语言，它能让构造简单、可靠且高效的软件变得容易。

### Go 语言特色

- 简洁、快速、安全
- 并行、有趣、开源
- 内存管理、数组安全、编译迅速

### 第一个 Go 程序

###### 我们来编写第一个 Go 程序 hello.go（Go 语言源文件的扩展是 .go），代码如下：

```go
package main

import "fmt"

func main() {
  
  fmt.Println("Hello, World!")//打印 Hello,world!

}
```



###### 要执行 Go 语言代码可以使用  **go run** 命令。

###### 执行以上代码输出:

```go
$ go build hello.go 
$ ls
hello    hello.go
$ ./hello 
Hello, World!
```

### go bulid , go install ,go run 的区别。

###### go run：go run 编译并直接运行程序，它会产生一个临时文件（但不会生成 .exe 文件），直接在命令行输出程序执行结果，方便用户调试。

###### go build：go build 用于测试编译包，主要检查是否会有编译错误，如果是一个可执行文件的源码（即是 main 包），就会直接生成一个可执行文件。

###### go install：go install 的作用有两步：第一步是编译导入的包文件，所有导入的包文件编译完才会编译主程序；第二步是将编译后生成的可执行文件放到 bin 目录下，编译后的包文件放到 pkg 目录下（$GOPATH/pkg）。

### Go常用命令

```go
build: 编译包和依赖
doc: 显示包或者符号的文档
env: 打印go的环境信息
fmt: 格式化代码
get: 下载并安装包和依赖
install: 编译并安装包和依赖
run: 编译并运行go程序
test: 运行测试
version: 显示go的版本
```

——————————————————————————————————————————————————