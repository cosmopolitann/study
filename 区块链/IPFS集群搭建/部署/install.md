[toc]

### 搭建IPFS 集群

### 第一步

#### 安装Golang

##### 1.Mac 安装

```go
1.先打开下载安装 Golang的网站
https://studygolang.com/
也可以直接打开 
https://studygolang.com/dl      这个是下载地址 

2.可以使用源码编译，也可以下载压缩版

go1.16.4.darwin-arm64.pkg

下载 https://studygolang.com/dl/golang/go1.16.4.darwin-arm64.pkg
直接 下载 然后安装 就行了

===============配置  GOlang 环境=======================

3.打开 /etc/profile
添加如下 信息：
=======================================
#GOROOT
export GOROOT=/usr/local/go

#GOPATH
export GOPATH=$HOME/go

#PATH root bin
export PATH=$PATH:$GOROOT/bin

#GO PROXY
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
========================================

配置完成  查看 版本

go version

go version go1.16.3 darwin/amd64

```

##### 2.Linux 安装

```go
1.先打开下载安装 Golang的网站
https://studygolang.com/
也可以直接打开 
https://studygolang.com/dl      这个是下载地址 

2.可以使用源码编译，也可以下载压缩版

https://studygolang.com/dl/golang/go1.16.4.linux-amd64.tar.gz

使用 wget https://studygolang.com/dl/golang/go1.16.4.linux-amd64.tar.gz

下载完成后 解压

tar -C /home/server -zxvf go1.16.4.linux-amd64.tar.gz

======================= 配置 GOlang 环境 =================
#GOROOT
export GOROOT=/usr/local/go

#GOPATH
export GOPATH=$HOME/go

#PATH root bin
export PATH=$PATH:$GOROOT/bin

#GO PROXY
export GOPROXY=https://goproxy.cn
=========================================================

配置完成 查看版本信息
go verison
go version go1.16.4 linux/amd64

安装完成
```

### 第二步

#### 安装go-ipfs

##### 1.Mac 安装

```
网址 ： https://dist.ipfs.io/#go-ipfs

直接点击 下载 就行了

```



##### 2.Linxu 安装

```
https://dist.ipfs.io/#go-ipfs

下载 64-bit

wget https://dist.ipfs.io/go-ipfs/v0.8.0/go-ipfs_v0.8.0_linux-amd64.tar.gz
```



### 第三步

#### 安装 ipfs-cluster-service

#### 安装 ipfs-cluster-cli

##### 1.Mac 安装

```
https://dist.ipfs.io/#ipfs-cluster-service
直接下载安装

```



##### 2.Linxu 安装

```
https://dist.ipfs.io/#ipfs-cluster-service
复制连接地址

wget 下载

```



### 第四步

#### 搭建集群节点开始

##### 1.搭建

```go
1. 节点 分为  A  B   C 节点

第一步  分别安装 以上  GOlang  go-ipfs  ipfs-cluster-service  ipfs-cluster-cli 

在  A 节点  初始化 ipfs 
ipfs init

启动 ipfs
ipfs daemon

第二步
在 A 节点上  初始化 cluster-service

ipfs-cluster-service init

在生成的文件夹 中有个 秘钥 。留着 复制给  B C节点

ipfs-cluster-service daemon  
启动

==========
在 B  C 节点上  同样安装  要把秘钥复制过去。

删除 所有引导节点
命令：
ipfs bootstrap rm --all

连接 节点
ipfs bootstrap add /ip4/192.168.11.11/tcp/4001/ipfs/QmSyQFFm5KdB9zoTNQJhPnMs4LYsVmEpY427QYxH2CaFqL

QmSyQFFm5KdB9zoTNQJhPnMs4LYsVmEpY427QYxH2CaFqL 为ipfs init 时生成的节点 ID，也可以通过ipfs id 查看当前节点的 ID。
我们还需要设置环境变量LIBP2P FORCE PNET来强制我们的网络进入私有模式

其他节点启动--bootstrap添加主节点：

ipfs-cluster-service daemon --bootstrap /ip4/192.168.11.11/tcp/9096/ipfs/12D3KooWEGrD9d3n6UJNzAJDyhfTUZNQmQz4k56Hb6TrYEyxyW2F

这里注意下，12D3KooWEGrD9d3n6UJNzAJDyhfTUZNQmQz4k56Hb6TrYEyxyW2F 是 IPFS-Cluster 节点 ID，不是 IPFS 节点 ID，可以通过ipfs-cluster-service id 查看。
可以通过命令查看集群节点状态:

ipfs-cluster-ctl peers ls


```

##### 2.测试

```
ipfs-cluster-ctl add test.txt

ipfs-cluster-ctl status CID

另外一个 就可以查看到文件信息。

```

##### 参考信息

```
https://segmentfault.com/a/1190000021539562
https://cloud.tencent.com/developer/search/article-IPFS
https://cloud.tencent.com/developer/article/1430873
https://www.cnblogs.com/sumingk/articles/9434253.html
```

