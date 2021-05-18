## Cli

```go
  ipfs - Global p2p merkle-dag filesystem.  //全局 p2p 默克尔有向无环图 文件系统
	
  ipfs [<flags>] <command> [<arg>] ...

SUBCOMMANDS // 子命令   
  BASIC COMMANDS //基本命令
    init          Initialize local IPFS configuration  //初始化 本地 IPFS 配置
    add <path>    Add a file to IPFS                   //添加一个文件到 IPFS
    cat <ref>     Show IPFS object data								 //展示IPFS目标数据 
    get <ref>     Download IPFS objects								 //下载IPFS目标数据
    ls <ref>      List links from an object						 //目标列表链接
    refs <ref>    List hashes of links from an object  //目标列表链接哈希

  DATA STRUCTURE COMMANDS // 数据结构命令
    dag           Interact with IPLD DAG nodes         //与 LPLD DAG 交互
    files         Interact with files as if they were a unix filesystem //与文件交互
    block         Interact with raw blocks in the datastore//与数据存储中的原始块交互

  ADVANCED COMMANDS       //高级命令
    daemon        Start a long-running daemon process  // 启动一个长期运行的守护进程
    mount         Mount an IPFS read-only mount point  // 挂载IPFS只读挂载点
    resolve       Resolve any type of name             // 解析任何类型的名称
    name          Publish and resolve IPNS names       // 发布和解析IPNS名称
    key           Create and list IPNS name keypairs   // 创建并列出IPNS名称密钥对
    dns           Resolve DNS links                    // 解决域名链接
    pin           Pin objects to local storage         // 将对象引到本地存储
    repo          Manipulate the IPFS repository       // 操作IPFS存储库
    stats         Various operational stats            // 各种运营数据
    p2p           Libp2p stream mounting               // Libp2p流
    filestore     Manage the filestore (experimental)  // 管理文件存储库(实验性的
 
  NETWORK COMMANDS      //网络命令
    id            Show info about IPFS peers           // 显示IPFS节点信息
    bootstrap     Add or remove bootstrap peers        // 添加或者移除引导节点
    swarm         Manage connections to the p2p network// 管理p2p 网络连接
    dht           Query the DHT for values or peers    // 查询DHT的值或对等体
    ping          Measure the latency of a connection  // 测量连接的延迟时间
    diag          Print diagnostics                    // 打印诊断

  TOOL COMMANDS
    config        Manage configuration                 // 管理配置
    version       Show IPFS version information        // IPFS信息版本
    update        Download and apply go-ipfs updates   // 下载和申请ipfs更新
    commands      List all available commands          // 所有可用命令的列表
    cid           Convert and discover properties of CIDs//转换和发现cid的属性
    log           Manage and show logs of running daemon //管理和显示运行daemon的日志

  Use 'ipfs <command> --help' to learn more about each command.

  ipfs uses a repository in the local file system. By default, the repo is located at
  ~/.ipfs. To change the repo location, set the $IPFS_PATH environment variable:
//Ipfs在本地文件系统中使用存储库。默认情况下，是在 ~/.ipfs  你想改变本地，设置环境变量 $IPFS_PATH
    export IPFS_PATH=/path/to/ipfsrepo
```

