# IPFS Cli 使用

##  BASIC COMMAND

### init

```go
//初始化节点信息
 ipfs init

//命令使用
$ ipfs init                                                                                                   [14:00:38]
generating ED25519 keypair...done        //使用ED25519算法，生成秘钥对完成，
peer identity: 12D3KooWP18zfh57ryRe9DaQ6aYr97dCJbD9fU9Dg75pVbpbNHqw       //节点身份信息 ,就是节点 ID
initializing IPFS node at /Users/apple/.ipfs                         //生成ipfs 节点信息 存放在/Users/apple/.ipfs
to get started, enter:               //开始  加入

	ipfs cat /ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme
这个是 查看 ipfs 上面的 readme 信息。

```

### cat

```go
ipfs cat fileCid

后面跟 文件 的 CID

eg初始化节点的时候  可以查看 readme 信息：

ipfs cat /ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme

$ ipfs cat /ipfs/QmQPeNsJPyVWPFDVHb77w8G42Fvo15z4bG2X8D2GhfbSXc/readme                                        [14:09:25]
Hello and Welcome to IPFS!

██╗██████╗ ███████╗███████╗
██║██╔══██╗██╔════╝██╔════╝
██║██████╔╝█████╗  ███████╗
██║██╔═══╝ ██╔══╝  ╚════██║
██║██║     ██║     ███████║
╚═╝╚═╝     ╚═╝     ╚══════╝

If you're seeing this, you have successfully installed
IPFS and are now interfacing with the ipfs merkledag!

 -------------------------------------------------------
| Warning:                                              |
|   This is alpha software. Use at your own discretion! |
|   Much is missing or lacking polish. There are bugs.  |
|   Not yet secure. Read the security notes for more.   |
 -------------------------------------------------------

Check out some of the other files in this directory:

  ./about
  ./help
  ./quick-start     <-- usage examples
  ./readme          <-- this file
  ./security-notes


//查看 cat  的命令 帮助 
ipfs cat -h 
//USAGE
  ipfs cat <ipfs-path>... - Show IPFS object data.

  ipfs cat [--offset=<offset> | -o] [--length=<length> | -l] [--] <ipfs-path>...

  Displays the data contained by an IPFS or IPNS object(s) at the given path.

  For more information about each command, use:
  'ipfs cat <subcmd> --help

eg:
//如果 要查看 一个 文件信息 内容  使用
ipfs cat QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T 

hello

//可以加上 参数   -o= ？   -l=？     ？等于数字 

-o  就是 offset  从多少开始 偏移    -l  就是 输出 多少个字节。
///例如 ：
ipfs cat -o=2 -l=2  QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T

则 ： ll      (hello)

```



### add

```go
//添加信息
ipfs add [path]
//添加文件
eg:
$ ipfs add /Users/apple/winter/ipfs/i.txt                                                                     [14:37:51]
added QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T i.txt
 6 B / 6 B [====================================================================================================] 100.00%
 
 如果 是当前目录下 就直接 添加文件就行了
 ipfs add i.txt

 //添加文件夹
$ ipfs add -r /Users/apple/winter/ipfs                                                                        [14:39:20]
added QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T ipfs/i.txt
added QmSGmPkgZGQhJ9uunAVmMRB5dULMfEPRZ3ob4qBJvijhwF ipfs
 6 B / 6 B [====================================================================================================] 100.00%
 会生成两个 CID 信息


cat 的时候  不能 cat  第二个 CID   报错说是 tag dag node is directory。
不知道 为何  后面在探讨一下

```



### get

