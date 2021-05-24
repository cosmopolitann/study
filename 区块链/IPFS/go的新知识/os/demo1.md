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



#### 1.Hostname 返回内核提供的主机名。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("这是 获取 主机 名字 ")
	// Host name
	//func Hostname() (name string, err error)
	name, err := os.Hostname()
	if err != nil {
		fmt.Println("获取主机名字错误 ：", err)

	}
	fmt.Println("name = ", name)
}

//使用 debug 一步一步走进去看 返回的是字节  buf[:n]


apple@appledeMacBook-Air: ~/project/our/os/Hosename
$ go run hostname.go  

name =  appledeMacBook-Air.local

返回的 主机 名称 是 appledeMacBook-Air.local
```



#### 2.Getpagesize    返回底层的系统内存页的尺寸。

```go
package main

import (
	"fmt"
	"os"
)

func main() {


	size:=os.Getpagesize()
	fmt.Println("size = ", size)

  
  //返回结果 
  size =  4096

```



#### 3.Environ   返回表示环境变量的格式为"key=value"的字符串的切片拷贝

```go
package main

import (
	"fmt"
	"os"
)

func main() {

d:=os.Environ()
fmt.Println("d = ",d)
}

//返回的就是  env  一样的结果的   []string  是一个 数组。
//如果以后需要 可以直接 通过这个方法 然后 通过数组下标获取 对应所需要的  环境变量值

for 或者 range 。

```



#### 4.Getenv 检索并返回名为key的环境变量的值。如果不存在该环境变量会返回空字符串。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
d:=os.Getenv("PATH")
fmt.Println("d = ",d)
}

//结果

apple@appledeMacBook-Air: ~/project/our/os/Hosename
$ go run hostname.go                                                                                                                                                                                                         [16:51:54]

d =  /usr/local/go/bin:/Users/apple/go/bin:/usr/local/bin:/usr/local/sbin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/Library/Apple/usr/bin
```

#### 5.Setenv  设置名为key的环境变量。如果出错会返回该错误。

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	err:=os.Setenv("name","hello")
	if err!=nil{
		fmt.Println("err = ",err)

	}
	d:=os.Getenv("name")
   fmt.Println("d = ",d)
}

//结果

apple@appledeMacBook-Air: ~/project/our/os/Hosename
$ go run hostname.go                                                                                                                                                                                                         [16:54:25]
d =  hello

```



#### 6.Clearenv删除所有环境变量。

```go
// os.Clearenv()

```



#### 7.Exit  让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。

```go
package main

import (
	"os"
)

func main() {

	os.Exit(0)
   //fmt.Println("d = ",d)
}
```



#### 8.Getuid返回调用者的用户ID。

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	uid:=os.Getuid()
   fmt.Println("uid = ",uid)
}

//我返回的 uid =501
// 试了好几次  都是 固定的   应该是一样的  但是 不是很确定。试了 3次 都是 501 
// 应该 是 当前用户的 id 
```



#### 9.Getuid返回调用者的用户ID。

```go
package main

import (
   "fmt"
   "os"
)

func main() {

   uid:=os.Getuid()
   d:=os.Getuid()
   fmt.Println("d = ",d)

   fmt.Println("uid = ",uid)
}

//这是  Getuid  和  Geteuid   都是返回 的 501 这个值。

```



#### 10.Getegid返回调用者的有效组ID。

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	g:=os.Getegid()
	fmt.Println("g = ",g)

}

//这个 结果  g=20
// 不知道 为什么 这个 和 上面获取的  id  是不一样的 
// 返回调用者的有效组 ID。
//  这个 有效组 ID 到底是个什么？ 具体 到时候 在去查找一下。

```



#### 11.Getpid  返回调用者所在进程的进程ID。

```go
package main

import (
	"fmt"
	"os"
)
func main() {

p:=os.Getpid()

	fmt.Println("p = ",p)
}

//调用者  当前程序的 进程 pid
eg：

pid = 85108
```



#### 12.Stat 返回一个描述name指定的文件对象的FileInfo。如果指定的文件对象是一个符号链接，返回的FileInfo描述该符号链接指向的文件的信息，本函数会尝试跳转该链接。如果出错，返回的错误值为*PathError类型。

```go
package main

