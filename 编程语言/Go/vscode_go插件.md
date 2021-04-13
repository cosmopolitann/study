# vscode安装go插件

### https://blog.csdn.net/lisulong1/article/details/86688514





注意：以下方式可能不能够安装go install github.com/golang/lint/golint

安装方式：D:\study_other\go\go_path\src\github.com\golang\lint目录下执行git clone https://github.com/golang/lint.git

然后将lint目录复制到D:\study_other\go\go_path\src\github.com\golang\和D:\study_other\go\go_path\src\golang.org\x\

然后执行：go install github.com/golang/lint/golint

注意：该演示环境是windows环境，linux和mac环境操作思路一样

vscode中有很多go的相关插件，非常好用如下： gocode gopkgs go-outline go-symbols guru gorename gomodifytags goplay impl godef goreturns golint gotests dlv

但是由于各种原因，这些插件无法安装，甚至你FQ之后发现也还是无法安装，加上最近FQ被限制的这么严格，所以总结了如下方法，让你在不FQ的情况下还能将这些插件安装成功，下图是我直接通过vscode安装提示的错误：



貌似运气还不错，还安装成功了几个，但是大部分还是没有安装成功，下面是vscode详细的安装日志：

Installing 14 tools at D:\go_project\bin
  gocode
  gopkgs
  go-outline
  go-symbols
  guru
  gorename
  gomodifytags
  goplay
  impl
  godef
  goreturns
  golint
  gotests
  dlv

Installing github.com/nsf/gocode SUCCEEDED
Installing github.com/uudashr/gopkgs/cmd/gopkgs SUCCEEDED
Installing github.com/ramya-rao-a/go-outline FAILED
Installing github.com/acroca/go-symbols FAILED
Installing golang.org/x/tools/cmd/guru FAILED
Installing golang.org/x/tools/cmd/gorename FAILED
Installing github.com/fatih/gomodifytags SUCCEEDED
Installing github.com/haya14busa/goplay/cmd/goplay SUCCEEDED
Installing github.com/josharian/impl FAILED
Installing github.com/rogpeppe/godef SUCCEEDED
Installing sourcegraph.com/sqs/goreturns FAILED
Installing github.com/golang/lint/golint FAILED
Installing github.com/cweill/gotests/... FAILED
Installing github.com/derekparker/delve/cmd/dlv SUCCEEDED

8 tools failed to install.

