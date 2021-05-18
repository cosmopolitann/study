Go安装启动实现的Libp2p的网络栈
The Go implementation of the libp2p Networking Stack.

Table of Contents（内容目录）

Background        背景
Usage             用法
API               API
Examples          案例
Development       发展
Using the go-libp2p Workspace   使用libp2p的工作空间
About gx          关于gx
Tests             测试
Packages          包
Contribute        贡献


Background       背景

libp2p 是IPFS项目的一个模块化的网络堆栈和仓库，并单独捆绑供其他工具使用。
libp2p is a networking stack and library modularized out of The IPFS Project, and bundled separately for other tools to use.



libp2p2 是一个长期追求探索理解的产物，对互联网网络栈深入的研究，过去许多点对点协议。再过去的15年里，构建了一个大规模的点对点系统是非常负责和困难的，libp2p2是解决这个问题的一种方法，它是一个网络栈， 一个协议套件。它可以清晰分离关注点，并且允许复杂的应用仅仅使用这个协议，而不会放弃互操性和可升级性。libp2p起源于 IPFS，但是它的构建能让更多人，在不同的项目中使用它。

libp2p is the product of a long, and arduous quest of understanding -- a deep dive into the internet's network stack, and plentiful peer-to-peer protocols from the past. Building large-scale peer-to-peer systems has been complex and difficult in the last 15 years, and libp2p is a way to fix that. It is a "network stack" -- a protocol suite -- that cleanly separates concerns, and enables sophisticated applications to only use the protocols they absolutely need, without giving up interoperability and upgradeability. libp2p grew out of IPFS, but it is built so that lots of people can use it, for lots of different projects.

我们会编写一些文档，邮件，教程，告诉并且解释什么是 p2p，为什么它非常的有用，以及它如何帮助你现有的项目和新的项目，但与此同时，还需检验。
We will be writing a set of docs, posts, tutorials, and talks to explain what p2p is, why it is tremendously useful, and how it can help your existing and new projects. But in the meantime, check out

我们开发的文档集合
Our developing collection of docs
我们社区论坛
Our community discussion forums

libp2p的规范
The libp2p Specification

go-libp2p的实现
go-libp2p implementation

js-libp2p的实现
js-libp2p implementation

rust-libp2p的实现
rust-libp2p implementation



Usage  用法

这个仓库服务（go-libp2p）作为进入这个领域的模块，go的libp2p安装装置，Libp2p  要求 go 的版本在 1.12 以上
This repository (go-libp2p) serves as the entrypoint to the universe of modules that compose the Go implementation of the libp2p stack. Libp2p requires go 1.12+.

我们主要使用Go模块进行依赖和发布管理（这个要求go的版本在 1.12以上）。为了获得最好的开发人员体验，我们建议您也这样做。否则，您可能会偶尔遇到一个中断的构建，因为您将运行master(根据定义，它不能保证是稳定的)。
We mainly use Go modules for our dependency and release management (and thus require go >= 1.12+). In order to get the best developer experience, we recommend you do too. Otherwise, you may ocassionally encounter a breaking build as you'll be running off master (which, by definition, is not guaranteed to be stable).

你可以在Go应用中使用  go-libp2p，只需从我们的repo 中 添加导入，例如:
You can start using go-libp2p in your Go application simply by adding imports from our repos, e.g.:

import "github.com/libp2p/go-libp2p"

大概意思 就是 导入 libp2p2的模块


运行go get 或者go build 排除go模块中 使用代理，你仅需要第一次导入 go-libp2p2 以确保你能锁定当前的版本。
Run go get or go build, excluding the libp2p repos from Go modules proxy usage. You only need to do this the first time you import go-libp2p to make sure you latch onto the correct version lineage (see golang/go#34189 for context):

$ GOPRIVATE='github.com/libp2p/*' go get ./...


go构建工具会查看可用的版本，会挑选高可用，版本最高的那个
The Go build tools will look for available releases, and will pick the highest available one.

作为go-libp2p2 可用新的模块，  你能 手动 修改  go.mod 文件 来 升级你的 应用。或者使用go 工具去 维护这个模块要求。
As new releases of go-libp2p are made available, you can upgrade your application by manually editing your go.mod file, or using the Go tools to maintain module requirements.



API                    api
GoDoc                  Go文档


案例
Examples
案例可以在 examples 文件夹被发现 
Examples can be found in the examples folder.


发展
Development

使用 go-libp2p 工作空间
Using the go-libp2p Workspace

当正在发展中，你可能需要多次修改几个模块， 或者，您可能希望在一个模块中本地所做的更改可以被另一个模块导入。
While developing, you may need to make changes to several modules at once, or you may want changes made locally in one module to be available for import by another.

go-libp2p工作区提供了包含go-libp2p的模块的面向开发人员的视图。
The go-libp2p workspace provides a developer-oriented view of the modules that comprise go-libp2p.


使用这个工具，在这个工作区仓库。
Using the tooling in the workspace repository, you can checkout all of go-libp2p's module repos and enter "local mode", which adds replace directives to the go.mod files in each local working copy. When you build locally, the libp2p depdendencies will be resolved from your local working copies.

Once you've committed your changes, you can switch back to "remote mode", which removes the replace directives and pulls imports from the main go module cache.

See the workspace repo for more information.

About gx
Before adopting gomod, libp2p used gx to manage dependencies using IPFS.

Due to the difficulties in keeping both dependency management solutions up-to-date, gx support was ended in April 2019.

Ending gx support does not mean that existing gx builds will break. Because gx references dependencies by their immutable IPFS hash, any currently working gx builds will continue to work for as long as the dependencies are resolvable in IPFS.

However, new changes to go-libp2p will not be published via gx, and users are encouraged to adopt gomod to stay up-to-date.

If you experience any issues migrating from gx to gomod, please join the discussion at the libp2p forums.




Contribute   （贡献）
go-libp2p 是 IPFS 项目重要的一部分，并且 是 MIT 开源的软件。我们 欢迎 大大 小小 的 贡献。 看看社区贡献文档，请务必检查问题。在报告这个事情之后，请关闭查找搜索。 帮助我们处理打开的问题。
go-libp2p is part of The IPFS Project, and is MIT-licensed open source software. We welcome contributions big and small! Take a look at the community contributing notes. Please make sure to check the issues. Search the closed ones before reporting things, and help us with the open ones.

指南：
Guidelines:

阅读 libp2p2 规范
read the libp2p spec

请使用  这个分支  pull       即使在主 仓库。
please make branches + pull-request, even if working on the main repository

在我们的论坛，讨论中 ， 问问题 ，或者 谈论一些事情。 
ask questions or talk about things in issues, our discussion forums, or #libp2p or #ipfs on freenode.
ensure you are able to contribute (no legal issues please -- we use the DCO)

运行 go fmt  在 推送任何代码之前
run go fmt before pushing any code

run golint and go vet too -- some things (like protobuf files) are expected to fail.
get in touch with @raulk and @mgoelzer about how best to contribute
have fun!
There's a few things you can do right now to help out:

Go through the modules below and check out existing issues. This would be especially useful for modules in active development. Some knowledge of IPFS/libp2p may be required, as well as the infrasture behind it - for instance, you may need to read up on p2p and more complex operations like muxing to be able to help technically.
Perform code reviews.
Add tests. There can never be enough tests.