import (
	"fmt"
	"os"
)

func main() {

    s,_:=os.Stat("/Users/apple/project/our/os/Hosename")
	fmt.Println("s =",s)

    //是否 是一个文件夹
    b:=s.IsDir()
	fmt.Println("b =",b)

    file:=s.Mode()
	fmt.Println(" file =",file)
    s.Size()
   p:= s.Mode().Perm()
	fmt.Println(" perm  =",p)

}

//结果
apple@appledeMacBook-Air: ~/project/our/os/Hosename
$ go run hostname.go                                                                                                                                                                                                         [17:18:36]
s = &{Hosename 96 2147484141 {741071112 63757444714 0x115f0e0} {16777223 16877 3 12888190913 501 20 0 [0 0 0 0] {1621847915 928682680} {1621847914 741071112} {1621847914 741071112} {1621837221 26108249} 96 0 4096 0 0 0 [0 0]}}
b = true
 file = drwxr-xr-x
 perm  = -rwxr-xr-x

```



#### 13. File 文件操作

##### 首先，file 类是在 os 包中的，封装了底层的文件描述符和相关信息，同时封装了 Read 和 Write 的实现。



FileInfo 接口

FileInfo 接口中定义了 File 信息相关的方法。

```csharp
type FileInfo interface {
    Name() string       // 文件的名字（不含扩展名）
    Size() int64        // 普通文件返回值表示其大小；其他文件的返回值含义各系统不同
    Mode() FileMode     // 文件的模式位
    ModTime() time.Time // 文件的修改时间
    IsDir() bool        // 等价于Mode().IsDir()
    Sys() interface{}   // 底层数据来源（可以返回nil）
}
```

##### 权限操作：

至于操作权限 perm，除非创建文件时才需要指定，不需要创建新文件时可以将其设定为０。虽然 Golang 语言给 perm 权限设定了很多的常量，但是习惯上也可以直接使用数字，如 0666 (具体含义和 Unix 系统的一致)。
 权限控制：

```go
linux 下有2种文件权限表示方式，即“符号表示”和“八进制表示”。

（1）符号表示方式:
-      ---         ---        ---
type   owner       group      others
文件的权限是这样子分配的 读 写 可执行 分别对应的是 r w x 如果没有那一个权限，用 - 代替
(-文件 d目录 |连接符号)
例如：-rwxr-xr-x

（2）八进制表示方式： 
r ——> 004
w ——> 002
x ——> 001
- ——> 000

0755
0777
0555
0444
0666
```

实例代码

```go
package main

import (
    "os"
    "fmt"
)

func main() {
    /*
    FileInfo：文件信息
        interface
            Name()，文件名
            Size()，文件大小，字节为单位
            IsDir()，是否是目录
            ModTime()，修改时间
            Mode()，权限

     */
    fileInfo,err :=  os.Stat("/Users/ruby/Documents/pro/a/aa.txt")
    if err != nil{
        fmt.Println("err :",err)
        return
    }
    fmt.Printf("%T\n",fileInfo)
    //文件名
    fmt.Println(fileInfo.Name())
    //文件大小
    fmt.Println(fileInfo.Size())
    //是否是目录
    fmt.Println(fileInfo.IsDir()) //IsDirectory
    //修改时间
    fmt.Println(fileInfo.ModTime())
    //权限
    fmt.Println(fileInfo.Mode()) //-rw-r--r--
}
```



##### File 操作

```go
type File
//File代表一个打开的文件对象。

func Create(name string) (file *File, err error)
//Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）。如果成功，返回的文件对象可用于I/O；对应的文件描述符具有O_RDWR模式。如果出错，错误底层类型是*PathError。

func Open(name string) (file *File, err error)
//Open打开一个文件用于读取。如果操作成功，返回的文件对象的方法可用于读取数据；对应的文件描述符具有O_RDONLY模式。如果出错，错误底层类型是*PathError。

func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
//OpenFile是一个更一般性的文件打开函数，大多数调用者都应用Open或Create代替本函数。它会使用指定的选项（如O_RDONLY等）、指定的模式（如0666等）打开指定名称的文件。如果操作成功，返回的文件对象可用于I/O。如果出错，错误底层类型是*PathError。