```go
//获取 文件 
$ ipfs get -h                                                                                                 [14:40:16]
USAGE
  ipfs get <ipfs-path> - Download IPFS objects.
  //       ipfs  路径    -d 
  ipfs get [--output=<output> | -o] [--archive | -a] [--compress | -C] [--compression-level=<compression-level> | -l]
           [--] <ipfs-path>

  Stores to disk the data contained an IPFS or IPNS object(s) at the given path.

  By default, the output will be stored at './<ipfs-path>', but an alternate
  path can be specified with '--output=<path>' or '-o=<path>'.

  To output a TAR archive instead of unpacked files, use '--archive' or '-a'.

  To compress the output with GZIP compression, use '--compress' or '-C'. You
  may also specify the level of compression by specifying '-l=<1-9>'.

  For more information about each command, use:
  'ipfs get <subcmd> --help'
  
QmSGmPkgZGQhJ9uunAVmMRB5dULMfEPRZ3ob4qBJvijhwF


//例如
$ ipfs add -r ipfs                                                                                            [15:00:41]
added QmPje32NWnjRYWagtB61VF1QhzWEaCvPwWMvUeouUbGwe7 ipfs/2.txt
added QmPje32NWnjRYWagtB61VF1QhzWEaCvPwWMvUeouUbGwe7 ipfs/3.txt
added QmUtzmZ6Aot5TTMdoDUTyiTm1EViJVJrU7dGd79KjSEiJA ipfs/4.txt
added QmZyTztEF1UfJ1Qw8HzaWpQcv98ogu4kfnKteUwav2gg6T ipfs/i.txt
added QmdhZ2yN2DZ6pyE7Xumb9oR68yM8ckL4dX1MHbddtDxhf9 ipfs
 818.64 KiB / 818.64 KiB [======================================================================================] 100.00%

这是添加了一个文件夹
如果  文件夹下面的内容  都是小于  256kb 的话  都不会分片  
大于  256kb 分片 

可以通过 ipfs ls CID  查看 link的 cid信息


```



### ls

```go
$ ipfs ls -h                                                                                                                         [15:03:40]
USAGE
  ipfs ls <ipfs-path>... - List directory contents for Unix filesystem objects.

  ipfs ls [--headers | -v] [--resolve-type=false] [--size=false] [--stream | -s] [--] <ipfs-path>...

  Displays the contents of an IPFS or IPNS object(s) at the given path, with
  the following format:

    <link base58 hash> <link size in bytes> <link name>

  The JSON output contains type information.

  For more information about each command, use:
  'ipfs ls <subcmd> --help'
  
  例如：
$ ipfs ls -v=true -s=true QmUtzmZ6Aot5TTMdoDUTyiTm1EViJVJrU7dGd79KjSEiJA                                                             [15:05:08]
Hash                                           Size      Name
QmRBFEuvwzjLshfUVWJL3JWU428xjUMLFT2Tc8A27mNAcz 262144
QmSAsJRfRiqMD6f41Pu39PJgodbr32qyUn2iKR8FgUXk3W 262144
QmVPk7rMbiD8duqkA49zeop2JWAjtysEirXAqG84BXQwjK 262144
QmXkdWnqzgmoU1kvFpoR4HeMuAvAH4dAiAyU7UxqwoJ6wV 51846



```





## DATA STRUCTURE COMMANDS

### block

```go
ipfs block -h
USAGE
  ipfs block - Interact with raw IPFS blocks.
			
  ipfs block

  'ipfs block' is a plumbing command used to manipulate raw IPFS blocks.
  Reads from stdin or writes to stdout, and <key> is a base58 encoded
  multihash.

  // 'ipfs block'是一个用于操作原始ipfs块的管道命令。
  // 从stdin读取或写入到stdout， <key>是base58编码的
  // multihash。

SUBCOMMANDS
  ipfs block get <key>     - Get a raw IPFS block.          //获取原始IPFS块  
  ipfs block put <data>... - Store input as an IPFS block.  //将输入存储为IPFS块
  ipfs block rm <hash>...  - Remove IPFS block(s).          //移除IPFS块
  ipfs block stat <key>    - Print information of a raw IPFS block.//打印原始IPFS 块信息

  For more information about each command, use:
  'ipfs block <subcmd> --help
```





## ipfs

### Ipfs 命令行简介

ipfs是基于默克尔有向无环图（merkle dag）的全球性p2p文件系统。

#### 命令行

```
ipfs [--config=<config> | -c] [--debug=<debug> | -D] 
     [--help=<help>] [-h=<h>] [--local=<local> | -L] 
     [--api=<api>] <command> ...
```

#### 命令行选项

