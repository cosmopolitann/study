

# Orbit-db

[toc]



<img src="https://github.com/orbitdb/orbit-db/blob/main/images/orbit_db_logo_color.png?raw=true" style="zoom:50%;" />



![](https://img.shields.io/badge/github-on%20gitter-green)![](https://img.shields.io/badge/matrix-orbit--db%3Amatrix.org-orange)![](https://img.shields.io/badge/node-%3E%3D10.0.0-green)![](https://img.shields.io/badge/npm%20package-0.26.1-green)![](https://img.shields.io/badge/curckeci-passing-green)



## Introduce

OrbitDB is a **serverless, distributed, peer-to-peer database**. OrbitDB uses [IPFS](https://ipfs.io/) as its data storage and [IPFS Pubsub](https://github.com/ipfs/go-ipfs/blob/master/core/commands/pubsub.go#L23) to automatically sync databases with peers. It's an eventually consistent database that uses [CRDTs](https://en.wikipedia.org/wiki/Conflict-free_replicated_data_type) for conflict-free database merges making OrbitDB an excellent choice for decentralized apps (dApps), blockchain applications and offline-first web applications.

---

OrbitDB 是一个 无服务器，分布式的，点对点的数据库。OrbitDB 使用 IPFS 作为他的数据存储，并且使用 IPFS 的 pubsub 去 自动的和节点同步数据库。它是一个一致性的数据库，使用 CRDTs 进行无冲突合并数据库使 OrbitDB 成为一个去中心化的 Dapps，区块链应用和离线web应用的一个最佳选择。

---



### Test it live 

at  Live demo ，Live demo2 or P2P TodoMVC app.

在线测试： Live demo ，Live demo2 or P2P TodoMVC app.

```
下面是链接地址，github。
// 1.
Live demo https://ipfs.io/ipfs/QmUsoSkGzUQnCgzfjL549KKf29m5EMYky3Y6gQp5HptLTG/
// 2.
Live demo2 https://ipfs.io/ipfs/QmasHFRj6unJ3nSmtPn97tWDaQWEZw3W9Eh3gUgZktuZDZ/
// 3.
P2P TodoMVC app  https://ipfs.io/ipfs/QmVWQMLUM3o4ZFbLtLMS1PMLfodeEeBkBPR2a2R3hqQ337/
```

---



### Database Type.

OrbitDB provides various types of databases for different data models and use cases:

- **[log](https://github.com/orbitdb/orbit-db/blob/master/API.md#orbitdblognameaddress)**: an immutable (append-only) log with traversable history. Useful for *"latest N"* use cases or as a message queue.
- **[feed](https://github.com/orbitdb/orbit-db/blob/master/API.md#orbitdbfeednameaddress)**: a mutable log with traversable history. Entries can be added and removed. Useful for *"shopping cart"* type of use cases, or for example as a feed of blog posts or "tweets".
- **[keyvalue](https://github.com/orbitdb/orbit-db/blob/master/API.md#orbitdbkeyvaluenameaddress)**: a key-value database just like your favourite key-value database.
- **[docs](https://github.com/orbitdb/orbit-db/blob/master/API.md#orbitdbdocsnameaddress-options)**: a document database to which JSON documents can be stored and indexed by a specified key. Useful for building search indices or version controlling documents and data.
- **[counter](https://github.com/orbitdb/orbit-db/blob/master/API.md#orbitdbcounternameaddress)**: Useful for counting events separate from log/feed data.



---

OrbitDB为不同的数据模型和用例提供了多种类型的数据库:

- log                    一个可遍历 历史  不可变的  log （仅追加）

- feed                 一个可遍历 历史 记录 可变的 log。词目(entries)（数据）可以 添加, 移除. 可以使用在 购物车 案例上类型上，或者 文章博客 tweets的 提要

- keyvalue        键值数据库  就像你喜欢键值对数据库 例如 mongo  redis 数据库。

- docs               可以将JSON文档存储到其中并按指定键建立索引的文档数据库。用于构建搜索索引或版本控制文档和数据。

- counter         对于独立于日志/提要数据的事件计数很有用

  

---

All databases are [implemented](https://github.com/orbitdb/orbit-db-store) on top of [ipfs-log](https://github.com/orbitdb/ipfs-log), an immutable, operation-based  conflict-free(无冲突的) replicated data structure (CRDT) for distributed systems. If none of the OrbitDB database types match your needs and/or you need case-specific functionality, you can easily [implement and use a custom database store](https://github.com/orbitdb/orbit-db/blob/master/GUIDE.md#custom-stores) of your own.

---

所有数据库都是在 IPFS-log 之上实现的， 它是 一种 用于 分布式的系统的不可变，基于操作，不可变的数据结构（CRDTs）.如果没有任何一种 OrbitDB 数据类型 满足 符合 你的 需求 ，你可以 很容易 地 实现 使用你自己的自定义的数据库存储。

### Project status & support 

项目状态 和支持

- Status: **in active development**   在积极开发
- Compatible with **js-ipfs versions <= 0.52** and **go-ipfs versions <= 0.6.0**   兼容**js-ipfs版本<= 0.52**和**go-ipfs版本<= 0.6.0**

---



***NOTE!*** *OrbitDB is **alpha-stage** software. It means OrbitDB hasn't been security audited and programming APIs and data formats can still change. We encourage you to [reach out to the maintainers](https://gitter.im/orbitdb/Lobby) if you plan to use OrbitDB in mission critical systems.*

This is the Javascript implementation and it works both in **Browsers** and **Node.js** with support for Linux, OS X, and Windows. LTS versions (even numbered versions 8, 10, etc) are supported.

To use with older versions of Node.js, we provide an ES5-compatible build through the npm package, located in `dist/es5/` when installed through npm.

---

提示笔记 !（注意 ！！!）  OrbitDB是 开发阶段 的 软件。这意味着OrbitDB还没有经过安全审计，编程api和数据格式仍然可以改变。如果你计划在关键任务系统中使用OrbitDB，我们鼓励你联系维护人员。

这是Javascript实现，它工作在浏览器和Node.js支持Linux, OS X，和Windows。支持LTS版本(甚至编号版本8、10等)。

为了与旧版本的Node.js一起使用，我们通过npm包提供了一个与es5兼容的构建，当通过npm安装时，它位于dist/es5/中。

### Table of Contents

- Usage
  - [Database browser UI](https://github.com/orbitdb/orbit-db#database-browser-ui)
  - [Module with IPFS Instance](https://github.com/orbitdb/orbit-db#module-with-ipfs-instance)
  - [Module with IPFS Daemon](https://github.com/orbitdb/orbit-db#module-with-ipfs-daemon)
- [API](https://github.com/orbitdb/orbit-db#api)
- Examples
  - [Install dependencies](https://github.com/orbitdb/orbit-db#install-dependencies)
  - [Browser example](https://github.com/orbitdb/orbit-db#browser-example)
  - [Node.js example](https://github.com/orbitdb/orbit-db#nodejs-example)
  - [Workshop](https://github.com/orbitdb/orbit-db#workshop)
- Packages
  - [OrbitDB Store Packages](https://github.com/orbitdb/orbit-db#orbitdb-store-packages)
- Development
  - [Run Tests](https://github.com/orbitdb/orbit-db#run-tests)
  - [Build](https://github.com/orbitdb/orbit-db#build)
  - [Benchmark](https://github.com/orbitdb/orbit-db#benchmark)
  - [Logging](https://github.com/orbitdb/orbit-db#logging)
- Frequently Asked Questions
  - [Are there implementations in other languages?](https://github.com/orbitdb/orbit-db#are-there-implementations-in-other-languages)
- [Contributing](https://github.com/orbitdb/orbit-db#contributing)
- [Sponsors](https://github.com/orbitdb/orbit-db#sponsors)
- [License](https://github.com/orbitdb/orbit-db#license)





---

### Usage

Read the **[GETTING STARTED](https://github.com/orbitdb/orbit-db/blob/master/GUIDE.md)** guide for a quick tutorial on how to use OrbitDB.

For a more in-depth tutorial and exploration of OrbitDB's architecture, please check out the **[OrbitDB Field Manual](https://github.com/orbitdb/field-manual)**.

### Database browser UI

OrbitDB databases can easily be managed using a web UI, see **[OrbitDB Control Center](https://github.com/orbitdb/orbit-db-control-center)**.

Install and run it locally:

---

使用说明

快速阅读这个指南 怎么去使用 orbitdb。

关于OrbitDB架构的更深入的教程和探索，请查阅**[OrbitDB Field Manual](https://github.com/orbitdb/field-manual)**.

数据库浏览器UI

OrbitDB数据库可以通过web UI轻松管理，请看  **[OrbitDB Control Center](https://github.com/orbitdb/orbit-db-control-center)**.

在本地安装 并运行它。



安装，运行。

```
git clone https://github.com/orbitdb/orbit-db-control-center.git
cd orbit-db-control-center/
npm i && npm start
```

![](https://raw.githubusercontent.com/orbitdb/orbit-db-control-center/master/screenshot1.png) 



![](https://raw.githubusercontent.com/orbitdb/orbit-db-control-center/master/screenshot2.png)



### Module with IPFS Instance

If you're using `orbit-db` to develop **browser** or **Node.js** applications, use it as a module with the javascript instance of IPFS

Install dependencies:

---

IPFS 实例模块

如果你正在使用 orbit-db 去开发浏览器 和 node.js 应用。使用它作为 IPFS 的 JavaScript 实例模块。 

安装依赖：

```
npm install orbit-db ipfs
```

```js
const IPFS = require('ipfs')
const OrbitDB = require('orbit-db')

// For js-ipfs >= 0.38

// Create IPFS instance
const initIPFSInstance = async () => {
  return await IPFS.create({ repo: "./path-for-js-ipfs-repo" });
};

initIPFSInstance().then(async ipfs => {
  const orbitdb = await OrbitDB.createInstance(ipfs);

  // Create / Open a database
  const db = await orbitdb.log("hello");
  await db.load();

  // Listen for updates from peers
  db.events.on("replicated", address => {
    console.log(db.iterator({ limit: -1 }).collect());
  });

  // Add an entry
  const hash = await db.add("world");
  console.log(hash);

  // Query
  const result = db.iterator({ limit: -1 }).collect();
  console.log(JSON.stringify(result, null, 2));
});


// For js-ipfs < 0.38

// Create IPFS instance
const ipfsOptions = {
    EXPERIMENTAL: {
      pubsub: true
    }
  };

ipfs = new IPFS(ipfsOptions);

initIPFSInstance().then(ipfs => {
  ipfs.on("error", e => console.error(e));
  ipfs.on("ready", async () => {
    const orbitdb = await OrbitDB.createInstance(ipfs);

    // Create / Open a database
    const db = await orbitdb.log("hello");
    await db.load();

    // Listen for updates from peers
    db.events.on("replicated", address => {
      console.log(db.iterator({ limit: -1 }).collect());
    });

    // Add an entry
    const hash = await db.add("world");
    console.log(hash);

    // Query
    const result = db.iterator({ limit: -1 }).collect();
    console.log(JSON.stringify(result, null, 2));
  });
});
```

### Module with IPFS Daemon

Alternatively, you can use [ipfs-http-client](https://www.npmjs.com/package/ipfs-http-client) to use `orbit-db` with a locally running IPFS daemon. Use this method if you're using `orbitd-db` to develop **backend** or **desktop** applications, eg. with [Electron](https://electron.atom.io/).

Install dependencies:

---

IPFS 进程模块

或者 你可以 使用 ipfs-http-client 去使用 orbit-db 在本地运行 IPFS 进程。如果你正在使用 orbit-db 去开 后端在 或者 在桌面开发，请使用此方法。

安装依赖：

---



```js
npm install orbit-db ipfs-http-client
```



```js
const IpfsClient = require('ipfs-http-client')
const OrbitDB = require('orbit-db')

const ipfs = IpfsClient('localhost', '5001')

const orbitdb = await OrbitDB.createInstance(ipfs)
const db = await orbitdb.log('hello')
// Do something with your db.
// Of course, you may want to wrap these in an async function.
```



### API

See [API.md](https://github.com/orbitdb/orbit-db/blob/master/API.md) for the full documentation.

完整文档请看 API.md

```
https://github.com/orbitdb/orbit-db/blob/master/API.md
```

### Examples  

案例

### Install dependencies

安装依赖。

```
git clone https://github.com/orbitdb/orbit-db.git
cd orbit-db
npm install
```

Some dependencies depend on native addon modules, so you'll also need to meet [node-gyp's](https://github.com/nodejs/node-gyp#installation) installation prerequisites. Therefore, Linux users may need to

有些依赖项依赖于本地插件模块，所以您还需要满足 node-gyp's 安装先决条件。因此，Linux用户可能需要.

```
make clean-dependencies && make deps
```

to redo the local package-lock.json with working native dependencies.

重做本地包锁，使用本机依赖项的Json。

### Browser example

```
npm run build
npm run examples:browser
```

Using Webpack:

```
npm run build
npm run examples:browser-webpack
```



Webpack

本质上，**webpack** 是一个用于现代 JavaScript 应用程序的 *静态模块打包工具*。当 webpack 处理应用程序时，它会在内部构建一个依赖图 ，此依赖图对应映射到项目所需的每个模块，并生成一个或多个 *bundle*。

```
https://webpack.docschina.org/concepts/
```

### Node.js example

```
npm run examples:node
```

**Eventlog** 

系统日志纪录服务 -- 事件日志

See the code in [examples/eventlog.js](https://github.com/orbitdb/orbit-db/blob/master/examples/eventlog.js) and run it with:

在 examples/eventlog.js 看这个代码，使用下面的命令运行它。

```
node examples/eventlog.js
```

### Workshop

讨论。

We have a field manual which has much more detailed examples and a run-through of how to understand OrbitDB, at [orbitdb/field-manual](https://github.com/orbitdb/field-manual). There is also a workshop you can follow, which shows how to build an app, at [orbit-db/web3-workshop](https://github.com/orbitdb/web3-workshop).

More examples at [examples](https://github.com/orbitdb/orbit-db/tree/master/examples).

---

我们在 [orbitdb/field-manual](https://github.com/orbitdb/field-manual).有一个手册 ，里面有更详细的例子和如何理解OrbitDB的演练。你也可以在[orbit-db/web3-workshop](https://github.com/orbitdb/web3-workshop). 上面学习如何构建应用程序。

更多案例在 [examples](https://github.com/orbitdb/orbit-db/tree/master/examples).。

---

### Packages

OrbitDB uses the following modules:

- [ipfs](https://github.com/ipfs/js-ipfs)
- [ipfs-log](https://github.com/orbitdb/ipfs-log)
- [ipfs-pubsub-room](https://github.com/ipfs-shipyard/ipfs-pubsub-room)
- [crdts](https://github.com/orbitdb/crdts)
- [orbit-db-cache](https://github.com/orbitdb/orbit-db-cache)
- [orbit-db-pubsub](https://github.com/orbitdb/orbit-db-pubsub)
- [orbit-db-identity-provider](https://github.com/orbitdb/orbit-db-identity-provider)
- [orbit-db-access-controllers](https://github.com/orbitdb/orbit-db-access-controllers)

使用的库。

OrbitDB使用了下面的模块。

---

### OrbitDB Store Packages

- [orbit-db-store](https://github.com/orbitdb/orbit-db-store)
- [orbit-db-eventstore](https://github.com/orbitdb/orbit-db-eventstore)
- [orbit-db-feedstore](https://github.com/orbitdb/orbit-db-feedstore)
- [orbit-db-kvstore](https://github.com/orbitdb/orbit-db-kvstore)
- [orbit-db-docstore](https://github.com/orbitdb/orbit-db-docstore)
- [orbit-db-counterstore](https://github.com/orbitdb/orbit-db-counterstore)

OrbitDB 存储库。

---

To understand a little bit about the architecture, check out a visualization of the data flow at https://github.com/haadcode/proto2 or a live demo: http://celebdil.benet.ai:8080/ipfs/Qmezm7g8mBpWyuPk6D84CNcfLKJwU6mpXuEN5GJZNkX3XK/.

Community-maintained Typescript typings are available here: https://github.com/orbitdb/orbit-db-types

---

关于想对架构有一点了解，可以在 https://github.com/haadcode/proto2 查看数据流，或者  查看 现场演示案例，地址是这个：

http://celebdil.benet.ai:8080/ipfs/Qmezm7g8mBpWyuPk6D84CNcfLKJwU6mpXuEN5GJZNkX3XK/.

社区维护的 Typescript 类型，是能够在这里找到的。 https://github.com/orbitdb/orbit-db-types

---

### Development

### Run Tests

```
npm test
```

### Build

```
npm run build
```

### Benchmark

基准

```
node benchmarks/benchmark-add.js
```

See [benchmarks/](https://github.com/orbitdb/orbit-db/tree/master/benchmarks) for more benchmarks.

看这个了解更多的基准。

### Logging

To enable OrbitDB's logging output, set a global ENV variable called `LOG` to `debug`,`warn` or `error`:

要启用OrbitDB的日志输出，设置一个名为' LOG '的全局ENV变量为' debug '， ' warn '或' error '。

```js
LOG=debug node <file>
```

---



### Frequently Asked Questions

We have an FAQ! [Go take a look at it](https://github.com/orbitdb/orbit-db/blob/main/FAQ.md). If a question isn't there, open an issue and suggest adding it. We can work on the best answer together.

### Are there implementations in other languages?

Yes! Take a look at these implementations:

- Golang: [berty/go-orbit-db](https://github.com/berty/go-orbit-db)
- Python: [orbitdb/py-orbit-db-http-client](https://github.com/orbitdb/py-orbit-db-http-client)

The best place to find out what is out there and what is being actively worked on is likely by asking in the [Gitter](https://gitter.im/orbitdb/Lobby). If you know of any other repos that ought to be included in this section, please open a PR and add them.

## 

常见问题

我们有一个常见问题解答，点击  [Go take a look at it](https://github.com/orbitdb/orbit-db/blob/main/FAQ.md). ，去看看吧。如果那个上面没有你的问题，打开 issue ，建议把它添加进去，我们一起去找到好的答案和方法，去解决它。

是否还有其他语言的实现？

是的！ 看看这些实现：

Golang  ： [berty/go-orbit-db](https://github.com/berty/go-orbit-db)      

Python:     [orbitdb/py-orbit-db-http-client](https://github.com/orbitdb/py-orbit-db-http-client)

---

### Contributing

**Take a look at our organization-wide [Contributing Guide](https://github.com/orbitdb/welcome/blob/master/contributing.md).** You'll find most of your questions answered there. Some questions may be answered in the [FAQ](https://github.com/orbitdb/orbit-db/blob/main/FAQ.md), as well.

As far as code goes, we would be happy to accept PRs! If you want to work on something, it'd be good to talk beforehand to make sure nobody else is working on it. You can reach us [on Gitter](https://gitter.im/orbitdb/Lobby), or in the [issues section](https://github.com/orbitdb/orbit-db/issues).

If you want to code but don't know where to start, check out the issues labelled ["help wanted"](https://github.com/orbitdb/orbit-db/issues?q=is%3Aopen+is%3Aissue+label%3A"help+wanted"+sort%3Areactions-%2B1-desc).

Please note that we have a [Code of Conduct](https://github.com/orbitdb/orbit-db/blob/main/CODE_OF_CONDUCT.md), and that all activity in the [@orbitdb](https://github.com/orbitdb) organization falls under it. Read it when you get the chance, as being part of this community means that you agree to abide by it. Thanks.

---

贡献

看看我们组织的贡献指南，你在那里会发现更多的问题答案。一些问题也许可以在 FAQ 常见问题解答  同样找到答案。

就代码而言，我们会更乐意接收PRs! 如果 你想做一些事，你最好在那之前先和别人聊聊，确保没有人在正在做这件事，你可以通过 Gitter 或者 在 issues 这个板块联系我们。

如果你想编写代码，但是不知道从哪里开始，可以查看 isssues 板块里面 标记为 需要帮助的问题。

请注意，我们有一个行为准则，orbitdb组织的所有活动 都属于该准则。当你有机会去阅读它，作为即将成为这个社区的一员，意味着你要遵守这个准则。谢谢。

---



```sequence
Orbit -> IPFS: 数据通过IPFS的pubsub传输
Note left of Orbit: 点对点
Note right of IPFS: 点对点
IPFS --> Orbit: 同步去拉取数据
```

