### 通过学习go-ipfs 源码，学习新语法

### 1.os 包

```go
1.通过 os 包 获取 环境变量的值
2.在终端中 输入 env  可以获取 所有环境变量，也可以在 profile 文件中查看。mac 也可以在 .bashsc 文件中查看
//如下是 env


$ env                                                                                                                                                                                                                                   [14:08:05]
USER=apple
PATH=/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/Library/Apple/usr/bin:/usr/local/sbin
LOGNAME=apple
SSH_AUTH_SOCK=/private/tmp/com.apple.launchd.NznQN3VAci/Listeners
HOME=/Users/apple
SHELL=/bin/zsh
__CF_USER_TEXT_ENCODING=0x1F5:0x19:0x34
TMPDIR=/var/folders/g9/1hx2zwss3wd5c3km241g07080000gn/T/
XPC_SERVICE_NAME=0
XPC_FLAGS=0x0
ORIGINAL_XDG_CURRENT_DESKTOP=undefined
SHLVL=1
PAGER=less

太多了就 不 过多粘贴了   




```

#### GO 中 使用  举个例子  获取 PAGER=less 的值

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Test~~~~")
	//获取 env 中 某个参数的值

	pager := os.Getenv("PAGER")
	fmt.Println("pager := ", pager)
}

打印 输出 ：
==============
$ go run main.go                                                                                                                                                                                                                        [14:08:01]
Test~~~~
pager :=  less

```