```
-c, --config string - 配置文件路径
-D, --debug  bool   - 开启调试模式，默认值：false
--help       bool   - 是否显示完整的命令帮助文档，默认值：false
-h           bool   - 显示简明版的命令帮助文档，默认值：false
-L, --local  bool   - 仅在本地执行命令，不使用后台进程。默认值：false
--api        string - 使用指定的API实例，默认值：`/ip4/127.0.0.1/tcp/5001`
--cid-base          - 字符串-多基编码用于版本1的cid输出。
--upgrade-cidv0-in-output  bool   - 将输出中的cid从版本0升级到版本1。
--enc, --encoding          string - 输出应该使用的编码类型(json、xml或text)。
Default: text.                    -默认  text
  --stream-channels          bool   - 流管道输出
  --timeout                  string - 设置该命令的全局超时时间。
```



#### 基本子命令

```
init          初始化ipfs本地配置
add <path>    将指定文件添加到IPFS
cat <ref>     显示指定的IPFS对象数据
get <ref>     下载指定的IPFS对象
ls <ref>      列表显示指定对象的链接
refs <ref>    列表显示指定对象的链接哈希
```



#### 数据结构子命令

```
block         操作数据仓中的裸块
object        操作有向图中的裸节点
files         以unix文件系统方式操作IPFS对象
dag           操作IPLD文档，目前处于实验阶段
```



#### 高级子命令

```
daemon        启动后台服务进程
mount         挂接只读IPFS
resolve       名称解析
name          发布、解析IPNS名称
key           创建、列表IPNS名称键值对
dns           解析DNS链接
pin           在本地存储中固定IPFS对象
repo          操作IPFS仓库
stats         各种运营统计
p2p           挂载libp2p 流
filestore     管理文件仓，目前处于实验阶段
```



#### 网络子命令

```
id            显示IPFS节点信息
bootstrap     添加、删除启动节点
swarm         管理p2p网络的连接
dht           查询分布哈希表中的值或节点信息
ping          检测连接延时
diag          打印诊断信息
```



#### 工具子命令

```
config        管理配置信息
version       显示ipfs版本信息
update        下载并应用go-ipfs更新
commands      列表显示全部可用命令
cid           转换和发现cid的属性 
log           管理和显示正在运行的守护进程的日志信息
```

使用`ipfs <command> --help`来了解特定命令的详细帮助信息。



#### 本地仓库地址

```go
ipfs使用本地文件系统中的仓库存储内容。默认情况下，本地仓库位于 ~/.ipfs。你可以设置IPFS_PATH环境变量来定义本地仓库的位置：

eg:
export IPFS_PATH=/home/server

```





#### 命令行退出状态

```shell
命令行的退出码
如下：

- 0：执行成功
- 1：执行失败
```



## Ipfs add

### 添加 文件 或者 文件夹

#### 命令行

```
USAGE
  ipfs add <path>... - Add a file or directory to ipfs.

  ipfs add [--recursive | -r] [--dereference-args] [--stdin-name=<stdin-name>] [--hidden | -H] [--ignore=<ignore>]...
           [--ignore-rules-path=<ignore-rules-path>] [--quiet | -q] [--quieter | -Q] [--silent] [--progress | -p] [--trickle | -t] [--only-hash | -n]
           [--wrap-with-directory | -w] [--chunker=<chunker> | -s] [--pin=false] [--raw-leaves] [--nocopy] [--fscache] [--cid-version=<cid-version>]
           [--hash=<hash>] [--inline] [--inline-limit=<inline-limit>] [--] <path>...

  Adds contents of <path> to ipfs. Use -r to add directories (recursively).

  For more information about each command, use:
  'ipfs add <subcmd> --help'
```

`<path>...` - 要添加到ipfs中的文件的路径

#### 选项