func NewFile(fd uintptr, name string) *File
//NewFile使用给出的Unix文件描述符和名称创建一个文件。

func Pipe() (r *File, w *File, err error)
//Pipe返回一对关联的文件对象。从r的读取将返回写入w的数据。本函数会返回两个文件对象和可能的错误。

func (f *File) Name() string
//Name方法返回（提供给Open/Create等方法的）文件名称。

func (f *File) Stat() (fi FileInfo, err error)
//Stat返回描述文件f的FileInfo类型值。如果出错，错误底层类型是*PathError。

func (f *File) Fd() uintptr
//Fd返回与文件f对应的整数类型的Unix文件描述符。

func (f *File) Chdir() error
//Chdir将当前工作目录修改为f，f必须是一个目录。如果出错，错误底层类型是*PathError。

func (f *File) Chmod(mode FileMode) error
//Chmod修改文件的模式。如果出错，错误底层类型是*PathError。

func (f *File) Chown(uid, gid int) error
//Chown修改文件的用户ID和组ID。如果出错，错误底层类型是*PathError。

func (f *File) Close() error
//Close关闭文件f，使文件不能用于读写。它返回可能出现的错误。

func (f *File) Readdir(n int) (fi []FileInfo, err error)
//Readdir读取目录f的内容，返回一个有n个成员的[]FileInfo，这些FileInfo是被Lstat返回的，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息。如果n>0，Readdir函数会返回一个最多n个成员的切片。这时，如果Readdir返回一个空切片，它会返回一个非nil的错误说明原因。如果到达了目录f的结尾，返回值err会是io.EOF。如果n<=0，Readdir函数返回目录中剩余所有文件对象的FileInfo构成的切片。此时，如果Readdir调用成功（读取所有内容直到结尾），它会返回该切片和nil的错误值。如果在到达结尾前遇到错误，会返回之前成功读取的FileInfo构成的切片和该错误。

func (f *File) Readdirnames(n int) (names []string, err error)
//Readdir读取目录f的内容，返回一个有n个成员的[]string，切片成员为目录中文件对象的名字，采用目录顺序。对本函数的下一次调用会返回上一次调用剩余未读取的内容的信息。如果n>0，Readdir函数会返回一个最多n个成员的切片。这时，如果Readdir返回一个空切片，它会返回一个非nil的错误说明原因。如果到达了目录f的结尾，返回值err会是io.EOF。如果n<=0，Readdir函数返回目录中剩余所有文件对象的名字构成的切片。此时，如果Readdir调用成功（读取所有内容直到结尾），它会返回该切片和nil的错误值。如果在到达结尾前遇到错误，会返回之前成功读取的名字构成的切片和该错误。

func (f *File) Truncate(size int64) error
//Truncate改变文件的大小，它不会改变I/O的当前位置。 如果截断文件，多出的部分就会被丢弃。如果出错，错误底层类型是*PathError。

```



##### 打开模式

```go
const (
    O_RDONLY int = syscall.O_RDONLY // 只读模式打开文件
    O_WRONLY int = syscall.O_WRONLY // 只写模式打开文件
    O_RDWR   int = syscall.O_RDWR   // 读写模式打开文件
    O_APPEND int = syscall.O_APPEND // 写操作时将数据附加到文件尾部
    O_CREATE int = syscall.O_CREAT  // 如果不存在将创建一个新文件
    O_EXCL   int = syscall.O_EXCL   // 和 O_CREATE 配合使用，文件必须不存在
    O_SYNC   int = syscall.O_SYNC   // 打开文件用于同步 I/O
    O_TRUNC  int = syscall.O_TRUNC  // 如果可能，打开时清空文件
)
```



##### 文件操作实例：

```go
package main

import (
    "fmt"
    "path/filepath"
    "path"
    "os"
)