go-outline:
Error: Command failed: D:\Go\bin\go.exe get -u -v github.com/ramya-rao-a/go-outline
github.com/ramya-rao-a/go-outline (download)
Fetching https://golang.org/x/tools/go/buildutil?go-get=1
https fetch failed: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/go/buildutil: unrecognized import path "golang.org/x/tools/go/buildutil" (https fetch: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
github.com/ramya-rao-a/go-outline (download)
Fetching https://golang.org/x/tools/go/buildutil?go-get=1
https fetch failed: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/go/buildutil: unrecognized import path "golang.org/x/tools/go/buildutil" (https fetch: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)

go-symbols:
Error: Command failed: D:\Go\bin\go.exe get -u -v github.com/acroca/go-symbols
github.com/acroca/go-symbols (download)
Fetching https://golang.org/x/tools/go/buildutil?go-get=1
https fetch failed: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/go/buildutil: unrecognized import path "golang.org/x/tools/go/buildutil" (https fetch: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
github.com/acroca/go-symbols (download)
Fetching https://golang.org/x/tools/go/buildutil?go-get=1
https fetch failed: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/go/buildutil: unrecognized import path "golang.org/x/tools/go/buildutil" (https fetch: Get https://golang.org/x/tools/go/buildutil?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)

guru:
Error: Command failed: D:\Go\bin\go.exe get -u -v golang.org/x/tools/cmd/guru
Fetching https://golang.org/x/tools/cmd/guru?go-get=1
https fetch failed: Get https://golang.org/x/tools/cmd/guru?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/cmd/guru: unrecognized import path "golang.org/x/tools/cmd/guru" (https fetch: Get https://golang.org/x/tools/cmd/guru?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
Fetching https://golang.org/x/tools/cmd/guru?go-get=1
https fetch failed: Get https://golang.org/x/tools/cmd/guru?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/cmd/guru: unrecognized import path "golang.org/x/tools/cmd/guru" (https fetch: Get https://golang.org/x/tools/cmd/guru?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)

gorename:
Error: Command failed: D:\Go\bin\go.exe get -u -v golang.org/x/tools/cmd/gorename
Fetching https://golang.org/x/tools/cmd/gorename?go-get=1
https fetch failed: Get https://golang.org/x/tools/cmd/gorename?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/cmd/gorename: unrecognized import path "golang.org/x/tools/cmd/gorename" (https fetch: Get https://golang.org/x/tools/cmd/gorename?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
Fetching https://golang.org/x/tools/cmd/gorename?go-get=1
https fetch failed: Get https://golang.org/x/tools/cmd/gorename?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/cmd/gorename: unrecognized import path "golang.org/x/tools/cmd/gorename" (https fetch: Get https://golang.org/x/tools/cmd/gorename?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)

impl:
Error: Command failed: D:\Go\bin\go.exe get -u -v github.com/josharian/impl
github.com/josharian/impl (download)
Fetching https://golang.org/x/tools/imports?go-get=1
https fetch failed: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/imports: unrecognized import path "golang.org/x/tools/imports" (https fetch: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
github.com/josharian/impl (download)
Fetching https://golang.org/x/tools/imports?go-get=1
https fetch failed: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/imports: unrecognized import path "golang.org/x/tools/imports" (https fetch: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)

goreturns:
Error: Command failed: D:\Go\bin\go.exe get -u -v sourcegraph.com/sqs/goreturns
Fetching https://sourcegraph.com/sqs/goreturns?go-get=1
Parsing meta tags from https://sourcegraph.com/sqs/goreturns?go-get=1 (status code 200)
get "sourcegraph.com/sqs/goreturns": found meta tag get.metaImport{Prefix:"sourcegraph.com/sqs/goreturns", VCS:"git", RepoRoot:"https://github.com/sqs/goreturns"} at https://sourcegraph.com/sqs/goreturns?go-get=1
sourcegraph.com/sqs/goreturns (download)
github.com/sqs/goreturns (download)
Fetching https://golang.org/x/tools/imports?go-get=1
https fetch failed: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/imports: unrecognized import path "golang.org/x/tools/imports" (https fetch: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
Fetching https://sourcegraph.com/sqs/goreturns?go-get=1
Parsing meta tags from https://sourcegraph.com/sqs/goreturns?go-get=1 (status code 200)
get "sourcegraph.com/sqs/goreturns": found meta tag get.metaImport{Prefix:"sourcegraph.com/sqs/goreturns", VCS:"git", RepoRoot:"https://github.com/sqs/goreturns"} at https://sourcegraph.com/sqs/goreturns?go-get=1
sourcegraph.com/sqs/goreturns (download)
github.com/sqs/goreturns (download)
Fetching https://golang.org/x/tools/imports?go-get=1
https fetch failed: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/imports: unrecognized import path "golang.org/x/tools/imports" (https fetch: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)

golint:
Error: Command failed: D:\Go\bin\go.exe get -u -v github.com/golang/lint/golint
github.com/golang/lint (download)
Fetching https://golang.org/x/tools/go/gcexportdata?go-get=1
https fetch failed: Get https://golang.org/x/tools/go/gcexportdata?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/go/gcexportdata: unrecognized import path "golang.org/x/tools/go/gcexportdata" (https fetch: Get https://golang.org/x/tools/go/gcexportdata?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
github.com/golang/lint (download)
Fetching https://golang.org/x/tools/go/gcexportdata?go-get=1
https fetch failed: Get https://golang.org/x/tools/go/gcexportdata?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package golang.org/x/tools/go/gcexportdata: unrecognized import path "golang.org/x/tools/go/gcexportdata" (https fetch: Get https://golang.org/x/tools/go/gcexportdata?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)

gotests:
Error: Command failed: D:\Go\bin\go.exe get -u -v github.com/cweill/gotests/...
github.com/cweill/gotests (download)
Fetching https://golang.org/x/tools/imports?go-get=1
https fetch failed: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package github.com/cweill/gotests
    imports golang.org/x/tools/imports: unrecognized import path "golang.org/x/tools/imports" (https fetch: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
github.com/cweill/gotests (download)
Fetching https://golang.org/x/tools/imports?go-get=1
https fetch failed: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
package github.com/cweill/gotests
    imports golang.org/x/tools/imports: unrecognized import path "golang.org/x/tools/imports" (https fetch: Get https://golang.org/x/tools/imports?go-get=1: dial tcp 216.239.37.1:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.)
其实去src目录下看的话，是下载成功了，但是没有安装成功，并且我们也可以看出有几个是可以直接安装成功的 github.com/nsf/gocode github.com/tpng/gopkgs github.com/fatih/gomodifytags github.com/haya14busa/goplay github.com/rogpeppe/gode github.com/derekparker/delve/cmd/dlv

解决方法

关于go开发目录的结构这里不做过多解释，之前已经说过了

进行如下命令进行目录切换： cd %GOPATH%\src\github.com\golang 我这里的GOPATH是在D:\go_project 如果src目录下面没有github.com\golang请自行创建

完成目录切换后，开始下载插件包： git clone https://github.com/golang/tools.git tools

当下载完成后，你会发现%GOPATH%\src\github.com\golang多了一个tools目录 需要把tools目录下的所有文件拷贝到%GOPATH%\src\golang.org\x\tools下，如果没有自行创建 当然如果你是windows环境，如果你当前是在%GOPATH%\src\golang.org\x\tools 目录下，你可以直接使用如下命令进行拷贝： xcopy /s /e %GOPATH%\src\github.com\golang\tools 关于这个命令的使用可以具体百度查看，如果对该命令不熟悉就手动拷贝,直接将你下载的tools目录下的所有文件拷贝到%GOPATH%\src\golang.org\x\tools目录下

经过多次测试，插件中有几个其实不用FQ或其他方法就可以安装成功： github.com/nsf/gocode github.com/uudashr/gopkgs/cmd/gopkgs github.com/fatih/gomodifytags github.com/haya14busa/goplay/cmd/goplay github.com/derekparker/delve/cmd/dlv

下面安装无法安装的插件 开始安装： 切换到GOPATH目录下，执行相关的go install 命令

go install github.com/ramya-rao-a/go-outline

go install github.com/acroca/go-symbols

go install golang.org/x/tools/cmd/guru

go install golang.org/x/tools/cmd/gorename

go install github.com/josharian/impl

go install github.com/rogpeppe/godef

go install github.com/sqs/goreturns

go install github.com/golang/lint/golint

go install github.com/cweill/gotests/gotests

这样vscode下go开发需要安装的插件都已经安装成功。







```go

func CreateTableIndex() {
	//     abi         表
	//     user        表
	//  transaction_detail 交易信息表
	//  parameter_abi 匹配表

	//contract_bas9   BAS9合约表
	//contract        合约表

	//block_detail   区块表
	//user_asset     用户资产表
	//user_total     用户总数表
	// dailiy 表     每日增加表         

	CreateIndex()

}
```

