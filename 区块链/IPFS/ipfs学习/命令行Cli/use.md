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

Ipfs

























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

