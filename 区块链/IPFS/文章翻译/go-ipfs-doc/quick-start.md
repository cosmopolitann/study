## Quick-start

```
# 0.1 - Quick Start

This is a set of short examples with minimal explanation. It is meant as
a "quick start".


Add a file to ipfs:

  echo "hello world" >hello
  ipfs add hello


View it:

  ipfs cat <the-hash-you-got-here>


Try a directory:

  mkdir foo
  mkdir foo/bar
  echo "baz" > foo/baz
  echo "baz" > foo/bar/baz
  ipfs add -r foo


View things:

  ipfs ls <the-hash-here>
  ipfs ls <the-hash-here>/bar
  ipfs cat <the-hash-here>/baz
  ipfs cat <the-hash-here>/bar/baz
  ipfs cat <the-hash-here>/bar
  ipfs ls <the-hash-here>/baz


References:

  ipfs refs <the-hash-here>
  ipfs refs -r <the-hash-here>
  ipfs refs --help


Get:

  ipfs get <the-hash-here> -o foo2
  diff foo foo2


Objects:

  ipfs object get <the-hash-here>
  ipfs object get <the-hash-here>/foo2
  ipfs object --help


Pin + GC:

  ipfs pin add <the-hash-here>
  ipfs repo gc
  ipfs ls <the-hash-here>
  ipfs pin rm <the-hash-here>
  ipfs repo gc


Daemon:

  ipfs daemon  (in another terminal)
  ipfs id


Network:

  (must be online)
  ipfs swarm peers
  ipfs id
  ipfs cat <hash-of-remote-object>


Mount:

  (warning: fuse is finicky!)
  ipfs mount
  cd /ipfs/<the-hash-here>
  ls


Tool:

  ipfs version
  ipfs update
  ipfs commands
  ipfs config --help
  open http://localhost:5001/webui


Browse:

  WebUI:

    http://localhost:5001/webui

  video:

    http://localhost:8080/ipfs/QmVc6zuAneKJzicnJpfrqCH9gSy6bz54JhcypfJYhGUFQu/play#/ipfs/QmTKZgRNwDNZwHtJSjCp6r5FYefzpULfy37JvMt9DwvXse

  images:

    http://localhost:8080/ipfs/QmZpc3HvfjEXvLWGQPWbHk3AjD5j8NEN4gmFN8Jmrd5g83/cs

  markdown renderer app:

    http://localhost:8080/ipfs/QmX7M9CiYXjVeFnkfVGf3y5ixTZ2ACeSGyL1vBJY1HvQPp/mdown

```



## 中文

```
# 0.1 -快速入门

这是一组简短的示例，意思是一个“快速启动”。


添加一个文件到ipfs:

Echo "hello world" >hello
ipfs 添加 这个文件


视图:

ipfs cat < the-hash-you-got-here >


试试一个目录:

mkdir foo
mkdir foo / bar
Echo "baz" > foo/baz
Echo "baz" > foo/bar/baz
Ipfs add -r foo


查看:

ipfs ls < the-hash-here >
ipfs ls < the-hash-here > /bar
ipfs cat< the-hash-here > /baz
ipfs cat< the-hash-here > /bar/baz
ipfs cat< the-hash-here > /bar
ipfs ls < the-hash-here > /baz


参考:

ipfs refs < the-hash-here >
ipfs refs -r <the-hash-here>
ipfs refs --help


得到:

ipfs get <the-hash-here> -o foo2
diff foo foo2


对象:

Ipfs object get <the-hash-here>
Ipfs object get <the-hash-here>/foo2
ipfs object --help


gc:

Ipfs pin add <the-hash-here>
ipfs repo gc
ipfs ls < the-hash-here >
Ipfs pin rm <the-hash-here>
ipfs  repo gc


守护进程:

Ipfs daemon(在另一个终端中)
ipfs id


网络:

(必须是在线)
ipfs swarm peers
ipfs id
ipfs cat < hash-of-remote-object >


挂载:

(警告:fuse 是 危险的!)
ipfs mount
cd m/ipfs /< the-hash-here >
ls


工具:

ipfs版本
ipfs更新
ipfs命令
ipfs config --help
打开http://localhost: 5001 / webui


浏览:

WebUI:

http://localhost:5001/webui

视频:

http://localhost:8080/ipfs/QmVc6zuAneKJzicnJpfrqCH9gSy6bz54JhcypfJYhGUFQu/play#/ipfs/QmTKZgRNwDNZwHtJSjCp6r5FYefzpULfy37JvMt9DwvXse

图片:

http://localhost:8080/ipfs/QmZpc3HvfjEXvLWGQPWbHk3AjD5j8NEN4gmFN8Jmrd5g83/cs

渲染器的应用:

http://localhost:8080/ipfs/QmX7M9CiYXjVeFnkfVGf3y5ixTZ2ACeSGyL1vBJY1HvQPp/mdown
```