```go
-r,         --recursive           bool   - 递归添加目录内容，默认值：false
-q,         --quiet               bool   - 安静模式，执行过程中输出显示尽可能少的信息
-Q,         --quieter             bool   - 更安静模式，仅输出最终的结果哈希值
--silent                          bool   - 静默模式，不输出任何信息
-p,         --progress            bool   - 流式输出过程数据
-t,         --trickle             bool   - 使用trickle-dag格式进行有向图生成
-n,         --only-hash           bool   - Only chunk and hash - do not write to disk.
-w,         --wrap-with-directory bool   - 使用目录对象包装文件
-H,         --hidden              bool   - 包含隐藏文件，仅在进行递归添加时有效
-s,         --chunker             string - 使用的分块算法
--pin                             bool   - 添加时固定对象，默认值：true
--raw-leaves                      bool   - 叶节点使用裸块，实验特性
--nocopy                          bool   - 使用filestore添加文件，实验特性
--fscache                         bool   - 为已有块检查filestore，实验特性
=======
之前版本
===============================================================================================================
之后版本
=======
-r, --recursive               bool     - 递归添加目录内容，默认值：false
  --dereference-args          bool     - 参数中提供的符号链接将被解引用。
  --stdin-name                string   - 如果文件源是stdin，则指定一个名称。
  -H, --hidden                bool     - 包含隐藏的文件。只对递归添加生效。
  --ignore                    array    - 一个规则(.gitignore-stype)，定义应该忽略哪个文件(可变的，实验性的)。
  --ignore-rules-path        string    - 使用.gitignore风格忽略规则的文件路径(实验性)。
  -q, --quiet                bool      - 安静模式，执行过程中输出显示尽可能少的信息
  -Q, --quieter              bool      - 更安静模式，仅输出最终的结果哈希值
  --silent                   bool      - 静默模式，不输出任何信息
  -p, --progress             bool      - 流式输出过程数据
  -t, --trickle              bool      - 使用trickle-dag格式进行有向图生成
  -n, --only-hash            bool      - Only chunk and hash - do not write to disk.
  -w, --wrap-with-directory  bool      - 使用目录对象包装文件
  -s, --chunker              string    - 使用的分块算法
Default: size-262144.                  - 默认的是 256kb   256 *1024
  --pin                      bool      - 添加时固定对象，默认值：true
  --raw-leaves               bool      - 叶节点使用裸块，实验特性
  --nocopy                   bool      - 使用filestore添加文件，实验特性
  --fscache                  bool      - 为已有块检查filestore，实验特性
  --cid-version              int       - CID 版本. 默认为0，除非传递了一个依赖于CIDv1的选项.实验特性
  --hash                     string    - 要使用的哈希函数。如果不是sha2-256，则表示CIDv1
Default: sha2-256.                     - 默认使用 sha2-256
  --inline                   bool      - 将小块内联到cid中. 实验特性
  --inline-limit             int       - 内联的最大块大小。 实验特性 默认: 32.
```

#### 说明

将`<path>`的内容添加到ipfs中。使用`-r`来添加目录。目录内容的添加 是递归进行的，以便生成ipfs的默克尔DAG图。

包装选项`-w`将文件包装到一个目录中，该目录仅包含已经添加的文件，这意味着 文件将保留其文件名。

如果守护进程没有开启运行，它只是在本地添加，如果等会开启本地进程，它会在几秒之后发布。

例如：

```
> ipfs add example.jpg
added QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH example.jpg
> ipfs add example.jpg -w
added QmbFMke1KXqnYyBBWxB74N4c5SBnJMVAiMNRcGu6x1AwQH example.jpg
added QmaG4FuMqEBnQNn3C8XJ5bpW8kLs7zq2ZXgHptJHbKDDVx
```

你可以在网关上查看你的文件引用，eg：

/ipfs/<cid>



Chunker Option  选项   -s 制定了分块策略

如何将文件分成块，具有相同内容的块能被处理，不同的块策略会有不同的结果，同一个散列，默认值是 256kb，或者可以自己定义使用分块大小。

最后，关于哈希注释的警告注意，虽然不能保证，但是添加相同的文件，输出的哈希值总是相同的。



## ipfs cat

`ipfs cat <ipfs-path>...`命令用来显示IPFS对象数据.

#### 命令行

```go
USAGE
  ipfs cat <ipfs-path>...         - 显示文件信息内容

SYNOPSIS
  ipfs cat [--offset=<offset> | -o] [--length=<length> | -l] [--] <ipfs-path>...

ARGUMENTS

  <ipfs-path>...                 - 输出路径

OPTIONS

  -o, --offset  int64           - 从读开始偏移的字节
  -l, --length  int64           - 读出的最大长度

DESCRIPTION

  Displays the data contained by an IPFS or IPNS object(s) at the given path.
  显示指定路径下的IPFS或IPNS对象所包含的数据。
```