func main() {
    /*
    文件操作：
    1.路径：
        相对路径：relative
            ab.txt
            相对于当前工程
        绝对路径：absolute
            /Users/ruby/Documents/pro/a/aa.txt

        .当前目录
        ..上一层
    2.创建文件夹，如果文件夹存在，创建失败
        os.MkDir()，创建一层
        os.MkDirAll()，可以创建多层

    3.创建文件，Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
        os.Create()，创建文件

    4.打开文件：让当前的程序，和指定的文件之间建立一个连接
        os.Open(filename)
        os.OpenFile(filename,mode,perm)

    5.关闭文件：程序和文件之间的链接断开。
        file.Close()

    5.删除文件或目录：慎用，慎用，再慎用
        os.Remove()，删除文件和空目录
        os.RemoveAll()，删除所有
     */
     //1.路径
     fileName1:="/Users/ruby/Documents/pro/a/aa.txt"
     fileName2:="bb.txt"
     fmt.Println(filepath.IsAbs(fileName1)) //true
     fmt.Println(filepath.IsAbs(fileName2)) //false
     fmt.Println(filepath.Abs(fileName1))
     fmt.Println(filepath.Abs(fileName2)) // /Users/ruby/go/src/l_file/bb.txt

     fmt.Println("获取父目录：",path.Join(fileName1,".."))

     //2.创建目录
     //err := os.Mkdir("/Users/ruby/Documents/pro/a/bb",os.ModePerm)
     //if err != nil{
     // fmt.Println("err:",err)
     // return
     //}
     //fmt.Println("文件夹创建成功。。")
     //err :=os.MkdirAll("/Users/ruby/Documents/pro/a/cc/dd/ee",os.ModePerm)
     //if err != nil{
     // fmt.Println("err:",err)
     // return
     //}
     //fmt.Println("多层文件夹创建成功")

     //3.创建文件:Create采用模式0666（任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（为空文件）
     //file1,err :=os.Create("/Users/ruby/Documents/pro/a/ab.txt")
     //if err != nil{
     // fmt.Println("err：",err)
     // return
     //}
     //fmt.Println(file1)

     //file2,err := os.Create(fileName2)//创建相对路径的文件，是以当前工程为参照的
     //if err != nil{
     // fmt.Println("err :",err)
     // return
     //}
     //fmt.Println(file2)

     //4.打开文件：
     //file3 ,err := os.Open(fileName1) //只读的
     //if err != nil{
     // fmt.Println("err:",err)
     // return
     //}
     //fmt.Println(file3)
    /*
    第一个参数：文件名称
    第二个参数：文件的打开方式
        const (
    // Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
        O_RDONLY int = syscall.O_RDONLY // open the file read-only.
        O_WRONLY int = syscall.O_WRONLY // open the file write-only.
        O_RDWR   int = syscall.O_RDWR   // open the file read-write.
        // The remaining values may be or'ed in to control behavior.
        O_APPEND int = syscall.O_APPEND // append data to the file when writing.
        O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
        O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
        O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
        O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
    )
    第三个参数：文件的权限：文件不存在创建文件，需要指定权限
     */
     //file4,err := os.OpenFile(fileName1,os.O_RDONLY|os.O_WRONLY,os.ModePerm)
     //if err != nil{
     // fmt.Println("err:",err)
     // return
     //}
     //fmt.Println(file4)

     //5关闭文件，
     //file4.Close()

     //6.删除文件或文件夹：
     //删除文件
    //err :=  os.Remove("/Users/ruby/Documents/pro/a/aa.txt")
    //if err != nil{
    //  fmt.Println("err:",err)
    //  return
    //}
    //fmt.Println("删除文件成功。。")
    //删除目录
    err :=  os.RemoveAll("/Users/ruby/Documents/pro/a/cc")
    if err != nil{
        fmt.Println("err:",err)
        return
    }
    fmt.Println("删除目录成功。。")
}
```



#### 判断文件是否存在

Golang 判断文件或文件夹是否存在的方法为使用 os.Stat() 函数返回的错误值进行判断。

1. 如果返回的错误为 nil，说明文件或文件夹存在
2. 如果返回的错误类型使用 os.IsNotExist() 判断为 true，说明文件或文件夹不存在

```go
package main
import (
    "log"
    "os"
)

func main() {
    fileInfo,err:=os.Stat("/Users/ruby/Documents/pro/a/aa.txt")
    if err!=nil{
        if os.IsNotExist(err){
            log.Fatalln("file does not exist")
        }
    }
    log.Println("file does exist. file information:")
    log.Println(fileInfo)
}
```

