# About

## Content

```


                  IPFS -- Inter-Planetary File system

IPFS is a global, versioned, peer-to-peer filesystem. It combines good ideas
from Git, BitTorrent, Kademlia, SFS, and the Web. It is like a single bit-
torrent swarm, exchanging git objects. IPFS provides an interface as simple
as the HTTP web, but with permanence built-in. You can also mount the world
at /ipfs.

IPFS is a protocol:
- defines a content-addressed file system
- coordinates content delivery
- combines Kademlia + BitTorrent + Git

IPFS is a filesystem:
- has directories and files
- mountable filesystem (via FUSE)

IPFS is a web:
- can be used to view documents like the web
- files accessible via HTTP at `http://ipfs.io/<path>`
- browsers or extensions can learn to use `ipfs://` directly
- hash-addressed content guarantees the authenticity

IPFS is modular:
- connection layer over any network protocol
- routing layer
- uses a routing layer DHT (kademlia/coral)
- uses a path-based naming service
- uses BitTorrent-inspired block exchange

IPFS uses crypto:
- cryptographic-hash content addressing
- block-level deduplication
- file integrity + versioning
- filesystem-level encryption + signing support

IPFS is p2p:
- worldwide peer-to-peer file transfers
- completely decentralized architecture
- **no** central point of failure

IPFS is a CDN:
- add a file to the filesystem locally, and it's now available to the world
- caching-friendly (content-hash naming)
- BitTorrent-based bandwidth distribution

IPFS has a name service:
- IPNS, an SFS inspired name system
- global namespace based on PKI
- serves to build trust chains
- compatible with other NSes
- can map DNS, .onion, .bit, etc to IPNS

```



### 中文

```

IPFS——星际文件系统

IPFS是一个全局的、版本化的、对等的文件系统。它结合了好主意
从Git, BitTorrent, Kademlia, SFS，和Web。就像一个比特
洪流群，交换git对象。IPFS提供了一个同样简单的接口
就像HTTP web一样，但具有内置的永久性。你也可以登上世界
在/ ipf。

IPFS是一个协议:
-定义一个内容寻址的文件系统
-协调内容交付
-结合Kademlia + BitTorrent + Git

IPFS是一个文件系统:
-包含目录和文件
-挂载文件系统(通过FUSE)

IPFS是一个网络:
-可以用来查看文件，如网页
-通过HTTP访问的文件:http://ipfs.io/<path>
-浏览器或扩展可以学习直接使用ipfs://
-哈希地址内容保证了真实性

ipf模块:
-连接层超过任何网络协议
-路由层
-使用路由层DHT(山茱萸/珊瑚)
—使用基于路径的命名服务
-使用bittorrent启发的块交换

ipf使用密码:
-加密散列内容寻址
——块级重复数据删除
-文件完整性+版本控制
-文件系统级加密+签名支持

ipf p2p:
-全球点对点文件传输
-完全去中心化架构
- **没有**失败中心点

IPFS是一个CDN:
-本地添加一个文件到文件系统，它现在对世界可用
-缓存友好(内容哈希命名)
—基于bt的带宽分布

IPFS有一个名称服务:
- IPNS，一个受SFS启发的名称系统
—基于PKI的全局命名空间
-服务于建立信任链
-兼容其他nse
-可以映射DNS， .onion， .bit等到IPNS

```