#### 说明

```
显示指定路径下的IPFS或IPNS对象所包含的数据。
```

#### 使用

```
echo Hello,world >hello |ipfs add hello
ipfs cat QmaZMLejnjNKex6Nrs2RGLC8n7NvWQP8RFPn2dLs2XviYb

Hello,world

ipfs cat -o=0 -l=5 QmaZMLejnjNKex6Nrs2RGLC8n7NvWQP8RFPn2dLs2XviYb

hello

```



## Ipfs commands

#### 命令行

```go
USAGE
  ipfs commands - List all available commands.            -显示所有可用的命令
               
  ipfs commands [--flags | -f]                            -显示后面的子命令

  Lists all available commands (and subcommands) and exits. -列出所有可用的命令(和子命令)和退出

  For more information about each command, use:
  'ipfs commands <subcmd> --help'
```



#### 说明

```go
$ ipfs commands |wc -l                                                                                                                          [14:10:21]
     168
    
168 个 命令
```



## Ipfs daemon

### 命令行

```go
  ipfs daemon - Run a network-connected IPFS node.

SYNOPSIS
  ipfs daemon [--init] [--init-config=<init-config>] [--init-profile=<init-profile>] [--routing=<routing>] [--mount] [--writable]
              [--mount-ipfs=<mount-ipfs>] [--mount-ipns=<mount-ipns>] [--unrestricted-api] [--disable-transport-encryption] [--enable-gc]
              [--manage-fdlimit=false] [--migrate] [--enable-pubsub-experiment] [--enable-namesys-pubsub] [--enable-mplex-experiment]

OPTIONS

  --init                          bool   - 是否使用默认设置自动初始化ipfs，默认值：false
  --init-config                   string - Path to existing configuration file to be loaded during --init.
  --init-profile                  string - Configuration profiles to apply for --init. See ipfs init --help for more.
  --routing                       string - 路由选项，默认值：dht
  --mount                         bool   - 是否将IPFS挂载到文件系统，默认值：false
  --writable                      bool   - 是否允许使用`POST/PUT/DELETE`修改对象，默认值： false.
  --mount-ipfs                    string - 当使用--mount选项时IPFS的挂接点，默认值采用配置文件中的设置
  
  --mount-ipns                    string - 当使用--mount选项时IPNS的挂接点，默认值采用配置文件中的设置
  --unrestricted-api              bool   - 是否允许API访问未列出的哈希，默认值：false
  --disable-transport-encryption  bool   - 是否进制传输层加密，默认值：false。当调试协议时可开启该选项
  --enable-gc                     bool   - 是否启用自动定时仓库垃圾回收，默认值：false
  --manage-fdlimit                bool   - 是否按需自动提高文件描述符上限，默认值：false
  --migrate                       bool   - 是否离线运行，即不连接到网络，仅提供本地API，默认值：false
  --enable-pubsub-experiment      bool   - true对应于mirage提示时输入yes，false对应于输入no
  --enable-namesys-pubsub         bool   - 是否启用发布订阅（pubsub）特性，该特性目前尚处于实验阶段
  --enable-mplex-experiment       bool   - 是否启用`go-multiplex`流多路处理器，默认值：true
```

#### 说明

```
服务进程将在指定的端口监听网络连接。使用ipfs config Addresses 命令修改默认端口。

例如，修改网关监听端口：

ipfs config Addresses.Gateway /ip4/127.0.0.1/tcp/8082
同样的方式修改API地址：

ipfs config Addresses.API /ip4/127.0.0.1/tcp/5002
在修改地址后，确保重新启动服务进程以便生效。

默认情况下，网络仅在本地可以访问，如果希望允许其他计算机访问，可以 使用地址0.0.0.0。例如：

ipfs config Addresses.Gateway /ip4/0.0.0.0/tcp/8080
当开放API访问时请千万小心，这存在一定的安全风险，因为任何人都可以 远程控制你的节点。如果你希望远程控制节点，请使用防火墙、授权代理 或其他服务来保护该API访问地址。
```

#### HTTP头

ipfs支持向API和网关传入任意HTTP头信息。你可以使用`API.HTTPHeaders`和 `Gateway.HTTPHeaders`配置项进行配置。例如：

```
ipfs config --json API.HTTPHeaders.X-Special-Header '["so special :)"]'
ipfs config --json Gateway.HTTPHeaders.X-Special-Header '["so special :)"]'
```

需要指出的是，值应当是字符串数组，因为HTTP头可以包含多个值，而且这样也 方面传给其他的库。

对于API而言，可以同样的方式为其设置CORS头：

```
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Origin '["example.com"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Methods '["PUT", "GET", "POST"]'
ipfs config --json API.HTTPHeaders.Access-Control-Allow-Credentials '["true"]'
```

#### 停止服务

要停止服务进程，发送`SIGINT`信号即可，例如，使用`Ctrl-C`组合键。也可以发送 `SIGTERM`信号，例如，使用`kill`命令。服务进程需要稍等一下以便优雅退出，但是你 可以继续发送一次信息来强制服务进程立刻退出。

#### IPFS_PATH环境变量

ipfs使用本地文件系统建立本地仓库。默认情况下，本地仓库的目录是`~/.ipfs`， 可以设置`IPFS_PATH`环境变量来自定义本地仓库路径：

```
export IPFS_PATH=/path/to/ipfsrepo
```

#### 路由

默认情况下，ipfs使用分布式哈希表（DHT）进行内容的路由。目前有一个尚处于 试验阶段的替代方案，使用纯客户端模式来操作分布式哈希表，可以在启动服务 进程时，使用如下命令启动这一替代路由方案：

```
ipfs daemon --routing=dhtclient
```

该选项在实验阶段结束后将转变为一个配置选项。

#### 弃用通知

ipfs之前使用过环境变量`API_ORIGIN`：

```
export API_ORIGIN="http://localhost:8888/"
```

该环境变量已经被弃用。在当前版本中还可以使用，但在将来的版本中将会删除对 此环境变量的支持。请使用前述HTTP头信息设置方法来取代该环境变量。













### 其他知识点

#### 1.ED25519 算法

```go
网址 ：  https://zhuanlan.zhihu.com/p/110413836
ED25519 算法 生成秘钥对
常见的 SSH 登录密钥使用 RSA 算法。RSA 经典且可靠，但性能不够理想。

只要你的服务器上 OpenSSH 版本大于 6.5（2014 年的古早版本），就可以利用 Ed25519 算法生成的密钥对，减少你的登录时间。如果你使用 SSH 访问 Git，那么就更值得一试。

Ed25519 的安全性在 RSA 2048 与 RSA 4096 之间，且性能在数十倍以上。

顺便回顾一下 ssh 的 rsa 生成秘钥对算法。

生成秘钥：

mkdir .ssh     //新建  .ssh  文件夹目录

ssh-keygen -t ed25519 -f my_dir -C "me@.github.com"


解释： ssh-keygen 的命令含义

其中：

[-t rsa] 表示使用 RSA 算法。
[-b 4096] 表示 RSA 密钥长度 4096 bits （默认 2048 bits）。Ed25519 算法不需要指定。
[-f my_id] 表示在【当前工作目录】下生成一个私钥文件 my_id （同时也会生成一个公钥文件 my_id.pub）。
[-C "email@example.com"] 表示在公钥文件中添加注释，即为这个公钥“起个别名”（不是 id，可以更改）。
在敲下该命令后，会提示输入 passphrase，即为私钥添加一个“解锁口令”。

解释：最佳实践
私钥必须要有 passphrase。如果私钥文件遗失，没有 passphrase 也无法解锁（只能暴力破解）。不要偷懒，passphrase 一定要加。
一对密钥只对应一个 Git 服务。一对密钥通吃各 Git 服务不太明智。
严格来讲，你应该在不同的机器上用不同的密钥，出了问题好排查处理。但实际上复杂的管理反而更容易让人犯错，选择你能 hold 住的方式更为重要。
```

